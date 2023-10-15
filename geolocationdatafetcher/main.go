package geolocationdatafetcher

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "github.com/tidwall/gjson"
    "net/http"
    "net/url"
    "strconv"
)

type Branch struct {
    Name      string  `json:"name"`
    Address   string  `json:"address"`
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Occupancy int     `json:"occupancy"`
    Distance  float64 `json:"distance"`
}

type BankAPIResponse struct {
    Branches []Branch `json:"items"`
}

func main() {
    // Создание маршрутизатора HTTP
    router := mux.NewRouter()

    // Запрос данных о ближайших отделениях банка ВТБ
    router.HandleFunc("/branches", getVTBBranches).Methods("GET")

    // Запрос локации отделения на карте
    router.HandleFunc("/branches/{id}/location", getBranchLocation).Methods("GET")

    // Запуск HTTP-сервера на порту 8080
    http.ListenAndServe(":8080", router)
}

func getVTBBranches(w http.ResponseWriter, r *http.Request) {
    // Получение координат пользователя для определения ближайшего отделения
    userLatStr := r.URL.Query().Get("lat")
    userLngStr := r.URL.Query().Get("lng")

    userLat, err := strconv.ParseFloat(userLatStr, 64)
    if err != nil {
        http.Error(w, "Invalid user latitude", http.StatusBadRequest)
        return
    }

    userLng, err := strconv.ParseFloat(userLngStr, 64)
    if err != nil {
        http.Error(w, "Invalid user longitude", http.StatusBadRequest)
        return
    }

    // Формирование GET-запроса к API ВТБ
    requestURL := "https://www.vtb.ru/api/branchlocator/branchlist?CityID=1&ServiceID=1&HasCB=Moscow"
    resp, err := http.Get(requestURL)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Чтение данных из ответа
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Извлечение информации о отделениях из JSON-данных
    branchesJSON := gjson.GetBytes(body, "items")

    var branches []Branch
    err = json.Unmarshal([]byte(branchesJSON.Raw), &branches)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Фильтрация отделений по близости к пользователю и загруженности
    filteredBranches := filterBranches(branches, userLat, userLng)

    // Отправка отфильтрованных данных в формате JSON
    jsonResponse, err := json.Marshal(filteredBranches)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

func getBranchLocation(w http.ResponseWriter, r *http.Request) {
    // Извлечение идентификатора отделения из URL-параметра
    vars := mux.Vars(r)
    branchID := vars["id"]

    // Получение данных о локации отделения из API Google Maps
    location, err := getBranchLocationFromMapsAPI(branchID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Отправка данных о локации в формате JSON
    jsonResponse, err := json.Marshal(location)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

func filterBranches(branches []Branch, userLat, userLng float64) []Branch {
    var filteredBranches []Branch

    for _, branch := range branches {
        // Рассчет расстояния между отделением и пользователем
        distance := calculateDistance(userLat, userLng, branch.Latitude, branch.Longitude)
        branch.Distance = distance

        // Фильтрация отделений по близости к пользователю (например, 5 км)
        if distance <= 5 {
            filteredBranches = append(filteredBranches, branch)
        }
    }

    return filteredBranches
}

func calculateDistance(lat1, lng1, lat2, lng2 float64) float64 {
    const EarthRadius = 6371 // радиус Земли в километрах

    lat1Rad := degToRad(lat1)
    lng1Rad := degToRad(lng1)
    lat2Rad := degToRad(lat2)
    lng2Rad := degToRad(lng2)

    deltaLat := lat2Rad - lat1Rad
    deltaLng := lng2Rad - lng1Rad

    a := (math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
        math.Cos(lat1Rad)*math.Cos(lat2Rad)*
            math.Sin(deltaLng/2)*math.Sin(deltaLng/2))

    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

    distance := EarthRadius * c

    return distance
}

func degToRad(deg float64) float64 {
    return deg * (math.Pi / 180)
}

func getBranchLocationFromMapsAPI(branchID string) (map[string]interface{}, error) {
    // Формирование запроса к API Google Maps для получения локации отделения
    requestURL := "https://maps.googleapis.com/maps/api/geocode/json?address=" + url.QueryEscape(branchID)

    resp, err := http.Get(requestURL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    // Извлечение нужных данных из JSON-ответа
    location := gjson.GetBytes(body, "results.0.geometry.location").Value().(map[string]interface{})

    return location, nil
}
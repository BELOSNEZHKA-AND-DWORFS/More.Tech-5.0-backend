package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"more.tech/structs"
	"more.tech/util"
)

func RootTest(w http.ResponseWriter, r *http.Request) {
	// result := getOfficeByTitleInArea("втб", 55.649865, 37.622646)
	// fmt.Fprintf(w, "%v\n", result)
}

func readFile(path string) []byte {
	// file, err := os.Open(path)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// var lines []string
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// }

	// var text string
	// for _, line := range lines {
	// 	text += line
	// }

	// return text

	v, err := ioutil.ReadFile(path) //read the content of file
	if err != nil {
		fmt.Println(err)
	}
	return v
}

func GetOfficeInfo(w http.ResponseWriter, r *http.Request) {
	text := readFile("file.txt")
	var objects []structs.Object
	json.Unmarshal([]byte(text), &objects)

	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var requestObject structs.GetObjectsBody
	err = json.Unmarshal(b, &requestObject)

	var result []structs.Object
	for _, obj := range objects {
		if util.Distance(obj.Location, requestObject.Location) > float64(requestObject.Radius) {
			continue
		}

		if !util.CheckFilter(obj.Services, requestObject.Filter) {
			continue
		}

		if len(obj.Services) == 0 {
			result = append(result, obj)
			continue
		}

		if obj.ServiceType != "office" {
			continue
		}

		result = append(result, obj)
	}

	jsonResp, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(jsonResp))
}

func GetOffices(w http.ResponseWriter, r *http.Request) {
	text := readFile("file.txt")
	var objects []structs.Object
	err := json.Unmarshal([]byte(text), &objects)

	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var requestObject structs.GetObjectsBody
	err = json.Unmarshal(b, &requestObject)

	var result []structs.Object
	for _, obj := range objects {
		if util.Distance(obj.Location, requestObject.Location) > float64(requestObject.Radius) {
			continue
		}

		if !util.CheckFilter(obj.Services, requestObject.Filter) {
			continue
		}

		if len(requestObject.Filter) == 0 {
			result = append(result, obj)
			continue
		}

		if !util.CheckOneFilter(obj.Services, "office") {
			continue
		}

		result = append(result, obj)
	}

	jsonResp, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(jsonResp))
}

func GetAtmInfo(w http.ResponseWriter, r *http.Request) {
	text := readFile("file.txt")
	var objects []structs.Object
	json.Unmarshal([]byte(text), &objects)

	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var requestObject structs.GetObjectsBody
	err = json.Unmarshal(b, &requestObject)

	id := r.URL.Query().Get("atmid")
	for _, obj := range objects {
		if id_i, _ := strconv.Atoi(id); id_i == obj.Id {
			jsonResp, err := json.Marshal(obj)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprintf(w, string(jsonResp))
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func GetAtms(w http.ResponseWriter, r *http.Request) {
	text := readFile("file.txt")
	var objects []structs.Object
	json.Unmarshal([]byte(text), &objects)

	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var requestObject structs.GetObjectsBody
	json.Unmarshal(b, &requestObject)

	var result []structs.Object
	for _, obj := range objects {
		if util.Distance(obj.Location, requestObject.Location) > float64(requestObject.Radius) {
			continue
		}

		if len(requestObject.Filter) == 0 {
			result = append(result, obj)
			continue
		}

		if !util.CheckFilter(obj.Services, requestObject.Filter) {
			continue
		}

		if obj.ServiceType != "atm" {
			continue
		}

		result = append(result, obj)
	}

	jsonResp, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(jsonResp))
}

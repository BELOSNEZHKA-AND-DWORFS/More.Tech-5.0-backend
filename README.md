# More.Tech-5.0 backend

## main package

Для реализации указанных уточнений вам потребуются следующие дополнительные пакеты:

1) "github.com/tidwall/gjson" — для работы с JSON-данными и извлечения информации из них.
2) "github.com/gorilla/mux" — для создания HTTP-маршрутов и обработки запросов.
3) "googlemaps.github.io/maps" — для взаимодействия с API Google Maps.

Ниже представлен доработанный пример программы, учитывающий указанные уточнения.

В этом примере создается HTTP-сервер, обрабатывающий два типа запросов. Первый запрос (/branches) возвращает список ближайших отделений банка ВТБ, учитывая координаты пользователя. Второй запрос (/branches/{id}/location) возвращает локацию выбранного отделения на карте, используя API Google Maps.

Программа также включает функцию filterBranches, которая фильтрует отделения по близости к пользователю (например, в пределах 5 км), а также функции calculateDistance и degToRad для расчета расстояния между координатами на основе формулы гаверсинусов.

Обратите внимание, что вам необходимо подставить свой API-ключ Google Maps в запросе к API Google Maps (в функции getBranchLocationFromMapsAPI).

Также следует учесть, что данная программа является только примером и может потребоваться дополнительная доработка и обработка ошибок для использования в реальном проекте.
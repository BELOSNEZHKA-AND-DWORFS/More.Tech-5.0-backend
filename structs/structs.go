package structs

type Location struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Object struct {
	Id              int       `json:"id"`
	ServiceType     string    `json:"type"`
	Services        []string  `json:"services"`
	Location        []float64 `json:"location"`
	Working_time_fl string    `json:"working_time_fl"`
	Working_time_ul string    `json:"working_time_ul"`
}

type GetObjectsBody struct {
	Location  []float64 `json:"location"`
	Radius    float64   `json:"radius"`
	Filter    []string  `json:"filter"`
	Sorted_by []string  `json:"sorted_by"`
}

type YandexBasicOrganisationData struct {
	Title    string
	Phone    string
	Distance float32
	Datetime string
}

type OfficeYandexData struct {
	YandexBasicOrganisationData
}

type AtmYandexData struct {
	YandexBasicOrganisationData
}

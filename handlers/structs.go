package handlers

type Location struct {
	Latitude  float32
	Longitude float32
}

type yandexBasicOrganisationData struct {
	title    string
	phone    string
	distance float32
	datetime string
}

type officeYandexData struct {
	yandexBasicOrganisationData
}

type atmYandexData struct {
	yandexBasicOrganisationData
}

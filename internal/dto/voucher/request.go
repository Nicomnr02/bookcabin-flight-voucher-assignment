package voucherdto

type CheckRequest struct {
	FlightNumber string `json:"flightNumber"`
	Date         string `json:"date"`
}

type GenerateRequest struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	FlightNumber string `json:"flightNumber"`
	Date         string `json:"date"`
	Aircraft     string `json:"aircraft"`
}

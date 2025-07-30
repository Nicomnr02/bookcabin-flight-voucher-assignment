package voucherdto

type CheckResponse struct {
	Exists bool `json:"exists"`
}

type GenerateResponse struct {
	Seats []string `json:"seats"`
}

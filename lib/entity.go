package lib

type Person struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	HitPoints    int64  `json:"hitPoints"`
	Strength     int64  `json:"strength"`
	Defense      int64  `json:"defense"`
	Intelligence int64  `json:"intelligence"`
	Class        int64  `json:"class"`
}

type ParticipantResponse struct {
	Data    []Person `json:"data"`
	Success bool     `json:"success"`
	Message *string  `json:"message"`
}

type GeneralResponse struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Message *string     `json:"message"`
}

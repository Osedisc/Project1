package createUser

type Request struct {
	Name        string `json:"name"`
	SurName     string `json:"surName"`
	DateOfBirth string `json:"dateOfBirth"`
}

type Response struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	SurName     string `json:"surName"`
	DateOfBirth string `json:"dateOfBirth"`
}

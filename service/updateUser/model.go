package updateUser

type Request struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	SurName     string `json:"surName"`
	DateOfBirth string `json:"dateOfBirth"`
}

type Response struct {
	Id string `json:"id"`
}

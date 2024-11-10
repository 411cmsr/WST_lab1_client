package models

type AddPersonResponse struct {
	Content struct {
		ID int `xml:"id"`
	} `xml:"Body"`
}

type UpdatePersonResponse struct {
	Success bool `xml:"success"`
}

type DeletePersonResponse struct {
	Success bool `xml:"success"`
}

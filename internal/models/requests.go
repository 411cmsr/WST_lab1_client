package models

import "encoding/xml"

type AddPersonRequest struct {
	XMLName xml.Name `xml:"AddPerson"`
	Person  Person   `xml:"person"`
}

type GetPersonRequest struct {
	XMLName xml.Name `xml:"GetPerson"`
	ID      int      `xml:"id"`
}

type UpdatePersonRequest struct {
	XMLName xml.Name `xml:"UpdatePerson"`
	Person  Person   `xml:"person"`
}

type DeletePersonRequest struct {
	ID int `xml:"id"`
}

type SearchPersonRequest struct {
	XMLName xml.Name `xml:"SearchPerson"`
	Query   string   `xml:"query"`
}

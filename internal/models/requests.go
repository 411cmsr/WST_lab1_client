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
	XMLName xml.Name `xml:"DeletePerson"`
	ID      int      `xml:"id"`
}

type SearchPersonRequest struct {
	XMLName xml.Name `xml:"SearchPerson"`
	Query   string   `xml:"query"`
}

type Header struct {
}

type FaultType struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
}
type ContentType struct {
	Persons []Person `xml:"persons"`
}

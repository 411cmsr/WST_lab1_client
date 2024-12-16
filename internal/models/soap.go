package models

import (
	"encoding/xml"
)

//	type Fault struct {
//		FaultString string `xml:"faultstring"`
//	}
//
//	type Content struct {
//		Persons []Person `xml:"person"`
//	}
//
//	type Body struct {
//		Fault   *Fault   `xml:"Fault"`
//		Content *Content `xml:"Content"`
//	}
//
//	type Envelope struct {
//		Body Body `xml:"Body"`
//	}
type Envelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Header  Header   `xml:"Header"`
	Body    Body     `xml:"Body"`
}

type Header struct{}

// Body представляет тело SOAP сообщения
type Body struct {
	AddPerson     *AddPersonRequest     `xml:"AddPerson,omitempty"`
	DeletePerson  *DeletePersonRequest  `xml:"DeletePerson,omitempty"`
	UpdatePerson  *UpdatePersonRequest  `xml:"UpdatePerson,omitempty"`
	GetPerson     *GetPersonRequest     `xml:"GetPerson,omitempty"`
	GetAllPersons *GetAllPersonsRequest `xml:"GetAllPersons,omitempty"`
	SearchPerson  *SearchPersonRequest  `xml:"SearchPerson,omitempty"`

	AddPersonResponse     *AddPersonResponse     `xml:"AddPersonResponse,omitempty"`
	DeletePersonResponse  *DeletePersonResponse  `xml:"DeletePersonResponse,omitempty"`
	UpdatePersonResponse  *UpdatePersonResponse  `xml:"UpdatePersonResponse,omitempty"`
	GetPersonResponse     *GetPersonResponse     `xml:"GetPersonResponse,omitempty"`
	GetAllPersonsResponse *GetAllPersonsResponse `xml:"GetAllPersonsResponse,omitempty"`
	SearchPersonResponse  *SearchPersonResponse  `xml:"SearchPersonResponse,omitempty"`

	Fault   *Fault  `xml:"Fault,omitempty"`
	Content Content `xml:",any"`
}

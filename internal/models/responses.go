package models

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Body    Body     `xml:"soap:Body"`
}

type Body struct {
	Fault   *Fault      `xml:"Fault,omitempty"`
	Content interface{} `xml:",any"` 
}
type SOAPFault struct {
	XMLName xml.Name `xml:"SOAPFault"`
	Envelope struct {
		Body struct {
			Fault Fault `xml:"Fault"`
		 } `xml:"Body"`
			 
   	}`xml:"Envelope"`
}

type Fault struct {
	FaultCode string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	FaultDetail FaultDetail `xml:"detail"`
}
type FaultDetail struct {
	ErrorCode string `xml:"errorCode"`
	ErrorMessage string `xml:"errorMessage"`
}

type DeleteResponse struct {
	Status bool `xml:"status"`
}

type SearchPersonResponse struct {
	Persons []Person 
}

type AddPersonResponse struct {
	ID int `xml:"ID"`
}

// GetPersonResponse представляет ответ на запрос получения информации о человеке
type GetPersonResponse struct {
	Content Person `xml:"Content"` // Предполагается, что Content содержит информацию о человеке
}

// GetAllPersonsResponse представляет ответ на запрос получения всех людей
type GetAllPersonsResponse struct {
	Content PersonsList `xml:"Content"` // Предполагается, что Content содержит список людей
}

// UpdatePersonResponse представляет ответ на запрос обновления информации о человеке
type UpdatePersonResponse struct {
	Status bool `xml:"status"`
}

type GetPersonResponse struct {
	Person Person `xml:"Person"`
}

type GetAllPersonsResponse struct {
	
	Persons []Person `xml:"persons"`
}

type ErrorResponse struct {
	
	Envelope struct {} `xml:"Envelope"`

}

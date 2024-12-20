package models

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Body    Body     `xml:"soap:Body"`
}

type Body struct {
	Fault   *Fault      `xml:"Fault,omitempty"`
	Content interface{} `xml:",any"` // This allows for any content
}

type Fault struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	FaultDetail string `xml:"detail"`

}
type FaultDetail struct {
	ErrorCode string `xml:"errorCode"`
	ErrorMessage string `xml:"errorMessage"`
}

type DeleteResponse struct {
	Status bool `xml:"status"`
}

type SearchPersonResponse struct {
	Persons []Person `xml:"Persons>Person"`
}

type AddPersonResponse struct {
	ID int `xml:"ID"`
}

type UpdatePersonResponse struct {
	Status bool `xml:"status"`
}

type GetPersonResponse struct {
	Person Person `xml:"Person"`
}

type GetAllPersonsResponse struct {
	//Persons []Person `xml:"Persons>Person"`
	Persons []Person `xml:"persons"`
}

type ErrorResponse struct {
	
	Envelope struct {} `xml:"Envelope"`

}
// type SOAPFault struct {

// }




// type ErrorResponse struct {
// 	Type     string `xml:"type"`
// 	Title    string `xml:"title"`
// 	Status   int    `xml:"status"`
// 	Detail   string `xml:"detail"`
	
// }

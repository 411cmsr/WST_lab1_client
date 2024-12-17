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

// ////////////////////////////////////////

// Envelope представляет корневой элемент SOAP
// type Envelope struct {
// 	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
// 	Header  Header   `xml:"Header"`
// 	Body    Body     `xml:"Body"`
// }

// Header представляет заголовок SOAP
type Header struct {
	// Здесь можно добавить поля, если они нужны
}

// Body представляет тело SOAP сообщения
// type Body struct {
// 	//Content interface{} `xml:",any"`
// 	Content ContentType `xml:"content"`
// 	Fault   *FaultType    `xml:"fault,omitempty"`
// 	 // Используем интерфейс для динамического содержимого
// }
type FaultType struct {
    FaultCode   string `xml:"faultcode"`
    FaultString string `xml:"faultstring"`
}
type ContentType struct {
    Persons []Person `xml:"persons"`
}
package models

//type Envelope struct {
//	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
//	Header  Header   `xml:"Header"`
//	Body    Body    `xml:"Body"`
//}

// Header представляет заголовок SOAP

// Person представляет информацию о человеке

// Запросы
type AddPersonRequest struct {
	Person Person `xml:"Person"` // Убедитесь, что структура соответствует формату
}

type DeletePersonRequest struct {
	XMLName xml.Name `xml:"DeletePerson"`
	ID      int      `xml:"id"`
}

type UpdatePersonRequest struct {
	Person Person `xml:"Person"` // Убедитесь, что структура соответствует формату
}

type GetPersonRequest struct {
	ID uint `xml:"ID"`
}

type GetAllPersonsRequest struct{}

type SearchPersonRequest struct {
	Query string `xml:"query"`
}

// Fault представляет ошибку в SOAP сообщении
type Fault struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
}

type Content struct {
	Persons []Person `xml:"Persons>Person"` // Предполагается, что это структура для списка людей
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
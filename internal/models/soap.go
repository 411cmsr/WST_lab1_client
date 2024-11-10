package models

type Fault struct {
	FaultString string `xml:"faultstring"`
}

type Content struct {
	Message string   `xml:"chardata"`
	Persons []Person `xml:"person"`
}

type Body struct {
	Fault   *Fault   `xml:"Fault"`
	Content *Content `xml:"Content"`
}

type Envelope struct {
	Body Body `xml:"Body"`
}

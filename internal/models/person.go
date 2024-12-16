package models

type Person struct {
	ID        int    `xml:"ID"`
	Name      string `xml:"Name"`
	Surname   string `xml:"Surname"`
	Age       int    `xml:"Age"`
	Email     string `xml:"Email"`
	Telephone string `xml:"Telephone"`
}

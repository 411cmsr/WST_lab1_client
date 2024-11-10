package models

type Person struct {
	ID      int    `xml:"id"`
	Name    string `xml:"name"`
	Surname string `xml:"surname"`
	Age     int    `xml:"age"`
}

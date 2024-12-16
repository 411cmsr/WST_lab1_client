package models

//type AddPersonResponse struct {
//	Content struct {
//		ID uint `xml:"id"`
//	} `xml:"Body"`
//}
//
//type UpdatePersonResponse struct {
//	Success bool `xml:"success"`
//}
//
//type DeletePersonResponse struct {
//	Success bool `xml:"success"`
//}

// Envelope представляет корневой элемент SOAP

// Header представляет заголовок SOAP

// Body представляет тело SOAP сообщения
//type Body struct {
//	AddPersonResponse     *AddPersonResponse     `xml:"AddPersonResponse,omitempty"`
//	DeletePersonResponse  *DeletePersonResponse  `xml:"DeletePersonResponse,omitempty"`
//	UpdatePersonResponse  *UpdatePersonResponse  `xml:"UpdatePersonResponse,omitempty"`
//	GetPersonResponse     *GetPersonResponse     `xml:"GetPersonResponse,omitempty"`
//	GetAllPersonsResponse *GetAllPersonsResponse `xml:"GetAllPersonsResponse,omitempty"`
//	SearchPersonResponse  *SearchPersonResponse `xml:"SearchPersonResponse,omitempty"`
//	Fault                 *Fault                `xml:"Fault,omitempty"`
//}

// Fault представляет ошибку в SOAP сообщении

// AddPersonResponse представляет ответ на запрос добавления человека
type AddPersonResponse struct {
	Content PersonID `xml:"Content"` // Предполагается, что Content содержит ID добавленного человека
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
	Success bool `xml:"Success"` // Успех операции обновления
}

// DeletePersonResponse представляет ответ на запрос удаления человека
type DeletePersonResponse struct {
	Success bool `xml:"Success"` // Успех операции удаления
}

// SearchPersonResponse представляет ответ на запрос поиска людей
type SearchPersonResponse struct {
	Content PersonsList `xml:"Content"`
}

// Person представляет информацию о человеке

// PersonsList представляет список людей
type PersonsList struct {
	Persons []Person `xml:"Persons>Person"`
}

// PersonID представляет ID добавленного человека
type PersonID struct {
	ID uint `xml:"ID"` // ID добавленного человека
}

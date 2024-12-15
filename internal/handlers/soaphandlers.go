package handlers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"

	"WST_lab1_client/internal/models"
)

func AddPersonHandler(url string, name string, surname string, age int, email string, telephone string, logger *zap.Logger) {
	person := models.Person{Name: name, Surname: surname, Age: age, Email: email, Telephone: telephone}
	request := models.AddPersonRequest{Person: person}
	requestXML, err := xml.Marshal(request)
	if err != nil {
		logger.Fatal("Error marshaling request", zap.Error(err))
	}

	body, err := sendSOAPRequest(url, requestXML, logger)
	if err != nil {
		logger.Warn("Error calling AddPerson", zap.Error(err))
		return
	}

	var response models.AddPersonResponse
	if err := xml.Unmarshal(body, &response); err != nil {
		logger.Fatal("Error unmarshalling response", zap.Error(err))
	}

	fmt.Printf("Added person with ID: %d\n", response.Content.ID)
}

func GetPersonHandler(url string, id int, logger *zap.Logger) {
	requestXML := []byte(fmt.Sprintf(`<GetPerson><id>%d</id></GetPerson>`, id))

	body, err := sendSOAPRequest(url, requestXML, logger)
	if err != nil {
		logger.Warn("Error calling GetPerson", zap.Error(err))
		return
	}

	err = printFullResponse(body, logger)
	if err != nil {
		return
	}
}

func GetAllPersonsHandler(url string, logger *zap.Logger) {
	requestXML := []byte(`<GetAllPersons/>`)

	body, err := sendSOAPRequest(url, requestXML, logger)
	if err != nil {
		logger.Warn("Error calling GetAllPersons", zap.Error(err))
		return
	}

	err = printFullResponse(body, logger)
	if err != nil {
		return
	}
}

func UpdatePersonHandler(url string, id int, name string, surname string, age int, email string, telephone string, logger *zap.Logger) {
	person := models.Person{ID: id, Name: name, Surname: surname, Age: age, Email: email, Telephone: telephone}
	request := models.UpdatePersonRequest{Person: person}
	requestXML, err := xml.Marshal(request)
	if err != nil {
		logger.Fatal("Error marshaling request", zap.Error(err))
		return
	}

	body, err := sendSOAPRequest(url, requestXML, logger)
	if err != nil {
		logger.Warn("Error calling UpdatePerson", zap.Error(err))
		return
	}

	var response models.UpdatePersonResponse
	if err := xml.Unmarshal(body, &response); err != nil {
		logger.Fatal("Error unmarshalling response", zap.Error(err))
		return
	}

	fmt.Printf("Updated person successfully: %v\n", response.Success)
}

func DeletePersonHandler(url string, id int, logger *zap.Logger) {
	request := models.DeletePersonRequest{ID: id}
	requestXML, err := xml.Marshal(request)
	if err != nil {
		logger.Fatal("Error marshaling request", zap.Error(err))
		return
	}

	body, err := sendSOAPRequest(url, requestXML, logger)
	if err != nil {
		logger.Warn("Error calling DeletePerson", zap.Error(err))
		return
	}
	fmt.Printf(string(body))
	var response models.Envelope
	//var response models.DeletePersonResponse
	if err := xml.Unmarshal(body, &response); err != nil {
		logger.Fatal("Error unmarshalling response", zap.Error(err))
		return
	}
	fmt.Println("RRRRRRRRRRRRRRRRRR", response)
	//if response.Body.Fault != nil {
	//	logger.Warn("Received Fault", zap.String("faultstring", response.Body.Fault.FaultString))
	//	fmt.Println(response.Body.Fault.FaultString)
	//	return
	//}

	fmt.Printf("Status of deleting person: %v\n", response.Body.Content)
}

func SearchPersonsHandler(url string, query string, logger *zap.Logger) {
	request := models.SearchPersonRequest{Query: query}
	requestXML, err := xml.Marshal(request)
	if err != nil {
		logger.Fatal("Error marshaling request", zap.Error(err))
		return
	}

	body, err := sendSOAPRequest(url, requestXML, logger)
	if err != nil {
		logger.Warn("Error calling SearchPersons", zap.Error(err))
		return
	}

	err = printFullResponse(body, logger)

	if err != nil {
		return
	}
}

//func sendSOAPRequest(url string, request interface{}, logger *zap.Logger) ([]byte, error) {
//	//func sendSOAPRequest(url string, requestXML []byte, logger *zap.Logger) ([]byte, error) {
//	//soapEnvelope := fmt.Sprintf(`
//	//    <Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
//	//        <Header></Header>
//	//        <Body>
//	//            %s
//	//        </Body>
//	//    </Envelope>`, string(requestXML))
//
//	bodyContent, err := xml.Marshal(request)
//	if err != nil {
//		logger.Fatal("Error marshaling request", zap.Error(err))
//		return nil, err
//	}
//
//	envelope := models.Envelope{
//		Header: models.Header{},
//		Body: models.Body{
//			Content: bodyContent,
//		},
//	}
//
//	soapEnvelope, err := xml.Marshal(envelope)
//	if err != nil {
//		logger.Fatal("Error marshaling envelope", zap.Error(err))
//		return nil, err
//
//	}
//
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(soapEnvelope)))
//	if err != nil {
//		logger.Error("Error creating request", zap.Error(err))
//		return nil, err
//	}
//
//	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
//	req.Header.Set("SOAPAction", "Request")
//
//	logger.Info("Sending SOAP request", zap.String("url", url), zap.String("request", soapEnvelope))
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		logger.Error("Error sending request", zap.Error(err))
//		return nil, err
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//
//		}
//	}(resp.Body)
//
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		logger.Error("Error reading response body", zap.Error(err))
//		return nil, err
//	}
//
//	return body, nil
//}

// sendSOAPRequest отправляет SOAP запрос с использованием структуры запроса
func sendSOAPRequest(url string, request interface{}, logger *zap.Logger) ([]byte, error) {
	// Создаем тело SOAP сообщения
	bodyContent, err := xml.Marshal(request)
	if err != nil {
		logger.Error("Error marshaling request to XML", zap.Error(err))
		return nil, err
	}

	// Создаем Envelope
	envelope := models.Envelope{
		Header: models.Header{}, // Если заголовки не нужны, оставляем пустым
		Body: models.Body{
			Content: bodyContent,
		},
	}

	// Маршалим Envelope в XML
	soapEnvelope, err := xml.Marshal(envelope)
	if err != nil {
		logger.Error("Error marshaling envelope to XML", zap.Error(err))
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(soapEnvelope))
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "Request")

	logger.Info("Sending SOAP request", zap.String("url", url), zap.String("request", string(soapEnvelope)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error sending request", zap.Error(err))
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error("Error closing response body", zap.Error(err))
		}
	}(resp.Body)

	bodyResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body", zap.Error(err))
		return nil, err
	}

	return bodyResponse, nil
}

func printFullResponse(body []byte, logger *zap.Logger) error {
	//var response models.Envelope
	//err := xml.Unmarshal(body, &response)
	//if err != nil {
	//	logger.Fatal("Error unmarshalling response", zap.Error(err))
	//	return err
	//}
	//
	//if response.Body.Fault != nil {
	//	logger.Warn("Received Fault", zap.String("faultstring", response.Body.Fault.FaultString))
	//	fmt.Println(response.Body.Fault.FaultString)
	//	return nil
	//}
	//
	//if response.Body.Content == nil || len(response.Body.Content.Persons) == 0 {
	//	message := "Response is empty"
	//	logger.Info(message)
	//	fmt.Println(message)
	//	return nil
	//}
	//
	//fmt.Println("Result of request execution:")
	//for _, person := range response.Body.Content.Persons {
	//	fmt.Printf("- ID: %d, Name: %s, Surname: %s, Age: %d, Email: %s, Telephone: %s\n",
	//		person.ID,
	//		person.Name,
	//		person.Surname,
	//		person.Age,
	//		person.Email,
	//		person.Telephone,
	//	)
	//
	//	logger.Info("Result of request execution",
	//		zap.Int("ID", person.ID),
	//		zap.String("Name", person.Name),
	//		zap.String("Surname", person.Surname),
	//		zap.Int("Age", person.Age),
	//		zap.String("Email", person.Email),
	//		zap.String("Telephone", person.Telephone),
	//	)
	//}
	return nil
}

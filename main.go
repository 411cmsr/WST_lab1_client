package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.uber.org/zap"
)

type Person struct {
	ID      int    `xml:"id"`
	Name    string `xml:"name"`
	Surname string `xml:"surname"`
	Age     int    `xml:"age"`
}

type AddPersonRequest struct {
	XMLName xml.Name `xml:"AddPerson"`
	Person  Person   `xml:"person"`
}

type AddPersonResponse struct {
	Content struct {
		ID int `xml:"id"`
	} `xml:"Body"`
}

type GetPersonRequest struct {
	XMLName xml.Name `xml:"GetPerson"`
	ID      int      `xml:"id"`
}

type GetPersonResponse struct {
	Content Person `xml:"person"`
}

type UpdatePersonRequest struct {
	XMLName xml.Name `xml:"UpdatePerson"`
	Person  Person   `xml:"person"`
}

type UpdatePersonResponse struct {
	Success bool `xml:"success"`
}

type DeletePersonRequest struct {
	XMLName xml.Name `xml:"DeletePerson"`
	ID      int      `xml:"id"`
}

type DeletePersonResponse struct {
	Success bool `xml:"success"`
}

type SearchPersonRequest struct {
	XMLName xml.Name `xml:"SearchPerson"`
	Query   string   `xml:"query"`
}

type Fault struct {
	FaultString string `xml:"faultstring"`
}

type Content struct {
	Message string   `xml:",chardata"`
	Persons []Person `xml:"person"`
}

type Body struct {
	Fault   *Fault   `xml:"Fault"`
	Content *Content `xml:"Content"`
}

type Envelope struct {
	Body Body `xml:"Body"`
}

func sendSOAPRequest(url string, requestXML []byte, logger *zap.Logger) ([]byte, error) {
	soapEnvelope := fmt.Sprintf(`
        <Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
            <Header></Header>
            <Body>
                %s
            </Body>
        </Envelope>`, string(requestXML))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(soapEnvelope)))
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "Request")

	logger.Info("Sending SOAP request", zap.String("url", url), zap.String("request", soapEnvelope))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error sending request", zap.Error(err))
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body", zap.Error(err))
		return nil, err
	}

	return body, nil
}

func printFullResponse(body []byte, logger *zap.Logger) error {
	var response Envelope
	err := xml.Unmarshal(body, &response)
	if err != nil {
		logger.Fatal("Error unmarshalling response", zap.Error(err))
		return err
	}

	if response.Body.Fault != nil {
		logger.Warn("Received Fault", zap.String("faultstring", response.Body.Fault.FaultString))
		fmt.Println(response.Body.Fault.FaultString)
		return nil
	}

	if response.Body.Content == nil || response.Body.Content.Message == "" {
		message := "Response is empty"
		logger.Info(message)
		fmt.Println(message)
		return nil
	}

	for _, person := range response.Body.Content.Persons {
		fmt.Printf("- ID: %d, Name: %s, Surname: %s, Age: %d\n",
			person.ID,
			person.Name,
			person.Surname,
			person.Age,
		)

		logger.Info("Result of request execution",
			zap.Int("ID", person.ID),
			zap.String("Name", person.Name),
			zap.String("Surname", person.Surname),
			zap.Int("Age", person.Age),
		)
	}
	return nil
}

func addPerson(url string, person Person, logger *zap.Logger) {
	request := AddPersonRequest{Person: person}
	requestXML, err := xml.Marshal(request)
	if err != nil {
		logger.Fatal("Error marshaling request", zap.Error(err))
	}

	body, err := sendSOAPRequest(url, requestXML, logger)
	if err != nil {
		logger.Warn("Error calling AddPerson", zap.Error(err))
		return
	}

	var response AddPersonResponse
	if err := xml.Unmarshal(body, &response); err != nil {
		logger.Fatal("Error unmarshalling response", zap.Error(err))
	}

	fmt.Printf("Added person with ID: %d\n", response.Content.ID)
}

func getPerson(url string, id int, logger *zap.Logger) {
	request := GetPersonRequest{ID: id}
	requestXML, err := xml.Marshal(request)
	if err != nil {
		logger.Fatal("Error marshaling request", zap.Error(err))
	}

	body, err := sendSOAPRequest(url, requestXML, logger)
	if err != nil {
		logger.Warn("Error calling GetPerson", zap.Error(err))
		return
	}
	fmt.Println(string(body))
	err = printFullResponse(body, logger)
	if err != nil {
		return
	}

}

func getAllPersons(url string, logger *zap.Logger) {
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

func updatePerson(url string, person Person, logger *zap.Logger) {
	request := UpdatePersonRequest{Person: person}
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

	var response UpdatePersonResponse
	if err := xml.Unmarshal(body, &response); err != nil {
		logger.Fatal("Error unmarshalling response", zap.Error(err))
		return
	}

	fmt.Printf("Updated person successfully: %v\n", response.Success)
}

func deletePerson(url string, id int, logger *zap.Logger) {
	request := DeletePersonRequest{ID: id}
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
	var response DeletePersonResponse
	if err := xml.Unmarshal(body, &response); err != nil {
		logger.Fatal("Error unmarshalling response", zap.Error(err))
		return
	}

	fmt.Printf("Deleted person successfully: %v\n", response.Success)
}

func searchPersons(url string, query string, logger *zap.Logger) {
	request := SearchPersonRequest{Query: query}
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

func main() {

	loggerConfig := zap.NewProductionConfig()
	loggerConfig.OutputPaths = []string{"soapclient.log"}

	logger, _ := loggerConfig.Build()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	url := flag.String("url", "http://localhost:8094/soap", "SOAP server URL")
	method := flag.String("method", "", "Method to call (addperson|getperson|getallpersons|updateperson|deleteperson|searchperson)")
	name := flag.String("name", "", "Name of the person (required for addperson and updateperson)")
	surname := flag.String("surname", "", "Surname of the person (required for addperson and updateperson)")
	id := flag.Int("id", 0, "ID of the person (required for getperson, updateperson and deleteperson)")
	query := flag.String("query", "", "Query for searching person (required for searchperson)")
	age := flag.Int("age", 0, "Age of the person (required for addperson and updateperson)")

	flag.Parse()

	switch *method {
	case "addperson":
		if *name == "" || *surname == "" || *age <= 0 {
			log.Fatal("Both name and surname are required for addperson. Age must be greater than 0.")
		}
		addPerson(*url, Person{Name: *name, Surname: *surname, Age: *age}, logger)
	case "getperson":
		if *id <= 0 {
			log.Fatal("ID must be greater than 0 for getperson.")
		}
		getPerson(*url, *id, logger)
	case "getallpersons":
		getAllPersons(*url, logger)
	case "updateperson":
		if *id <= 0 || *name == "" || *surname == "" || *age <= 0 {
			log.Fatal("ID must be greater than 0 and both name and surname are required for updateperson. Age must be greater than 0.")
		}
		updateInfo := Person{ID: *id, Name: *name, Surname: *surname, Age: *age}
		updatePerson(*url, updateInfo, logger)
	case "deleteperson":
		if *id <= 0 {
			log.Fatal("ID must be greater than 0 for deleteperson.")
		}
		deletePerson(*url, *id, logger)
	case "searchperson":
		if *query == "" {
			log.Fatal("Query is required for searchperson.")
		}
		searchPersons(*url, *query, logger)
	default:
		log.Fatal("Unknown method. Use one of addperson|getperson|getallpersons|updateperson|deleteperson|searchperson.")
	}
}
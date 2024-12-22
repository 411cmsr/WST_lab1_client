package main

import (
	"WST_lab1_client/internal/handlers"
	"WST_lab1_client/internal/logger"
	"WST_lab1_client/internal/models"


	"flag"
	"fmt"

)

func main() {
	logConfig := logger.NewLoggerConfig()
	log, err := logger.NewLogger(logConfig)
	if err != nil {
		log.Fatal("Failed logger")
	}
	defer log.Sync()

	url := flag.String("url", "http://localhost:8094/soap", "SOAP server URL")
	method := flag.String("method", "", "Method to call (addperson|getperson|getallpersons|updateperson|deleteperson|searchperson)")
	name := flag.String("name", "", "Name of the person (required for addperson and updateperson)")
	surname := flag.String("surname", "", "Surname of the person (required for addperson and updateperson)")
	email := flag.String("email", "", "Email of the person (required for addperson and updateperson)")
	telephone := flag.String("telephone", "", "Telephone of the person (required for addperson and updateperson)")
	id := flag.Int("id", 0, "ID of the person (required for getperson, updateperson and deleteperson)")
	query := flag.String("query", "", "Query for searching person (required for searchperson)")
	age := flag.Int("age", 0, "Age of the person (required for addperson and updateperson)")

	flag.Parse()

	var requestXML []byte
	

	switch *method {
	case "addperson":
		if *name == "" || *surname == "" || *email == "" || *telephone == "" || *age <= 0 {
			log.Fatal("Both name and surname are required for addperson. Age must be greater than 0.")
		}
		requestXML = []byte(fmt.Sprintf(`
            <soapenv:Envelope xmlns:soapenv="http://www.w3.org/2003/05/soap-envelope">
                <soapenv:Body>
                    <AddPerson>
                        <Name>%s</Name>
                        <Surname>%s</Surname>
                        <Age>%d</Age>
                        <Email>%s</Email>
                        <Telephone>%s</Telephone>
                    </AddPerson>
                </soapenv:Body>
            </soapenv:Envelope>`, *name, *surname, *age, *email, *telephone))

	
		body, err := handlers.SendRequest(*url, requestXML, log)
		if err != nil {
			handlers.PrintError(body, log)
			return
		}
		var addPersonResp models.AddPersonResponse
		if err = handlers.ParseResponse(body, &addPersonResp, log); err == nil {
			handlers.PrintResult(addPersonResp)
		}

	case "deleteperson":
		if *id <= 0 {
			log.Fatal("ID must be greater than 0 for deleteperson.")
		}
		requestXML = []byte(fmt.Sprintf(`
            <soapenv:Envelope xmlns:soapenv="http://www.w3.org/2003/05/soap-envelope">
                <soapenv:Body>
                    <DeletePerson>
                        <ID>%d</ID>
                    </DeletePerson>
                </soapenv:Body>
            </soapenv:Envelope>`, *id))

		body, err := handlers.SendRequest(*url, requestXML, log)
		if err != nil {
			handlers.PrintError(body, log)
			return
		}
		var deleteResp models.DeleteResponse
		if err = handlers.ParseResponse(body, &deleteResp, log); err == nil {
			handlers.PrintResult(deleteResp)
		}


	case "getperson":
		if *id <= 0 {
			log.Fatal("ID must be greater than 0 for getperson.")
		}
		requestXML = []byte(fmt.Sprintf(`
            <soapenv:Envelope xmlns:soapenv="http://www.w3.org/2003/05/soap-envelope">
                <soapenv:Body>
                    <GetPerson>
                        <ID>%d</ID>
                    </GetPerson>
                </soapenv:Body>
            </soapenv:Envelope>`, *id))

		body, err := handlers.SendRequest(*url, requestXML, log)
		if err != nil {
			handlers.PrintError(body, log)
			return
		}
		var getResp models.GetPersonResponse
		if err = handlers.ParseResponse(body, &getResp, log); err == nil {
			handlers.PrintResult(getResp)
		}

	
	case "getallpersons":
		requestXML = []byte(fmt.Sprintf(`
            <soapenv:Envelope xmlns:soapenv="http://www.w3.org/2003/05/soap-envelope">
                <soapenv:Body>
                    <GetAllPersons>                       
                    </GetAllPersons>
                </soapenv:Body>
            </soapenv:Envelope>`, *id))

		body, err := handlers.SendRequest(*url, requestXML, log)
		if err != nil {
			handlers.PrintError(body, log)
			return
		}
		var getAllResp models.GetAllPersonsResponse
		if err = handlers.ParseResponse(body, &getAllResp, log); err == nil {
			handlers.PrintResult(getAllResp)
		}
	case "updateperson":
		if *id <= 0 || *name == "" || *surname == "" || *email == "" || *telephone == "" || *age <= 0 {
			log.Fatal("Both name and surname are required for addperson. Age must be greater than 0.")
		}
		requestXML = []byte(fmt.Sprintf(`
            <soapenv:Envelope xmlns:soapenv="http://www.w3.org/2003/05/soap-envelope">
                <soapenv:Body>
                    <UpdatePerson>
						<ID>%d</ID>
                        <Name>%s</Name>
                        <Surname>%s</Surname>
                        <Age>%d</Age>
                        <Email>%s</Email>
                        <Telephone>%s</Telephone>
                    </UpdatePerson>
                </soapenv:Body>
            </soapenv:Envelope>`, *id, *name, *surname, *age, *email, *telephone))

		
		body, err := handlers.SendRequest(*url, requestXML, log)
		if err != nil {
			handlers.PrintError(body, log)
			return
		}
		var updatePersonResp models.UpdatePersonResponse
		if err = handlers.ParseResponse(body, &updatePersonResp, log); err == nil {
			handlers.PrintResult(updatePersonResp)
		}
	case "searchperson":
		if *query == "" {
			log.Fatal("Query must be provided for searchperson.")
		}
		requestXML = []byte(fmt.Sprintf(`
			<soapenv:Envelope xmlns:soapenv="http://www.w3.org/2003/05/soap-envelope">
				<soapenv:Body>
					<SearchPerson>
						<Query>%s</Query>
					</SearchPerson>
				</soapenv:Body>
			</soapenv:Envelope>`, *query))
	
		body, err := handlers.SendRequest(*url, requestXML, log)
		if err != nil {
			handlers.PrintError(body, log)
			return
		}
		var searchResp models.SearchPersonResponse
		if err = handlers.ParseResponse(body, &searchResp, log); err == nil {
			handlers.PrintResult(searchResp)
		}
	default:
		log.Fatal("Unknown method. Use one of addperson|getperson|getallpersons|updateperson|deleteperson|searchperson.")
	}
}

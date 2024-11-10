package main

import (
	"WST_lab1_client/internal/handlers"
	"WST_lab1_client/internal/logger"
	"flag"
	"go.uber.org/zap"
)

func main() {
	logConfig := logger.NewLoggerConfig()

	log, err := logger.NewLogger(logConfig)
	if err != nil {
		log.Fatal("Failed logger")
	}
	defer func(log *zap.Logger) {
		err := log.Sync()
		if err != nil {

		}
	}(log)

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
		handlers.AddPersonHandler(*url, *name, *surname, *age, log)
	case "getperson":
		if *id <= 0 {
			log.Fatal("ID must be greater than 0 for getperson.")
		}
		handlers.GetPersonHandler(*url, *id, log)
	case "getallpersons":
		handlers.GetAllPersonsHandler(*url, log)
	case "updateperson":
		if *id <= 0 || *name == "" || *surname == "" || *age <= 0 {
			log.Fatal("ID must be greater than 0 and both name and surname are required for updateperson. Age must be greater than 0.")
		}
		handlers.UpdatePersonHandler(*url, *id, *name, *surname, *age, log)
	case "deleteperson":
		if *id <= 0 {
			log.Fatal("ID must be greater than 0 for deleteperson.")
		}
		handlers.DeletePersonHandler(*url, *id, log)
	case "searchperson":
		if *query == "" {
			log.Fatal("Query is required for searchperson.")
		}
		handlers.SearchPersonsHandler(*url, *query, log)
	default:
		log.Fatal("Unknown method. Use one of addperson|getperson|getallpersons|updateperson|deleteperson|searchperson.")
	}
}

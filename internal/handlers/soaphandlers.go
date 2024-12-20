package handlers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	

	"WST_lab1_client/internal/models"

	"go.uber.org/zap"
)

// SendRequest sends a SOAP request to the specified URL and returns the response.
func SendRequest(url string, requestXML []byte, logger *zap.Logger) ([]byte, error) {
	reqBody := bytes.NewBuffer(requestXML)

	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		logger.Error("Error creating request", zap.Error(err))
		return nil, err
	}

	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")

	logger.Info("Sending SOAP request", zap.String("url", url), zap.String("request", string(requestXML)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error sending request", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	bodyResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body", zap.Error(err))
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		logger.Warn("Received non-OK response status", zap.Int("status", resp.StatusCode), zap.String("response", string(bodyResponse)))
		return bodyResponse, fmt.Errorf("received non-OK response status: %d", resp.StatusCode)
	}
	//fmt.Println(string(bodyResponse))
	return bodyResponse, nil
}

// ParseResponse parses the SOAP response based on the expected type.
func ParseResponse(body []byte, response interface{}, logger *zap.Logger) error {
	if err := xml.Unmarshal(body, response); err != nil {
		logger.Fatal("Error unmarshalling response", zap.Error(err), zap.String("response_body", string(body)))
		return err
	}
	return nil
}

// PrintResult prints the result to the console.
func PrintResult(result interface{}) {
	switch res := result.(type) {
	case models.DeleteResponse:
		fmt.Printf("Delete status: %v\n", res.Status)
	case models.AddPersonResponse:
		fmt.Printf("Added person with ID: %d\n", res.ID)
	case models.UpdatePersonResponse:
		fmt.Printf("Update status: %v\n", res.Status)
	case models.GetPersonResponse:
		fmt.Printf("Retrieved person:\n ID: %+v\n Name: %v\n Surname: %v\n Age: %v\n Email: %v\n Telephone: %v\n", res.Person.ID, res.Person.Name, res.Person.Surname, res.Person.Age, res.Person.Email, res.Person.Telephone)
	case models.GetAllPersonsResponse:
		for _, p := range res.Persons {
			fmt.Printf("Person: ID: %+v Name: %v Surname: %v Age: %v Email: %v Telephone: %v\n", p.ID, p.Name, p.Surname, p.Age, p.Email, p.Telephone)
			//fmt.Printf("Person: %+v\n", p)
		}
		//fmt.Printf("Retrieved persons: %+v\n", res.Persons)
	case models.SearchPersonResponse:
		fmt.Printf("Search results: %+v\n", res.Persons)
	default:
		fmt.Println("Unknown response type")
	}
}

// PrintError prints error details from the SOAP error response.
func PrintError(body []byte, logger *zap.Logger) {
	var errorResponse models.Fault
	fmt.Println(string(body))
	if err := xml.Unmarshal(body, &errorResponse); err == nil {
		logger.Warn("SOAP Error Response",
			zap.String("FaultCode", errorResponse.FaultCode),
			zap.String("FaultString", errorResponse.FaultString),


		)
		fmt.Printf("Error: %s (Status: %d)\nDetail: %s\nInstance: %s\n",
			errorResponse.FaultCode,
			errorResponse.FaultString,
			
			
		)
	} else {
		logger.Warn("Failed to parse error response",
			zap.String("response_body", string(body)),
			zap.Error(err))
	}
}

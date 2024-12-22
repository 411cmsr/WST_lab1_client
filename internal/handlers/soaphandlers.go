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
	
	return bodyResponse, nil
}


func ParseResponse(body []byte, response interface{}, logger *zap.Logger) error {
	if err := xml.Unmarshal(body, response); err != nil {
		logger.Fatal("Error unmarshalling response", zap.Error(err), zap.String("response_body", string(body)))
		return err
	}
	return nil
}

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
		}
	case models.SearchPersonResponse:
		for _, p := range res.Persons {
			fmt.Printf("Person: ID: %+v Name: %v Surname: %v Age: %v Email: %v Telephone: %v\n", p.ID, p.Name, p.Surname, p.Age, p.Email, p.Telephone)
		}
	default:
		fmt.Println("Unknown response type")
	}
}

func PrintError(body []byte, logger *zap.Logger) {
	var errorResponse models.SOAPFault
	//fmt.Println(string(body))
	if err := xml.Unmarshal(body, &errorResponse); err == nil {
		logger.Warn("SOAP Error Response",
			zap.String("FaultCode", errorResponse.Envelope.Body.Fault.FaultCode),
			zap.String("FaultString", errorResponse.Envelope.Body.Fault.FaultString),
			zap.String("Detail", errorResponse.Envelope.Body.Fault.FaultDetail.ErrorCode),
			zap.String("Message", errorResponse.Envelope.Body.Fault.FaultDetail.ErrorMessage),

		)
		fmt.Printf("Error: \n FaultCode: %s \n FaultString: %s\n Detail:\n ErrorCode: %s\n ErrorMessage: %s\n", errorResponse.Envelope.Body.Fault.FaultCode, errorResponse.Envelope.Body.Fault.FaultString, errorResponse.Envelope.Body.Fault.FaultDetail.ErrorCode, errorResponse.Envelope.Body.Fault.FaultDetail.ErrorMessage)
	} else {
		logger.Warn("Failed to parse error response",
			zap.String("response_body", string(body)),
			zap.Error(err))
	}
}

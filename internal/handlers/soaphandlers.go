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
		fmt.Printf("Retrieved person: %+v\n", res.Person)
	case models.GetAllPersonsResponse:
		fmt.Printf("Retrieved persons: %+v\n", res.Persons)
	case models.SearchPersonResponse:
		fmt.Printf("Search results: %+v\n", res.Persons)
	default:
		fmt.Println("Unknown response type")
	}
}

// PrintError prints error details from the SOAP error response.
func PrintError(body []byte, logger *zap.Logger) {
	var errorResponse models.ErrorResponse
	if err := xml.Unmarshal(body, &errorResponse); err == nil {
		logger.Warn("SOAP Error Response",
			zap.String("title", errorResponse.Title),
			zap.Int("status", errorResponse.Status),
			zap.String("detail", errorResponse.Detail),
			zap.String("instance", errorResponse.Instance),
		)
		fmt.Printf("Error: %s (Status: %d)\nDetail: %s\nInstance: %s\n",
			errorResponse.Title,
			errorResponse.Status,
			errorResponse.Detail,
			errorResponse.Instance,
		)
	} else {
		logger.Warn("Failed to parse error response",
			zap.String("response_body", string(body)),
			zap.Error(err))
	}
}

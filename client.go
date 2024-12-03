package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// AddRequest structure for the addition request
type AddRequest struct {
	XMLName xml.Name `xml:"Add"`
	A       int      `xml:"a"`
	B       int      `xml:"b"`
}

// sendSOAPRequest sends the SOAP request and returns the response
func sendSOAPRequest(url string, requestBody AddRequest) ([]byte, error) {
	// Create the SOAP envelope
	envelope := struct {
		XMLName xml.Name `xml:"soap:Envelope"`
		Xmlns   string   `xml:"xmlns:soap,attr"`
		Body    struct {
			XMLName xml.Name
			Content AddRequest
		}
	}{
		Xmlns: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: struct {
			XMLName xml.Name
			Content AddRequest
		}{
			XMLName: xml.Name{Local: "soap:Body"},
			Content: requestBody,
		},
	}

	// Serialize the SOAP envelope to XML
	xmlBody, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling XML: %v", err)
	}

	// Send the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xmlBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("SOAPAction", "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	return body, nil
}

func main() {
	// SOAP server URL
	url := "http://localhost:8080/math"

	// Create the request for the Add operation
	request := AddRequest{A: 25, B: 25}

	// Send the request
	responseData, err := sendSOAPRequest(url, request)
	if err != nil {
		log.Fatalf("Error making SOAP request: %v", err)
	}

	// Print the response
	fmt.Println("Response:", string(responseData))
}

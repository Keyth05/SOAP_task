package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// AddRequest estructura para la solicitud de suma
type AddRequest struct {
	XMLName xml.Name `xml:"Add"`
	A       int      `xml:"a"`
	B       int      `xml:"b"`
}

// sendSOAPRequest envía la solicitud SOAP y devuelve la respuesta
func sendSOAPRequest(url string, requestBody AddRequest) ([]byte, error) {
	// Crear el sobre SOAP
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

	// Serializar el sobre SOAP a XML
	xmlBody, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling XML: %v", err)
	}

	// Enviar la solicitud
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

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	return body, nil
}

func main() {
	// Dirección del servidor SOAP
	url := "http://localhost:8080/math"

	// Crear la solicitud para la operación Add
	request := AddRequest{A: 25, B: 25}

	// Enviar la solicitud
	responseData, err := sendSOAPRequest(url, request)
	if err != nil {
		log.Fatalf("Error making SOAP request: %v", err)
	}

	// Mostrar la respuesta
	fmt.Println("Response:", string(responseData))
}

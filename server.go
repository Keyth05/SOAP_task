package main

import (
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

// AddResponse structure for the addition response
type AddResponse struct {
	XMLName xml.Name `xml:"AddResponse"`
	Result  int      `xml:"result"`
}

// handleSOAP handles SOAP requests
func handleSOAP(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Deserialize the XML request body
	var envelope struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			XMLName xml.Name
			Add     *AddRequest `xml:"Add"`
		}
	}

	err = xml.Unmarshal(body, &envelope)
	if err != nil {
		http.Error(w, "Invalid SOAP format", http.StatusBadRequest)
		return
	}

	// Check if the request contains the Add operation data
	if envelope.Body.Add != nil {
		// Perform the addition operation
		result := envelope.Body.Add.A + envelope.Body.Add.B
		response := AddResponse{Result: result}

		// Prepare the SOAP response
		w.Header().Set("Content-Type", "text/xml")
		fmt.Fprintf(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
		<soap:Body>
		<AddResponse>
		<result>%d</result>
		</AddResponse>
		</soap:Body>
		</soap:Envelope>`, response.Result)
		return
	}

	http.Error(w, "Unknown operation", http.StatusBadRequest)
}

func main() {
	// Define the SOAP service route
	http.HandleFunc("/math", handleSOAP)

	// Start the server on port 8080
	fmt.Println("SOAP server running on http://localhost:8080/math")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

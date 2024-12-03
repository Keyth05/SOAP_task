package main

import (
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

// AddResponse estructura para la respuesta de suma
type AddResponse struct {
	XMLName xml.Name `xml:"AddResponse"`
	Result  int      `xml:"result"`
}

// handleSOAP maneja las solicitudes SOAP
func handleSOAP(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Deserializar el cuerpo XML de la solicitud
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

	// Verificar si la solicitud contiene los datos de la operación Add
	if envelope.Body.Add != nil {
		// Realizar la operación de suma
		result := envelope.Body.Add.A + envelope.Body.Add.B
		response := AddResponse{Result: result}

		// Preparar la respuesta SOAP
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
	// Definir la ruta del servicio SOAP
	http.HandleFunc("/math", handleSOAP)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("SOAP server running on http://localhost:8080/math")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

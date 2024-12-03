# SOAP Server in Go: Sum of Two Numbers 🧮

This project implements a SOAP server in Go that exposes a service to perform the **sum of two numbers** using SOAP requests. SOAP (Simple Object Access Protocol) is an XML-based protocol for exchanging messages between distributed applications.

---
## Description 📝

This SOAP server allows a client to send a request with two numbers to sum (in XML format), and the server will return the sum of these two numbers in XML format. 

---
### Service Features 🚀

- Performs the **sum** of two numbers.
- Exposes the service as a SOAP operation.

---
## Execution Process ⚙️

1. Open your terminal and navigate to the folder where server.go is located:
``` go run server.go```

    The server will be available at http://localhost:8080/math and will listen for SOAP requests.


2. Then, open new terminal and navigate to the folder where client.go is located:
``` go run client.go```

    Finaly, the client will send a request to the SOAP server and display the sum of the numbers. 

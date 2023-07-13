package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailRequest struct {
	Recipient string `json:"recipient"`
}


func sendEmail( recipient string) {
	from := mail.NewEmail("Ayush", "ayushshreeshreemal@gmail.com")
	subject := "React"
	to := mail.NewEmail("Example User", recipient)
	plainTextContent := "This email was sent through React."
	htmlContent := "<strong>This email was sent through React.</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func handleEmailRequest(w http.ResponseWriter, r *http.Request) {
	var req EmailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendEmail(req.Recipient)
}
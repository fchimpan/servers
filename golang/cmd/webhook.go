package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	fmt.Println("Received GitHub Webhook Payload:")
	fmt.Println(string(body))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Webhook received successfully")
}

func main() {
	http.HandleFunc("/webhook", WebhookHandler)

	fmt.Println("GitHub Webhook server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

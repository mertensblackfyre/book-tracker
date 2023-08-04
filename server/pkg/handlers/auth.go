package handlders

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)


func handleWebhook(w http.ResponseWriter, r *http.Request) {
	// Verify webhook signature
	err := clerk.Ver  VerifyWebhookSignature(r)
	if err != nil {
		// Handle error
		log.Println(err)
	}

	// Parse webhook payload
	var payload client.Webhook
		err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		// Handle error
	}

	// Handle webhook event
	switch payload.EventType {
	case "user.created":
		// Sync new user with database
		fmt.Println("New user created:", payload.EventData.(clerk.User))
	case "user.updated":
		// Sync updated user with database
		fmt.Println("User updated:", payload.EventData.(clerk.User))
	case "user.deleted":
		// Sync deleted user with database
		fmt.Println("User deleted:", payload.EventData.(clerk.User))
	default:
		// Ignore other events
		fmt.Println("Unknown event type:", payload.EventType)
	}
}

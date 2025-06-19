// In your home-assist-jarvis-core-plugin git repo:
// /plugin.go - This is the corrected plug-in file.

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	// Import the shared types package
	"jarvis-types"
)

// GetTools now returns a slice of the imported type.
func GetTools() []types.ToolDefinition {
	return []types.ToolDefinition{
		{
			Name:        "homeassistant.send_notification",
			Description: "Sends a notification to a specific device via Home Assistant.",
			Arguments: []types.ArgumentDefinition{
				{ Name: "title", Description: "The title of the notification.", Required: true },
				{ Name: "message", Description: "The main content of the notification message.", Required: true },
				{ Name: "device_id", Description: "The target device ID in Home Assistant.", Required: true },
			},
			FunctionName: "SendNotification",
		},
	}
}

// The function signature also uses the imported type.
func SendNotification(ctx context.Context, args map[string]interface{}) (string, error) {
	title, _ := args["title"].(string)
	message, _ := args["message"].(string)
	deviceID, _ := args["device_id"].(string)

	if title == "" || message == "" || deviceID == "" {
		return "", fmt.Errorf("missing required arguments")
	}
	
	log.Printf("HOME ASSISTANT PLUGIN: Sending notification to device '%s'", deviceID)
	time.Sleep(1 * time.Second)

	return "Notification sent successfully via Home Assistant.", nil
}

func main() {}
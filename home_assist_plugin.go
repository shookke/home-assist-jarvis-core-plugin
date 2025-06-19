// This is the corrected main file for the Home Assistant Assist plug-in.
// It includes all necessary, matching struct definitions to ensure compatibility
// with the main Jarvis Core service that loads it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// =========================================================================
// Models and Core Structs
// These MUST exactly match the definitions in the main Jarvis system.
// =========================================================================

type ToolFunction func(ctx context.Context, args map[string]interface{}) (string, error)

type ArgumentDefinition struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type ToolDefinition struct {
	Name         string               `json:"name"`
	Description  string               `json:"description"`
	Arguments    []ArgumentDefinition `json:"arguments"`
	FunctionName string               `json:"function_name"`
}

// =========================================================================
// Plug-in Implementation
// =========================================================================

// GetTools is the required public entry point that Jarvis Core will look for.
// It returns a slice of all tool definitions provided by this plug-in.
func GetTools() []ToolDefinition {
	return []ToolDefinition{
		{
			Name:        "homeassistant.send_notification",
			Description: "Sends a notification to a specific device via Home Assistant.",
			Arguments: []ArgumentDefinition{
				{
					Name:        "title",
					Description: "The title of the notification.",
					Required:    true,
				},
				{
					Name:        "message",
					Description: "The main content of the notification message.",
					Required:    true,
				},
				{
					Name:        "device_id",
					Description: "The target device ID in Home Assistant (e.g., 'mobile_app_shookkes_iphone').",
					Required:    true,
				},
			},
			FunctionName: "SendNotification", // This must match the function name below.
		},
	}
}

// SendNotification is the actual implementation of the tool.
// Its name matches the FunctionName in the definition above.
func SendNotification(ctx context.Context, args map[string]interface{}) (string, error) {
	title, _ := args["title"].(string)
	message, _ := args["message"].(string)
	deviceID, _ := args["device_id"].(string)

	if title == "" || message == "" || deviceID == "" {
		return "", fmt.Errorf("missing required arguments for homeassistant.send_notification")
	}

	// In a real implementation, this would make an API call to the Home Assistant instance.
	log.Printf("HOME ASSISTANT PLUGIN: Sending notification to device '%s' with title: '%s'", deviceID, title)
	time.Sleep(1 * time.Second) // Simulate network latency

	return "Notification successfully sent via Home Assistant.", nil
}

// The main function is required for the file to be a valid Go program,
// but it is ignored when the code is built as a plug-in.
func main() {}
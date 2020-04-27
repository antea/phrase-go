package phrase
// WebhookCreate struct for WebhookCreate
type WebhookCreate struct {
	// Callback URL to send requests to
	CallbackUrl string `json:"callback_url,omitempty"`
	// Webhook description
	Description string `json:"description,omitempty"`
	// List of event names to trigger the webhook (separated by comma)
	Events string `json:"events,omitempty"`
	// Whether webhook is active or inactive
	Active string `json:"active,omitempty"`
}

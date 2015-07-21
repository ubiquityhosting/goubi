package goubi

import (
	"strconv"
)

/*****************************************************************************/

// Represents an array of welcome emails
type WelcomeEmails struct {
	WelcomeEmails []WelcomeEmail `json:"WelcomeEmails"`
}

// Represents a welcome email
type WelcomeEmail struct {
	Id          int    `json:"id,string"`
	SendTo      string `json:"send_to"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	Date        int    `json:"date,string"`
	PlacedAt    int    `json:"placed_at,string"`
	ServiceId   int    `json:"service_id,string"`
	ProductType string `json:"product_type"`
	InvoiceId   int    `json:"invoice_id,string"`
	Action      string `json:"action"`
}

/*****************************************************************************/

// Retrieves all welcome emails associated with a VM
func (c *Cloud) GetWelcomeEmails(vm_id int) ([]WelcomeEmail, error) {
	wm := &WelcomeEmails{}
	err := c.callWithParams("cloud.get_welcome_emails", map[string]string{"vm_id": strconv.Itoa(vm_id)}, &wm)
	return wm.WelcomeEmails, err
}

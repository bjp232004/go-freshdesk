package models

type Source uint8

const (
	Email Source = iota + 1
	Portal
	Phone
	Chat Source = iota + 4
	Mobihelp
	FeedbackWidget
	OutboundEmail
)

type Priority uint8

const (
	Low Priority = iota + 1
	Medium
	High
	Urgent
)

type Status uint8

const (
	Open Status = iota + 2
	Pending
	Resolved
	Closed
)

type Ticket struct {
	ID uint64 `json:"id,omitempty"`

	// Name of the requester
	Name string `json:"name"`

	// User ID of the requester. For existing contacts, the requester_id can be passed instead of the requester's email.
	RequesterID uint64 `json:"requester_id"`

	// Email address of the requester. If no contact exists with this email address in Freshdesk, it will be added as a new contact.
	Email string `json:"email"`

	// Facebook ID of the requester. A contact should exist with this facebook_id in Freshdesk.
	// FacebookID string `json:"facebook_id"`

	// Phone number of the requester. If no contact exists with this phone number in Freshdesk, it will be added as a new contact. If the phone number is set and the email address is not, then the name attribute is mandatory.
	// Phone string `json:"phone"`

	// Twitter handle of the requester. If no contact exists with this handle in Freshdesk, it will be added as a new contact.
	// TwitterID string `json:"twitter_id"`

	// External ID of the requester. If no contact exists with this external ID in Freshdesk, they will be added as a new contact.
	// UniqueExternalID string `json:"unique_external_id"`

	// Subject of the ticket. The default Value is null.
	Subject string `json:"subject"`

	// Helps categorize the ticket according to the different kinds of issues your support team deals with. The default Value is null.
	Type string `json:"type"`

	// Status of the ticket. The default Value is 2.
	Status Status `json:"status"`

	// Priority of the ticket. The default value is 1.
	Priority Priority `json:"priority"`

	// HTML content of the ticket.
	Description string `json:"description"`

	// ID of the agent to whom the ticket has been assigned
	ResponderID uint64 `json:"responder_id"`

	// Ticket attachments. The total size of these attachments cannot exceed 15MB.
	// Not implemented
	// Attachments []interface{} `json:"attachments"`

	// Email address added in the 'cc' field of the incoming ticket email
	// CCEmails []string `json:"cc_emails"`

	// FWDEmails []string `json:"fwd_emails"`

	// REPLYCCEmails []string `json:"reply_cc_emails"`

	// Key value pairs containing the names and values of custom fields. Read more here
	// CustomFields map[string]interface{} `json:"custom_fields"`

	// Timestamp that denotes when the ticket is due to be resolved
	// DueBy time.Time `json:"due_by"`

	// ID of email config which is used for this ticket. (i.e., support@yourcompany.com/sales@yourcompany.com)
	// If product_id is given and email_config_id is not given, product's primary email_config_id will be set
	// EmailConfigID uint64 `json:"email_config_id"`

	// Timestamp that denotes when the first response is due
	// FirstResponseDueBy time.Time `json:"fr_due_by"`

	// ID of the group to which the ticket has been assigned. The default value is the ID of the group that is associated with the given email_config_id
	// GroupID uint64 `json:"group_id"`

	// ID of the product to which the ticket is associated. It will be ignored if the email_config_id attribute is set in the request.
	// ProductID uint64 `json:"product_id"`

	// The channel through which the ticket was created. The default value is 2.
	Source Source `json:"source"`

	// Tags that have been associated with the ticket
	// Tags []string `json:"tags"`

	// Company ID of the requester. This attribute can only be set if the Multiple Companies feature is enabled (Estate plan and above)
	// CompanyID uint64 `json:"company_id"`
}

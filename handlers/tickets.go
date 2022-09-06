package handlers

import (
	"fmt"

	"github.com/bjp232004/go-freshdesk/models"
)

type Ticket models.Ticket

type TicketsClient struct {
	client *Client
}

// Create a Ticket
func (c *TicketsClient) Create(t *Ticket) (*Ticket, error) {
	// foo_marshalled, _ := json.Marshal(t)
	// foo_marshalled_str := string(foo_marshalled)
	// fmt.Print(foo_marshalled_str)

	req, err := c.client.freshDeskRequest("POST", "tickets", t)

	if err != nil {
		fmt.Print("inside create if")
		return nil, err
	}
	fmt.Print(req)
	res := new(Ticket)
	if err := c.client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Update a Ticket
func (c *TicketsClient) Update(id uint64, t *Ticket) (*Ticket, error) {
	req, err := c.client.freshDeskRequest("PUT", fmt.Sprintf("tickets/%d", id), t)
	if err != nil {
		return nil, err
	}

	res := new(Ticket)
	if err := c.client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// View a Ticket
func (c *TicketsClient) View(id uint64) (*Ticket, error) {
	req, err := c.client.freshDeskRequest("GET", fmt.Sprintf("tickets/%d", id), nil)
	if err != nil {
		return nil, err
	}

	res := new(Ticket)
	if err := c.client.do(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// List All Tickets
func (c *TicketsClient) ListAll() ([]*Ticket, error) {
	req, err := c.client.freshDeskRequest("GET", "tickets", nil)
	if err != nil {
		return nil, err
	}

	var res []*Ticket
	if err := c.client.do(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// Delete a Ticket
func (c *TicketsClient) Delete(id uint64) error {
	req, err := c.client.freshDeskRequest("DELETE", fmt.Sprintf("tickets/%d", id), nil)
	if err != nil {
		return err
	}

	return c.client.do(req, nil)
}

// Restore a Ticket
func (c *TicketsClient) Restore(id uint64) error {
	req, err := c.client.freshDeskRequest("PUT", fmt.Sprintf("tickets/%d/restore", id), nil)
	if err != nil {
		return err
	}

	return c.client.do(req, nil)
}

package main

func (ticket Ticket) getNameById(id string) string {
	return ("Ticket dengan Id: " + id + " Ticket Name: " + ticket.Name)
}

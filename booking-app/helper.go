package main

import (
	"strings"
)

func validateFields(user UserData) (bool, string) {

	var valid = true
	var err = ""
	if user.fname == "" {
		err = "Incorrect first name"
		valid = false
	}
	if user.lname == "" {
		err = "Incorrect last name"
		valid = false
	}
	if user.email == "" || !strings.Contains(user.email, "@") {
		err = "Invalid email id"
		valid = false
	}
	if user.tickets > totalTickets {
		err = "Unable to book as the amount of tickets exceeds the what we have"
		valid = false
	}

	return valid, err
}

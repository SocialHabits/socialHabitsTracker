package util

import (
	"fmt"
	"net"
	"net/mail"
	"strings"
)

func CheckEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	if err != nil {
		fmt.Println("Email or domain is not valid!")
		return false
	}

	parts := strings.Split(email, "@")
	mx, err := net.LookupMX(parts[1])

	if err != nil || len(mx) == 0 {
		fmt.Println("Domain is not valid!")
		return false
	}

	return true
}

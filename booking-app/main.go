package main

import (
	"fmt"
	"time"
)

var totalTickets int = 100

var confMap = make(map[int]string)

type UserData struct {
	fname      string
	lname      string
	email      string
	tickets    int
	conference string
}

// create the default userData struct
var user UserData = UserData{" ", " ", " ", 0, " "}

// slice of user data
var listOfUsers []ticketBooking

type ticketBooking interface {
	bookTicket()
}

func config(cNum int) {
	confMap[1] = "Tokyo"
	confMap[2] = "Mumbai"
	confMap[3] = "New York"
	confMap[4] = "Singapore"

	fmt.Printf("\n You have selected %v conference", confMap[cNum])

}

func welcomeMsg() {

	var cNum int
	fmt.Print("Welcome to the tickect booking portal !\n")
	fmt.Print("What conference do you which to attend ? \n")
	fmt.Print("\n 1. Tokyo \n 2. Mumbai \n 3. New York \n 4. Singapore")
	fmt.Print("\n Enter the corresponding number next to the location.\n")
	fmt.Scan(&cNum)
	config(cNum)
	switch cNum {
	case 1:
		fmt.Println("You have selected the Tokyo conference")
	case 2:
		fmt.Println("You have selected the Mumbai conference")
	case 3:
		fmt.Println("You have selected the New York conference")
	case 4:
		fmt.Println("You have selected the Singapore conference")
	default:
		fmt.Println("Invalid Choice")
	}

	fmt.Println("Enter your first name")
	fmt.Scan(&user.fname)
	fmt.Println("Enter your last name")
	fmt.Scan(&user.lname)
	fmt.Println("Enter your email")
	fmt.Scan(&user.email)
	fmt.Printf("\nEnter the tickets you want to book, we have %v\n", totalTickets)
	fmt.Scan(&user.tickets)
	totalTickets = totalTickets - user.tickets

}

func (u UserData) bookTicket() {

	valid, err := validateFields(u)
	if !valid {
		fmt.Println(err)
	} else {
		time.Sleep(2 * time.Second)
		fmt.Println("#######################")
		fmt.Println("\nTicket booked!")
		fmt.Printf("\nThank You %v for booking with us. See you at the conference!!\n", user.fname)
	}

}

func main() {
	//var exit int
	for {
		welcomeMsg()
		listOfUsers = append(listOfUsers, user)
		fmt.Printf("\nEnter -1 to exit ")
		//fmt.Scan(&exit)
		go user.bookTicket()
		// if exit < 0 {
		// 	break
		// }

	}
	// for _, u := range listOfUsers {
	// 	u.bookTicket()
	// }
}

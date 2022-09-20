package main

import (
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	Email string
	Name  string
}

var db = []User{
	{Email: "lwithull0@shareasale.com", Name: "Lonna Withull"},
	{Email: "nmccaig1@istockphoto.com", Name: "Nancy McCaig"},
	{Email: "ode2@odnoklassniki.ru", Name: "Ora De Mars"},
	{Email: "tgidney3@github.io", Name: "Theda Gidney"},
	{Email: "sbeurich4@prnewswire.com", Name: "Syd Beurich"},
	{Email: "scerman5@wufoo.com", Name: "Sancho Cerman"},
	{Email: "ldabbotdoyle6@cdbaby.com", Name: "Laurice D'Abbot-Doyle"},
	{Email: "ajoseff7@state.gov", Name: "Arther Joseff"},
	{Email: "ttuma8@t-online.de", Name: "Tadeo Tuma"},
	{Email: "dblowfield9@ehow.com", Name: "Derrek Blowfield"},
	{Email: "mvona@livejournal.com", Name: "Manny von Hagt"},
	{Email: "mpearceyb@weebly.com", Name: "Marcelo Pearcey"},
	{Email: "mjendrysc@myspace.com", Name: "Melonie Jendrys"},
	{Email: "baronind@geocities.com", Name: "Billie Aronin"},
	{Email: "ecrowthe@oaic.gov.au", Name: "Ellsworth Crowth"},
	{Email: "ebaldellif@reuters.com", Name: "Edmund Baldelli"},
	{Email: "gmiong@walmart.com", Name: "Grantham Mion"},
	{Email: "ayoungeh@admin.ch", Name: "Ansell Younge"},
	{Email: "sredingi@cornell.edu", Name: "Staci Reding"},
	{Email: "ggilardengoj@google.com.hk", Name: "Georgia Gilardengo"},
}

type Worker struct {
	users []User
	chOut chan<- *User
}

func NewWorker(users []User, chOut chan<- *User) *Worker {
	return &Worker{users: users, chOut: chOut}
}

func (w *Worker) Find(email string) {
	for _, user := range w.users {
		if strings.Contains(user.Email, email) {
			w.chOut <- &user
		}
	}
}

func main() {
	// get the email to search for
	email := os.Args[1]

	// create channel and worker
	ch := make(chan *User)

	// attemt to find the user with the entered email
	log.Printf("searching for %s\n", email)
	go NewWorker(db[:5], ch).Find(email)
	go NewWorker(db[5:10], ch).Find(email)
	go NewWorker(db[10:15], ch).Find(email)
	go NewWorker(db[15:], ch).Find(email)

	// retrieve and return result
	for {
		select {
		case user := <-ch:
			log.Printf("email %s belongs to %s", user.Email, user.Name)
		case <-time.After(1 * time.Second):
			log.Printf("email %s was not found", email)
		}
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// tells Go we want to create a new type
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

// reciever is like a method to an instance spec to type; binds function to vars of bound type
// (input type) any variable of type deck gets access to this method
// uses one/two letters from type instead of 'self' or 'this'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// regular func with return stated
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// uses reciever to bind to type deck
// crates type conversion to return string and joins each with comma
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// returns error type if ioutil is unsuccessful
// converts deck type to string; passes in file and converted type
func (d deck) saveToFile(filename string) error {
	stringToByte := []byte(d.toString())
	return ioutil.WriteFile(filename, stringToByte, 0666)
}

// no reciever bc there is no deck upon attempt to create one
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	// kills program if ioutil returns err
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// loop; random num; swap current card w card at rand num i
	for i := range d {
		newPos := rand.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

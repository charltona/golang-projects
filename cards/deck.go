package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of Deck
// Which is slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"♠", "♥", "♦", "♣"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// use _ to silence codechecks on unused variables.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" "+value+" "+suit)
		}
	}

	return cards
}

// Create a new function that prints out the deck.
// Receivers!
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// No receiver here. Takes the deck in as a param.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	handleError(err)

	s := strings.Split(string(bs), ",")

	return deck(s)

}

func handleError(err error) {
	if err != nil {
		fmt.Println("[Error]: ", err)
		os.Exit(1)
	}
}

func (d deck) shuffle() {

	r := createNewRNG()

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func createNewRNG() *rand.Rand {
	source := rand.NewSource(time.Now().UnixNano())
	return rand.New(source)
}

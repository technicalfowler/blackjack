package main

import ( 
	"fmt"
	"math/rand"
	"time"
)

func main() {
	deck := createDeck()
	deck.randomize()

	var hand = deck[0:2]
	deck = deck[2:]
	hand = append(hand, deck.hit())
}

func scoreHand(h []card) int {
	score := 0
	for i := range h {
		score += h[i].numValue()
	}
	return score
}

/* Card */
type card struct {
        suit    string
        value   string
}


func (c *card) toString() {
        fmt.Println(c.value + " " + c.suit )
}
func (c *card) isAce() bool {
	if c.value != "A" {
		return false
	}
	return true 
}

var numValues = map[string]int {
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"10": 10,
	"J": 10,
	"Q": 10,
	"K": 10,
	"A": 11,
}
func (c *card) numValue() int {
	return numValues[c.value]	
}

/* Deck */
type deck []card

var suits = []string{"H","S","D","C"}
var faces = []string{ "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A" }

func createDeck() deck {
	var d []card
	var faceSize = len(faces)
	for i := 0; i < 4; i++ {
		suit := suits[i]
		for j := 0; j < faceSize; j++ {
			face := faces[j]
			c := card{ suit: suit, value: face }
			d = append(d,c)
		}
	}

	return d
}

func (d deck) randomize() {
	rand.Seed(time.Now().UTC().UnixNano())
	var newDeck []card
	var deckSize = len(d)
	for i := 0; i < deckSize; i++ {
		/* Copy card from random index */
		randidx := rand.Intn(len(d))
		c := d[ randidx ]
		newDeck = append(newDeck, c)

		/* Set random index element to last element of deck */
		d[ randidx ] = d[ len(d)-1 ]

		/* Reallocate deck to be one element smaller */
		d = d[ 0 : len(d)-1 ] 
	}
}

/* return a card from the deck */
func (d deck) hit() card {
	fmt.Println("size of deck", len(d))
	c := d[0]
	fmt.Println("card drawn", c)

	/* slice the array */
	d = d[1:]
	fmt.Println("new size of deck")

	return c
}




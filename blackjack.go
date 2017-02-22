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


/* Hand */
type hand struct {
	cards []card
	// all potential scores the hand could have (dealing with aces)
	scores []int
}

func (h *hand) score() int {
	numAces := 0
        score := 0
	var scores = []int{0}
        for i := range h.cards {
      		score += h.cards[i].numValue()
		// for each Ace, record a new potential score with value of 11 or 2
                if h.cards[i].isAce() {
			numAces++
			//11 has already been added to score, make new score with potential of 2
			//create new score	
                }
        }

	// for each ace seen, create potential new scores based on 2 <=> 11

        // If hand has an Ace, and busts, rescore with Ace = 2?
        if (score > 21 && numAces > 0) {
                fmt.Println("Hand has an ace, busts with score of %d, hand is:",
                        score, h)
        }

	return score
}

/* does this hand bust */
func (h *hand) busts() bool {
	if h.score() > 21 { 
		return true
	}
	return false
}

/* given another hand, does this hand beat that hand? */
func (h *hand) beats(opp *hand) bool {
	// assume the opponent hand is valid, so if this hand busts, it does not beat opp
	if h.busts() {
		return false
	}

	
	return false
}


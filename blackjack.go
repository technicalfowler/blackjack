package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []card
type player struct {
    //can i name an instance var of type `hand` to have name `hand`?
    hand        hand
    isDealer    bool
}

//#TODO: game holds the shoe, the collection of players, 
//knows how to deal a hand then deliver the score
type game struct {
    shoe    deck
}

func main() {
	deck := createDeck()
	deck = deck.randomize()

    var a = hand{cards: deck[0:2]}
    a.toString()
	var hand = hand{cards: deck[0:2]}
	deck = deck[2:]

	reader := bufio.NewReader(os.Stdin)

	for loopctrl := true; loopctrl != false; {
		hand.toString()

		fmt.Print("Options:\n1)stand\n2)hit\n")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\n")

		fmt.Print("you chose:", cmd, ":")
		fmt.Print("----\n")

		switch cmd {
		case "hit":
            // #TODO clean this junk up -- gross
			var newCardArray []card
			deck, newCardArray = hit(deck)
			hand.cards = append(hand.cards, newCardArray[0])
			if hand.busts() {
				hand.toString()
				loopctrl = false
			}
		case "stand":
			fmt.Print("You chose to stand\n")
			hand.toString()
			loopctrl = false
		default:
			loopctrl = false
		}
	}
}

/* Card */
type card struct {
	suit  string
	value string
}

func (c *card) toString() {
	fmt.Println(c.value + " " + c.suit)
}
func (c *card) isAce() bool {
	if c.value != "A" {
		return false
	}
	return true
}

var numValues = map[string]int{
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  10,
	"Q":  10,
	"K":  10,
	"A":  11,
}

func (c *card) numValue() int {
	return numValues[c.value]
}

var suits = []string{"H", "S", "D", "C"}
var faces = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

func createDeck() deck {
	var d []card
	var faceSize = len(faces)
	for i := 0; i < 4; i++ {
		suit := suits[i]
		for j := 0; j < faceSize; j++ {
			face := faces[j]
			c := card{suit: suit, value: face}
			d = append(d, c)
		}
	}

	return d
}

func (d deck) randomize() []card {
	rand.Seed(time.Now().UTC().UnixNano())
	var newDeck deck
	for {
		/* Copy card from random index */
		randidx := rand.Intn(len(d))
		c := d[randidx]
		newDeck = append(newDeck, c)

		/* Set random index element to last element of deck */
		d[randidx] = d[len(d)-1]

		/* Reallocate deck to be one element smaller */
		d = d[0 : len(d)-1]

		if len(d) == 0 {
			break
		}
	}

	return newDeck
}

/* return a card from the deck */
/* #TODO: make hit a function of hand(), that way we can know when you bust on a hit in this method
 * need to return the array of card[] representing the deck since we can't alter a slice by reference
 *  also need to return the card drawn as a card[] with a single element since a function cannot have multiple return
 *  types
 *      #TODO -- change so hit returns only the card drawn and operates on the deck without needing to return?
 */
func hit(d []card) ([]card, []card) {
	fmt.Println("size of deck", len(d))
	var c []card
	c = append(c, d[0])
	fmt.Println("card drawn", c)

	/* slice the array */
	d = d[1:]
	fmt.Println("new size of deck", len(d))

	return d, c
}

/* Hand */
type hand struct {
	cards []card
	// all potential scores the hand could have (dealing with aces)
	scores []int
}

func (h *hand) score() int {
	// initialize scores array for hand
	h.scores = []int{0}
	for _, c := range h.cards {
		h._updateScores(c)
	}

	/* Get "best" score for this hand. Best is defined as either:
	1. Highest total without breaking 21
	2. If no scores < 21, the lowest total score */
	score := 0
	for _, s := range h.scores {
		if s > score && s <= 21 {
			score = s
		}
	}

	// Didn't find a valid score that was <= 21
	// Return the smallest score > 21
	if score == 0 {
		for _, s := range h.scores {
			if score < s {
				score = s
			}
		}

	}

	return score
}

func (h *hand) _updateScores(c card) {
	score := c.numValue()
	/* If we see any Aces, capture potential scores if it is valued at 1 instead of 11 */
	var lowAceScores []int

	/* Account for Aces being valued at either 1 or 11
	Double the scores array in a hand, apply 11 to half, apply 1 to half */
	if c.isAce() {
		//current scores will need to be doubled, apply 11 to one half, 1 to the other
		lowAceScores = make([]int, len(h.scores))
		copy(lowAceScores, h.scores)
		for i := range lowAceScores {
			lowAceScores[i] += 1
		}

	}

	for i := range h.scores {
		h.scores[i] += score
	}

	if len(lowAceScores) > 0 {
		//make enough room for all score permutations
		newScores := make([]int, len(h.scores)*2)
		//populate our score permutations
		copy(newScores, h.scores)
		//add in any scores where Ace could be valued at 1
		copy(newScores[len(h.scores):], lowAceScores)
		//write to the hand
		h.scores = newScores
	}
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

func (h *hand) toString() {
	curScore := h.score()
	for _, c := range h.cards {
		c.toString()
	}
	fmt.Println("Score:", curScore)
	if h.busts() {
		fmt.Println("Busted!!")
	}
}

/* Helper functions to determine if a given score exists in a hands scores permutations
   and how many times is it seen */

func (h *hand) hasScore(score int) bool {
	timesScoreSeen := h.seenScore(score)

	var sawScore bool = false
	if timesScoreSeen > 0 {
		sawScore = true
	}

	return sawScore
}

func (h *hand) seenScore(score int) int {
	// how many elements in h.scores match our given score
	seen := 0
	for _, s := range h.scores {
		if s == score {
			seen++
		}
	}

	return seen
}

func (p *player) hit(g game) (card) {
    fmt.Println("in player* hit -- p:", p)

    return card{ suit: "K", value: "12" }

}

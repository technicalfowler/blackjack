package main

import (
	"testing"
)


func TestDeck(t *testing.T) {
	d := createDeck()	
	/* Decks should have 52 cards */
	if len(d) != 52 {
		t.Error(
			"For", "new deck",
			"expected", 52,
			"got",	len(d),
		)
	}
}

func TestCard(t *testing.T) {
	// 2 of Clubs -- numValue: 2
	a := card{suit: "C", value: "2"}
	if a.numValue() != 2 {
		t.Error("2 of clubs does not evaluate to 2")
	}
	// Does this card know it is not an Ace
	if a.isAce() {
		t.Error("2 of clubs thinks its an Ace")
	}

	//Ace of Spades -- numValue should be 2 or 11
	b := card{suit: "S", value: "A"}
	if b.numValue() != 11 {
		t.Error("Ace of spades does not evaluate to 11")
	}
	// does the card know its an Ace
	if !b.isAce() {
		t.Error("Ace of spades doesnt think its an Ace")
	}

	// King of Spades -- numValue: 10
	c := card{suit: "S", value: "K"}
	if c.numValue() != 10 {
		t.Error("King of Spades does not evaluate to 10")
	}
	
}

func TestHand(t *testing.T) {
	// 2 cards, 9 and 10 of spades
	var cards []card
	cards = append(cards, card{suit:"S", value: "10"})
	cards = append(cards, card{suit:"S", value: "9"})

	h := hand{ cards:cards}
	// score of hand should be 19
	if h.score() != 19 {
		t.Error("Hand does not correctly score",
			"Got", h.score(),
			"expected", 19)
	}

	// 3 cards, 9 and 10 of spades and Ace of spades
	h.cards = append(h.cards, card{suit:"S", value: "A"})

	if h.score() != 21 {
		t.Error("Hand doesnt score correctly\n",
			"Got", h.score(),
			"expected", 21)
	}

}

func TestBust(t *testing.T) {
	// 3 cards, all 10s
	var cards []card
	cards = append(cards, card{suit:"S", value: "10"})
	cards = append(cards, card{suit:"H", value: "10"})
	cards = append(cards, card{suit:"D", value: "10"})

	h := hand{ cards:cards}
	
	if h.score() != 30 {
		t.Error("Hand doesnt score correctly\n",
			"Got", h.score(),
			"Expected", 30)
	}	
	
	if !h.busts() { 
		t.Error("Hand doesn't bust")
	}
}

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
	var hand []card
	hand = append(hand, card{suit:"S", value: "10"})
	hand = append(hand, card{suit:"S", value: "9"})
	// score of hand should be 19
	if scoreHand(hand) != 19 {
		t.Error("Hand does not correctly score",
			"Got", scoreHand(hand),
			"expected", 19)
	}

	// 3 cards, 9 and 10 of spades and Ace of spades
	hand = append(hand, card{suit:"S", value: "A"})

	if scoreHand(hand) != 21 {
		t.Error("Hand doesnt score correctly\n",
			"Got", scoreHand(hand),
			"expected", 21)
	}
}

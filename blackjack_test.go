package main

import (
	"testing"
	"fmt"
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

	//Ace of Spades -- numValue should be 2 or 11
	b := card{suit: "S", value: "A"}
	if b.numValue() != 11 {
		t.Error("Ace of spades does not evaluate to 11")
	}
	// King of Spades -- numValue: 10
	c := card{suit: "S", value: "K"}
	if c.numValue() != 10 {
		t.Error("King of Spades does not evaluate to 10")
	}
	
	
}

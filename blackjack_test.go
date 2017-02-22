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

func TestDeckUniques(t *testing.T) {
	for i := 0; i < 10; i++ {
		d := createDeck()
		d = d.randomize()
		
		//iterate through each card in the deck to see if there are duplicates
		for j := 0; j < len(d); j++ {
			c := d[j]
			for _, sib := range d[j+1:] {
				if (sib.suit == c.suit && sib.value == c.value) {
					t.Error("Deck is not unique, found multiples of",	
						c.suit, c.value)
				}
			}
		}
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

	if h.score() != 20 {
		t.Error("Hand doesnt score correctly\n",
			"Got", h.score(),
			"expected", 20)
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

func TestHandMultipleAces(t *testing.T) {
	//3 cards, two Aces
	aa := card{suit:"S", value: "5"}
	ab := card{suit:"S", value: "A"}
	ac := card{suit:"H", value: "A"}
	cardsa := []card{aa,ab,ac}
	ha := hand{cards:cardsa}
	ha.score()

	/* Possible scores include:
		7: 5 + A(1) + A(1)
		17: 5 + A(11) + A(1)
		17: 5 + A(1) + A(11)
		27: 5 + A(11) + A(11)
	*/
	// 7, 17, 27
	for _, score := range []int{7, 17, 27} {
		if !ha.hasScore(score) {
			t.Error("hand does not calculate possible score",
				"Got", ha.hasScore(score),
				"expected", true)
		}
	}
	// Did we calculate 17 twice?
	if ha.seenScore(17) != 2 {
		t.Error("hand did not find score of 17 calculated in two separate permutations",	
			"Got", ha.seenScore(17),
			"expected", 2)
	}
	// "Best" score is 17 -- highest of three unique scores without going over 21
	if ha.score() != 17 {
		t.Error("hand did not calculate 17 as the best score",
			"Got", ha.score(),
			"expected", 17)
	}

	//4 cards, three Aces
	ba := card{suit:"S", value:"4"}
	bb := card{suit:"S", value:"A"}
	bc := card{suit:"D", value:"A"}
	bd := card{suit:"H", value:"A"}
	cardsb := []card{ba,bb,bc,bd}
	hb := hand{ cards:cardsb }
	hb.score()

	/* Possible scores include:
		7: 4 + A(1) + A(1) + A(1)
		17: 4 + A(11) + A(1) + A(1)
		27: 4 + A(11) + A(11) + A(1)
		37: 4 + A(11) + A(11) + A(11)
		27: 4 + A(11) + A(1) + A(11)
		17: 4 + A(1) + A(11) + A(1)
		27: 4 + A(1) + A(11) + A(11)
		17: 4 + A(1) + A(1) + A(11)	
	*/
	for _, score := range []int{7, 17, 27, 37} {
		if !hb.hasScore(score) {
			t.Error("Hand does not recognize score", score, "as a potential score",
				"Got", ha.hasScore(score),
				"expected", true)
		}
	}


	//4 cards, four Aces

	//5 cards, four Aces
}

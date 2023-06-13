package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	Clubs    string = "clubs"
	Spades   string = "spades"
	Hearts   string = "hearts"
	Diamonds string = "diamonds"
	Joker    string = "joker"
	Ace      int    = 1
	King     int    = 13
	Queen    int    = 12
	Jack     int    = 11
)

type Card struct {
	Value int
	Suit  string
}

func Shuffle(deck []Card) {
	rand.New(rand.NewSource(time.Now().UnixMilli()))
	rand.Shuffle(len(deck), func(i, j int) {
		(deck)[i], (deck)[j] = (deck)[j], (deck)[i]
	})
}

func NewMultideck(numberOfDecks int) []Card {
	multiDeck := make([]Card, 0, 100*numberOfDecks)
	for i := 0; i < numberOfDecks; i++ {
		multiDeck = append(multiDeck, New()...)
	}
	return multiDeck
}

func New() []Card {
	deck := make([]Card, 0, 100)
	for _, suit := range []string{Spades, Diamonds, Clubs, Hearts} {
		for _, value := range []int{Ace, 2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King} {
			deck = append(deck, Card{Value: value, Suit: suit})
		}
	}
	return deck
}

func PrintDeck(deck []Card) {
	for _, card := range deck {
		fmt.Println(card)
	}
}

func DrawCard(deck *[]Card) Card {
	topCardIndex := len(*deck) - 1
	drawnCard := (*deck)[topCardIndex]
	*deck = (*deck)[:topCardIndex]
	return drawnCard
}

func AddJokers(numberOfJokers int, deck *[]Card) {
	for i := 0; i < numberOfJokers; i++ {
		*deck = append(*deck, Card{Value: 0, Suit: Joker})
	}
}

func FilterOutValue(deck *[]Card, value int) {
	for i := len(*deck) - 1; i >= 0; i-- {
		if (*deck)[i].Value == value {
			*deck = append((*deck)[:i], (*deck)[i+1:]...)
		}
	}
}

func FilterOutSuit(deck *[]Card, suit string) {
	for i := len(*deck) - 1; i >= 0; i-- {
		if (*deck)[i].Suit == suit {
			*deck = append((*deck)[:i], (*deck)[i+1:]...)
		}
	}
}

func FilterOutCards(cardsToRemove []Card, deck *[]Card) {
	for _, card := range cardsToRemove {
		for i := len(*deck) - 1; i >= 0; i-- {
			if compareCardEquality(card, (*deck)[i]) {
				*deck = append((*deck)[:i], (*deck)[i+1:]...)
			}
		}
	}
}

func compareCardEquality(card1 Card, card2 Card) bool {
	return (card1.Suit == card2.Suit) && (card2.Value == card1.Value)
}

func SortDeckValueFirst(deck *[]Card) {
	derefDeck := *deck
	sort.Slice(derefDeck, func(i, j int) bool {
		if derefDeck[i].Value == derefDeck[j].Value {
			return sortBySuit(derefDeck, i, j)
		} else {
			return sortByValue(derefDeck, i, j)
		}
	})
}

func SortDeckSuitFirst(deck *[]Card) {
	derefDeck := *deck
	sort.Slice(derefDeck, func(i, j int) bool {
		if derefDeck[i].Suit != derefDeck[j].Suit {
			return sortBySuit(derefDeck, i, j)
		} else {
			return sortByValue(derefDeck, i, j)
		}
	})
}

func CustomSort(deck *[]Card, lessFn func(i, j int) bool) {
	sort.Slice((*deck), lessFn)
}

func sortBySuit(deck []Card, i, j int) bool {
	if deck[i].Suit == Spades && deck[j].Suit != Spades {
		return true
	} else if deck[i].Suit == Diamonds && (deck[j].Suit != Spades && deck[j].Suit != Diamonds) {
		return true
	} else if deck[i].Suit == Clubs && (deck[j].Suit == Hearts) {
		return true
	} else {
		return false
	}
}

func sortByValue(deck []Card, i, j int) bool {
	return deck[i].Value < deck[j].Value
}

package main

import (
	"github.com/lukemoran01/deck"
)

func main() {
	playingDeck := deck.NewMultideck(4)
	deck.Shuffle(playingDeck)
	deck.SortDeckSuitFirst(&playingDeck)
	deck.PrintDeck(playingDeck)

}

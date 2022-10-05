package main

import "fmt"

func main(){
	cards := newDeck()
	hand,remainingDeck := deal(cards,5)
	hand.print()
	remainingDeck.print()

	//writing into file
	cardsToSave :=newDeck()
	fmt.Println(cards.toString()) 
	cardsToSave.saveToFile("my_cards")

	//retrieving data from file
	cardsToFromFile := newDeckFromFile("my_ards")
	cardsToFromFile.print()


	//shuffling 
	cardShuffle := newDeck()
	cardShuffle.print()
	cardShuffle.shuffle()
	cardShuffle.print()
	
}


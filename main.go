package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("mycards")

	cards := newDeckFromFile("mycards")
	cards.print()

}

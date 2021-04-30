package main

import "fmt"

type deck []string

//Creating a new deck of cards
func newDeck() deck{
	cards:=deck{}

	cardSuits:=[]string{"Spades","Diamonds","Hearts","Clubs"}

	cardValues:=[]string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","Queen","King"}

	for _,suit:=range cardSuits{
		for _,value:=range cardValues{
			cards=append(cards,value+" of "+suit)
		}
	}

	return cards
}

//Printing the deck of cards
func(d deck) print(){
	for i,card:=range d{
		fmt.Println(i,card)
	}
}

//Dealing with the deck of cards 
func deal(d deck,handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}
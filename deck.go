package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

//converting deck(slice of strings) to single string
func(d deck) toString() string{
	return strings.Join([] string(d),",")
}

//saving deck of cards in a file in local drive
func(d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()),0666)
}

//retrieving deck of cards from saved file 
func newDeckFromFile(filename string) deck{
	bs,err:=ioutil.ReadFile(filename)

	if(err!=nil){
		fmt.Println("Error: ",err)
		os.Exit(1)
	}

	s:=strings.Split(string(bs),",")

	return deck(s)
}

//suffle deck of cards
func (d deck) suffle(){
	source:=rand.NewSource(time.Now().UnixNano())
	r:=rand.New(source)
	for i:=range d{
		newPosition:=r.Intn(len(d))

		d[i],d[newPosition]=d[newPosition],d[i]
	}
}

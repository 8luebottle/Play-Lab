package main

import (
	"fmt"
	"time"
)

// Switch statements
/*
	1. Basic switch
	2. Case List
	3. No condition
*/

const (
	tonnyBirthdayMonth = time.November
	jennyBirthdayMonth = time.May
	brownBirthdayMonth = time.June
)

var thisMonth = time.Now().Month()

func main() {
	fmt.Println("[Basic Switch]")
	isBirthdayMonth("JellyBelly")
	isBirthdayMonth("Jenny")
	isBirthdayMonth("Brown")
	fmt.Println()

	fmt.Println("[Case List]")
	firstBean := ISLAND
	secondBean := WATERMELON
	lastBean := PUDDING
	firstBean.Taste()
	secondBean.Taste()
	lastBean.Taste()
	fmt.Println()

	fmt.Println("[No Condition]")
	firstBean.Length(3)
	secondBean.Length(10)
}

// Basic Switch
// loop over all the cases from top to bottom.
func isBirthdayMonth(name string) {
	fmt.Printf("Is this %s's birhtday month?\n", name)

	switch name {
	case "Tonny":
		fmt.Println(thisMonth == tonnyBirthdayMonth)
	case "Jenny":
		fmt.Println(thisMonth == jennyBirthdayMonth)
	case "Brown":
		fmt.Println(thisMonth == brownBirthdayMonth)
	default:
		fmt.Println("Who is that?")
	}
}

type JellyBellyColor string

const (
	LEMON      JellyBellyColor = "lemon"
	COCONUT    JellyBellyColor = "coconut white"
	ISLAND     JellyBellyColor = "island punch"
	WATERMELON JellyBellyColor = "watermelon"
	CANDY      JellyBellyColor = "cotton candy"
	BERRY      JellyBellyColor = "berry blue"
	PUDDING    JellyBellyColor = "choco pudding"
)

// Case List
func (c JellyBellyColor) Taste() {
	fmt.Printf("%s taste : ", c)
	switch c {
	case LEMON, ISLAND, CANDY, BERRY:
		fmt.Print("Yummy yum yum !!!\n")
	case COCONUT, WATERMELON:
		fmt.Print("Yucky yuck yuck !!!!!!!!!!!!!\n")
	case PUDDING:
		fmt.Print("Literally taste like Poop...💩...\n")
	}
}

// No Condition
// No condition is same as switch true.
func (c JellyBellyColor) Length(m int) {
	switch {
	case len(c) < m:
		fmt.Println("less than")
	case len(c) > m:
		fmt.Println("greater than")
	case len(c) == m:
		fmt.Println("You got it!")
	}
}

package main

import (
	"fmt"
	"slices"
)

const (
	RichInitMoney = 30
	PoorInitMoney = 12
)

var poorPlayerCount = 0

type Player struct {
	money           int
	isRich          bool
	currentPosition int
	inventorySize   int
	inventory       []Item
}

func (p Player) initRich() {
	p.isRich = true
	p.money = RichInitMoney
	p.currentPosition = 1
	p.inventorySize = 9
	p.inventory = append(p.inventory, RiseCost)
}

func (p Player) initPoor() {
	p.isRich = false
	p.money = PoorInitMoney
	p.inventorySize = 2
	p.inventory = append(p.inventory, Knife)
	switch poorPlayerCount {
	case 0:
		p.currentPosition = 21
	case 1:
		p.currentPosition = 24
	default:
		p.currentPosition = 26
	}
	poorPlayerCount++
}

func (p Player) becomePoor() {
	p.isRich = false
}

func (p Player) getMovement() []int {
	closest := mapSlice[p.currentPosition][1:]
	if !p.isRich {
		return closest
	}

	var movement []int

	for _, i := range closest {
		movement = append(movement, mapSlice[closest[i]][1:]...)
	}

	return movement
}

func (p Player) getCurrentSpaceState() {
	switch mapSlice[p.currentPosition][0] {
	case BlankSpace:
		fmt.Println("BlankSpace")
	case BadSpace:
		fmt.Println("BadSpace")
	case GoodSpace:
		fmt.Println("GoodSpace")
	case ShopSpace:
		fmt.Println("ShopSpace")
	case GambleSpace:
		fmt.Println("GambleSpace")
	}
}

func (p Player) move(dest int) {
	if slices.Contains(p.getMovement(), dest) {
		p.currentPosition = dest
	} else {
		fmt.Println("You can`t move here")
	}
}

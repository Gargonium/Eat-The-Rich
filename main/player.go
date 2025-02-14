package main

import (
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
	//inventory       []Item
}

func (p Player) initRich() Player {
	p.isRich = true
	p.money = RichInitMoney
	p.currentPosition = 1
	p.inventorySize = 9
	//p.inventory = append(p.inventory, RiseCost)
	return p
}

func (p Player) initPoor() Player {
	p.isRich = false
	p.money = PoorInitMoney
	p.inventorySize = 2
	//p.inventory = append(p.inventory, Knife)
	switch poorPlayerCount {
	case 0:
		p.currentPosition = 21
	case 1:
		p.currentPosition = 24
	default:
		p.currentPosition = 26
	}
	poorPlayerCount++
	return p
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

func (p Player) getCurrentSpaceState() string {
	switch mapSlice[p.currentPosition][0] {
	case BlankSpace:
		return "BlankSpace"
	case BadSpace:
		return "BadSpace"
	case GoodSpace:
		return "GoodSpace"
	case ShopSpace:
		return "ShopSpace"
	case GambleSpace:
		return "GambleSpace"
	default:
		return "Error, unknown space"
	}
}

func (p Player) move(dest int, err error) (string, Player) {
	if slices.Contains(p.getMovement(), dest) {
		p.currentPosition = dest
		return "", p
	} else {
		return "You can`t move here", p
	}
}

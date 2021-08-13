package main

import "fmt"

//607,2
//7394.12
//7394.12

//               Gallons    Pounds
//Center1      = 4530.000   30351
//LeftMain     = 1780.000   11926
//RightMain    = 1780.000   11926

const wingTanksMax float32 = 11926
const centerTanksMax float32 = 30351

type Fuel struct {
	centerTank float32
	leftTank   float32
	rightTank  float32
}

func NewFuel() Fuel {
	var fuel = Fuel{0, 0, 0}
	return fuel
}

func (fuel *Fuel) Load(payload int64) {

	var eqWeight = float32(payload) / 3

	if eqWeight > wingTanksMax {
		fuel.leftTank = wingTanksMax
		fuel.rightTank = wingTanksMax
		fuel.centerTank = float32(payload) - (wingTanksMax * 2)
	} else {
		fuel.leftTank = eqWeight
		fuel.rightTank = eqWeight
		fuel.centerTank = eqWeight
	}
}

func (fuel Fuel) Print() {
	var cprValue = (fuel.centerTank * 100) / centerTanksMax
	var wprValue = (fuel.leftTank * 100) / wingTanksMax
	fmt.Println("------------Fuel weights-----------------")
	fmt.Printf("Center: %.2f (%.2f%%)\n", fuel.centerTank, cprValue)
	fmt.Printf("Left  : %.2f (%.2f%%)\n", fuel.leftTank, wprValue)
	fmt.Printf("Right : %.2f (%.2f%%)\n", fuel.rightTank, wprValue)
}

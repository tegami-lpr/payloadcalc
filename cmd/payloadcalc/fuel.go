package main

import "fmt"

//607,2
//7394.12
//7394.12

//               Gallons    Pounds
//Center1      = 4530.000   30351
//LeftMain     = 1780.000   11926
//RightMain    = 1780.000   11926

type Fuel struct {
	centerTank int64
	leftTank float32
	rightTank float32
}

func NewFuel() Fuel {
	var fuel = Fuel{0, 0, 0}
	return fuel
}

func (fuel *Fuel) Load(payload int64) {
	var wingTanksMax int64 = 11926
	var centerWeight = payload - (wingTanksMax * 2)
	fuel.centerTank = centerWeight
	var leftRight = float32(payload - centerWeight) / 2
	fuel.leftTank = leftRight
	fuel.rightTank = leftRight
}

func (fuel Fuel) Print() {
	fmt.Println("------------Fuel weights-----------------")
	fmt.Printf("Center: %d\n", fuel.centerTank)
	fmt.Printf("Left  : %f\n", fuel.leftTank)
	fmt.Printf("Right : %f\n", fuel.rightTank)
}
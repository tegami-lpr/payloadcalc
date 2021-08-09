package main

import (
	"fmt"
	"os"
	"strconv"
)



func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tpayloadcalc BlockOfFuel FuelAtDestination")
	fmt.Println("\tAll weight on lbs")
}

func main() {
	//All weight in lbs
	var landingWeight int64 = 150000
	var blockOfFuel int64 = 16000 //must be readed from args
	var fuelAtDestination int64 = 1000 //must be readed from args
	var emptyWeight int64 = 100700
	var err error

	if len(os.Args) < 2 {
		fmt.Println("No filename given. Exiting.")
		PrintUsage()
		os.Exit(0)
	}

	blockOfFuel, err = strconv.ParseInt(os.Args[1], 10 ,32)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fuelAtDestination, err = strconv.ParseInt(os.Args[2], 10 ,32)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	var takeOfWeight = landingWeight + (blockOfFuel - fuelAtDestination)
	var payload = takeOfWeight - emptyWeight - blockOfFuel

	if payload < 0 {
		fmt.Println("Unable calc payload")
		os.Exit(0)
	}

	fmt.Printf("BoF: %d\n", blockOfFuel)
	fmt.Printf("FaD: %d\n", fuelAtDestination)
	fmt.Printf("ToW: %d\n", takeOfWeight)
	fmt.Printf("Payload: %d\n", payload)

	var cargo = NewCargo()
	cargo.Load(payload)
	cargo.PrintStations()

	var fuel = NewFuel()
	fuel.Load(blockOfFuel)
	fuel.Print()

	os.Exit(1)
}
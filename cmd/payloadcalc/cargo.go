package main

import "fmt"

//station_load.0= 9503, -36.82, 0.00, 0.00, Aft Cargo
//station_load.1= 9503, -23.37, 0.00, 0.00, Container 11 12
//station_load.2= 9503, -5.08,  0.00, 0.00, Container 9 10
//station_load.3= 9503, 13.21,  0.00, 0.00, Container 7 8
//station_load.4= 9503, 18.27,  0.00, 0.00, Container 5 6
//station_load.5= 9503, 31.51,  0.00, 0.00, Container 3 4
//station_load.6= 9503, 44.29,  0.00, 0.00, Container 1 2
//station_load.7= 9503, 11.98,  0.00, 0.00, Fwd Cargo


type Cargo struct {
	stations [8]int64
}

func NewCargo() Cargo {
	var cargo = Cargo{}
	for i := range cargo.stations {
		cargo.stations[i] = 0
	}
	return cargo
}

func (cargo *Cargo) Load(payload int64) {
	var maxPayload int64 = 9503
	for i := range cargo.stations {
		if payload <= 0 {
			break
		}
		var stationPayload = payload
		if stationPayload > maxPayload {
			stationPayload = maxPayload
		}
		cargo.stations[i] = stationPayload
		payload = payload - stationPayload
	}
}

func (cargo Cargo)PrintStations() {
	fmt.Println("---------Container weights---------------")
	fmt.Printf("Container 1 2  : %d\n", cargo.stations[6])
	fmt.Printf("Container 3 4  : %d\n", cargo.stations[5])
	fmt.Printf("Container 5 6  : %d\n", cargo.stations[4])
	fmt.Printf("Container 7 8  : %d\n", cargo.stations[3])
	fmt.Printf("Container 9 10 : %d\n", cargo.stations[2])
	fmt.Printf("Container 11 12: %d\n", cargo.stations[1])
	fmt.Printf("Fwd Cargo: %d\n", cargo.stations[7])
	fmt.Printf("Aft Cargo: %d\n", cargo.stations[0])
}
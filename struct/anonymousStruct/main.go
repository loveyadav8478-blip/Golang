package main

type car struct{
	brand string
	model string
	name string
	mileage string
	doors int


	wheel struct {
		radius int
		material string
	}
}

var myCar = car{
	brand: "LAND ROVER",
	model: "VELAR",
	name: "RANGE ROVER",
	mileage: "100",
	doors: 6,

	wheel:struct{radius int; material string}{
		radius: 10,
		material: "ALLOY",
	},
}
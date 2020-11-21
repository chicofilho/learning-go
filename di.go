// This has the unique purppose of using a DI under go
// The program might not make a lot of sense, it's a very small subset
// And the user case is totally imaginary
package learninggo

import (
	"time"
)

// this is the entity that will be injected
type parkingRule struct{}

func (pr *parkingRule) parkingPrice(cr carRecord) int {
	now := time.Now()
	diff := now.Sub(cr.entryTime).Hours()
	if diff > 1 {
		return 5 + int(diff)*2
	}
	return 5
}

// And an interface to add any entity that implements interface requested methods
type parkingRuler interface {
	parkingPrice(cr carRecord) int
}

// A very simple car record for a parking lot
type carRecord struct {
	id        int
	entryTime time.Time // time when the car was parked
}

// A parking lot, the entity that will receive the injected dep.
type parkingLot struct {
	freeSpots int
	balance   int
	cars      map[int]carRecord
	config    parkingRuler
}

// Just a small fabric method to initialize parking lots
func Create(freeSpots int, config parkingRuler) *parkingLot {
	return &parkingLot{config: config, freeSpots: 2, cars: make(map[int]carRecord)}
}

// Method to park a car
func (pl *parkingLot) Park(cr carRecord) {
	if pl.freeSpots == 0 {
		panic("No space left")
	}
	pl.freeSpots--
	pl.cars[cr.id] = cr
}

// Method to make a car leave the parking lot
// Price is applied
func (pl *parkingLot) Leave(carId int) {
	car := pl.cars[carId]
	pl.balance += pl.config.parkingPrice(car)
	pl.freeSpots++
	delete(pl.cars, car.id)
}

// getters
func (pl *parkingLot) GetFreeSpots() int {
	return pl.freeSpots
}

func (pl *parkingLot) GetBalance() int {
	return pl.balance
}

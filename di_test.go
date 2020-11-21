package learninggo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Fake ParkingRuler to test DI
type fakeParkingRule struct{}

func (pr *fakeParkingRule) parkingPrice(cr carRecord) int {
	return 1
}

// testing the injection of a fake Parking Ruler
func TestMockedParkingRuler(t *testing.T) {
	fpr := fakeParkingRule{}
	plot := Create(2, &fpr)

	plot.Park(carRecord{1, time.Now().Add(-time.Hour * 2)})
	assert.Equal(t, plot.GetFreeSpots(), 1)
	assert.Equal(t, plot.GetBalance(), 0)
	plot.Leave(1)
	assert.Equal(t, plot.GetBalance(), 1)
}

// injection of the default parking ruler
func TestParkingRuler(t *testing.T) {
	prl := parkingRule{}
	plot := Create(2, &prl)

	// parking one car
	plot.Park(carRecord{2, time.Now().Add(-time.Hour * 5)})
	assert.Equal(t, plot.GetFreeSpots(), 1)

	// parking another
	plot.Park(carRecord{1, time.Now().Add(-time.Hour * 2)})
	assert.Equal(t, plot.GetFreeSpots(), 0)

	// start to make the cars leave. Balance starts to grow
	assert.Equal(t, plot.GetBalance(), 0)
	plot.Leave(1)
	assert.Equal(t, plot.GetBalance(), 9)
	plot.Leave(2)
	assert.Equal(t, plot.GetBalance(), 24)
}

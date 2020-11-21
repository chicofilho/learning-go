package learninggo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple Testing. It doesn't test big entries, as it's not the intended approach here
func TestSolution(t *testing.T) {
	assert.Equal(t, minMinutes([][]int{{0}}), 0)
	assert.Equal(t, minMinutes([][]int{{0, 2}}), 0)
	assert.Equal(t, minMinutes([][]int{{0, 0}}), 0)
	assert.Equal(t, minMinutes([][]int{{2, 2}, {2, 2}}), 0)
	assert.Equal(t, minMinutes([][]int{{2, 1}, {2, 2}}), 1)
	assert.Equal(t, minMinutes([][]int{{2, 1}, {1, 2}}), 1)
	assert.Equal(t, minMinutes([][]int{{2, 1}, {1, 1}}), 2)
	assert.Equal(t, minMinutes([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}), 4)
	assert.Equal(t, minMinutes([][]int{{2, 1, 0}, {0, 1, 1}, {1, 0, 1}}), -1)
	assert.Equal(t, minMinutes([][]int{{2, 1, 1}, {0, 1, 1}, {1, 1, 1}}), 4)
	assert.Equal(t, minMinutes([][]int{{2, 1, 1}, {1, 1, 1}, {1, 1, 2}}), 2)
	assert.Equal(t, minMinutes([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}), -1)
}

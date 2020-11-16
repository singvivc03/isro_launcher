package counter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldVerifyTimerAlwaysReturnsOne(t *testing.T) {
	counter := timer(1 * time.Second)
	val := <-counter
	assert.Equal(t, val, 1)

	counter = timer(1 * time.Second)
	val = <-counter
	assert.Equal(t, val, 1)
}

func TestShouldStartCounter(t *testing.T) {

}

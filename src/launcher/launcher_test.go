package launcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLaunchSatelliteInAGroupOfTwo(t *testing.T) {
	launcher := New("LS1")
	signal := launcher.Launch(1, 2)
	assert.Equal(t, <-signal, 3)
}

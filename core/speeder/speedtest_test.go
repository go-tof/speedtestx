package speeder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallSpeedTestCommand(t *testing.T) {
	command, err := CallSpeedTestCommand()
	assert.NoError(t, err)
	assert.NotNil(t, command)
}

package speeder

import (
	"os/exec"
)

// CallSpeedTestCommand call os speeder cli application.
func CallSpeedTestCommand() ([]byte, error) {
	command := exec.Command("speedtest", "-A", "--progress=no", "--format=json")

	// Command run and catch output information.
	output, err := command.Output()
	if err != nil {
		return nil, err
	}

	return output, nil
}

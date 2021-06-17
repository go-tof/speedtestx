package speedtest

import (
	"fmt"

	"github.com/go-tof/speedtestx/core/speeder"
	"github.com/urfave/cli/v2"
)

// ActionFuncForSpeedtest Screenshot handler.
var ActionFuncForSpeedtest = func() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		output, err := speeder.CallSpeedTestCommand()
		if err != nil {
			fmt.Fprintf(ctx.App.Writer, fmt.Sprintf("Call os envariment speedtest has error: %s\n", err))
			return nil
		}

		// Gets the first parameter.
		// all := ctx.Bool("all")

		// initialize speeder with bytes data
		speed, err := speeder.NewSpeederWithBytes(output)
		if err != nil {
			fmt.Fprintf(ctx.App.Writer, fmt.Sprintf("Create speedtest network check result."))
		}

		fmt.Fprintf(ctx.App.Writer, fmt.Sprintf("%v", &speed))

		return nil
	}
}()

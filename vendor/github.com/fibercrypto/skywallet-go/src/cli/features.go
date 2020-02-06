package cli

import (
	messages "github.com/fibercrypto/skywallet-protob/go"
	gcli "github.com/urfave/cli"
)

func featuresCmd() gcli.Command {
	name := "features"
	return gcli.Command{
		Name:         name,
		Usage:        "Ask the device Features.",
		Description:  "",
		OnUsageError: onCommandUsageError(name),
		Flags: []gcli.Flag{
			gcli.StringFlag{
				Name:   "deviceType",
				Usage:  "Device type to send instructions to, hardware wallet (USB) or emulator.",
				EnvVar: "DEVICE_TYPE",
			},
		},
		Action: func(c *gcli.Context) {
			sq, err := createDevice(c.String("deviceType"))
			if err != nil {
				return
			}
			msg, err := sq.GetFeatures()
			handleFinalResponse(msg, err, "unable to get features", messages.MessageType_MessageType_Features)
		},
	}
}

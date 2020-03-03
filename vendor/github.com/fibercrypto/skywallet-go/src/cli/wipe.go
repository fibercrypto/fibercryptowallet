package cli

import (
	messages "github.com/fibercrypto/skywallet-protob/go"

	gcli "github.com/urfave/cli"
)

func wipeCmd() gcli.Command {
	name := "wipe"
	return gcli.Command{
		Name:         name,
		Usage:        "Ask the device to wipe clean all the configuration it contains.",
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
			msg, err := sq.Wipe()
			handleFinalResponse(msg, err, "unable to wipe device", messages.MessageType_MessageType_Success)
		},
	}
}

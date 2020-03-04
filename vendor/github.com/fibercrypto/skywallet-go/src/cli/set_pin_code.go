package cli

import (
	gcli "github.com/urfave/cli"

	messages "github.com/fibercrypto/skywallet-protob/go"
)

func setPinCode() gcli.Command {
	name := "setPinCode"
	return gcli.Command{
		Name:        name,
		Usage:       "Configure a PIN code on a device.",
		Description: "",
		Flags: []gcli.Flag{
			gcli.StringFlag{
				Name:   "deviceType",
				Usage:  "Device type to send instructions to, hardware wallet (USB) or emulator.",
				EnvVar: "DEVICE_TYPE",
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			sq, err := createDevice(c.String("deviceType"))
			if err != nil {
				return
			}
			msg, err := sq.ChangePin(new(bool))
			handleFinalResponse(msg, err, "unable to change pin", messages.MessageType_MessageType_Success)
		},
	}
}

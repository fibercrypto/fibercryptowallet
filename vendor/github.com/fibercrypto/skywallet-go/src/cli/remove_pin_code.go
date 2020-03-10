package cli

import (
	gcli "github.com/urfave/cli"

	messages "github.com/fibercrypto/skywallet-protob/go"
)

func removePinCode() gcli.Command {
	name := "removePinCode"
	return gcli.Command{
		Name:        name,
		Usage:       "Remove a PIN code on a device.",
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
			removePin := new(bool)
			*removePin = true
			sq, err := createDevice(c.String("deviceType"))
			if err != nil {
				return
			}
			msg, err := sq.ChangePin(removePin)
			handleFinalResponse(msg, err, "unable to remove pin", messages.MessageType_MessageType_Success)
		},
	}
}

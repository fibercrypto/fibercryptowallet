package cli

import (
	messages "github.com/fibercrypto/skywallet-protob/go"

	gcli "github.com/urfave/cli"
)

func setMnemonicCmd() gcli.Command {
	name := "setMnemonic"
	return gcli.Command{
		Name:        name,
		Usage:       "Configure the device with a mnemonic.",
		Description: "",
		Flags: []gcli.Flag{
			gcli.StringFlag{
				Name:  "mnemonic",
				Usage: "Mnemonic that will be stored in the device to generate addresses.",
			},
			gcli.StringFlag{
				Name:   "deviceType",
				Usage:  "Device type to send instructions to, hardware wallet (USB) or emulator.",
				EnvVar: "DEVICE_TYPE",
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			mnemonic := c.String("mnemonic")
			sq, err := createDevice(c.String("deviceType"))
			if err != nil {
				return
			}
			msg, err := sq.SetMnemonic(mnemonic)
			handleFinalResponse(msg, err, "unable to set mnemonic", messages.MessageType_MessageType_Success)
		},
	}
}

package cli

import (
	messages "github.com/fibercrypto/skywallet-protob/go"

	gcli "github.com/urfave/cli"
)

func generateMnemonicCmd() gcli.Command {
	name := "generateMnemonic"
	return gcli.Command{
		Name:        name,
		Usage:       "Ask the device to generate a mnemonic and configure itself with it.",
		Description: "",
		Flags: []gcli.Flag{
			gcli.BoolFlag{
				Name:  "usePassphrase",
				Usage: "Configure a passphrase",
			},
			gcli.IntFlag{
				Name:  "wordCount",
				Usage: "Use a specific (12 | 24) number of words for the Mnemonic",
				Value: 12,
			},
			gcli.StringFlag{
				Name:   "deviceType",
				Usage:  "Device type to send instructions to, hardware wallet (USB) or emulator.",
				EnvVar: "DEVICE_TYPE",
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			usePassphrase := c.Bool("usePassphrase")
			wordCount := uint32(c.Uint64("wordCount"))
			sq, err := createDevice(c.String("deviceType"))
			if err != nil {
				return
			}
			msg, err := sq.GenerateMnemonic(wordCount, usePassphrase)
			handleFinalResponse(msg, err, "unable to generate mnemonic", messages.MessageType_MessageType_Success)
		},
	}
}

package cli

import (
	gcli "github.com/urfave/cli"

	messages "github.com/fibercrypto/skywallet-protob/go"
)

func signMessageCmd() gcli.Command {
	name := "signMessage"
	return gcli.Command{
		Name:        name,
		Usage:       "Ask the device to sign a message using the secret key at given index.",
		Description: "",
		Flags: []gcli.Flag{
			gcli.IntFlag{
				Name:  "addressIndex",
				Value: 0,
				Usage: "Index of the address that will issue the signature. Assume 0 if not set.",
			},
			gcli.StringFlag{
				Name:  "message",
				Usage: "The message that the signature claims to be signing.",
			},
			gcli.StringFlag{
				Name:   "deviceType",
				Usage:  "Device type to send instructions to, hardware wallet (USB) or emulator.",
				EnvVar: "DEVICE_TYPE",
			},
			gcli.StringFlag{
				Name:  "walletType",
				Usage: "Wallet type. Types are \"deterministic\" or \"bip44\"",
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			addressIndex := c.Int("addressIndex")
			message := c.String("message")
			walletType := c.String("walletType")
			sq, err := createDevice(c.String("deviceType"))
			if err != nil {
				return
			}
			msg, err := sq.SignMessage(1, addressIndex, message, walletType)
			handleFinalResponse(msg, err, "unable to sign message", messages.MessageType_MessageType_ResponseSkycoinSignMessage)
		},
	}
}

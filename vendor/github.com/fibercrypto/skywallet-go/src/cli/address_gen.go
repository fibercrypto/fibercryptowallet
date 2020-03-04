package cli

import (
	messages "github.com/fibercrypto/skywallet-protob/go"

	gcli "github.com/urfave/cli"
)

func addressGenCmd() gcli.Command {
	name := "addressGen"
	return gcli.Command{
		Name:        name,
		Usage:       "Generate skycoin addresses using the firmware",
		Description: "",
		Flags: []gcli.Flag{
			gcli.IntFlag{
				Name:  "addressN",
				Value: 1,
				Usage: "Number of addresses to generate. Assume 1 if not set.",
			},
			gcli.IntFlag{
				Name:  "startIndex",
				Value: 0,
				Usage: "Index where deterministic key generation will start from. Assume 0 if not set.",
			},
			gcli.BoolFlag{
				Name:  "confirmAddress",
				Usage: "If requesting one address it will be sent only if user confirms operation by pressing device's button.",
			},
			gcli.StringFlag{
				Name:   "deviceType",
				Usage:  "Device type to send instructions to, hardware wallet (USB) or emulator.",
				EnvVar: "DEVICE_TYPE",
			},
			gcli.StringFlag{
				Name:     "walletType",
				Usage:    "Wallet type. Types are \"deterministic\" or \"bip44\"",
				Required: true,
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			walletType := c.String("walletType")
			addressN := c.Int("addressN")
			startIndex := c.Int("startIndex")
			confirmAddress := c.Bool("confirmAddress")
			sq, err := createDevice(c.String("deviceType"))
			if err != nil {
				return
			}
			msg, err := sq.AddressGen(uint32(addressN), uint32(startIndex), confirmAddress, walletType)
			handleFinalResponse(msg, err, "unable to get address", messages.MessageType_MessageType_ResponseSkycoinAddress)
		},
	}
}

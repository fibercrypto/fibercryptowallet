package cli

import (
	gcli "github.com/urfave/cli"

	messages "github.com/fibercrypto/skywallet-protob/go"
)

func applySettingsCmd() gcli.Command {
	name := "applySettings"
	return gcli.Command{
		Name:        name,
		Usage:       "Apply settings.",
		Description: "",
		Flags: []gcli.Flag{
			gcli.StringFlag{
				Name:  "usePassphrase",
				Usage: "Configure a passphrase (true or false)",
			},
			gcli.StringFlag{
				Name:  "label",
				Usage: "Configure a device label",
			},
			gcli.StringFlag{
				Name:   "deviceType",
				Usage:  "Device type to send instructions to, hardware wallet (USB) or emulator.",
				EnvVar: "DEVICE_TYPE",
			},
			gcli.StringFlag{
				Name:  "language",
				Usage: "Configure a device language",
				Value: "",
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			passphrase := c.String("usePassphrase")
			label := c.String("label")
			language := c.String("language")
			usePassphrase, _err := parseBool(passphrase)
			if _err != nil {
				log.Errorln("Valid values for usePassphrase are true or false")
				return
			}
			sq, err := createDevice(c.String("deviceType"))
			if err != nil {
				return
			}
			msg, err := sq.ApplySettings(usePassphrase, label, language)
			handleFinalResponse(msg, err, "unable to apply settings", messages.MessageType_MessageType_Success)
		},
	}
}

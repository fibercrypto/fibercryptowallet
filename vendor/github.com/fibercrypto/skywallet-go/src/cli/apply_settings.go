package cli

import (
	"fmt"
	"os"
	"runtime"

	gcli "github.com/urfave/cli"

	messages "github.com/fibercrypto/skywallet-protob/go"

	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
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

			device := skyWallet.NewDevice(skyWallet.DeviceTypeFromString(c.String("deviceType")))
			if device == nil {
				return
			}
			defer device.Close()

			if os.Getenv("AUTO_PRESS_BUTTONS") == "1" && device.Driver.DeviceType() == skyWallet.DeviceTypeEmulator && runtime.GOOS == "linux" {
				err := device.SetAutoPressButton(true, skyWallet.ButtonRight)
				if err != nil {
					log.Error(err)
					return
				}
			}

			usePassphrase, _err := parseBool(passphrase)
			if _err != nil {
				log.Errorln("Valid values for usePassphrase are true or false")
				return
			}
			msg, err := device.ApplySettings(usePassphrase, label, language)
			if err != nil {
				log.Error(err)
				return
			}

			for msg.Kind != uint16(messages.MessageType_MessageType_Failure) && msg.Kind != uint16(messages.MessageType_MessageType_Success) {
				if msg.Kind == uint16(messages.MessageType_MessageType_ButtonRequest) {
					msg, err = device.ButtonAck()
					if err != nil {
						log.Error(err)
						return
					}
					continue
				}

				if msg.Kind == uint16(messages.MessageType_MessageType_PinMatrixRequest) {
					var pinEnc string
					fmt.Printf("PinMatrixRequest response: ")
					fmt.Scanln(&pinEnc)
					pinAckResponse, err := device.PinMatrixAck(pinEnc)
					if err != nil {
						log.Error(err)
						return
					}
					log.Infof("PinMatrixAck response: %s", pinAckResponse)
					continue
				}
			}

			if msg.Kind == uint16(messages.MessageType_MessageType_Failure) {
				failMsg, err := skyWallet.DecodeFailMsg(msg)
				if err != nil {
					log.Error(err)
					return
				}
				fmt.Println(failMsg)
				return
			}

			if msg.Kind == uint16(messages.MessageType_MessageType_Success) {
				successMsg, err := skyWallet.DecodeSuccessMsg(msg)
				if err != nil {
					log.Error(err)
					return
				}
				fmt.Println(successMsg)
				return
			}
		},
	}
}

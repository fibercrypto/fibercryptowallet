package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/gogo/protobuf/proto"
	gcli "github.com/urfave/cli"

	messages "github.com/fibercrypto/skywallet-protob/go"

	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
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

			msg, err := device.GetFeatures()
			if err != nil {
				log.Error(err)
				return
			}

			switch msg.Kind {
			case uint16(messages.MessageType_MessageType_Features):
				features := &messages.Features{}
				err = proto.Unmarshal(msg.Data, features)
				if err != nil {
					log.Error(err)
					return
				}

				enc := json.NewEncoder(os.Stdout)
				if err = enc.Encode(features); err != nil {
					log.Errorln(err)
					return
				}
				ff := skyWallet.NewFirmwareFeatures(uint64(*features.FirmwareFeatures))
				if err := ff.Unmarshal(); err != nil {
					log.Errorln(err)
					return
				}
				log.Printf("\n\nFirmware features:\n%s", ff)
			// TODO: figure out if this method can even return success or failure msg.
			case uint16(messages.MessageType_MessageType_Failure), uint16(messages.MessageType_MessageType_Success):
				msgData, err := skyWallet.DecodeSuccessOrFailMsg(msg)
				if err != nil {
					log.Error(err)
					return
				}

				fmt.Println(msgData)
			default:
				log.Errorf("received unexpected message type: %s", messages.MessageType(msg.Kind))
			}
		},
	}
}

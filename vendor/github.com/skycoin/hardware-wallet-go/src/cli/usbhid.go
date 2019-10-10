package cli

import (
	gcli "github.com/urfave/cli"

	skyWallet "github.com/skycoin/hardware-wallet-go/src/skywallet"
)

func getUsbDetails() gcli.Command {
	name := "getUsbDetails"
	return gcli.Command{
		Name:         name,
		Usage:        "Ask host usb about details for the hardware wallet",
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

			infos, err := device.GetUsbInfo()
			if err != nil {
				log.Errorln(err)
			}
			for infoIdx := range infos {
				log.Infoln("-----------------------------------------")
				if infos[infoIdx].VendorID == skyWallet.SkycoinVendorID {
					log.Printf("%-13d%-5s%s", infos[infoIdx].VendorID, "==>", "Skycoin Foundation")
				}
				if infos[infoIdx].ProductID == skyWallet.SkycoinHwProductID {
					log.Printf("%-13d%-5s%s", infos[infoIdx].ProductID, "==>", "Hardware Wallet")
				}
				log.Printf("%-13s%-5s%s", "Device path", "==>", infos[infoIdx].Path)
			}
		},
	}
}

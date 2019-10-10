package cli

import (
	gcli "github.com/urfave/cli"

	skyWallet "github.com/skycoin/hardware-wallet-go/src/skywallet"
)

func getRawEntropyCmd() gcli.Command {
	name := "getRawEntropy"
	return gcli.Command{
		Name:  name,
		Usage: "Get device raw internal entropy and write it down to a file",
		Action: func(c *gcli.Context) {
			entropyBytes := uint32(c.Int("entropyBytes"))
			outFile := c.String("outFile")
			if len(outFile) == 0 {
				log.Error("outFile is mandatory")
				return
			}

			device := skyWallet.NewDevice(skyWallet.DeviceTypeFromString(c.String("deviceType")))
			if device == nil {
				return
			}
			defer device.Close()

			log.Infoln("Getting raw entropy from device")
			if err := device.SaveDeviceEntropyInFile(outFile, entropyBytes, skyWallet.MessageDeviceGetRawEntropy); err != nil {
				log.Error(err)
				return
			}
		},
		OnUsageError: onCommandUsageError(name),
		Subcommands:  nil,
		Flags: []gcli.Flag{
			gcli.IntFlag{
				Name:  "entropyBytes",
				Value: 1048576,
				Usage: "Total number of how many bytes of raw entropy to read.",
			},
			gcli.StringFlag{
				Name:  "outFile",
				Usage: `File path to write out the raw entropy buffers, a "-" set the file to stdout.`,
				Value: "-",
			},
			gcli.StringFlag{
				Name:   "deviceType",
				Usage:  "Device type to send instructions to, hardware wallet (USB) or emulator.",
				EnvVar: "DEVICE_TYPE",
			},
		},
		SkipFlagParsing: false,
		SkipArgReorder:  false,
		HideHelp:        false,
		Hidden:          false,
	}
}

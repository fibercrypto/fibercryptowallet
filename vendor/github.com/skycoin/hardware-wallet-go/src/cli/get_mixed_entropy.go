package cli

import (
	gcli "github.com/urfave/cli"

	skyWallet "github.com/skycoin/hardware-wallet-go/src/skywallet"
)

func getMixedEntropyCmd() gcli.Command {
	name := "getMixedEntropy"
	return gcli.Command{
		Name:  name,
		Usage: "Get device internal mixed entropy and write it down to a file",
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

			log.Infoln("Getting mixed entropy from device")
			if err := device.SaveDeviceEntropyInFile(outFile, entropyBytes, skyWallet.MessageDeviceGetMixedEntropy); err != nil {
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
				Usage: "Total number of how many bytes of mixed entropy to read.",
			},
			gcli.StringFlag{
				Name:  "outFile",
				Usage: `File path to write out the mixed entropy buffers, a "-" set the file to stdout.`,
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

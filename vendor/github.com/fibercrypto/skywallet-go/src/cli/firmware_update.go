package cli

import (
	"crypto/sha256"
	"fmt"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"io/ioutil"

	gcli "github.com/urfave/cli"
)

func firmwareUpdate() gcli.Command {
	name := "firmwareUpdate"
	return gcli.Command{
		Name:        name,
		Usage:       "Update device's firmware.",
		Description: "",
		Flags: []gcli.Flag{
			gcli.StringFlag{
				Name:  "f, file",
				Usage: "path to the firmware .bin file",
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			filePath := c.String("file")
			fmt.Printf("File : %s\n", filePath)
			firmware, err := ioutil.ReadFile(filePath)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Hash: %x\n", sha256.Sum256(firmware[0x100:]))
			sq, err := createDevice(skywallet.DeviceTypeUSB.String())
			if err != nil {
				return
			}
			msg, err := sq.UploadFirmware(firmware, sha256.Sum256(firmware[0x100:]))
			handleFinalResponse(msg, err, "unable to update firmware", messages.MessageType_MessageType_Success)
			if err != nil {
				log.Error(err)
				return
			}
		},
	}
}

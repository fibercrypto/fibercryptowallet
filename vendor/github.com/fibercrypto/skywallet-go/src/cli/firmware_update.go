package cli

import (
	"crypto/sha256"
	"fmt"
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
			sq, err := createDevice(c.String("deviceType"))
			if err != nil {
				return
			}
			err = sq.FirmwareUpload(firmware, sha256.Sum256(firmware[0x100:]))
			if err != nil {
				log.Error(err)
				return
			}
		},
	}
}

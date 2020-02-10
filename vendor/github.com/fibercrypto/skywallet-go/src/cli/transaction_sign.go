package cli

import (
	"fmt"
	"github.com/gogo/protobuf/proto"

	gcli "github.com/urfave/cli"

	messages "github.com/fibercrypto/skywallet-protob/go"
)

func transactionSignCmd() gcli.Command {
	name := "transactionSign"
	return gcli.Command{
		Name:        name,
		Usage:       "Ask the device to sign a transaction using the provided information.",
		Description: "",
		Flags: []gcli.Flag{
			gcli.StringSliceFlag{
				Name:  "inputHash",
				Usage: "Hash of the Input of the transaction we expect the device to sign",
			},
			gcli.IntSliceFlag{
				Name:  "inputIndex",
				Usage: "Index of the input in the wallet",
			},
			gcli.StringSliceFlag{
				Name:  "outputAddress",
				Usage: "Addresses of the output for the transaction",
			},
			gcli.Int64SliceFlag{
				Name:  "coin",
				Usage: "Amount of coins",
			},
			gcli.Int64SliceFlag{
				Name:  "hour",
				Usage: "Number of hours",
			},
			gcli.IntSliceFlag{
				Name:  "addressIndex",
				Usage: "If the address is a return address tell its index in the wallet",
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
			inputs := c.StringSlice("inputHash")
			inputIndex := c.IntSlice("inputIndex")
			outputs := c.StringSlice("outputAddress")
			coins := c.Int64Slice("coin")
			hours := c.Int64Slice("hour")
			addressIndex := c.IntSlice("addressIndex")
			walletType := c.String("walletType")
			if len(inputs) != len(inputIndex) {
				fmt.Println("Every given input hash should have the an inputIndex")
				return
			}
			if len(outputs) != len(coins) || len(outputs) != len(hours) {
				fmt.Println("Every given output should have a coin and hour value")
				return
			}
			var transactionInputs []*messages.SkycoinTransactionInput
			var transactionOutputs []*messages.SkycoinTransactionOutput
			for i, input := range inputs {
				var transactionInput messages.SkycoinTransactionInput
				transactionInput.HashIn = proto.String(input)
				transactionInput.Index = proto.Uint32(uint32(inputIndex[i]))
				transactionInputs = append(transactionInputs, &transactionInput)
			}
			for i, output := range outputs {
				var transactionOutput messages.SkycoinTransactionOutput
				transactionOutput.Address = proto.String(output)
				transactionOutput.Coin = proto.Uint64(uint64(coins[i]))
				transactionOutput.Hour = proto.Uint64(uint64(hours[i]))
				if i < len(addressIndex) {
					transactionOutput.AddressIndex = proto.Uint32(uint32(addressIndex[i]))
				}
				transactionOutputs = append(transactionOutputs, &transactionOutput)
			}
			sq, err := createDevice(c.String("deviceType"))
			if err != nil {
				return
			}
			msg, err := sq.TransactionSign(transactionInputs, transactionOutputs, walletType)
			handleFinalResponse(msg, err, "unable to sign transaction", messages.MessageType_MessageType_ResponseTransactionSign)
		},
	}
}

package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	"github.com/sirupsen/logrus"
	"github.com/fibercrypto/skywallet-go/src/integration/proxy"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/gogo/protobuf/proto"
	"os"
	"runtime"
)

func parseBool(s string) (*bool, error) {
	var b bool
	switch s {
	case "true":
		b = true
	case "false":
		b = false
	case "":
		return nil, nil
	default:
		return nil, errors.New("Invalid boolean argument")
	}
	return &b, nil
}

func createDevice(devType string) (skywallet.Devicer, error) {
	device := skywallet.NewDevice(skywallet.DeviceTypeFromString(devType))
	if device == nil {
		return nil, errors.New("got null device")
	}
	if os.Getenv("AUTO_PRESS_BUTTONS") == "1" && device.Driver.DeviceType() == skywallet.DeviceTypeEmulator && runtime.GOOS == "linux" {
		err := device.SetAutoPressButton(true, skywallet.ButtonRight)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}
	return proxy.NewSequencer(device, false, func(kind skywallet.InputRequestKind, title, message string) (string, error) {
		if kind != skywallet.RequestInformUserOnlyOk && kind != skywallet.RequestInformUserOnlyCancel && kind != skywallet.RequestInformUserOkAndCancel {
			var line string
			if kind == skywallet.RequestKindWord {
				logging.NewMasterLogger().Printf("Word:")
				fmt.Scan(&line)
			} else {
				logging.NewMasterLogger().Printf(title)
				fmt.Scanln(&line)
			}
			return line, nil
		}
		return "", nil
	}), nil
}

func handleFinalResponse(msg wire.Message, err error, errMsg string, expectedMsgKind messages.MessageType) {
	if err != nil {
		logrus.WithError(err).Errorln(errMsg)
	}
	if msg.Kind != uint16(expectedMsgKind) {
		logrus.Errorln(errMsg + "invalid state")
	}
	switch messages.MessageType(msg.Kind) {
	case messages.MessageType_MessageType_ResponseSkycoinAddress:
		msgStr, err := skywallet.DecodeResponseSkycoinAddress(msg)
		if err != nil {
			logrus.WithError(err).Errorln("unable to decode response")
		}
		fmt.Println(msgStr)
	case messages.MessageType_MessageType_Success:
		msgStr, err := skywallet.DecodeSuccessMsg(msg)
		if err != nil {
			logrus.WithError(err).Errorln("unable to decode response")
		}
		fmt.Println(msgStr)
	case messages.MessageType_MessageType_ResponseSkycoinSignMessage:
		msgStr, err := skywallet.DecodeResponseSkycoinSignMessage(msg)
		if err != nil {
			logrus.WithError(err).Errorln("unable to decode response")
		}
		fmt.Println(msgStr)
	case messages.MessageType_MessageType_ResponseTransactionSign:
		msgStr, err := skywallet.DecodeResponseTransactionSign(msg)
		if err != nil {
			logrus.WithError(err).Errorln("unable to decode response")
		}
		fmt.Println(msgStr)
	case messages.MessageType_MessageType_Features:
		features := &messages.Features{}
		if err = proto.Unmarshal(msg.Data, features); err != nil {
			log.Error(err)
		}
		enc := json.NewEncoder(os.Stdout)
		if err = enc.Encode(features); err != nil {
			log.Errorln(err)
		}
		ff := skywallet.NewFirmwareFeatures(uint64(*features.FirmwareFeatures))
		if err := ff.Unmarshal(); err != nil {
			log.Errorln(err)
		}
		log.Printf("\n\nFirmware features:\n%s", ff)
	default:
		logrus.Errorln("invalid state")
	}
}
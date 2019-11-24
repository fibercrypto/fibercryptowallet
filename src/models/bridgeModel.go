package models

import (
	"github.com/therecipe/qt/core"
)

type QBridge struct {
	core.QObject
	_ func(string) string `signal:"getPassword"`
}

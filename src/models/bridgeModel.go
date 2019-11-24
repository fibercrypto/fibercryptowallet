package models

import (
	"sync"

	"github.com/therecipe/qt/core"
)

type QBridge struct {
	core.QObject
	_   func(message string) `signal:"getPassword"`
	_   string               `property:"returnGetPassword"`
	_   func()               `slot:"lock"`
	_   func()               `slot:"unlock"`
	sem *sync.Mutex
}

func (b *QBridge) init() {
	b.ConnectLock(b.lock)
	b.ConnectUnlock(b.unlock)
}

func (b *QBridge) lock() {
	b.sem.Lock()
}

func (b *QBridge) unlock() {
	b.sem.Unlock()
}

package models

import (
	"sync"

	"github.com/therecipe/qt/core"
)

type QBridge struct {
	core.QObject
	_      func()               `constructor:"init"`
	_      func()               `slot:"onCompleted"`
	_      func()               `slot:"lock"`
	_      func()               `slot:"unlock"`
	_      func(message string) `signal:"getPassword"`
	_      func(message string) `signal:"getSkyHardwareWalletPin"`
	_      func(string)         `slot:"setResult"`
	_      func() string        `slot:"getResult"`
	result string
	sem    sync.Mutex
	use    sync.Mutex
}

func (b *QBridge) init() {

	b.ConnectLock(b.lock)
	b.ConnectUnlock(b.unlock)
	b.ConnectSetResult(b.setResult)
	b.ConnectGetResult(b.getResult)
	b.ConnectOnCompleted(b.onCompleted)
}

func (b *QBridge) onCompleted() {
	createSkyHardwareWallet(b)
}

func (b *QBridge) BeginUse() {
	b.use.Lock()
}

func (b *QBridge) EndUse() {
	b.use.Unlock()
}

func (b *QBridge) lock() {
	b.sem.Lock()
}

func (b *QBridge) setResult(result string) {
	b.result = result
}

func (b *QBridge) getResult() string {
	return b.result
}

func (b *QBridge) unlock() {
	b.sem.Unlock()
}

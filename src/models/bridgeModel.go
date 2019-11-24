package models

import (
	"sync"

	"github.com/therecipe/qt/core"
)

type QBridge struct {
	core.QObject
	_      func()               `constructor:"init"`
	_      func()               `slot:"lock"`
	_      func()               `slot:"unlock"`
	_      func(message string) `signal:"getPassword"`
	_      func(string)         `slot:"setResult"`
	_      func() string        `slot:"getResult"`
	result string
	sem    sync.Mutex
}

func (b *QBridge) init() {

	b.ConnectLock(b.lock)
	b.ConnectUnlock(b.unlock)
	b.ConnectSetResult(b.setResult)
	b.ConnectGetResult(b.getResult)
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

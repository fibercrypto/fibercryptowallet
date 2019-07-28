package walletsManager

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	custom_wallets_445aa6m "github.com/fibercrypto/FiberCryptoWallet/src/models/wallets"
	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type WalletManager_ITF interface {
	std_core.QObject_ITF
	WalletManager_PTR() *WalletManager
}

func (ptr *WalletManager) WalletManager_PTR() *WalletManager {
	return ptr
}

func (ptr *WalletManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *WalletManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromWalletManager(ptr WalletManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WalletManager_PTR().Pointer()
	}
	return nil
}

func NewWalletManagerFromPointer(ptr unsafe.Pointer) (n *WalletManager) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(WalletManager)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *WalletManager:
			n = deduced

		case *std_core.QObject:
			n = &WalletManager{QObject: *deduced}

		default:
			n = new(WalletManager)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *WalletManager) Init() { this.init() }

//export callbackWalletManager64bdd5_Constructor
func callbackWalletManager64bdd5_Constructor(ptr unsafe.Pointer) {
	this := NewWalletManagerFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackWalletManager64bdd5_CreateEncryptedWallet
func callbackWalletManager64bdd5_CreateEncryptedWallet(ptr unsafe.Pointer, seed C.struct_Moc_PackedString, label C.struct_Moc_PackedString, password C.struct_Moc_PackedString, scanN C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createEncryptedWallet"); signal != nil {
		return custom_wallets_445aa6m.PointerFromQWallet((*(*func(string, string, string, int) *custom_wallets_445aa6m.QWallet)(signal))(cGoUnpackString(seed), cGoUnpackString(label), cGoUnpackString(password), int(int32(scanN))))
	}

	return custom_wallets_445aa6m.PointerFromQWallet(custom_wallets_445aa6m.NewQWallet(nil))
}

func (ptr *WalletManager) ConnectCreateEncryptedWallet(f func(seed string, label string, password string, scanN int) *custom_wallets_445aa6m.QWallet) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createEncryptedWallet"); signal != nil {
			f := func(seed string, label string, password string, scanN int) *custom_wallets_445aa6m.QWallet {
				(*(*func(string, string, string, int) *custom_wallets_445aa6m.QWallet)(signal))(seed, label, password, scanN)
				return f(seed, label, password, scanN)
			}
			qt.ConnectSignal(ptr.Pointer(), "createEncryptedWallet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createEncryptedWallet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletManager) DisconnectCreateEncryptedWallet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createEncryptedWallet")
	}
}

func (ptr *WalletManager) CreateEncryptedWallet(seed string, label string, password string, scanN int) *custom_wallets_445aa6m.QWallet {
	if ptr.Pointer() != nil {
		var seedC *C.char
		if seed != "" {
			seedC = C.CString(seed)
			defer C.free(unsafe.Pointer(seedC))
		}
		var labelC *C.char
		if label != "" {
			labelC = C.CString(label)
			defer C.free(unsafe.Pointer(labelC))
		}
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		tmpValue := custom_wallets_445aa6m.NewQWalletFromPointer(C.WalletManager64bdd5_CreateEncryptedWallet(ptr.Pointer(), C.struct_Moc_PackedString{data: seedC, len: C.longlong(len(seed))}, C.struct_Moc_PackedString{data: labelC, len: C.longlong(len(label))}, C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))}, C.int(int32(scanN))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackWalletManager64bdd5_CreateUnencryptedWallet
func callbackWalletManager64bdd5_CreateUnencryptedWallet(ptr unsafe.Pointer, seed C.struct_Moc_PackedString, label C.struct_Moc_PackedString, scanN C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createUnencryptedWallet"); signal != nil {
		return custom_wallets_445aa6m.PointerFromQWallet((*(*func(string, string, int) *custom_wallets_445aa6m.QWallet)(signal))(cGoUnpackString(seed), cGoUnpackString(label), int(int32(scanN))))
	}

	return custom_wallets_445aa6m.PointerFromQWallet(custom_wallets_445aa6m.NewQWallet(nil))
}

func (ptr *WalletManager) ConnectCreateUnencryptedWallet(f func(seed string, label string, scanN int) *custom_wallets_445aa6m.QWallet) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createUnencryptedWallet"); signal != nil {
			f := func(seed string, label string, scanN int) *custom_wallets_445aa6m.QWallet {
				(*(*func(string, string, int) *custom_wallets_445aa6m.QWallet)(signal))(seed, label, scanN)
				return f(seed, label, scanN)
			}
			qt.ConnectSignal(ptr.Pointer(), "createUnencryptedWallet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createUnencryptedWallet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletManager) DisconnectCreateUnencryptedWallet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createUnencryptedWallet")
	}
}

func (ptr *WalletManager) CreateUnencryptedWallet(seed string, label string, scanN int) *custom_wallets_445aa6m.QWallet {
	if ptr.Pointer() != nil {
		var seedC *C.char
		if seed != "" {
			seedC = C.CString(seed)
			defer C.free(unsafe.Pointer(seedC))
		}
		var labelC *C.char
		if label != "" {
			labelC = C.CString(label)
			defer C.free(unsafe.Pointer(labelC))
		}
		tmpValue := custom_wallets_445aa6m.NewQWalletFromPointer(C.WalletManager64bdd5_CreateUnencryptedWallet(ptr.Pointer(), C.struct_Moc_PackedString{data: seedC, len: C.longlong(len(seed))}, C.struct_Moc_PackedString{data: labelC, len: C.longlong(len(label))}, C.int(int32(scanN))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackWalletManager64bdd5_GetNewSeed
func callbackWalletManager64bdd5_GetNewSeed(ptr unsafe.Pointer, entropy C.int) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "getNewSeed"); signal != nil {
		tempVal := (*(*func(int) string)(signal))(int(int32(entropy)))
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *WalletManager) ConnectGetNewSeed(f func(entropy int) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getNewSeed"); signal != nil {
			f := func(entropy int) string {
				(*(*func(int) string)(signal))(entropy)
				return f(entropy)
			}
			qt.ConnectSignal(ptr.Pointer(), "getNewSeed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getNewSeed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletManager) DisconnectGetNewSeed() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getNewSeed")
	}
}

func (ptr *WalletManager) GetNewSeed(entropy int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.WalletManager64bdd5_GetNewSeed(ptr.Pointer(), C.int(int32(entropy))))
	}
	return ""
}

//export callbackWalletManager64bdd5_VerifySeed
func callbackWalletManager64bdd5_VerifySeed(ptr unsafe.Pointer, seed C.struct_Moc_PackedString) C.int {
	if signal := qt.GetSignal(ptr, "verifySeed"); signal != nil {
		return C.int(int32((*(*func(string) int)(signal))(cGoUnpackString(seed))))
	}

	return C.int(int32(0))
}

func (ptr *WalletManager) ConnectVerifySeed(f func(seed string) int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "verifySeed"); signal != nil {
			f := func(seed string) int {
				(*(*func(string) int)(signal))(seed)
				return f(seed)
			}
			qt.ConnectSignal(ptr.Pointer(), "verifySeed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "verifySeed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletManager) DisconnectVerifySeed() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "verifySeed")
	}
}

func (ptr *WalletManager) VerifySeed(seed string) int {
	if ptr.Pointer() != nil {
		var seedC *C.char
		if seed != "" {
			seedC = C.CString(seed)
			defer C.free(unsafe.Pointer(seedC))
		}
		return int(int32(C.WalletManager64bdd5_VerifySeed(ptr.Pointer(), C.struct_Moc_PackedString{data: seedC, len: C.longlong(len(seed))})))
	}
	return 0
}

//export callbackWalletManager64bdd5_NewWalletAddress
func callbackWalletManager64bdd5_NewWalletAddress(ptr unsafe.Pointer, id C.struct_Moc_PackedString, n C.int, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "newWalletAddress"); signal != nil {
		(*(*func(string, int, string))(signal))(cGoUnpackString(id), int(int32(n)), cGoUnpackString(password))
	}

}

func (ptr *WalletManager) ConnectNewWalletAddress(f func(id string, n int, password string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "newWalletAddress"); signal != nil {
			f := func(id string, n int, password string) {
				(*(*func(string, int, string))(signal))(id, n, password)
				f(id, n, password)
			}
			qt.ConnectSignal(ptr.Pointer(), "newWalletAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newWalletAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletManager) DisconnectNewWalletAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "newWalletAddress")
	}
}

func (ptr *WalletManager) NewWalletAddress(id string, n int, password string) {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.WalletManager64bdd5_NewWalletAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))}, C.int(int32(n)), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackWalletManager64bdd5_EncryptWallet
func callbackWalletManager64bdd5_EncryptWallet(ptr unsafe.Pointer, id C.struct_Moc_PackedString, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "encryptWallet"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(id), cGoUnpackString(password))
	}

}

func (ptr *WalletManager) ConnectEncryptWallet(f func(id string, password string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "encryptWallet"); signal != nil {
			f := func(id string, password string) {
				(*(*func(string, string))(signal))(id, password)
				f(id, password)
			}
			qt.ConnectSignal(ptr.Pointer(), "encryptWallet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "encryptWallet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletManager) DisconnectEncryptWallet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "encryptWallet")
	}
}

func (ptr *WalletManager) EncryptWallet(id string, password string) {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.WalletManager64bdd5_EncryptWallet(ptr.Pointer(), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))}, C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackWalletManager64bdd5_DecryptWallet
func callbackWalletManager64bdd5_DecryptWallet(ptr unsafe.Pointer, id C.struct_Moc_PackedString, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "decryptWallet"); signal != nil {
		(*(*func(string, string))(signal))(cGoUnpackString(id), cGoUnpackString(password))
	}

}

func (ptr *WalletManager) ConnectDecryptWallet(f func(id string, password string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "decryptWallet"); signal != nil {
			f := func(id string, password string) {
				(*(*func(string, string))(signal))(id, password)
				f(id, password)
			}
			qt.ConnectSignal(ptr.Pointer(), "decryptWallet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "decryptWallet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletManager) DisconnectDecryptWallet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "decryptWallet")
	}
}

func (ptr *WalletManager) DecryptWallet(id string, password string) {
	if ptr.Pointer() != nil {
		var idC *C.char
		if id != "" {
			idC = C.CString(id)
			defer C.free(unsafe.Pointer(idC))
		}
		var passwordC *C.char
		if password != "" {
			passwordC = C.CString(password)
			defer C.free(unsafe.Pointer(passwordC))
		}
		C.WalletManager64bdd5_DecryptWallet(ptr.Pointer(), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))}, C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackWalletManager64bdd5_GetWallets
func callbackWalletManager64bdd5_GetWallets(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "getWallets"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewWalletManagerFromPointer(NewWalletManagerFromPointer(nil).__getWallets_newList())
			for _, v := range (*(*func() []*custom_wallets_445aa6m.QWallet)(signal))() {
				tmpList.__getWallets_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewWalletManagerFromPointer(NewWalletManagerFromPointer(nil).__getWallets_newList())
		for _, v := range make([]*custom_wallets_445aa6m.QWallet, 0) {
			tmpList.__getWallets_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *WalletManager) ConnectGetWallets(f func() []*custom_wallets_445aa6m.QWallet) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "getWallets"); signal != nil {
			f := func() []*custom_wallets_445aa6m.QWallet {
				(*(*func() []*custom_wallets_445aa6m.QWallet)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "getWallets", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "getWallets", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletManager) DisconnectGetWallets() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "getWallets")
	}
}

func (ptr *WalletManager) GetWallets() []*custom_wallets_445aa6m.QWallet {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*custom_wallets_445aa6m.QWallet {
			out := make([]*custom_wallets_445aa6m.QWallet, int(l.len))
			tmpList := NewWalletManagerFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__getWallets_atList(i)
			}
			return out
		}(C.WalletManager64bdd5_GetWallets(ptr.Pointer()))
	}
	return make([]*custom_wallets_445aa6m.QWallet, 0)
}

func WalletManager_QRegisterMetaType() int {
	return int(int32(C.WalletManager64bdd5_WalletManager64bdd5_QRegisterMetaType()))
}

func (ptr *WalletManager) QRegisterMetaType() int {
	return int(int32(C.WalletManager64bdd5_WalletManager64bdd5_QRegisterMetaType()))
}

func WalletManager_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.WalletManager64bdd5_WalletManager64bdd5_QRegisterMetaType2(typeNameC)))
}

func (ptr *WalletManager) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.WalletManager64bdd5_WalletManager64bdd5_QRegisterMetaType2(typeNameC)))
}

func WalletManager_QmlRegisterType() int {
	return int(int32(C.WalletManager64bdd5_WalletManager64bdd5_QmlRegisterType()))
}

func (ptr *WalletManager) QmlRegisterType() int {
	return int(int32(C.WalletManager64bdd5_WalletManager64bdd5_QmlRegisterType()))
}

func WalletManager_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.WalletManager64bdd5_WalletManager64bdd5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *WalletManager) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.WalletManager64bdd5_WalletManager64bdd5_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *WalletManager) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletManager64bdd5___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletManager) __children_newList() unsafe.Pointer {
	return C.WalletManager64bdd5___children_newList(ptr.Pointer())
}

func (ptr *WalletManager) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.WalletManager64bdd5___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *WalletManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.WalletManager64bdd5___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *WalletManager) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletManager64bdd5___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletManager) __findChildren_newList() unsafe.Pointer {
	return C.WalletManager64bdd5___findChildren_newList(ptr.Pointer())
}

func (ptr *WalletManager) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletManager64bdd5___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletManager) __findChildren_newList3() unsafe.Pointer {
	return C.WalletManager64bdd5___findChildren_newList3(ptr.Pointer())
}

func (ptr *WalletManager) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletManager64bdd5___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletManager) __qFindChildren_newList2() unsafe.Pointer {
	return C.WalletManager64bdd5___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *WalletManager) __getWallets_atList(i int) *custom_wallets_445aa6m.QWallet {
	if ptr.Pointer() != nil {
		tmpValue := custom_wallets_445aa6m.NewQWalletFromPointer(C.WalletManager64bdd5___getWallets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __getWallets_setList(i custom_wallets_445aa6m.QWallet_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5___getWallets_setList(ptr.Pointer(), custom_wallets_445aa6m.PointerFromQWallet(i))
	}
}

func (ptr *WalletManager) __getWallets_newList() unsafe.Pointer {
	return C.WalletManager64bdd5___getWallets_newList(ptr.Pointer())
}

func NewWalletManager(parent std_core.QObject_ITF) *WalletManager {
	tmpValue := NewWalletManagerFromPointer(C.WalletManager64bdd5_NewWalletManager(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackWalletManager64bdd5_DestroyWalletManager
func callbackWalletManager64bdd5_DestroyWalletManager(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~WalletManager"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewWalletManagerFromPointer(ptr).DestroyWalletManagerDefault()
	}
}

func (ptr *WalletManager) ConnectDestroyWalletManager(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~WalletManager"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~WalletManager", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~WalletManager", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletManager) DisconnectDestroyWalletManager() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~WalletManager")
	}
}

func (ptr *WalletManager) DestroyWalletManager() {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5_DestroyWalletManager(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *WalletManager) DestroyWalletManagerDefault() {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5_DestroyWalletManagerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackWalletManager64bdd5_ChildEvent
func callbackWalletManager64bdd5_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewWalletManagerFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *WalletManager) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackWalletManager64bdd5_ConnectNotify
func callbackWalletManager64bdd5_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewWalletManagerFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *WalletManager) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackWalletManager64bdd5_CustomEvent
func callbackWalletManager64bdd5_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewWalletManagerFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *WalletManager) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackWalletManager64bdd5_DeleteLater
func callbackWalletManager64bdd5_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewWalletManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *WalletManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackWalletManager64bdd5_Destroyed
func callbackWalletManager64bdd5_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackWalletManager64bdd5_DisconnectNotify
func callbackWalletManager64bdd5_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewWalletManagerFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *WalletManager) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackWalletManager64bdd5_Event
func callbackWalletManager64bdd5_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletManagerFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *WalletManager) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletManager64bdd5_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackWalletManager64bdd5_EventFilter
func callbackWalletManager64bdd5_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletManagerFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *WalletManager) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletManager64bdd5_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackWalletManager64bdd5_ObjectNameChanged
func callbackWalletManager64bdd5_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackWalletManager64bdd5_TimerEvent
func callbackWalletManager64bdd5_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewWalletManagerFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *WalletManager) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager64bdd5_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

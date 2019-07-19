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

//export callbackWalletManager268d39_Constructor
func callbackWalletManager268d39_Constructor(ptr unsafe.Pointer) {
	this := NewWalletManagerFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackWalletManager268d39_CreateEncryptedWallet
func callbackWalletManager268d39_CreateEncryptedWallet(ptr unsafe.Pointer, seed C.struct_Moc_PackedString, label C.struct_Moc_PackedString, password C.struct_Moc_PackedString, scanN C.int) {
	if signal := qt.GetSignal(ptr, "createEncryptedWallet"); signal != nil {
		(*(*func(string, string, string, int))(signal))(cGoUnpackString(seed), cGoUnpackString(label), cGoUnpackString(password), int(int32(scanN)))
	}

}

func (ptr *WalletManager) ConnectCreateEncryptedWallet(f func(seed string, label string, password string, scanN int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createEncryptedWallet"); signal != nil {
			f := func(seed string, label string, password string, scanN int) {
				(*(*func(string, string, string, int))(signal))(seed, label, password, scanN)
				f(seed, label, password, scanN)
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

func (ptr *WalletManager) CreateEncryptedWallet(seed string, label string, password string, scanN int) {
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
		C.WalletManager268d39_CreateEncryptedWallet(ptr.Pointer(), C.struct_Moc_PackedString{data: seedC, len: C.longlong(len(seed))}, C.struct_Moc_PackedString{data: labelC, len: C.longlong(len(label))}, C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))}, C.int(int32(scanN)))
	}
}

//export callbackWalletManager268d39_CreateUnencryptedWallet
func callbackWalletManager268d39_CreateUnencryptedWallet(ptr unsafe.Pointer, seed C.struct_Moc_PackedString, label C.struct_Moc_PackedString, scanN C.int) {
	if signal := qt.GetSignal(ptr, "createUnencryptedWallet"); signal != nil {
		(*(*func(string, string, int))(signal))(cGoUnpackString(seed), cGoUnpackString(label), int(int32(scanN)))
	}

}

func (ptr *WalletManager) ConnectCreateUnencryptedWallet(f func(seed string, label string, scanN int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createUnencryptedWallet"); signal != nil {
			f := func(seed string, label string, scanN int) {
				(*(*func(string, string, int))(signal))(seed, label, scanN)
				f(seed, label, scanN)
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

func (ptr *WalletManager) CreateUnencryptedWallet(seed string, label string, scanN int) {
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
		C.WalletManager268d39_CreateUnencryptedWallet(ptr.Pointer(), C.struct_Moc_PackedString{data: seedC, len: C.longlong(len(seed))}, C.struct_Moc_PackedString{data: labelC, len: C.longlong(len(label))}, C.int(int32(scanN)))
	}
}

//export callbackWalletManager268d39_GetNewSeed
func callbackWalletManager268d39_GetNewSeed(ptr unsafe.Pointer, entropy C.int) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.WalletManager268d39_GetNewSeed(ptr.Pointer(), C.int(int32(entropy))))
	}
	return ""
}

//export callbackWalletManager268d39_VerifySeed
func callbackWalletManager268d39_VerifySeed(ptr unsafe.Pointer, seed C.struct_Moc_PackedString) C.int {
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
		return int(int32(C.WalletManager268d39_VerifySeed(ptr.Pointer(), C.struct_Moc_PackedString{data: seedC, len: C.longlong(len(seed))})))
	}
	return 0
}

//export callbackWalletManager268d39_NewWalletAddress
func callbackWalletManager268d39_NewWalletAddress(ptr unsafe.Pointer, id C.struct_Moc_PackedString, n C.int, password C.struct_Moc_PackedString) {
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
		C.WalletManager268d39_NewWalletAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))}, C.int(int32(n)), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackWalletManager268d39_EncryptWallet
func callbackWalletManager268d39_EncryptWallet(ptr unsafe.Pointer, id C.struct_Moc_PackedString, password C.struct_Moc_PackedString) {
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
		C.WalletManager268d39_EncryptWallet(ptr.Pointer(), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))}, C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackWalletManager268d39_DecryptWallet
func callbackWalletManager268d39_DecryptWallet(ptr unsafe.Pointer, id C.struct_Moc_PackedString, password C.struct_Moc_PackedString) {
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
		C.WalletManager268d39_DecryptWallet(ptr.Pointer(), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))}, C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

func WalletManager_QRegisterMetaType() int {
	return int(int32(C.WalletManager268d39_WalletManager268d39_QRegisterMetaType()))
}

func (ptr *WalletManager) QRegisterMetaType() int {
	return int(int32(C.WalletManager268d39_WalletManager268d39_QRegisterMetaType()))
}

func WalletManager_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.WalletManager268d39_WalletManager268d39_QRegisterMetaType2(typeNameC)))
}

func (ptr *WalletManager) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.WalletManager268d39_WalletManager268d39_QRegisterMetaType2(typeNameC)))
}

func WalletManager_QmlRegisterType() int {
	return int(int32(C.WalletManager268d39_WalletManager268d39_QmlRegisterType()))
}

func (ptr *WalletManager) QmlRegisterType() int {
	return int(int32(C.WalletManager268d39_WalletManager268d39_QmlRegisterType()))
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
	return int(int32(C.WalletManager268d39_WalletManager268d39_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.WalletManager268d39_WalletManager268d39_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *WalletManager) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletManager268d39___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager268d39___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletManager) __children_newList() unsafe.Pointer {
	return C.WalletManager268d39___children_newList(ptr.Pointer())
}

func (ptr *WalletManager) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.WalletManager268d39___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager268d39___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *WalletManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.WalletManager268d39___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *WalletManager) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletManager268d39___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager268d39___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletManager) __findChildren_newList() unsafe.Pointer {
	return C.WalletManager268d39___findChildren_newList(ptr.Pointer())
}

func (ptr *WalletManager) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletManager268d39___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager268d39___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletManager) __findChildren_newList3() unsafe.Pointer {
	return C.WalletManager268d39___findChildren_newList3(ptr.Pointer())
}

func (ptr *WalletManager) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletManager268d39___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager268d39___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletManager) __qFindChildren_newList2() unsafe.Pointer {
	return C.WalletManager268d39___qFindChildren_newList2(ptr.Pointer())
}

func NewWalletManager(parent std_core.QObject_ITF) *WalletManager {
	tmpValue := NewWalletManagerFromPointer(C.WalletManager268d39_NewWalletManager(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackWalletManager268d39_DestroyWalletManager
func callbackWalletManager268d39_DestroyWalletManager(ptr unsafe.Pointer) {
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
		C.WalletManager268d39_DestroyWalletManager(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *WalletManager) DestroyWalletManagerDefault() {
	if ptr.Pointer() != nil {
		C.WalletManager268d39_DestroyWalletManagerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackWalletManager268d39_ChildEvent
func callbackWalletManager268d39_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewWalletManagerFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *WalletManager) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager268d39_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackWalletManager268d39_ConnectNotify
func callbackWalletManager268d39_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewWalletManagerFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *WalletManager) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager268d39_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackWalletManager268d39_CustomEvent
func callbackWalletManager268d39_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewWalletManagerFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *WalletManager) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager268d39_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackWalletManager268d39_DeleteLater
func callbackWalletManager268d39_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewWalletManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *WalletManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.WalletManager268d39_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackWalletManager268d39_Destroyed
func callbackWalletManager268d39_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackWalletManager268d39_DisconnectNotify
func callbackWalletManager268d39_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewWalletManagerFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *WalletManager) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager268d39_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackWalletManager268d39_Event
func callbackWalletManager268d39_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletManagerFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *WalletManager) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletManager268d39_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackWalletManager268d39_EventFilter
func callbackWalletManager268d39_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletManagerFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *WalletManager) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletManager268d39_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackWalletManager268d39_ObjectNameChanged
func callbackWalletManager268d39_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackWalletManager268d39_TimerEvent
func callbackWalletManager268d39_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewWalletManagerFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *WalletManager) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager268d39_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type WalletModel_ITF interface {
	std_core.QAbstractListModel_ITF
	WalletModel_PTR() *WalletModel
}

func (ptr *WalletModel) WalletModel_PTR() *WalletModel {
	return ptr
}

func (ptr *WalletModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *WalletModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromWalletModel(ptr WalletModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.WalletModel_PTR().Pointer()
	}
	return nil
}

func NewWalletModelFromPointer(ptr unsafe.Pointer) (n *WalletModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(WalletModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *WalletModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &WalletModel{QAbstractListModel: *deduced}

		default:
			n = new(WalletModel)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *WalletModel) Init() { this.init() }

//export callbackWalletModel268d39_Constructor
func callbackWalletModel268d39_Constructor(ptr unsafe.Pointer) {
	this := NewWalletModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackWalletModel268d39_AddWallet
func callbackWalletModel268d39_AddWallet(ptr unsafe.Pointer, v0 unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addWallet"); signal != nil {
		(*(*func(*QWallet))(signal))(NewQWalletFromPointer(v0))
	}

}

func (ptr *WalletModel) ConnectAddWallet(f func(v0 *QWallet)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addWallet"); signal != nil {
			f := func(v0 *QWallet) {
				(*(*func(*QWallet))(signal))(v0)
				f(v0)
			}
			qt.ConnectSignal(ptr.Pointer(), "addWallet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addWallet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectAddWallet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addWallet")
	}
}

func (ptr *WalletModel) AddWallet(v0 QWallet_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_AddWallet(ptr.Pointer(), PointerFromQWallet(v0))
	}
}

//export callbackWalletModel268d39_EditWallet
func callbackWalletModel268d39_EditWallet(ptr unsafe.Pointer, row C.int, name C.struct_Moc_PackedString, encryptionEnabled C.char, sky C.int, coinHours C.int) {
	if signal := qt.GetSignal(ptr, "editWallet"); signal != nil {
		(*(*func(int, string, bool, int, int))(signal))(int(int32(row)), cGoUnpackString(name), int8(encryptionEnabled) != 0, int(int32(sky)), int(int32(coinHours)))
	}

}

func (ptr *WalletModel) ConnectEditWallet(f func(row int, name string, encryptionEnabled bool, sky int, coinHours int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "editWallet"); signal != nil {
			f := func(row int, name string, encryptionEnabled bool, sky int, coinHours int) {
				(*(*func(int, string, bool, int, int))(signal))(row, name, encryptionEnabled, sky, coinHours)
				f(row, name, encryptionEnabled, sky, coinHours)
			}
			qt.ConnectSignal(ptr.Pointer(), "editWallet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "editWallet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectEditWallet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "editWallet")
	}
}

func (ptr *WalletModel) EditWallet(row int, name string, encryptionEnabled bool, sky int, coinHours int) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.WalletModel268d39_EditWallet(ptr.Pointer(), C.int(int32(row)), C.struct_Moc_PackedString{data: nameC, len: C.longlong(len(name))}, C.char(int8(qt.GoBoolToInt(encryptionEnabled))), C.int(int32(sky)), C.int(int32(coinHours)))
	}
}

//export callbackWalletModel268d39_RemoveWallet
func callbackWalletModel268d39_RemoveWallet(ptr unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "removeWallet"); signal != nil {
		(*(*func(int))(signal))(int(int32(row)))
	}

}

func (ptr *WalletModel) ConnectRemoveWallet(f func(row int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeWallet"); signal != nil {
			f := func(row int) {
				(*(*func(int))(signal))(row)
				f(row)
			}
			qt.ConnectSignal(ptr.Pointer(), "removeWallet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeWallet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectRemoveWallet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeWallet")
	}
}

func (ptr *WalletModel) RemoveWallet(row int) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_RemoveWallet(ptr.Pointer(), C.int(int32(row)))
	}
}

//export callbackWalletModel268d39_LoadModel
func callbackWalletModel268d39_LoadModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "loadModel"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *WalletModel) ConnectLoadModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "loadModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "loadModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectLoadModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "loadModel")
	}
}

func (ptr *WalletModel) LoadModel() {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_LoadModel(ptr.Pointer())
	}
}

//export callbackWalletModel268d39_Roles
func callbackWalletModel268d39_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__roles_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__roles_newList())
		for k, v := range NewWalletModelFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *WalletModel) ConnectRoles(f func() map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "roles"); signal != nil {
			f := func() map[int]*std_core.QByteArray {
				(*(*func() map[int]*std_core.QByteArray)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "roles", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "roles", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *WalletModel) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.WalletModel268d39_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *WalletModel) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.WalletModel268d39_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackWalletModel268d39_SetRoles
func callbackWalletModel268d39_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	} else {
		NewWalletModelFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	}
}

func (ptr *WalletModel) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRoles"); signal != nil {
			f := func(roles map[int]*std_core.QByteArray) {
				(*(*func(map[int]*std_core.QByteArray))(signal))(roles)
				f(roles)
			}
			qt.ConnectSignal(ptr.Pointer(), "setRoles", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *WalletModel) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *WalletModel) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackWalletModel268d39_RolesChanged
func callbackWalletModel268d39_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i, v := range tmpList.__rolesChanged_roles_keyList() {
				out[v] = tmpList.__rolesChanged_roles_atList(v, i)
			}
			return out
		}(roles))
	}

}

func (ptr *WalletModel) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.WalletModel268d39_ConnectRolesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rolesChanged"); signal != nil {
			f := func(roles map[int]*std_core.QByteArray) {
				(*(*func(map[int]*std_core.QByteArray))(signal))(roles)
				f(roles)
			}
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *WalletModel) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackWalletModel268d39_Wallets
func callbackWalletModel268d39_Wallets(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "wallets"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__wallets_newList())
			for _, v := range (*(*func() []*QWallet)(signal))() {
				tmpList.__wallets_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__wallets_newList())
		for _, v := range NewWalletModelFromPointer(ptr).WalletsDefault() {
			tmpList.__wallets_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *WalletModel) ConnectWallets(f func() []*QWallet) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "wallets"); signal != nil {
			f := func() []*QWallet {
				(*(*func() []*QWallet)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "wallets", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "wallets", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectWallets() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "wallets")
	}
}

func (ptr *WalletModel) Wallets() []*QWallet {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*QWallet {
			out := make([]*QWallet, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__wallets_atList(i)
			}
			return out
		}(C.WalletModel268d39_Wallets(ptr.Pointer()))
	}
	return make([]*QWallet, 0)
}

func (ptr *WalletModel) WalletsDefault() []*QWallet {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*QWallet {
			out := make([]*QWallet, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__wallets_atList(i)
			}
			return out
		}(C.WalletModel268d39_WalletsDefault(ptr.Pointer()))
	}
	return make([]*QWallet, 0)
}

//export callbackWalletModel268d39_SetWallets
func callbackWalletModel268d39_SetWallets(ptr unsafe.Pointer, wallets C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setWallets"); signal != nil {
		(*(*func([]*QWallet))(signal))(func(l C.struct_Moc_PackedList) []*QWallet {
			out := make([]*QWallet, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setWallets_wallets_atList(i)
			}
			return out
		}(wallets))
	} else {
		NewWalletModelFromPointer(ptr).SetWalletsDefault(func(l C.struct_Moc_PackedList) []*QWallet {
			out := make([]*QWallet, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setWallets_wallets_atList(i)
			}
			return out
		}(wallets))
	}
}

func (ptr *WalletModel) ConnectSetWallets(f func(wallets []*QWallet)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setWallets"); signal != nil {
			f := func(wallets []*QWallet) {
				(*(*func([]*QWallet))(signal))(wallets)
				f(wallets)
			}
			qt.ConnectSignal(ptr.Pointer(), "setWallets", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setWallets", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectSetWallets() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setWallets")
	}
}

func (ptr *WalletModel) SetWallets(wallets []*QWallet) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_SetWallets(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__setWallets_wallets_newList())
			for _, v := range wallets {
				tmpList.__setWallets_wallets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *WalletModel) SetWalletsDefault(wallets []*QWallet) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_SetWalletsDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__setWallets_wallets_newList())
			for _, v := range wallets {
				tmpList.__setWallets_wallets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackWalletModel268d39_WalletsChanged
func callbackWalletModel268d39_WalletsChanged(ptr unsafe.Pointer, wallets C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "walletsChanged"); signal != nil {
		(*(*func([]*QWallet))(signal))(func(l C.struct_Moc_PackedList) []*QWallet {
			out := make([]*QWallet, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__walletsChanged_wallets_atList(i)
			}
			return out
		}(wallets))
	}

}

func (ptr *WalletModel) ConnectWalletsChanged(f func(wallets []*QWallet)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "walletsChanged") {
			C.WalletModel268d39_ConnectWalletsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "walletsChanged"); signal != nil {
			f := func(wallets []*QWallet) {
				(*(*func([]*QWallet))(signal))(wallets)
				f(wallets)
			}
			qt.ConnectSignal(ptr.Pointer(), "walletsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "walletsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectWalletsChanged() {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_DisconnectWalletsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "walletsChanged")
	}
}

func (ptr *WalletModel) WalletsChanged(wallets []*QWallet) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_WalletsChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__walletsChanged_wallets_newList())
			for _, v := range wallets {
				tmpList.__walletsChanged_wallets_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func WalletModel_QRegisterMetaType() int {
	return int(int32(C.WalletModel268d39_WalletModel268d39_QRegisterMetaType()))
}

func (ptr *WalletModel) QRegisterMetaType() int {
	return int(int32(C.WalletModel268d39_WalletModel268d39_QRegisterMetaType()))
}

func WalletModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.WalletModel268d39_WalletModel268d39_QRegisterMetaType2(typeNameC)))
}

func (ptr *WalletModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.WalletModel268d39_WalletModel268d39_QRegisterMetaType2(typeNameC)))
}

func WalletModel_QmlRegisterType() int {
	return int(int32(C.WalletModel268d39_WalletModel268d39_QmlRegisterType()))
}

func (ptr *WalletModel) QmlRegisterType() int {
	return int(int32(C.WalletModel268d39_WalletModel268d39_QmlRegisterType()))
}

func WalletModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.WalletModel268d39_WalletModel268d39_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *WalletModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.WalletModel268d39_WalletModel268d39_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *WalletModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WalletModel268d39_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *WalletModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *WalletModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.WalletModel268d39_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *WalletModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WalletModel268d39_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *WalletModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *WalletModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.WalletModel268d39_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *WalletModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WalletModel268d39_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *WalletModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *WalletModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.WalletModel268d39_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *WalletModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.WalletModel268d39___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *WalletModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.WalletModel268d39___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *WalletModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.WalletModel268d39___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *WalletModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.WalletModel268d39___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *WalletModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WalletModel268d39___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *WalletModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *WalletModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.WalletModel268d39___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *WalletModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.WalletModel268d39___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *WalletModel) __itemData_newList() unsafe.Pointer {
	return C.WalletModel268d39___itemData_newList(ptr.Pointer())
}

func (ptr *WalletModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.WalletModel268d39___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *WalletModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.WalletModel268d39___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *WalletModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.WalletModel268d39___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *WalletModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.WalletModel268d39___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *WalletModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.WalletModel268d39___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *WalletModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.WalletModel268d39___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *WalletModel) __match_newList() unsafe.Pointer {
	return C.WalletModel268d39___match_newList(ptr.Pointer())
}

func (ptr *WalletModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.WalletModel268d39___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *WalletModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.WalletModel268d39___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *WalletModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.WalletModel268d39___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *WalletModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.WalletModel268d39___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *WalletModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.WalletModel268d39___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *WalletModel) __roleNames_newList() unsafe.Pointer {
	return C.WalletModel268d39___roleNames_newList(ptr.Pointer())
}

func (ptr *WalletModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.WalletModel268d39___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *WalletModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.WalletModel268d39___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *WalletModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.WalletModel268d39___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *WalletModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.WalletModel268d39___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *WalletModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WalletModel268d39_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *WalletModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *WalletModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.WalletModel268d39_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *WalletModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WalletModel268d39_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *WalletModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *WalletModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.WalletModel268d39_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *WalletModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletModel268d39___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletModel) __children_newList() unsafe.Pointer {
	return C.WalletModel268d39___children_newList(ptr.Pointer())
}

func (ptr *WalletModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.WalletModel268d39___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *WalletModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.WalletModel268d39___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *WalletModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletModel268d39___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletModel) __findChildren_newList() unsafe.Pointer {
	return C.WalletModel268d39___findChildren_newList(ptr.Pointer())
}

func (ptr *WalletModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletModel268d39___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletModel) __findChildren_newList3() unsafe.Pointer {
	return C.WalletModel268d39___findChildren_newList3(ptr.Pointer())
}

func (ptr *WalletModel) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletModel268d39___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletModel) __qFindChildren_newList2() unsafe.Pointer {
	return C.WalletModel268d39___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *WalletModel) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.WalletModel268d39___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *WalletModel) __roles_newList() unsafe.Pointer {
	return C.WalletModel268d39___roles_newList(ptr.Pointer())
}

func (ptr *WalletModel) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roles_keyList_atList(i)
			}
			return out
		}(C.WalletModel268d39___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *WalletModel) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.WalletModel268d39___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *WalletModel) __setRoles_roles_newList() unsafe.Pointer {
	return C.WalletModel268d39___setRoles_roles_newList(ptr.Pointer())
}

func (ptr *WalletModel) __setRoles_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoles_roles_keyList_atList(i)
			}
			return out
		}(C.WalletModel268d39___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *WalletModel) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.WalletModel268d39___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *WalletModel) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.WalletModel268d39___rolesChanged_roles_newList(ptr.Pointer())
}

func (ptr *WalletModel) __rolesChanged_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____rolesChanged_roles_keyList_atList(i)
			}
			return out
		}(C.WalletModel268d39___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *WalletModel) __wallets_atList(i int) *QWallet {
	if ptr.Pointer() != nil {
		tmpValue := NewQWalletFromPointer(C.WalletModel268d39___wallets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __wallets_setList(i QWallet_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___wallets_setList(ptr.Pointer(), PointerFromQWallet(i))
	}
}

func (ptr *WalletModel) __wallets_newList() unsafe.Pointer {
	return C.WalletModel268d39___wallets_newList(ptr.Pointer())
}

func (ptr *WalletModel) __setWallets_wallets_atList(i int) *QWallet {
	if ptr.Pointer() != nil {
		tmpValue := NewQWalletFromPointer(C.WalletModel268d39___setWallets_wallets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __setWallets_wallets_setList(i QWallet_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___setWallets_wallets_setList(ptr.Pointer(), PointerFromQWallet(i))
	}
}

func (ptr *WalletModel) __setWallets_wallets_newList() unsafe.Pointer {
	return C.WalletModel268d39___setWallets_wallets_newList(ptr.Pointer())
}

func (ptr *WalletModel) __walletsChanged_wallets_atList(i int) *QWallet {
	if ptr.Pointer() != nil {
		tmpValue := NewQWalletFromPointer(C.WalletModel268d39___walletsChanged_wallets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletModel) __walletsChanged_wallets_setList(i QWallet_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39___walletsChanged_wallets_setList(ptr.Pointer(), PointerFromQWallet(i))
	}
}

func (ptr *WalletModel) __walletsChanged_wallets_newList() unsafe.Pointer {
	return C.WalletModel268d39___walletsChanged_wallets_newList(ptr.Pointer())
}

func (ptr *WalletModel) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WalletModel268d39_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *WalletModel) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *WalletModel) ____roles_keyList_newList() unsafe.Pointer {
	return C.WalletModel268d39_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *WalletModel) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WalletModel268d39_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *WalletModel) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *WalletModel) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.WalletModel268d39_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *WalletModel) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WalletModel268d39_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *WalletModel) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *WalletModel) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.WalletModel268d39_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewWalletModel(parent std_core.QObject_ITF) *WalletModel {
	tmpValue := NewWalletModelFromPointer(C.WalletModel268d39_NewWalletModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackWalletModel268d39_DestroyWalletModel
func callbackWalletModel268d39_DestroyWalletModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~WalletModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewWalletModelFromPointer(ptr).DestroyWalletModelDefault()
	}
}

func (ptr *WalletModel) ConnectDestroyWalletModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~WalletModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~WalletModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~WalletModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectDestroyWalletModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~WalletModel")
	}
}

func (ptr *WalletModel) DestroyWalletModel() {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_DestroyWalletModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *WalletModel) DestroyWalletModelDefault() {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_DestroyWalletModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackWalletModel268d39_DropMimeData
func callbackWalletModel268d39_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *WalletModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackWalletModel268d39_Flags
func callbackWalletModel268d39_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewWalletModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *WalletModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.WalletModel268d39_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackWalletModel268d39_Index
func callbackWalletModel268d39_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewWalletModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *WalletModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.WalletModel268d39_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackWalletModel268d39_Sibling
func callbackWalletModel268d39_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewWalletModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *WalletModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.WalletModel268d39_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackWalletModel268d39_Buddy
func callbackWalletModel268d39_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewWalletModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *WalletModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.WalletModel268d39_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackWalletModel268d39_CanDropMimeData
func callbackWalletModel268d39_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *WalletModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackWalletModel268d39_CanFetchMore
func callbackWalletModel268d39_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *WalletModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackWalletModel268d39_ColumnCount
func callbackWalletModel268d39_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewWalletModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *WalletModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WalletModel268d39_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackWalletModel268d39_ColumnsAboutToBeInserted
func callbackWalletModel268d39_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackWalletModel268d39_ColumnsAboutToBeMoved
func callbackWalletModel268d39_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackWalletModel268d39_ColumnsAboutToBeRemoved
func callbackWalletModel268d39_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackWalletModel268d39_ColumnsInserted
func callbackWalletModel268d39_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackWalletModel268d39_ColumnsMoved
func callbackWalletModel268d39_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackWalletModel268d39_ColumnsRemoved
func callbackWalletModel268d39_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackWalletModel268d39_Data
func callbackWalletModel268d39_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewWalletModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *WalletModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.WalletModel268d39_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackWalletModel268d39_DataChanged
func callbackWalletModel268d39_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(signal))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackWalletModel268d39_FetchMore
func callbackWalletModel268d39_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewWalletModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *WalletModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackWalletModel268d39_HasChildren
func callbackWalletModel268d39_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *WalletModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackWalletModel268d39_HeaderData
func callbackWalletModel268d39_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewWalletModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *WalletModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.WalletModel268d39_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackWalletModel268d39_HeaderDataChanged
func callbackWalletModel268d39_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackWalletModel268d39_InsertColumns
func callbackWalletModel268d39_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *WalletModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackWalletModel268d39_InsertRows
func callbackWalletModel268d39_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *WalletModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackWalletModel268d39_ItemData
func callbackWalletModel268d39_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*std_core.QModelIndex) map[int]*std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__itemData_newList())
		for k, v := range NewWalletModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *WalletModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.WalletModel268d39_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackWalletModel268d39_LayoutAboutToBeChanged
func callbackWalletModel268d39_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackWalletModel268d39_LayoutChanged
func callbackWalletModel268d39_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackWalletModel268d39_Match
func callbackWalletModel268d39_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__match_newList())
		for _, v := range NewWalletModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *WalletModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.WalletModel268d39_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackWalletModel268d39_MimeData
func callbackWalletModel268d39_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func([]*std_core.QModelIndex) *std_core.QMimeData)(signal))(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewWalletModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewWalletModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *WalletModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.WalletModel268d39_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackWalletModel268d39_MimeTypes
func callbackWalletModel268d39_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewWalletModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *WalletModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.WalletModel268d39_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackWalletModel268d39_ModelAboutToBeReset
func callbackWalletModel268d39_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackWalletModel268d39_ModelReset
func callbackWalletModel268d39_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackWalletModel268d39_MoveColumns
func callbackWalletModel268d39_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *WalletModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackWalletModel268d39_MoveRows
func callbackWalletModel268d39_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *WalletModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackWalletModel268d39_Parent
func callbackWalletModel268d39_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewWalletModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *WalletModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.WalletModel268d39_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackWalletModel268d39_RemoveColumns
func callbackWalletModel268d39_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *WalletModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackWalletModel268d39_RemoveRows
func callbackWalletModel268d39_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *WalletModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackWalletModel268d39_ResetInternalData
func callbackWalletModel268d39_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewWalletModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *WalletModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackWalletModel268d39_Revert
func callbackWalletModel268d39_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewWalletModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *WalletModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_RevertDefault(ptr.Pointer())
	}
}

//export callbackWalletModel268d39_RoleNames
func callbackWalletModel268d39_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewWalletModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *WalletModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.WalletModel268d39_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackWalletModel268d39_RowCount
func callbackWalletModel268d39_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewWalletModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *WalletModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.WalletModel268d39_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackWalletModel268d39_RowsAboutToBeInserted
func callbackWalletModel268d39_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackWalletModel268d39_RowsAboutToBeMoved
func callbackWalletModel268d39_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackWalletModel268d39_RowsAboutToBeRemoved
func callbackWalletModel268d39_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackWalletModel268d39_RowsInserted
func callbackWalletModel268d39_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackWalletModel268d39_RowsMoved
func callbackWalletModel268d39_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackWalletModel268d39_RowsRemoved
func callbackWalletModel268d39_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackWalletModel268d39_SetData
func callbackWalletModel268d39_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *WalletModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackWalletModel268d39_SetHeaderData
func callbackWalletModel268d39_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *WalletModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackWalletModel268d39_SetItemData
func callbackWalletModel268d39_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(signal))(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewWalletModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewWalletModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *WalletModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewWalletModelFromPointer(NewWalletModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackWalletModel268d39_Sort
func callbackWalletModel268d39_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewWalletModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *WalletModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackWalletModel268d39_Span
func callbackWalletModel268d39_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewWalletModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *WalletModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.WalletModel268d39_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackWalletModel268d39_Submit
func callbackWalletModel268d39_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *WalletModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackWalletModel268d39_SupportedDragActions
func callbackWalletModel268d39_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewWalletModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *WalletModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.WalletModel268d39_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackWalletModel268d39_SupportedDropActions
func callbackWalletModel268d39_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewWalletModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *WalletModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.WalletModel268d39_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackWalletModel268d39_ChildEvent
func callbackWalletModel268d39_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewWalletModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *WalletModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackWalletModel268d39_ConnectNotify
func callbackWalletModel268d39_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewWalletModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *WalletModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackWalletModel268d39_CustomEvent
func callbackWalletModel268d39_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewWalletModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *WalletModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackWalletModel268d39_DeleteLater
func callbackWalletModel268d39_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewWalletModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *WalletModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackWalletModel268d39_Destroyed
func callbackWalletModel268d39_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackWalletModel268d39_DisconnectNotify
func callbackWalletModel268d39_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewWalletModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *WalletModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackWalletModel268d39_Event
func callbackWalletModel268d39_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *WalletModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackWalletModel268d39_EventFilter
func callbackWalletModel268d39_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *WalletModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletModel268d39_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackWalletModel268d39_ObjectNameChanged
func callbackWalletModel268d39_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackWalletModel268d39_TimerEvent
func callbackWalletModel268d39_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewWalletModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *WalletModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type AddressesModel_ITF interface {
	std_core.QAbstractListModel_ITF
	AddressesModel_PTR() *AddressesModel
}

func (ptr *AddressesModel) AddressesModel_PTR() *AddressesModel {
	return ptr
}

func (ptr *AddressesModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *AddressesModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromAddressesModel(ptr AddressesModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AddressesModel_PTR().Pointer()
	}
	return nil
}

func NewAddressesModelFromPointer(ptr unsafe.Pointer) (n *AddressesModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(AddressesModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *AddressesModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &AddressesModel{QAbstractListModel: *deduced}

		default:
			n = new(AddressesModel)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *AddressesModel) Init() { this.init() }

//export callbackAddressesModel268d39_Constructor
func callbackAddressesModel268d39_Constructor(ptr unsafe.Pointer) {
	this := NewAddressesModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackAddressesModel268d39_AddAddress
func callbackAddressesModel268d39_AddAddress(ptr unsafe.Pointer, v0 unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addAddress"); signal != nil {
		(*(*func(*QAddress))(signal))(NewQAddressFromPointer(v0))
	}

}

func (ptr *AddressesModel) ConnectAddAddress(f func(v0 *QAddress)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addAddress"); signal != nil {
			f := func(v0 *QAddress) {
				(*(*func(*QAddress))(signal))(v0)
				f(v0)
			}
			qt.ConnectSignal(ptr.Pointer(), "addAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectAddAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addAddress")
	}
}

func (ptr *AddressesModel) AddAddress(v0 QAddress_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_AddAddress(ptr.Pointer(), PointerFromQAddress(v0))
	}
}

//export callbackAddressesModel268d39_RemoveAddress
func callbackAddressesModel268d39_RemoveAddress(ptr unsafe.Pointer, v0 C.int) {
	if signal := qt.GetSignal(ptr, "removeAddress"); signal != nil {
		(*(*func(int))(signal))(int(int32(v0)))
	}

}

func (ptr *AddressesModel) ConnectRemoveAddress(f func(v0 int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "removeAddress"); signal != nil {
			f := func(v0 int) {
				(*(*func(int))(signal))(v0)
				f(v0)
			}
			qt.ConnectSignal(ptr.Pointer(), "removeAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectRemoveAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "removeAddress")
	}
}

func (ptr *AddressesModel) RemoveAddress(v0 int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_RemoveAddress(ptr.Pointer(), C.int(int32(v0)))
	}
}

//export callbackAddressesModel268d39_EditAddress
func callbackAddressesModel268d39_EditAddress(ptr unsafe.Pointer, v0 C.int, v1 C.struct_Moc_PackedString, v2 C.int, v3 C.int) {
	if signal := qt.GetSignal(ptr, "editAddress"); signal != nil {
		(*(*func(int, string, int, int))(signal))(int(int32(v0)), cGoUnpackString(v1), int(int32(v2)), int(int32(v3)))
	}

}

func (ptr *AddressesModel) ConnectEditAddress(f func(v0 int, v1 string, v2 int, v3 int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "editAddress"); signal != nil {
			f := func(v0 int, v1 string, v2 int, v3 int) {
				(*(*func(int, string, int, int))(signal))(v0, v1, v2, v3)
				f(v0, v1, v2, v3)
			}
			qt.ConnectSignal(ptr.Pointer(), "editAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "editAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectEditAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "editAddress")
	}
}

func (ptr *AddressesModel) EditAddress(v0 int, v1 string, v2 int, v3 int) {
	if ptr.Pointer() != nil {
		var v1C *C.char
		if v1 != "" {
			v1C = C.CString(v1)
			defer C.free(unsafe.Pointer(v1C))
		}
		C.AddressesModel268d39_EditAddress(ptr.Pointer(), C.int(int32(v0)), C.struct_Moc_PackedString{data: v1C, len: C.longlong(len(v1))}, C.int(int32(v2)), C.int(int32(v3)))
	}
}

//export callbackAddressesModel268d39_LoadModel
func callbackAddressesModel268d39_LoadModel(ptr unsafe.Pointer, v0 C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "loadModel"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(v0))
	}

}

func (ptr *AddressesModel) ConnectLoadModel(f func(v0 string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "loadModel"); signal != nil {
			f := func(v0 string) {
				(*(*func(string))(signal))(v0)
				f(v0)
			}
			qt.ConnectSignal(ptr.Pointer(), "loadModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectLoadModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "loadModel")
	}
}

func (ptr *AddressesModel) LoadModel(v0 string) {
	if ptr.Pointer() != nil {
		var v0C *C.char
		if v0 != "" {
			v0C = C.CString(v0)
			defer C.free(unsafe.Pointer(v0C))
		}
		C.AddressesModel268d39_LoadModel(ptr.Pointer(), C.struct_Moc_PackedString{data: v0C, len: C.longlong(len(v0))})
	}
}

//export callbackAddressesModel268d39_Roles
func callbackAddressesModel268d39_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__roles_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__roles_newList())
		for k, v := range NewAddressesModelFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressesModel) ConnectRoles(f func() map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "roles"); signal != nil {
			f := func() map[int]*std_core.QByteArray {
				(*(*func() map[int]*std_core.QByteArray)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "roles", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "roles", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *AddressesModel) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.AddressesModel268d39_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *AddressesModel) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.AddressesModel268d39_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackAddressesModel268d39_SetRoles
func callbackAddressesModel268d39_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	} else {
		NewAddressesModelFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	}
}

func (ptr *AddressesModel) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRoles"); signal != nil {
			f := func(roles map[int]*std_core.QByteArray) {
				(*(*func(map[int]*std_core.QByteArray))(signal))(roles)
				f(roles)
			}
			qt.ConnectSignal(ptr.Pointer(), "setRoles", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRoles", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *AddressesModel) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *AddressesModel) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackAddressesModel268d39_RolesChanged
func callbackAddressesModel268d39_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i, v := range tmpList.__rolesChanged_roles_keyList() {
				out[v] = tmpList.__rolesChanged_roles_atList(v, i)
			}
			return out
		}(roles))
	}

}

func (ptr *AddressesModel) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.AddressesModel268d39_ConnectRolesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rolesChanged"); signal != nil {
			f := func(roles map[int]*std_core.QByteArray) {
				(*(*func(map[int]*std_core.QByteArray))(signal))(roles)
				f(roles)
			}
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rolesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *AddressesModel) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackAddressesModel268d39_Addresses
func callbackAddressesModel268d39_Addresses(ptr unsafe.Pointer) C.uintptr_t {
	if signal := qt.GetSignal(ptr, "addresses"); signal != nil {
		oP := (*(*func() []*QAddress)(signal))()
		qt.RegisterTemp(unsafe.Pointer(&oP), unsafe.Pointer(&oP))
		return C.uintptr_t(uintptr(unsafe.Pointer(&oP)))
	}
	oP := NewAddressesModelFromPointer(ptr).AddressesDefault()
	qt.RegisterTemp(unsafe.Pointer(&oP), unsafe.Pointer(&oP))
	return C.uintptr_t(uintptr(unsafe.Pointer(&oP)))
}

func (ptr *AddressesModel) ConnectAddresses(f func() []*QAddress) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addresses"); signal != nil {
			f := func() []*QAddress {
				(*(*func() []*QAddress)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "addresses", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addresses", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectAddresses() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addresses")
	}
}

func (ptr *AddressesModel) Addresses() []*QAddress {
	if ptr.Pointer() != nil {
		oP := unsafe.Pointer(uintptr(C.AddressesModel268d39_Addresses(ptr.Pointer())))
		if oI, ok := qt.ReceiveTemp(oP); ok {
			oD := (*(*[]*QAddress)(oI))
			return oD
		}

	}
	var out []*QAddress
	return out
}

func (ptr *AddressesModel) AddressesDefault() []*QAddress {
	if ptr.Pointer() != nil {
		oP := unsafe.Pointer(uintptr(C.AddressesModel268d39_AddressesDefault(ptr.Pointer())))
		if oI, ok := qt.ReceiveTemp(oP); ok {
			oD := (*(*[]*QAddress)(oI))
			return oD
		}

	}
	var out []*QAddress
	return out
}

//export callbackAddressesModel268d39_SetAddresses
func callbackAddressesModel268d39_SetAddresses(ptr unsafe.Pointer, addresses C.uintptr_t) {
	var addressesD []*QAddress
	if addressesI, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(addresses))); ok {
		qt.UnregisterTemp(unsafe.Pointer(uintptr(addresses)))
		addressesD = (*(*[]*QAddress)(addressesI))
	}
	if signal := qt.GetSignal(ptr, "setAddresses"); signal != nil {
		(*(*func([]*QAddress))(signal))(addressesD)
	} else {
		NewAddressesModelFromPointer(ptr).SetAddressesDefault(addressesD)
	}
}

func (ptr *AddressesModel) ConnectSetAddresses(f func(addresses []*QAddress)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddresses"); signal != nil {
			f := func(addresses []*QAddress) {
				(*(*func([]*QAddress))(signal))(addresses)
				f(addresses)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAddresses", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAddresses", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectSetAddresses() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddresses")
	}
}

func (ptr *AddressesModel) SetAddresses(addresses []*QAddress) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(&addresses), unsafe.Pointer(&addresses))
		C.AddressesModel268d39_SetAddresses(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(&addresses))))
	}
}

func (ptr *AddressesModel) SetAddressesDefault(addresses []*QAddress) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(&addresses), unsafe.Pointer(&addresses))
		C.AddressesModel268d39_SetAddressesDefault(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(&addresses))))
	}
}

//export callbackAddressesModel268d39_AddressesChanged
func callbackAddressesModel268d39_AddressesChanged(ptr unsafe.Pointer, addresses C.uintptr_t) {
	var addressesD []*QAddress
	if addressesI, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(addresses))); ok {
		addressesD = (*(*[]*QAddress)(addressesI))
	}
	if signal := qt.GetSignal(ptr, "addressesChanged"); signal != nil {
		(*(*func([]*QAddress))(signal))(addressesD)
	}

}

func (ptr *AddressesModel) ConnectAddressesChanged(f func(addresses []*QAddress)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressesChanged") {
			C.AddressesModel268d39_ConnectAddressesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressesChanged"); signal != nil {
			f := func(addresses []*QAddress) {
				(*(*func([]*QAddress))(signal))(addresses)
				f(addresses)
			}
			qt.ConnectSignal(ptr.Pointer(), "addressesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectAddressesChanged() {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_DisconnectAddressesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressesChanged")
	}
}

func (ptr *AddressesModel) AddressesChanged(addresses []*QAddress) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(&addresses), unsafe.Pointer(&addresses))
		C.AddressesModel268d39_AddressesChanged(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(&addresses))))
	}
}

//export callbackAddressesModel268d39_Loaded
func callbackAddressesModel268d39_Loaded(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "loaded"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewAddressesModelFromPointer(ptr).LoadedDefault()))
}

func (ptr *AddressesModel) ConnectLoaded(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "loaded"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "loaded", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loaded", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectLoaded() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "loaded")
	}
}

func (ptr *AddressesModel) Loaded() int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39_Loaded(ptr.Pointer())))
	}
	return 0
}

func (ptr *AddressesModel) LoadedDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39_LoadedDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackAddressesModel268d39_SetLoaded
func callbackAddressesModel268d39_SetLoaded(ptr unsafe.Pointer, loaded C.int) {
	if signal := qt.GetSignal(ptr, "setLoaded"); signal != nil {
		(*(*func(int))(signal))(int(int32(loaded)))
	} else {
		NewAddressesModelFromPointer(ptr).SetLoadedDefault(int(int32(loaded)))
	}
}

func (ptr *AddressesModel) ConnectSetLoaded(f func(loaded int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLoaded"); signal != nil {
			f := func(loaded int) {
				(*(*func(int))(signal))(loaded)
				f(loaded)
			}
			qt.ConnectSignal(ptr.Pointer(), "setLoaded", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLoaded", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectSetLoaded() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLoaded")
	}
}

func (ptr *AddressesModel) SetLoaded(loaded int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_SetLoaded(ptr.Pointer(), C.int(int32(loaded)))
	}
}

func (ptr *AddressesModel) SetLoadedDefault(loaded int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_SetLoadedDefault(ptr.Pointer(), C.int(int32(loaded)))
	}
}

//export callbackAddressesModel268d39_LoadedChanged
func callbackAddressesModel268d39_LoadedChanged(ptr unsafe.Pointer, loaded C.int) {
	if signal := qt.GetSignal(ptr, "loadedChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(loaded)))
	}

}

func (ptr *AddressesModel) ConnectLoadedChanged(f func(loaded int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "loadedChanged") {
			C.AddressesModel268d39_ConnectLoadedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "loadedChanged"); signal != nil {
			f := func(loaded int) {
				(*(*func(int))(signal))(loaded)
				f(loaded)
			}
			qt.ConnectSignal(ptr.Pointer(), "loadedChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadedChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectLoadedChanged() {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_DisconnectLoadedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "loadedChanged")
	}
}

func (ptr *AddressesModel) LoadedChanged(loaded int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_LoadedChanged(ptr.Pointer(), C.int(int32(loaded)))
	}
}

func AddressesModel_QRegisterMetaType() int {
	return int(int32(C.AddressesModel268d39_AddressesModel268d39_QRegisterMetaType()))
}

func (ptr *AddressesModel) QRegisterMetaType() int {
	return int(int32(C.AddressesModel268d39_AddressesModel268d39_QRegisterMetaType()))
}

func AddressesModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressesModel268d39_AddressesModel268d39_QRegisterMetaType2(typeNameC)))
}

func (ptr *AddressesModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressesModel268d39_AddressesModel268d39_QRegisterMetaType2(typeNameC)))
}

func AddressesModel_QmlRegisterType() int {
	return int(int32(C.AddressesModel268d39_AddressesModel268d39_QmlRegisterType()))
}

func (ptr *AddressesModel) QmlRegisterType() int {
	return int(int32(C.AddressesModel268d39_AddressesModel268d39_QmlRegisterType()))
}

func AddressesModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.AddressesModel268d39_AddressesModel268d39_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *AddressesModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.AddressesModel268d39_AddressesModel268d39_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *AddressesModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressesModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressesModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.AddressesModel268d39_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *AddressesModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressesModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressesModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.AddressesModel268d39_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *AddressesModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressesModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressesModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.AddressesModel268d39_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressesModel268d39___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressesModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.AddressesModel268d39___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressesModel268d39___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressesModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.AddressesModel268d39___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressesModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressesModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.AddressesModel268d39___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressesModel268d39___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *AddressesModel) __itemData_newList() unsafe.Pointer {
	return C.AddressesModel268d39___itemData_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.AddressesModel268d39___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressesModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.AddressesModel268d39___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *AddressesModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.AddressesModel268d39___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.AddressesModel268d39___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *AddressesModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.AddressesModel268d39___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressesModel268d39___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressesModel) __match_newList() unsafe.Pointer {
	return C.AddressesModel268d39___match_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressesModel268d39___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressesModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.AddressesModel268d39___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressesModel268d39___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressesModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.AddressesModel268d39___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressesModel268d39___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressesModel) __roleNames_newList() unsafe.Pointer {
	return C.AddressesModel268d39___roleNames_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.AddressesModel268d39___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressesModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressesModel268d39___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *AddressesModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.AddressesModel268d39___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.AddressesModel268d39___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressesModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressesModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressesModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.AddressesModel268d39_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *AddressesModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressesModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressesModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.AddressesModel268d39_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressesModel268d39___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressesModel) __children_newList() unsafe.Pointer {
	return C.AddressesModel268d39___children_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressesModel268d39___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressesModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.AddressesModel268d39___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressesModel268d39___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressesModel) __findChildren_newList() unsafe.Pointer {
	return C.AddressesModel268d39___findChildren_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressesModel268d39___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressesModel) __findChildren_newList3() unsafe.Pointer {
	return C.AddressesModel268d39___findChildren_newList3(ptr.Pointer())
}

func (ptr *AddressesModel) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressesModel268d39___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressesModel) __qFindChildren_newList2() unsafe.Pointer {
	return C.AddressesModel268d39___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *AddressesModel) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressesModel268d39___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressesModel) __roles_newList() unsafe.Pointer {
	return C.AddressesModel268d39___roles_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roles_keyList_atList(i)
			}
			return out
		}(C.AddressesModel268d39___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressesModel) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressesModel268d39___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressesModel) __setRoles_roles_newList() unsafe.Pointer {
	return C.AddressesModel268d39___setRoles_roles_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __setRoles_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoles_roles_keyList_atList(i)
			}
			return out
		}(C.AddressesModel268d39___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressesModel) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressesModel268d39___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressesModel) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressesModel) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.AddressesModel268d39___rolesChanged_roles_newList(ptr.Pointer())
}

func (ptr *AddressesModel) __rolesChanged_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____rolesChanged_roles_keyList_atList(i)
			}
			return out
		}(C.AddressesModel268d39___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressesModel) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressesModel) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressesModel) ____roles_keyList_newList() unsafe.Pointer {
	return C.AddressesModel268d39_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *AddressesModel) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressesModel) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressesModel) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.AddressesModel268d39_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *AddressesModel) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressesModel) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressesModel) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.AddressesModel268d39_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewAddressesModel(parent std_core.QObject_ITF) *AddressesModel {
	tmpValue := NewAddressesModelFromPointer(C.AddressesModel268d39_NewAddressesModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackAddressesModel268d39_DestroyAddressesModel
func callbackAddressesModel268d39_DestroyAddressesModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~AddressesModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressesModelFromPointer(ptr).DestroyAddressesModelDefault()
	}
}

func (ptr *AddressesModel) ConnectDestroyAddressesModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~AddressesModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~AddressesModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~AddressesModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressesModel) DisconnectDestroyAddressesModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~AddressesModel")
	}
}

func (ptr *AddressesModel) DestroyAddressesModel() {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_DestroyAddressesModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *AddressesModel) DestroyAddressesModelDefault() {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_DestroyAddressesModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressesModel268d39_DropMimeData
func callbackAddressesModel268d39_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressesModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_Flags
func callbackAddressesModel268d39_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewAddressesModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressesModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.AddressesModel268d39_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackAddressesModel268d39_Index
func callbackAddressesModel268d39_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewAddressesModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *AddressesModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressesModel268d39_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressesModel268d39_Sibling
func callbackAddressesModel268d39_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewAddressesModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *AddressesModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressesModel268d39_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressesModel268d39_Buddy
func callbackAddressesModel268d39_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewAddressesModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressesModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressesModel268d39_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressesModel268d39_CanDropMimeData
func callbackAddressesModel268d39_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressesModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_CanFetchMore
func callbackAddressesModel268d39_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressesModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_ColumnCount
func callbackAddressesModel268d39_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewAddressesModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *AddressesModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackAddressesModel268d39_ColumnsAboutToBeInserted
func callbackAddressesModel268d39_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressesModel268d39_ColumnsAboutToBeMoved
func callbackAddressesModel268d39_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackAddressesModel268d39_ColumnsAboutToBeRemoved
func callbackAddressesModel268d39_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressesModel268d39_ColumnsInserted
func callbackAddressesModel268d39_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressesModel268d39_ColumnsMoved
func callbackAddressesModel268d39_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackAddressesModel268d39_ColumnsRemoved
func callbackAddressesModel268d39_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressesModel268d39_Data
func callbackAddressesModel268d39_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewAddressesModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *AddressesModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressesModel268d39_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackAddressesModel268d39_DataChanged
func callbackAddressesModel268d39_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(signal))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackAddressesModel268d39_FetchMore
func callbackAddressesModel268d39_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewAddressesModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *AddressesModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackAddressesModel268d39_HasChildren
func callbackAddressesModel268d39_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressesModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_HeaderData
func callbackAddressesModel268d39_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewAddressesModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *AddressesModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressesModel268d39_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackAddressesModel268d39_HeaderDataChanged
func callbackAddressesModel268d39_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressesModel268d39_InsertColumns
func callbackAddressesModel268d39_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressesModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_InsertRows
func callbackAddressesModel268d39_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressesModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_ItemData
func callbackAddressesModel268d39_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*std_core.QModelIndex) map[int]*std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__itemData_newList())
		for k, v := range NewAddressesModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressesModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.AddressesModel268d39_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackAddressesModel268d39_LayoutAboutToBeChanged
func callbackAddressesModel268d39_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackAddressesModel268d39_LayoutChanged
func callbackAddressesModel268d39_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackAddressesModel268d39_Match
func callbackAddressesModel268d39_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__match_newList())
		for _, v := range NewAddressesModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressesModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.AddressesModel268d39_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackAddressesModel268d39_MimeData
func callbackAddressesModel268d39_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func([]*std_core.QModelIndex) *std_core.QMimeData)(signal))(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewAddressesModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewAddressesModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *AddressesModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.AddressesModel268d39_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__mimeData_indexes_newList())
			for _, v := range indexes {
				tmpList.__mimeData_indexes_setList(v)
			}
			return tmpList.Pointer()
		}()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackAddressesModel268d39_MimeTypes
func callbackAddressesModel268d39_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewAddressesModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *AddressesModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.AddressesModel268d39_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackAddressesModel268d39_ModelAboutToBeReset
func callbackAddressesModel268d39_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackAddressesModel268d39_ModelReset
func callbackAddressesModel268d39_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackAddressesModel268d39_MoveColumns
func callbackAddressesModel268d39_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *AddressesModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_MoveRows
func callbackAddressesModel268d39_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *AddressesModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_Parent
func callbackAddressesModel268d39_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewAddressesModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressesModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressesModel268d39_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressesModel268d39_RemoveColumns
func callbackAddressesModel268d39_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressesModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_RemoveRows
func callbackAddressesModel268d39_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressesModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_ResetInternalData
func callbackAddressesModel268d39_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressesModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *AddressesModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackAddressesModel268d39_Revert
func callbackAddressesModel268d39_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressesModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *AddressesModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_RevertDefault(ptr.Pointer())
	}
}

//export callbackAddressesModel268d39_RoleNames
func callbackAddressesModel268d39_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewAddressesModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressesModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.AddressesModel268d39_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackAddressesModel268d39_RowCount
func callbackAddressesModel268d39_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewAddressesModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *AddressesModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressesModel268d39_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackAddressesModel268d39_RowsAboutToBeInserted
func callbackAddressesModel268d39_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackAddressesModel268d39_RowsAboutToBeMoved
func callbackAddressesModel268d39_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackAddressesModel268d39_RowsAboutToBeRemoved
func callbackAddressesModel268d39_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressesModel268d39_RowsInserted
func callbackAddressesModel268d39_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressesModel268d39_RowsMoved
func callbackAddressesModel268d39_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackAddressesModel268d39_RowsRemoved
func callbackAddressesModel268d39_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressesModel268d39_SetData
func callbackAddressesModel268d39_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *AddressesModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_SetHeaderData
func callbackAddressesModel268d39_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *AddressesModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_SetItemData
func callbackAddressesModel268d39_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(signal))(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewAddressesModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewAddressesModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *AddressesModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewAddressesModelFromPointer(NewAddressesModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackAddressesModel268d39_Sort
func callbackAddressesModel268d39_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewAddressesModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *AddressesModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackAddressesModel268d39_Span
func callbackAddressesModel268d39_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewAddressesModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressesModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.AddressesModel268d39_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackAddressesModel268d39_Submit
func callbackAddressesModel268d39_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *AddressesModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackAddressesModel268d39_SupportedDragActions
func callbackAddressesModel268d39_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewAddressesModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *AddressesModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.AddressesModel268d39_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackAddressesModel268d39_SupportedDropActions
func callbackAddressesModel268d39_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewAddressesModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *AddressesModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.AddressesModel268d39_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackAddressesModel268d39_ChildEvent
func callbackAddressesModel268d39_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewAddressesModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *AddressesModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackAddressesModel268d39_ConnectNotify
func callbackAddressesModel268d39_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressesModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressesModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressesModel268d39_CustomEvent
func callbackAddressesModel268d39_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewAddressesModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *AddressesModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackAddressesModel268d39_DeleteLater
func callbackAddressesModel268d39_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressesModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *AddressesModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressesModel268d39_Destroyed
func callbackAddressesModel268d39_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackAddressesModel268d39_DisconnectNotify
func callbackAddressesModel268d39_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressesModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressesModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressesModel268d39_Event
func callbackAddressesModel268d39_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *AddressesModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_EventFilter
func callbackAddressesModel268d39_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressesModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *AddressesModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressesModel268d39_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackAddressesModel268d39_ObjectNameChanged
func callbackAddressesModel268d39_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackAddressesModel268d39_TimerEvent
func callbackAddressesModel268d39_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewAddressesModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *AddressesModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressesModel268d39_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type QAddress_ITF interface {
	std_core.QObject_ITF
	QAddress_PTR() *QAddress
}

func (ptr *QAddress) QAddress_PTR() *QAddress {
	return ptr
}

func (ptr *QAddress) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAddress) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQAddress(ptr QAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAddress_PTR().Pointer()
	}
	return nil
}

func NewQAddressFromPointer(ptr unsafe.Pointer) (n *QAddress) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QAddress)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QAddress:
			n = deduced

		case *std_core.QObject:
			n = &QAddress{QObject: *deduced}

		default:
			n = new(QAddress)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackQAddress268d39_Constructor
func callbackQAddress268d39_Constructor(ptr unsafe.Pointer) {
	this := NewQAddressFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackQAddress268d39_Address
func callbackQAddress268d39_Address(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "address"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQAddressFromPointer(ptr).AddressDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QAddress) ConnectAddress(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "address"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "address", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "address", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAddress) DisconnectAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "address")
	}
}

func (ptr *QAddress) Address() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAddress268d39_Address(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAddress) AddressDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QAddress268d39_AddressDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQAddress268d39_SetAddress
func callbackQAddress268d39_SetAddress(ptr unsafe.Pointer, address C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setAddress"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(address))
	} else {
		NewQAddressFromPointer(ptr).SetAddressDefault(cGoUnpackString(address))
	}
}

func (ptr *QAddress) ConnectSetAddress(f func(address string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddress"); signal != nil {
			f := func(address string) {
				(*(*func(string))(signal))(address)
				f(address)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAddress) DisconnectSetAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddress")
	}
}

func (ptr *QAddress) SetAddress(address string) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		C.QAddress268d39_SetAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))})
	}
}

func (ptr *QAddress) SetAddressDefault(address string) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		C.QAddress268d39_SetAddressDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))})
	}
}

//export callbackQAddress268d39_AddressChanged
func callbackQAddress268d39_AddressChanged(ptr unsafe.Pointer, address C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addressChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(address))
	}

}

func (ptr *QAddress) ConnectAddressChanged(f func(address string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressChanged") {
			C.QAddress268d39_ConnectAddressChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressChanged"); signal != nil {
			f := func(address string) {
				(*(*func(string))(signal))(address)
				f(address)
			}
			qt.ConnectSignal(ptr.Pointer(), "addressChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAddress) DisconnectAddressChanged() {
	if ptr.Pointer() != nil {
		C.QAddress268d39_DisconnectAddressChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressChanged")
	}
}

func (ptr *QAddress) AddressChanged(address string) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		C.QAddress268d39_AddressChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))})
	}
}

//export callbackQAddress268d39_AddressSky
func callbackQAddress268d39_AddressSky(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "addressSky"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQAddressFromPointer(ptr).AddressSkyDefault()))
}

func (ptr *QAddress) ConnectAddressSky(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addressSky"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "addressSky", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressSky", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAddress) DisconnectAddressSky() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addressSky")
	}
}

func (ptr *QAddress) AddressSky() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAddress268d39_AddressSky(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAddress) AddressSkyDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAddress268d39_AddressSkyDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQAddress268d39_SetAddressSky
func callbackQAddress268d39_SetAddressSky(ptr unsafe.Pointer, addressSky C.int) {
	if signal := qt.GetSignal(ptr, "setAddressSky"); signal != nil {
		(*(*func(int))(signal))(int(int32(addressSky)))
	} else {
		NewQAddressFromPointer(ptr).SetAddressSkyDefault(int(int32(addressSky)))
	}
}

func (ptr *QAddress) ConnectSetAddressSky(f func(addressSky int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddressSky"); signal != nil {
			f := func(addressSky int) {
				(*(*func(int))(signal))(addressSky)
				f(addressSky)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAddressSky", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAddressSky", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAddress) DisconnectSetAddressSky() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddressSky")
	}
}

func (ptr *QAddress) SetAddressSky(addressSky int) {
	if ptr.Pointer() != nil {
		C.QAddress268d39_SetAddressSky(ptr.Pointer(), C.int(int32(addressSky)))
	}
}

func (ptr *QAddress) SetAddressSkyDefault(addressSky int) {
	if ptr.Pointer() != nil {
		C.QAddress268d39_SetAddressSkyDefault(ptr.Pointer(), C.int(int32(addressSky)))
	}
}

//export callbackQAddress268d39_AddressSkyChanged
func callbackQAddress268d39_AddressSkyChanged(ptr unsafe.Pointer, addressSky C.int) {
	if signal := qt.GetSignal(ptr, "addressSkyChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(addressSky)))
	}

}

func (ptr *QAddress) ConnectAddressSkyChanged(f func(addressSky int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressSkyChanged") {
			C.QAddress268d39_ConnectAddressSkyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressSkyChanged"); signal != nil {
			f := func(addressSky int) {
				(*(*func(int))(signal))(addressSky)
				f(addressSky)
			}
			qt.ConnectSignal(ptr.Pointer(), "addressSkyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressSkyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAddress) DisconnectAddressSkyChanged() {
	if ptr.Pointer() != nil {
		C.QAddress268d39_DisconnectAddressSkyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressSkyChanged")
	}
}

func (ptr *QAddress) AddressSkyChanged(addressSky int) {
	if ptr.Pointer() != nil {
		C.QAddress268d39_AddressSkyChanged(ptr.Pointer(), C.int(int32(addressSky)))
	}
}

//export callbackQAddress268d39_AddressCoinHours
func callbackQAddress268d39_AddressCoinHours(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "addressCoinHours"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQAddressFromPointer(ptr).AddressCoinHoursDefault()))
}

func (ptr *QAddress) ConnectAddressCoinHours(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addressCoinHours"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "addressCoinHours", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressCoinHours", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAddress) DisconnectAddressCoinHours() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addressCoinHours")
	}
}

func (ptr *QAddress) AddressCoinHours() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAddress268d39_AddressCoinHours(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAddress) AddressCoinHoursDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QAddress268d39_AddressCoinHoursDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQAddress268d39_SetAddressCoinHours
func callbackQAddress268d39_SetAddressCoinHours(ptr unsafe.Pointer, addressCoinHours C.int) {
	if signal := qt.GetSignal(ptr, "setAddressCoinHours"); signal != nil {
		(*(*func(int))(signal))(int(int32(addressCoinHours)))
	} else {
		NewQAddressFromPointer(ptr).SetAddressCoinHoursDefault(int(int32(addressCoinHours)))
	}
}

func (ptr *QAddress) ConnectSetAddressCoinHours(f func(addressCoinHours int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddressCoinHours"); signal != nil {
			f := func(addressCoinHours int) {
				(*(*func(int))(signal))(addressCoinHours)
				f(addressCoinHours)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAddressCoinHours", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAddressCoinHours", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAddress) DisconnectSetAddressCoinHours() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddressCoinHours")
	}
}

func (ptr *QAddress) SetAddressCoinHours(addressCoinHours int) {
	if ptr.Pointer() != nil {
		C.QAddress268d39_SetAddressCoinHours(ptr.Pointer(), C.int(int32(addressCoinHours)))
	}
}

func (ptr *QAddress) SetAddressCoinHoursDefault(addressCoinHours int) {
	if ptr.Pointer() != nil {
		C.QAddress268d39_SetAddressCoinHoursDefault(ptr.Pointer(), C.int(int32(addressCoinHours)))
	}
}

//export callbackQAddress268d39_AddressCoinHoursChanged
func callbackQAddress268d39_AddressCoinHoursChanged(ptr unsafe.Pointer, addressCoinHours C.int) {
	if signal := qt.GetSignal(ptr, "addressCoinHoursChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(addressCoinHours)))
	}

}

func (ptr *QAddress) ConnectAddressCoinHoursChanged(f func(addressCoinHours int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressCoinHoursChanged") {
			C.QAddress268d39_ConnectAddressCoinHoursChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressCoinHoursChanged"); signal != nil {
			f := func(addressCoinHours int) {
				(*(*func(int))(signal))(addressCoinHours)
				f(addressCoinHours)
			}
			qt.ConnectSignal(ptr.Pointer(), "addressCoinHoursChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressCoinHoursChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAddress) DisconnectAddressCoinHoursChanged() {
	if ptr.Pointer() != nil {
		C.QAddress268d39_DisconnectAddressCoinHoursChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressCoinHoursChanged")
	}
}

func (ptr *QAddress) AddressCoinHoursChanged(addressCoinHours int) {
	if ptr.Pointer() != nil {
		C.QAddress268d39_AddressCoinHoursChanged(ptr.Pointer(), C.int(int32(addressCoinHours)))
	}
}

func QAddress_QRegisterMetaType() int {
	return int(int32(C.QAddress268d39_QAddress268d39_QRegisterMetaType()))
}

func (ptr *QAddress) QRegisterMetaType() int {
	return int(int32(C.QAddress268d39_QAddress268d39_QRegisterMetaType()))
}

func QAddress_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QAddress268d39_QAddress268d39_QRegisterMetaType2(typeNameC)))
}

func (ptr *QAddress) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QAddress268d39_QAddress268d39_QRegisterMetaType2(typeNameC)))
}

func QAddress_QmlRegisterType() int {
	return int(int32(C.QAddress268d39_QAddress268d39_QmlRegisterType()))
}

func (ptr *QAddress) QmlRegisterType() int {
	return int(int32(C.QAddress268d39_QAddress268d39_QmlRegisterType()))
}

func QAddress_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QAddress268d39_QAddress268d39_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QAddress) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QAddress268d39_QAddress268d39_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QAddress) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QAddress268d39___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAddress) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAddress268d39___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QAddress) __children_newList() unsafe.Pointer {
	return C.QAddress268d39___children_newList(ptr.Pointer())
}

func (ptr *QAddress) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.QAddress268d39___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QAddress) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QAddress268d39___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QAddress) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QAddress268d39___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QAddress) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QAddress268d39___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAddress) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAddress268d39___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QAddress) __findChildren_newList() unsafe.Pointer {
	return C.QAddress268d39___findChildren_newList(ptr.Pointer())
}

func (ptr *QAddress) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QAddress268d39___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAddress) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAddress268d39___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QAddress) __findChildren_newList3() unsafe.Pointer {
	return C.QAddress268d39___findChildren_newList3(ptr.Pointer())
}

func (ptr *QAddress) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QAddress268d39___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QAddress) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QAddress268d39___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QAddress) __qFindChildren_newList2() unsafe.Pointer {
	return C.QAddress268d39___qFindChildren_newList2(ptr.Pointer())
}

func NewQAddress(parent std_core.QObject_ITF) *QAddress {
	tmpValue := NewQAddressFromPointer(C.QAddress268d39_NewQAddress(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAddress268d39_DestroyQAddress
func callbackQAddress268d39_DestroyQAddress(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QAddress"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAddressFromPointer(ptr).DestroyQAddressDefault()
	}
}

func (ptr *QAddress) ConnectDestroyQAddress(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QAddress"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QAddress) DisconnectDestroyQAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QAddress")
	}
}

func (ptr *QAddress) DestroyQAddress() {
	if ptr.Pointer() != nil {
		C.QAddress268d39_DestroyQAddress(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QAddress) DestroyQAddressDefault() {
	if ptr.Pointer() != nil {
		C.QAddress268d39_DestroyQAddressDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQAddress268d39_ChildEvent
func callbackQAddress268d39_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQAddressFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAddress) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAddress268d39_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQAddress268d39_ConnectNotify
func callbackQAddress268d39_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAddressFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAddress) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAddress268d39_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAddress268d39_CustomEvent
func callbackQAddress268d39_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewQAddressFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QAddress) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAddress268d39_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQAddress268d39_DeleteLater
func callbackQAddress268d39_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQAddressFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAddress) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QAddress268d39_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQAddress268d39_Destroyed
func callbackQAddress268d39_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackQAddress268d39_DisconnectNotify
func callbackQAddress268d39_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAddressFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAddress) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QAddress268d39_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAddress268d39_Event
func callbackQAddress268d39_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAddressFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QAddress) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAddress268d39_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQAddress268d39_EventFilter
func callbackQAddress268d39_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAddressFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QAddress) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QAddress268d39_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQAddress268d39_ObjectNameChanged
func callbackQAddress268d39_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQAddress268d39_TimerEvent
func callbackQAddress268d39_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQAddressFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAddress) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QAddress268d39_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type QWallet_ITF interface {
	std_core.QObject_ITF
	QWallet_PTR() *QWallet
}

func (ptr *QWallet) QWallet_PTR() *QWallet {
	return ptr
}

func (ptr *QWallet) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWallet) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWallet(ptr QWallet_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWallet_PTR().Pointer()
	}
	return nil
}

func NewQWalletFromPointer(ptr unsafe.Pointer) (n *QWallet) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QWallet)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QWallet:
			n = deduced

		case *std_core.QObject:
			n = &QWallet{QObject: *deduced}

		default:
			n = new(QWallet)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackQWallet268d39_Constructor
func callbackQWallet268d39_Constructor(ptr unsafe.Pointer) {
	this := NewQWalletFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackQWallet268d39_Name
func callbackQWallet268d39_Name(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "name"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQWalletFromPointer(ptr).NameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QWallet) ConnectName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "name"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "name", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "name", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "name")
	}
}

func (ptr *QWallet) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWallet268d39_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWallet) NameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWallet268d39_NameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQWallet268d39_SetName
func callbackQWallet268d39_SetName(ptr unsafe.Pointer, name C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setName"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(name))
	} else {
		NewQWalletFromPointer(ptr).SetNameDefault(cGoUnpackString(name))
	}
}

func (ptr *QWallet) ConnectSetName(f func(name string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setName"); signal != nil {
			f := func(name string) {
				(*(*func(string))(signal))(name)
				f(name)
			}
			qt.ConnectSignal(ptr.Pointer(), "setName", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setName", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectSetName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setName")
	}
}

func (ptr *QWallet) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QWallet268d39_SetName(ptr.Pointer(), C.struct_Moc_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QWallet) SetNameDefault(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QWallet268d39_SetNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

//export callbackQWallet268d39_NameChanged
func callbackQWallet268d39_NameChanged(ptr unsafe.Pointer, name C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "nameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(name))
	}

}

func (ptr *QWallet) ConnectNameChanged(f func(name string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "nameChanged") {
			C.QWallet268d39_ConnectNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "nameChanged"); signal != nil {
			f := func(name string) {
				(*(*func(string))(signal))(name)
				f(name)
			}
			qt.ConnectSignal(ptr.Pointer(), "nameChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nameChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectNameChanged() {
	if ptr.Pointer() != nil {
		C.QWallet268d39_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "nameChanged")
	}
}

func (ptr *QWallet) NameChanged(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QWallet268d39_NameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

//export callbackQWallet268d39_EncryptionEnabled
func callbackQWallet268d39_EncryptionEnabled(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "encryptionEnabled"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQWalletFromPointer(ptr).EncryptionEnabledDefault()))
}

func (ptr *QWallet) ConnectEncryptionEnabled(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "encryptionEnabled"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "encryptionEnabled", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "encryptionEnabled", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectEncryptionEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "encryptionEnabled")
	}
}

func (ptr *QWallet) EncryptionEnabled() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWallet268d39_EncryptionEnabled(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWallet) EncryptionEnabledDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWallet268d39_EncryptionEnabledDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQWallet268d39_SetEncryptionEnabled
func callbackQWallet268d39_SetEncryptionEnabled(ptr unsafe.Pointer, encryptionEnabled C.int) {
	if signal := qt.GetSignal(ptr, "setEncryptionEnabled"); signal != nil {
		(*(*func(int))(signal))(int(int32(encryptionEnabled)))
	} else {
		NewQWalletFromPointer(ptr).SetEncryptionEnabledDefault(int(int32(encryptionEnabled)))
	}
}

func (ptr *QWallet) ConnectSetEncryptionEnabled(f func(encryptionEnabled int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setEncryptionEnabled"); signal != nil {
			f := func(encryptionEnabled int) {
				(*(*func(int))(signal))(encryptionEnabled)
				f(encryptionEnabled)
			}
			qt.ConnectSignal(ptr.Pointer(), "setEncryptionEnabled", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setEncryptionEnabled", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectSetEncryptionEnabled() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setEncryptionEnabled")
	}
}

func (ptr *QWallet) SetEncryptionEnabled(encryptionEnabled int) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_SetEncryptionEnabled(ptr.Pointer(), C.int(int32(encryptionEnabled)))
	}
}

func (ptr *QWallet) SetEncryptionEnabledDefault(encryptionEnabled int) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_SetEncryptionEnabledDefault(ptr.Pointer(), C.int(int32(encryptionEnabled)))
	}
}

//export callbackQWallet268d39_EncryptionEnabledChanged
func callbackQWallet268d39_EncryptionEnabledChanged(ptr unsafe.Pointer, encryptionEnabled C.int) {
	if signal := qt.GetSignal(ptr, "encryptionEnabledChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(encryptionEnabled)))
	}

}

func (ptr *QWallet) ConnectEncryptionEnabledChanged(f func(encryptionEnabled int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "encryptionEnabledChanged") {
			C.QWallet268d39_ConnectEncryptionEnabledChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "encryptionEnabledChanged"); signal != nil {
			f := func(encryptionEnabled int) {
				(*(*func(int))(signal))(encryptionEnabled)
				f(encryptionEnabled)
			}
			qt.ConnectSignal(ptr.Pointer(), "encryptionEnabledChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "encryptionEnabledChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectEncryptionEnabledChanged() {
	if ptr.Pointer() != nil {
		C.QWallet268d39_DisconnectEncryptionEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "encryptionEnabledChanged")
	}
}

func (ptr *QWallet) EncryptionEnabledChanged(encryptionEnabled int) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_EncryptionEnabledChanged(ptr.Pointer(), C.int(int32(encryptionEnabled)))
	}
}

//export callbackQWallet268d39_Sky
func callbackQWallet268d39_Sky(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "sky"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQWalletFromPointer(ptr).SkyDefault()))
}

func (ptr *QWallet) ConnectSky(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sky"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "sky", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sky", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectSky() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sky")
	}
}

func (ptr *QWallet) Sky() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWallet268d39_Sky(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWallet) SkyDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWallet268d39_SkyDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQWallet268d39_SetSky
func callbackQWallet268d39_SetSky(ptr unsafe.Pointer, sky C.int) {
	if signal := qt.GetSignal(ptr, "setSky"); signal != nil {
		(*(*func(int))(signal))(int(int32(sky)))
	} else {
		NewQWalletFromPointer(ptr).SetSkyDefault(int(int32(sky)))
	}
}

func (ptr *QWallet) ConnectSetSky(f func(sky int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSky"); signal != nil {
			f := func(sky int) {
				(*(*func(int))(signal))(sky)
				f(sky)
			}
			qt.ConnectSignal(ptr.Pointer(), "setSky", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSky", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectSetSky() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSky")
	}
}

func (ptr *QWallet) SetSky(sky int) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_SetSky(ptr.Pointer(), C.int(int32(sky)))
	}
}

func (ptr *QWallet) SetSkyDefault(sky int) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_SetSkyDefault(ptr.Pointer(), C.int(int32(sky)))
	}
}

//export callbackQWallet268d39_SkyChanged
func callbackQWallet268d39_SkyChanged(ptr unsafe.Pointer, sky C.int) {
	if signal := qt.GetSignal(ptr, "skyChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(sky)))
	}

}

func (ptr *QWallet) ConnectSkyChanged(f func(sky int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "skyChanged") {
			C.QWallet268d39_ConnectSkyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "skyChanged"); signal != nil {
			f := func(sky int) {
				(*(*func(int))(signal))(sky)
				f(sky)
			}
			qt.ConnectSignal(ptr.Pointer(), "skyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "skyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectSkyChanged() {
	if ptr.Pointer() != nil {
		C.QWallet268d39_DisconnectSkyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "skyChanged")
	}
}

func (ptr *QWallet) SkyChanged(sky int) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_SkyChanged(ptr.Pointer(), C.int(int32(sky)))
	}
}

//export callbackQWallet268d39_CoinHours
func callbackQWallet268d39_CoinHours(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "coinHours"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewQWalletFromPointer(ptr).CoinHoursDefault()))
}

func (ptr *QWallet) ConnectCoinHours(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "coinHours"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "coinHours", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "coinHours", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectCoinHours() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "coinHours")
	}
}

func (ptr *QWallet) CoinHours() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWallet268d39_CoinHours(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWallet) CoinHoursDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWallet268d39_CoinHoursDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQWallet268d39_SetCoinHours
func callbackQWallet268d39_SetCoinHours(ptr unsafe.Pointer, coinHours C.int) {
	if signal := qt.GetSignal(ptr, "setCoinHours"); signal != nil {
		(*(*func(int))(signal))(int(int32(coinHours)))
	} else {
		NewQWalletFromPointer(ptr).SetCoinHoursDefault(int(int32(coinHours)))
	}
}

func (ptr *QWallet) ConnectSetCoinHours(f func(coinHours int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCoinHours"); signal != nil {
			f := func(coinHours int) {
				(*(*func(int))(signal))(coinHours)
				f(coinHours)
			}
			qt.ConnectSignal(ptr.Pointer(), "setCoinHours", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setCoinHours", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectSetCoinHours() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setCoinHours")
	}
}

func (ptr *QWallet) SetCoinHours(coinHours int) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_SetCoinHours(ptr.Pointer(), C.int(int32(coinHours)))
	}
}

func (ptr *QWallet) SetCoinHoursDefault(coinHours int) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_SetCoinHoursDefault(ptr.Pointer(), C.int(int32(coinHours)))
	}
}

//export callbackQWallet268d39_CoinHoursChanged
func callbackQWallet268d39_CoinHoursChanged(ptr unsafe.Pointer, coinHours C.int) {
	if signal := qt.GetSignal(ptr, "coinHoursChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(coinHours)))
	}

}

func (ptr *QWallet) ConnectCoinHoursChanged(f func(coinHours int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "coinHoursChanged") {
			C.QWallet268d39_ConnectCoinHoursChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "coinHoursChanged"); signal != nil {
			f := func(coinHours int) {
				(*(*func(int))(signal))(coinHours)
				f(coinHours)
			}
			qt.ConnectSignal(ptr.Pointer(), "coinHoursChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "coinHoursChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectCoinHoursChanged() {
	if ptr.Pointer() != nil {
		C.QWallet268d39_DisconnectCoinHoursChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "coinHoursChanged")
	}
}

func (ptr *QWallet) CoinHoursChanged(coinHours int) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_CoinHoursChanged(ptr.Pointer(), C.int(int32(coinHours)))
	}
}

//export callbackQWallet268d39_FileName
func callbackQWallet268d39_FileName(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "fileName"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQWalletFromPointer(ptr).FileNameDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QWallet) ConnectFileName(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "fileName"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "fileName", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fileName", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectFileName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "fileName")
	}
}

func (ptr *QWallet) FileName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWallet268d39_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWallet) FileNameDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWallet268d39_FileNameDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQWallet268d39_SetFileName
func callbackQWallet268d39_SetFileName(ptr unsafe.Pointer, fileName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setFileName"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(fileName))
	} else {
		NewQWalletFromPointer(ptr).SetFileNameDefault(cGoUnpackString(fileName))
	}
}

func (ptr *QWallet) ConnectSetFileName(f func(fileName string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFileName"); signal != nil {
			f := func(fileName string) {
				(*(*func(string))(signal))(fileName)
				f(fileName)
			}
			qt.ConnectSignal(ptr.Pointer(), "setFileName", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFileName", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectSetFileName() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFileName")
	}
}

func (ptr *QWallet) SetFileName(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QWallet268d39_SetFileName(ptr.Pointer(), C.struct_Moc_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

func (ptr *QWallet) SetFileNameDefault(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QWallet268d39_SetFileNameDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

//export callbackQWallet268d39_FileNameChanged
func callbackQWallet268d39_FileNameChanged(ptr unsafe.Pointer, fileName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "fileNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(fileName))
	}

}

func (ptr *QWallet) ConnectFileNameChanged(f func(fileName string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "fileNameChanged") {
			C.QWallet268d39_ConnectFileNameChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "fileNameChanged"); signal != nil {
			f := func(fileName string) {
				(*(*func(string))(signal))(fileName)
				f(fileName)
			}
			qt.ConnectSignal(ptr.Pointer(), "fileNameChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "fileNameChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectFileNameChanged() {
	if ptr.Pointer() != nil {
		C.QWallet268d39_DisconnectFileNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "fileNameChanged")
	}
}

func (ptr *QWallet) FileNameChanged(fileName string) {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		C.QWallet268d39_FileNameChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: fileNameC, len: C.longlong(len(fileName))})
	}
}

func QWallet_QRegisterMetaType() int {
	return int(int32(C.QWallet268d39_QWallet268d39_QRegisterMetaType()))
}

func (ptr *QWallet) QRegisterMetaType() int {
	return int(int32(C.QWallet268d39_QWallet268d39_QRegisterMetaType()))
}

func QWallet_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QWallet268d39_QWallet268d39_QRegisterMetaType2(typeNameC)))
}

func (ptr *QWallet) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QWallet268d39_QWallet268d39_QRegisterMetaType2(typeNameC)))
}

func QWallet_QmlRegisterType() int {
	return int(int32(C.QWallet268d39_QWallet268d39_QmlRegisterType()))
}

func (ptr *QWallet) QmlRegisterType() int {
	return int(int32(C.QWallet268d39_QWallet268d39_QmlRegisterType()))
}

func QWallet_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QWallet268d39_QWallet268d39_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QWallet) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QWallet268d39_QWallet268d39_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QWallet) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QWallet268d39___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWallet) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWallet268d39___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QWallet) __children_newList() unsafe.Pointer {
	return C.QWallet268d39___children_newList(ptr.Pointer())
}

func (ptr *QWallet) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.QWallet268d39___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWallet) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWallet268d39___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QWallet) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWallet268d39___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWallet) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QWallet268d39___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWallet) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWallet268d39___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QWallet) __findChildren_newList() unsafe.Pointer {
	return C.QWallet268d39___findChildren_newList(ptr.Pointer())
}

func (ptr *QWallet) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QWallet268d39___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWallet) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWallet268d39___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QWallet) __findChildren_newList3() unsafe.Pointer {
	return C.QWallet268d39___findChildren_newList3(ptr.Pointer())
}

func (ptr *QWallet) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QWallet268d39___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWallet) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWallet268d39___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QWallet) __qFindChildren_newList2() unsafe.Pointer {
	return C.QWallet268d39___qFindChildren_newList2(ptr.Pointer())
}

func NewQWallet(parent std_core.QObject_ITF) *QWallet {
	tmpValue := NewQWalletFromPointer(C.QWallet268d39_NewQWallet(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWallet268d39_DestroyQWallet
func callbackQWallet268d39_DestroyQWallet(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWallet"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWalletFromPointer(ptr).DestroyQWalletDefault()
	}
}

func (ptr *QWallet) ConnectDestroyQWallet(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWallet"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWallet", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWallet", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWallet) DisconnectDestroyQWallet() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWallet")
	}
}

func (ptr *QWallet) DestroyQWallet() {
	if ptr.Pointer() != nil {
		C.QWallet268d39_DestroyQWallet(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QWallet) DestroyQWalletDefault() {
	if ptr.Pointer() != nil {
		C.QWallet268d39_DestroyQWalletDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWallet268d39_ChildEvent
func callbackQWallet268d39_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQWalletFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWallet) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQWallet268d39_ConnectNotify
func callbackQWallet268d39_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWalletFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWallet) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWallet268d39_CustomEvent
func callbackQWallet268d39_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewQWalletFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QWallet) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQWallet268d39_DeleteLater
func callbackQWallet268d39_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWalletFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWallet) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWallet268d39_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQWallet268d39_Destroyed
func callbackQWallet268d39_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackQWallet268d39_DisconnectNotify
func callbackQWallet268d39_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWalletFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWallet) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWallet268d39_Event
func callbackQWallet268d39_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWalletFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QWallet) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWallet268d39_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWallet268d39_EventFilter
func callbackQWallet268d39_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWalletFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QWallet) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWallet268d39_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWallet268d39_ObjectNameChanged
func callbackQWallet268d39_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWallet268d39_TimerEvent
func callbackQWallet268d39_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQWalletFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWallet) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWallet268d39_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

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

	custom_models_baab1em "github.com/FiberCryptoWallet/src/coin/skycoin/models"
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

//export callbackWalletManager3cf9d2_Constructor
func callbackWalletManager3cf9d2_Constructor(ptr unsafe.Pointer) {
	this := NewWalletManagerFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackWalletManager3cf9d2_CreateEncryptedWallet
func callbackWalletManager3cf9d2_CreateEncryptedWallet(ptr unsafe.Pointer, walletM unsafe.Pointer, seed C.struct_Moc_PackedString, label C.struct_Moc_PackedString, password C.struct_Moc_PackedString, scanN C.int) {
	if signal := qt.GetSignal(ptr, "createEncryptedWallet"); signal != nil {
		(*(*func(*custom_models_baab1em.WalletModel, string, string, string, int))(signal))(custom_models_baab1em.NewWalletModelFromPointer(walletM), cGoUnpackString(seed), cGoUnpackString(label), cGoUnpackString(password), int(int32(scanN)))
	}

}

func (ptr *WalletManager) ConnectCreateEncryptedWallet(f func(walletM *custom_models_baab1em.WalletModel, seed string, label string, password string, scanN int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createEncryptedWallet"); signal != nil {
			f := func(walletM *custom_models_baab1em.WalletModel, seed string, label string, password string, scanN int) {
				(*(*func(*custom_models_baab1em.WalletModel, string, string, string, int))(signal))(walletM, seed, label, password, scanN)
				f(walletM, seed, label, password, scanN)
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

func (ptr *WalletManager) CreateEncryptedWallet(walletM custom_models_baab1em.WalletModel_ITF, seed string, label string, password string, scanN int) {
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
		C.WalletManager3cf9d2_CreateEncryptedWallet(ptr.Pointer(), custom_models_baab1em.PointerFromWalletModel(walletM), C.struct_Moc_PackedString{data: seedC, len: C.longlong(len(seed))}, C.struct_Moc_PackedString{data: labelC, len: C.longlong(len(label))}, C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))}, C.int(int32(scanN)))
	}
}

//export callbackWalletManager3cf9d2_CreateUnencryptedWallet
func callbackWalletManager3cf9d2_CreateUnencryptedWallet(ptr unsafe.Pointer, walletM unsafe.Pointer, seed C.struct_Moc_PackedString, label C.struct_Moc_PackedString, scanN C.int) {
	if signal := qt.GetSignal(ptr, "createUnencryptedWallet"); signal != nil {
		(*(*func(*custom_models_baab1em.WalletModel, string, string, int))(signal))(custom_models_baab1em.NewWalletModelFromPointer(walletM), cGoUnpackString(seed), cGoUnpackString(label), int(int32(scanN)))
	}

}

func (ptr *WalletManager) ConnectCreateUnencryptedWallet(f func(walletM *custom_models_baab1em.WalletModel, seed string, label string, scanN int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createUnencryptedWallet"); signal != nil {
			f := func(walletM *custom_models_baab1em.WalletModel, seed string, label string, scanN int) {
				(*(*func(*custom_models_baab1em.WalletModel, string, string, int))(signal))(walletM, seed, label, scanN)
				f(walletM, seed, label, scanN)
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

func (ptr *WalletManager) CreateUnencryptedWallet(walletM custom_models_baab1em.WalletModel_ITF, seed string, label string, scanN int) {
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
		C.WalletManager3cf9d2_CreateUnencryptedWallet(ptr.Pointer(), custom_models_baab1em.PointerFromWalletModel(walletM), C.struct_Moc_PackedString{data: seedC, len: C.longlong(len(seed))}, C.struct_Moc_PackedString{data: labelC, len: C.longlong(len(label))}, C.int(int32(scanN)))
	}
}

//export callbackWalletManager3cf9d2_GetNewSeed
func callbackWalletManager3cf9d2_GetNewSeed(ptr unsafe.Pointer, entropy C.int) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.WalletManager3cf9d2_GetNewSeed(ptr.Pointer(), C.int(int32(entropy))))
	}
	return ""
}

//export callbackWalletManager3cf9d2_VerifySeed
func callbackWalletManager3cf9d2_VerifySeed(ptr unsafe.Pointer, seed C.struct_Moc_PackedString) C.int {
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
		return int(int32(C.WalletManager3cf9d2_VerifySeed(ptr.Pointer(), C.struct_Moc_PackedString{data: seedC, len: C.longlong(len(seed))})))
	}
	return 0
}

//export callbackWalletManager3cf9d2_NewWalletAddress
func callbackWalletManager3cf9d2_NewWalletAddress(ptr unsafe.Pointer, addressM unsafe.Pointer, id C.struct_Moc_PackedString, n C.int, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "newWalletAddress"); signal != nil {
		(*(*func(*custom_models_baab1em.AddressesModel, string, int, string))(signal))(custom_models_baab1em.NewAddressesModelFromPointer(addressM), cGoUnpackString(id), int(int32(n)), cGoUnpackString(password))
	}

}

func (ptr *WalletManager) ConnectNewWalletAddress(f func(addressM *custom_models_baab1em.AddressesModel, id string, n int, password string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "newWalletAddress"); signal != nil {
			f := func(addressM *custom_models_baab1em.AddressesModel, id string, n int, password string) {
				(*(*func(*custom_models_baab1em.AddressesModel, string, int, string))(signal))(addressM, id, n, password)
				f(addressM, id, n, password)
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

func (ptr *WalletManager) NewWalletAddress(addressM custom_models_baab1em.AddressesModel_ITF, id string, n int, password string) {
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
		C.WalletManager3cf9d2_NewWalletAddress(ptr.Pointer(), custom_models_baab1em.PointerFromAddressesModel(addressM), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))}, C.int(int32(n)), C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackWalletManager3cf9d2_EncryptWallet
func callbackWalletManager3cf9d2_EncryptWallet(ptr unsafe.Pointer, walletM unsafe.Pointer, id C.struct_Moc_PackedString, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "encryptWallet"); signal != nil {
		(*(*func(*custom_models_baab1em.WalletModel, string, string))(signal))(custom_models_baab1em.NewWalletModelFromPointer(walletM), cGoUnpackString(id), cGoUnpackString(password))
	}

}

func (ptr *WalletManager) ConnectEncryptWallet(f func(walletM *custom_models_baab1em.WalletModel, id string, password string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "encryptWallet"); signal != nil {
			f := func(walletM *custom_models_baab1em.WalletModel, id string, password string) {
				(*(*func(*custom_models_baab1em.WalletModel, string, string))(signal))(walletM, id, password)
				f(walletM, id, password)
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

func (ptr *WalletManager) EncryptWallet(walletM custom_models_baab1em.WalletModel_ITF, id string, password string) {
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
		C.WalletManager3cf9d2_EncryptWallet(ptr.Pointer(), custom_models_baab1em.PointerFromWalletModel(walletM), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))}, C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

//export callbackWalletManager3cf9d2_DecryptWallet
func callbackWalletManager3cf9d2_DecryptWallet(ptr unsafe.Pointer, walletM unsafe.Pointer, id C.struct_Moc_PackedString, password C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "decryptWallet"); signal != nil {
		(*(*func(*custom_models_baab1em.WalletModel, string, string))(signal))(custom_models_baab1em.NewWalletModelFromPointer(walletM), cGoUnpackString(id), cGoUnpackString(password))
	}

}

func (ptr *WalletManager) ConnectDecryptWallet(f func(walletM *custom_models_baab1em.WalletModel, id string, password string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "decryptWallet"); signal != nil {
			f := func(walletM *custom_models_baab1em.WalletModel, id string, password string) {
				(*(*func(*custom_models_baab1em.WalletModel, string, string))(signal))(walletM, id, password)
				f(walletM, id, password)
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

func (ptr *WalletManager) DecryptWallet(walletM custom_models_baab1em.WalletModel_ITF, id string, password string) {
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
		C.WalletManager3cf9d2_DecryptWallet(ptr.Pointer(), custom_models_baab1em.PointerFromWalletModel(walletM), C.struct_Moc_PackedString{data: idC, len: C.longlong(len(id))}, C.struct_Moc_PackedString{data: passwordC, len: C.longlong(len(password))})
	}
}

func WalletManager_QRegisterMetaType() int {
	return int(int32(C.WalletManager3cf9d2_WalletManager3cf9d2_QRegisterMetaType()))
}

func (ptr *WalletManager) QRegisterMetaType() int {
	return int(int32(C.WalletManager3cf9d2_WalletManager3cf9d2_QRegisterMetaType()))
}

func WalletManager_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.WalletManager3cf9d2_WalletManager3cf9d2_QRegisterMetaType2(typeNameC)))
}

func (ptr *WalletManager) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.WalletManager3cf9d2_WalletManager3cf9d2_QRegisterMetaType2(typeNameC)))
}

func WalletManager_QmlRegisterType() int {
	return int(int32(C.WalletManager3cf9d2_WalletManager3cf9d2_QmlRegisterType()))
}

func (ptr *WalletManager) QmlRegisterType() int {
	return int(int32(C.WalletManager3cf9d2_WalletManager3cf9d2_QmlRegisterType()))
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
	return int(int32(C.WalletManager3cf9d2_WalletManager3cf9d2_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.WalletManager3cf9d2_WalletManager3cf9d2_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *WalletManager) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletManager3cf9d2___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager3cf9d2___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletManager) __children_newList() unsafe.Pointer {
	return C.WalletManager3cf9d2___children_newList(ptr.Pointer())
}

func (ptr *WalletManager) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.WalletManager3cf9d2___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager3cf9d2___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *WalletManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.WalletManager3cf9d2___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *WalletManager) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletManager3cf9d2___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager3cf9d2___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletManager) __findChildren_newList() unsafe.Pointer {
	return C.WalletManager3cf9d2___findChildren_newList(ptr.Pointer())
}

func (ptr *WalletManager) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletManager3cf9d2___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager3cf9d2___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletManager) __findChildren_newList3() unsafe.Pointer {
	return C.WalletManager3cf9d2___findChildren_newList3(ptr.Pointer())
}

func (ptr *WalletManager) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.WalletManager3cf9d2___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *WalletManager) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager3cf9d2___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *WalletManager) __qFindChildren_newList2() unsafe.Pointer {
	return C.WalletManager3cf9d2___qFindChildren_newList2(ptr.Pointer())
}

func NewWalletManager(parent std_core.QObject_ITF) *WalletManager {
	tmpValue := NewWalletManagerFromPointer(C.WalletManager3cf9d2_NewWalletManager(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackWalletManager3cf9d2_DestroyWalletManager
func callbackWalletManager3cf9d2_DestroyWalletManager(ptr unsafe.Pointer) {
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
		C.WalletManager3cf9d2_DestroyWalletManager(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *WalletManager) DestroyWalletManagerDefault() {
	if ptr.Pointer() != nil {
		C.WalletManager3cf9d2_DestroyWalletManagerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackWalletManager3cf9d2_ChildEvent
func callbackWalletManager3cf9d2_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewWalletManagerFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *WalletManager) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager3cf9d2_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackWalletManager3cf9d2_ConnectNotify
func callbackWalletManager3cf9d2_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewWalletManagerFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *WalletManager) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager3cf9d2_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackWalletManager3cf9d2_CustomEvent
func callbackWalletManager3cf9d2_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewWalletManagerFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *WalletManager) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager3cf9d2_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackWalletManager3cf9d2_DeleteLater
func callbackWalletManager3cf9d2_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewWalletManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *WalletManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.WalletManager3cf9d2_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackWalletManager3cf9d2_Destroyed
func callbackWalletManager3cf9d2_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackWalletManager3cf9d2_DisconnectNotify
func callbackWalletManager3cf9d2_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewWalletManagerFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *WalletManager) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager3cf9d2_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackWalletManager3cf9d2_Event
func callbackWalletManager3cf9d2_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletManagerFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *WalletManager) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletManager3cf9d2_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackWalletManager3cf9d2_EventFilter
func callbackWalletManager3cf9d2_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewWalletManagerFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *WalletManager) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.WalletManager3cf9d2_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackWalletManager3cf9d2_ObjectNameChanged
func callbackWalletManager3cf9d2_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackWalletManager3cf9d2_TimerEvent
func callbackWalletManager3cf9d2_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewWalletManagerFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *WalletManager) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.WalletManager3cf9d2_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
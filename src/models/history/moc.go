package history

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	custom_transactions_ecff1cm "github.com/fibercrypto/FiberCryptoWallet/src/models/transactions"
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

type HistoryManager_ITF interface {
	std_core.QObject_ITF
	HistoryManager_PTR() *HistoryManager
}

func (ptr *HistoryManager) HistoryManager_PTR() *HistoryManager {
	return ptr
}

func (ptr *HistoryManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *HistoryManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromHistoryManager(ptr HistoryManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.HistoryManager_PTR().Pointer()
	}
	return nil
}

func NewHistoryManagerFromPointer(ptr unsafe.Pointer) (n *HistoryManager) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(HistoryManager)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *HistoryManager:
			n = deduced

		case *std_core.QObject:
			n = &HistoryManager{QObject: *deduced}

		default:
			n = new(HistoryManager)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *HistoryManager) Init() { this.init() }

//export callbackHistoryManager554044_Constructor
func callbackHistoryManager554044_Constructor(ptr unsafe.Pointer) {
	this := NewHistoryManagerFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackHistoryManager554044_LoadHistoryWithFilters
func callbackHistoryManager554044_LoadHistoryWithFilters(ptr unsafe.Pointer, addresses C.struct_Moc_PackedList, filterAddresses C.struct_Moc_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "loadHistoryWithFilters"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewHistoryManagerFromPointer(NewHistoryManagerFromPointer(nil).__loadHistoryWithFilters_newList())
			for _, v := range (*(*func(map[string]string, []string) []*custom_transactions_ecff1cm.TransactionDetails)(signal))(func(l C.struct_Moc_PackedList) map[string]string {
				out := make(map[string]string, int(l.len))
				tmpList := NewHistoryManagerFromPointer(l.data)
				for i, v := range tmpList.__loadHistoryWithFilters_addresses_keyList() {
					out[v] = tmpList.__loadHistoryWithFilters_addresses_atList(v, i).ToInterface().(string)
				}
				return out
			}(addresses), unpackStringList(cGoUnpackString(filterAddresses))) {
				tmpList.__loadHistoryWithFilters_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewHistoryManagerFromPointer(NewHistoryManagerFromPointer(nil).__loadHistoryWithFilters_newList())
		for _, v := range make([]*custom_transactions_ecff1cm.TransactionDetails, 0) {
			tmpList.__loadHistoryWithFilters_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *HistoryManager) ConnectLoadHistoryWithFilters(f func(addresses map[string]string, filterAddresses []string) []*custom_transactions_ecff1cm.TransactionDetails) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "loadHistoryWithFilters"); signal != nil {
			f := func(addresses map[string]string, filterAddresses []string) []*custom_transactions_ecff1cm.TransactionDetails {
				(*(*func(map[string]string, []string) []*custom_transactions_ecff1cm.TransactionDetails)(signal))(addresses, filterAddresses)
				return f(addresses, filterAddresses)
			}
			qt.ConnectSignal(ptr.Pointer(), "loadHistoryWithFilters", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadHistoryWithFilters", unsafe.Pointer(&f))
		}
	}
}

func (ptr *HistoryManager) DisconnectLoadHistoryWithFilters() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "loadHistoryWithFilters")
	}
}

func (ptr *HistoryManager) LoadHistoryWithFilters(addresses map[string]string, filterAddresses []string) []*custom_transactions_ecff1cm.TransactionDetails {
	if ptr.Pointer() != nil {
		filterAddressesC := C.CString(strings.Join(filterAddresses, "¡¦!"))
		defer C.free(unsafe.Pointer(filterAddressesC))
		return func(l C.struct_Moc_PackedList) []*custom_transactions_ecff1cm.TransactionDetails {
			out := make([]*custom_transactions_ecff1cm.TransactionDetails, int(l.len))
			tmpList := NewHistoryManagerFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__loadHistoryWithFilters_atList(i)
			}
			return out
		}(C.HistoryManager554044_LoadHistoryWithFilters(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewHistoryManagerFromPointer(NewHistoryManagerFromPointer(nil).__loadHistoryWithFilters_addresses_newList())
			for k, v := range addresses {
				tmpList.__loadHistoryWithFilters_addresses_setList(k, std_core.NewQVariant1(v))
			}
			return tmpList.Pointer()
		}(), C.struct_Moc_PackedString{data: filterAddressesC, len: C.longlong(len(strings.Join(filterAddresses, "¡¦!")))}))
	}
	return make([]*custom_transactions_ecff1cm.TransactionDetails, 0)
}

//export callbackHistoryManager554044_LoadHistory
func callbackHistoryManager554044_LoadHistory(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "loadHistory"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewHistoryManagerFromPointer(NewHistoryManagerFromPointer(nil).__loadHistory_newList())
			for _, v := range (*(*func() []*custom_transactions_ecff1cm.TransactionDetails)(signal))() {
				tmpList.__loadHistory_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewHistoryManagerFromPointer(NewHistoryManagerFromPointer(nil).__loadHistory_newList())
		for _, v := range make([]*custom_transactions_ecff1cm.TransactionDetails, 0) {
			tmpList.__loadHistory_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *HistoryManager) ConnectLoadHistory(f func() []*custom_transactions_ecff1cm.TransactionDetails) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "loadHistory"); signal != nil {
			f := func() []*custom_transactions_ecff1cm.TransactionDetails {
				(*(*func() []*custom_transactions_ecff1cm.TransactionDetails)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "loadHistory", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "loadHistory", unsafe.Pointer(&f))
		}
	}
}

func (ptr *HistoryManager) DisconnectLoadHistory() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "loadHistory")
	}
}

func (ptr *HistoryManager) LoadHistory() []*custom_transactions_ecff1cm.TransactionDetails {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*custom_transactions_ecff1cm.TransactionDetails {
			out := make([]*custom_transactions_ecff1cm.TransactionDetails, int(l.len))
			tmpList := NewHistoryManagerFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__loadHistory_atList(i)
			}
			return out
		}(C.HistoryManager554044_LoadHistory(ptr.Pointer()))
	}
	return make([]*custom_transactions_ecff1cm.TransactionDetails, 0)
}

func HistoryManager_QRegisterMetaType() int {
	return int(int32(C.HistoryManager554044_HistoryManager554044_QRegisterMetaType()))
}

func (ptr *HistoryManager) QRegisterMetaType() int {
	return int(int32(C.HistoryManager554044_HistoryManager554044_QRegisterMetaType()))
}

func HistoryManager_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.HistoryManager554044_HistoryManager554044_QRegisterMetaType2(typeNameC)))
}

func (ptr *HistoryManager) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.HistoryManager554044_HistoryManager554044_QRegisterMetaType2(typeNameC)))
}

func HistoryManager_QmlRegisterType() int {
	return int(int32(C.HistoryManager554044_HistoryManager554044_QmlRegisterType()))
}

func (ptr *HistoryManager) QmlRegisterType() int {
	return int(int32(C.HistoryManager554044_HistoryManager554044_QmlRegisterType()))
}

func HistoryManager_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.HistoryManager554044_HistoryManager554044_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *HistoryManager) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.HistoryManager554044_HistoryManager554044_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *HistoryManager) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.HistoryManager554044___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryManager) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryManager554044___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *HistoryManager) __children_newList() unsafe.Pointer {
	return C.HistoryManager554044___children_newList(ptr.Pointer())
}

func (ptr *HistoryManager) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.HistoryManager554044___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryManager) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryManager554044___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *HistoryManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.HistoryManager554044___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *HistoryManager) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.HistoryManager554044___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryManager) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryManager554044___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *HistoryManager) __findChildren_newList() unsafe.Pointer {
	return C.HistoryManager554044___findChildren_newList(ptr.Pointer())
}

func (ptr *HistoryManager) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.HistoryManager554044___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryManager) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryManager554044___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *HistoryManager) __findChildren_newList3() unsafe.Pointer {
	return C.HistoryManager554044___findChildren_newList3(ptr.Pointer())
}

func (ptr *HistoryManager) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.HistoryManager554044___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryManager) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryManager554044___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *HistoryManager) __qFindChildren_newList2() unsafe.Pointer {
	return C.HistoryManager554044___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *HistoryManager) __loadHistoryWithFilters_atList(i int) *custom_transactions_ecff1cm.TransactionDetails {
	if ptr.Pointer() != nil {
		tmpValue := custom_transactions_ecff1cm.NewTransactionDetailsFromPointer(C.HistoryManager554044___loadHistoryWithFilters_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryManager) __loadHistoryWithFilters_setList(i custom_transactions_ecff1cm.TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryManager554044___loadHistoryWithFilters_setList(ptr.Pointer(), custom_transactions_ecff1cm.PointerFromTransactionDetails(i))
	}
}

func (ptr *HistoryManager) __loadHistoryWithFilters_newList() unsafe.Pointer {
	return C.HistoryManager554044___loadHistoryWithFilters_newList(ptr.Pointer())
}

func (ptr *HistoryManager) __loadHistoryWithFilters_addresses_atList(v string, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := std_core.NewQVariantFromPointer(C.HistoryManager554044___loadHistoryWithFilters_addresses_atList(ptr.Pointer(), C.struct_Moc_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryManager) __loadHistoryWithFilters_addresses_setList(key string, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.HistoryManager554044___loadHistoryWithFilters_addresses_setList(ptr.Pointer(), C.struct_Moc_PackedString{data: keyC, len: C.longlong(len(key))}, std_core.PointerFromQVariant(i))
	}
}

func (ptr *HistoryManager) __loadHistoryWithFilters_addresses_newList() unsafe.Pointer {
	return C.HistoryManager554044___loadHistoryWithFilters_addresses_newList(ptr.Pointer())
}

func (ptr *HistoryManager) __loadHistoryWithFilters_addresses_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewHistoryManagerFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____loadHistoryWithFilters_addresses_keyList_atList(i)
			}
			return out
		}(C.HistoryManager554044___loadHistoryWithFilters_addresses_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *HistoryManager) __loadHistory_atList(i int) *custom_transactions_ecff1cm.TransactionDetails {
	if ptr.Pointer() != nil {
		tmpValue := custom_transactions_ecff1cm.NewTransactionDetailsFromPointer(C.HistoryManager554044___loadHistory_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryManager) __loadHistory_setList(i custom_transactions_ecff1cm.TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryManager554044___loadHistory_setList(ptr.Pointer(), custom_transactions_ecff1cm.PointerFromTransactionDetails(i))
	}
}

func (ptr *HistoryManager) __loadHistory_newList() unsafe.Pointer {
	return C.HistoryManager554044___loadHistory_newList(ptr.Pointer())
}

func (ptr *HistoryManager) ____loadHistoryWithFilters_addresses_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.HistoryManager554044_____loadHistoryWithFilters_addresses_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *HistoryManager) ____loadHistoryWithFilters_addresses_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.HistoryManager554044_____loadHistoryWithFilters_addresses_keyList_setList(ptr.Pointer(), C.struct_Moc_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *HistoryManager) ____loadHistoryWithFilters_addresses_keyList_newList() unsafe.Pointer {
	return C.HistoryManager554044_____loadHistoryWithFilters_addresses_keyList_newList(ptr.Pointer())
}

func NewHistoryManager(parent std_core.QObject_ITF) *HistoryManager {
	tmpValue := NewHistoryManagerFromPointer(C.HistoryManager554044_NewHistoryManager(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackHistoryManager554044_DestroyHistoryManager
func callbackHistoryManager554044_DestroyHistoryManager(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~HistoryManager"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHistoryManagerFromPointer(ptr).DestroyHistoryManagerDefault()
	}
}

func (ptr *HistoryManager) ConnectDestroyHistoryManager(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~HistoryManager"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~HistoryManager", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~HistoryManager", unsafe.Pointer(&f))
		}
	}
}

func (ptr *HistoryManager) DisconnectDestroyHistoryManager() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~HistoryManager")
	}
}

func (ptr *HistoryManager) DestroyHistoryManager() {
	if ptr.Pointer() != nil {
		C.HistoryManager554044_DestroyHistoryManager(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *HistoryManager) DestroyHistoryManagerDefault() {
	if ptr.Pointer() != nil {
		C.HistoryManager554044_DestroyHistoryManagerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackHistoryManager554044_ChildEvent
func callbackHistoryManager554044_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewHistoryManagerFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *HistoryManager) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryManager554044_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackHistoryManager554044_ConnectNotify
func callbackHistoryManager554044_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewHistoryManagerFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *HistoryManager) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryManager554044_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackHistoryManager554044_CustomEvent
func callbackHistoryManager554044_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewHistoryManagerFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *HistoryManager) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryManager554044_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackHistoryManager554044_DeleteLater
func callbackHistoryManager554044_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHistoryManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *HistoryManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.HistoryManager554044_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackHistoryManager554044_Destroyed
func callbackHistoryManager554044_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackHistoryManager554044_DisconnectNotify
func callbackHistoryManager554044_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewHistoryManagerFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *HistoryManager) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryManager554044_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackHistoryManager554044_Event
func callbackHistoryManager554044_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryManagerFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *HistoryManager) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryManager554044_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackHistoryManager554044_EventFilter
func callbackHistoryManager554044_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryManagerFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *HistoryManager) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryManager554044_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackHistoryManager554044_ObjectNameChanged
func callbackHistoryManager554044_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackHistoryManager554044_TimerEvent
func callbackHistoryManager554044_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewHistoryManagerFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *HistoryManager) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryManager554044_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type TransactionList_ITF interface {
	std_core.QAbstractListModel_ITF
	TransactionList_PTR() *TransactionList
}

func (ptr *TransactionList) TransactionList_PTR() *TransactionList {
	return ptr
}

func (ptr *TransactionList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *TransactionList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromTransactionList(ptr TransactionList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.TransactionList_PTR().Pointer()
	}
	return nil
}

func NewTransactionListFromPointer(ptr unsafe.Pointer) (n *TransactionList) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(TransactionList)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *TransactionList:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &TransactionList{QAbstractListModel: *deduced}

		default:
			n = new(TransactionList)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *TransactionList) Init() { this.init() }

//export callbackTransactionList554044_Constructor
func callbackTransactionList554044_Constructor(ptr unsafe.Pointer) {
	this := NewTransactionListFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectAddTransaction(this.addTransaction)
	this.ConnectRemoveTransaction(this.removeTransaction)
	this.init()
}

//export callbackTransactionList554044_AddTransaction
func callbackTransactionList554044_AddTransaction(ptr unsafe.Pointer, transaction unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addTransaction"); signal != nil {
		(*(*func(*custom_transactions_ecff1cm.TransactionDetails))(signal))(custom_transactions_ecff1cm.NewTransactionDetailsFromPointer(transaction))
	}

}

func (ptr *TransactionList) ConnectAddTransaction(f func(transaction *custom_transactions_ecff1cm.TransactionDetails)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addTransaction") {
			C.TransactionList554044_ConnectAddTransaction(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addTransaction"); signal != nil {
			f := func(transaction *custom_transactions_ecff1cm.TransactionDetails) {
				(*(*func(*custom_transactions_ecff1cm.TransactionDetails))(signal))(transaction)
				f(transaction)
			}
			qt.ConnectSignal(ptr.Pointer(), "addTransaction", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addTransaction", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionList) DisconnectAddTransaction() {
	if ptr.Pointer() != nil {
		C.TransactionList554044_DisconnectAddTransaction(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addTransaction")
	}
}

func (ptr *TransactionList) AddTransaction(transaction custom_transactions_ecff1cm.TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_AddTransaction(ptr.Pointer(), custom_transactions_ecff1cm.PointerFromTransactionDetails(transaction))
	}
}

//export callbackTransactionList554044_RemoveTransaction
func callbackTransactionList554044_RemoveTransaction(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "removeTransaction"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *TransactionList) ConnectRemoveTransaction(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "removeTransaction") {
			C.TransactionList554044_ConnectRemoveTransaction(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "removeTransaction"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "removeTransaction", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeTransaction", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionList) DisconnectRemoveTransaction() {
	if ptr.Pointer() != nil {
		C.TransactionList554044_DisconnectRemoveTransaction(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "removeTransaction")
	}
}

func (ptr *TransactionList) RemoveTransaction(index int) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_RemoveTransaction(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackTransactionList554044_AddMultipleTransactions
func callbackTransactionList554044_AddMultipleTransactions(ptr unsafe.Pointer, txns C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "addMultipleTransactions"); signal != nil {
		(*(*func([]*custom_transactions_ecff1cm.TransactionDetails))(signal))(func(l C.struct_Moc_PackedList) []*custom_transactions_ecff1cm.TransactionDetails {
			out := make([]*custom_transactions_ecff1cm.TransactionDetails, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__addMultipleTransactions_txns_atList(i)
			}
			return out
		}(txns))
	}

}

func (ptr *TransactionList) ConnectAddMultipleTransactions(f func(txns []*custom_transactions_ecff1cm.TransactionDetails)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addMultipleTransactions"); signal != nil {
			f := func(txns []*custom_transactions_ecff1cm.TransactionDetails) {
				(*(*func([]*custom_transactions_ecff1cm.TransactionDetails))(signal))(txns)
				f(txns)
			}
			qt.ConnectSignal(ptr.Pointer(), "addMultipleTransactions", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addMultipleTransactions", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionList) DisconnectAddMultipleTransactions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addMultipleTransactions")
	}
}

func (ptr *TransactionList) AddMultipleTransactions(txns []*custom_transactions_ecff1cm.TransactionDetails) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_AddMultipleTransactions(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__addMultipleTransactions_txns_newList())
			for _, v := range txns {
				tmpList.__addMultipleTransactions_txns_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackTransactionList554044_Roles
func callbackTransactionList554044_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__roles_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__roles_newList())
		for k, v := range NewTransactionListFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TransactionList) ConnectRoles(f func() map[int]*std_core.QByteArray) {
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

func (ptr *TransactionList) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *TransactionList) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.TransactionList554044_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *TransactionList) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.TransactionList554044_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackTransactionList554044_SetRoles
func callbackTransactionList554044_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	} else {
		NewTransactionListFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	}
}

func (ptr *TransactionList) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
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

func (ptr *TransactionList) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *TransactionList) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *TransactionList) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackTransactionList554044_RolesChanged
func callbackTransactionList554044_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i, v := range tmpList.__rolesChanged_roles_keyList() {
				out[v] = tmpList.__rolesChanged_roles_atList(v, i)
			}
			return out
		}(roles))
	}

}

func (ptr *TransactionList) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.TransactionList554044_ConnectRolesChanged(ptr.Pointer())
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

func (ptr *TransactionList) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.TransactionList554044_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *TransactionList) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackTransactionList554044_Transactions
func callbackTransactionList554044_Transactions(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "transactions"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__transactions_newList())
			for _, v := range (*(*func() []*custom_transactions_ecff1cm.TransactionDetails)(signal))() {
				tmpList.__transactions_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__transactions_newList())
		for _, v := range NewTransactionListFromPointer(ptr).TransactionsDefault() {
			tmpList.__transactions_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TransactionList) ConnectTransactions(f func() []*custom_transactions_ecff1cm.TransactionDetails) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "transactions"); signal != nil {
			f := func() []*custom_transactions_ecff1cm.TransactionDetails {
				(*(*func() []*custom_transactions_ecff1cm.TransactionDetails)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "transactions", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transactions", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionList) DisconnectTransactions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "transactions")
	}
}

func (ptr *TransactionList) Transactions() []*custom_transactions_ecff1cm.TransactionDetails {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*custom_transactions_ecff1cm.TransactionDetails {
			out := make([]*custom_transactions_ecff1cm.TransactionDetails, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__transactions_atList(i)
			}
			return out
		}(C.TransactionList554044_Transactions(ptr.Pointer()))
	}
	return make([]*custom_transactions_ecff1cm.TransactionDetails, 0)
}

func (ptr *TransactionList) TransactionsDefault() []*custom_transactions_ecff1cm.TransactionDetails {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*custom_transactions_ecff1cm.TransactionDetails {
			out := make([]*custom_transactions_ecff1cm.TransactionDetails, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__transactions_atList(i)
			}
			return out
		}(C.TransactionList554044_TransactionsDefault(ptr.Pointer()))
	}
	return make([]*custom_transactions_ecff1cm.TransactionDetails, 0)
}

//export callbackTransactionList554044_SetTransactions
func callbackTransactionList554044_SetTransactions(ptr unsafe.Pointer, transactions C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setTransactions"); signal != nil {
		(*(*func([]*custom_transactions_ecff1cm.TransactionDetails))(signal))(func(l C.struct_Moc_PackedList) []*custom_transactions_ecff1cm.TransactionDetails {
			out := make([]*custom_transactions_ecff1cm.TransactionDetails, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setTransactions_transactions_atList(i)
			}
			return out
		}(transactions))
	} else {
		NewTransactionListFromPointer(ptr).SetTransactionsDefault(func(l C.struct_Moc_PackedList) []*custom_transactions_ecff1cm.TransactionDetails {
			out := make([]*custom_transactions_ecff1cm.TransactionDetails, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setTransactions_transactions_atList(i)
			}
			return out
		}(transactions))
	}
}

func (ptr *TransactionList) ConnectSetTransactions(f func(transactions []*custom_transactions_ecff1cm.TransactionDetails)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTransactions"); signal != nil {
			f := func(transactions []*custom_transactions_ecff1cm.TransactionDetails) {
				(*(*func([]*custom_transactions_ecff1cm.TransactionDetails))(signal))(transactions)
				f(transactions)
			}
			qt.ConnectSignal(ptr.Pointer(), "setTransactions", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTransactions", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionList) DisconnectSetTransactions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTransactions")
	}
}

func (ptr *TransactionList) SetTransactions(transactions []*custom_transactions_ecff1cm.TransactionDetails) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_SetTransactions(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__setTransactions_transactions_newList())
			for _, v := range transactions {
				tmpList.__setTransactions_transactions_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *TransactionList) SetTransactionsDefault(transactions []*custom_transactions_ecff1cm.TransactionDetails) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_SetTransactionsDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__setTransactions_transactions_newList())
			for _, v := range transactions {
				tmpList.__setTransactions_transactions_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackTransactionList554044_TransactionsChanged
func callbackTransactionList554044_TransactionsChanged(ptr unsafe.Pointer, transactions C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "transactionsChanged"); signal != nil {
		(*(*func([]*custom_transactions_ecff1cm.TransactionDetails))(signal))(func(l C.struct_Moc_PackedList) []*custom_transactions_ecff1cm.TransactionDetails {
			out := make([]*custom_transactions_ecff1cm.TransactionDetails, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__transactionsChanged_transactions_atList(i)
			}
			return out
		}(transactions))
	}

}

func (ptr *TransactionList) ConnectTransactionsChanged(f func(transactions []*custom_transactions_ecff1cm.TransactionDetails)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transactionsChanged") {
			C.TransactionList554044_ConnectTransactionsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "transactionsChanged"); signal != nil {
			f := func(transactions []*custom_transactions_ecff1cm.TransactionDetails) {
				(*(*func([]*custom_transactions_ecff1cm.TransactionDetails))(signal))(transactions)
				f(transactions)
			}
			qt.ConnectSignal(ptr.Pointer(), "transactionsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transactionsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionList) DisconnectTransactionsChanged() {
	if ptr.Pointer() != nil {
		C.TransactionList554044_DisconnectTransactionsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transactionsChanged")
	}
}

func (ptr *TransactionList) TransactionsChanged(transactions []*custom_transactions_ecff1cm.TransactionDetails) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_TransactionsChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__transactionsChanged_transactions_newList())
			for _, v := range transactions {
				tmpList.__transactionsChanged_transactions_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func TransactionList_QRegisterMetaType() int {
	return int(int32(C.TransactionList554044_TransactionList554044_QRegisterMetaType()))
}

func (ptr *TransactionList) QRegisterMetaType() int {
	return int(int32(C.TransactionList554044_TransactionList554044_QRegisterMetaType()))
}

func TransactionList_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TransactionList554044_TransactionList554044_QRegisterMetaType2(typeNameC)))
}

func (ptr *TransactionList) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TransactionList554044_TransactionList554044_QRegisterMetaType2(typeNameC)))
}

func TransactionList_QmlRegisterType() int {
	return int(int32(C.TransactionList554044_TransactionList554044_QmlRegisterType()))
}

func (ptr *TransactionList) QmlRegisterType() int {
	return int(int32(C.TransactionList554044_TransactionList554044_QmlRegisterType()))
}

func TransactionList_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.TransactionList554044_TransactionList554044_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TransactionList) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.TransactionList554044_TransactionList554044_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TransactionList) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList554044_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____itemData_keyList_newList() unsafe.Pointer {
	return C.TransactionList554044_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList554044_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.TransactionList554044_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList554044_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.TransactionList554044_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList554044___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TransactionList) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.TransactionList554044___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *TransactionList) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList554044___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TransactionList) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.TransactionList554044___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *TransactionList) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList554044___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) __dataChanged_roles_newList() unsafe.Pointer {
	return C.TransactionList554044___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *TransactionList) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TransactionList554044___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *TransactionList) __itemData_newList() unsafe.Pointer {
	return C.TransactionList554044___itemData_newList(ptr.Pointer())
}

func (ptr *TransactionList) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.TransactionList554044___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TransactionList) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.TransactionList554044___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *TransactionList) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.TransactionList554044___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *TransactionList) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.TransactionList554044___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *TransactionList) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.TransactionList554044___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *TransactionList) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList554044___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TransactionList) __match_newList() unsafe.Pointer {
	return C.TransactionList554044___match_newList(ptr.Pointer())
}

func (ptr *TransactionList) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList554044___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TransactionList) __mimeData_indexes_newList() unsafe.Pointer {
	return C.TransactionList554044___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *TransactionList) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList554044___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TransactionList) __persistentIndexList_newList() unsafe.Pointer {
	return C.TransactionList554044___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *TransactionList) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionList554044___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionList) __roleNames_newList() unsafe.Pointer {
	return C.TransactionList554044___roleNames_newList(ptr.Pointer())
}

func (ptr *TransactionList) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.TransactionList554044___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TransactionList) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TransactionList554044___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *TransactionList) __setItemData_roles_newList() unsafe.Pointer {
	return C.TransactionList554044___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *TransactionList) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.TransactionList554044___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TransactionList) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList554044_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.TransactionList554044_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList554044_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.TransactionList554044_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionList554044___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionList) __children_newList() unsafe.Pointer {
	return C.TransactionList554044___children_newList(ptr.Pointer())
}

func (ptr *TransactionList) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionList554044___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionList) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.TransactionList554044___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *TransactionList) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionList554044___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionList) __findChildren_newList() unsafe.Pointer {
	return C.TransactionList554044___findChildren_newList(ptr.Pointer())
}

func (ptr *TransactionList) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionList554044___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionList) __findChildren_newList3() unsafe.Pointer {
	return C.TransactionList554044___findChildren_newList3(ptr.Pointer())
}

func (ptr *TransactionList) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionList554044___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionList) __qFindChildren_newList2() unsafe.Pointer {
	return C.TransactionList554044___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *TransactionList) __addMultipleTransactions_txns_atList(i int) *custom_transactions_ecff1cm.TransactionDetails {
	if ptr.Pointer() != nil {
		tmpValue := custom_transactions_ecff1cm.NewTransactionDetailsFromPointer(C.TransactionList554044___addMultipleTransactions_txns_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __addMultipleTransactions_txns_setList(i custom_transactions_ecff1cm.TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___addMultipleTransactions_txns_setList(ptr.Pointer(), custom_transactions_ecff1cm.PointerFromTransactionDetails(i))
	}
}

func (ptr *TransactionList) __addMultipleTransactions_txns_newList() unsafe.Pointer {
	return C.TransactionList554044___addMultipleTransactions_txns_newList(ptr.Pointer())
}

func (ptr *TransactionList) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionList554044___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionList) __roles_newList() unsafe.Pointer {
	return C.TransactionList554044___roles_newList(ptr.Pointer())
}

func (ptr *TransactionList) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roles_keyList_atList(i)
			}
			return out
		}(C.TransactionList554044___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TransactionList) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionList554044___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionList) __setRoles_roles_newList() unsafe.Pointer {
	return C.TransactionList554044___setRoles_roles_newList(ptr.Pointer())
}

func (ptr *TransactionList) __setRoles_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoles_roles_keyList_atList(i)
			}
			return out
		}(C.TransactionList554044___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TransactionList) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionList554044___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionList) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.TransactionList554044___rolesChanged_roles_newList(ptr.Pointer())
}

func (ptr *TransactionList) __rolesChanged_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____rolesChanged_roles_keyList_atList(i)
			}
			return out
		}(C.TransactionList554044___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TransactionList) __transactions_atList(i int) *custom_transactions_ecff1cm.TransactionDetails {
	if ptr.Pointer() != nil {
		tmpValue := custom_transactions_ecff1cm.NewTransactionDetailsFromPointer(C.TransactionList554044___transactions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __transactions_setList(i custom_transactions_ecff1cm.TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___transactions_setList(ptr.Pointer(), custom_transactions_ecff1cm.PointerFromTransactionDetails(i))
	}
}

func (ptr *TransactionList) __transactions_newList() unsafe.Pointer {
	return C.TransactionList554044___transactions_newList(ptr.Pointer())
}

func (ptr *TransactionList) __setTransactions_transactions_atList(i int) *custom_transactions_ecff1cm.TransactionDetails {
	if ptr.Pointer() != nil {
		tmpValue := custom_transactions_ecff1cm.NewTransactionDetailsFromPointer(C.TransactionList554044___setTransactions_transactions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __setTransactions_transactions_setList(i custom_transactions_ecff1cm.TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___setTransactions_transactions_setList(ptr.Pointer(), custom_transactions_ecff1cm.PointerFromTransactionDetails(i))
	}
}

func (ptr *TransactionList) __setTransactions_transactions_newList() unsafe.Pointer {
	return C.TransactionList554044___setTransactions_transactions_newList(ptr.Pointer())
}

func (ptr *TransactionList) __transactionsChanged_transactions_atList(i int) *custom_transactions_ecff1cm.TransactionDetails {
	if ptr.Pointer() != nil {
		tmpValue := custom_transactions_ecff1cm.NewTransactionDetailsFromPointer(C.TransactionList554044___transactionsChanged_transactions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __transactionsChanged_transactions_setList(i custom_transactions_ecff1cm.TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044___transactionsChanged_transactions_setList(ptr.Pointer(), custom_transactions_ecff1cm.PointerFromTransactionDetails(i))
	}
}

func (ptr *TransactionList) __transactionsChanged_transactions_newList() unsafe.Pointer {
	return C.TransactionList554044___transactionsChanged_transactions_newList(ptr.Pointer())
}

func (ptr *TransactionList) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList554044_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____roles_keyList_newList() unsafe.Pointer {
	return C.TransactionList554044_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList554044_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.TransactionList554044_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList554044_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.TransactionList554044_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewTransactionList(parent std_core.QObject_ITF) *TransactionList {
	tmpValue := NewTransactionListFromPointer(C.TransactionList554044_NewTransactionList(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTransactionList554044_DestroyTransactionList
func callbackTransactionList554044_DestroyTransactionList(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~TransactionList"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionListFromPointer(ptr).DestroyTransactionListDefault()
	}
}

func (ptr *TransactionList) ConnectDestroyTransactionList(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~TransactionList"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~TransactionList", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~TransactionList", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionList) DisconnectDestroyTransactionList() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~TransactionList")
	}
}

func (ptr *TransactionList) DestroyTransactionList() {
	if ptr.Pointer() != nil {
		C.TransactionList554044_DestroyTransactionList(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *TransactionList) DestroyTransactionListDefault() {
	if ptr.Pointer() != nil {
		C.TransactionList554044_DestroyTransactionListDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTransactionList554044_DropMimeData
func callbackTransactionList554044_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList554044_Flags
func callbackTransactionList554044_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewTransactionListFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TransactionList) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.TransactionList554044_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackTransactionList554044_Index
func callbackTransactionList554044_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewTransactionListFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *TransactionList) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList554044_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList554044_Sibling
func callbackTransactionList554044_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewTransactionListFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *TransactionList) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList554044_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList554044_Buddy
func callbackTransactionList554044_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewTransactionListFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TransactionList) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList554044_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList554044_CanDropMimeData
func callbackTransactionList554044_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList554044_CanFetchMore
func callbackTransactionList554044_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList554044_ColumnCount
func callbackTransactionList554044_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewTransactionListFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *TransactionList) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList554044_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackTransactionList554044_ColumnsAboutToBeInserted
func callbackTransactionList554044_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList554044_ColumnsAboutToBeMoved
func callbackTransactionList554044_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackTransactionList554044_ColumnsAboutToBeRemoved
func callbackTransactionList554044_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList554044_ColumnsInserted
func callbackTransactionList554044_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList554044_ColumnsMoved
func callbackTransactionList554044_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackTransactionList554044_ColumnsRemoved
func callbackTransactionList554044_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList554044_Data
func callbackTransactionList554044_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewTransactionListFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *TransactionList) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TransactionList554044_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList554044_DataChanged
func callbackTransactionList554044_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(signal))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackTransactionList554044_FetchMore
func callbackTransactionList554044_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewTransactionListFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *TransactionList) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackTransactionList554044_HasChildren
func callbackTransactionList554044_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList554044_HeaderData
func callbackTransactionList554044_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewTransactionListFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *TransactionList) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TransactionList554044_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList554044_HeaderDataChanged
func callbackTransactionList554044_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList554044_InsertColumns
func callbackTransactionList554044_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList554044_InsertRows
func callbackTransactionList554044_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList554044_ItemData
func callbackTransactionList554044_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*std_core.QModelIndex) map[int]*std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__itemData_newList())
		for k, v := range NewTransactionListFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TransactionList) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.TransactionList554044_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackTransactionList554044_LayoutAboutToBeChanged
func callbackTransactionList554044_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackTransactionList554044_LayoutChanged
func callbackTransactionList554044_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackTransactionList554044_Match
func callbackTransactionList554044_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__match_newList())
		for _, v := range NewTransactionListFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TransactionList) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.TransactionList554044_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackTransactionList554044_MimeData
func callbackTransactionList554044_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func([]*std_core.QModelIndex) *std_core.QMimeData)(signal))(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewTransactionListFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewTransactionListFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *TransactionList) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.TransactionList554044_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__mimeData_indexes_newList())
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

//export callbackTransactionList554044_MimeTypes
func callbackTransactionList554044_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewTransactionListFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *TransactionList) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.TransactionList554044_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackTransactionList554044_ModelAboutToBeReset
func callbackTransactionList554044_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackTransactionList554044_ModelReset
func callbackTransactionList554044_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackTransactionList554044_MoveColumns
func callbackTransactionList554044_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *TransactionList) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackTransactionList554044_MoveRows
func callbackTransactionList554044_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *TransactionList) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackTransactionList554044_Parent
func callbackTransactionList554044_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewTransactionListFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TransactionList) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList554044_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList554044_RemoveColumns
func callbackTransactionList554044_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList554044_RemoveRows
func callbackTransactionList554044_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList554044_ResetInternalData
func callbackTransactionList554044_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionListFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *TransactionList) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.TransactionList554044_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackTransactionList554044_Revert
func callbackTransactionList554044_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionListFromPointer(ptr).RevertDefault()
	}
}

func (ptr *TransactionList) RevertDefault() {
	if ptr.Pointer() != nil {
		C.TransactionList554044_RevertDefault(ptr.Pointer())
	}
}

//export callbackTransactionList554044_RoleNames
func callbackTransactionList554044_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__roleNames_newList())
		for k, v := range NewTransactionListFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *TransactionList) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.TransactionList554044_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackTransactionList554044_RowCount
func callbackTransactionList554044_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewTransactionListFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *TransactionList) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList554044_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackTransactionList554044_RowsAboutToBeInserted
func callbackTransactionList554044_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackTransactionList554044_RowsAboutToBeMoved
func callbackTransactionList554044_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackTransactionList554044_RowsAboutToBeRemoved
func callbackTransactionList554044_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList554044_RowsInserted
func callbackTransactionList554044_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList554044_RowsMoved
func callbackTransactionList554044_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackTransactionList554044_RowsRemoved
func callbackTransactionList554044_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList554044_SetData
func callbackTransactionList554044_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *TransactionList) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackTransactionList554044_SetHeaderData
func callbackTransactionList554044_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *TransactionList) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackTransactionList554044_SetItemData
func callbackTransactionList554044_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(signal))(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewTransactionListFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *TransactionList) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackTransactionList554044_Sort
func callbackTransactionList554044_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewTransactionListFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *TransactionList) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackTransactionList554044_Span
func callbackTransactionList554044_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewTransactionListFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TransactionList) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.TransactionList554044_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList554044_Submit
func callbackTransactionList554044_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).SubmitDefault())))
}

func (ptr *TransactionList) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackTransactionList554044_SupportedDragActions
func callbackTransactionList554044_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewTransactionListFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *TransactionList) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.TransactionList554044_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackTransactionList554044_SupportedDropActions
func callbackTransactionList554044_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewTransactionListFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *TransactionList) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.TransactionList554044_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackTransactionList554044_ChildEvent
func callbackTransactionList554044_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTransactionListFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *TransactionList) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTransactionList554044_ConnectNotify
func callbackTransactionList554044_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTransactionListFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TransactionList) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTransactionList554044_CustomEvent
func callbackTransactionList554044_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewTransactionListFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *TransactionList) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTransactionList554044_DeleteLater
func callbackTransactionList554044_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionListFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *TransactionList) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.TransactionList554044_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTransactionList554044_Destroyed
func callbackTransactionList554044_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackTransactionList554044_DisconnectNotify
func callbackTransactionList554044_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTransactionListFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TransactionList) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTransactionList554044_Event
func callbackTransactionList554044_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *TransactionList) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackTransactionList554044_EventFilter
func callbackTransactionList554044_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *TransactionList) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList554044_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackTransactionList554044_ObjectNameChanged
func callbackTransactionList554044_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackTransactionList554044_TimerEvent
func callbackTransactionList554044_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTransactionListFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *TransactionList) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList554044_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

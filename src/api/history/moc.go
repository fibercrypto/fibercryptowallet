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

type HistoryModel_ITF interface {
	std_core.QAbstractListModel_ITF
	HistoryModel_PTR() *HistoryModel
}

func (ptr *HistoryModel) HistoryModel_PTR() *HistoryModel {
	return ptr
}

func (ptr *HistoryModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *HistoryModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromHistoryModel(ptr HistoryModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.HistoryModel_PTR().Pointer()
	}
	return nil
}

func NewHistoryModelFromPointer(ptr unsafe.Pointer) (n *HistoryModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(HistoryModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *HistoryModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &HistoryModel{QAbstractListModel: *deduced}

		default:
			n = new(HistoryModel)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *HistoryModel) Init() { this.init() }

//export callbackHistoryModel8ba275_Constructor
func callbackHistoryModel8ba275_Constructor(ptr unsafe.Pointer) {
	this := NewHistoryModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackHistoryModel8ba275_AddTransaction
func callbackHistoryModel8ba275_AddTransaction(ptr unsafe.Pointer, transaction unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addTransaction"); signal != nil {
		(*(*func(*TransactionDetails))(signal))(NewTransactionDetailsFromPointer(transaction))
	}

}

func (ptr *HistoryModel) ConnectAddTransaction(f func(transaction *TransactionDetails)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addTransaction") {
			C.HistoryModel8ba275_ConnectAddTransaction(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addTransaction"); signal != nil {
			f := func(transaction *TransactionDetails) {
				(*(*func(*TransactionDetails))(signal))(transaction)
				f(transaction)
			}
			qt.ConnectSignal(ptr.Pointer(), "addTransaction", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addTransaction", unsafe.Pointer(&f))
		}
	}
}

func (ptr *HistoryModel) DisconnectAddTransaction() {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_DisconnectAddTransaction(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addTransaction")
	}
}

func (ptr *HistoryModel) AddTransaction(transaction TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_AddTransaction(ptr.Pointer(), PointerFromTransactionDetails(transaction))
	}
}

//export callbackHistoryModel8ba275_Roles
func callbackHistoryModel8ba275_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__roles_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__roles_newList())
		for k, v := range NewHistoryModelFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *HistoryModel) ConnectRoles(f func() map[int]*std_core.QByteArray) {
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

func (ptr *HistoryModel) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *HistoryModel) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.HistoryModel8ba275_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *HistoryModel) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.HistoryModel8ba275_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackHistoryModel8ba275_SetRoles
func callbackHistoryModel8ba275_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	} else {
		NewHistoryModelFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	}
}

func (ptr *HistoryModel) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
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

func (ptr *HistoryModel) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *HistoryModel) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *HistoryModel) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackHistoryModel8ba275_RolesChanged
func callbackHistoryModel8ba275_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i, v := range tmpList.__rolesChanged_roles_keyList() {
				out[v] = tmpList.__rolesChanged_roles_atList(v, i)
			}
			return out
		}(roles))
	}

}

func (ptr *HistoryModel) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.HistoryModel8ba275_ConnectRolesChanged(ptr.Pointer())
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

func (ptr *HistoryModel) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *HistoryModel) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackHistoryModel8ba275_Transactions
func callbackHistoryModel8ba275_Transactions(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "transactions"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__transactions_newList())
			for _, v := range (*(*func() []*TransactionDetails)(signal))() {
				tmpList.__transactions_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__transactions_newList())
		for _, v := range NewHistoryModelFromPointer(ptr).TransactionsDefault() {
			tmpList.__transactions_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *HistoryModel) ConnectTransactions(f func() []*TransactionDetails) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "transactions"); signal != nil {
			f := func() []*TransactionDetails {
				(*(*func() []*TransactionDetails)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "transactions", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transactions", unsafe.Pointer(&f))
		}
	}
}

func (ptr *HistoryModel) DisconnectTransactions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "transactions")
	}
}

func (ptr *HistoryModel) Transactions() []*TransactionDetails {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*TransactionDetails {
			out := make([]*TransactionDetails, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__transactions_atList(i)
			}
			return out
		}(C.HistoryModel8ba275_Transactions(ptr.Pointer()))
	}
	return make([]*TransactionDetails, 0)
}

func (ptr *HistoryModel) TransactionsDefault() []*TransactionDetails {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*TransactionDetails {
			out := make([]*TransactionDetails, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__transactions_atList(i)
			}
			return out
		}(C.HistoryModel8ba275_TransactionsDefault(ptr.Pointer()))
	}
	return make([]*TransactionDetails, 0)
}

//export callbackHistoryModel8ba275_SetTransactions
func callbackHistoryModel8ba275_SetTransactions(ptr unsafe.Pointer, transactions C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setTransactions"); signal != nil {
		(*(*func([]*TransactionDetails))(signal))(func(l C.struct_Moc_PackedList) []*TransactionDetails {
			out := make([]*TransactionDetails, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setTransactions_transactions_atList(i)
			}
			return out
		}(transactions))
	} else {
		NewHistoryModelFromPointer(ptr).SetTransactionsDefault(func(l C.struct_Moc_PackedList) []*TransactionDetails {
			out := make([]*TransactionDetails, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setTransactions_transactions_atList(i)
			}
			return out
		}(transactions))
	}
}

func (ptr *HistoryModel) ConnectSetTransactions(f func(transactions []*TransactionDetails)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTransactions"); signal != nil {
			f := func(transactions []*TransactionDetails) {
				(*(*func([]*TransactionDetails))(signal))(transactions)
				f(transactions)
			}
			qt.ConnectSignal(ptr.Pointer(), "setTransactions", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTransactions", unsafe.Pointer(&f))
		}
	}
}

func (ptr *HistoryModel) DisconnectSetTransactions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTransactions")
	}
}

func (ptr *HistoryModel) SetTransactions(transactions []*TransactionDetails) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_SetTransactions(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__setTransactions_transactions_newList())
			for _, v := range transactions {
				tmpList.__setTransactions_transactions_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *HistoryModel) SetTransactionsDefault(transactions []*TransactionDetails) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_SetTransactionsDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__setTransactions_transactions_newList())
			for _, v := range transactions {
				tmpList.__setTransactions_transactions_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackHistoryModel8ba275_TransactionsChanged
func callbackHistoryModel8ba275_TransactionsChanged(ptr unsafe.Pointer, transactions C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "transactionsChanged"); signal != nil {
		(*(*func([]*TransactionDetails))(signal))(func(l C.struct_Moc_PackedList) []*TransactionDetails {
			out := make([]*TransactionDetails, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__transactionsChanged_transactions_atList(i)
			}
			return out
		}(transactions))
	}

}

func (ptr *HistoryModel) ConnectTransactionsChanged(f func(transactions []*TransactionDetails)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transactionsChanged") {
			C.HistoryModel8ba275_ConnectTransactionsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "transactionsChanged"); signal != nil {
			f := func(transactions []*TransactionDetails) {
				(*(*func([]*TransactionDetails))(signal))(transactions)
				f(transactions)
			}
			qt.ConnectSignal(ptr.Pointer(), "transactionsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transactionsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *HistoryModel) DisconnectTransactionsChanged() {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_DisconnectTransactionsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transactionsChanged")
	}
}

func (ptr *HistoryModel) TransactionsChanged(transactions []*TransactionDetails) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_TransactionsChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__transactionsChanged_transactions_newList())
			for _, v := range transactions {
				tmpList.__transactionsChanged_transactions_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func HistoryModel_QRegisterMetaType() int {
	return int(int32(C.HistoryModel8ba275_HistoryModel8ba275_QRegisterMetaType()))
}

func (ptr *HistoryModel) QRegisterMetaType() int {
	return int(int32(C.HistoryModel8ba275_HistoryModel8ba275_QRegisterMetaType()))
}

func HistoryModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.HistoryModel8ba275_HistoryModel8ba275_QRegisterMetaType2(typeNameC)))
}

func (ptr *HistoryModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.HistoryModel8ba275_HistoryModel8ba275_QRegisterMetaType2(typeNameC)))
}

func HistoryModel_QmlRegisterType() int {
	return int(int32(C.HistoryModel8ba275_HistoryModel8ba275_QmlRegisterType()))
}

func (ptr *HistoryModel) QmlRegisterType() int {
	return int(int32(C.HistoryModel8ba275_HistoryModel8ba275_QmlRegisterType()))
}

func HistoryModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.HistoryModel8ba275_HistoryModel8ba275_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *HistoryModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.HistoryModel8ba275_HistoryModel8ba275_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *HistoryModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HistoryModel8ba275_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *HistoryModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *HistoryModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.HistoryModel8ba275_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *HistoryModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HistoryModel8ba275_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *HistoryModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *HistoryModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.HistoryModel8ba275_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *HistoryModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HistoryModel8ba275_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *HistoryModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *HistoryModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.HistoryModel8ba275_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.HistoryModel8ba275___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *HistoryModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.HistoryModel8ba275___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *HistoryModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HistoryModel8ba275___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *HistoryModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *HistoryModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.HistoryModel8ba275___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *HistoryModel) __itemData_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___itemData_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.HistoryModel8ba275___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *HistoryModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.HistoryModel8ba275___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *HistoryModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.HistoryModel8ba275___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *HistoryModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.HistoryModel8ba275___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *HistoryModel) __match_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___match_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.HistoryModel8ba275___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *HistoryModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.HistoryModel8ba275___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *HistoryModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.HistoryModel8ba275___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *HistoryModel) __roleNames_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___roleNames_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.HistoryModel8ba275___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *HistoryModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.HistoryModel8ba275___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *HistoryModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.HistoryModel8ba275___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *HistoryModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HistoryModel8ba275_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *HistoryModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *HistoryModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.HistoryModel8ba275_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *HistoryModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HistoryModel8ba275_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *HistoryModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *HistoryModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.HistoryModel8ba275_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.HistoryModel8ba275___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *HistoryModel) __children_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___children_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.HistoryModel8ba275___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *HistoryModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.HistoryModel8ba275___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *HistoryModel) __findChildren_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___findChildren_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.HistoryModel8ba275___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *HistoryModel) __findChildren_newList3() unsafe.Pointer {
	return C.HistoryModel8ba275___findChildren_newList3(ptr.Pointer())
}

func (ptr *HistoryModel) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.HistoryModel8ba275___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *HistoryModel) __qFindChildren_newList2() unsafe.Pointer {
	return C.HistoryModel8ba275___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *HistoryModel) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.HistoryModel8ba275___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *HistoryModel) __roles_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___roles_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roles_keyList_atList(i)
			}
			return out
		}(C.HistoryModel8ba275___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *HistoryModel) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.HistoryModel8ba275___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *HistoryModel) __setRoles_roles_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___setRoles_roles_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __setRoles_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoles_roles_keyList_atList(i)
			}
			return out
		}(C.HistoryModel8ba275___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *HistoryModel) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.HistoryModel8ba275___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *HistoryModel) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___rolesChanged_roles_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __rolesChanged_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____rolesChanged_roles_keyList_atList(i)
			}
			return out
		}(C.HistoryModel8ba275___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *HistoryModel) __transactions_atList(i int) *TransactionDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewTransactionDetailsFromPointer(C.HistoryModel8ba275___transactions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __transactions_setList(i TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___transactions_setList(ptr.Pointer(), PointerFromTransactionDetails(i))
	}
}

func (ptr *HistoryModel) __transactions_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___transactions_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __setTransactions_transactions_atList(i int) *TransactionDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewTransactionDetailsFromPointer(C.HistoryModel8ba275___setTransactions_transactions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __setTransactions_transactions_setList(i TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___setTransactions_transactions_setList(ptr.Pointer(), PointerFromTransactionDetails(i))
	}
}

func (ptr *HistoryModel) __setTransactions_transactions_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___setTransactions_transactions_newList(ptr.Pointer())
}

func (ptr *HistoryModel) __transactionsChanged_transactions_atList(i int) *TransactionDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewTransactionDetailsFromPointer(C.HistoryModel8ba275___transactionsChanged_transactions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *HistoryModel) __transactionsChanged_transactions_setList(i TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275___transactionsChanged_transactions_setList(ptr.Pointer(), PointerFromTransactionDetails(i))
	}
}

func (ptr *HistoryModel) __transactionsChanged_transactions_newList() unsafe.Pointer {
	return C.HistoryModel8ba275___transactionsChanged_transactions_newList(ptr.Pointer())
}

func (ptr *HistoryModel) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HistoryModel8ba275_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *HistoryModel) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *HistoryModel) ____roles_keyList_newList() unsafe.Pointer {
	return C.HistoryModel8ba275_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *HistoryModel) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HistoryModel8ba275_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *HistoryModel) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *HistoryModel) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.HistoryModel8ba275_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *HistoryModel) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HistoryModel8ba275_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *HistoryModel) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *HistoryModel) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.HistoryModel8ba275_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewHistoryModel(parent std_core.QObject_ITF) *HistoryModel {
	tmpValue := NewHistoryModelFromPointer(C.HistoryModel8ba275_NewHistoryModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackHistoryModel8ba275_DestroyHistoryModel
func callbackHistoryModel8ba275_DestroyHistoryModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~HistoryModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHistoryModelFromPointer(ptr).DestroyHistoryModelDefault()
	}
}

func (ptr *HistoryModel) ConnectDestroyHistoryModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~HistoryModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~HistoryModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~HistoryModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *HistoryModel) DisconnectDestroyHistoryModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~HistoryModel")
	}
}

func (ptr *HistoryModel) DestroyHistoryModel() {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_DestroyHistoryModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *HistoryModel) DestroyHistoryModelDefault() {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_DestroyHistoryModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackHistoryModel8ba275_DropMimeData
func callbackHistoryModel8ba275_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *HistoryModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_Flags
func callbackHistoryModel8ba275_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewHistoryModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *HistoryModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.HistoryModel8ba275_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackHistoryModel8ba275_Index
func callbackHistoryModel8ba275_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewHistoryModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *HistoryModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.HistoryModel8ba275_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackHistoryModel8ba275_Sibling
func callbackHistoryModel8ba275_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewHistoryModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *HistoryModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.HistoryModel8ba275_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackHistoryModel8ba275_Buddy
func callbackHistoryModel8ba275_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewHistoryModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *HistoryModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.HistoryModel8ba275_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackHistoryModel8ba275_CanDropMimeData
func callbackHistoryModel8ba275_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *HistoryModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_CanFetchMore
func callbackHistoryModel8ba275_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *HistoryModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_ColumnCount
func callbackHistoryModel8ba275_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewHistoryModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *HistoryModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HistoryModel8ba275_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackHistoryModel8ba275_ColumnsAboutToBeInserted
func callbackHistoryModel8ba275_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackHistoryModel8ba275_ColumnsAboutToBeMoved
func callbackHistoryModel8ba275_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackHistoryModel8ba275_ColumnsAboutToBeRemoved
func callbackHistoryModel8ba275_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackHistoryModel8ba275_ColumnsInserted
func callbackHistoryModel8ba275_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackHistoryModel8ba275_ColumnsMoved
func callbackHistoryModel8ba275_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackHistoryModel8ba275_ColumnsRemoved
func callbackHistoryModel8ba275_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackHistoryModel8ba275_Data
func callbackHistoryModel8ba275_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewHistoryModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *HistoryModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.HistoryModel8ba275_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackHistoryModel8ba275_DataChanged
func callbackHistoryModel8ba275_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(signal))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackHistoryModel8ba275_FetchMore
func callbackHistoryModel8ba275_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewHistoryModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *HistoryModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackHistoryModel8ba275_HasChildren
func callbackHistoryModel8ba275_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *HistoryModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_HeaderData
func callbackHistoryModel8ba275_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewHistoryModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *HistoryModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.HistoryModel8ba275_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackHistoryModel8ba275_HeaderDataChanged
func callbackHistoryModel8ba275_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackHistoryModel8ba275_InsertColumns
func callbackHistoryModel8ba275_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *HistoryModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_InsertRows
func callbackHistoryModel8ba275_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *HistoryModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_ItemData
func callbackHistoryModel8ba275_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*std_core.QModelIndex) map[int]*std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__itemData_newList())
		for k, v := range NewHistoryModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *HistoryModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.HistoryModel8ba275_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackHistoryModel8ba275_LayoutAboutToBeChanged
func callbackHistoryModel8ba275_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackHistoryModel8ba275_LayoutChanged
func callbackHistoryModel8ba275_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackHistoryModel8ba275_Match
func callbackHistoryModel8ba275_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__match_newList())
		for _, v := range NewHistoryModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *HistoryModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.HistoryModel8ba275_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackHistoryModel8ba275_MimeData
func callbackHistoryModel8ba275_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func([]*std_core.QModelIndex) *std_core.QMimeData)(signal))(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewHistoryModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewHistoryModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *HistoryModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.HistoryModel8ba275_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__mimeData_indexes_newList())
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

//export callbackHistoryModel8ba275_MimeTypes
func callbackHistoryModel8ba275_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewHistoryModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *HistoryModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.HistoryModel8ba275_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackHistoryModel8ba275_ModelAboutToBeReset
func callbackHistoryModel8ba275_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackHistoryModel8ba275_ModelReset
func callbackHistoryModel8ba275_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackHistoryModel8ba275_MoveColumns
func callbackHistoryModel8ba275_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *HistoryModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_MoveRows
func callbackHistoryModel8ba275_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *HistoryModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_Parent
func callbackHistoryModel8ba275_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewHistoryModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *HistoryModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.HistoryModel8ba275_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackHistoryModel8ba275_RemoveColumns
func callbackHistoryModel8ba275_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *HistoryModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_RemoveRows
func callbackHistoryModel8ba275_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *HistoryModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_ResetInternalData
func callbackHistoryModel8ba275_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHistoryModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *HistoryModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackHistoryModel8ba275_Revert
func callbackHistoryModel8ba275_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHistoryModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *HistoryModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_RevertDefault(ptr.Pointer())
	}
}

//export callbackHistoryModel8ba275_RoleNames
func callbackHistoryModel8ba275_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewHistoryModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *HistoryModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.HistoryModel8ba275_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackHistoryModel8ba275_RowCount
func callbackHistoryModel8ba275_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewHistoryModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *HistoryModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.HistoryModel8ba275_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackHistoryModel8ba275_RowsAboutToBeInserted
func callbackHistoryModel8ba275_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackHistoryModel8ba275_RowsAboutToBeMoved
func callbackHistoryModel8ba275_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackHistoryModel8ba275_RowsAboutToBeRemoved
func callbackHistoryModel8ba275_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackHistoryModel8ba275_RowsInserted
func callbackHistoryModel8ba275_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackHistoryModel8ba275_RowsMoved
func callbackHistoryModel8ba275_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackHistoryModel8ba275_RowsRemoved
func callbackHistoryModel8ba275_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackHistoryModel8ba275_SetData
func callbackHistoryModel8ba275_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *HistoryModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_SetHeaderData
func callbackHistoryModel8ba275_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *HistoryModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_SetItemData
func callbackHistoryModel8ba275_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(signal))(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewHistoryModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewHistoryModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *HistoryModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewHistoryModelFromPointer(NewHistoryModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_Sort
func callbackHistoryModel8ba275_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewHistoryModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *HistoryModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackHistoryModel8ba275_Span
func callbackHistoryModel8ba275_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewHistoryModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *HistoryModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.HistoryModel8ba275_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackHistoryModel8ba275_Submit
func callbackHistoryModel8ba275_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *HistoryModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_SupportedDragActions
func callbackHistoryModel8ba275_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewHistoryModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *HistoryModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.HistoryModel8ba275_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackHistoryModel8ba275_SupportedDropActions
func callbackHistoryModel8ba275_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewHistoryModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *HistoryModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.HistoryModel8ba275_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackHistoryModel8ba275_ChildEvent
func callbackHistoryModel8ba275_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewHistoryModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *HistoryModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackHistoryModel8ba275_ConnectNotify
func callbackHistoryModel8ba275_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewHistoryModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *HistoryModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackHistoryModel8ba275_CustomEvent
func callbackHistoryModel8ba275_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewHistoryModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *HistoryModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackHistoryModel8ba275_DeleteLater
func callbackHistoryModel8ba275_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewHistoryModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *HistoryModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackHistoryModel8ba275_Destroyed
func callbackHistoryModel8ba275_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackHistoryModel8ba275_DisconnectNotify
func callbackHistoryModel8ba275_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewHistoryModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *HistoryModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackHistoryModel8ba275_Event
func callbackHistoryModel8ba275_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *HistoryModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_EventFilter
func callbackHistoryModel8ba275_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewHistoryModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *HistoryModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.HistoryModel8ba275_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackHistoryModel8ba275_ObjectNameChanged
func callbackHistoryModel8ba275_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackHistoryModel8ba275_TimerEvent
func callbackHistoryModel8ba275_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewHistoryModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *HistoryModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.HistoryModel8ba275_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type TransactionDetails_ITF interface {
	std_core.QObject_ITF
	TransactionDetails_PTR() *TransactionDetails
}

func (ptr *TransactionDetails) TransactionDetails_PTR() *TransactionDetails {
	return ptr
}

func (ptr *TransactionDetails) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *TransactionDetails) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromTransactionDetails(ptr TransactionDetails_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.TransactionDetails_PTR().Pointer()
	}
	return nil
}

func NewTransactionDetailsFromPointer(ptr unsafe.Pointer) (n *TransactionDetails) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(TransactionDetails)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *TransactionDetails:
			n = deduced

		case *std_core.QObject:
			n = &TransactionDetails{QObject: *deduced}

		default:
			n = new(TransactionDetails)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackTransactionDetails8ba275_Constructor
func callbackTransactionDetails8ba275_Constructor(ptr unsafe.Pointer) {
	this := NewTransactionDetailsFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackTransactionDetails8ba275_Date
func callbackTransactionDetails8ba275_Date(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "date"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTransactionDetailsFromPointer(ptr).DateDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *TransactionDetails) ConnectDate(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "date"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "date", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "date", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectDate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "date")
	}
}

func (ptr *TransactionDetails) Date() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails8ba275_Date(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) DateDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails8ba275_DateDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetails8ba275_SetDate
func callbackTransactionDetails8ba275_SetDate(ptr unsafe.Pointer, date C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setDate"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(date))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetDateDefault(cGoUnpackString(date))
	}
}

func (ptr *TransactionDetails) ConnectSetDate(f func(date string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDate"); signal != nil {
			f := func(date string) {
				(*(*func(string))(signal))(date)
				f(date)
			}
			qt.ConnectSignal(ptr.Pointer(), "setDate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setDate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetDate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setDate")
	}
}

func (ptr *TransactionDetails) SetDate(date string) {
	if ptr.Pointer() != nil {
		var dateC *C.char
		if date != "" {
			dateC = C.CString(date)
			defer C.free(unsafe.Pointer(dateC))
		}
		C.TransactionDetails8ba275_SetDate(ptr.Pointer(), C.struct_Moc_PackedString{data: dateC, len: C.longlong(len(date))})
	}
}

func (ptr *TransactionDetails) SetDateDefault(date string) {
	if ptr.Pointer() != nil {
		var dateC *C.char
		if date != "" {
			dateC = C.CString(date)
			defer C.free(unsafe.Pointer(dateC))
		}
		C.TransactionDetails8ba275_SetDateDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: dateC, len: C.longlong(len(date))})
	}
}

//export callbackTransactionDetails8ba275_DateChanged
func callbackTransactionDetails8ba275_DateChanged(ptr unsafe.Pointer, date C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "dateChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(date))
	}

}

func (ptr *TransactionDetails) ConnectDateChanged(f func(date string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dateChanged") {
			C.TransactionDetails8ba275_ConnectDateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dateChanged"); signal != nil {
			f := func(date string) {
				(*(*func(string))(signal))(date)
				f(date)
			}
			qt.ConnectSignal(ptr.Pointer(), "dateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dateChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectDateChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DisconnectDateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dateChanged")
	}
}

func (ptr *TransactionDetails) DateChanged(date string) {
	if ptr.Pointer() != nil {
		var dateC *C.char
		if date != "" {
			dateC = C.CString(date)
			defer C.free(unsafe.Pointer(dateC))
		}
		C.TransactionDetails8ba275_DateChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: dateC, len: C.longlong(len(date))})
	}
}

//export callbackTransactionDetails8ba275_Status
func callbackTransactionDetails8ba275_Status(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "status"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewTransactionDetailsFromPointer(ptr).StatusDefault()))
}

func (ptr *TransactionDetails) ConnectStatus(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "status"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "status", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "status", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectStatus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "status")
	}
}

func (ptr *TransactionDetails) Status() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails8ba275_Status(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) StatusDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails8ba275_StatusDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails8ba275_SetStatus
func callbackTransactionDetails8ba275_SetStatus(ptr unsafe.Pointer, status C.int) {
	if signal := qt.GetSignal(ptr, "setStatus"); signal != nil {
		(*(*func(int))(signal))(int(int32(status)))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetStatusDefault(int(int32(status)))
	}
}

func (ptr *TransactionDetails) ConnectSetStatus(f func(status int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setStatus"); signal != nil {
			f := func(status int) {
				(*(*func(int))(signal))(status)
				f(status)
			}
			qt.ConnectSignal(ptr.Pointer(), "setStatus", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setStatus", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetStatus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setStatus")
	}
}

func (ptr *TransactionDetails) SetStatus(status int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_SetStatus(ptr.Pointer(), C.int(int32(status)))
	}
}

func (ptr *TransactionDetails) SetStatusDefault(status int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_SetStatusDefault(ptr.Pointer(), C.int(int32(status)))
	}
}

//export callbackTransactionDetails8ba275_StatusChanged
func callbackTransactionDetails8ba275_StatusChanged(ptr unsafe.Pointer, status C.int) {
	if signal := qt.GetSignal(ptr, "statusChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(status)))
	}

}

func (ptr *TransactionDetails) ConnectStatusChanged(f func(status int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "statusChanged") {
			C.TransactionDetails8ba275_ConnectStatusChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "statusChanged"); signal != nil {
			f := func(status int) {
				(*(*func(int))(signal))(status)
				f(status)
			}
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "statusChanged")
	}
}

func (ptr *TransactionDetails) StatusChanged(status int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_StatusChanged(ptr.Pointer(), C.int(int32(status)))
	}
}

//export callbackTransactionDetails8ba275_Type
func callbackTransactionDetails8ba275_Type(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewTransactionDetailsFromPointer(ptr).TypeDefault()))
}

func (ptr *TransactionDetails) ConnectType(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *TransactionDetails) Type() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails8ba275_Type(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) TypeDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails8ba275_TypeDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails8ba275_SetType
func callbackTransactionDetails8ba275_SetType(ptr unsafe.Pointer, ty C.int) {
	if signal := qt.GetSignal(ptr, "setType"); signal != nil {
		(*(*func(int))(signal))(int(int32(ty)))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetTypeDefault(int(int32(ty)))
	}
}

func (ptr *TransactionDetails) ConnectSetType(f func(ty int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setType"); signal != nil {
			f := func(ty int) {
				(*(*func(int))(signal))(ty)
				f(ty)
			}
			qt.ConnectSignal(ptr.Pointer(), "setType", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setType", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setType")
	}
}

func (ptr *TransactionDetails) SetType(ty int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_SetType(ptr.Pointer(), C.int(int32(ty)))
	}
}

func (ptr *TransactionDetails) SetTypeDefault(ty int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_SetTypeDefault(ptr.Pointer(), C.int(int32(ty)))
	}
}

//export callbackTransactionDetails8ba275_TypeChanged
func callbackTransactionDetails8ba275_TypeChanged(ptr unsafe.Pointer, ty C.int) {
	if signal := qt.GetSignal(ptr, "typeChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(ty)))
	}

}

func (ptr *TransactionDetails) ConnectTypeChanged(f func(ty int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "typeChanged") {
			C.TransactionDetails8ba275_ConnectTypeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "typeChanged"); signal != nil {
			f := func(ty int) {
				(*(*func(int))(signal))(ty)
				f(ty)
			}
			qt.ConnectSignal(ptr.Pointer(), "typeChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "typeChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectTypeChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "typeChanged")
	}
}

func (ptr *TransactionDetails) TypeChanged(ty int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_TypeChanged(ptr.Pointer(), C.int(int32(ty)))
	}
}

//export callbackTransactionDetails8ba275_Amount
func callbackTransactionDetails8ba275_Amount(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "amount"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewTransactionDetailsFromPointer(ptr).AmountDefault()))
}

func (ptr *TransactionDetails) ConnectAmount(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "amount"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "amount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "amount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectAmount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "amount")
	}
}

func (ptr *TransactionDetails) Amount() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails8ba275_Amount(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) AmountDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails8ba275_AmountDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails8ba275_SetAmount
func callbackTransactionDetails8ba275_SetAmount(ptr unsafe.Pointer, amount C.int) {
	if signal := qt.GetSignal(ptr, "setAmount"); signal != nil {
		(*(*func(int))(signal))(int(int32(amount)))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetAmountDefault(int(int32(amount)))
	}
}

func (ptr *TransactionDetails) ConnectSetAmount(f func(amount int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAmount"); signal != nil {
			f := func(amount int) {
				(*(*func(int))(signal))(amount)
				f(amount)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAmount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAmount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetAmount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAmount")
	}
}

func (ptr *TransactionDetails) SetAmount(amount int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_SetAmount(ptr.Pointer(), C.int(int32(amount)))
	}
}

func (ptr *TransactionDetails) SetAmountDefault(amount int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_SetAmountDefault(ptr.Pointer(), C.int(int32(amount)))
	}
}

//export callbackTransactionDetails8ba275_AmountChanged
func callbackTransactionDetails8ba275_AmountChanged(ptr unsafe.Pointer, amount C.int) {
	if signal := qt.GetSignal(ptr, "amountChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(amount)))
	}

}

func (ptr *TransactionDetails) ConnectAmountChanged(f func(amount int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "amountChanged") {
			C.TransactionDetails8ba275_ConnectAmountChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "amountChanged"); signal != nil {
			f := func(amount int) {
				(*(*func(int))(signal))(amount)
				f(amount)
			}
			qt.ConnectSignal(ptr.Pointer(), "amountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "amountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectAmountChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DisconnectAmountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "amountChanged")
	}
}

func (ptr *TransactionDetails) AmountChanged(amount int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_AmountChanged(ptr.Pointer(), C.int(int32(amount)))
	}
}

//export callbackTransactionDetails8ba275_HoursReceived
func callbackTransactionDetails8ba275_HoursReceived(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "hoursReceived"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewTransactionDetailsFromPointer(ptr).HoursReceivedDefault()))
}

func (ptr *TransactionDetails) ConnectHoursReceived(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hoursReceived"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hoursReceived", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoursReceived", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectHoursReceived() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hoursReceived")
	}
}

func (ptr *TransactionDetails) HoursReceived() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails8ba275_HoursReceived(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) HoursReceivedDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails8ba275_HoursReceivedDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails8ba275_SetHoursReceived
func callbackTransactionDetails8ba275_SetHoursReceived(ptr unsafe.Pointer, hoursReceived C.int) {
	if signal := qt.GetSignal(ptr, "setHoursReceived"); signal != nil {
		(*(*func(int))(signal))(int(int32(hoursReceived)))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetHoursReceivedDefault(int(int32(hoursReceived)))
	}
}

func (ptr *TransactionDetails) ConnectSetHoursReceived(f func(hoursReceived int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHoursReceived"); signal != nil {
			f := func(hoursReceived int) {
				(*(*func(int))(signal))(hoursReceived)
				f(hoursReceived)
			}
			qt.ConnectSignal(ptr.Pointer(), "setHoursReceived", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setHoursReceived", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetHoursReceived() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHoursReceived")
	}
}

func (ptr *TransactionDetails) SetHoursReceived(hoursReceived int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_SetHoursReceived(ptr.Pointer(), C.int(int32(hoursReceived)))
	}
}

func (ptr *TransactionDetails) SetHoursReceivedDefault(hoursReceived int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_SetHoursReceivedDefault(ptr.Pointer(), C.int(int32(hoursReceived)))
	}
}

//export callbackTransactionDetails8ba275_HoursReceivedChanged
func callbackTransactionDetails8ba275_HoursReceivedChanged(ptr unsafe.Pointer, hoursReceived C.int) {
	if signal := qt.GetSignal(ptr, "hoursReceivedChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(hoursReceived)))
	}

}

func (ptr *TransactionDetails) ConnectHoursReceivedChanged(f func(hoursReceived int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hoursReceivedChanged") {
			C.TransactionDetails8ba275_ConnectHoursReceivedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hoursReceivedChanged"); signal != nil {
			f := func(hoursReceived int) {
				(*(*func(int))(signal))(hoursReceived)
				f(hoursReceived)
			}
			qt.ConnectSignal(ptr.Pointer(), "hoursReceivedChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoursReceivedChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectHoursReceivedChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DisconnectHoursReceivedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hoursReceivedChanged")
	}
}

func (ptr *TransactionDetails) HoursReceivedChanged(hoursReceived int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_HoursReceivedChanged(ptr.Pointer(), C.int(int32(hoursReceived)))
	}
}

//export callbackTransactionDetails8ba275_HoursBurned
func callbackTransactionDetails8ba275_HoursBurned(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "hoursBurned"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewTransactionDetailsFromPointer(ptr).HoursBurnedDefault()))
}

func (ptr *TransactionDetails) ConnectHoursBurned(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hoursBurned"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hoursBurned", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoursBurned", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectHoursBurned() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hoursBurned")
	}
}

func (ptr *TransactionDetails) HoursBurned() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails8ba275_HoursBurned(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) HoursBurnedDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails8ba275_HoursBurnedDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails8ba275_SetHoursBurned
func callbackTransactionDetails8ba275_SetHoursBurned(ptr unsafe.Pointer, hoursBurned C.int) {
	if signal := qt.GetSignal(ptr, "setHoursBurned"); signal != nil {
		(*(*func(int))(signal))(int(int32(hoursBurned)))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetHoursBurnedDefault(int(int32(hoursBurned)))
	}
}

func (ptr *TransactionDetails) ConnectSetHoursBurned(f func(hoursBurned int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHoursBurned"); signal != nil {
			f := func(hoursBurned int) {
				(*(*func(int))(signal))(hoursBurned)
				f(hoursBurned)
			}
			qt.ConnectSignal(ptr.Pointer(), "setHoursBurned", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setHoursBurned", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetHoursBurned() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHoursBurned")
	}
}

func (ptr *TransactionDetails) SetHoursBurned(hoursBurned int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_SetHoursBurned(ptr.Pointer(), C.int(int32(hoursBurned)))
	}
}

func (ptr *TransactionDetails) SetHoursBurnedDefault(hoursBurned int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_SetHoursBurnedDefault(ptr.Pointer(), C.int(int32(hoursBurned)))
	}
}

//export callbackTransactionDetails8ba275_HoursBurnedChanged
func callbackTransactionDetails8ba275_HoursBurnedChanged(ptr unsafe.Pointer, hoursBurned C.int) {
	if signal := qt.GetSignal(ptr, "hoursBurnedChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(hoursBurned)))
	}

}

func (ptr *TransactionDetails) ConnectHoursBurnedChanged(f func(hoursBurned int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hoursBurnedChanged") {
			C.TransactionDetails8ba275_ConnectHoursBurnedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hoursBurnedChanged"); signal != nil {
			f := func(hoursBurned int) {
				(*(*func(int))(signal))(hoursBurned)
				f(hoursBurned)
			}
			qt.ConnectSignal(ptr.Pointer(), "hoursBurnedChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoursBurnedChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectHoursBurnedChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DisconnectHoursBurnedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hoursBurnedChanged")
	}
}

func (ptr *TransactionDetails) HoursBurnedChanged(hoursBurned int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_HoursBurnedChanged(ptr.Pointer(), C.int(int32(hoursBurned)))
	}
}

//export callbackTransactionDetails8ba275_TransactionID
func callbackTransactionDetails8ba275_TransactionID(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "transactionID"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTransactionDetailsFromPointer(ptr).TransactionIDDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *TransactionDetails) ConnectTransactionID(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "transactionID"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "transactionID", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transactionID", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectTransactionID() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "transactionID")
	}
}

func (ptr *TransactionDetails) TransactionID() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails8ba275_TransactionID(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) TransactionIDDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails8ba275_TransactionIDDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetails8ba275_SetTransactionID
func callbackTransactionDetails8ba275_SetTransactionID(ptr unsafe.Pointer, transactionID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTransactionID"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(transactionID))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetTransactionIDDefault(cGoUnpackString(transactionID))
	}
}

func (ptr *TransactionDetails) ConnectSetTransactionID(f func(transactionID string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTransactionID"); signal != nil {
			f := func(transactionID string) {
				(*(*func(string))(signal))(transactionID)
				f(transactionID)
			}
			qt.ConnectSignal(ptr.Pointer(), "setTransactionID", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTransactionID", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetTransactionID() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTransactionID")
	}
}

func (ptr *TransactionDetails) SetTransactionID(transactionID string) {
	if ptr.Pointer() != nil {
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		C.TransactionDetails8ba275_SetTransactionID(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

func (ptr *TransactionDetails) SetTransactionIDDefault(transactionID string) {
	if ptr.Pointer() != nil {
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		C.TransactionDetails8ba275_SetTransactionIDDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

//export callbackTransactionDetails8ba275_TransactionIDChanged
func callbackTransactionDetails8ba275_TransactionIDChanged(ptr unsafe.Pointer, transactionID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transactionIDChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(transactionID))
	}

}

func (ptr *TransactionDetails) ConnectTransactionIDChanged(f func(transactionID string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transactionIDChanged") {
			C.TransactionDetails8ba275_ConnectTransactionIDChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "transactionIDChanged"); signal != nil {
			f := func(transactionID string) {
				(*(*func(string))(signal))(transactionID)
				f(transactionID)
			}
			qt.ConnectSignal(ptr.Pointer(), "transactionIDChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transactionIDChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectTransactionIDChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DisconnectTransactionIDChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transactionIDChanged")
	}
}

func (ptr *TransactionDetails) TransactionIDChanged(transactionID string) {
	if ptr.Pointer() != nil {
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		C.TransactionDetails8ba275_TransactionIDChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

//export callbackTransactionDetails8ba275_SentAddress
func callbackTransactionDetails8ba275_SentAddress(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "sentAddress"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTransactionDetailsFromPointer(ptr).SentAddressDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *TransactionDetails) ConnectSentAddress(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "sentAddress"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "sentAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sentAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSentAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "sentAddress")
	}
}

func (ptr *TransactionDetails) SentAddress() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails8ba275_SentAddress(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) SentAddressDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails8ba275_SentAddressDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetails8ba275_SetSentAddress
func callbackTransactionDetails8ba275_SetSentAddress(ptr unsafe.Pointer, sentAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setSentAddress"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(sentAddress))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetSentAddressDefault(cGoUnpackString(sentAddress))
	}
}

func (ptr *TransactionDetails) ConnectSetSentAddress(f func(sentAddress string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSentAddress"); signal != nil {
			f := func(sentAddress string) {
				(*(*func(string))(signal))(sentAddress)
				f(sentAddress)
			}
			qt.ConnectSignal(ptr.Pointer(), "setSentAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSentAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetSentAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSentAddress")
	}
}

func (ptr *TransactionDetails) SetSentAddress(sentAddress string) {
	if ptr.Pointer() != nil {
		var sentAddressC *C.char
		if sentAddress != "" {
			sentAddressC = C.CString(sentAddress)
			defer C.free(unsafe.Pointer(sentAddressC))
		}
		C.TransactionDetails8ba275_SetSentAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: sentAddressC, len: C.longlong(len(sentAddress))})
	}
}

func (ptr *TransactionDetails) SetSentAddressDefault(sentAddress string) {
	if ptr.Pointer() != nil {
		var sentAddressC *C.char
		if sentAddress != "" {
			sentAddressC = C.CString(sentAddress)
			defer C.free(unsafe.Pointer(sentAddressC))
		}
		C.TransactionDetails8ba275_SetSentAddressDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: sentAddressC, len: C.longlong(len(sentAddress))})
	}
}

//export callbackTransactionDetails8ba275_SentAddressChanged
func callbackTransactionDetails8ba275_SentAddressChanged(ptr unsafe.Pointer, sentAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sentAddressChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(sentAddress))
	}

}

func (ptr *TransactionDetails) ConnectSentAddressChanged(f func(sentAddress string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sentAddressChanged") {
			C.TransactionDetails8ba275_ConnectSentAddressChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sentAddressChanged"); signal != nil {
			f := func(sentAddress string) {
				(*(*func(string))(signal))(sentAddress)
				f(sentAddress)
			}
			qt.ConnectSignal(ptr.Pointer(), "sentAddressChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sentAddressChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSentAddressChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DisconnectSentAddressChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sentAddressChanged")
	}
}

func (ptr *TransactionDetails) SentAddressChanged(sentAddress string) {
	if ptr.Pointer() != nil {
		var sentAddressC *C.char
		if sentAddress != "" {
			sentAddressC = C.CString(sentAddress)
			defer C.free(unsafe.Pointer(sentAddressC))
		}
		C.TransactionDetails8ba275_SentAddressChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: sentAddressC, len: C.longlong(len(sentAddress))})
	}
}

//export callbackTransactionDetails8ba275_ReceivedAddress
func callbackTransactionDetails8ba275_ReceivedAddress(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "receivedAddress"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTransactionDetailsFromPointer(ptr).ReceivedAddressDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *TransactionDetails) ConnectReceivedAddress(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "receivedAddress"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "receivedAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "receivedAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectReceivedAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "receivedAddress")
	}
}

func (ptr *TransactionDetails) ReceivedAddress() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails8ba275_ReceivedAddress(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) ReceivedAddressDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails8ba275_ReceivedAddressDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetails8ba275_SetReceivedAddress
func callbackTransactionDetails8ba275_SetReceivedAddress(ptr unsafe.Pointer, receivedAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setReceivedAddress"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(receivedAddress))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetReceivedAddressDefault(cGoUnpackString(receivedAddress))
	}
}

func (ptr *TransactionDetails) ConnectSetReceivedAddress(f func(receivedAddress string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setReceivedAddress"); signal != nil {
			f := func(receivedAddress string) {
				(*(*func(string))(signal))(receivedAddress)
				f(receivedAddress)
			}
			qt.ConnectSignal(ptr.Pointer(), "setReceivedAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setReceivedAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetReceivedAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setReceivedAddress")
	}
}

func (ptr *TransactionDetails) SetReceivedAddress(receivedAddress string) {
	if ptr.Pointer() != nil {
		var receivedAddressC *C.char
		if receivedAddress != "" {
			receivedAddressC = C.CString(receivedAddress)
			defer C.free(unsafe.Pointer(receivedAddressC))
		}
		C.TransactionDetails8ba275_SetReceivedAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: receivedAddressC, len: C.longlong(len(receivedAddress))})
	}
}

func (ptr *TransactionDetails) SetReceivedAddressDefault(receivedAddress string) {
	if ptr.Pointer() != nil {
		var receivedAddressC *C.char
		if receivedAddress != "" {
			receivedAddressC = C.CString(receivedAddress)
			defer C.free(unsafe.Pointer(receivedAddressC))
		}
		C.TransactionDetails8ba275_SetReceivedAddressDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: receivedAddressC, len: C.longlong(len(receivedAddress))})
	}
}

//export callbackTransactionDetails8ba275_ReceivedAddressChanged
func callbackTransactionDetails8ba275_ReceivedAddressChanged(ptr unsafe.Pointer, receivedAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "receivedAddressChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(receivedAddress))
	}

}

func (ptr *TransactionDetails) ConnectReceivedAddressChanged(f func(receivedAddress string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "receivedAddressChanged") {
			C.TransactionDetails8ba275_ConnectReceivedAddressChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "receivedAddressChanged"); signal != nil {
			f := func(receivedAddress string) {
				(*(*func(string))(signal))(receivedAddress)
				f(receivedAddress)
			}
			qt.ConnectSignal(ptr.Pointer(), "receivedAddressChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "receivedAddressChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectReceivedAddressChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DisconnectReceivedAddressChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "receivedAddressChanged")
	}
}

func (ptr *TransactionDetails) ReceivedAddressChanged(receivedAddress string) {
	if ptr.Pointer() != nil {
		var receivedAddressC *C.char
		if receivedAddress != "" {
			receivedAddressC = C.CString(receivedAddress)
			defer C.free(unsafe.Pointer(receivedAddressC))
		}
		C.TransactionDetails8ba275_ReceivedAddressChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: receivedAddressC, len: C.longlong(len(receivedAddress))})
	}
}

func TransactionDetails_QRegisterMetaType() int {
	return int(int32(C.TransactionDetails8ba275_TransactionDetails8ba275_QRegisterMetaType()))
}

func (ptr *TransactionDetails) QRegisterMetaType() int {
	return int(int32(C.TransactionDetails8ba275_TransactionDetails8ba275_QRegisterMetaType()))
}

func TransactionDetails_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TransactionDetails8ba275_TransactionDetails8ba275_QRegisterMetaType2(typeNameC)))
}

func (ptr *TransactionDetails) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TransactionDetails8ba275_TransactionDetails8ba275_QRegisterMetaType2(typeNameC)))
}

func TransactionDetails_QmlRegisterType() int {
	return int(int32(C.TransactionDetails8ba275_TransactionDetails8ba275_QmlRegisterType()))
}

func (ptr *TransactionDetails) QmlRegisterType() int {
	return int(int32(C.TransactionDetails8ba275_TransactionDetails8ba275_QmlRegisterType()))
}

func TransactionDetails_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.TransactionDetails8ba275_TransactionDetails8ba275_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TransactionDetails) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.TransactionDetails8ba275_TransactionDetails8ba275_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TransactionDetails) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetails8ba275___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __children_newList() unsafe.Pointer {
	return C.TransactionDetails8ba275___children_newList(ptr.Pointer())
}

func (ptr *TransactionDetails) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionDetails8ba275___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionDetails) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.TransactionDetails8ba275___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *TransactionDetails) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetails8ba275___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __findChildren_newList() unsafe.Pointer {
	return C.TransactionDetails8ba275___findChildren_newList(ptr.Pointer())
}

func (ptr *TransactionDetails) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetails8ba275___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __findChildren_newList3() unsafe.Pointer {
	return C.TransactionDetails8ba275___findChildren_newList3(ptr.Pointer())
}

func (ptr *TransactionDetails) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetails8ba275___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __qFindChildren_newList2() unsafe.Pointer {
	return C.TransactionDetails8ba275___qFindChildren_newList2(ptr.Pointer())
}

func NewTransactionDetails(parent std_core.QObject_ITF) *TransactionDetails {
	tmpValue := NewTransactionDetailsFromPointer(C.TransactionDetails8ba275_NewTransactionDetails(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTransactionDetails8ba275_DestroyTransactionDetails
func callbackTransactionDetails8ba275_DestroyTransactionDetails(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~TransactionDetails"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionDetailsFromPointer(ptr).DestroyTransactionDetailsDefault()
	}
}

func (ptr *TransactionDetails) ConnectDestroyTransactionDetails(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~TransactionDetails"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~TransactionDetails", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~TransactionDetails", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectDestroyTransactionDetails() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~TransactionDetails")
	}
}

func (ptr *TransactionDetails) DestroyTransactionDetails() {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DestroyTransactionDetails(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *TransactionDetails) DestroyTransactionDetailsDefault() {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DestroyTransactionDetailsDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTransactionDetails8ba275_ChildEvent
func callbackTransactionDetails8ba275_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTransactionDetailsFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *TransactionDetails) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTransactionDetails8ba275_ConnectNotify
func callbackTransactionDetails8ba275_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTransactionDetailsFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TransactionDetails) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTransactionDetails8ba275_CustomEvent
func callbackTransactionDetails8ba275_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewTransactionDetailsFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *TransactionDetails) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTransactionDetails8ba275_DeleteLater
func callbackTransactionDetails8ba275_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionDetailsFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *TransactionDetails) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTransactionDetails8ba275_Destroyed
func callbackTransactionDetails8ba275_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackTransactionDetails8ba275_DisconnectNotify
func callbackTransactionDetails8ba275_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTransactionDetailsFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TransactionDetails) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTransactionDetails8ba275_Event
func callbackTransactionDetails8ba275_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionDetailsFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *TransactionDetails) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionDetails8ba275_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackTransactionDetails8ba275_EventFilter
func callbackTransactionDetails8ba275_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionDetailsFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *TransactionDetails) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionDetails8ba275_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackTransactionDetails8ba275_ObjectNameChanged
func callbackTransactionDetails8ba275_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackTransactionDetails8ba275_TimerEvent
func callbackTransactionDetails8ba275_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTransactionDetailsFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *TransactionDetails) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails8ba275_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

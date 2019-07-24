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

//export callbackTransactionList742340_Constructor
func callbackTransactionList742340_Constructor(ptr unsafe.Pointer) {
	this := NewTransactionListFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectAddTransaction(this.addTransaction)
	this.ConnectRemoveTransaction(this.removeTransaction)
	this.init()
}

//export callbackTransactionList742340_AddTransaction
func callbackTransactionList742340_AddTransaction(ptr unsafe.Pointer, transaction unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addTransaction"); signal != nil {
		(*(*func(*TransactionDetails))(signal))(NewTransactionDetailsFromPointer(transaction))
	}

}

func (ptr *TransactionList) ConnectAddTransaction(f func(transaction *TransactionDetails)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addTransaction") {
			C.TransactionList742340_ConnectAddTransaction(ptr.Pointer())
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

func (ptr *TransactionList) DisconnectAddTransaction() {
	if ptr.Pointer() != nil {
		C.TransactionList742340_DisconnectAddTransaction(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addTransaction")
	}
}

func (ptr *TransactionList) AddTransaction(transaction TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_AddTransaction(ptr.Pointer(), PointerFromTransactionDetails(transaction))
	}
}

//export callbackTransactionList742340_RemoveTransaction
func callbackTransactionList742340_RemoveTransaction(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "removeTransaction"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *TransactionList) ConnectRemoveTransaction(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "removeTransaction") {
			C.TransactionList742340_ConnectRemoveTransaction(ptr.Pointer())
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
		C.TransactionList742340_DisconnectRemoveTransaction(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "removeTransaction")
	}
}

func (ptr *TransactionList) RemoveTransaction(index int) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_RemoveTransaction(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackTransactionList742340_Roles
func callbackTransactionList742340_Roles(ptr unsafe.Pointer) unsafe.Pointer {
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
		}(C.TransactionList742340_Roles(ptr.Pointer()))
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
		}(C.TransactionList742340_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackTransactionList742340_SetRoles
func callbackTransactionList742340_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
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
		C.TransactionList742340_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
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
		C.TransactionList742340_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackTransactionList742340_RolesChanged
func callbackTransactionList742340_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
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
			C.TransactionList742340_ConnectRolesChanged(ptr.Pointer())
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
		C.TransactionList742340_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *TransactionList) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackTransactionList742340_Transactions
func callbackTransactionList742340_Transactions(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "transactions"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__transactions_newList())
			for _, v := range (*(*func() []*TransactionDetails)(signal))() {
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

func (ptr *TransactionList) ConnectTransactions(f func() []*TransactionDetails) {
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

func (ptr *TransactionList) DisconnectTransactions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "transactions")
	}
}

func (ptr *TransactionList) Transactions() []*TransactionDetails {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*TransactionDetails {
			out := make([]*TransactionDetails, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__transactions_atList(i)
			}
			return out
		}(C.TransactionList742340_Transactions(ptr.Pointer()))
	}
	return make([]*TransactionDetails, 0)
}

func (ptr *TransactionList) TransactionsDefault() []*TransactionDetails {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*TransactionDetails {
			out := make([]*TransactionDetails, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__transactions_atList(i)
			}
			return out
		}(C.TransactionList742340_TransactionsDefault(ptr.Pointer()))
	}
	return make([]*TransactionDetails, 0)
}

//export callbackTransactionList742340_SetTransactions
func callbackTransactionList742340_SetTransactions(ptr unsafe.Pointer, transactions C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setTransactions"); signal != nil {
		(*(*func([]*TransactionDetails))(signal))(func(l C.struct_Moc_PackedList) []*TransactionDetails {
			out := make([]*TransactionDetails, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setTransactions_transactions_atList(i)
			}
			return out
		}(transactions))
	} else {
		NewTransactionListFromPointer(ptr).SetTransactionsDefault(func(l C.struct_Moc_PackedList) []*TransactionDetails {
			out := make([]*TransactionDetails, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setTransactions_transactions_atList(i)
			}
			return out
		}(transactions))
	}
}

func (ptr *TransactionList) ConnectSetTransactions(f func(transactions []*TransactionDetails)) {
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

func (ptr *TransactionList) DisconnectSetTransactions() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTransactions")
	}
}

func (ptr *TransactionList) SetTransactions(transactions []*TransactionDetails) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_SetTransactions(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__setTransactions_transactions_newList())
			for _, v := range transactions {
				tmpList.__setTransactions_transactions_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *TransactionList) SetTransactionsDefault(transactions []*TransactionDetails) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_SetTransactionsDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__setTransactions_transactions_newList())
			for _, v := range transactions {
				tmpList.__setTransactions_transactions_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackTransactionList742340_TransactionsChanged
func callbackTransactionList742340_TransactionsChanged(ptr unsafe.Pointer, transactions C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "transactionsChanged"); signal != nil {
		(*(*func([]*TransactionDetails))(signal))(func(l C.struct_Moc_PackedList) []*TransactionDetails {
			out := make([]*TransactionDetails, int(l.len))
			tmpList := NewTransactionListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__transactionsChanged_transactions_atList(i)
			}
			return out
		}(transactions))
	}

}

func (ptr *TransactionList) ConnectTransactionsChanged(f func(transactions []*TransactionDetails)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transactionsChanged") {
			C.TransactionList742340_ConnectTransactionsChanged(ptr.Pointer())
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

func (ptr *TransactionList) DisconnectTransactionsChanged() {
	if ptr.Pointer() != nil {
		C.TransactionList742340_DisconnectTransactionsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transactionsChanged")
	}
}

func (ptr *TransactionList) TransactionsChanged(transactions []*TransactionDetails) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_TransactionsChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__transactionsChanged_transactions_newList())
			for _, v := range transactions {
				tmpList.__transactionsChanged_transactions_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func TransactionList_QRegisterMetaType() int {
	return int(int32(C.TransactionList742340_TransactionList742340_QRegisterMetaType()))
}

func (ptr *TransactionList) QRegisterMetaType() int {
	return int(int32(C.TransactionList742340_TransactionList742340_QRegisterMetaType()))
}

func TransactionList_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TransactionList742340_TransactionList742340_QRegisterMetaType2(typeNameC)))
}

func (ptr *TransactionList) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TransactionList742340_TransactionList742340_QRegisterMetaType2(typeNameC)))
}

func TransactionList_QmlRegisterType() int {
	return int(int32(C.TransactionList742340_TransactionList742340_QmlRegisterType()))
}

func (ptr *TransactionList) QmlRegisterType() int {
	return int(int32(C.TransactionList742340_TransactionList742340_QmlRegisterType()))
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
	return int(int32(C.TransactionList742340_TransactionList742340_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.TransactionList742340_TransactionList742340_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TransactionList) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList742340_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____itemData_keyList_newList() unsafe.Pointer {
	return C.TransactionList742340_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList742340_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.TransactionList742340_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList742340_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.TransactionList742340_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList742340___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TransactionList) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.TransactionList742340___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *TransactionList) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList742340___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TransactionList) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.TransactionList742340___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *TransactionList) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList742340___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) __dataChanged_roles_newList() unsafe.Pointer {
	return C.TransactionList742340___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *TransactionList) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TransactionList742340___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *TransactionList) __itemData_newList() unsafe.Pointer {
	return C.TransactionList742340___itemData_newList(ptr.Pointer())
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
		}(C.TransactionList742340___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TransactionList) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.TransactionList742340___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *TransactionList) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.TransactionList742340___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *TransactionList) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.TransactionList742340___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *TransactionList) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.TransactionList742340___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *TransactionList) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList742340___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TransactionList) __match_newList() unsafe.Pointer {
	return C.TransactionList742340___match_newList(ptr.Pointer())
}

func (ptr *TransactionList) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList742340___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TransactionList) __mimeData_indexes_newList() unsafe.Pointer {
	return C.TransactionList742340___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *TransactionList) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList742340___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *TransactionList) __persistentIndexList_newList() unsafe.Pointer {
	return C.TransactionList742340___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *TransactionList) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionList742340___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionList) __roleNames_newList() unsafe.Pointer {
	return C.TransactionList742340___roleNames_newList(ptr.Pointer())
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
		}(C.TransactionList742340___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TransactionList) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TransactionList742340___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *TransactionList) __setItemData_roles_newList() unsafe.Pointer {
	return C.TransactionList742340___setItemData_roles_newList(ptr.Pointer())
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
		}(C.TransactionList742340___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TransactionList) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList742340_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.TransactionList742340_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList742340_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.TransactionList742340_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionList742340___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionList) __children_newList() unsafe.Pointer {
	return C.TransactionList742340___children_newList(ptr.Pointer())
}

func (ptr *TransactionList) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionList742340___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionList) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.TransactionList742340___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *TransactionList) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionList742340___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionList) __findChildren_newList() unsafe.Pointer {
	return C.TransactionList742340___findChildren_newList(ptr.Pointer())
}

func (ptr *TransactionList) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionList742340___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionList) __findChildren_newList3() unsafe.Pointer {
	return C.TransactionList742340___findChildren_newList3(ptr.Pointer())
}

func (ptr *TransactionList) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionList742340___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionList) __qFindChildren_newList2() unsafe.Pointer {
	return C.TransactionList742340___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *TransactionList) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionList742340___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionList) __roles_newList() unsafe.Pointer {
	return C.TransactionList742340___roles_newList(ptr.Pointer())
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
		}(C.TransactionList742340___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TransactionList) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionList742340___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionList) __setRoles_roles_newList() unsafe.Pointer {
	return C.TransactionList742340___setRoles_roles_newList(ptr.Pointer())
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
		}(C.TransactionList742340___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TransactionList) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionList742340___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionList) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.TransactionList742340___rolesChanged_roles_newList(ptr.Pointer())
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
		}(C.TransactionList742340___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *TransactionList) __transactions_atList(i int) *TransactionDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewTransactionDetailsFromPointer(C.TransactionList742340___transactions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __transactions_setList(i TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___transactions_setList(ptr.Pointer(), PointerFromTransactionDetails(i))
	}
}

func (ptr *TransactionList) __transactions_newList() unsafe.Pointer {
	return C.TransactionList742340___transactions_newList(ptr.Pointer())
}

func (ptr *TransactionList) __setTransactions_transactions_atList(i int) *TransactionDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewTransactionDetailsFromPointer(C.TransactionList742340___setTransactions_transactions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __setTransactions_transactions_setList(i TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___setTransactions_transactions_setList(ptr.Pointer(), PointerFromTransactionDetails(i))
	}
}

func (ptr *TransactionList) __setTransactions_transactions_newList() unsafe.Pointer {
	return C.TransactionList742340___setTransactions_transactions_newList(ptr.Pointer())
}

func (ptr *TransactionList) __transactionsChanged_transactions_atList(i int) *TransactionDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewTransactionDetailsFromPointer(C.TransactionList742340___transactionsChanged_transactions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionList) __transactionsChanged_transactions_setList(i TransactionDetails_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340___transactionsChanged_transactions_setList(ptr.Pointer(), PointerFromTransactionDetails(i))
	}
}

func (ptr *TransactionList) __transactionsChanged_transactions_newList() unsafe.Pointer {
	return C.TransactionList742340___transactionsChanged_transactions_newList(ptr.Pointer())
}

func (ptr *TransactionList) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList742340_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____roles_keyList_newList() unsafe.Pointer {
	return C.TransactionList742340_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList742340_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.TransactionList742340_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *TransactionList) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList742340_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *TransactionList) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *TransactionList) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.TransactionList742340_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewTransactionList(parent std_core.QObject_ITF) *TransactionList {
	tmpValue := NewTransactionListFromPointer(C.TransactionList742340_NewTransactionList(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTransactionList742340_DestroyTransactionList
func callbackTransactionList742340_DestroyTransactionList(ptr unsafe.Pointer) {
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
		C.TransactionList742340_DestroyTransactionList(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *TransactionList) DestroyTransactionListDefault() {
	if ptr.Pointer() != nil {
		C.TransactionList742340_DestroyTransactionListDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTransactionList742340_DropMimeData
func callbackTransactionList742340_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList742340_Flags
func callbackTransactionList742340_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewTransactionListFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TransactionList) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.TransactionList742340_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackTransactionList742340_Index
func callbackTransactionList742340_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewTransactionListFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *TransactionList) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList742340_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList742340_Sibling
func callbackTransactionList742340_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewTransactionListFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *TransactionList) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList742340_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList742340_Buddy
func callbackTransactionList742340_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewTransactionListFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TransactionList) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList742340_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList742340_CanDropMimeData
func callbackTransactionList742340_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList742340_CanFetchMore
func callbackTransactionList742340_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList742340_ColumnCount
func callbackTransactionList742340_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewTransactionListFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *TransactionList) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList742340_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackTransactionList742340_ColumnsAboutToBeInserted
func callbackTransactionList742340_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList742340_ColumnsAboutToBeMoved
func callbackTransactionList742340_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackTransactionList742340_ColumnsAboutToBeRemoved
func callbackTransactionList742340_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList742340_ColumnsInserted
func callbackTransactionList742340_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList742340_ColumnsMoved
func callbackTransactionList742340_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackTransactionList742340_ColumnsRemoved
func callbackTransactionList742340_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList742340_Data
func callbackTransactionList742340_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewTransactionListFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *TransactionList) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TransactionList742340_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList742340_DataChanged
func callbackTransactionList742340_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
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

//export callbackTransactionList742340_FetchMore
func callbackTransactionList742340_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewTransactionListFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *TransactionList) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackTransactionList742340_HasChildren
func callbackTransactionList742340_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList742340_HeaderData
func callbackTransactionList742340_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewTransactionListFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *TransactionList) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.TransactionList742340_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList742340_HeaderDataChanged
func callbackTransactionList742340_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList742340_InsertColumns
func callbackTransactionList742340_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList742340_InsertRows
func callbackTransactionList742340_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList742340_ItemData
func callbackTransactionList742340_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
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
		}(C.TransactionList742340_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackTransactionList742340_LayoutAboutToBeChanged
func callbackTransactionList742340_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
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

//export callbackTransactionList742340_LayoutChanged
func callbackTransactionList742340_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
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

//export callbackTransactionList742340_Match
func callbackTransactionList742340_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
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
		}(C.TransactionList742340_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackTransactionList742340_MimeData
func callbackTransactionList742340_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
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
		tmpValue := std_core.NewQMimeDataFromPointer(C.TransactionList742340_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
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

//export callbackTransactionList742340_MimeTypes
func callbackTransactionList742340_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewTransactionListFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *TransactionList) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.TransactionList742340_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackTransactionList742340_ModelAboutToBeReset
func callbackTransactionList742340_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackTransactionList742340_ModelReset
func callbackTransactionList742340_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackTransactionList742340_MoveColumns
func callbackTransactionList742340_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *TransactionList) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackTransactionList742340_MoveRows
func callbackTransactionList742340_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *TransactionList) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackTransactionList742340_Parent
func callbackTransactionList742340_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewTransactionListFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TransactionList) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.TransactionList742340_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList742340_RemoveColumns
func callbackTransactionList742340_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList742340_RemoveRows
func callbackTransactionList742340_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *TransactionList) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackTransactionList742340_ResetInternalData
func callbackTransactionList742340_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionListFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *TransactionList) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.TransactionList742340_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackTransactionList742340_Revert
func callbackTransactionList742340_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionListFromPointer(ptr).RevertDefault()
	}
}

func (ptr *TransactionList) RevertDefault() {
	if ptr.Pointer() != nil {
		C.TransactionList742340_RevertDefault(ptr.Pointer())
	}
}

//export callbackTransactionList742340_RoleNames
func callbackTransactionList742340_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
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
		}(C.TransactionList742340_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackTransactionList742340_RowCount
func callbackTransactionList742340_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewTransactionListFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *TransactionList) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionList742340_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackTransactionList742340_RowsAboutToBeInserted
func callbackTransactionList742340_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackTransactionList742340_RowsAboutToBeMoved
func callbackTransactionList742340_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackTransactionList742340_RowsAboutToBeRemoved
func callbackTransactionList742340_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList742340_RowsInserted
func callbackTransactionList742340_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList742340_RowsMoved
func callbackTransactionList742340_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackTransactionList742340_RowsRemoved
func callbackTransactionList742340_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackTransactionList742340_SetData
func callbackTransactionList742340_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *TransactionList) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackTransactionList742340_SetHeaderData
func callbackTransactionList742340_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *TransactionList) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackTransactionList742340_SetItemData
func callbackTransactionList742340_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
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
		return int8(C.TransactionList742340_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewTransactionListFromPointer(NewTransactionListFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackTransactionList742340_Sort
func callbackTransactionList742340_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewTransactionListFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *TransactionList) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackTransactionList742340_Span
func callbackTransactionList742340_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewTransactionListFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *TransactionList) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.TransactionList742340_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackTransactionList742340_Submit
func callbackTransactionList742340_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).SubmitDefault())))
}

func (ptr *TransactionList) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackTransactionList742340_SupportedDragActions
func callbackTransactionList742340_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewTransactionListFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *TransactionList) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.TransactionList742340_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackTransactionList742340_SupportedDropActions
func callbackTransactionList742340_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewTransactionListFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *TransactionList) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.TransactionList742340_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackTransactionList742340_ChildEvent
func callbackTransactionList742340_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTransactionListFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *TransactionList) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTransactionList742340_ConnectNotify
func callbackTransactionList742340_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTransactionListFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TransactionList) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTransactionList742340_CustomEvent
func callbackTransactionList742340_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewTransactionListFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *TransactionList) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTransactionList742340_DeleteLater
func callbackTransactionList742340_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionListFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *TransactionList) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.TransactionList742340_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTransactionList742340_Destroyed
func callbackTransactionList742340_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackTransactionList742340_DisconnectNotify
func callbackTransactionList742340_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTransactionListFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TransactionList) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTransactionList742340_Event
func callbackTransactionList742340_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *TransactionList) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackTransactionList742340_EventFilter
func callbackTransactionList742340_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionListFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *TransactionList) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionList742340_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackTransactionList742340_ObjectNameChanged
func callbackTransactionList742340_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackTransactionList742340_TimerEvent
func callbackTransactionList742340_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTransactionListFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *TransactionList) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionList742340_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type AddressDetails_ITF interface {
	std_core.QObject_ITF
	AddressDetails_PTR() *AddressDetails
}

func (ptr *AddressDetails) AddressDetails_PTR() *AddressDetails {
	return ptr
}

func (ptr *AddressDetails) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *AddressDetails) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromAddressDetails(ptr AddressDetails_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AddressDetails_PTR().Pointer()
	}
	return nil
}

func NewAddressDetailsFromPointer(ptr unsafe.Pointer) (n *AddressDetails) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(AddressDetails)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *AddressDetails:
			n = deduced

		case *std_core.QObject:
			n = &AddressDetails{QObject: *deduced}

		default:
			n = new(AddressDetails)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackAddressDetails742340_Constructor
func callbackAddressDetails742340_Constructor(ptr unsafe.Pointer) {
	this := NewAddressDetailsFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackAddressDetails742340_Address
func callbackAddressDetails742340_Address(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "address"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewAddressDetailsFromPointer(ptr).AddressDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *AddressDetails) ConnectAddress(f func() string) {
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

func (ptr *AddressDetails) DisconnectAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "address")
	}
}

func (ptr *AddressDetails) Address() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.AddressDetails742340_Address(ptr.Pointer()))
	}
	return ""
}

func (ptr *AddressDetails) AddressDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.AddressDetails742340_AddressDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackAddressDetails742340_SetAddress
func callbackAddressDetails742340_SetAddress(ptr unsafe.Pointer, address C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setAddress"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(address))
	} else {
		NewAddressDetailsFromPointer(ptr).SetAddressDefault(cGoUnpackString(address))
	}
}

func (ptr *AddressDetails) ConnectSetAddress(f func(address string)) {
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

func (ptr *AddressDetails) DisconnectSetAddress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddress")
	}
}

func (ptr *AddressDetails) SetAddress(address string) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		C.AddressDetails742340_SetAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))})
	}
}

func (ptr *AddressDetails) SetAddressDefault(address string) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		C.AddressDetails742340_SetAddressDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))})
	}
}

//export callbackAddressDetails742340_AddressChanged
func callbackAddressDetails742340_AddressChanged(ptr unsafe.Pointer, address C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addressChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(address))
	}

}

func (ptr *AddressDetails) ConnectAddressChanged(f func(address string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressChanged") {
			C.AddressDetails742340_ConnectAddressChanged(ptr.Pointer())
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

func (ptr *AddressDetails) DisconnectAddressChanged() {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_DisconnectAddressChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressChanged")
	}
}

func (ptr *AddressDetails) AddressChanged(address string) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		C.AddressDetails742340_AddressChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))})
	}
}

//export callbackAddressDetails742340_AddressSky
func callbackAddressDetails742340_AddressSky(ptr unsafe.Pointer) C.float {
	if signal := qt.GetSignal(ptr, "addressSky"); signal != nil {
		return C.float((*(*func() float32)(signal))())
	}

	return C.float(NewAddressDetailsFromPointer(ptr).AddressSkyDefault())
}

func (ptr *AddressDetails) ConnectAddressSky(f func() float32) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addressSky"); signal != nil {
			f := func() float32 {
				(*(*func() float32)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "addressSky", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressSky", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectAddressSky() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addressSky")
	}
}

func (ptr *AddressDetails) AddressSky() float32 {
	if ptr.Pointer() != nil {
		return float32(C.AddressDetails742340_AddressSky(ptr.Pointer()))
	}
	return 0
}

func (ptr *AddressDetails) AddressSkyDefault() float32 {
	if ptr.Pointer() != nil {
		return float32(C.AddressDetails742340_AddressSkyDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackAddressDetails742340_SetAddressSky
func callbackAddressDetails742340_SetAddressSky(ptr unsafe.Pointer, addressSky C.float) {
	if signal := qt.GetSignal(ptr, "setAddressSky"); signal != nil {
		(*(*func(float32))(signal))(float32(addressSky))
	} else {
		NewAddressDetailsFromPointer(ptr).SetAddressSkyDefault(float32(addressSky))
	}
}

func (ptr *AddressDetails) ConnectSetAddressSky(f func(addressSky float32)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddressSky"); signal != nil {
			f := func(addressSky float32) {
				(*(*func(float32))(signal))(addressSky)
				f(addressSky)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAddressSky", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAddressSky", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectSetAddressSky() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddressSky")
	}
}

func (ptr *AddressDetails) SetAddressSky(addressSky float32) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_SetAddressSky(ptr.Pointer(), C.float(addressSky))
	}
}

func (ptr *AddressDetails) SetAddressSkyDefault(addressSky float32) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_SetAddressSkyDefault(ptr.Pointer(), C.float(addressSky))
	}
}

//export callbackAddressDetails742340_AddressSkyChanged
func callbackAddressDetails742340_AddressSkyChanged(ptr unsafe.Pointer, addressSky C.float) {
	if signal := qt.GetSignal(ptr, "addressSkyChanged"); signal != nil {
		(*(*func(float32))(signal))(float32(addressSky))
	}

}

func (ptr *AddressDetails) ConnectAddressSkyChanged(f func(addressSky float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressSkyChanged") {
			C.AddressDetails742340_ConnectAddressSkyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressSkyChanged"); signal != nil {
			f := func(addressSky float32) {
				(*(*func(float32))(signal))(addressSky)
				f(addressSky)
			}
			qt.ConnectSignal(ptr.Pointer(), "addressSkyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressSkyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectAddressSkyChanged() {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_DisconnectAddressSkyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressSkyChanged")
	}
}

func (ptr *AddressDetails) AddressSkyChanged(addressSky float32) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_AddressSkyChanged(ptr.Pointer(), C.float(addressSky))
	}
}

//export callbackAddressDetails742340_AddressCoinHours
func callbackAddressDetails742340_AddressCoinHours(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "addressCoinHours"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewAddressDetailsFromPointer(ptr).AddressCoinHoursDefault()))
}

func (ptr *AddressDetails) ConnectAddressCoinHours(f func() int) {
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

func (ptr *AddressDetails) DisconnectAddressCoinHours() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addressCoinHours")
	}
}

func (ptr *AddressDetails) AddressCoinHours() int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressDetails742340_AddressCoinHours(ptr.Pointer())))
	}
	return 0
}

func (ptr *AddressDetails) AddressCoinHoursDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressDetails742340_AddressCoinHoursDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackAddressDetails742340_SetAddressCoinHours
func callbackAddressDetails742340_SetAddressCoinHours(ptr unsafe.Pointer, addressCoinHours C.int) {
	if signal := qt.GetSignal(ptr, "setAddressCoinHours"); signal != nil {
		(*(*func(int))(signal))(int(int32(addressCoinHours)))
	} else {
		NewAddressDetailsFromPointer(ptr).SetAddressCoinHoursDefault(int(int32(addressCoinHours)))
	}
}

func (ptr *AddressDetails) ConnectSetAddressCoinHours(f func(addressCoinHours int)) {
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

func (ptr *AddressDetails) DisconnectSetAddressCoinHours() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddressCoinHours")
	}
}

func (ptr *AddressDetails) SetAddressCoinHours(addressCoinHours int) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_SetAddressCoinHours(ptr.Pointer(), C.int(int32(addressCoinHours)))
	}
}

func (ptr *AddressDetails) SetAddressCoinHoursDefault(addressCoinHours int) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_SetAddressCoinHoursDefault(ptr.Pointer(), C.int(int32(addressCoinHours)))
	}
}

//export callbackAddressDetails742340_AddressCoinHoursChanged
func callbackAddressDetails742340_AddressCoinHoursChanged(ptr unsafe.Pointer, addressCoinHours C.int) {
	if signal := qt.GetSignal(ptr, "addressCoinHoursChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(addressCoinHours)))
	}

}

func (ptr *AddressDetails) ConnectAddressCoinHoursChanged(f func(addressCoinHours int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressCoinHoursChanged") {
			C.AddressDetails742340_ConnectAddressCoinHoursChanged(ptr.Pointer())
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

func (ptr *AddressDetails) DisconnectAddressCoinHoursChanged() {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_DisconnectAddressCoinHoursChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressCoinHoursChanged")
	}
}

func (ptr *AddressDetails) AddressCoinHoursChanged(addressCoinHours int) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_AddressCoinHoursChanged(ptr.Pointer(), C.int(int32(addressCoinHours)))
	}
}

func AddressDetails_QRegisterMetaType() int {
	return int(int32(C.AddressDetails742340_AddressDetails742340_QRegisterMetaType()))
}

func (ptr *AddressDetails) QRegisterMetaType() int {
	return int(int32(C.AddressDetails742340_AddressDetails742340_QRegisterMetaType()))
}

func AddressDetails_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressDetails742340_AddressDetails742340_QRegisterMetaType2(typeNameC)))
}

func (ptr *AddressDetails) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressDetails742340_AddressDetails742340_QRegisterMetaType2(typeNameC)))
}

func AddressDetails_QmlRegisterType() int {
	return int(int32(C.AddressDetails742340_AddressDetails742340_QmlRegisterType()))
}

func (ptr *AddressDetails) QmlRegisterType() int {
	return int(int32(C.AddressDetails742340_AddressDetails742340_QmlRegisterType()))
}

func AddressDetails_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.AddressDetails742340_AddressDetails742340_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *AddressDetails) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.AddressDetails742340_AddressDetails742340_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *AddressDetails) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressDetails742340___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressDetails) __children_newList() unsafe.Pointer {
	return C.AddressDetails742340___children_newList(ptr.Pointer())
}

func (ptr *AddressDetails) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressDetails742340___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressDetails) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.AddressDetails742340___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *AddressDetails) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressDetails742340___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressDetails) __findChildren_newList() unsafe.Pointer {
	return C.AddressDetails742340___findChildren_newList(ptr.Pointer())
}

func (ptr *AddressDetails) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressDetails742340___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressDetails) __findChildren_newList3() unsafe.Pointer {
	return C.AddressDetails742340___findChildren_newList3(ptr.Pointer())
}

func (ptr *AddressDetails) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressDetails742340___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressDetails) __qFindChildren_newList2() unsafe.Pointer {
	return C.AddressDetails742340___qFindChildren_newList2(ptr.Pointer())
}

func NewAddressDetails(parent std_core.QObject_ITF) *AddressDetails {
	tmpValue := NewAddressDetailsFromPointer(C.AddressDetails742340_NewAddressDetails(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackAddressDetails742340_DestroyAddressDetails
func callbackAddressDetails742340_DestroyAddressDetails(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~AddressDetails"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressDetailsFromPointer(ptr).DestroyAddressDetailsDefault()
	}
}

func (ptr *AddressDetails) ConnectDestroyAddressDetails(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~AddressDetails"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~AddressDetails", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~AddressDetails", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressDetails) DisconnectDestroyAddressDetails() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~AddressDetails")
	}
}

func (ptr *AddressDetails) DestroyAddressDetails() {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_DestroyAddressDetails(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *AddressDetails) DestroyAddressDetailsDefault() {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_DestroyAddressDetailsDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressDetails742340_ChildEvent
func callbackAddressDetails742340_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewAddressDetailsFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *AddressDetails) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackAddressDetails742340_ConnectNotify
func callbackAddressDetails742340_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressDetailsFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressDetails) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressDetails742340_CustomEvent
func callbackAddressDetails742340_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewAddressDetailsFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *AddressDetails) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackAddressDetails742340_DeleteLater
func callbackAddressDetails742340_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressDetailsFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *AddressDetails) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressDetails742340_Destroyed
func callbackAddressDetails742340_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackAddressDetails742340_DisconnectNotify
func callbackAddressDetails742340_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressDetailsFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressDetails) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressDetails742340_Event
func callbackAddressDetails742340_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressDetailsFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *AddressDetails) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressDetails742340_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackAddressDetails742340_EventFilter
func callbackAddressDetails742340_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressDetailsFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *AddressDetails) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressDetails742340_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackAddressDetails742340_ObjectNameChanged
func callbackAddressDetails742340_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackAddressDetails742340_TimerEvent
func callbackAddressDetails742340_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewAddressDetailsFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *AddressDetails) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetails742340_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type AddressList_ITF interface {
	std_core.QAbstractListModel_ITF
	AddressList_PTR() *AddressList
}

func (ptr *AddressList) AddressList_PTR() *AddressList {
	return ptr
}

func (ptr *AddressList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *AddressList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromAddressList(ptr AddressList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.AddressList_PTR().Pointer()
	}
	return nil
}

func NewAddressListFromPointer(ptr unsafe.Pointer) (n *AddressList) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(AddressList)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *AddressList:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &AddressList{QAbstractListModel: *deduced}

		default:
			n = new(AddressList)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *AddressList) Init() { this.init() }

//export callbackAddressList742340_Constructor
func callbackAddressList742340_Constructor(ptr unsafe.Pointer) {
	this := NewAddressListFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectAddAddress(this.addAddress)
	this.ConnectRemoveAddress(this.removeAddress)
	this.init()
}

//export callbackAddressList742340_AddAddress
func callbackAddressList742340_AddAddress(ptr unsafe.Pointer, transaction unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addAddress"); signal != nil {
		(*(*func(*AddressDetails))(signal))(NewAddressDetailsFromPointer(transaction))
	}

}

func (ptr *AddressList) ConnectAddAddress(f func(transaction *AddressDetails)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addAddress") {
			C.AddressList742340_ConnectAddAddress(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addAddress"); signal != nil {
			f := func(transaction *AddressDetails) {
				(*(*func(*AddressDetails))(signal))(transaction)
				f(transaction)
			}
			qt.ConnectSignal(ptr.Pointer(), "addAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectAddAddress() {
	if ptr.Pointer() != nil {
		C.AddressList742340_DisconnectAddAddress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addAddress")
	}
}

func (ptr *AddressList) AddAddress(transaction AddressDetails_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340_AddAddress(ptr.Pointer(), PointerFromAddressDetails(transaction))
	}
}

//export callbackAddressList742340_RemoveAddress
func callbackAddressList742340_RemoveAddress(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "removeAddress"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *AddressList) ConnectRemoveAddress(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "removeAddress") {
			C.AddressList742340_ConnectRemoveAddress(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "removeAddress"); signal != nil {
			f := func(index int) {
				(*(*func(int))(signal))(index)
				f(index)
			}
			qt.ConnectSignal(ptr.Pointer(), "removeAddress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "removeAddress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectRemoveAddress() {
	if ptr.Pointer() != nil {
		C.AddressList742340_DisconnectRemoveAddress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "removeAddress")
	}
}

func (ptr *AddressList) RemoveAddress(index int) {
	if ptr.Pointer() != nil {
		C.AddressList742340_RemoveAddress(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackAddressList742340_Roles
func callbackAddressList742340_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__roles_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__roles_newList())
		for k, v := range NewAddressListFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressList) ConnectRoles(f func() map[int]*std_core.QByteArray) {
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

func (ptr *AddressList) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *AddressList) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.AddressList742340_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *AddressList) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.AddressList742340_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackAddressList742340_SetRoles
func callbackAddressList742340_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	} else {
		NewAddressListFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	}
}

func (ptr *AddressList) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
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

func (ptr *AddressList) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *AddressList) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.AddressList742340_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *AddressList) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.AddressList742340_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackAddressList742340_RolesChanged
func callbackAddressList742340_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__rolesChanged_roles_keyList() {
				out[v] = tmpList.__rolesChanged_roles_atList(v, i)
			}
			return out
		}(roles))
	}

}

func (ptr *AddressList) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.AddressList742340_ConnectRolesChanged(ptr.Pointer())
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

func (ptr *AddressList) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.AddressList742340_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *AddressList) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.AddressList742340_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackAddressList742340_Addresses
func callbackAddressList742340_Addresses(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "addresses"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__addresses_newList())
			for _, v := range (*(*func() []*AddressDetails)(signal))() {
				tmpList.__addresses_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__addresses_newList())
		for _, v := range NewAddressListFromPointer(ptr).AddressesDefault() {
			tmpList.__addresses_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressList) ConnectAddresses(f func() []*AddressDetails) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addresses"); signal != nil {
			f := func() []*AddressDetails {
				(*(*func() []*AddressDetails)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "addresses", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addresses", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectAddresses() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addresses")
	}
}

func (ptr *AddressList) Addresses() []*AddressDetails {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*AddressDetails {
			out := make([]*AddressDetails, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__addresses_atList(i)
			}
			return out
		}(C.AddressList742340_Addresses(ptr.Pointer()))
	}
	return make([]*AddressDetails, 0)
}

func (ptr *AddressList) AddressesDefault() []*AddressDetails {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*AddressDetails {
			out := make([]*AddressDetails, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__addresses_atList(i)
			}
			return out
		}(C.AddressList742340_AddressesDefault(ptr.Pointer()))
	}
	return make([]*AddressDetails, 0)
}

//export callbackAddressList742340_SetAddresses
func callbackAddressList742340_SetAddresses(ptr unsafe.Pointer, addresses C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setAddresses"); signal != nil {
		(*(*func([]*AddressDetails))(signal))(func(l C.struct_Moc_PackedList) []*AddressDetails {
			out := make([]*AddressDetails, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setAddresses_addresses_atList(i)
			}
			return out
		}(addresses))
	} else {
		NewAddressListFromPointer(ptr).SetAddressesDefault(func(l C.struct_Moc_PackedList) []*AddressDetails {
			out := make([]*AddressDetails, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setAddresses_addresses_atList(i)
			}
			return out
		}(addresses))
	}
}

func (ptr *AddressList) ConnectSetAddresses(f func(addresses []*AddressDetails)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddresses"); signal != nil {
			f := func(addresses []*AddressDetails) {
				(*(*func([]*AddressDetails))(signal))(addresses)
				f(addresses)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAddresses", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAddresses", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectSetAddresses() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddresses")
	}
}

func (ptr *AddressList) SetAddresses(addresses []*AddressDetails) {
	if ptr.Pointer() != nil {
		C.AddressList742340_SetAddresses(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setAddresses_addresses_newList())
			for _, v := range addresses {
				tmpList.__setAddresses_addresses_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *AddressList) SetAddressesDefault(addresses []*AddressDetails) {
	if ptr.Pointer() != nil {
		C.AddressList742340_SetAddressesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setAddresses_addresses_newList())
			for _, v := range addresses {
				tmpList.__setAddresses_addresses_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackAddressList742340_AddressesChanged
func callbackAddressList742340_AddressesChanged(ptr unsafe.Pointer, addresses C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "addressesChanged"); signal != nil {
		(*(*func([]*AddressDetails))(signal))(func(l C.struct_Moc_PackedList) []*AddressDetails {
			out := make([]*AddressDetails, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__addressesChanged_addresses_atList(i)
			}
			return out
		}(addresses))
	}

}

func (ptr *AddressList) ConnectAddressesChanged(f func(addresses []*AddressDetails)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressesChanged") {
			C.AddressList742340_ConnectAddressesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressesChanged"); signal != nil {
			f := func(addresses []*AddressDetails) {
				(*(*func([]*AddressDetails))(signal))(addresses)
				f(addresses)
			}
			qt.ConnectSignal(ptr.Pointer(), "addressesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectAddressesChanged() {
	if ptr.Pointer() != nil {
		C.AddressList742340_DisconnectAddressesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressesChanged")
	}
}

func (ptr *AddressList) AddressesChanged(addresses []*AddressDetails) {
	if ptr.Pointer() != nil {
		C.AddressList742340_AddressesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__addressesChanged_addresses_newList())
			for _, v := range addresses {
				tmpList.__addressesChanged_addresses_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func AddressList_QRegisterMetaType() int {
	return int(int32(C.AddressList742340_AddressList742340_QRegisterMetaType()))
}

func (ptr *AddressList) QRegisterMetaType() int {
	return int(int32(C.AddressList742340_AddressList742340_QRegisterMetaType()))
}

func AddressList_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressList742340_AddressList742340_QRegisterMetaType2(typeNameC)))
}

func (ptr *AddressList) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressList742340_AddressList742340_QRegisterMetaType2(typeNameC)))
}

func AddressList_QmlRegisterType() int {
	return int(int32(C.AddressList742340_AddressList742340_QmlRegisterType()))
}

func (ptr *AddressList) QmlRegisterType() int {
	return int(int32(C.AddressList742340_AddressList742340_QmlRegisterType()))
}

func AddressList_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.AddressList742340_AddressList742340_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *AddressList) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.AddressList742340_AddressList742340_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *AddressList) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList742340_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList742340_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____itemData_keyList_newList() unsafe.Pointer {
	return C.AddressList742340_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList742340_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList742340_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.AddressList742340_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList742340_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList742340_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.AddressList742340_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList742340___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.AddressList742340___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *AddressList) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList742340___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.AddressList742340___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *AddressList) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList742340___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList742340___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) __dataChanged_roles_newList() unsafe.Pointer {
	return C.AddressList742340___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *AddressList) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressList742340___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *AddressList) __itemData_newList() unsafe.Pointer {
	return C.AddressList742340___itemData_newList(ptr.Pointer())
}

func (ptr *AddressList) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.AddressList742340___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.AddressList742340___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *AddressList) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.AddressList742340___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *AddressList) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.AddressList742340___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *AddressList) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.AddressList742340___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *AddressList) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList742340___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __match_newList() unsafe.Pointer {
	return C.AddressList742340___match_newList(ptr.Pointer())
}

func (ptr *AddressList) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList742340___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __mimeData_indexes_newList() unsafe.Pointer {
	return C.AddressList742340___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *AddressList) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList742340___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __persistentIndexList_newList() unsafe.Pointer {
	return C.AddressList742340___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *AddressList) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressList742340___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __roleNames_newList() unsafe.Pointer {
	return C.AddressList742340___roleNames_newList(ptr.Pointer())
}

func (ptr *AddressList) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.AddressList742340___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressList742340___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *AddressList) __setItemData_roles_newList() unsafe.Pointer {
	return C.AddressList742340___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *AddressList) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.AddressList742340___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList742340_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList742340_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.AddressList742340_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList742340_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList742340_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.AddressList742340_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressList742340___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressList) __children_newList() unsafe.Pointer {
	return C.AddressList742340___children_newList(ptr.Pointer())
}

func (ptr *AddressList) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressList742340___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.AddressList742340___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *AddressList) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressList742340___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressList) __findChildren_newList() unsafe.Pointer {
	return C.AddressList742340___findChildren_newList(ptr.Pointer())
}

func (ptr *AddressList) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressList742340___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressList) __findChildren_newList3() unsafe.Pointer {
	return C.AddressList742340___findChildren_newList3(ptr.Pointer())
}

func (ptr *AddressList) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressList742340___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressList) __qFindChildren_newList2() unsafe.Pointer {
	return C.AddressList742340___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *AddressList) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressList742340___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __roles_newList() unsafe.Pointer {
	return C.AddressList742340___roles_newList(ptr.Pointer())
}

func (ptr *AddressList) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roles_keyList_atList(i)
			}
			return out
		}(C.AddressList742340___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressList742340___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __setRoles_roles_newList() unsafe.Pointer {
	return C.AddressList742340___setRoles_roles_newList(ptr.Pointer())
}

func (ptr *AddressList) __setRoles_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoles_roles_keyList_atList(i)
			}
			return out
		}(C.AddressList742340___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressList742340___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.AddressList742340___rolesChanged_roles_newList(ptr.Pointer())
}

func (ptr *AddressList) __rolesChanged_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____rolesChanged_roles_keyList_atList(i)
			}
			return out
		}(C.AddressList742340___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __addresses_atList(i int) *AddressDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressDetailsFromPointer(C.AddressList742340___addresses_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __addresses_setList(i AddressDetails_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___addresses_setList(ptr.Pointer(), PointerFromAddressDetails(i))
	}
}

func (ptr *AddressList) __addresses_newList() unsafe.Pointer {
	return C.AddressList742340___addresses_newList(ptr.Pointer())
}

func (ptr *AddressList) __setAddresses_addresses_atList(i int) *AddressDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressDetailsFromPointer(C.AddressList742340___setAddresses_addresses_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __setAddresses_addresses_setList(i AddressDetails_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___setAddresses_addresses_setList(ptr.Pointer(), PointerFromAddressDetails(i))
	}
}

func (ptr *AddressList) __setAddresses_addresses_newList() unsafe.Pointer {
	return C.AddressList742340___setAddresses_addresses_newList(ptr.Pointer())
}

func (ptr *AddressList) __addressesChanged_addresses_atList(i int) *AddressDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressDetailsFromPointer(C.AddressList742340___addressesChanged_addresses_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __addressesChanged_addresses_setList(i AddressDetails_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340___addressesChanged_addresses_setList(ptr.Pointer(), PointerFromAddressDetails(i))
	}
}

func (ptr *AddressList) __addressesChanged_addresses_newList() unsafe.Pointer {
	return C.AddressList742340___addressesChanged_addresses_newList(ptr.Pointer())
}

func (ptr *AddressList) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList742340_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList742340_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____roles_keyList_newList() unsafe.Pointer {
	return C.AddressList742340_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList742340_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList742340_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.AddressList742340_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList742340_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressList742340_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.AddressList742340_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewAddressList(parent std_core.QObject_ITF) *AddressList {
	tmpValue := NewAddressListFromPointer(C.AddressList742340_NewAddressList(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackAddressList742340_DestroyAddressList
func callbackAddressList742340_DestroyAddressList(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~AddressList"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressListFromPointer(ptr).DestroyAddressListDefault()
	}
}

func (ptr *AddressList) ConnectDestroyAddressList(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~AddressList"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~AddressList", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~AddressList", unsafe.Pointer(&f))
		}
	}
}

func (ptr *AddressList) DisconnectDestroyAddressList() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~AddressList")
	}
}

func (ptr *AddressList) DestroyAddressList() {
	if ptr.Pointer() != nil {
		C.AddressList742340_DestroyAddressList(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *AddressList) DestroyAddressListDefault() {
	if ptr.Pointer() != nil {
		C.AddressList742340_DestroyAddressListDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressList742340_DropMimeData
func callbackAddressList742340_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList742340_Flags
func callbackAddressList742340_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewAddressListFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressList) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.AddressList742340_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackAddressList742340_Index
func callbackAddressList742340_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewAddressListFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *AddressList) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList742340_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressList742340_Sibling
func callbackAddressList742340_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewAddressListFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *AddressList) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList742340_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressList742340_Buddy
func callbackAddressList742340_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewAddressListFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressList) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList742340_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressList742340_CanDropMimeData
func callbackAddressList742340_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList742340_CanFetchMore
func callbackAddressList742340_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList742340_ColumnCount
func callbackAddressList742340_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewAddressListFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *AddressList) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList742340_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackAddressList742340_ColumnsAboutToBeInserted
func callbackAddressList742340_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList742340_ColumnsAboutToBeMoved
func callbackAddressList742340_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackAddressList742340_ColumnsAboutToBeRemoved
func callbackAddressList742340_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList742340_ColumnsInserted
func callbackAddressList742340_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList742340_ColumnsMoved
func callbackAddressList742340_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackAddressList742340_ColumnsRemoved
func callbackAddressList742340_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList742340_Data
func callbackAddressList742340_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewAddressListFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *AddressList) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressList742340_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackAddressList742340_DataChanged
func callbackAddressList742340_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(signal))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbackAddressList742340_FetchMore
func callbackAddressList742340_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewAddressListFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *AddressList) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackAddressList742340_HasChildren
func callbackAddressList742340_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList742340_HeaderData
func callbackAddressList742340_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewAddressListFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *AddressList) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressList742340_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackAddressList742340_HeaderDataChanged
func callbackAddressList742340_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList742340_InsertColumns
func callbackAddressList742340_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList742340_InsertRows
func callbackAddressList742340_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList742340_ItemData
func callbackAddressList742340_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*std_core.QModelIndex) map[int]*std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__itemData_newList())
		for k, v := range NewAddressListFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressList) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.AddressList742340_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackAddressList742340_LayoutAboutToBeChanged
func callbackAddressList742340_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackAddressList742340_LayoutChanged
func callbackAddressList742340_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbackAddressList742340_Match
func callbackAddressList742340_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__match_newList())
		for _, v := range NewAddressListFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressList) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.AddressList742340_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackAddressList742340_MimeData
func callbackAddressList742340_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func([]*std_core.QModelIndex) *std_core.QMimeData)(signal))(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewAddressListFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewAddressListFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *AddressList) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.AddressList742340_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__mimeData_indexes_newList())
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

//export callbackAddressList742340_MimeTypes
func callbackAddressList742340_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewAddressListFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *AddressList) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.AddressList742340_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackAddressList742340_ModelAboutToBeReset
func callbackAddressList742340_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackAddressList742340_ModelReset
func callbackAddressList742340_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackAddressList742340_MoveColumns
func callbackAddressList742340_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *AddressList) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackAddressList742340_MoveRows
func callbackAddressList742340_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *AddressList) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackAddressList742340_Parent
func callbackAddressList742340_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewAddressListFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressList) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressList742340_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressList742340_RemoveColumns
func callbackAddressList742340_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList742340_RemoveRows
func callbackAddressList742340_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressList742340_ResetInternalData
func callbackAddressList742340_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressListFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *AddressList) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.AddressList742340_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackAddressList742340_Revert
func callbackAddressList742340_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressListFromPointer(ptr).RevertDefault()
	}
}

func (ptr *AddressList) RevertDefault() {
	if ptr.Pointer() != nil {
		C.AddressList742340_RevertDefault(ptr.Pointer())
	}
}

//export callbackAddressList742340_RoleNames
func callbackAddressList742340_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__roleNames_newList())
		for k, v := range NewAddressListFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *AddressList) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.AddressList742340_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackAddressList742340_RowCount
func callbackAddressList742340_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewAddressListFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *AddressList) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressList742340_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackAddressList742340_RowsAboutToBeInserted
func callbackAddressList742340_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackAddressList742340_RowsAboutToBeMoved
func callbackAddressList742340_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackAddressList742340_RowsAboutToBeRemoved
func callbackAddressList742340_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList742340_RowsInserted
func callbackAddressList742340_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList742340_RowsMoved
func callbackAddressList742340_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackAddressList742340_RowsRemoved
func callbackAddressList742340_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressList742340_SetData
func callbackAddressList742340_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *AddressList) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackAddressList742340_SetHeaderData
func callbackAddressList742340_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *AddressList) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackAddressList742340_SetItemData
func callbackAddressList742340_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(signal))(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewAddressListFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewAddressListFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *AddressList) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackAddressList742340_Sort
func callbackAddressList742340_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewAddressListFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *AddressList) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.AddressList742340_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackAddressList742340_Span
func callbackAddressList742340_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewAddressListFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressList) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.AddressList742340_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackAddressList742340_Submit
func callbackAddressList742340_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).SubmitDefault())))
}

func (ptr *AddressList) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackAddressList742340_SupportedDragActions
func callbackAddressList742340_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewAddressListFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *AddressList) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.AddressList742340_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackAddressList742340_SupportedDropActions
func callbackAddressList742340_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewAddressListFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *AddressList) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.AddressList742340_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackAddressList742340_ChildEvent
func callbackAddressList742340_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewAddressListFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *AddressList) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackAddressList742340_ConnectNotify
func callbackAddressList742340_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressListFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressList) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressList742340_CustomEvent
func callbackAddressList742340_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewAddressListFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *AddressList) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackAddressList742340_DeleteLater
func callbackAddressList742340_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressListFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *AddressList) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.AddressList742340_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressList742340_Destroyed
func callbackAddressList742340_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackAddressList742340_DisconnectNotify
func callbackAddressList742340_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressListFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressList) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressList742340_Event
func callbackAddressList742340_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *AddressList) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackAddressList742340_EventFilter
func callbackAddressList742340_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *AddressList) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressList742340_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackAddressList742340_ObjectNameChanged
func callbackAddressList742340_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackAddressList742340_TimerEvent
func callbackAddressList742340_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewAddressListFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *AddressList) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressList742340_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
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

//export callbackTransactionDetails742340_Constructor
func callbackTransactionDetails742340_Constructor(ptr unsafe.Pointer) {
	this := NewTransactionDetailsFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackTransactionDetails742340_Date
func callbackTransactionDetails742340_Date(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "date"); signal != nil {
		return std_core.PointerFromQDateTime((*(*func() *std_core.QDateTime)(signal))())
	}

	return std_core.PointerFromQDateTime(NewTransactionDetailsFromPointer(ptr).DateDefault())
}

func (ptr *TransactionDetails) ConnectDate(f func() *std_core.QDateTime) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "date"); signal != nil {
			f := func() *std_core.QDateTime {
				(*(*func() *std_core.QDateTime)(signal))()
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

func (ptr *TransactionDetails) Date() *std_core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQDateTimeFromPointer(C.TransactionDetails742340_Date(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) DateDefault() *std_core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQDateTimeFromPointer(C.TransactionDetails742340_DateDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

//export callbackTransactionDetails742340_SetDate
func callbackTransactionDetails742340_SetDate(ptr unsafe.Pointer, date unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setDate"); signal != nil {
		(*(*func(*std_core.QDateTime))(signal))(std_core.NewQDateTimeFromPointer(date))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetDateDefault(std_core.NewQDateTimeFromPointer(date))
	}
}

func (ptr *TransactionDetails) ConnectSetDate(f func(date *std_core.QDateTime)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setDate"); signal != nil {
			f := func(date *std_core.QDateTime) {
				(*(*func(*std_core.QDateTime))(signal))(date)
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

func (ptr *TransactionDetails) SetDate(date std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_SetDate(ptr.Pointer(), std_core.PointerFromQDateTime(date))
	}
}

func (ptr *TransactionDetails) SetDateDefault(date std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_SetDateDefault(ptr.Pointer(), std_core.PointerFromQDateTime(date))
	}
}

//export callbackTransactionDetails742340_DateChanged
func callbackTransactionDetails742340_DateChanged(ptr unsafe.Pointer, date unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dateChanged"); signal != nil {
		(*(*func(*std_core.QDateTime))(signal))(std_core.NewQDateTimeFromPointer(date))
	}

}

func (ptr *TransactionDetails) ConnectDateChanged(f func(date *std_core.QDateTime)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dateChanged") {
			C.TransactionDetails742340_ConnectDateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dateChanged"); signal != nil {
			f := func(date *std_core.QDateTime) {
				(*(*func(*std_core.QDateTime))(signal))(date)
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
		C.TransactionDetails742340_DisconnectDateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dateChanged")
	}
}

func (ptr *TransactionDetails) DateChanged(date std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_DateChanged(ptr.Pointer(), std_core.PointerFromQDateTime(date))
	}
}

//export callbackTransactionDetails742340_Status
func callbackTransactionDetails742340_Status(ptr unsafe.Pointer) C.int {
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
		return int(int32(C.TransactionDetails742340_Status(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) StatusDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails742340_StatusDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails742340_SetStatus
func callbackTransactionDetails742340_SetStatus(ptr unsafe.Pointer, status C.int) {
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
		C.TransactionDetails742340_SetStatus(ptr.Pointer(), C.int(int32(status)))
	}
}

func (ptr *TransactionDetails) SetStatusDefault(status int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_SetStatusDefault(ptr.Pointer(), C.int(int32(status)))
	}
}

//export callbackTransactionDetails742340_StatusChanged
func callbackTransactionDetails742340_StatusChanged(ptr unsafe.Pointer, status C.int) {
	if signal := qt.GetSignal(ptr, "statusChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(status)))
	}

}

func (ptr *TransactionDetails) ConnectStatusChanged(f func(status int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "statusChanged") {
			C.TransactionDetails742340_ConnectStatusChanged(ptr.Pointer())
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
		C.TransactionDetails742340_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "statusChanged")
	}
}

func (ptr *TransactionDetails) StatusChanged(status int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_StatusChanged(ptr.Pointer(), C.int(int32(status)))
	}
}

//export callbackTransactionDetails742340_Type
func callbackTransactionDetails742340_Type(ptr unsafe.Pointer) C.int {
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
		return int(int32(C.TransactionDetails742340_Type(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) TypeDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails742340_TypeDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails742340_SetType
func callbackTransactionDetails742340_SetType(ptr unsafe.Pointer, ty C.int) {
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
		C.TransactionDetails742340_SetType(ptr.Pointer(), C.int(int32(ty)))
	}
}

func (ptr *TransactionDetails) SetTypeDefault(ty int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_SetTypeDefault(ptr.Pointer(), C.int(int32(ty)))
	}
}

//export callbackTransactionDetails742340_TypeChanged
func callbackTransactionDetails742340_TypeChanged(ptr unsafe.Pointer, ty C.int) {
	if signal := qt.GetSignal(ptr, "typeChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(ty)))
	}

}

func (ptr *TransactionDetails) ConnectTypeChanged(f func(ty int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "typeChanged") {
			C.TransactionDetails742340_ConnectTypeChanged(ptr.Pointer())
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
		C.TransactionDetails742340_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "typeChanged")
	}
}

func (ptr *TransactionDetails) TypeChanged(ty int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_TypeChanged(ptr.Pointer(), C.int(int32(ty)))
	}
}

//export callbackTransactionDetails742340_Amount
func callbackTransactionDetails742340_Amount(ptr unsafe.Pointer) C.int {
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
		return int(int32(C.TransactionDetails742340_Amount(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) AmountDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails742340_AmountDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails742340_SetAmount
func callbackTransactionDetails742340_SetAmount(ptr unsafe.Pointer, amount C.int) {
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
		C.TransactionDetails742340_SetAmount(ptr.Pointer(), C.int(int32(amount)))
	}
}

func (ptr *TransactionDetails) SetAmountDefault(amount int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_SetAmountDefault(ptr.Pointer(), C.int(int32(amount)))
	}
}

//export callbackTransactionDetails742340_AmountChanged
func callbackTransactionDetails742340_AmountChanged(ptr unsafe.Pointer, amount C.int) {
	if signal := qt.GetSignal(ptr, "amountChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(amount)))
	}

}

func (ptr *TransactionDetails) ConnectAmountChanged(f func(amount int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "amountChanged") {
			C.TransactionDetails742340_ConnectAmountChanged(ptr.Pointer())
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
		C.TransactionDetails742340_DisconnectAmountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "amountChanged")
	}
}

func (ptr *TransactionDetails) AmountChanged(amount int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_AmountChanged(ptr.Pointer(), C.int(int32(amount)))
	}
}

//export callbackTransactionDetails742340_HoursReceived
func callbackTransactionDetails742340_HoursReceived(ptr unsafe.Pointer) C.int {
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
		return int(int32(C.TransactionDetails742340_HoursReceived(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) HoursReceivedDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails742340_HoursReceivedDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails742340_SetHoursReceived
func callbackTransactionDetails742340_SetHoursReceived(ptr unsafe.Pointer, hoursReceived C.int) {
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
		C.TransactionDetails742340_SetHoursReceived(ptr.Pointer(), C.int(int32(hoursReceived)))
	}
}

func (ptr *TransactionDetails) SetHoursReceivedDefault(hoursReceived int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_SetHoursReceivedDefault(ptr.Pointer(), C.int(int32(hoursReceived)))
	}
}

//export callbackTransactionDetails742340_HoursReceivedChanged
func callbackTransactionDetails742340_HoursReceivedChanged(ptr unsafe.Pointer, hoursReceived C.int) {
	if signal := qt.GetSignal(ptr, "hoursReceivedChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(hoursReceived)))
	}

}

func (ptr *TransactionDetails) ConnectHoursReceivedChanged(f func(hoursReceived int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hoursReceivedChanged") {
			C.TransactionDetails742340_ConnectHoursReceivedChanged(ptr.Pointer())
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
		C.TransactionDetails742340_DisconnectHoursReceivedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hoursReceivedChanged")
	}
}

func (ptr *TransactionDetails) HoursReceivedChanged(hoursReceived int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_HoursReceivedChanged(ptr.Pointer(), C.int(int32(hoursReceived)))
	}
}

//export callbackTransactionDetails742340_HoursBurned
func callbackTransactionDetails742340_HoursBurned(ptr unsafe.Pointer) C.int {
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
		return int(int32(C.TransactionDetails742340_HoursBurned(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) HoursBurnedDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetails742340_HoursBurnedDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetails742340_SetHoursBurned
func callbackTransactionDetails742340_SetHoursBurned(ptr unsafe.Pointer, hoursBurned C.int) {
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
		C.TransactionDetails742340_SetHoursBurned(ptr.Pointer(), C.int(int32(hoursBurned)))
	}
}

func (ptr *TransactionDetails) SetHoursBurnedDefault(hoursBurned int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_SetHoursBurnedDefault(ptr.Pointer(), C.int(int32(hoursBurned)))
	}
}

//export callbackTransactionDetails742340_HoursBurnedChanged
func callbackTransactionDetails742340_HoursBurnedChanged(ptr unsafe.Pointer, hoursBurned C.int) {
	if signal := qt.GetSignal(ptr, "hoursBurnedChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(hoursBurned)))
	}

}

func (ptr *TransactionDetails) ConnectHoursBurnedChanged(f func(hoursBurned int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hoursBurnedChanged") {
			C.TransactionDetails742340_ConnectHoursBurnedChanged(ptr.Pointer())
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
		C.TransactionDetails742340_DisconnectHoursBurnedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hoursBurnedChanged")
	}
}

func (ptr *TransactionDetails) HoursBurnedChanged(hoursBurned int) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_HoursBurnedChanged(ptr.Pointer(), C.int(int32(hoursBurned)))
	}
}

//export callbackTransactionDetails742340_TransactionID
func callbackTransactionDetails742340_TransactionID(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.TransactionDetails742340_TransactionID(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) TransactionIDDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails742340_TransactionIDDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetails742340_SetTransactionID
func callbackTransactionDetails742340_SetTransactionID(ptr unsafe.Pointer, transactionID C.struct_Moc_PackedString) {
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
		C.TransactionDetails742340_SetTransactionID(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

func (ptr *TransactionDetails) SetTransactionIDDefault(transactionID string) {
	if ptr.Pointer() != nil {
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		C.TransactionDetails742340_SetTransactionIDDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

//export callbackTransactionDetails742340_TransactionIDChanged
func callbackTransactionDetails742340_TransactionIDChanged(ptr unsafe.Pointer, transactionID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transactionIDChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(transactionID))
	}

}

func (ptr *TransactionDetails) ConnectTransactionIDChanged(f func(transactionID string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transactionIDChanged") {
			C.TransactionDetails742340_ConnectTransactionIDChanged(ptr.Pointer())
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
		C.TransactionDetails742340_DisconnectTransactionIDChanged(ptr.Pointer())
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
		C.TransactionDetails742340_TransactionIDChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

//export callbackTransactionDetails742340_SentAddress
func callbackTransactionDetails742340_SentAddress(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.TransactionDetails742340_SentAddress(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) SentAddressDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails742340_SentAddressDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetails742340_SetSentAddress
func callbackTransactionDetails742340_SetSentAddress(ptr unsafe.Pointer, sentAddress C.struct_Moc_PackedString) {
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
		C.TransactionDetails742340_SetSentAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: sentAddressC, len: C.longlong(len(sentAddress))})
	}
}

func (ptr *TransactionDetails) SetSentAddressDefault(sentAddress string) {
	if ptr.Pointer() != nil {
		var sentAddressC *C.char
		if sentAddress != "" {
			sentAddressC = C.CString(sentAddress)
			defer C.free(unsafe.Pointer(sentAddressC))
		}
		C.TransactionDetails742340_SetSentAddressDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: sentAddressC, len: C.longlong(len(sentAddress))})
	}
}

//export callbackTransactionDetails742340_SentAddressChanged
func callbackTransactionDetails742340_SentAddressChanged(ptr unsafe.Pointer, sentAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sentAddressChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(sentAddress))
	}

}

func (ptr *TransactionDetails) ConnectSentAddressChanged(f func(sentAddress string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sentAddressChanged") {
			C.TransactionDetails742340_ConnectSentAddressChanged(ptr.Pointer())
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
		C.TransactionDetails742340_DisconnectSentAddressChanged(ptr.Pointer())
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
		C.TransactionDetails742340_SentAddressChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: sentAddressC, len: C.longlong(len(sentAddress))})
	}
}

//export callbackTransactionDetails742340_ReceivedAddress
func callbackTransactionDetails742340_ReceivedAddress(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.TransactionDetails742340_ReceivedAddress(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) ReceivedAddressDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetails742340_ReceivedAddressDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetails742340_SetReceivedAddress
func callbackTransactionDetails742340_SetReceivedAddress(ptr unsafe.Pointer, receivedAddress C.struct_Moc_PackedString) {
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
		C.TransactionDetails742340_SetReceivedAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: receivedAddressC, len: C.longlong(len(receivedAddress))})
	}
}

func (ptr *TransactionDetails) SetReceivedAddressDefault(receivedAddress string) {
	if ptr.Pointer() != nil {
		var receivedAddressC *C.char
		if receivedAddress != "" {
			receivedAddressC = C.CString(receivedAddress)
			defer C.free(unsafe.Pointer(receivedAddressC))
		}
		C.TransactionDetails742340_SetReceivedAddressDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: receivedAddressC, len: C.longlong(len(receivedAddress))})
	}
}

//export callbackTransactionDetails742340_ReceivedAddressChanged
func callbackTransactionDetails742340_ReceivedAddressChanged(ptr unsafe.Pointer, receivedAddress C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "receivedAddressChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(receivedAddress))
	}

}

func (ptr *TransactionDetails) ConnectReceivedAddressChanged(f func(receivedAddress string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "receivedAddressChanged") {
			C.TransactionDetails742340_ConnectReceivedAddressChanged(ptr.Pointer())
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
		C.TransactionDetails742340_DisconnectReceivedAddressChanged(ptr.Pointer())
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
		C.TransactionDetails742340_ReceivedAddressChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: receivedAddressC, len: C.longlong(len(receivedAddress))})
	}
}

//export callbackTransactionDetails742340_Inputs
func callbackTransactionDetails742340_Inputs(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputs"); signal != nil {
		return PointerFromAddressList((*(*func() *AddressList)(signal))())
	}

	return PointerFromAddressList(NewTransactionDetailsFromPointer(ptr).InputsDefault())
}

func (ptr *TransactionDetails) ConnectInputs(f func() *AddressList) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "inputs"); signal != nil {
			f := func() *AddressList {
				(*(*func() *AddressList)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "inputs", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputs", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectInputs() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "inputs")
	}
}

func (ptr *TransactionDetails) Inputs() *AddressList {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressListFromPointer(C.TransactionDetails742340_Inputs(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) InputsDefault() *AddressList {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressListFromPointer(C.TransactionDetails742340_InputsDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackTransactionDetails742340_SetInputs
func callbackTransactionDetails742340_SetInputs(ptr unsafe.Pointer, inputs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setInputs"); signal != nil {
		(*(*func(*AddressList))(signal))(NewAddressListFromPointer(inputs))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetInputsDefault(NewAddressListFromPointer(inputs))
	}
}

func (ptr *TransactionDetails) ConnectSetInputs(f func(inputs *AddressList)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setInputs"); signal != nil {
			f := func(inputs *AddressList) {
				(*(*func(*AddressList))(signal))(inputs)
				f(inputs)
			}
			qt.ConnectSignal(ptr.Pointer(), "setInputs", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setInputs", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetInputs() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setInputs")
	}
}

func (ptr *TransactionDetails) SetInputs(inputs AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_SetInputs(ptr.Pointer(), PointerFromAddressList(inputs))
	}
}

func (ptr *TransactionDetails) SetInputsDefault(inputs AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_SetInputsDefault(ptr.Pointer(), PointerFromAddressList(inputs))
	}
}

//export callbackTransactionDetails742340_InputsChanged
func callbackTransactionDetails742340_InputsChanged(ptr unsafe.Pointer, inputs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputsChanged"); signal != nil {
		(*(*func(*AddressList))(signal))(NewAddressListFromPointer(inputs))
	}

}

func (ptr *TransactionDetails) ConnectInputsChanged(f func(inputs *AddressList)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "inputsChanged") {
			C.TransactionDetails742340_ConnectInputsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "inputsChanged"); signal != nil {
			f := func(inputs *AddressList) {
				(*(*func(*AddressList))(signal))(inputs)
				f(inputs)
			}
			qt.ConnectSignal(ptr.Pointer(), "inputsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "inputsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectInputsChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_DisconnectInputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "inputsChanged")
	}
}

func (ptr *TransactionDetails) InputsChanged(inputs AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_InputsChanged(ptr.Pointer(), PointerFromAddressList(inputs))
	}
}

//export callbackTransactionDetails742340_Outputs
func callbackTransactionDetails742340_Outputs(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "outputs"); signal != nil {
		return PointerFromAddressList((*(*func() *AddressList)(signal))())
	}

	return PointerFromAddressList(NewTransactionDetailsFromPointer(ptr).OutputsDefault())
}

func (ptr *TransactionDetails) ConnectOutputs(f func() *AddressList) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "outputs"); signal != nil {
			f := func() *AddressList {
				(*(*func() *AddressList)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "outputs", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "outputs", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectOutputs() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "outputs")
	}
}

func (ptr *TransactionDetails) Outputs() *AddressList {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressListFromPointer(C.TransactionDetails742340_Outputs(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) OutputsDefault() *AddressList {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressListFromPointer(C.TransactionDetails742340_OutputsDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackTransactionDetails742340_SetOutputs
func callbackTransactionDetails742340_SetOutputs(ptr unsafe.Pointer, outputs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setOutputs"); signal != nil {
		(*(*func(*AddressList))(signal))(NewAddressListFromPointer(outputs))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetOutputsDefault(NewAddressListFromPointer(outputs))
	}
}

func (ptr *TransactionDetails) ConnectSetOutputs(f func(outputs *AddressList)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setOutputs"); signal != nil {
			f := func(outputs *AddressList) {
				(*(*func(*AddressList))(signal))(outputs)
				f(outputs)
			}
			qt.ConnectSignal(ptr.Pointer(), "setOutputs", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setOutputs", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetOutputs() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setOutputs")
	}
}

func (ptr *TransactionDetails) SetOutputs(outputs AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_SetOutputs(ptr.Pointer(), PointerFromAddressList(outputs))
	}
}

func (ptr *TransactionDetails) SetOutputsDefault(outputs AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_SetOutputsDefault(ptr.Pointer(), PointerFromAddressList(outputs))
	}
}

//export callbackTransactionDetails742340_OutputsChanged
func callbackTransactionDetails742340_OutputsChanged(ptr unsafe.Pointer, outputs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "outputsChanged"); signal != nil {
		(*(*func(*AddressList))(signal))(NewAddressListFromPointer(outputs))
	}

}

func (ptr *TransactionDetails) ConnectOutputsChanged(f func(outputs *AddressList)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "outputsChanged") {
			C.TransactionDetails742340_ConnectOutputsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "outputsChanged"); signal != nil {
			f := func(outputs *AddressList) {
				(*(*func(*AddressList))(signal))(outputs)
				f(outputs)
			}
			qt.ConnectSignal(ptr.Pointer(), "outputsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "outputsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectOutputsChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_DisconnectOutputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "outputsChanged")
	}
}

func (ptr *TransactionDetails) OutputsChanged(outputs AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_OutputsChanged(ptr.Pointer(), PointerFromAddressList(outputs))
	}
}

func TransactionDetails_QRegisterMetaType() int {
	return int(int32(C.TransactionDetails742340_TransactionDetails742340_QRegisterMetaType()))
}

func (ptr *TransactionDetails) QRegisterMetaType() int {
	return int(int32(C.TransactionDetails742340_TransactionDetails742340_QRegisterMetaType()))
}

func TransactionDetails_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TransactionDetails742340_TransactionDetails742340_QRegisterMetaType2(typeNameC)))
}

func (ptr *TransactionDetails) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TransactionDetails742340_TransactionDetails742340_QRegisterMetaType2(typeNameC)))
}

func TransactionDetails_QmlRegisterType() int {
	return int(int32(C.TransactionDetails742340_TransactionDetails742340_QmlRegisterType()))
}

func (ptr *TransactionDetails) QmlRegisterType() int {
	return int(int32(C.TransactionDetails742340_TransactionDetails742340_QmlRegisterType()))
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
	return int(int32(C.TransactionDetails742340_TransactionDetails742340_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.TransactionDetails742340_TransactionDetails742340_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TransactionDetails) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetails742340___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __children_newList() unsafe.Pointer {
	return C.TransactionDetails742340___children_newList(ptr.Pointer())
}

func (ptr *TransactionDetails) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionDetails742340___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionDetails) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.TransactionDetails742340___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *TransactionDetails) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetails742340___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __findChildren_newList() unsafe.Pointer {
	return C.TransactionDetails742340___findChildren_newList(ptr.Pointer())
}

func (ptr *TransactionDetails) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetails742340___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __findChildren_newList3() unsafe.Pointer {
	return C.TransactionDetails742340___findChildren_newList3(ptr.Pointer())
}

func (ptr *TransactionDetails) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetails742340___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __qFindChildren_newList2() unsafe.Pointer {
	return C.TransactionDetails742340___qFindChildren_newList2(ptr.Pointer())
}

func NewTransactionDetails(parent std_core.QObject_ITF) *TransactionDetails {
	tmpValue := NewTransactionDetailsFromPointer(C.TransactionDetails742340_NewTransactionDetails(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTransactionDetails742340_DestroyTransactionDetails
func callbackTransactionDetails742340_DestroyTransactionDetails(ptr unsafe.Pointer) {
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
		C.TransactionDetails742340_DestroyTransactionDetails(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *TransactionDetails) DestroyTransactionDetailsDefault() {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_DestroyTransactionDetailsDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTransactionDetails742340_ChildEvent
func callbackTransactionDetails742340_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTransactionDetailsFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *TransactionDetails) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTransactionDetails742340_ConnectNotify
func callbackTransactionDetails742340_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTransactionDetailsFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TransactionDetails) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTransactionDetails742340_CustomEvent
func callbackTransactionDetails742340_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewTransactionDetailsFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *TransactionDetails) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTransactionDetails742340_DeleteLater
func callbackTransactionDetails742340_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionDetailsFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *TransactionDetails) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTransactionDetails742340_Destroyed
func callbackTransactionDetails742340_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackTransactionDetails742340_DisconnectNotify
func callbackTransactionDetails742340_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTransactionDetailsFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TransactionDetails) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTransactionDetails742340_Event
func callbackTransactionDetails742340_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionDetailsFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *TransactionDetails) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionDetails742340_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackTransactionDetails742340_EventFilter
func callbackTransactionDetails742340_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionDetailsFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *TransactionDetails) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionDetails742340_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackTransactionDetails742340_ObjectNameChanged
func callbackTransactionDetails742340_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackTransactionDetails742340_TimerEvent
func callbackTransactionDetails742340_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTransactionDetailsFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *TransactionDetails) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetails742340_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

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

//export callbackWalletModel268d39_TestFunc
func callbackWalletModel268d39_TestFunc(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "testFunc"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *WalletModel) ConnectTestFunc(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "testFunc"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "testFunc", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "testFunc", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectTestFunc() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "testFunc")
	}
}

func (ptr *WalletModel) TestFunc() {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_TestFunc(ptr.Pointer())
	}
}

//export callbackWalletModel268d39_Created
func callbackWalletModel268d39_Created(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "created"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *WalletModel) ConnectCreated(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "created") {
			C.WalletModel268d39_ConnectCreated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "created"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "created", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "created", unsafe.Pointer(&f))
		}
	}
}

func (ptr *WalletModel) DisconnectCreated() {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_DisconnectCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "created")
	}
}

func (ptr *WalletModel) Created() {
	if ptr.Pointer() != nil {
		C.WalletModel268d39_Created(ptr.Pointer())
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

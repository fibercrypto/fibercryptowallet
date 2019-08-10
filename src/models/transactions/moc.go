package transactions

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	custom_address_fb4c44m "github.com/fibercrypto/FiberCryptoWallet/src/models/address"
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

//export callbackTransactionDetailsecff1c_Constructor
func callbackTransactionDetailsecff1c_Constructor(ptr unsafe.Pointer) {
	this := NewTransactionDetailsFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackTransactionDetailsecff1c_Date
func callbackTransactionDetailsecff1c_Date(ptr unsafe.Pointer) unsafe.Pointer {
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
		tmpValue := std_core.NewQDateTimeFromPointer(C.TransactionDetailsecff1c_Date(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) DateDefault() *std_core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQDateTimeFromPointer(C.TransactionDetailsecff1c_DateDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

//export callbackTransactionDetailsecff1c_SetDate
func callbackTransactionDetailsecff1c_SetDate(ptr unsafe.Pointer, date unsafe.Pointer) {
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
		C.TransactionDetailsecff1c_SetDate(ptr.Pointer(), std_core.PointerFromQDateTime(date))
	}
}

func (ptr *TransactionDetails) SetDateDefault(date std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_SetDateDefault(ptr.Pointer(), std_core.PointerFromQDateTime(date))
	}
}

//export callbackTransactionDetailsecff1c_DateChanged
func callbackTransactionDetailsecff1c_DateChanged(ptr unsafe.Pointer, date unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dateChanged"); signal != nil {
		(*(*func(*std_core.QDateTime))(signal))(std_core.NewQDateTimeFromPointer(date))
	}

}

func (ptr *TransactionDetails) ConnectDateChanged(f func(date *std_core.QDateTime)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dateChanged") {
			C.TransactionDetailsecff1c_ConnectDateChanged(ptr.Pointer())
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
		C.TransactionDetailsecff1c_DisconnectDateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dateChanged")
	}
}

func (ptr *TransactionDetails) DateChanged(date std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_DateChanged(ptr.Pointer(), std_core.PointerFromQDateTime(date))
	}
}

//export callbackTransactionDetailsecff1c_Status
func callbackTransactionDetailsecff1c_Status(ptr unsafe.Pointer) C.int {
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
		return int(int32(C.TransactionDetailsecff1c_Status(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) StatusDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetailsecff1c_StatusDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetailsecff1c_SetStatus
func callbackTransactionDetailsecff1c_SetStatus(ptr unsafe.Pointer, status C.int) {
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
		C.TransactionDetailsecff1c_SetStatus(ptr.Pointer(), C.int(int32(status)))
	}
}

func (ptr *TransactionDetails) SetStatusDefault(status int) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_SetStatusDefault(ptr.Pointer(), C.int(int32(status)))
	}
}

//export callbackTransactionDetailsecff1c_StatusChanged
func callbackTransactionDetailsecff1c_StatusChanged(ptr unsafe.Pointer, status C.int) {
	if signal := qt.GetSignal(ptr, "statusChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(status)))
	}

}

func (ptr *TransactionDetails) ConnectStatusChanged(f func(status int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "statusChanged") {
			C.TransactionDetailsecff1c_ConnectStatusChanged(ptr.Pointer())
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
		C.TransactionDetailsecff1c_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "statusChanged")
	}
}

func (ptr *TransactionDetails) StatusChanged(status int) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_StatusChanged(ptr.Pointer(), C.int(int32(status)))
	}
}

//export callbackTransactionDetailsecff1c_Type
func callbackTransactionDetailsecff1c_Type(ptr unsafe.Pointer) C.int {
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
		return int(int32(C.TransactionDetailsecff1c_Type(ptr.Pointer())))
	}
	return 0
}

func (ptr *TransactionDetails) TypeDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.TransactionDetailsecff1c_TypeDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackTransactionDetailsecff1c_SetType
func callbackTransactionDetailsecff1c_SetType(ptr unsafe.Pointer, ty C.int) {
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
		C.TransactionDetailsecff1c_SetType(ptr.Pointer(), C.int(int32(ty)))
	}
}

func (ptr *TransactionDetails) SetTypeDefault(ty int) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_SetTypeDefault(ptr.Pointer(), C.int(int32(ty)))
	}
}

//export callbackTransactionDetailsecff1c_TypeChanged
func callbackTransactionDetailsecff1c_TypeChanged(ptr unsafe.Pointer, ty C.int) {
	if signal := qt.GetSignal(ptr, "typeChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(ty)))
	}

}

func (ptr *TransactionDetails) ConnectTypeChanged(f func(ty int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "typeChanged") {
			C.TransactionDetailsecff1c_ConnectTypeChanged(ptr.Pointer())
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
		C.TransactionDetailsecff1c_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "typeChanged")
	}
}

func (ptr *TransactionDetails) TypeChanged(ty int) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_TypeChanged(ptr.Pointer(), C.int(int32(ty)))
	}
}

//export callbackTransactionDetailsecff1c_Amount
func callbackTransactionDetailsecff1c_Amount(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "amount"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTransactionDetailsFromPointer(ptr).AmountDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *TransactionDetails) ConnectAmount(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "amount"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
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

func (ptr *TransactionDetails) Amount() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetailsecff1c_Amount(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) AmountDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetailsecff1c_AmountDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetailsecff1c_SetAmount
func callbackTransactionDetailsecff1c_SetAmount(ptr unsafe.Pointer, amount C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setAmount"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(amount))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetAmountDefault(cGoUnpackString(amount))
	}
}

func (ptr *TransactionDetails) ConnectSetAmount(f func(amount string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAmount"); signal != nil {
			f := func(amount string) {
				(*(*func(string))(signal))(amount)
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

func (ptr *TransactionDetails) SetAmount(amount string) {
	if ptr.Pointer() != nil {
		var amountC *C.char
		if amount != "" {
			amountC = C.CString(amount)
			defer C.free(unsafe.Pointer(amountC))
		}
		C.TransactionDetailsecff1c_SetAmount(ptr.Pointer(), C.struct_Moc_PackedString{data: amountC, len: C.longlong(len(amount))})
	}
}

func (ptr *TransactionDetails) SetAmountDefault(amount string) {
	if ptr.Pointer() != nil {
		var amountC *C.char
		if amount != "" {
			amountC = C.CString(amount)
			defer C.free(unsafe.Pointer(amountC))
		}
		C.TransactionDetailsecff1c_SetAmountDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: amountC, len: C.longlong(len(amount))})
	}
}

//export callbackTransactionDetailsecff1c_AmountChanged
func callbackTransactionDetailsecff1c_AmountChanged(ptr unsafe.Pointer, amount C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "amountChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(amount))
	}

}

func (ptr *TransactionDetails) ConnectAmountChanged(f func(amount string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "amountChanged") {
			C.TransactionDetailsecff1c_ConnectAmountChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "amountChanged"); signal != nil {
			f := func(amount string) {
				(*(*func(string))(signal))(amount)
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
		C.TransactionDetailsecff1c_DisconnectAmountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "amountChanged")
	}
}

func (ptr *TransactionDetails) AmountChanged(amount string) {
	if ptr.Pointer() != nil {
		var amountC *C.char
		if amount != "" {
			amountC = C.CString(amount)
			defer C.free(unsafe.Pointer(amountC))
		}
		C.TransactionDetailsecff1c_AmountChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: amountC, len: C.longlong(len(amount))})
	}
}

//export callbackTransactionDetailsecff1c_HoursTraspassed
func callbackTransactionDetailsecff1c_HoursTraspassed(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "hoursTraspassed"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTransactionDetailsFromPointer(ptr).HoursTraspassedDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *TransactionDetails) ConnectHoursTraspassed(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hoursTraspassed"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hoursTraspassed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoursTraspassed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectHoursTraspassed() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hoursTraspassed")
	}
}

func (ptr *TransactionDetails) HoursTraspassed() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetailsecff1c_HoursTraspassed(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) HoursTraspassedDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetailsecff1c_HoursTraspassedDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetailsecff1c_SetHoursTraspassed
func callbackTransactionDetailsecff1c_SetHoursTraspassed(ptr unsafe.Pointer, hoursTraspassed C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setHoursTraspassed"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(hoursTraspassed))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetHoursTraspassedDefault(cGoUnpackString(hoursTraspassed))
	}
}

func (ptr *TransactionDetails) ConnectSetHoursTraspassed(f func(hoursTraspassed string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHoursTraspassed"); signal != nil {
			f := func(hoursTraspassed string) {
				(*(*func(string))(signal))(hoursTraspassed)
				f(hoursTraspassed)
			}
			qt.ConnectSignal(ptr.Pointer(), "setHoursTraspassed", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setHoursTraspassed", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetHoursTraspassed() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHoursTraspassed")
	}
}

func (ptr *TransactionDetails) SetHoursTraspassed(hoursTraspassed string) {
	if ptr.Pointer() != nil {
		var hoursTraspassedC *C.char
		if hoursTraspassed != "" {
			hoursTraspassedC = C.CString(hoursTraspassed)
			defer C.free(unsafe.Pointer(hoursTraspassedC))
		}
		C.TransactionDetailsecff1c_SetHoursTraspassed(ptr.Pointer(), C.struct_Moc_PackedString{data: hoursTraspassedC, len: C.longlong(len(hoursTraspassed))})
	}
}

func (ptr *TransactionDetails) SetHoursTraspassedDefault(hoursTraspassed string) {
	if ptr.Pointer() != nil {
		var hoursTraspassedC *C.char
		if hoursTraspassed != "" {
			hoursTraspassedC = C.CString(hoursTraspassed)
			defer C.free(unsafe.Pointer(hoursTraspassedC))
		}
		C.TransactionDetailsecff1c_SetHoursTraspassedDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: hoursTraspassedC, len: C.longlong(len(hoursTraspassed))})
	}
}

//export callbackTransactionDetailsecff1c_HoursTraspassedChanged
func callbackTransactionDetailsecff1c_HoursTraspassedChanged(ptr unsafe.Pointer, hoursTraspassed C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "hoursTraspassedChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(hoursTraspassed))
	}

}

func (ptr *TransactionDetails) ConnectHoursTraspassedChanged(f func(hoursTraspassed string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hoursTraspassedChanged") {
			C.TransactionDetailsecff1c_ConnectHoursTraspassedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hoursTraspassedChanged"); signal != nil {
			f := func(hoursTraspassed string) {
				(*(*func(string))(signal))(hoursTraspassed)
				f(hoursTraspassed)
			}
			qt.ConnectSignal(ptr.Pointer(), "hoursTraspassedChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hoursTraspassedChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectHoursTraspassedChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_DisconnectHoursTraspassedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hoursTraspassedChanged")
	}
}

func (ptr *TransactionDetails) HoursTraspassedChanged(hoursTraspassed string) {
	if ptr.Pointer() != nil {
		var hoursTraspassedC *C.char
		if hoursTraspassed != "" {
			hoursTraspassedC = C.CString(hoursTraspassed)
			defer C.free(unsafe.Pointer(hoursTraspassedC))
		}
		C.TransactionDetailsecff1c_HoursTraspassedChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: hoursTraspassedC, len: C.longlong(len(hoursTraspassed))})
	}
}

//export callbackTransactionDetailsecff1c_HoursBurned
func callbackTransactionDetailsecff1c_HoursBurned(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "hoursBurned"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewTransactionDetailsFromPointer(ptr).HoursBurnedDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *TransactionDetails) ConnectHoursBurned(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hoursBurned"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
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

func (ptr *TransactionDetails) HoursBurned() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetailsecff1c_HoursBurned(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) HoursBurnedDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetailsecff1c_HoursBurnedDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetailsecff1c_SetHoursBurned
func callbackTransactionDetailsecff1c_SetHoursBurned(ptr unsafe.Pointer, hoursBurned C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setHoursBurned"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(hoursBurned))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetHoursBurnedDefault(cGoUnpackString(hoursBurned))
	}
}

func (ptr *TransactionDetails) ConnectSetHoursBurned(f func(hoursBurned string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHoursBurned"); signal != nil {
			f := func(hoursBurned string) {
				(*(*func(string))(signal))(hoursBurned)
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

func (ptr *TransactionDetails) SetHoursBurned(hoursBurned string) {
	if ptr.Pointer() != nil {
		var hoursBurnedC *C.char
		if hoursBurned != "" {
			hoursBurnedC = C.CString(hoursBurned)
			defer C.free(unsafe.Pointer(hoursBurnedC))
		}
		C.TransactionDetailsecff1c_SetHoursBurned(ptr.Pointer(), C.struct_Moc_PackedString{data: hoursBurnedC, len: C.longlong(len(hoursBurned))})
	}
}

func (ptr *TransactionDetails) SetHoursBurnedDefault(hoursBurned string) {
	if ptr.Pointer() != nil {
		var hoursBurnedC *C.char
		if hoursBurned != "" {
			hoursBurnedC = C.CString(hoursBurned)
			defer C.free(unsafe.Pointer(hoursBurnedC))
		}
		C.TransactionDetailsecff1c_SetHoursBurnedDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: hoursBurnedC, len: C.longlong(len(hoursBurned))})
	}
}

//export callbackTransactionDetailsecff1c_HoursBurnedChanged
func callbackTransactionDetailsecff1c_HoursBurnedChanged(ptr unsafe.Pointer, hoursBurned C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "hoursBurnedChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(hoursBurned))
	}

}

func (ptr *TransactionDetails) ConnectHoursBurnedChanged(f func(hoursBurned string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hoursBurnedChanged") {
			C.TransactionDetailsecff1c_ConnectHoursBurnedChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hoursBurnedChanged"); signal != nil {
			f := func(hoursBurned string) {
				(*(*func(string))(signal))(hoursBurned)
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
		C.TransactionDetailsecff1c_DisconnectHoursBurnedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hoursBurnedChanged")
	}
}

func (ptr *TransactionDetails) HoursBurnedChanged(hoursBurned string) {
	if ptr.Pointer() != nil {
		var hoursBurnedC *C.char
		if hoursBurned != "" {
			hoursBurnedC = C.CString(hoursBurned)
			defer C.free(unsafe.Pointer(hoursBurnedC))
		}
		C.TransactionDetailsecff1c_HoursBurnedChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: hoursBurnedC, len: C.longlong(len(hoursBurned))})
	}
}

//export callbackTransactionDetailsecff1c_TransactionID
func callbackTransactionDetailsecff1c_TransactionID(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.TransactionDetailsecff1c_TransactionID(ptr.Pointer()))
	}
	return ""
}

func (ptr *TransactionDetails) TransactionIDDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.TransactionDetailsecff1c_TransactionIDDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackTransactionDetailsecff1c_SetTransactionID
func callbackTransactionDetailsecff1c_SetTransactionID(ptr unsafe.Pointer, transactionID C.struct_Moc_PackedString) {
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
		C.TransactionDetailsecff1c_SetTransactionID(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

func (ptr *TransactionDetails) SetTransactionIDDefault(transactionID string) {
	if ptr.Pointer() != nil {
		var transactionIDC *C.char
		if transactionID != "" {
			transactionIDC = C.CString(transactionID)
			defer C.free(unsafe.Pointer(transactionIDC))
		}
		C.TransactionDetailsecff1c_SetTransactionIDDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

//export callbackTransactionDetailsecff1c_TransactionIDChanged
func callbackTransactionDetailsecff1c_TransactionIDChanged(ptr unsafe.Pointer, transactionID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "transactionIDChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(transactionID))
	}

}

func (ptr *TransactionDetails) ConnectTransactionIDChanged(f func(transactionID string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transactionIDChanged") {
			C.TransactionDetailsecff1c_ConnectTransactionIDChanged(ptr.Pointer())
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
		C.TransactionDetailsecff1c_DisconnectTransactionIDChanged(ptr.Pointer())
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
		C.TransactionDetailsecff1c_TransactionIDChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: transactionIDC, len: C.longlong(len(transactionID))})
	}
}

//export callbackTransactionDetailsecff1c_Addresses
func callbackTransactionDetailsecff1c_Addresses(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "addresses"); signal != nil {
		return custom_address_fb4c44m.PointerFromAddressList((*(*func() *custom_address_fb4c44m.AddressList)(signal))())
	}

	return custom_address_fb4c44m.PointerFromAddressList(NewTransactionDetailsFromPointer(ptr).AddressesDefault())
}

func (ptr *TransactionDetails) ConnectAddresses(f func() *custom_address_fb4c44m.AddressList) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addresses"); signal != nil {
			f := func() *custom_address_fb4c44m.AddressList {
				(*(*func() *custom_address_fb4c44m.AddressList)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "addresses", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addresses", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectAddresses() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addresses")
	}
}

func (ptr *TransactionDetails) Addresses() *custom_address_fb4c44m.AddressList {
	if ptr.Pointer() != nil {
		tmpValue := custom_address_fb4c44m.NewAddressListFromPointer(C.TransactionDetailsecff1c_Addresses(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) AddressesDefault() *custom_address_fb4c44m.AddressList {
	if ptr.Pointer() != nil {
		tmpValue := custom_address_fb4c44m.NewAddressListFromPointer(C.TransactionDetailsecff1c_AddressesDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackTransactionDetailsecff1c_SetAddresses
func callbackTransactionDetailsecff1c_SetAddresses(ptr unsafe.Pointer, addresses unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setAddresses"); signal != nil {
		(*(*func(*custom_address_fb4c44m.AddressList))(signal))(custom_address_fb4c44m.NewAddressListFromPointer(addresses))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetAddressesDefault(custom_address_fb4c44m.NewAddressListFromPointer(addresses))
	}
}

func (ptr *TransactionDetails) ConnectSetAddresses(f func(addresses *custom_address_fb4c44m.AddressList)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddresses"); signal != nil {
			f := func(addresses *custom_address_fb4c44m.AddressList) {
				(*(*func(*custom_address_fb4c44m.AddressList))(signal))(addresses)
				f(addresses)
			}
			qt.ConnectSignal(ptr.Pointer(), "setAddresses", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setAddresses", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectSetAddresses() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setAddresses")
	}
}

func (ptr *TransactionDetails) SetAddresses(addresses custom_address_fb4c44m.AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_SetAddresses(ptr.Pointer(), custom_address_fb4c44m.PointerFromAddressList(addresses))
	}
}

func (ptr *TransactionDetails) SetAddressesDefault(addresses custom_address_fb4c44m.AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_SetAddressesDefault(ptr.Pointer(), custom_address_fb4c44m.PointerFromAddressList(addresses))
	}
}

//export callbackTransactionDetailsecff1c_AddressesChanged
func callbackTransactionDetailsecff1c_AddressesChanged(ptr unsafe.Pointer, addresses unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addressesChanged"); signal != nil {
		(*(*func(*custom_address_fb4c44m.AddressList))(signal))(custom_address_fb4c44m.NewAddressListFromPointer(addresses))
	}

}

func (ptr *TransactionDetails) ConnectAddressesChanged(f func(addresses *custom_address_fb4c44m.AddressList)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressesChanged") {
			C.TransactionDetailsecff1c_ConnectAddressesChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressesChanged"); signal != nil {
			f := func(addresses *custom_address_fb4c44m.AddressList) {
				(*(*func(*custom_address_fb4c44m.AddressList))(signal))(addresses)
				f(addresses)
			}
			qt.ConnectSignal(ptr.Pointer(), "addressesChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addressesChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TransactionDetails) DisconnectAddressesChanged() {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_DisconnectAddressesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressesChanged")
	}
}

func (ptr *TransactionDetails) AddressesChanged(addresses custom_address_fb4c44m.AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_AddressesChanged(ptr.Pointer(), custom_address_fb4c44m.PointerFromAddressList(addresses))
	}
}

//export callbackTransactionDetailsecff1c_Inputs
func callbackTransactionDetailsecff1c_Inputs(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputs"); signal != nil {
		return custom_address_fb4c44m.PointerFromAddressList((*(*func() *custom_address_fb4c44m.AddressList)(signal))())
	}

	return custom_address_fb4c44m.PointerFromAddressList(NewTransactionDetailsFromPointer(ptr).InputsDefault())
}

func (ptr *TransactionDetails) ConnectInputs(f func() *custom_address_fb4c44m.AddressList) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "inputs"); signal != nil {
			f := func() *custom_address_fb4c44m.AddressList {
				(*(*func() *custom_address_fb4c44m.AddressList)(signal))()
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

func (ptr *TransactionDetails) Inputs() *custom_address_fb4c44m.AddressList {
	if ptr.Pointer() != nil {
		tmpValue := custom_address_fb4c44m.NewAddressListFromPointer(C.TransactionDetailsecff1c_Inputs(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) InputsDefault() *custom_address_fb4c44m.AddressList {
	if ptr.Pointer() != nil {
		tmpValue := custom_address_fb4c44m.NewAddressListFromPointer(C.TransactionDetailsecff1c_InputsDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackTransactionDetailsecff1c_SetInputs
func callbackTransactionDetailsecff1c_SetInputs(ptr unsafe.Pointer, inputs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setInputs"); signal != nil {
		(*(*func(*custom_address_fb4c44m.AddressList))(signal))(custom_address_fb4c44m.NewAddressListFromPointer(inputs))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetInputsDefault(custom_address_fb4c44m.NewAddressListFromPointer(inputs))
	}
}

func (ptr *TransactionDetails) ConnectSetInputs(f func(inputs *custom_address_fb4c44m.AddressList)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setInputs"); signal != nil {
			f := func(inputs *custom_address_fb4c44m.AddressList) {
				(*(*func(*custom_address_fb4c44m.AddressList))(signal))(inputs)
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

func (ptr *TransactionDetails) SetInputs(inputs custom_address_fb4c44m.AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_SetInputs(ptr.Pointer(), custom_address_fb4c44m.PointerFromAddressList(inputs))
	}
}

func (ptr *TransactionDetails) SetInputsDefault(inputs custom_address_fb4c44m.AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_SetInputsDefault(ptr.Pointer(), custom_address_fb4c44m.PointerFromAddressList(inputs))
	}
}

//export callbackTransactionDetailsecff1c_InputsChanged
func callbackTransactionDetailsecff1c_InputsChanged(ptr unsafe.Pointer, inputs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputsChanged"); signal != nil {
		(*(*func(*custom_address_fb4c44m.AddressList))(signal))(custom_address_fb4c44m.NewAddressListFromPointer(inputs))
	}

}

func (ptr *TransactionDetails) ConnectInputsChanged(f func(inputs *custom_address_fb4c44m.AddressList)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "inputsChanged") {
			C.TransactionDetailsecff1c_ConnectInputsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "inputsChanged"); signal != nil {
			f := func(inputs *custom_address_fb4c44m.AddressList) {
				(*(*func(*custom_address_fb4c44m.AddressList))(signal))(inputs)
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
		C.TransactionDetailsecff1c_DisconnectInputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "inputsChanged")
	}
}

func (ptr *TransactionDetails) InputsChanged(inputs custom_address_fb4c44m.AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_InputsChanged(ptr.Pointer(), custom_address_fb4c44m.PointerFromAddressList(inputs))
	}
}

//export callbackTransactionDetailsecff1c_Outputs
func callbackTransactionDetailsecff1c_Outputs(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "outputs"); signal != nil {
		return custom_address_fb4c44m.PointerFromAddressList((*(*func() *custom_address_fb4c44m.AddressList)(signal))())
	}

	return custom_address_fb4c44m.PointerFromAddressList(NewTransactionDetailsFromPointer(ptr).OutputsDefault())
}

func (ptr *TransactionDetails) ConnectOutputs(f func() *custom_address_fb4c44m.AddressList) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "outputs"); signal != nil {
			f := func() *custom_address_fb4c44m.AddressList {
				(*(*func() *custom_address_fb4c44m.AddressList)(signal))()
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

func (ptr *TransactionDetails) Outputs() *custom_address_fb4c44m.AddressList {
	if ptr.Pointer() != nil {
		tmpValue := custom_address_fb4c44m.NewAddressListFromPointer(C.TransactionDetailsecff1c_Outputs(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) OutputsDefault() *custom_address_fb4c44m.AddressList {
	if ptr.Pointer() != nil {
		tmpValue := custom_address_fb4c44m.NewAddressListFromPointer(C.TransactionDetailsecff1c_OutputsDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackTransactionDetailsecff1c_SetOutputs
func callbackTransactionDetailsecff1c_SetOutputs(ptr unsafe.Pointer, outputs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setOutputs"); signal != nil {
		(*(*func(*custom_address_fb4c44m.AddressList))(signal))(custom_address_fb4c44m.NewAddressListFromPointer(outputs))
	} else {
		NewTransactionDetailsFromPointer(ptr).SetOutputsDefault(custom_address_fb4c44m.NewAddressListFromPointer(outputs))
	}
}

func (ptr *TransactionDetails) ConnectSetOutputs(f func(outputs *custom_address_fb4c44m.AddressList)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setOutputs"); signal != nil {
			f := func(outputs *custom_address_fb4c44m.AddressList) {
				(*(*func(*custom_address_fb4c44m.AddressList))(signal))(outputs)
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

func (ptr *TransactionDetails) SetOutputs(outputs custom_address_fb4c44m.AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_SetOutputs(ptr.Pointer(), custom_address_fb4c44m.PointerFromAddressList(outputs))
	}
}

func (ptr *TransactionDetails) SetOutputsDefault(outputs custom_address_fb4c44m.AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_SetOutputsDefault(ptr.Pointer(), custom_address_fb4c44m.PointerFromAddressList(outputs))
	}
}

//export callbackTransactionDetailsecff1c_OutputsChanged
func callbackTransactionDetailsecff1c_OutputsChanged(ptr unsafe.Pointer, outputs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "outputsChanged"); signal != nil {
		(*(*func(*custom_address_fb4c44m.AddressList))(signal))(custom_address_fb4c44m.NewAddressListFromPointer(outputs))
	}

}

func (ptr *TransactionDetails) ConnectOutputsChanged(f func(outputs *custom_address_fb4c44m.AddressList)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "outputsChanged") {
			C.TransactionDetailsecff1c_ConnectOutputsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "outputsChanged"); signal != nil {
			f := func(outputs *custom_address_fb4c44m.AddressList) {
				(*(*func(*custom_address_fb4c44m.AddressList))(signal))(outputs)
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
		C.TransactionDetailsecff1c_DisconnectOutputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "outputsChanged")
	}
}

func (ptr *TransactionDetails) OutputsChanged(outputs custom_address_fb4c44m.AddressList_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_OutputsChanged(ptr.Pointer(), custom_address_fb4c44m.PointerFromAddressList(outputs))
	}
}

func TransactionDetails_QRegisterMetaType() int {
	return int(int32(C.TransactionDetailsecff1c_TransactionDetailsecff1c_QRegisterMetaType()))
}

func (ptr *TransactionDetails) QRegisterMetaType() int {
	return int(int32(C.TransactionDetailsecff1c_TransactionDetailsecff1c_QRegisterMetaType()))
}

func TransactionDetails_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TransactionDetailsecff1c_TransactionDetailsecff1c_QRegisterMetaType2(typeNameC)))
}

func (ptr *TransactionDetails) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TransactionDetailsecff1c_TransactionDetailsecff1c_QRegisterMetaType2(typeNameC)))
}

func TransactionDetails_QmlRegisterType() int {
	return int(int32(C.TransactionDetailsecff1c_TransactionDetailsecff1c_QmlRegisterType()))
}

func (ptr *TransactionDetails) QmlRegisterType() int {
	return int(int32(C.TransactionDetailsecff1c_TransactionDetailsecff1c_QmlRegisterType()))
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
	return int(int32(C.TransactionDetailsecff1c_TransactionDetailsecff1c_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.TransactionDetailsecff1c_TransactionDetailsecff1c_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TransactionDetails) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetailsecff1c___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __children_newList() unsafe.Pointer {
	return C.TransactionDetailsecff1c___children_newList(ptr.Pointer())
}

func (ptr *TransactionDetails) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TransactionDetailsecff1c___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TransactionDetails) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.TransactionDetailsecff1c___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *TransactionDetails) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetailsecff1c___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __findChildren_newList() unsafe.Pointer {
	return C.TransactionDetailsecff1c___findChildren_newList(ptr.Pointer())
}

func (ptr *TransactionDetails) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetailsecff1c___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __findChildren_newList3() unsafe.Pointer {
	return C.TransactionDetailsecff1c___findChildren_newList3(ptr.Pointer())
}

func (ptr *TransactionDetails) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TransactionDetailsecff1c___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TransactionDetails) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TransactionDetails) __qFindChildren_newList2() unsafe.Pointer {
	return C.TransactionDetailsecff1c___qFindChildren_newList2(ptr.Pointer())
}

func NewTransactionDetails(parent std_core.QObject_ITF) *TransactionDetails {
	tmpValue := NewTransactionDetailsFromPointer(C.TransactionDetailsecff1c_NewTransactionDetails(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTransactionDetailsecff1c_DestroyTransactionDetails
func callbackTransactionDetailsecff1c_DestroyTransactionDetails(ptr unsafe.Pointer) {
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
		C.TransactionDetailsecff1c_DestroyTransactionDetails(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *TransactionDetails) DestroyTransactionDetailsDefault() {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_DestroyTransactionDetailsDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTransactionDetailsecff1c_ChildEvent
func callbackTransactionDetailsecff1c_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTransactionDetailsFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *TransactionDetails) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTransactionDetailsecff1c_ConnectNotify
func callbackTransactionDetailsecff1c_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTransactionDetailsFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TransactionDetails) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTransactionDetailsecff1c_CustomEvent
func callbackTransactionDetailsecff1c_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewTransactionDetailsFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *TransactionDetails) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTransactionDetailsecff1c_DeleteLater
func callbackTransactionDetailsecff1c_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTransactionDetailsFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *TransactionDetails) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTransactionDetailsecff1c_Destroyed
func callbackTransactionDetailsecff1c_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackTransactionDetailsecff1c_DisconnectNotify
func callbackTransactionDetailsecff1c_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTransactionDetailsFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TransactionDetails) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTransactionDetailsecff1c_Event
func callbackTransactionDetailsecff1c_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionDetailsFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *TransactionDetails) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionDetailsecff1c_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackTransactionDetailsecff1c_EventFilter
func callbackTransactionDetailsecff1c_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTransactionDetailsFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *TransactionDetails) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TransactionDetailsecff1c_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackTransactionDetailsecff1c_ObjectNameChanged
func callbackTransactionDetailsecff1c_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackTransactionDetailsecff1c_TimerEvent
func callbackTransactionDetailsecff1c_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTransactionDetailsFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *TransactionDetails) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TransactionDetailsecff1c_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

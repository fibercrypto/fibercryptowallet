package address

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

//export callbackAddressDetailsfb4c44_Constructor
func callbackAddressDetailsfb4c44_Constructor(ptr unsafe.Pointer) {
	this := NewAddressDetailsFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackAddressDetailsfb4c44_Address
func callbackAddressDetailsfb4c44_Address(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.AddressDetailsfb4c44_Address(ptr.Pointer()))
	}
	return ""
}

func (ptr *AddressDetails) AddressDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.AddressDetailsfb4c44_AddressDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackAddressDetailsfb4c44_SetAddress
func callbackAddressDetailsfb4c44_SetAddress(ptr unsafe.Pointer, address C.struct_Moc_PackedString) {
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
		C.AddressDetailsfb4c44_SetAddress(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))})
	}
}

func (ptr *AddressDetails) SetAddressDefault(address string) {
	if ptr.Pointer() != nil {
		var addressC *C.char
		if address != "" {
			addressC = C.CString(address)
			defer C.free(unsafe.Pointer(addressC))
		}
		C.AddressDetailsfb4c44_SetAddressDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))})
	}
}

//export callbackAddressDetailsfb4c44_AddressChanged
func callbackAddressDetailsfb4c44_AddressChanged(ptr unsafe.Pointer, address C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addressChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(address))
	}

}

func (ptr *AddressDetails) ConnectAddressChanged(f func(address string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressChanged") {
			C.AddressDetailsfb4c44_ConnectAddressChanged(ptr.Pointer())
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
		C.AddressDetailsfb4c44_DisconnectAddressChanged(ptr.Pointer())
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
		C.AddressDetailsfb4c44_AddressChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: addressC, len: C.longlong(len(address))})
	}
}

//export callbackAddressDetailsfb4c44_AddressSky
func callbackAddressDetailsfb4c44_AddressSky(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "addressSky"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewAddressDetailsFromPointer(ptr).AddressSkyDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *AddressDetails) ConnectAddressSky(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addressSky"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
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

func (ptr *AddressDetails) AddressSky() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.AddressDetailsfb4c44_AddressSky(ptr.Pointer()))
	}
	return ""
}

func (ptr *AddressDetails) AddressSkyDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.AddressDetailsfb4c44_AddressSkyDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackAddressDetailsfb4c44_SetAddressSky
func callbackAddressDetailsfb4c44_SetAddressSky(ptr unsafe.Pointer, addressSky C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setAddressSky"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(addressSky))
	} else {
		NewAddressDetailsFromPointer(ptr).SetAddressSkyDefault(cGoUnpackString(addressSky))
	}
}

func (ptr *AddressDetails) ConnectSetAddressSky(f func(addressSky string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddressSky"); signal != nil {
			f := func(addressSky string) {
				(*(*func(string))(signal))(addressSky)
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

func (ptr *AddressDetails) SetAddressSky(addressSky string) {
	if ptr.Pointer() != nil {
		var addressSkyC *C.char
		if addressSky != "" {
			addressSkyC = C.CString(addressSky)
			defer C.free(unsafe.Pointer(addressSkyC))
		}
		C.AddressDetailsfb4c44_SetAddressSky(ptr.Pointer(), C.struct_Moc_PackedString{data: addressSkyC, len: C.longlong(len(addressSky))})
	}
}

func (ptr *AddressDetails) SetAddressSkyDefault(addressSky string) {
	if ptr.Pointer() != nil {
		var addressSkyC *C.char
		if addressSky != "" {
			addressSkyC = C.CString(addressSky)
			defer C.free(unsafe.Pointer(addressSkyC))
		}
		C.AddressDetailsfb4c44_SetAddressSkyDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: addressSkyC, len: C.longlong(len(addressSky))})
	}
}

//export callbackAddressDetailsfb4c44_AddressSkyChanged
func callbackAddressDetailsfb4c44_AddressSkyChanged(ptr unsafe.Pointer, addressSky C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addressSkyChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(addressSky))
	}

}

func (ptr *AddressDetails) ConnectAddressSkyChanged(f func(addressSky string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressSkyChanged") {
			C.AddressDetailsfb4c44_ConnectAddressSkyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressSkyChanged"); signal != nil {
			f := func(addressSky string) {
				(*(*func(string))(signal))(addressSky)
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
		C.AddressDetailsfb4c44_DisconnectAddressSkyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressSkyChanged")
	}
}

func (ptr *AddressDetails) AddressSkyChanged(addressSky string) {
	if ptr.Pointer() != nil {
		var addressSkyC *C.char
		if addressSky != "" {
			addressSkyC = C.CString(addressSky)
			defer C.free(unsafe.Pointer(addressSkyC))
		}
		C.AddressDetailsfb4c44_AddressSkyChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: addressSkyC, len: C.longlong(len(addressSky))})
	}
}

//export callbackAddressDetailsfb4c44_AddressCoinHours
func callbackAddressDetailsfb4c44_AddressCoinHours(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "addressCoinHours"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewAddressDetailsFromPointer(ptr).AddressCoinHoursDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *AddressDetails) ConnectAddressCoinHours(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addressCoinHours"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
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

func (ptr *AddressDetails) AddressCoinHours() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.AddressDetailsfb4c44_AddressCoinHours(ptr.Pointer()))
	}
	return ""
}

func (ptr *AddressDetails) AddressCoinHoursDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.AddressDetailsfb4c44_AddressCoinHoursDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackAddressDetailsfb4c44_SetAddressCoinHours
func callbackAddressDetailsfb4c44_SetAddressCoinHours(ptr unsafe.Pointer, addressCoinHours C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setAddressCoinHours"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(addressCoinHours))
	} else {
		NewAddressDetailsFromPointer(ptr).SetAddressCoinHoursDefault(cGoUnpackString(addressCoinHours))
	}
}

func (ptr *AddressDetails) ConnectSetAddressCoinHours(f func(addressCoinHours string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setAddressCoinHours"); signal != nil {
			f := func(addressCoinHours string) {
				(*(*func(string))(signal))(addressCoinHours)
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

func (ptr *AddressDetails) SetAddressCoinHours(addressCoinHours string) {
	if ptr.Pointer() != nil {
		var addressCoinHoursC *C.char
		if addressCoinHours != "" {
			addressCoinHoursC = C.CString(addressCoinHours)
			defer C.free(unsafe.Pointer(addressCoinHoursC))
		}
		C.AddressDetailsfb4c44_SetAddressCoinHours(ptr.Pointer(), C.struct_Moc_PackedString{data: addressCoinHoursC, len: C.longlong(len(addressCoinHours))})
	}
}

func (ptr *AddressDetails) SetAddressCoinHoursDefault(addressCoinHours string) {
	if ptr.Pointer() != nil {
		var addressCoinHoursC *C.char
		if addressCoinHours != "" {
			addressCoinHoursC = C.CString(addressCoinHours)
			defer C.free(unsafe.Pointer(addressCoinHoursC))
		}
		C.AddressDetailsfb4c44_SetAddressCoinHoursDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: addressCoinHoursC, len: C.longlong(len(addressCoinHours))})
	}
}

//export callbackAddressDetailsfb4c44_AddressCoinHoursChanged
func callbackAddressDetailsfb4c44_AddressCoinHoursChanged(ptr unsafe.Pointer, addressCoinHours C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addressCoinHoursChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(addressCoinHours))
	}

}

func (ptr *AddressDetails) ConnectAddressCoinHoursChanged(f func(addressCoinHours string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addressCoinHoursChanged") {
			C.AddressDetailsfb4c44_ConnectAddressCoinHoursChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "addressCoinHoursChanged"); signal != nil {
			f := func(addressCoinHours string) {
				(*(*func(string))(signal))(addressCoinHours)
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
		C.AddressDetailsfb4c44_DisconnectAddressCoinHoursChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressCoinHoursChanged")
	}
}

func (ptr *AddressDetails) AddressCoinHoursChanged(addressCoinHours string) {
	if ptr.Pointer() != nil {
		var addressCoinHoursC *C.char
		if addressCoinHours != "" {
			addressCoinHoursC = C.CString(addressCoinHours)
			defer C.free(unsafe.Pointer(addressCoinHoursC))
		}
		C.AddressDetailsfb4c44_AddressCoinHoursChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: addressCoinHoursC, len: C.longlong(len(addressCoinHours))})
	}
}

func AddressDetails_QRegisterMetaType() int {
	return int(int32(C.AddressDetailsfb4c44_AddressDetailsfb4c44_QRegisterMetaType()))
}

func (ptr *AddressDetails) QRegisterMetaType() int {
	return int(int32(C.AddressDetailsfb4c44_AddressDetailsfb4c44_QRegisterMetaType()))
}

func AddressDetails_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressDetailsfb4c44_AddressDetailsfb4c44_QRegisterMetaType2(typeNameC)))
}

func (ptr *AddressDetails) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressDetailsfb4c44_AddressDetailsfb4c44_QRegisterMetaType2(typeNameC)))
}

func AddressDetails_QmlRegisterType() int {
	return int(int32(C.AddressDetailsfb4c44_AddressDetailsfb4c44_QmlRegisterType()))
}

func (ptr *AddressDetails) QmlRegisterType() int {
	return int(int32(C.AddressDetailsfb4c44_AddressDetailsfb4c44_QmlRegisterType()))
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
	return int(int32(C.AddressDetailsfb4c44_AddressDetailsfb4c44_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.AddressDetailsfb4c44_AddressDetailsfb4c44_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *AddressDetails) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressDetailsfb4c44___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetailsfb4c44___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressDetails) __children_newList() unsafe.Pointer {
	return C.AddressDetailsfb4c44___children_newList(ptr.Pointer())
}

func (ptr *AddressDetails) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressDetailsfb4c44___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetailsfb4c44___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressDetails) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.AddressDetailsfb4c44___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *AddressDetails) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressDetailsfb4c44___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetailsfb4c44___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressDetails) __findChildren_newList() unsafe.Pointer {
	return C.AddressDetailsfb4c44___findChildren_newList(ptr.Pointer())
}

func (ptr *AddressDetails) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressDetailsfb4c44___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetailsfb4c44___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressDetails) __findChildren_newList3() unsafe.Pointer {
	return C.AddressDetailsfb4c44___findChildren_newList3(ptr.Pointer())
}

func (ptr *AddressDetails) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressDetailsfb4c44___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressDetails) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetailsfb4c44___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressDetails) __qFindChildren_newList2() unsafe.Pointer {
	return C.AddressDetailsfb4c44___qFindChildren_newList2(ptr.Pointer())
}

func NewAddressDetails(parent std_core.QObject_ITF) *AddressDetails {
	tmpValue := NewAddressDetailsFromPointer(C.AddressDetailsfb4c44_NewAddressDetails(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackAddressDetailsfb4c44_DestroyAddressDetails
func callbackAddressDetailsfb4c44_DestroyAddressDetails(ptr unsafe.Pointer) {
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
		C.AddressDetailsfb4c44_DestroyAddressDetails(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *AddressDetails) DestroyAddressDetailsDefault() {
	if ptr.Pointer() != nil {
		C.AddressDetailsfb4c44_DestroyAddressDetailsDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressDetailsfb4c44_ChildEvent
func callbackAddressDetailsfb4c44_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewAddressDetailsFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *AddressDetails) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetailsfb4c44_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackAddressDetailsfb4c44_ConnectNotify
func callbackAddressDetailsfb4c44_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressDetailsFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressDetails) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetailsfb4c44_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressDetailsfb4c44_CustomEvent
func callbackAddressDetailsfb4c44_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewAddressDetailsFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *AddressDetails) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetailsfb4c44_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackAddressDetailsfb4c44_DeleteLater
func callbackAddressDetailsfb4c44_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressDetailsFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *AddressDetails) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.AddressDetailsfb4c44_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressDetailsfb4c44_Destroyed
func callbackAddressDetailsfb4c44_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackAddressDetailsfb4c44_DisconnectNotify
func callbackAddressDetailsfb4c44_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressDetailsFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressDetails) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetailsfb4c44_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressDetailsfb4c44_Event
func callbackAddressDetailsfb4c44_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressDetailsFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *AddressDetails) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressDetailsfb4c44_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackAddressDetailsfb4c44_EventFilter
func callbackAddressDetailsfb4c44_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressDetailsFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *AddressDetails) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressDetailsfb4c44_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackAddressDetailsfb4c44_ObjectNameChanged
func callbackAddressDetailsfb4c44_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackAddressDetailsfb4c44_TimerEvent
func callbackAddressDetailsfb4c44_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewAddressDetailsFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *AddressDetails) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressDetailsfb4c44_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
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

//export callbackAddressListfb4c44_Constructor
func callbackAddressListfb4c44_Constructor(ptr unsafe.Pointer) {
	this := NewAddressListFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectAddAddress(this.addAddress)
	this.ConnectRemoveAddress(this.removeAddress)
	this.init()
}

//export callbackAddressListfb4c44_AddAddress
func callbackAddressListfb4c44_AddAddress(ptr unsafe.Pointer, transaction unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addAddress"); signal != nil {
		(*(*func(*AddressDetails))(signal))(NewAddressDetailsFromPointer(transaction))
	}

}

func (ptr *AddressList) ConnectAddAddress(f func(transaction *AddressDetails)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "addAddress") {
			C.AddressListfb4c44_ConnectAddAddress(ptr.Pointer())
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
		C.AddressListfb4c44_DisconnectAddAddress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addAddress")
	}
}

func (ptr *AddressList) AddAddress(transaction AddressDetails_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_AddAddress(ptr.Pointer(), PointerFromAddressDetails(transaction))
	}
}

//export callbackAddressListfb4c44_RemoveAddress
func callbackAddressListfb4c44_RemoveAddress(ptr unsafe.Pointer, index C.int) {
	if signal := qt.GetSignal(ptr, "removeAddress"); signal != nil {
		(*(*func(int))(signal))(int(int32(index)))
	}

}

func (ptr *AddressList) ConnectRemoveAddress(f func(index int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "removeAddress") {
			C.AddressListfb4c44_ConnectRemoveAddress(ptr.Pointer())
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
		C.AddressListfb4c44_DisconnectRemoveAddress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "removeAddress")
	}
}

func (ptr *AddressList) RemoveAddress(index int) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_RemoveAddress(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackAddressListfb4c44_Roles
func callbackAddressListfb4c44_Roles(ptr unsafe.Pointer) unsafe.Pointer {
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
		}(C.AddressListfb4c44_Roles(ptr.Pointer()))
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
		}(C.AddressListfb4c44_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackAddressListfb4c44_SetRoles
func callbackAddressListfb4c44_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
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
		C.AddressListfb4c44_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
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
		C.AddressListfb4c44_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackAddressListfb4c44_RolesChanged
func callbackAddressListfb4c44_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
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
			C.AddressListfb4c44_ConnectRolesChanged(ptr.Pointer())
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
		C.AddressListfb4c44_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *AddressList) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackAddressListfb4c44_Addresses
func callbackAddressListfb4c44_Addresses(ptr unsafe.Pointer) unsafe.Pointer {
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
		}(C.AddressListfb4c44_Addresses(ptr.Pointer()))
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
		}(C.AddressListfb4c44_AddressesDefault(ptr.Pointer()))
	}
	return make([]*AddressDetails, 0)
}

//export callbackAddressListfb4c44_SetAddresses
func callbackAddressListfb4c44_SetAddresses(ptr unsafe.Pointer, addresses C.struct_Moc_PackedList) {
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
		C.AddressListfb4c44_SetAddresses(ptr.Pointer(), func() unsafe.Pointer {
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
		C.AddressListfb4c44_SetAddressesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setAddresses_addresses_newList())
			for _, v := range addresses {
				tmpList.__setAddresses_addresses_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackAddressListfb4c44_AddressesChanged
func callbackAddressListfb4c44_AddressesChanged(ptr unsafe.Pointer, addresses C.struct_Moc_PackedList) {
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
			C.AddressListfb4c44_ConnectAddressesChanged(ptr.Pointer())
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
		C.AddressListfb4c44_DisconnectAddressesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "addressesChanged")
	}
}

func (ptr *AddressList) AddressesChanged(addresses []*AddressDetails) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_AddressesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__addressesChanged_addresses_newList())
			for _, v := range addresses {
				tmpList.__addressesChanged_addresses_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func AddressList_QRegisterMetaType() int {
	return int(int32(C.AddressListfb4c44_AddressListfb4c44_QRegisterMetaType()))
}

func (ptr *AddressList) QRegisterMetaType() int {
	return int(int32(C.AddressListfb4c44_AddressListfb4c44_QRegisterMetaType()))
}

func AddressList_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressListfb4c44_AddressListfb4c44_QRegisterMetaType2(typeNameC)))
}

func (ptr *AddressList) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.AddressListfb4c44_AddressListfb4c44_QRegisterMetaType2(typeNameC)))
}

func AddressList_QmlRegisterType() int {
	return int(int32(C.AddressListfb4c44_AddressListfb4c44_QmlRegisterType()))
}

func (ptr *AddressList) QmlRegisterType() int {
	return int(int32(C.AddressListfb4c44_AddressListfb4c44_QmlRegisterType()))
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
	return int(int32(C.AddressListfb4c44_AddressListfb4c44_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.AddressListfb4c44_AddressListfb4c44_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *AddressList) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressListfb4c44_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____itemData_keyList_newList() unsafe.Pointer {
	return C.AddressListfb4c44_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressListfb4c44_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.AddressListfb4c44_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressListfb4c44_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.AddressListfb4c44_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressListfb4c44___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.AddressListfb4c44___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *AddressList) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressListfb4c44___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.AddressListfb4c44___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *AddressList) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressListfb4c44___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) __dataChanged_roles_newList() unsafe.Pointer {
	return C.AddressListfb4c44___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *AddressList) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressListfb4c44___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *AddressList) __itemData_newList() unsafe.Pointer {
	return C.AddressListfb4c44___itemData_newList(ptr.Pointer())
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
		}(C.AddressListfb4c44___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.AddressListfb4c44___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *AddressList) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.AddressListfb4c44___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *AddressList) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.AddressListfb4c44___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *AddressList) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.AddressListfb4c44___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *AddressList) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressListfb4c44___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __match_newList() unsafe.Pointer {
	return C.AddressListfb4c44___match_newList(ptr.Pointer())
}

func (ptr *AddressList) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressListfb4c44___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __mimeData_indexes_newList() unsafe.Pointer {
	return C.AddressListfb4c44___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *AddressList) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressListfb4c44___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *AddressList) __persistentIndexList_newList() unsafe.Pointer {
	return C.AddressListfb4c44___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *AddressList) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressListfb4c44___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __roleNames_newList() unsafe.Pointer {
	return C.AddressListfb4c44___roleNames_newList(ptr.Pointer())
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
		}(C.AddressListfb4c44___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressListfb4c44___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *AddressList) __setItemData_roles_newList() unsafe.Pointer {
	return C.AddressListfb4c44___setItemData_roles_newList(ptr.Pointer())
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
		}(C.AddressListfb4c44___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressListfb4c44_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.AddressListfb4c44_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressListfb4c44_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.AddressListfb4c44_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressListfb4c44___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressList) __children_newList() unsafe.Pointer {
	return C.AddressListfb4c44___children_newList(ptr.Pointer())
}

func (ptr *AddressList) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressListfb4c44___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.AddressListfb4c44___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *AddressList) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressListfb4c44___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressList) __findChildren_newList() unsafe.Pointer {
	return C.AddressListfb4c44___findChildren_newList(ptr.Pointer())
}

func (ptr *AddressList) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressListfb4c44___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressList) __findChildren_newList3() unsafe.Pointer {
	return C.AddressListfb4c44___findChildren_newList3(ptr.Pointer())
}

func (ptr *AddressList) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.AddressListfb4c44___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *AddressList) __qFindChildren_newList2() unsafe.Pointer {
	return C.AddressListfb4c44___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *AddressList) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressListfb4c44___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __roles_newList() unsafe.Pointer {
	return C.AddressListfb4c44___roles_newList(ptr.Pointer())
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
		}(C.AddressListfb4c44___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressListfb4c44___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __setRoles_roles_newList() unsafe.Pointer {
	return C.AddressListfb4c44___setRoles_roles_newList(ptr.Pointer())
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
		}(C.AddressListfb4c44___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.AddressListfb4c44___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *AddressList) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.AddressListfb4c44___rolesChanged_roles_newList(ptr.Pointer())
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
		}(C.AddressListfb4c44___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *AddressList) __addresses_atList(i int) *AddressDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressDetailsFromPointer(C.AddressListfb4c44___addresses_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __addresses_setList(i AddressDetails_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___addresses_setList(ptr.Pointer(), PointerFromAddressDetails(i))
	}
}

func (ptr *AddressList) __addresses_newList() unsafe.Pointer {
	return C.AddressListfb4c44___addresses_newList(ptr.Pointer())
}

func (ptr *AddressList) __setAddresses_addresses_atList(i int) *AddressDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressDetailsFromPointer(C.AddressListfb4c44___setAddresses_addresses_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __setAddresses_addresses_setList(i AddressDetails_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___setAddresses_addresses_setList(ptr.Pointer(), PointerFromAddressDetails(i))
	}
}

func (ptr *AddressList) __setAddresses_addresses_newList() unsafe.Pointer {
	return C.AddressListfb4c44___setAddresses_addresses_newList(ptr.Pointer())
}

func (ptr *AddressList) __addressesChanged_addresses_atList(i int) *AddressDetails {
	if ptr.Pointer() != nil {
		tmpValue := NewAddressDetailsFromPointer(C.AddressListfb4c44___addressesChanged_addresses_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *AddressList) __addressesChanged_addresses_setList(i AddressDetails_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44___addressesChanged_addresses_setList(ptr.Pointer(), PointerFromAddressDetails(i))
	}
}

func (ptr *AddressList) __addressesChanged_addresses_newList() unsafe.Pointer {
	return C.AddressListfb4c44___addressesChanged_addresses_newList(ptr.Pointer())
}

func (ptr *AddressList) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressListfb4c44_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____roles_keyList_newList() unsafe.Pointer {
	return C.AddressListfb4c44_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressListfb4c44_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.AddressListfb4c44_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *AddressList) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressListfb4c44_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *AddressList) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *AddressList) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.AddressListfb4c44_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewAddressList(parent std_core.QObject_ITF) *AddressList {
	tmpValue := NewAddressListFromPointer(C.AddressListfb4c44_NewAddressList(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackAddressListfb4c44_DestroyAddressList
func callbackAddressListfb4c44_DestroyAddressList(ptr unsafe.Pointer) {
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
		C.AddressListfb4c44_DestroyAddressList(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *AddressList) DestroyAddressListDefault() {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_DestroyAddressListDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressListfb4c44_DropMimeData
func callbackAddressListfb4c44_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_Flags
func callbackAddressListfb4c44_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewAddressListFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressList) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.AddressListfb4c44_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackAddressListfb4c44_Index
func callbackAddressListfb4c44_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewAddressListFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *AddressList) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressListfb4c44_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressListfb4c44_Sibling
func callbackAddressListfb4c44_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewAddressListFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *AddressList) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressListfb4c44_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressListfb4c44_Buddy
func callbackAddressListfb4c44_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewAddressListFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressList) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressListfb4c44_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressListfb4c44_CanDropMimeData
func callbackAddressListfb4c44_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_CanFetchMore
func callbackAddressListfb4c44_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_ColumnCount
func callbackAddressListfb4c44_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewAddressListFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *AddressList) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressListfb4c44_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackAddressListfb4c44_ColumnsAboutToBeInserted
func callbackAddressListfb4c44_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressListfb4c44_ColumnsAboutToBeMoved
func callbackAddressListfb4c44_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbackAddressListfb4c44_ColumnsAboutToBeRemoved
func callbackAddressListfb4c44_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressListfb4c44_ColumnsInserted
func callbackAddressListfb4c44_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressListfb4c44_ColumnsMoved
func callbackAddressListfb4c44_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbackAddressListfb4c44_ColumnsRemoved
func callbackAddressListfb4c44_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressListfb4c44_Data
func callbackAddressListfb4c44_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewAddressListFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *AddressList) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressListfb4c44_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackAddressListfb4c44_DataChanged
func callbackAddressListfb4c44_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
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

//export callbackAddressListfb4c44_FetchMore
func callbackAddressListfb4c44_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewAddressListFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *AddressList) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbackAddressListfb4c44_HasChildren
func callbackAddressListfb4c44_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_HeaderData
func callbackAddressListfb4c44_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewAddressListFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *AddressList) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.AddressListfb4c44_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackAddressListfb4c44_HeaderDataChanged
func callbackAddressListfb4c44_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressListfb4c44_InsertColumns
func callbackAddressListfb4c44_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_InsertRows
func callbackAddressListfb4c44_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_ItemData
func callbackAddressListfb4c44_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
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
		}(C.AddressListfb4c44_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbackAddressListfb4c44_LayoutAboutToBeChanged
func callbackAddressListfb4c44_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
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

//export callbackAddressListfb4c44_LayoutChanged
func callbackAddressListfb4c44_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
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

//export callbackAddressListfb4c44_Match
func callbackAddressListfb4c44_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
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
		}(C.AddressListfb4c44_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbackAddressListfb4c44_MimeData
func callbackAddressListfb4c44_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
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
		tmpValue := std_core.NewQMimeDataFromPointer(C.AddressListfb4c44_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
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

//export callbackAddressListfb4c44_MimeTypes
func callbackAddressListfb4c44_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewAddressListFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *AddressList) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.AddressListfb4c44_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbackAddressListfb4c44_ModelAboutToBeReset
func callbackAddressListfb4c44_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackAddressListfb4c44_ModelReset
func callbackAddressListfb4c44_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackAddressListfb4c44_MoveColumns
func callbackAddressListfb4c44_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *AddressList) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_MoveRows
func callbackAddressListfb4c44_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *AddressList) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_Parent
func callbackAddressListfb4c44_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewAddressListFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressList) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.AddressListfb4c44_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackAddressListfb4c44_RemoveColumns
func callbackAddressListfb4c44_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_RemoveRows
func callbackAddressListfb4c44_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *AddressList) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_ResetInternalData
func callbackAddressListfb4c44_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressListFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *AddressList) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackAddressListfb4c44_Revert
func callbackAddressListfb4c44_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressListFromPointer(ptr).RevertDefault()
	}
}

func (ptr *AddressList) RevertDefault() {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_RevertDefault(ptr.Pointer())
	}
}

//export callbackAddressListfb4c44_RoleNames
func callbackAddressListfb4c44_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
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
		}(C.AddressListfb4c44_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbackAddressListfb4c44_RowCount
func callbackAddressListfb4c44_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewAddressListFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *AddressList) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.AddressListfb4c44_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackAddressListfb4c44_RowsAboutToBeInserted
func callbackAddressListfb4c44_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbackAddressListfb4c44_RowsAboutToBeMoved
func callbackAddressListfb4c44_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbackAddressListfb4c44_RowsAboutToBeRemoved
func callbackAddressListfb4c44_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressListfb4c44_RowsInserted
func callbackAddressListfb4c44_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressListfb4c44_RowsMoved
func callbackAddressListfb4c44_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbackAddressListfb4c44_RowsRemoved
func callbackAddressListfb4c44_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbackAddressListfb4c44_SetData
func callbackAddressListfb4c44_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *AddressList) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_SetHeaderData
func callbackAddressListfb4c44_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *AddressList) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_SetItemData
func callbackAddressListfb4c44_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
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
		return int8(C.AddressListfb4c44_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewAddressListFromPointer(NewAddressListFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbackAddressListfb4c44_Sort
func callbackAddressListfb4c44_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewAddressListFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *AddressList) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackAddressListfb4c44_Span
func callbackAddressListfb4c44_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewAddressListFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *AddressList) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.AddressListfb4c44_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackAddressListfb4c44_Submit
func callbackAddressListfb4c44_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).SubmitDefault())))
}

func (ptr *AddressList) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackAddressListfb4c44_SupportedDragActions
func callbackAddressListfb4c44_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewAddressListFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *AddressList) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.AddressListfb4c44_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackAddressListfb4c44_SupportedDropActions
func callbackAddressListfb4c44_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewAddressListFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *AddressList) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.AddressListfb4c44_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackAddressListfb4c44_ChildEvent
func callbackAddressListfb4c44_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewAddressListFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *AddressList) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackAddressListfb4c44_ConnectNotify
func callbackAddressListfb4c44_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressListFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressList) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressListfb4c44_CustomEvent
func callbackAddressListfb4c44_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewAddressListFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *AddressList) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackAddressListfb4c44_DeleteLater
func callbackAddressListfb4c44_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewAddressListFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *AddressList) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackAddressListfb4c44_Destroyed
func callbackAddressListfb4c44_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackAddressListfb4c44_DisconnectNotify
func callbackAddressListfb4c44_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewAddressListFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *AddressList) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackAddressListfb4c44_Event
func callbackAddressListfb4c44_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *AddressList) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_EventFilter
func callbackAddressListfb4c44_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewAddressListFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *AddressList) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.AddressListfb4c44_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackAddressListfb4c44_ObjectNameChanged
func callbackAddressListfb4c44_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackAddressListfb4c44_TimerEvent
func callbackAddressListfb4c44_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewAddressListFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *AddressList) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.AddressListfb4c44_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

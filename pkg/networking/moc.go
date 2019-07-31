package networking

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

type QConnection_ITF interface {
	std_core.QObject_ITF
	QConnection_PTR() *QConnection
}

func (ptr *QConnection) QConnection_PTR() *QConnection {
	return ptr
}

func (ptr *QConnection) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QConnection) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQConnection(ptr QConnection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QConnection_PTR().Pointer()
	}
	return nil
}

func NewQConnectionFromPointer(ptr unsafe.Pointer) (n *QConnection) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QConnection)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QConnection:
			n = deduced

		case *std_core.QObject:
			n = &QConnection{QObject: *deduced}

		default:
			n = new(QConnection)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackQConnection580e15_Constructor
func callbackQConnection580e15_Constructor(ptr unsafe.Pointer) {
	this := NewQConnectionFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackQConnection580e15_Peer
func callbackQConnection580e15_Peer(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "peer"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQConnectionFromPointer(ptr).PeerDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QConnection) ConnectPeer(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "peer"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "peer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "peer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectPeer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "peer")
	}
}

func (ptr *QConnection) Peer() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QConnection580e15_Peer(ptr.Pointer()))
	}
	return ""
}

func (ptr *QConnection) PeerDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QConnection580e15_PeerDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQConnection580e15_SetPeer
func callbackQConnection580e15_SetPeer(ptr unsafe.Pointer, peer C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setPeer"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(peer))
	} else {
		NewQConnectionFromPointer(ptr).SetPeerDefault(cGoUnpackString(peer))
	}
}

func (ptr *QConnection) ConnectSetPeer(f func(peer string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPeer"); signal != nil {
			f := func(peer string) {
				(*(*func(string))(signal))(peer)
				f(peer)
			}
			qt.ConnectSignal(ptr.Pointer(), "setPeer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPeer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectSetPeer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPeer")
	}
}

func (ptr *QConnection) SetPeer(peer string) {
	if ptr.Pointer() != nil {
		var peerC *C.char
		if peer != "" {
			peerC = C.CString(peer)
			defer C.free(unsafe.Pointer(peerC))
		}
		C.QConnection580e15_SetPeer(ptr.Pointer(), C.struct_Moc_PackedString{data: peerC, len: C.longlong(len(peer))})
	}
}

func (ptr *QConnection) SetPeerDefault(peer string) {
	if ptr.Pointer() != nil {
		var peerC *C.char
		if peer != "" {
			peerC = C.CString(peer)
			defer C.free(unsafe.Pointer(peerC))
		}
		C.QConnection580e15_SetPeerDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: peerC, len: C.longlong(len(peer))})
	}
}

//export callbackQConnection580e15_PeerChanged
func callbackQConnection580e15_PeerChanged(ptr unsafe.Pointer, peer C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "peerChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(peer))
	}

}

func (ptr *QConnection) ConnectPeerChanged(f func(peer string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "peerChanged") {
			C.QConnection580e15_ConnectPeerChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "peerChanged"); signal != nil {
			f := func(peer string) {
				(*(*func(string))(signal))(peer)
				f(peer)
			}
			qt.ConnectSignal(ptr.Pointer(), "peerChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "peerChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectPeerChanged() {
	if ptr.Pointer() != nil {
		C.QConnection580e15_DisconnectPeerChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "peerChanged")
	}
}

func (ptr *QConnection) PeerChanged(peer string) {
	if ptr.Pointer() != nil {
		var peerC *C.char
		if peer != "" {
			peerC = C.CString(peer)
			defer C.free(unsafe.Pointer(peerC))
		}
		C.QConnection580e15_PeerChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: peerC, len: C.longlong(len(peer))})
	}
}

//export callbackQConnection580e15_Source
func callbackQConnection580e15_Source(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "source"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQConnectionFromPointer(ptr).SourceDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QConnection) ConnectSource(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "source"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "source", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "source", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "source")
	}
}

func (ptr *QConnection) Source() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QConnection580e15_Source(ptr.Pointer()))
	}
	return ""
}

func (ptr *QConnection) SourceDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QConnection580e15_SourceDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQConnection580e15_SetSource
func callbackQConnection580e15_SetSource(ptr unsafe.Pointer, source C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setSource"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(source))
	} else {
		NewQConnectionFromPointer(ptr).SetSourceDefault(cGoUnpackString(source))
	}
}

func (ptr *QConnection) ConnectSetSource(f func(source string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSource"); signal != nil {
			f := func(source string) {
				(*(*func(string))(signal))(source)
				f(source)
			}
			qt.ConnectSignal(ptr.Pointer(), "setSource", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSource", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectSetSource() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSource")
	}
}

func (ptr *QConnection) SetSource(source string) {
	if ptr.Pointer() != nil {
		var sourceC *C.char
		if source != "" {
			sourceC = C.CString(source)
			defer C.free(unsafe.Pointer(sourceC))
		}
		C.QConnection580e15_SetSource(ptr.Pointer(), C.struct_Moc_PackedString{data: sourceC, len: C.longlong(len(source))})
	}
}

func (ptr *QConnection) SetSourceDefault(source string) {
	if ptr.Pointer() != nil {
		var sourceC *C.char
		if source != "" {
			sourceC = C.CString(source)
			defer C.free(unsafe.Pointer(sourceC))
		}
		C.QConnection580e15_SetSourceDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: sourceC, len: C.longlong(len(source))})
	}
}

//export callbackQConnection580e15_SourceChanged
func callbackQConnection580e15_SourceChanged(ptr unsafe.Pointer, source C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sourceChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(source))
	}

}

func (ptr *QConnection) ConnectSourceChanged(f func(source string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sourceChanged") {
			C.QConnection580e15_ConnectSourceChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sourceChanged"); signal != nil {
			f := func(source string) {
				(*(*func(string))(signal))(source)
				f(source)
			}
			qt.ConnectSignal(ptr.Pointer(), "sourceChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sourceChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectSourceChanged() {
	if ptr.Pointer() != nil {
		C.QConnection580e15_DisconnectSourceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sourceChanged")
	}
}

func (ptr *QConnection) SourceChanged(source string) {
	if ptr.Pointer() != nil {
		var sourceC *C.char
		if source != "" {
			sourceC = C.CString(source)
			defer C.free(unsafe.Pointer(sourceC))
		}
		C.QConnection580e15_SourceChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: sourceC, len: C.longlong(len(source))})
	}
}

//export callbackQConnection580e15_BlockHeight
func callbackQConnection580e15_BlockHeight(ptr unsafe.Pointer) C.ulonglong {
	if signal := qt.GetSignal(ptr, "blockHeight"); signal != nil {
		return C.ulonglong((*(*func() uint64)(signal))())
	}

	return C.ulonglong(NewQConnectionFromPointer(ptr).BlockHeightDefault())
}

func (ptr *QConnection) ConnectBlockHeight(f func() uint64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "blockHeight"); signal != nil {
			f := func() uint64 {
				(*(*func() uint64)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "blockHeight", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "blockHeight", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectBlockHeight() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "blockHeight")
	}
}

func (ptr *QConnection) BlockHeight() uint64 {
	if ptr.Pointer() != nil {
		return uint64(C.QConnection580e15_BlockHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QConnection) BlockHeightDefault() uint64 {
	if ptr.Pointer() != nil {
		return uint64(C.QConnection580e15_BlockHeightDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQConnection580e15_SetBlockHeight
func callbackQConnection580e15_SetBlockHeight(ptr unsafe.Pointer, blockHeight C.ulonglong) {
	if signal := qt.GetSignal(ptr, "setBlockHeight"); signal != nil {
		(*(*func(uint64))(signal))(uint64(blockHeight))
	} else {
		NewQConnectionFromPointer(ptr).SetBlockHeightDefault(uint64(blockHeight))
	}
}

func (ptr *QConnection) ConnectSetBlockHeight(f func(blockHeight uint64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setBlockHeight"); signal != nil {
			f := func(blockHeight uint64) {
				(*(*func(uint64))(signal))(blockHeight)
				f(blockHeight)
			}
			qt.ConnectSignal(ptr.Pointer(), "setBlockHeight", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setBlockHeight", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectSetBlockHeight() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setBlockHeight")
	}
}

func (ptr *QConnection) SetBlockHeight(blockHeight uint64) {
	if ptr.Pointer() != nil {
		C.QConnection580e15_SetBlockHeight(ptr.Pointer(), C.ulonglong(blockHeight))
	}
}

func (ptr *QConnection) SetBlockHeightDefault(blockHeight uint64) {
	if ptr.Pointer() != nil {
		C.QConnection580e15_SetBlockHeightDefault(ptr.Pointer(), C.ulonglong(blockHeight))
	}
}

//export callbackQConnection580e15_BlockHeightChanged
func callbackQConnection580e15_BlockHeightChanged(ptr unsafe.Pointer, blockHeight C.ulonglong) {
	if signal := qt.GetSignal(ptr, "blockHeightChanged"); signal != nil {
		(*(*func(uint64))(signal))(uint64(blockHeight))
	}

}

func (ptr *QConnection) ConnectBlockHeightChanged(f func(blockHeight uint64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "blockHeightChanged") {
			C.QConnection580e15_ConnectBlockHeightChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "blockHeightChanged"); signal != nil {
			f := func(blockHeight uint64) {
				(*(*func(uint64))(signal))(blockHeight)
				f(blockHeight)
			}
			qt.ConnectSignal(ptr.Pointer(), "blockHeightChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "blockHeightChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectBlockHeightChanged() {
	if ptr.Pointer() != nil {
		C.QConnection580e15_DisconnectBlockHeightChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "blockHeightChanged")
	}
}

func (ptr *QConnection) BlockHeightChanged(blockHeight uint64) {
	if ptr.Pointer() != nil {
		C.QConnection580e15_BlockHeightChanged(ptr.Pointer(), C.ulonglong(blockHeight))
	}
}

//export callbackQConnection580e15_LastSeen
func callbackQConnection580e15_LastSeen(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "lastSeen"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQConnectionFromPointer(ptr).LastSeenDefault())
}

func (ptr *QConnection) ConnectLastSeen(f func() int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "lastSeen"); signal != nil {
			f := func() int64 {
				(*(*func() int64)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "lastSeen", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastSeen", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectLastSeen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "lastSeen")
	}
}

func (ptr *QConnection) LastSeen() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QConnection580e15_LastSeen(ptr.Pointer()))
	}
	return 0
}

func (ptr *QConnection) LastSeenDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QConnection580e15_LastSeenDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQConnection580e15_SetLastSeen
func callbackQConnection580e15_SetLastSeen(ptr unsafe.Pointer, lastSeen C.longlong) {
	if signal := qt.GetSignal(ptr, "setLastSeen"); signal != nil {
		(*(*func(int64))(signal))(int64(lastSeen))
	} else {
		NewQConnectionFromPointer(ptr).SetLastSeenDefault(int64(lastSeen))
	}
}

func (ptr *QConnection) ConnectSetLastSeen(f func(lastSeen int64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLastSeen"); signal != nil {
			f := func(lastSeen int64) {
				(*(*func(int64))(signal))(lastSeen)
				f(lastSeen)
			}
			qt.ConnectSignal(ptr.Pointer(), "setLastSeen", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLastSeen", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectSetLastSeen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLastSeen")
	}
}

func (ptr *QConnection) SetLastSeen(lastSeen int64) {
	if ptr.Pointer() != nil {
		C.QConnection580e15_SetLastSeen(ptr.Pointer(), C.longlong(lastSeen))
	}
}

func (ptr *QConnection) SetLastSeenDefault(lastSeen int64) {
	if ptr.Pointer() != nil {
		C.QConnection580e15_SetLastSeenDefault(ptr.Pointer(), C.longlong(lastSeen))
	}
}

//export callbackQConnection580e15_LastSeenChanged
func callbackQConnection580e15_LastSeenChanged(ptr unsafe.Pointer, lastSeen C.longlong) {
	if signal := qt.GetSignal(ptr, "lastSeenChanged"); signal != nil {
		(*(*func(int64))(signal))(int64(lastSeen))
	}

}

func (ptr *QConnection) ConnectLastSeenChanged(f func(lastSeen int64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "lastSeenChanged") {
			C.QConnection580e15_ConnectLastSeenChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "lastSeenChanged"); signal != nil {
			f := func(lastSeen int64) {
				(*(*func(int64))(signal))(lastSeen)
				f(lastSeen)
			}
			qt.ConnectSignal(ptr.Pointer(), "lastSeenChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "lastSeenChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectLastSeenChanged() {
	if ptr.Pointer() != nil {
		C.QConnection580e15_DisconnectLastSeenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "lastSeenChanged")
	}
}

func (ptr *QConnection) LastSeenChanged(lastSeen int64) {
	if ptr.Pointer() != nil {
		C.QConnection580e15_LastSeenChanged(ptr.Pointer(), C.longlong(lastSeen))
	}
}

func QConnection_QRegisterMetaType() int {
	return int(int32(C.QConnection580e15_QConnection580e15_QRegisterMetaType()))
}

func (ptr *QConnection) QRegisterMetaType() int {
	return int(int32(C.QConnection580e15_QConnection580e15_QRegisterMetaType()))
}

func QConnection_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QConnection580e15_QConnection580e15_QRegisterMetaType2(typeNameC)))
}

func (ptr *QConnection) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QConnection580e15_QConnection580e15_QRegisterMetaType2(typeNameC)))
}

func QConnection_QmlRegisterType() int {
	return int(int32(C.QConnection580e15_QConnection580e15_QmlRegisterType()))
}

func (ptr *QConnection) QmlRegisterType() int {
	return int(int32(C.QConnection580e15_QConnection580e15_QmlRegisterType()))
}

func QConnection_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QConnection580e15_QConnection580e15_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QConnection) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QConnection580e15_QConnection580e15_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QConnection) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QConnection580e15___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QConnection) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QConnection580e15___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QConnection) __children_newList() unsafe.Pointer {
	return C.QConnection580e15___children_newList(ptr.Pointer())
}

func (ptr *QConnection) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.QConnection580e15___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QConnection) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QConnection580e15___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QConnection) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QConnection580e15___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QConnection) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QConnection580e15___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QConnection) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QConnection580e15___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QConnection) __findChildren_newList() unsafe.Pointer {
	return C.QConnection580e15___findChildren_newList(ptr.Pointer())
}

func (ptr *QConnection) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QConnection580e15___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QConnection) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QConnection580e15___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QConnection) __findChildren_newList3() unsafe.Pointer {
	return C.QConnection580e15___findChildren_newList3(ptr.Pointer())
}

func (ptr *QConnection) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.QConnection580e15___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QConnection) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QConnection580e15___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QConnection) __qFindChildren_newList2() unsafe.Pointer {
	return C.QConnection580e15___qFindChildren_newList2(ptr.Pointer())
}

func NewQConnection(parent std_core.QObject_ITF) *QConnection {
	tmpValue := NewQConnectionFromPointer(C.QConnection580e15_NewQConnection(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQConnection580e15_DestroyQConnection
func callbackQConnection580e15_DestroyQConnection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QConnection"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQConnectionFromPointer(ptr).DestroyQConnectionDefault()
	}
}

func (ptr *QConnection) ConnectDestroyQConnection(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QConnection"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QConnection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QConnection", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QConnection) DisconnectDestroyQConnection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QConnection")
	}
}

func (ptr *QConnection) DestroyQConnection() {
	if ptr.Pointer() != nil {
		C.QConnection580e15_DestroyQConnection(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QConnection) DestroyQConnectionDefault() {
	if ptr.Pointer() != nil {
		C.QConnection580e15_DestroyQConnectionDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQConnection580e15_ChildEvent
func callbackQConnection580e15_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQConnectionFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QConnection) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QConnection580e15_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackQConnection580e15_ConnectNotify
func callbackQConnection580e15_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQConnectionFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QConnection) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QConnection580e15_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQConnection580e15_CustomEvent
func callbackQConnection580e15_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewQConnectionFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QConnection) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QConnection580e15_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQConnection580e15_DeleteLater
func callbackQConnection580e15_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQConnectionFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QConnection) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QConnection580e15_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQConnection580e15_Destroyed
func callbackQConnection580e15_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackQConnection580e15_DisconnectNotify
func callbackQConnection580e15_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQConnectionFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QConnection) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QConnection580e15_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQConnection580e15_Event
func callbackQConnection580e15_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQConnectionFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QConnection) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QConnection580e15_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQConnection580e15_EventFilter
func callbackQConnection580e15_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQConnectionFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QConnection) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QConnection580e15_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQConnection580e15_ObjectNameChanged
func callbackQConnection580e15_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQConnection580e15_TimerEvent
func callbackQConnection580e15_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQConnectionFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QConnection) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QConnection580e15_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type networkingModel_ITF interface {
	std_core.QAbstractListModel_ITF
	networkingModel_PTR() *networkingModel
}

func (ptr *networkingModel) networkingModel_PTR() *networkingModel {
	return ptr
}

func (ptr *networkingModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractListModel_PTR().Pointer()
	}
	return nil
}

func (ptr *networkingModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractListModel_PTR().SetPointer(p)
	}
}

func PointerFromNetworkingModel(ptr networkingModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.networkingModel_PTR().Pointer()
	}
	return nil
}

func NewNetworkingModelFromPointer(ptr unsafe.Pointer) (n *networkingModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(networkingModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *networkingModel:
			n = deduced

		case *std_core.QAbstractListModel:
			n = &networkingModel{QAbstractListModel: *deduced}

		default:
			n = new(networkingModel)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *networkingModel) Init() { this.init() }

//export callbacknetworkingModel580e15_Constructor
func callbacknetworkingModel580e15_Constructor(ptr unsafe.Pointer) {
	this := NewNetworkingModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbacknetworkingModel580e15_Roles
func callbacknetworkingModel580e15_Roles(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roles"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__roles_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roles_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__roles_newList())
		for k, v := range NewNetworkingModelFromPointer(ptr).RolesDefault() {
			tmpList.__roles_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *networkingModel) ConnectRoles(f func() map[int]*std_core.QByteArray) {
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

func (ptr *networkingModel) DisconnectRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "roles")
	}
}

func (ptr *networkingModel) Roles() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.networkingModel580e15_Roles(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

func (ptr *networkingModel) RolesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i, v := range tmpList.__roles_keyList() {
				out[v] = tmpList.__roles_atList(v, i)
			}
			return out
		}(C.networkingModel580e15_RolesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbacknetworkingModel580e15_SetRoles
func callbacknetworkingModel580e15_SetRoles(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setRoles"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	} else {
		NewNetworkingModelFromPointer(ptr).SetRolesDefault(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i, v := range tmpList.__setRoles_roles_keyList() {
				out[v] = tmpList.__setRoles_roles_atList(v, i)
			}
			return out
		}(roles))
	}
}

func (ptr *networkingModel) ConnectSetRoles(f func(roles map[int]*std_core.QByteArray)) {
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

func (ptr *networkingModel) DisconnectSetRoles() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRoles")
	}
}

func (ptr *networkingModel) SetRoles(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_SetRoles(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *networkingModel) SetRolesDefault(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_SetRolesDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__setRoles_roles_newList())
			for k, v := range roles {
				tmpList.__setRoles_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbacknetworkingModel580e15_RolesChanged
func callbacknetworkingModel580e15_RolesChanged(ptr unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "rolesChanged"); signal != nil {
		(*(*func(map[int]*std_core.QByteArray))(signal))(func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i, v := range tmpList.__rolesChanged_roles_keyList() {
				out[v] = tmpList.__rolesChanged_roles_atList(v, i)
			}
			return out
		}(roles))
	}

}

func (ptr *networkingModel) ConnectRolesChanged(f func(roles map[int]*std_core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rolesChanged") {
			C.networkingModel580e15_ConnectRolesChanged(ptr.Pointer())
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

func (ptr *networkingModel) DisconnectRolesChanged() {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_DisconnectRolesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rolesChanged")
	}
}

func (ptr *networkingModel) RolesChanged(roles map[int]*std_core.QByteArray) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_RolesChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__rolesChanged_roles_newList())
			for k, v := range roles {
				tmpList.__rolesChanged_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbacknetworkingModel580e15_Connections
func callbacknetworkingModel580e15_Connections(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "connections"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__connections_newList())
			for _, v := range (*(*func() []*QConnection)(signal))() {
				tmpList.__connections_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__connections_newList())
		for _, v := range NewNetworkingModelFromPointer(ptr).ConnectionsDefault() {
			tmpList.__connections_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *networkingModel) ConnectConnections(f func() []*QConnection) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "connections"); signal != nil {
			f := func() []*QConnection {
				(*(*func() []*QConnection)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "connections", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connections", unsafe.Pointer(&f))
		}
	}
}

func (ptr *networkingModel) DisconnectConnections() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "connections")
	}
}

func (ptr *networkingModel) Connections() []*QConnection {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*QConnection {
			out := make([]*QConnection, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__connections_atList(i)
			}
			return out
		}(C.networkingModel580e15_Connections(ptr.Pointer()))
	}
	return make([]*QConnection, 0)
}

func (ptr *networkingModel) ConnectionsDefault() []*QConnection {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*QConnection {
			out := make([]*QConnection, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__connections_atList(i)
			}
			return out
		}(C.networkingModel580e15_ConnectionsDefault(ptr.Pointer()))
	}
	return make([]*QConnection, 0)
}

//export callbacknetworkingModel580e15_SetConnections
func callbacknetworkingModel580e15_SetConnections(ptr unsafe.Pointer, connections C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "setConnections"); signal != nil {
		(*(*func([]*QConnection))(signal))(func(l C.struct_Moc_PackedList) []*QConnection {
			out := make([]*QConnection, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setConnections_connections_atList(i)
			}
			return out
		}(connections))
	} else {
		NewNetworkingModelFromPointer(ptr).SetConnectionsDefault(func(l C.struct_Moc_PackedList) []*QConnection {
			out := make([]*QConnection, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__setConnections_connections_atList(i)
			}
			return out
		}(connections))
	}
}

func (ptr *networkingModel) ConnectSetConnections(f func(connections []*QConnection)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setConnections"); signal != nil {
			f := func(connections []*QConnection) {
				(*(*func([]*QConnection))(signal))(connections)
				f(connections)
			}
			qt.ConnectSignal(ptr.Pointer(), "setConnections", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setConnections", unsafe.Pointer(&f))
		}
	}
}

func (ptr *networkingModel) DisconnectSetConnections() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setConnections")
	}
}

func (ptr *networkingModel) SetConnections(connections []*QConnection) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_SetConnections(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__setConnections_connections_newList())
			for _, v := range connections {
				tmpList.__setConnections_connections_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *networkingModel) SetConnectionsDefault(connections []*QConnection) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_SetConnectionsDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__setConnections_connections_newList())
			for _, v := range connections {
				tmpList.__setConnections_connections_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbacknetworkingModel580e15_ConnectionsChanged
func callbacknetworkingModel580e15_ConnectionsChanged(ptr unsafe.Pointer, connections C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "connectionsChanged"); signal != nil {
		(*(*func([]*QConnection))(signal))(func(l C.struct_Moc_PackedList) []*QConnection {
			out := make([]*QConnection, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__connectionsChanged_connections_atList(i)
			}
			return out
		}(connections))
	}

}

func (ptr *networkingModel) ConnectConnectionsChanged(f func(connections []*QConnection)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connectionsChanged") {
			C.networkingModel580e15_ConnectConnectionsChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connectionsChanged"); signal != nil {
			f := func(connections []*QConnection) {
				(*(*func([]*QConnection))(signal))(connections)
				f(connections)
			}
			qt.ConnectSignal(ptr.Pointer(), "connectionsChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectionsChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *networkingModel) DisconnectConnectionsChanged() {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_DisconnectConnectionsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "connectionsChanged")
	}
}

func (ptr *networkingModel) ConnectionsChanged(connections []*QConnection) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_ConnectionsChanged(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__connectionsChanged_connections_newList())
			for _, v := range connections {
				tmpList.__connectionsChanged_connections_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func networkingModel_QRegisterMetaType() int {
	return int(int32(C.networkingModel580e15_networkingModel580e15_QRegisterMetaType()))
}

func (ptr *networkingModel) QRegisterMetaType() int {
	return int(int32(C.networkingModel580e15_networkingModel580e15_QRegisterMetaType()))
}

func networkingModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.networkingModel580e15_networkingModel580e15_QRegisterMetaType2(typeNameC)))
}

func (ptr *networkingModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.networkingModel580e15_networkingModel580e15_QRegisterMetaType2(typeNameC)))
}

func networkingModel_QmlRegisterType() int {
	return int(int32(C.networkingModel580e15_networkingModel580e15_QmlRegisterType()))
}

func (ptr *networkingModel) QmlRegisterType() int {
	return int(int32(C.networkingModel580e15_networkingModel580e15_QmlRegisterType()))
}

func networkingModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.networkingModel580e15_networkingModel580e15_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *networkingModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.networkingModel580e15_networkingModel580e15_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *networkingModel) ____itemData_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.networkingModel580e15_____itemData_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *networkingModel) ____itemData_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_____itemData_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *networkingModel) ____itemData_keyList_newList() unsafe.Pointer {
	return C.networkingModel580e15_____itemData_keyList_newList(ptr.Pointer())
}

func (ptr *networkingModel) ____roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.networkingModel580e15_____roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *networkingModel) ____roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_____roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *networkingModel) ____roleNames_keyList_newList() unsafe.Pointer {
	return C.networkingModel580e15_____roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *networkingModel) ____setItemData_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.networkingModel580e15_____setItemData_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *networkingModel) ____setItemData_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_____setItemData_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *networkingModel) ____setItemData_roles_keyList_newList() unsafe.Pointer {
	return C.networkingModel580e15_____setItemData_roles_keyList_newList(ptr.Pointer())
}

func (ptr *networkingModel) __changePersistentIndexList_from_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.networkingModel580e15___changePersistentIndexList_from_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __changePersistentIndexList_from_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___changePersistentIndexList_from_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *networkingModel) __changePersistentIndexList_from_newList() unsafe.Pointer {
	return C.networkingModel580e15___changePersistentIndexList_from_newList(ptr.Pointer())
}

func (ptr *networkingModel) __changePersistentIndexList_to_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.networkingModel580e15___changePersistentIndexList_to_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __changePersistentIndexList_to_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___changePersistentIndexList_to_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *networkingModel) __changePersistentIndexList_to_newList() unsafe.Pointer {
	return C.networkingModel580e15___changePersistentIndexList_to_newList(ptr.Pointer())
}

func (ptr *networkingModel) __dataChanged_roles_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.networkingModel580e15___dataChanged_roles_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *networkingModel) __dataChanged_roles_setList(i int) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___dataChanged_roles_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *networkingModel) __dataChanged_roles_newList() unsafe.Pointer {
	return C.networkingModel580e15___dataChanged_roles_newList(ptr.Pointer())
}

func (ptr *networkingModel) __itemData_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.networkingModel580e15___itemData_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __itemData_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___itemData_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *networkingModel) __itemData_newList() unsafe.Pointer {
	return C.networkingModel580e15___itemData_newList(ptr.Pointer())
}

func (ptr *networkingModel) __itemData_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____itemData_keyList_atList(i)
			}
			return out
		}(C.networkingModel580e15___itemData_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *networkingModel) __layoutAboutToBeChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.networkingModel580e15___layoutAboutToBeChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __layoutAboutToBeChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___layoutAboutToBeChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *networkingModel) __layoutAboutToBeChanged_parents_newList() unsafe.Pointer {
	return C.networkingModel580e15___layoutAboutToBeChanged_parents_newList(ptr.Pointer())
}

func (ptr *networkingModel) __layoutChanged_parents_atList(i int) *std_core.QPersistentModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQPersistentModelIndexFromPointer(C.networkingModel580e15___layoutChanged_parents_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QPersistentModelIndex).DestroyQPersistentModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __layoutChanged_parents_setList(i std_core.QPersistentModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___layoutChanged_parents_setList(ptr.Pointer(), std_core.PointerFromQPersistentModelIndex(i))
	}
}

func (ptr *networkingModel) __layoutChanged_parents_newList() unsafe.Pointer {
	return C.networkingModel580e15___layoutChanged_parents_newList(ptr.Pointer())
}

func (ptr *networkingModel) __match_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.networkingModel580e15___match_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __match_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___match_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *networkingModel) __match_newList() unsafe.Pointer {
	return C.networkingModel580e15___match_newList(ptr.Pointer())
}

func (ptr *networkingModel) __mimeData_indexes_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.networkingModel580e15___mimeData_indexes_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __mimeData_indexes_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___mimeData_indexes_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *networkingModel) __mimeData_indexes_newList() unsafe.Pointer {
	return C.networkingModel580e15___mimeData_indexes_newList(ptr.Pointer())
}

func (ptr *networkingModel) __persistentIndexList_atList(i int) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.networkingModel580e15___persistentIndexList_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __persistentIndexList_setList(i std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___persistentIndexList_setList(ptr.Pointer(), std_core.PointerFromQModelIndex(i))
	}
}

func (ptr *networkingModel) __persistentIndexList_newList() unsafe.Pointer {
	return C.networkingModel580e15___persistentIndexList_newList(ptr.Pointer())
}

func (ptr *networkingModel) __roleNames_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.networkingModel580e15___roleNames_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __roleNames_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___roleNames_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *networkingModel) __roleNames_newList() unsafe.Pointer {
	return C.networkingModel580e15___roleNames_newList(ptr.Pointer())
}

func (ptr *networkingModel) __roleNames_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roleNames_keyList_atList(i)
			}
			return out
		}(C.networkingModel580e15___roleNames_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *networkingModel) __setItemData_roles_atList(v int, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.networkingModel580e15___setItemData_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __setItemData_roles_setList(key int, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___setItemData_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQVariant(i))
	}
}

func (ptr *networkingModel) __setItemData_roles_newList() unsafe.Pointer {
	return C.networkingModel580e15___setItemData_roles_newList(ptr.Pointer())
}

func (ptr *networkingModel) __setItemData_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setItemData_roles_keyList_atList(i)
			}
			return out
		}(C.networkingModel580e15___setItemData_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *networkingModel) ____doSetRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.networkingModel580e15_____doSetRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *networkingModel) ____doSetRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_____doSetRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *networkingModel) ____doSetRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.networkingModel580e15_____doSetRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *networkingModel) ____setRoleNames_roleNames_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.networkingModel580e15_____setRoleNames_roleNames_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *networkingModel) ____setRoleNames_roleNames_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_____setRoleNames_roleNames_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *networkingModel) ____setRoleNames_roleNames_keyList_newList() unsafe.Pointer {
	return C.networkingModel580e15_____setRoleNames_roleNames_keyList_newList(ptr.Pointer())
}

func (ptr *networkingModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.networkingModel580e15___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *networkingModel) __children_newList() unsafe.Pointer {
	return C.networkingModel580e15___children_newList(ptr.Pointer())
}

func (ptr *networkingModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.networkingModel580e15___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *networkingModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.networkingModel580e15___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *networkingModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.networkingModel580e15___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *networkingModel) __findChildren_newList() unsafe.Pointer {
	return C.networkingModel580e15___findChildren_newList(ptr.Pointer())
}

func (ptr *networkingModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.networkingModel580e15___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *networkingModel) __findChildren_newList3() unsafe.Pointer {
	return C.networkingModel580e15___findChildren_newList3(ptr.Pointer())
}

func (ptr *networkingModel) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.networkingModel580e15___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *networkingModel) __qFindChildren_newList2() unsafe.Pointer {
	return C.networkingModel580e15___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *networkingModel) __roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.networkingModel580e15___roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *networkingModel) __roles_newList() unsafe.Pointer {
	return C.networkingModel580e15___roles_newList(ptr.Pointer())
}

func (ptr *networkingModel) __roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____roles_keyList_atList(i)
			}
			return out
		}(C.networkingModel580e15___roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *networkingModel) __setRoles_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.networkingModel580e15___setRoles_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __setRoles_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___setRoles_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *networkingModel) __setRoles_roles_newList() unsafe.Pointer {
	return C.networkingModel580e15___setRoles_roles_newList(ptr.Pointer())
}

func (ptr *networkingModel) __setRoles_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____setRoles_roles_keyList_atList(i)
			}
			return out
		}(C.networkingModel580e15___setRoles_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *networkingModel) __rolesChanged_roles_atList(v int, i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.networkingModel580e15___rolesChanged_roles_atList(ptr.Pointer(), C.int(int32(v)), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __rolesChanged_roles_setList(key int, i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___rolesChanged_roles_setList(ptr.Pointer(), C.int(int32(key)), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *networkingModel) __rolesChanged_roles_newList() unsafe.Pointer {
	return C.networkingModel580e15___rolesChanged_roles_newList(ptr.Pointer())
}

func (ptr *networkingModel) __rolesChanged_roles_keyList() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____rolesChanged_roles_keyList_atList(i)
			}
			return out
		}(C.networkingModel580e15___rolesChanged_roles_keyList(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *networkingModel) __connections_atList(i int) *QConnection {
	if ptr.Pointer() != nil {
		tmpValue := NewQConnectionFromPointer(C.networkingModel580e15___connections_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __connections_setList(i QConnection_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___connections_setList(ptr.Pointer(), PointerFromQConnection(i))
	}
}

func (ptr *networkingModel) __connections_newList() unsafe.Pointer {
	return C.networkingModel580e15___connections_newList(ptr.Pointer())
}

func (ptr *networkingModel) __setConnections_connections_atList(i int) *QConnection {
	if ptr.Pointer() != nil {
		tmpValue := NewQConnectionFromPointer(C.networkingModel580e15___setConnections_connections_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __setConnections_connections_setList(i QConnection_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___setConnections_connections_setList(ptr.Pointer(), PointerFromQConnection(i))
	}
}

func (ptr *networkingModel) __setConnections_connections_newList() unsafe.Pointer {
	return C.networkingModel580e15___setConnections_connections_newList(ptr.Pointer())
}

func (ptr *networkingModel) __connectionsChanged_connections_atList(i int) *QConnection {
	if ptr.Pointer() != nil {
		tmpValue := NewQConnectionFromPointer(C.networkingModel580e15___connectionsChanged_connections_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *networkingModel) __connectionsChanged_connections_setList(i QConnection_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15___connectionsChanged_connections_setList(ptr.Pointer(), PointerFromQConnection(i))
	}
}

func (ptr *networkingModel) __connectionsChanged_connections_newList() unsafe.Pointer {
	return C.networkingModel580e15___connectionsChanged_connections_newList(ptr.Pointer())
}

func (ptr *networkingModel) ____roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.networkingModel580e15_____roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *networkingModel) ____roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_____roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *networkingModel) ____roles_keyList_newList() unsafe.Pointer {
	return C.networkingModel580e15_____roles_keyList_newList(ptr.Pointer())
}

func (ptr *networkingModel) ____setRoles_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.networkingModel580e15_____setRoles_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *networkingModel) ____setRoles_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_____setRoles_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *networkingModel) ____setRoles_roles_keyList_newList() unsafe.Pointer {
	return C.networkingModel580e15_____setRoles_roles_keyList_newList(ptr.Pointer())
}

func (ptr *networkingModel) ____rolesChanged_roles_keyList_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.networkingModel580e15_____rolesChanged_roles_keyList_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *networkingModel) ____rolesChanged_roles_keyList_setList(i int) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_____rolesChanged_roles_keyList_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *networkingModel) ____rolesChanged_roles_keyList_newList() unsafe.Pointer {
	return C.networkingModel580e15_____rolesChanged_roles_keyList_newList(ptr.Pointer())
}

func NewNetworkingModel(parent std_core.QObject_ITF) *networkingModel {
	tmpValue := NewNetworkingModelFromPointer(C.networkingModel580e15_NewNetworkingModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbacknetworkingModel580e15_DestroyNetworkingModel
func callbacknetworkingModel580e15_DestroyNetworkingModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~networkingModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNetworkingModelFromPointer(ptr).DestroyNetworkingModelDefault()
	}
}

func (ptr *networkingModel) ConnectDestroyNetworkingModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~networkingModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~networkingModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~networkingModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *networkingModel) DisconnectDestroyNetworkingModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~networkingModel")
	}
}

func (ptr *networkingModel) DestroyNetworkingModel() {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_DestroyNetworkingModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *networkingModel) DestroyNetworkingModelDefault() {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_DestroyNetworkingModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbacknetworkingModel580e15_DropMimeData
func callbacknetworkingModel580e15_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).DropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *networkingModel) DropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_DropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_Flags
func callbacknetworkingModel580e15_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "flags"); signal != nil {
		return C.longlong((*(*func(*std_core.QModelIndex) std_core.Qt__ItemFlag)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewNetworkingModelFromPointer(ptr).FlagsDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *networkingModel) FlagsDefault(index std_core.QModelIndex_ITF) std_core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return std_core.Qt__ItemFlag(C.networkingModel580e15_FlagsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbacknetworkingModel580e15_Index
func callbacknetworkingModel580e15_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "index"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
	}

	return std_core.PointerFromQModelIndex(NewNetworkingModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))
}

func (ptr *networkingModel) IndexDefault(row int, column int, parent std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.networkingModel580e15_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbacknetworkingModel580e15_Sibling
func callbacknetworkingModel580e15_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sibling"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(int, int, *std_core.QModelIndex) *std_core.QModelIndex)(signal))(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
	}

	return std_core.PointerFromQModelIndex(NewNetworkingModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(idx)))
}

func (ptr *networkingModel) SiblingDefault(row int, column int, idx std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.networkingModel580e15_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbacknetworkingModel580e15_Buddy
func callbacknetworkingModel580e15_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "buddy"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewNetworkingModelFromPointer(ptr).BuddyDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *networkingModel) BuddyDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.networkingModel580e15_BuddyDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbacknetworkingModel580e15_CanDropMimeData
func callbacknetworkingModel580e15_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData, std_core.Qt__DropAction, int, int, *std_core.QModelIndex) bool)(signal))(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).CanDropMimeDataDefault(std_core.NewQMimeDataFromPointer(data), std_core.Qt__DropAction(action), int(int32(row)), int(int32(column)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *networkingModel) CanDropMimeDataDefault(data std_core.QMimeData_ITF, action std_core.Qt__DropAction, row int, column int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_CanDropMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_CanFetchMore
func callbacknetworkingModel580e15_CanFetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canFetchMore"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).CanFetchMoreDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *networkingModel) CanFetchMoreDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_CanFetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_ColumnCount
func callbacknetworkingModel580e15_ColumnCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "columnCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewNetworkingModelFromPointer(ptr).ColumnCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *networkingModel) ColumnCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.networkingModel580e15_ColumnCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbacknetworkingModel580e15_ColumnsAboutToBeInserted
func callbacknetworkingModel580e15_ColumnsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbacknetworkingModel580e15_ColumnsAboutToBeMoved
func callbacknetworkingModel580e15_ColumnsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationColumn C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationColumn)))
	}

}

//export callbacknetworkingModel580e15_ColumnsAboutToBeRemoved
func callbacknetworkingModel580e15_ColumnsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbacknetworkingModel580e15_ColumnsInserted
func callbacknetworkingModel580e15_ColumnsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbacknetworkingModel580e15_ColumnsMoved
func callbacknetworkingModel580e15_ColumnsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, column C.int) {
	if signal := qt.GetSignal(ptr, "columnsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(column)))
	}

}

//export callbacknetworkingModel580e15_ColumnsRemoved
func callbacknetworkingModel580e15_ColumnsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "columnsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbacknetworkingModel580e15_Data
func callbacknetworkingModel580e15_Data(ptr unsafe.Pointer, index unsafe.Pointer, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "data"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(*std_core.QModelIndex, int) *std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewNetworkingModelFromPointer(ptr).DataDefault(std_core.NewQModelIndexFromPointer(index), int(int32(role))))
}

func (ptr *networkingModel) DataDefault(index std_core.QModelIndex_ITF, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.networkingModel580e15_DataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbacknetworkingModel580e15_DataChanged
func callbacknetworkingModel580e15_DataChanged(ptr unsafe.Pointer, topLeft unsafe.Pointer, bottomRight unsafe.Pointer, roles C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "dataChanged"); signal != nil {
		(*(*func(*std_core.QModelIndex, *std_core.QModelIndex, []int))(signal))(std_core.NewQModelIndexFromPointer(topLeft), std_core.NewQModelIndexFromPointer(bottomRight), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__dataChanged_roles_atList(i)
			}
			return out
		}(roles))
	}

}

//export callbacknetworkingModel580e15_FetchMore
func callbacknetworkingModel580e15_FetchMore(ptr unsafe.Pointer, parent unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "fetchMore"); signal != nil {
		(*(*func(*std_core.QModelIndex))(signal))(std_core.NewQModelIndexFromPointer(parent))
	} else {
		NewNetworkingModelFromPointer(ptr).FetchMoreDefault(std_core.NewQModelIndexFromPointer(parent))
	}
}

func (ptr *networkingModel) FetchMoreDefault(parent std_core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_FetchMoreDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))
	}
}

//export callbacknetworkingModel580e15_HasChildren
func callbacknetworkingModel580e15_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex) bool)(signal))(std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).HasChildrenDefault(std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *networkingModel) HasChildrenDefault(parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_HasChildrenDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_HeaderData
func callbacknetworkingModel580e15_HeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, role C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "headerData"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, std_core.Qt__Orientation, int) *std_core.QVariant)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
	}

	return std_core.PointerFromQVariant(NewNetworkingModelFromPointer(ptr).HeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), int(int32(role))))
}

func (ptr *networkingModel) HeaderDataDefault(section int, orientation std_core.Qt__Orientation, role int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.networkingModel580e15_HeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbacknetworkingModel580e15_HeaderDataChanged
func callbacknetworkingModel580e15_HeaderDataChanged(ptr unsafe.Pointer, orientation C.longlong, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "headerDataChanged"); signal != nil {
		(*(*func(std_core.Qt__Orientation, int, int))(signal))(std_core.Qt__Orientation(orientation), int(int32(first)), int(int32(last)))
	}

}

//export callbacknetworkingModel580e15_InsertColumns
func callbacknetworkingModel580e15_InsertColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).InsertColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *networkingModel) InsertColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_InsertColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_InsertRows
func callbacknetworkingModel580e15_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *networkingModel) InsertRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_ItemData
func callbacknetworkingModel580e15_ItemData(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemData"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__itemData_newList())
			for k, v := range (*(*func(*std_core.QModelIndex) map[int]*std_core.QVariant)(signal))(std_core.NewQModelIndexFromPointer(index)) {
				tmpList.__itemData_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__itemData_newList())
		for k, v := range NewNetworkingModelFromPointer(ptr).ItemDataDefault(std_core.NewQModelIndexFromPointer(index)) {
			tmpList.__itemData_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *networkingModel) ItemDataDefault(index std_core.QModelIndex_ITF) map[int]*std_core.QVariant {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i, v := range tmpList.__itemData_keyList() {
				out[v] = tmpList.__itemData_atList(v, i)
			}
			return out
		}(C.networkingModel580e15_ItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
	}
	return make(map[int]*std_core.QVariant, 0)
}

//export callbacknetworkingModel580e15_LayoutAboutToBeChanged
func callbacknetworkingModel580e15_LayoutAboutToBeChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutAboutToBeChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutAboutToBeChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbacknetworkingModel580e15_LayoutChanged
func callbacknetworkingModel580e15_LayoutChanged(ptr unsafe.Pointer, parents C.struct_Moc_PackedList, hint C.longlong) {
	if signal := qt.GetSignal(ptr, "layoutChanged"); signal != nil {
		(*(*func([]*std_core.QPersistentModelIndex, std_core.QAbstractItemModel__LayoutChangeHint))(signal))(func(l C.struct_Moc_PackedList) []*std_core.QPersistentModelIndex {
			out := make([]*std_core.QPersistentModelIndex, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__layoutChanged_parents_atList(i)
			}
			return out
		}(parents), std_core.QAbstractItemModel__LayoutChangeHint(hint))
	}

}

//export callbacknetworkingModel580e15_Match
func callbacknetworkingModel580e15_Match(ptr unsafe.Pointer, start unsafe.Pointer, role C.int, value unsafe.Pointer, hits C.int, flags C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "match"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__match_newList())
			for _, v := range (*(*func(*std_core.QModelIndex, int, *std_core.QVariant, int, std_core.Qt__MatchFlag) []*std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
				tmpList.__match_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__match_newList())
		for _, v := range NewNetworkingModelFromPointer(ptr).MatchDefault(std_core.NewQModelIndexFromPointer(start), int(int32(role)), std_core.NewQVariantFromPointer(value), int(int32(hits)), std_core.Qt__MatchFlag(flags)) {
			tmpList.__match_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *networkingModel) MatchDefault(start std_core.QModelIndex_ITF, role int, value std_core.QVariant_ITF, hits int, flags std_core.Qt__MatchFlag) []*std_core.QModelIndex {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__match_atList(i)
			}
			return out
		}(C.networkingModel580e15_MatchDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(start), C.int(int32(role)), std_core.PointerFromQVariant(value), C.int(int32(hits)), C.longlong(flags)))
	}
	return make([]*std_core.QModelIndex, 0)
}

//export callbacknetworkingModel580e15_MimeData
func callbacknetworkingModel580e15_MimeData(ptr unsafe.Pointer, indexes C.struct_Moc_PackedList) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "mimeData"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func([]*std_core.QModelIndex) *std_core.QMimeData)(signal))(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
			out := make([]*std_core.QModelIndex, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__mimeData_indexes_atList(i)
			}
			return out
		}(indexes)))
	}

	return std_core.PointerFromQMimeData(NewNetworkingModelFromPointer(ptr).MimeDataDefault(func(l C.struct_Moc_PackedList) []*std_core.QModelIndex {
		out := make([]*std_core.QModelIndex, int(l.len))
		tmpList := NewNetworkingModelFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__mimeData_indexes_atList(i)
		}
		return out
	}(indexes)))
}

func (ptr *networkingModel) MimeDataDefault(indexes []*std_core.QModelIndex) *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.networkingModel580e15_MimeDataDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__mimeData_indexes_newList())
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

//export callbacknetworkingModel580e15_MimeTypes
func callbacknetworkingModel580e15_MimeTypes(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "mimeTypes"); signal != nil {
		tempVal := (*(*func() []string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
	}
	tempVal := NewNetworkingModelFromPointer(ptr).MimeTypesDefault()
	return C.struct_Moc_PackedString{data: C.CString(strings.Join(tempVal, "¡¦!")), len: C.longlong(len(strings.Join(tempVal, "¡¦!")))}
}

func (ptr *networkingModel) MimeTypesDefault() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.networkingModel580e15_MimeTypesDefault(ptr.Pointer())))
	}
	return make([]string, 0)
}

//export callbacknetworkingModel580e15_ModelAboutToBeReset
func callbacknetworkingModel580e15_ModelAboutToBeReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelAboutToBeReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacknetworkingModel580e15_ModelReset
func callbacknetworkingModel580e15_ModelReset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "modelReset"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbacknetworkingModel580e15_MoveColumns
func callbacknetworkingModel580e15_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).MoveColumnsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *networkingModel) MoveColumnsDefault(sourceParent std_core.QModelIndex_ITF, sourceColumn int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_MoveColumnsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_MoveRows
func callbacknetworkingModel580e15_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	if signal := qt.GetSignal(ptr, "moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int) bool)(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).MoveRowsDefault(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *networkingModel) MoveRowsDefault(sourceParent std_core.QModelIndex_ITF, sourceRow int, count int, destinationParent std_core.QModelIndex_ITF, destinationChild int) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_MoveRowsDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), std_core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild)))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_Parent
func callbacknetworkingModel580e15_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "parent"); signal != nil {
		return std_core.PointerFromQModelIndex((*(*func(*std_core.QModelIndex) *std_core.QModelIndex)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQModelIndex(NewNetworkingModelFromPointer(ptr).ParentDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *networkingModel) ParentDefault(index std_core.QModelIndex_ITF) *std_core.QModelIndex {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQModelIndexFromPointer(C.networkingModel580e15_ParentDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbacknetworkingModel580e15_RemoveColumns
func callbacknetworkingModel580e15_RemoveColumns(ptr unsafe.Pointer, column C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).RemoveColumnsDefault(int(int32(column)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *networkingModel) RemoveColumnsDefault(column int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_RemoveColumnsDefault(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_RemoveRows
func callbacknetworkingModel580e15_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, int, *std_core.QModelIndex) bool)(signal))(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), std_core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *networkingModel) RemoveRowsDefault(row int, count int, parent std_core.QModelIndex_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), std_core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_ResetInternalData
func callbacknetworkingModel580e15_ResetInternalData(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resetInternalData"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNetworkingModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *networkingModel) ResetInternalDataDefault() {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbacknetworkingModel580e15_Revert
func callbacknetworkingModel580e15_Revert(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "revert"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNetworkingModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *networkingModel) RevertDefault() {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_RevertDefault(ptr.Pointer())
	}
}

//export callbacknetworkingModel580e15_RoleNames
func callbacknetworkingModel580e15_RoleNames(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "roleNames"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__roleNames_newList())
			for k, v := range (*(*func() map[int]*std_core.QByteArray)(signal))() {
				tmpList.__roleNames_setList(k, v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__roleNames_newList())
		for k, v := range NewNetworkingModelFromPointer(ptr).RoleNamesDefault() {
			tmpList.__roleNames_setList(k, v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *networkingModel) RoleNamesDefault() map[int]*std_core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) map[int]*std_core.QByteArray {
			out := make(map[int]*std_core.QByteArray, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i, v := range tmpList.__roleNames_keyList() {
				out[v] = tmpList.__roleNames_atList(v, i)
			}
			return out
		}(C.networkingModel580e15_RoleNamesDefault(ptr.Pointer()))
	}
	return make(map[int]*std_core.QByteArray, 0)
}

//export callbacknetworkingModel580e15_RowCount
func callbacknetworkingModel580e15_RowCount(ptr unsafe.Pointer, parent unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "rowCount"); signal != nil {
		return C.int(int32((*(*func(*std_core.QModelIndex) int)(signal))(std_core.NewQModelIndexFromPointer(parent))))
	}

	return C.int(int32(NewNetworkingModelFromPointer(ptr).RowCountDefault(std_core.NewQModelIndexFromPointer(parent))))
}

func (ptr *networkingModel) RowCountDefault(parent std_core.QModelIndex_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.networkingModel580e15_RowCountDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbacknetworkingModel580e15_RowsAboutToBeInserted
func callbacknetworkingModel580e15_RowsAboutToBeInserted(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)))
	}

}

//export callbacknetworkingModel580e15_RowsAboutToBeMoved
func callbacknetworkingModel580e15_RowsAboutToBeMoved(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceStart C.int, sourceEnd C.int, destinationParent unsafe.Pointer, destinationRow C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceStart)), int(int32(sourceEnd)), std_core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationRow)))
	}

}

//export callbacknetworkingModel580e15_RowsAboutToBeRemoved
func callbacknetworkingModel580e15_RowsAboutToBeRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsAboutToBeRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbacknetworkingModel580e15_RowsInserted
func callbacknetworkingModel580e15_RowsInserted(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsInserted"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbacknetworkingModel580e15_RowsMoved
func callbacknetworkingModel580e15_RowsMoved(ptr unsafe.Pointer, parent unsafe.Pointer, start C.int, end C.int, destination unsafe.Pointer, row C.int) {
	if signal := qt.GetSignal(ptr, "rowsMoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int, *std_core.QModelIndex, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(start)), int(int32(end)), std_core.NewQModelIndexFromPointer(destination), int(int32(row)))
	}

}

//export callbacknetworkingModel580e15_RowsRemoved
func callbacknetworkingModel580e15_RowsRemoved(ptr unsafe.Pointer, parent unsafe.Pointer, first C.int, last C.int) {
	if signal := qt.GetSignal(ptr, "rowsRemoved"); signal != nil {
		(*(*func(*std_core.QModelIndex, int, int))(signal))(std_core.NewQModelIndexFromPointer(parent), int(int32(first)), int(int32(last)))
	}

}

//export callbacknetworkingModel580e15_SetData
func callbacknetworkingModel580e15_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, *std_core.QVariant, int) bool)(signal))(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).SetDataDefault(std_core.NewQModelIndexFromPointer(index), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *networkingModel) SetDataDefault(index std_core.QModelIndex_ITF, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_SetDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_SetHeaderData
func callbacknetworkingModel580e15_SetHeaderData(ptr unsafe.Pointer, section C.int, orientation C.longlong, value unsafe.Pointer, role C.int) C.char {
	if signal := qt.GetSignal(ptr, "setHeaderData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int, std_core.Qt__Orientation, *std_core.QVariant, int) bool)(signal))(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).SetHeaderDataDefault(int(int32(section)), std_core.Qt__Orientation(orientation), std_core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *networkingModel) SetHeaderDataDefault(section int, orientation std_core.Qt__Orientation, value std_core.QVariant_ITF, role int) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_SetHeaderDataDefault(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), std_core.PointerFromQVariant(value), C.int(int32(role)))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_SetItemData
func callbacknetworkingModel580e15_SetItemData(ptr unsafe.Pointer, index unsafe.Pointer, roles C.struct_Moc_PackedList) C.char {
	if signal := qt.GetSignal(ptr, "setItemData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QModelIndex, map[int]*std_core.QVariant) bool)(signal))(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
			out := make(map[int]*std_core.QVariant, int(l.len))
			tmpList := NewNetworkingModelFromPointer(l.data)
			for i, v := range tmpList.__setItemData_roles_keyList() {
				out[v] = tmpList.__setItemData_roles_atList(v, i)
			}
			return out
		}(roles)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).SetItemDataDefault(std_core.NewQModelIndexFromPointer(index), func(l C.struct_Moc_PackedList) map[int]*std_core.QVariant {
		out := make(map[int]*std_core.QVariant, int(l.len))
		tmpList := NewNetworkingModelFromPointer(l.data)
		for i, v := range tmpList.__setItemData_roles_keyList() {
			out[v] = tmpList.__setItemData_roles_atList(v, i)
		}
		return out
	}(roles)))))
}

func (ptr *networkingModel) SetItemDataDefault(index std_core.QModelIndex_ITF, roles map[int]*std_core.QVariant) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_SetItemDataDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index), func() unsafe.Pointer {
			tmpList := NewNetworkingModelFromPointer(NewNetworkingModelFromPointer(nil).__setItemData_roles_newList())
			for k, v := range roles {
				tmpList.__setItemData_roles_setList(k, v)
			}
			return tmpList.Pointer()
		}())) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_Sort
func callbacknetworkingModel580e15_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	if signal := qt.GetSignal(ptr, "sort"); signal != nil {
		(*(*func(int, std_core.Qt__SortOrder))(signal))(int(int32(column)), std_core.Qt__SortOrder(order))
	} else {
		NewNetworkingModelFromPointer(ptr).SortDefault(int(int32(column)), std_core.Qt__SortOrder(order))
	}
}

func (ptr *networkingModel) SortDefault(column int, order std_core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbacknetworkingModel580e15_Span
func callbacknetworkingModel580e15_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "span"); signal != nil {
		return std_core.PointerFromQSize((*(*func(*std_core.QModelIndex) *std_core.QSize)(signal))(std_core.NewQModelIndexFromPointer(index)))
	}

	return std_core.PointerFromQSize(NewNetworkingModelFromPointer(ptr).SpanDefault(std_core.NewQModelIndexFromPointer(index)))
}

func (ptr *networkingModel) SpanDefault(index std_core.QModelIndex_ITF) *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.networkingModel580e15_SpanDefault(ptr.Pointer(), std_core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbacknetworkingModel580e15_Submit
func callbacknetworkingModel580e15_Submit(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *networkingModel) SubmitDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_SubmitDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_SupportedDragActions
func callbacknetworkingModel580e15_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDragActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewNetworkingModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *networkingModel) SupportedDragActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.networkingModel580e15_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbacknetworkingModel580e15_SupportedDropActions
func callbacknetworkingModel580e15_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "supportedDropActions"); signal != nil {
		return C.longlong((*(*func() std_core.Qt__DropAction)(signal))())
	}

	return C.longlong(NewNetworkingModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *networkingModel) SupportedDropActionsDefault() std_core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return std_core.Qt__DropAction(C.networkingModel580e15_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbacknetworkingModel580e15_ChildEvent
func callbacknetworkingModel580e15_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewNetworkingModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *networkingModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbacknetworkingModel580e15_ConnectNotify
func callbacknetworkingModel580e15_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewNetworkingModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *networkingModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbacknetworkingModel580e15_CustomEvent
func callbacknetworkingModel580e15_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewNetworkingModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *networkingModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbacknetworkingModel580e15_DeleteLater
func callbacknetworkingModel580e15_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewNetworkingModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *networkingModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbacknetworkingModel580e15_Destroyed
func callbacknetworkingModel580e15_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbacknetworkingModel580e15_DisconnectNotify
func callbacknetworkingModel580e15_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewNetworkingModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *networkingModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbacknetworkingModel580e15_Event
func callbacknetworkingModel580e15_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *networkingModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_EventFilter
func callbacknetworkingModel580e15_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewNetworkingModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *networkingModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.networkingModel580e15_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbacknetworkingModel580e15_ObjectNameChanged
func callbacknetworkingModel580e15_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbacknetworkingModel580e15_TimerEvent
func callbacknetworkingModel580e15_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewNetworkingModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *networkingModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.networkingModel580e15_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

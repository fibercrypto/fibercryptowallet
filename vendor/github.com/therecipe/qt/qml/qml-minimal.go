// +build minimal

package qml

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "qml-minimal.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtQml_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtQml_PackedString) []byte {
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

type QJSEngine struct {
	core.QObject
}

type QJSEngine_ITF interface {
	core.QObject_ITF
	QJSEngine_PTR() *QJSEngine
}

func (ptr *QJSEngine) QJSEngine_PTR() *QJSEngine {
	return ptr
}

func (ptr *QJSEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QJSEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQJSEngine(ptr QJSEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSEngine_PTR().Pointer()
	}
	return nil
}

func NewQJSEngineFromPointer(ptr unsafe.Pointer) (n *QJSEngine) {
	n = new(QJSEngine)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QJSEngine__Extension
//QJSEngine::Extension
type QJSEngine__Extension int64

const (
	QJSEngine__TranslationExtension       QJSEngine__Extension = QJSEngine__Extension(0x1)
	QJSEngine__ConsoleExtension           QJSEngine__Extension = QJSEngine__Extension(0x2)
	QJSEngine__GarbageCollectionExtension QJSEngine__Extension = QJSEngine__Extension(0x4)
	QJSEngine__AllExtensions              QJSEngine__Extension = QJSEngine__Extension(0xffffffff)
)

func NewQJSEngine() *QJSEngine {
	tmpValue := NewQJSEngineFromPointer(C.QJSEngine_NewQJSEngine())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQJSEngine2(parent core.QObject_ITF) *QJSEngine {
	tmpValue := NewQJSEngineFromPointer(C.QJSEngine_NewQJSEngine2(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QJSEngine) NewQObject(object core.QObject_ITF) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSEngine_NewQObject(ptr.Pointer(), core.PointerFromQObject(object)))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

//export callbackQJSEngine_DestroyQJSEngine
func callbackQJSEngine_DestroyQJSEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QJSEngine"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQJSEngineFromPointer(ptr).DestroyQJSEngineDefault()
	}
}

func (ptr *QJSEngine) ConnectDestroyQJSEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QJSEngine"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QJSEngine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QJSEngine", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QJSEngine) DisconnectDestroyQJSEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QJSEngine")
	}
}

func (ptr *QJSEngine) DestroyQJSEngine() {
	if ptr.Pointer() != nil {
		C.QJSEngine_DestroyQJSEngine(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QJSEngine) DestroyQJSEngineDefault() {
	if ptr.Pointer() != nil {
		C.QJSEngine_DestroyQJSEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QJSEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QJSEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QJSEngine) __children_newList() unsafe.Pointer {
	return C.QJSEngine___children_newList(ptr.Pointer())
}

func (ptr *QJSEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QJSEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QJSEngine) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QJSEngine___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QJSEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QJSEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QJSEngine) __findChildren_newList() unsafe.Pointer {
	return C.QJSEngine___findChildren_newList(ptr.Pointer())
}

func (ptr *QJSEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QJSEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QJSEngine) __findChildren_newList3() unsafe.Pointer {
	return C.QJSEngine___findChildren_newList3(ptr.Pointer())
}

func (ptr *QJSEngine) __qFindChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QJSEngine___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QJSEngine) __qFindChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine___qFindChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QJSEngine) __qFindChildren_newList2() unsafe.Pointer {
	return C.QJSEngine___qFindChildren_newList2(ptr.Pointer())
}

//export callbackQJSEngine_ChildEvent
func callbackQJSEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QJSEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQJSEngine_ConnectNotify
func callbackQJSEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQJSEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QJSEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQJSEngine_CustomEvent
func callbackQJSEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QJSEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQJSEngine_DeleteLater
func callbackQJSEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQJSEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QJSEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QJSEngine_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQJSEngine_Destroyed
func callbackQJSEngine_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQJSEngine_DisconnectNotify
func callbackQJSEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQJSEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QJSEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQJSEngine_Event
func callbackQJSEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQJSEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QJSEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQJSEngine_EventFilter
func callbackQJSEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQJSEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QJSEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QJSEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQJSEngine_ObjectNameChanged
func callbackQJSEngine_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtQml_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQJSEngine_TimerEvent
func callbackQJSEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQJSEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QJSEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QJSEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QJSValue struct {
	ptr unsafe.Pointer
}

type QJSValue_ITF interface {
	QJSValue_PTR() *QJSValue
}

func (ptr *QJSValue) QJSValue_PTR() *QJSValue {
	return ptr
}

func (ptr *QJSValue) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QJSValue) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQJSValue(ptr QJSValue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSValue_PTR().Pointer()
	}
	return nil
}

func NewQJSValueFromPointer(ptr unsafe.Pointer) (n *QJSValue) {
	n = new(QJSValue)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QJSValue__SpecialValue
//QJSValue::SpecialValue
type QJSValue__SpecialValue int64

const (
	QJSValue__NullValue      QJSValue__SpecialValue = QJSValue__SpecialValue(0)
	QJSValue__UndefinedValue QJSValue__SpecialValue = QJSValue__SpecialValue(1)
)

//go:generate stringer -type=QJSValue__ErrorType
//QJSValue::ErrorType
type QJSValue__ErrorType int64

const (
	QJSValue__NoError        QJSValue__ErrorType = QJSValue__ErrorType(0)
	QJSValue__GenericError   QJSValue__ErrorType = QJSValue__ErrorType(1)
	QJSValue__EvalError      QJSValue__ErrorType = QJSValue__ErrorType(2)
	QJSValue__RangeError     QJSValue__ErrorType = QJSValue__ErrorType(3)
	QJSValue__ReferenceError QJSValue__ErrorType = QJSValue__ErrorType(4)
	QJSValue__SyntaxError    QJSValue__ErrorType = QJSValue__ErrorType(5)
	QJSValue__TypeError      QJSValue__ErrorType = QJSValue__ErrorType(6)
	QJSValue__URIError       QJSValue__ErrorType = QJSValue__ErrorType(7)
)

func NewQJSValue(value QJSValue__SpecialValue) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue(C.longlong(value)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue2(other QJSValue_ITF) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue2(PointerFromQJSValue(other)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue3(other QJSValue_ITF) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue3(PointerFromQJSValue(other)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue4(value bool) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue4(C.char(int8(qt.GoBoolToInt(value)))))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue5(value int) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue5(C.int(int32(value))))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue6(value uint) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue6(C.uint(uint32(value))))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue7(value float64) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue7(C.double(value)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue8(value string) *QJSValue {
	var valueC *C.char
	if value != "" {
		valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
	}
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue8(C.struct_QtQml_PackedString{data: valueC, len: C.longlong(len(value))}))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue9(value core.QLatin1String_ITF) *QJSValue {
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue9(core.PointerFromQLatin1String(value)))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func NewQJSValue10(value string) *QJSValue {
	var valueC *C.char
	if value != "" {
		valueC = C.CString(value)
		defer C.free(unsafe.Pointer(valueC))
	}
	tmpValue := NewQJSValueFromPointer(C.QJSValue_NewQJSValue10(valueC))
	runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
	return tmpValue
}

func (ptr *QJSValue) Property(name string) *QJSValue {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		tmpValue := NewQJSValueFromPointer(C.QJSValue_Property(ptr.Pointer(), C.struct_QtQml_PackedString{data: nameC, len: C.longlong(len(name))}))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) Property2(arrayIndex uint) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSValue_Property2(ptr.Pointer(), C.uint(uint32(arrayIndex))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) ToInt() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QJSValue_ToInt(ptr.Pointer())))
	}
	return 0
}

func (ptr *QJSValue) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QJSValue_ToVariant(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) DestroyQJSValue() {
	if ptr.Pointer() != nil {
		C.QJSValue_DestroyQJSValue(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QJSValue) __call_args_atList(i int) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSValue___call_args_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) __call_args_setList(i QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSValue___call_args_setList(ptr.Pointer(), PointerFromQJSValue(i))
	}
}

func (ptr *QJSValue) __call_args_newList() unsafe.Pointer {
	return C.QJSValue___call_args_newList(ptr.Pointer())
}

func (ptr *QJSValue) __callAsConstructor_args_atList(i int) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSValue___callAsConstructor_args_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) __callAsConstructor_args_setList(i QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSValue___callAsConstructor_args_setList(ptr.Pointer(), PointerFromQJSValue(i))
	}
}

func (ptr *QJSValue) __callAsConstructor_args_newList() unsafe.Pointer {
	return C.QJSValue___callAsConstructor_args_newList(ptr.Pointer())
}

func (ptr *QJSValue) __callWithInstance_args_atList(i int) *QJSValue {
	if ptr.Pointer() != nil {
		tmpValue := NewQJSValueFromPointer(C.QJSValue___callWithInstance_args_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QJSValue).DestroyQJSValue)
		return tmpValue
	}
	return nil
}

func (ptr *QJSValue) __callWithInstance_args_setList(i QJSValue_ITF) {
	if ptr.Pointer() != nil {
		C.QJSValue___callWithInstance_args_setList(ptr.Pointer(), PointerFromQJSValue(i))
	}
}

func (ptr *QJSValue) __callWithInstance_args_newList() unsafe.Pointer {
	return C.QJSValue___callWithInstance_args_newList(ptr.Pointer())
}

//go:generate stringer -type=QQmlAbstractUrlInterceptor__DataType
//QQmlAbstractUrlInterceptor::DataType
type QQmlAbstractUrlInterceptor__DataType int64

const (
	QQmlAbstractUrlInterceptor__QmlFile        QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(0)
	QQmlAbstractUrlInterceptor__JavaScriptFile QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(1)
	QQmlAbstractUrlInterceptor__QmldirFile     QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(2)
	QQmlAbstractUrlInterceptor__UrlString      QQmlAbstractUrlInterceptor__DataType = QQmlAbstractUrlInterceptor__DataType(0x1000)
)

//go:generate stringer -type=QQmlComponent__CompilationMode
//QQmlComponent::CompilationMode
type QQmlComponent__CompilationMode int64

const (
	QQmlComponent__PreferSynchronous QQmlComponent__CompilationMode = QQmlComponent__CompilationMode(0)
	QQmlComponent__Asynchronous      QQmlComponent__CompilationMode = QQmlComponent__CompilationMode(1)
)

//go:generate stringer -type=QQmlComponent__Status
//QQmlComponent::Status
type QQmlComponent__Status int64

const (
	QQmlComponent__Null    QQmlComponent__Status = QQmlComponent__Status(0)
	QQmlComponent__Ready   QQmlComponent__Status = QQmlComponent__Status(1)
	QQmlComponent__Loading QQmlComponent__Status = QQmlComponent__Status(2)
	QQmlComponent__Error   QQmlComponent__Status = QQmlComponent__Status(3)
)

//go:generate stringer -type=QQmlDebuggingEnabler__StartMode
//QQmlDebuggingEnabler::StartMode
type QQmlDebuggingEnabler__StartMode int64

const (
	QQmlDebuggingEnabler__DoNotWaitForClient QQmlDebuggingEnabler__StartMode = QQmlDebuggingEnabler__StartMode(0)
	QQmlDebuggingEnabler__WaitForClient      QQmlDebuggingEnabler__StartMode = QQmlDebuggingEnabler__StartMode(1)
)

type QQmlEngine struct {
	QJSEngine
}

type QQmlEngine_ITF interface {
	QJSEngine_ITF
	QQmlEngine_PTR() *QQmlEngine
}

func (ptr *QQmlEngine) QQmlEngine_PTR() *QQmlEngine {
	return ptr
}

func (ptr *QQmlEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QJSEngine_PTR().SetPointer(p)
	}
}

func PointerFromQQmlEngine(ptr QQmlEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlEngine_PTR().Pointer()
	}
	return nil
}

func NewQQmlEngineFromPointer(ptr unsafe.Pointer) (n *QQmlEngine) {
	n = new(QQmlEngine)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QQmlEngine__ObjectOwnership
//QQmlEngine::ObjectOwnership
type QQmlEngine__ObjectOwnership int64

const (
	QQmlEngine__CppOwnership        QQmlEngine__ObjectOwnership = QQmlEngine__ObjectOwnership(0)
	QQmlEngine__JavaScriptOwnership QQmlEngine__ObjectOwnership = QQmlEngine__ObjectOwnership(1)
)

func NewQQmlEngine(parent core.QObject_ITF) *QQmlEngine {
	tmpValue := NewQQmlEngineFromPointer(C.QQmlEngine_NewQQmlEngine(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QQmlEngine) AddImportPath(path string) {
	if ptr.Pointer() != nil {
		var pathC *C.char
		if path != "" {
			pathC = C.CString(path)
			defer C.free(unsafe.Pointer(pathC))
		}
		C.QQmlEngine_AddImportPath(ptr.Pointer(), C.struct_QtQml_PackedString{data: pathC, len: C.longlong(len(path))})
	}
}

//export callbackQQmlEngine_DestroyQQmlEngine
func callbackQQmlEngine_DestroyQQmlEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QQmlEngine"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQQmlEngineFromPointer(ptr).DestroyQQmlEngineDefault()
	}
}

func (ptr *QQmlEngine) ConnectDestroyQQmlEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QQmlEngine"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QQmlEngine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QQmlEngine", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQmlEngine) DisconnectDestroyQQmlEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QQmlEngine")
	}
}

func (ptr *QQmlEngine) DestroyQQmlEngine() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DestroyQQmlEngine(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QQmlEngine) DestroyQQmlEngineDefault() {
	if ptr.Pointer() != nil {
		C.QQmlEngine_DestroyQQmlEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func QQmlEngine_QmlRegisterType(url core.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QQmlEngine_QQmlEngine_QmlRegisterType(core.PointerFromQUrl(url), uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QQmlEngine) QmlRegisterType(url core.QUrl_ITF, uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QQmlEngine_QQmlEngine_QmlRegisterType(core.PointerFromQUrl(url), uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QQmlEngine) __importPlugin_errors_atList(i int) *QQmlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlErrorFromPointer(C.QQmlEngine___importPlugin_errors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __importPlugin_errors_setList(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___importPlugin_errors_setList(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlEngine) __importPlugin_errors_newList() unsafe.Pointer {
	return C.QQmlEngine___importPlugin_errors_newList(ptr.Pointer())
}

func (ptr *QQmlEngine) __qmlDebug_errors_atList3(i int) *QQmlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlErrorFromPointer(C.QQmlEngine___qmlDebug_errors_atList3(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __qmlDebug_errors_setList3(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___qmlDebug_errors_setList3(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlEngine) __qmlDebug_errors_newList3() unsafe.Pointer {
	return C.QQmlEngine___qmlDebug_errors_newList3(ptr.Pointer())
}

func (ptr *QQmlEngine) __qmlInfo_errors_atList3(i int) *QQmlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlErrorFromPointer(C.QQmlEngine___qmlInfo_errors_atList3(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __qmlInfo_errors_setList3(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___qmlInfo_errors_setList3(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlEngine) __qmlInfo_errors_newList3() unsafe.Pointer {
	return C.QQmlEngine___qmlInfo_errors_newList3(ptr.Pointer())
}

func (ptr *QQmlEngine) __qmlWarning_errors_atList3(i int) *QQmlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlErrorFromPointer(C.QQmlEngine___qmlWarning_errors_atList3(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __qmlWarning_errors_setList3(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___qmlWarning_errors_setList3(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlEngine) __qmlWarning_errors_newList3() unsafe.Pointer {
	return C.QQmlEngine___qmlWarning_errors_newList3(ptr.Pointer())
}

func (ptr *QQmlEngine) __warnings_warnings_atList(i int) *QQmlError {
	if ptr.Pointer() != nil {
		tmpValue := NewQQmlErrorFromPointer(C.QQmlEngine___warnings_warnings_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
		return tmpValue
	}
	return nil
}

func (ptr *QQmlEngine) __warnings_warnings_setList(i QQmlError_ITF) {
	if ptr.Pointer() != nil {
		C.QQmlEngine___warnings_warnings_setList(ptr.Pointer(), PointerFromQQmlError(i))
	}
}

func (ptr *QQmlEngine) __warnings_warnings_newList() unsafe.Pointer {
	return C.QQmlEngine___warnings_warnings_newList(ptr.Pointer())
}

type QQmlError struct {
	ptr unsafe.Pointer
}

type QQmlError_ITF interface {
	QQmlError_PTR() *QQmlError
}

func (ptr *QQmlError) QQmlError_PTR() *QQmlError {
	return ptr
}

func (ptr *QQmlError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlError(ptr QQmlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlError_PTR().Pointer()
	}
	return nil
}

func NewQQmlErrorFromPointer(ptr unsafe.Pointer) (n *QQmlError) {
	n = new(QQmlError)
	n.SetPointer(ptr)
	return
}

func (ptr *QQmlError) DestroyQQmlError() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQQmlError() *QQmlError {
	tmpValue := NewQQmlErrorFromPointer(C.QQmlError_NewQQmlError())
	runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
	return tmpValue
}

func NewQQmlError2(other QQmlError_ITF) *QQmlError {
	tmpValue := NewQQmlErrorFromPointer(C.QQmlError_NewQQmlError2(PointerFromQQmlError(other)))
	runtime.SetFinalizer(tmpValue, (*QQmlError).DestroyQQmlError)
	return tmpValue
}

func (ptr *QQmlError) Column() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QQmlError_Column(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlError) Object() *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QQmlError_Object(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QQmlError) Url() *core.QUrl {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQUrlFromPointer(C.QQmlError_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

//go:generate stringer -type=QQmlImageProviderBase__ImageType
//QQmlImageProviderBase::ImageType
type QQmlImageProviderBase__ImageType int64

const (
	QQmlImageProviderBase__Image         QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(0)
	QQmlImageProviderBase__Pixmap        QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(1)
	QQmlImageProviderBase__Texture       QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(2)
	QQmlImageProviderBase__Invalid       QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(3)
	QQmlImageProviderBase__ImageResponse QQmlImageProviderBase__ImageType = QQmlImageProviderBase__ImageType(4)
)

//go:generate stringer -type=QQmlImageProviderBase__Flag
//QQmlImageProviderBase::Flag
type QQmlImageProviderBase__Flag int64

const (
	QQmlImageProviderBase__ForceAsynchronousImageLoading QQmlImageProviderBase__Flag = QQmlImageProviderBase__Flag(0x01)
)

//go:generate stringer -type=QQmlIncubator__IncubationMode
//QQmlIncubator::IncubationMode
type QQmlIncubator__IncubationMode int64

const (
	QQmlIncubator__Asynchronous         QQmlIncubator__IncubationMode = QQmlIncubator__IncubationMode(0)
	QQmlIncubator__AsynchronousIfNested QQmlIncubator__IncubationMode = QQmlIncubator__IncubationMode(1)
	QQmlIncubator__Synchronous          QQmlIncubator__IncubationMode = QQmlIncubator__IncubationMode(2)
)

//go:generate stringer -type=QQmlIncubator__Status
//QQmlIncubator::Status
type QQmlIncubator__Status int64

const (
	QQmlIncubator__Null    QQmlIncubator__Status = QQmlIncubator__Status(0)
	QQmlIncubator__Ready   QQmlIncubator__Status = QQmlIncubator__Status(1)
	QQmlIncubator__Loading QQmlIncubator__Status = QQmlIncubator__Status(2)
	QQmlIncubator__Error   QQmlIncubator__Status = QQmlIncubator__Status(3)
)

type QQmlParserStatus struct {
	ptr unsafe.Pointer
}

type QQmlParserStatus_ITF interface {
	QQmlParserStatus_PTR() *QQmlParserStatus
}

func (ptr *QQmlParserStatus) QQmlParserStatus_PTR() *QQmlParserStatus {
	return ptr
}

func (ptr *QQmlParserStatus) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QQmlParserStatus) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQQmlParserStatus(ptr QQmlParserStatus_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlParserStatus_PTR().Pointer()
	}
	return nil
}

func NewQQmlParserStatusFromPointer(ptr unsafe.Pointer) (n *QQmlParserStatus) {
	n = new(QQmlParserStatus)
	n.SetPointer(ptr)
	return
}

func (ptr *QQmlParserStatus) DestroyQQmlParserStatus() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQQmlParserStatus_ClassBegin
func callbackQQmlParserStatus_ClassBegin(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "classBegin"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQmlParserStatus) ConnectClassBegin(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "classBegin"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "classBegin", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "classBegin", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQmlParserStatus) DisconnectClassBegin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "classBegin")
	}
}

func (ptr *QQmlParserStatus) ClassBegin() {
	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ClassBegin(ptr.Pointer())
	}
}

//export callbackQQmlParserStatus_ComponentComplete
func callbackQQmlParserStatus_ComponentComplete(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "componentComplete"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QQmlParserStatus) ConnectComponentComplete(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "componentComplete"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "componentComplete", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "componentComplete", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QQmlParserStatus) DisconnectComponentComplete() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "componentComplete")
	}
}

func (ptr *QQmlParserStatus) ComponentComplete() {
	if ptr.Pointer() != nil {
		C.QQmlParserStatus_ComponentComplete(ptr.Pointer())
	}
}

//go:generate stringer -type=QQmlProperty__PropertyTypeCategory
//QQmlProperty::PropertyTypeCategory
type QQmlProperty__PropertyTypeCategory int64

const (
	QQmlProperty__InvalidCategory QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(0)
	QQmlProperty__List            QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(1)
	QQmlProperty__Object          QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(2)
	QQmlProperty__Normal          QQmlProperty__PropertyTypeCategory = QQmlProperty__PropertyTypeCategory(3)
)

//go:generate stringer -type=QQmlProperty__Type
//QQmlProperty::Type
type QQmlProperty__Type int64

const (
	QQmlProperty__Invalid        QQmlProperty__Type = QQmlProperty__Type(0)
	QQmlProperty__Property       QQmlProperty__Type = QQmlProperty__Type(1)
	QQmlProperty__SignalProperty QQmlProperty__Type = QQmlProperty__Type(2)
)

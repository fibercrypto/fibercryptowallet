package blockchain

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

type BlockchainStatusModel_ITF interface {
	std_core.QObject_ITF
	BlockchainStatusModel_PTR() *BlockchainStatusModel
}

func (ptr *BlockchainStatusModel) BlockchainStatusModel_PTR() *BlockchainStatusModel {
	return ptr
}

func (ptr *BlockchainStatusModel) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *BlockchainStatusModel) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromBlockchainStatusModel(ptr BlockchainStatusModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.BlockchainStatusModel_PTR().Pointer()
	}
	return nil
}

func NewBlockchainStatusModelFromPointer(ptr unsafe.Pointer) (n *BlockchainStatusModel) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(BlockchainStatusModel)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *BlockchainStatusModel:
			n = deduced

		case *std_core.QObject:
			n = &BlockchainStatusModel{QObject: *deduced}

		default:
			n = new(BlockchainStatusModel)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *BlockchainStatusModel) Init() { this.init() }

//export callbackBlockchainStatusModel9ae898_Constructor
func callbackBlockchainStatusModel9ae898_Constructor(ptr unsafe.Pointer) {
	this := NewBlockchainStatusModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackBlockchainStatusModel9ae898_NumberOfBlocks
func callbackBlockchainStatusModel9ae898_NumberOfBlocks(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "numberOfBlocks"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewBlockchainStatusModelFromPointer(ptr).NumberOfBlocksDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *BlockchainStatusModel) ConnectNumberOfBlocks(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "numberOfBlocks"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "numberOfBlocks", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "numberOfBlocks", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectNumberOfBlocks() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "numberOfBlocks")
	}
}

func (ptr *BlockchainStatusModel) NumberOfBlocks() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel9ae898_NumberOfBlocks(ptr.Pointer()))
	}
	return ""
}

func (ptr *BlockchainStatusModel) NumberOfBlocksDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel9ae898_NumberOfBlocksDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackBlockchainStatusModel9ae898_SetNumberOfBlocks
func callbackBlockchainStatusModel9ae898_SetNumberOfBlocks(ptr unsafe.Pointer, numberOfBlocks C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setNumberOfBlocks"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(numberOfBlocks))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).SetNumberOfBlocksDefault(cGoUnpackString(numberOfBlocks))
	}
}

func (ptr *BlockchainStatusModel) ConnectSetNumberOfBlocks(f func(numberOfBlocks string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setNumberOfBlocks"); signal != nil {
			f := func(numberOfBlocks string) {
				(*(*func(string))(signal))(numberOfBlocks)
				f(numberOfBlocks)
			}
			qt.ConnectSignal(ptr.Pointer(), "setNumberOfBlocks", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setNumberOfBlocks", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectSetNumberOfBlocks() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setNumberOfBlocks")
	}
}

func (ptr *BlockchainStatusModel) SetNumberOfBlocks(numberOfBlocks string) {
	if ptr.Pointer() != nil {
		var numberOfBlocksC *C.char
		if numberOfBlocks != "" {
			numberOfBlocksC = C.CString(numberOfBlocks)
			defer C.free(unsafe.Pointer(numberOfBlocksC))
		}
		C.BlockchainStatusModel9ae898_SetNumberOfBlocks(ptr.Pointer(), C.struct_Moc_PackedString{data: numberOfBlocksC, len: C.longlong(len(numberOfBlocks))})
	}
}

func (ptr *BlockchainStatusModel) SetNumberOfBlocksDefault(numberOfBlocks string) {
	if ptr.Pointer() != nil {
		var numberOfBlocksC *C.char
		if numberOfBlocks != "" {
			numberOfBlocksC = C.CString(numberOfBlocks)
			defer C.free(unsafe.Pointer(numberOfBlocksC))
		}
		C.BlockchainStatusModel9ae898_SetNumberOfBlocksDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: numberOfBlocksC, len: C.longlong(len(numberOfBlocks))})
	}
}

//export callbackBlockchainStatusModel9ae898_NumberOfBlocksChanged
func callbackBlockchainStatusModel9ae898_NumberOfBlocksChanged(ptr unsafe.Pointer, numberOfBlocks C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "numberOfBlocksChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(numberOfBlocks))
	}

}

func (ptr *BlockchainStatusModel) ConnectNumberOfBlocksChanged(f func(numberOfBlocks string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "numberOfBlocksChanged") {
			C.BlockchainStatusModel9ae898_ConnectNumberOfBlocksChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "numberOfBlocksChanged"); signal != nil {
			f := func(numberOfBlocks string) {
				(*(*func(string))(signal))(numberOfBlocks)
				f(numberOfBlocks)
			}
			qt.ConnectSignal(ptr.Pointer(), "numberOfBlocksChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "numberOfBlocksChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectNumberOfBlocksChanged() {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_DisconnectNumberOfBlocksChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "numberOfBlocksChanged")
	}
}

func (ptr *BlockchainStatusModel) NumberOfBlocksChanged(numberOfBlocks string) {
	if ptr.Pointer() != nil {
		var numberOfBlocksC *C.char
		if numberOfBlocks != "" {
			numberOfBlocksC = C.CString(numberOfBlocks)
			defer C.free(unsafe.Pointer(numberOfBlocksC))
		}
		C.BlockchainStatusModel9ae898_NumberOfBlocksChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: numberOfBlocksC, len: C.longlong(len(numberOfBlocks))})
	}
}

//export callbackBlockchainStatusModel9ae898_TimestampLastBlock
func callbackBlockchainStatusModel9ae898_TimestampLastBlock(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "timestampLastBlock"); signal != nil {
		return std_core.PointerFromQDateTime((*(*func() *std_core.QDateTime)(signal))())
	}

	return std_core.PointerFromQDateTime(NewBlockchainStatusModelFromPointer(ptr).TimestampLastBlockDefault())
}

func (ptr *BlockchainStatusModel) ConnectTimestampLastBlock(f func() *std_core.QDateTime) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "timestampLastBlock"); signal != nil {
			f := func() *std_core.QDateTime {
				(*(*func() *std_core.QDateTime)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "timestampLastBlock", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "timestampLastBlock", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectTimestampLastBlock() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "timestampLastBlock")
	}
}

func (ptr *BlockchainStatusModel) TimestampLastBlock() *std_core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQDateTimeFromPointer(C.BlockchainStatusModel9ae898_TimestampLastBlock(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *BlockchainStatusModel) TimestampLastBlockDefault() *std_core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQDateTimeFromPointer(C.BlockchainStatusModel9ae898_TimestampLastBlockDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

//export callbackBlockchainStatusModel9ae898_SetTimestampLastBlock
func callbackBlockchainStatusModel9ae898_SetTimestampLastBlock(ptr unsafe.Pointer, timestampLastBlock unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setTimestampLastBlock"); signal != nil {
		(*(*func(*std_core.QDateTime))(signal))(std_core.NewQDateTimeFromPointer(timestampLastBlock))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).SetTimestampLastBlockDefault(std_core.NewQDateTimeFromPointer(timestampLastBlock))
	}
}

func (ptr *BlockchainStatusModel) ConnectSetTimestampLastBlock(f func(timestampLastBlock *std_core.QDateTime)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTimestampLastBlock"); signal != nil {
			f := func(timestampLastBlock *std_core.QDateTime) {
				(*(*func(*std_core.QDateTime))(signal))(timestampLastBlock)
				f(timestampLastBlock)
			}
			qt.ConnectSignal(ptr.Pointer(), "setTimestampLastBlock", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTimestampLastBlock", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectSetTimestampLastBlock() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTimestampLastBlock")
	}
}

func (ptr *BlockchainStatusModel) SetTimestampLastBlock(timestampLastBlock std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_SetTimestampLastBlock(ptr.Pointer(), std_core.PointerFromQDateTime(timestampLastBlock))
	}
}

func (ptr *BlockchainStatusModel) SetTimestampLastBlockDefault(timestampLastBlock std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_SetTimestampLastBlockDefault(ptr.Pointer(), std_core.PointerFromQDateTime(timestampLastBlock))
	}
}

//export callbackBlockchainStatusModel9ae898_TimestampLastBlockChanged
func callbackBlockchainStatusModel9ae898_TimestampLastBlockChanged(ptr unsafe.Pointer, timestampLastBlock unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timestampLastBlockChanged"); signal != nil {
		(*(*func(*std_core.QDateTime))(signal))(std_core.NewQDateTimeFromPointer(timestampLastBlock))
	}

}

func (ptr *BlockchainStatusModel) ConnectTimestampLastBlockChanged(f func(timestampLastBlock *std_core.QDateTime)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "timestampLastBlockChanged") {
			C.BlockchainStatusModel9ae898_ConnectTimestampLastBlockChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "timestampLastBlockChanged"); signal != nil {
			f := func(timestampLastBlock *std_core.QDateTime) {
				(*(*func(*std_core.QDateTime))(signal))(timestampLastBlock)
				f(timestampLastBlock)
			}
			qt.ConnectSignal(ptr.Pointer(), "timestampLastBlockChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "timestampLastBlockChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectTimestampLastBlockChanged() {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_DisconnectTimestampLastBlockChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "timestampLastBlockChanged")
	}
}

func (ptr *BlockchainStatusModel) TimestampLastBlockChanged(timestampLastBlock std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_TimestampLastBlockChanged(ptr.Pointer(), std_core.PointerFromQDateTime(timestampLastBlock))
	}
}

//export callbackBlockchainStatusModel9ae898_HashLastBlock
func callbackBlockchainStatusModel9ae898_HashLastBlock(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "hashLastBlock"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewBlockchainStatusModelFromPointer(ptr).HashLastBlockDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *BlockchainStatusModel) ConnectHashLastBlock(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hashLastBlock"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hashLastBlock", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hashLastBlock", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectHashLastBlock() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hashLastBlock")
	}
}

func (ptr *BlockchainStatusModel) HashLastBlock() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel9ae898_HashLastBlock(ptr.Pointer()))
	}
	return ""
}

func (ptr *BlockchainStatusModel) HashLastBlockDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel9ae898_HashLastBlockDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackBlockchainStatusModel9ae898_SetHashLastBlock
func callbackBlockchainStatusModel9ae898_SetHashLastBlock(ptr unsafe.Pointer, hashLastBlock C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setHashLastBlock"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(hashLastBlock))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).SetHashLastBlockDefault(cGoUnpackString(hashLastBlock))
	}
}

func (ptr *BlockchainStatusModel) ConnectSetHashLastBlock(f func(hashLastBlock string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setHashLastBlock"); signal != nil {
			f := func(hashLastBlock string) {
				(*(*func(string))(signal))(hashLastBlock)
				f(hashLastBlock)
			}
			qt.ConnectSignal(ptr.Pointer(), "setHashLastBlock", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setHashLastBlock", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectSetHashLastBlock() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setHashLastBlock")
	}
}

func (ptr *BlockchainStatusModel) SetHashLastBlock(hashLastBlock string) {
	if ptr.Pointer() != nil {
		var hashLastBlockC *C.char
		if hashLastBlock != "" {
			hashLastBlockC = C.CString(hashLastBlock)
			defer C.free(unsafe.Pointer(hashLastBlockC))
		}
		C.BlockchainStatusModel9ae898_SetHashLastBlock(ptr.Pointer(), C.struct_Moc_PackedString{data: hashLastBlockC, len: C.longlong(len(hashLastBlock))})
	}
}

func (ptr *BlockchainStatusModel) SetHashLastBlockDefault(hashLastBlock string) {
	if ptr.Pointer() != nil {
		var hashLastBlockC *C.char
		if hashLastBlock != "" {
			hashLastBlockC = C.CString(hashLastBlock)
			defer C.free(unsafe.Pointer(hashLastBlockC))
		}
		C.BlockchainStatusModel9ae898_SetHashLastBlockDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: hashLastBlockC, len: C.longlong(len(hashLastBlock))})
	}
}

//export callbackBlockchainStatusModel9ae898_HashLastBlockChanged
func callbackBlockchainStatusModel9ae898_HashLastBlockChanged(ptr unsafe.Pointer, hashLastBlock C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "hashLastBlockChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(hashLastBlock))
	}

}

func (ptr *BlockchainStatusModel) ConnectHashLastBlockChanged(f func(hashLastBlock string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hashLastBlockChanged") {
			C.BlockchainStatusModel9ae898_ConnectHashLastBlockChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hashLastBlockChanged"); signal != nil {
			f := func(hashLastBlock string) {
				(*(*func(string))(signal))(hashLastBlock)
				f(hashLastBlock)
			}
			qt.ConnectSignal(ptr.Pointer(), "hashLastBlockChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hashLastBlockChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectHashLastBlockChanged() {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_DisconnectHashLastBlockChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hashLastBlockChanged")
	}
}

func (ptr *BlockchainStatusModel) HashLastBlockChanged(hashLastBlock string) {
	if ptr.Pointer() != nil {
		var hashLastBlockC *C.char
		if hashLastBlock != "" {
			hashLastBlockC = C.CString(hashLastBlock)
			defer C.free(unsafe.Pointer(hashLastBlockC))
		}
		C.BlockchainStatusModel9ae898_HashLastBlockChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: hashLastBlockC, len: C.longlong(len(hashLastBlock))})
	}
}

//export callbackBlockchainStatusModel9ae898_CurrentSkySupply
func callbackBlockchainStatusModel9ae898_CurrentSkySupply(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "currentSkySupply"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewBlockchainStatusModelFromPointer(ptr).CurrentSkySupplyDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *BlockchainStatusModel) ConnectCurrentSkySupply(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "currentSkySupply"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "currentSkySupply", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "currentSkySupply", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectCurrentSkySupply() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "currentSkySupply")
	}
}

func (ptr *BlockchainStatusModel) CurrentSkySupply() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel9ae898_CurrentSkySupply(ptr.Pointer()))
	}
	return ""
}

func (ptr *BlockchainStatusModel) CurrentSkySupplyDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel9ae898_CurrentSkySupplyDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackBlockchainStatusModel9ae898_SetCurrentSkySupply
func callbackBlockchainStatusModel9ae898_SetCurrentSkySupply(ptr unsafe.Pointer, currentSkySupply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setCurrentSkySupply"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(currentSkySupply))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).SetCurrentSkySupplyDefault(cGoUnpackString(currentSkySupply))
	}
}

func (ptr *BlockchainStatusModel) ConnectSetCurrentSkySupply(f func(currentSkySupply string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCurrentSkySupply"); signal != nil {
			f := func(currentSkySupply string) {
				(*(*func(string))(signal))(currentSkySupply)
				f(currentSkySupply)
			}
			qt.ConnectSignal(ptr.Pointer(), "setCurrentSkySupply", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setCurrentSkySupply", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectSetCurrentSkySupply() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setCurrentSkySupply")
	}
}

func (ptr *BlockchainStatusModel) SetCurrentSkySupply(currentSkySupply string) {
	if ptr.Pointer() != nil {
		var currentSkySupplyC *C.char
		if currentSkySupply != "" {
			currentSkySupplyC = C.CString(currentSkySupply)
			defer C.free(unsafe.Pointer(currentSkySupplyC))
		}
		C.BlockchainStatusModel9ae898_SetCurrentSkySupply(ptr.Pointer(), C.struct_Moc_PackedString{data: currentSkySupplyC, len: C.longlong(len(currentSkySupply))})
	}
}

func (ptr *BlockchainStatusModel) SetCurrentSkySupplyDefault(currentSkySupply string) {
	if ptr.Pointer() != nil {
		var currentSkySupplyC *C.char
		if currentSkySupply != "" {
			currentSkySupplyC = C.CString(currentSkySupply)
			defer C.free(unsafe.Pointer(currentSkySupplyC))
		}
		C.BlockchainStatusModel9ae898_SetCurrentSkySupplyDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: currentSkySupplyC, len: C.longlong(len(currentSkySupply))})
	}
}

//export callbackBlockchainStatusModel9ae898_CurrentSkySupplyChanged
func callbackBlockchainStatusModel9ae898_CurrentSkySupplyChanged(ptr unsafe.Pointer, currentSkySupply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "currentSkySupplyChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(currentSkySupply))
	}

}

func (ptr *BlockchainStatusModel) ConnectCurrentSkySupplyChanged(f func(currentSkySupply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "currentSkySupplyChanged") {
			C.BlockchainStatusModel9ae898_ConnectCurrentSkySupplyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "currentSkySupplyChanged"); signal != nil {
			f := func(currentSkySupply string) {
				(*(*func(string))(signal))(currentSkySupply)
				f(currentSkySupply)
			}
			qt.ConnectSignal(ptr.Pointer(), "currentSkySupplyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "currentSkySupplyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectCurrentSkySupplyChanged() {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_DisconnectCurrentSkySupplyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "currentSkySupplyChanged")
	}
}

func (ptr *BlockchainStatusModel) CurrentSkySupplyChanged(currentSkySupply string) {
	if ptr.Pointer() != nil {
		var currentSkySupplyC *C.char
		if currentSkySupply != "" {
			currentSkySupplyC = C.CString(currentSkySupply)
			defer C.free(unsafe.Pointer(currentSkySupplyC))
		}
		C.BlockchainStatusModel9ae898_CurrentSkySupplyChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: currentSkySupplyC, len: C.longlong(len(currentSkySupply))})
	}
}

//export callbackBlockchainStatusModel9ae898_TotalSkySupply
func callbackBlockchainStatusModel9ae898_TotalSkySupply(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "totalSkySupply"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewBlockchainStatusModelFromPointer(ptr).TotalSkySupplyDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *BlockchainStatusModel) ConnectTotalSkySupply(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "totalSkySupply"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "totalSkySupply", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "totalSkySupply", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectTotalSkySupply() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "totalSkySupply")
	}
}

func (ptr *BlockchainStatusModel) TotalSkySupply() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel9ae898_TotalSkySupply(ptr.Pointer()))
	}
	return ""
}

func (ptr *BlockchainStatusModel) TotalSkySupplyDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel9ae898_TotalSkySupplyDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackBlockchainStatusModel9ae898_SetTotalSkySupply
func callbackBlockchainStatusModel9ae898_SetTotalSkySupply(ptr unsafe.Pointer, totalSkySupply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTotalSkySupply"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(totalSkySupply))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).SetTotalSkySupplyDefault(cGoUnpackString(totalSkySupply))
	}
}

func (ptr *BlockchainStatusModel) ConnectSetTotalSkySupply(f func(totalSkySupply string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTotalSkySupply"); signal != nil {
			f := func(totalSkySupply string) {
				(*(*func(string))(signal))(totalSkySupply)
				f(totalSkySupply)
			}
			qt.ConnectSignal(ptr.Pointer(), "setTotalSkySupply", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTotalSkySupply", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectSetTotalSkySupply() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTotalSkySupply")
	}
}

func (ptr *BlockchainStatusModel) SetTotalSkySupply(totalSkySupply string) {
	if ptr.Pointer() != nil {
		var totalSkySupplyC *C.char
		if totalSkySupply != "" {
			totalSkySupplyC = C.CString(totalSkySupply)
			defer C.free(unsafe.Pointer(totalSkySupplyC))
		}
		C.BlockchainStatusModel9ae898_SetTotalSkySupply(ptr.Pointer(), C.struct_Moc_PackedString{data: totalSkySupplyC, len: C.longlong(len(totalSkySupply))})
	}
}

func (ptr *BlockchainStatusModel) SetTotalSkySupplyDefault(totalSkySupply string) {
	if ptr.Pointer() != nil {
		var totalSkySupplyC *C.char
		if totalSkySupply != "" {
			totalSkySupplyC = C.CString(totalSkySupply)
			defer C.free(unsafe.Pointer(totalSkySupplyC))
		}
		C.BlockchainStatusModel9ae898_SetTotalSkySupplyDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: totalSkySupplyC, len: C.longlong(len(totalSkySupply))})
	}
}

//export callbackBlockchainStatusModel9ae898_TotalSkySupplyChanged
func callbackBlockchainStatusModel9ae898_TotalSkySupplyChanged(ptr unsafe.Pointer, totalSkySupply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "totalSkySupplyChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(totalSkySupply))
	}

}

func (ptr *BlockchainStatusModel) ConnectTotalSkySupplyChanged(f func(totalSkySupply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "totalSkySupplyChanged") {
			C.BlockchainStatusModel9ae898_ConnectTotalSkySupplyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "totalSkySupplyChanged"); signal != nil {
			f := func(totalSkySupply string) {
				(*(*func(string))(signal))(totalSkySupply)
				f(totalSkySupply)
			}
			qt.ConnectSignal(ptr.Pointer(), "totalSkySupplyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "totalSkySupplyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectTotalSkySupplyChanged() {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_DisconnectTotalSkySupplyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "totalSkySupplyChanged")
	}
}

func (ptr *BlockchainStatusModel) TotalSkySupplyChanged(totalSkySupply string) {
	if ptr.Pointer() != nil {
		var totalSkySupplyC *C.char
		if totalSkySupply != "" {
			totalSkySupplyC = C.CString(totalSkySupply)
			defer C.free(unsafe.Pointer(totalSkySupplyC))
		}
		C.BlockchainStatusModel9ae898_TotalSkySupplyChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: totalSkySupplyC, len: C.longlong(len(totalSkySupply))})
	}
}

//export callbackBlockchainStatusModel9ae898_CurrentCoinHoursSupply
func callbackBlockchainStatusModel9ae898_CurrentCoinHoursSupply(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "currentCoinHoursSupply"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewBlockchainStatusModelFromPointer(ptr).CurrentCoinHoursSupplyDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *BlockchainStatusModel) ConnectCurrentCoinHoursSupply(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "currentCoinHoursSupply"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "currentCoinHoursSupply", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "currentCoinHoursSupply", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectCurrentCoinHoursSupply() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "currentCoinHoursSupply")
	}
}

func (ptr *BlockchainStatusModel) CurrentCoinHoursSupply() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel9ae898_CurrentCoinHoursSupply(ptr.Pointer()))
	}
	return ""
}

func (ptr *BlockchainStatusModel) CurrentCoinHoursSupplyDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel9ae898_CurrentCoinHoursSupplyDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackBlockchainStatusModel9ae898_SetCurrentCoinHoursSupply
func callbackBlockchainStatusModel9ae898_SetCurrentCoinHoursSupply(ptr unsafe.Pointer, currentCoinHoursSupply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setCurrentCoinHoursSupply"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(currentCoinHoursSupply))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).SetCurrentCoinHoursSupplyDefault(cGoUnpackString(currentCoinHoursSupply))
	}
}

func (ptr *BlockchainStatusModel) ConnectSetCurrentCoinHoursSupply(f func(currentCoinHoursSupply string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCurrentCoinHoursSupply"); signal != nil {
			f := func(currentCoinHoursSupply string) {
				(*(*func(string))(signal))(currentCoinHoursSupply)
				f(currentCoinHoursSupply)
			}
			qt.ConnectSignal(ptr.Pointer(), "setCurrentCoinHoursSupply", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setCurrentCoinHoursSupply", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectSetCurrentCoinHoursSupply() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setCurrentCoinHoursSupply")
	}
}

func (ptr *BlockchainStatusModel) SetCurrentCoinHoursSupply(currentCoinHoursSupply string) {
	if ptr.Pointer() != nil {
		var currentCoinHoursSupplyC *C.char
		if currentCoinHoursSupply != "" {
			currentCoinHoursSupplyC = C.CString(currentCoinHoursSupply)
			defer C.free(unsafe.Pointer(currentCoinHoursSupplyC))
		}
		C.BlockchainStatusModel9ae898_SetCurrentCoinHoursSupply(ptr.Pointer(), C.struct_Moc_PackedString{data: currentCoinHoursSupplyC, len: C.longlong(len(currentCoinHoursSupply))})
	}
}

func (ptr *BlockchainStatusModel) SetCurrentCoinHoursSupplyDefault(currentCoinHoursSupply string) {
	if ptr.Pointer() != nil {
		var currentCoinHoursSupplyC *C.char
		if currentCoinHoursSupply != "" {
			currentCoinHoursSupplyC = C.CString(currentCoinHoursSupply)
			defer C.free(unsafe.Pointer(currentCoinHoursSupplyC))
		}
		C.BlockchainStatusModel9ae898_SetCurrentCoinHoursSupplyDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: currentCoinHoursSupplyC, len: C.longlong(len(currentCoinHoursSupply))})
	}
}

//export callbackBlockchainStatusModel9ae898_CurrentCoinHoursSupplyChanged
func callbackBlockchainStatusModel9ae898_CurrentCoinHoursSupplyChanged(ptr unsafe.Pointer, currentCoinHoursSupply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "currentCoinHoursSupplyChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(currentCoinHoursSupply))
	}

}

func (ptr *BlockchainStatusModel) ConnectCurrentCoinHoursSupplyChanged(f func(currentCoinHoursSupply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "currentCoinHoursSupplyChanged") {
			C.BlockchainStatusModel9ae898_ConnectCurrentCoinHoursSupplyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "currentCoinHoursSupplyChanged"); signal != nil {
			f := func(currentCoinHoursSupply string) {
				(*(*func(string))(signal))(currentCoinHoursSupply)
				f(currentCoinHoursSupply)
			}
			qt.ConnectSignal(ptr.Pointer(), "currentCoinHoursSupplyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "currentCoinHoursSupplyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectCurrentCoinHoursSupplyChanged() {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_DisconnectCurrentCoinHoursSupplyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "currentCoinHoursSupplyChanged")
	}
}

func (ptr *BlockchainStatusModel) CurrentCoinHoursSupplyChanged(currentCoinHoursSupply string) {
	if ptr.Pointer() != nil {
		var currentCoinHoursSupplyC *C.char
		if currentCoinHoursSupply != "" {
			currentCoinHoursSupplyC = C.CString(currentCoinHoursSupply)
			defer C.free(unsafe.Pointer(currentCoinHoursSupplyC))
		}
		C.BlockchainStatusModel9ae898_CurrentCoinHoursSupplyChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: currentCoinHoursSupplyC, len: C.longlong(len(currentCoinHoursSupply))})
	}
}

//export callbackBlockchainStatusModel9ae898_TotalCoinHoursSupply
func callbackBlockchainStatusModel9ae898_TotalCoinHoursSupply(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "totalCoinHoursSupply"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewBlockchainStatusModelFromPointer(ptr).TotalCoinHoursSupplyDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *BlockchainStatusModel) ConnectTotalCoinHoursSupply(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "totalCoinHoursSupply"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "totalCoinHoursSupply", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "totalCoinHoursSupply", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectTotalCoinHoursSupply() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "totalCoinHoursSupply")
	}
}

func (ptr *BlockchainStatusModel) TotalCoinHoursSupply() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel9ae898_TotalCoinHoursSupply(ptr.Pointer()))
	}
	return ""
}

func (ptr *BlockchainStatusModel) TotalCoinHoursSupplyDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel9ae898_TotalCoinHoursSupplyDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackBlockchainStatusModel9ae898_SetTotalCoinHoursSupply
func callbackBlockchainStatusModel9ae898_SetTotalCoinHoursSupply(ptr unsafe.Pointer, totalCoinHoursSupply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setTotalCoinHoursSupply"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(totalCoinHoursSupply))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).SetTotalCoinHoursSupplyDefault(cGoUnpackString(totalCoinHoursSupply))
	}
}

func (ptr *BlockchainStatusModel) ConnectSetTotalCoinHoursSupply(f func(totalCoinHoursSupply string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTotalCoinHoursSupply"); signal != nil {
			f := func(totalCoinHoursSupply string) {
				(*(*func(string))(signal))(totalCoinHoursSupply)
				f(totalCoinHoursSupply)
			}
			qt.ConnectSignal(ptr.Pointer(), "setTotalCoinHoursSupply", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setTotalCoinHoursSupply", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectSetTotalCoinHoursSupply() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setTotalCoinHoursSupply")
	}
}

func (ptr *BlockchainStatusModel) SetTotalCoinHoursSupply(totalCoinHoursSupply string) {
	if ptr.Pointer() != nil {
		var totalCoinHoursSupplyC *C.char
		if totalCoinHoursSupply != "" {
			totalCoinHoursSupplyC = C.CString(totalCoinHoursSupply)
			defer C.free(unsafe.Pointer(totalCoinHoursSupplyC))
		}
		C.BlockchainStatusModel9ae898_SetTotalCoinHoursSupply(ptr.Pointer(), C.struct_Moc_PackedString{data: totalCoinHoursSupplyC, len: C.longlong(len(totalCoinHoursSupply))})
	}
}

func (ptr *BlockchainStatusModel) SetTotalCoinHoursSupplyDefault(totalCoinHoursSupply string) {
	if ptr.Pointer() != nil {
		var totalCoinHoursSupplyC *C.char
		if totalCoinHoursSupply != "" {
			totalCoinHoursSupplyC = C.CString(totalCoinHoursSupply)
			defer C.free(unsafe.Pointer(totalCoinHoursSupplyC))
		}
		C.BlockchainStatusModel9ae898_SetTotalCoinHoursSupplyDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: totalCoinHoursSupplyC, len: C.longlong(len(totalCoinHoursSupply))})
	}
}

//export callbackBlockchainStatusModel9ae898_TotalCoinHoursSupplyChanged
func callbackBlockchainStatusModel9ae898_TotalCoinHoursSupplyChanged(ptr unsafe.Pointer, totalCoinHoursSupply C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "totalCoinHoursSupplyChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(totalCoinHoursSupply))
	}

}

func (ptr *BlockchainStatusModel) ConnectTotalCoinHoursSupplyChanged(f func(totalCoinHoursSupply string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "totalCoinHoursSupplyChanged") {
			C.BlockchainStatusModel9ae898_ConnectTotalCoinHoursSupplyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "totalCoinHoursSupplyChanged"); signal != nil {
			f := func(totalCoinHoursSupply string) {
				(*(*func(string))(signal))(totalCoinHoursSupply)
				f(totalCoinHoursSupply)
			}
			qt.ConnectSignal(ptr.Pointer(), "totalCoinHoursSupplyChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "totalCoinHoursSupplyChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectTotalCoinHoursSupplyChanged() {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_DisconnectTotalCoinHoursSupplyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "totalCoinHoursSupplyChanged")
	}
}

func (ptr *BlockchainStatusModel) TotalCoinHoursSupplyChanged(totalCoinHoursSupply string) {
	if ptr.Pointer() != nil {
		var totalCoinHoursSupplyC *C.char
		if totalCoinHoursSupply != "" {
			totalCoinHoursSupplyC = C.CString(totalCoinHoursSupply)
			defer C.free(unsafe.Pointer(totalCoinHoursSupplyC))
		}
		C.BlockchainStatusModel9ae898_TotalCoinHoursSupplyChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: totalCoinHoursSupplyC, len: C.longlong(len(totalCoinHoursSupply))})
	}
}

func BlockchainStatusModel_QRegisterMetaType() int {
	return int(int32(C.BlockchainStatusModel9ae898_BlockchainStatusModel9ae898_QRegisterMetaType()))
}

func (ptr *BlockchainStatusModel) QRegisterMetaType() int {
	return int(int32(C.BlockchainStatusModel9ae898_BlockchainStatusModel9ae898_QRegisterMetaType()))
}

func BlockchainStatusModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.BlockchainStatusModel9ae898_BlockchainStatusModel9ae898_QRegisterMetaType2(typeNameC)))
}

func (ptr *BlockchainStatusModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.BlockchainStatusModel9ae898_BlockchainStatusModel9ae898_QRegisterMetaType2(typeNameC)))
}

func BlockchainStatusModel_QmlRegisterType() int {
	return int(int32(C.BlockchainStatusModel9ae898_BlockchainStatusModel9ae898_QmlRegisterType()))
}

func (ptr *BlockchainStatusModel) QmlRegisterType() int {
	return int(int32(C.BlockchainStatusModel9ae898_BlockchainStatusModel9ae898_QmlRegisterType()))
}

func BlockchainStatusModel_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.BlockchainStatusModel9ae898_BlockchainStatusModel9ae898_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *BlockchainStatusModel) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.BlockchainStatusModel9ae898_BlockchainStatusModel9ae898_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *BlockchainStatusModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.BlockchainStatusModel9ae898___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *BlockchainStatusModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *BlockchainStatusModel) __children_newList() unsafe.Pointer {
	return C.BlockchainStatusModel9ae898___children_newList(ptr.Pointer())
}

func (ptr *BlockchainStatusModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.BlockchainStatusModel9ae898___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *BlockchainStatusModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *BlockchainStatusModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.BlockchainStatusModel9ae898___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *BlockchainStatusModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.BlockchainStatusModel9ae898___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *BlockchainStatusModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *BlockchainStatusModel) __findChildren_newList() unsafe.Pointer {
	return C.BlockchainStatusModel9ae898___findChildren_newList(ptr.Pointer())
}

func (ptr *BlockchainStatusModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.BlockchainStatusModel9ae898___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *BlockchainStatusModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *BlockchainStatusModel) __findChildren_newList3() unsafe.Pointer {
	return C.BlockchainStatusModel9ae898___findChildren_newList3(ptr.Pointer())
}

func (ptr *BlockchainStatusModel) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.BlockchainStatusModel9ae898___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *BlockchainStatusModel) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *BlockchainStatusModel) __qFindChildren_newList2() unsafe.Pointer {
	return C.BlockchainStatusModel9ae898___qFindChildren_newList2(ptr.Pointer())
}

func NewBlockchainStatusModel(parent std_core.QObject_ITF) *BlockchainStatusModel {
	tmpValue := NewBlockchainStatusModelFromPointer(C.BlockchainStatusModel9ae898_NewBlockchainStatusModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackBlockchainStatusModel9ae898_DestroyBlockchainStatusModel
func callbackBlockchainStatusModel9ae898_DestroyBlockchainStatusModel(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~BlockchainStatusModel"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewBlockchainStatusModelFromPointer(ptr).DestroyBlockchainStatusModelDefault()
	}
}

func (ptr *BlockchainStatusModel) ConnectDestroyBlockchainStatusModel(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~BlockchainStatusModel"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~BlockchainStatusModel", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~BlockchainStatusModel", unsafe.Pointer(&f))
		}
	}
}

func (ptr *BlockchainStatusModel) DisconnectDestroyBlockchainStatusModel() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~BlockchainStatusModel")
	}
}

func (ptr *BlockchainStatusModel) DestroyBlockchainStatusModel() {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_DestroyBlockchainStatusModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *BlockchainStatusModel) DestroyBlockchainStatusModelDefault() {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_DestroyBlockchainStatusModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackBlockchainStatusModel9ae898_ChildEvent
func callbackBlockchainStatusModel9ae898_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *BlockchainStatusModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackBlockchainStatusModel9ae898_ConnectNotify
func callbackBlockchainStatusModel9ae898_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *BlockchainStatusModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackBlockchainStatusModel9ae898_CustomEvent
func callbackBlockchainStatusModel9ae898_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *BlockchainStatusModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackBlockchainStatusModel9ae898_DeleteLater
func callbackBlockchainStatusModel9ae898_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewBlockchainStatusModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *BlockchainStatusModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackBlockchainStatusModel9ae898_Destroyed
func callbackBlockchainStatusModel9ae898_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackBlockchainStatusModel9ae898_DisconnectNotify
func callbackBlockchainStatusModel9ae898_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *BlockchainStatusModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackBlockchainStatusModel9ae898_Event
func callbackBlockchainStatusModel9ae898_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewBlockchainStatusModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *BlockchainStatusModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.BlockchainStatusModel9ae898_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackBlockchainStatusModel9ae898_EventFilter
func callbackBlockchainStatusModel9ae898_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewBlockchainStatusModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *BlockchainStatusModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.BlockchainStatusModel9ae898_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackBlockchainStatusModel9ae898_ObjectNameChanged
func callbackBlockchainStatusModel9ae898_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackBlockchainStatusModel9ae898_TimerEvent
func callbackBlockchainStatusModel9ae898_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *BlockchainStatusModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel9ae898_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

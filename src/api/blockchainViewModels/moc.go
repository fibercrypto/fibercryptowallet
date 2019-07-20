package blockchainViewModels

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

//export callbackBlockchainStatusModel97d618_Constructor
func callbackBlockchainStatusModel97d618_Constructor(ptr unsafe.Pointer) {
	this := NewBlockchainStatusModelFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackBlockchainStatusModel97d618_NumberOfBlocks
func callbackBlockchainStatusModel97d618_NumberOfBlocks(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "numberOfBlocks"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewBlockchainStatusModelFromPointer(ptr).NumberOfBlocksDefault()))
}

func (ptr *BlockchainStatusModel) ConnectNumberOfBlocks(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "numberOfBlocks"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
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

func (ptr *BlockchainStatusModel) NumberOfBlocks() int {
	if ptr.Pointer() != nil {
		return int(int32(C.BlockchainStatusModel97d618_NumberOfBlocks(ptr.Pointer())))
	}
	return 0
}

func (ptr *BlockchainStatusModel) NumberOfBlocksDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.BlockchainStatusModel97d618_NumberOfBlocksDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackBlockchainStatusModel97d618_SetNumberOfBlocks
func callbackBlockchainStatusModel97d618_SetNumberOfBlocks(ptr unsafe.Pointer, numberOfBlocks C.int) {
	if signal := qt.GetSignal(ptr, "setNumberOfBlocks"); signal != nil {
		(*(*func(int))(signal))(int(int32(numberOfBlocks)))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).SetNumberOfBlocksDefault(int(int32(numberOfBlocks)))
	}
}

func (ptr *BlockchainStatusModel) ConnectSetNumberOfBlocks(f func(numberOfBlocks int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setNumberOfBlocks"); signal != nil {
			f := func(numberOfBlocks int) {
				(*(*func(int))(signal))(numberOfBlocks)
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

func (ptr *BlockchainStatusModel) SetNumberOfBlocks(numberOfBlocks int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_SetNumberOfBlocks(ptr.Pointer(), C.int(int32(numberOfBlocks)))
	}
}

func (ptr *BlockchainStatusModel) SetNumberOfBlocksDefault(numberOfBlocks int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_SetNumberOfBlocksDefault(ptr.Pointer(), C.int(int32(numberOfBlocks)))
	}
}

//export callbackBlockchainStatusModel97d618_NumberOfBlocksChanged
func callbackBlockchainStatusModel97d618_NumberOfBlocksChanged(ptr unsafe.Pointer, numberOfBlocks C.int) {
	if signal := qt.GetSignal(ptr, "numberOfBlocksChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(numberOfBlocks)))
	}

}

func (ptr *BlockchainStatusModel) ConnectNumberOfBlocksChanged(f func(numberOfBlocks int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "numberOfBlocksChanged") {
			C.BlockchainStatusModel97d618_ConnectNumberOfBlocksChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "numberOfBlocksChanged"); signal != nil {
			f := func(numberOfBlocks int) {
				(*(*func(int))(signal))(numberOfBlocks)
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
		C.BlockchainStatusModel97d618_DisconnectNumberOfBlocksChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "numberOfBlocksChanged")
	}
}

func (ptr *BlockchainStatusModel) NumberOfBlocksChanged(numberOfBlocks int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_NumberOfBlocksChanged(ptr.Pointer(), C.int(int32(numberOfBlocks)))
	}
}

//export callbackBlockchainStatusModel97d618_TimestampLastBlock
func callbackBlockchainStatusModel97d618_TimestampLastBlock(ptr unsafe.Pointer) unsafe.Pointer {
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
		tmpValue := std_core.NewQDateTimeFromPointer(C.BlockchainStatusModel97d618_TimestampLastBlock(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *BlockchainStatusModel) TimestampLastBlockDefault() *std_core.QDateTime {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQDateTimeFromPointer(C.BlockchainStatusModel97d618_TimestampLastBlockDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

//export callbackBlockchainStatusModel97d618_SetTimestampLastBlock
func callbackBlockchainStatusModel97d618_SetTimestampLastBlock(ptr unsafe.Pointer, timestampLastBlock unsafe.Pointer) {
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
		C.BlockchainStatusModel97d618_SetTimestampLastBlock(ptr.Pointer(), std_core.PointerFromQDateTime(timestampLastBlock))
	}
}

func (ptr *BlockchainStatusModel) SetTimestampLastBlockDefault(timestampLastBlock std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_SetTimestampLastBlockDefault(ptr.Pointer(), std_core.PointerFromQDateTime(timestampLastBlock))
	}
}

//export callbackBlockchainStatusModel97d618_TimestampLastBlockChanged
func callbackBlockchainStatusModel97d618_TimestampLastBlockChanged(ptr unsafe.Pointer, timestampLastBlock unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timestampLastBlockChanged"); signal != nil {
		(*(*func(*std_core.QDateTime))(signal))(std_core.NewQDateTimeFromPointer(timestampLastBlock))
	}

}

func (ptr *BlockchainStatusModel) ConnectTimestampLastBlockChanged(f func(timestampLastBlock *std_core.QDateTime)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "timestampLastBlockChanged") {
			C.BlockchainStatusModel97d618_ConnectTimestampLastBlockChanged(ptr.Pointer())
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
		C.BlockchainStatusModel97d618_DisconnectTimestampLastBlockChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "timestampLastBlockChanged")
	}
}

func (ptr *BlockchainStatusModel) TimestampLastBlockChanged(timestampLastBlock std_core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_TimestampLastBlockChanged(ptr.Pointer(), std_core.PointerFromQDateTime(timestampLastBlock))
	}
}

//export callbackBlockchainStatusModel97d618_HashLastBlock
func callbackBlockchainStatusModel97d618_HashLastBlock(ptr unsafe.Pointer) C.struct_Moc_PackedString {
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
		return cGoUnpackString(C.BlockchainStatusModel97d618_HashLastBlock(ptr.Pointer()))
	}
	return ""
}

func (ptr *BlockchainStatusModel) HashLastBlockDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.BlockchainStatusModel97d618_HashLastBlockDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackBlockchainStatusModel97d618_SetHashLastBlock
func callbackBlockchainStatusModel97d618_SetHashLastBlock(ptr unsafe.Pointer, hashLastBlock C.struct_Moc_PackedString) {
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
		C.BlockchainStatusModel97d618_SetHashLastBlock(ptr.Pointer(), C.struct_Moc_PackedString{data: hashLastBlockC, len: C.longlong(len(hashLastBlock))})
	}
}

func (ptr *BlockchainStatusModel) SetHashLastBlockDefault(hashLastBlock string) {
	if ptr.Pointer() != nil {
		var hashLastBlockC *C.char
		if hashLastBlock != "" {
			hashLastBlockC = C.CString(hashLastBlock)
			defer C.free(unsafe.Pointer(hashLastBlockC))
		}
		C.BlockchainStatusModel97d618_SetHashLastBlockDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: hashLastBlockC, len: C.longlong(len(hashLastBlock))})
	}
}

//export callbackBlockchainStatusModel97d618_HashLastBlockChanged
func callbackBlockchainStatusModel97d618_HashLastBlockChanged(ptr unsafe.Pointer, hashLastBlock C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "hashLastBlockChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(hashLastBlock))
	}

}

func (ptr *BlockchainStatusModel) ConnectHashLastBlockChanged(f func(hashLastBlock string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hashLastBlockChanged") {
			C.BlockchainStatusModel97d618_ConnectHashLastBlockChanged(ptr.Pointer())
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
		C.BlockchainStatusModel97d618_DisconnectHashLastBlockChanged(ptr.Pointer())
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
		C.BlockchainStatusModel97d618_HashLastBlockChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: hashLastBlockC, len: C.longlong(len(hashLastBlock))})
	}
}

//export callbackBlockchainStatusModel97d618_CurrentSkySupply
func callbackBlockchainStatusModel97d618_CurrentSkySupply(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "currentSkySupply"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewBlockchainStatusModelFromPointer(ptr).CurrentSkySupplyDefault()))
}

func (ptr *BlockchainStatusModel) ConnectCurrentSkySupply(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "currentSkySupply"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
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

func (ptr *BlockchainStatusModel) CurrentSkySupply() int {
	if ptr.Pointer() != nil {
		return int(int32(C.BlockchainStatusModel97d618_CurrentSkySupply(ptr.Pointer())))
	}
	return 0
}

func (ptr *BlockchainStatusModel) CurrentSkySupplyDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.BlockchainStatusModel97d618_CurrentSkySupplyDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackBlockchainStatusModel97d618_SetCurrentSkySupply
func callbackBlockchainStatusModel97d618_SetCurrentSkySupply(ptr unsafe.Pointer, currentSkySupply C.int) {
	if signal := qt.GetSignal(ptr, "setCurrentSkySupply"); signal != nil {
		(*(*func(int))(signal))(int(int32(currentSkySupply)))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).SetCurrentSkySupplyDefault(int(int32(currentSkySupply)))
	}
}

func (ptr *BlockchainStatusModel) ConnectSetCurrentSkySupply(f func(currentSkySupply int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCurrentSkySupply"); signal != nil {
			f := func(currentSkySupply int) {
				(*(*func(int))(signal))(currentSkySupply)
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

func (ptr *BlockchainStatusModel) SetCurrentSkySupply(currentSkySupply int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_SetCurrentSkySupply(ptr.Pointer(), C.int(int32(currentSkySupply)))
	}
}

func (ptr *BlockchainStatusModel) SetCurrentSkySupplyDefault(currentSkySupply int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_SetCurrentSkySupplyDefault(ptr.Pointer(), C.int(int32(currentSkySupply)))
	}
}

//export callbackBlockchainStatusModel97d618_CurrentSkySupplyChanged
func callbackBlockchainStatusModel97d618_CurrentSkySupplyChanged(ptr unsafe.Pointer, currentSkySupply C.int) {
	if signal := qt.GetSignal(ptr, "currentSkySupplyChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(currentSkySupply)))
	}

}

func (ptr *BlockchainStatusModel) ConnectCurrentSkySupplyChanged(f func(currentSkySupply int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "currentSkySupplyChanged") {
			C.BlockchainStatusModel97d618_ConnectCurrentSkySupplyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "currentSkySupplyChanged"); signal != nil {
			f := func(currentSkySupply int) {
				(*(*func(int))(signal))(currentSkySupply)
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
		C.BlockchainStatusModel97d618_DisconnectCurrentSkySupplyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "currentSkySupplyChanged")
	}
}

func (ptr *BlockchainStatusModel) CurrentSkySupplyChanged(currentSkySupply int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_CurrentSkySupplyChanged(ptr.Pointer(), C.int(int32(currentSkySupply)))
	}
}

//export callbackBlockchainStatusModel97d618_TotalSkySupply
func callbackBlockchainStatusModel97d618_TotalSkySupply(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "totalSkySupply"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewBlockchainStatusModelFromPointer(ptr).TotalSkySupplyDefault()))
}

func (ptr *BlockchainStatusModel) ConnectTotalSkySupply(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "totalSkySupply"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
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

func (ptr *BlockchainStatusModel) TotalSkySupply() int {
	if ptr.Pointer() != nil {
		return int(int32(C.BlockchainStatusModel97d618_TotalSkySupply(ptr.Pointer())))
	}
	return 0
}

func (ptr *BlockchainStatusModel) TotalSkySupplyDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.BlockchainStatusModel97d618_TotalSkySupplyDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackBlockchainStatusModel97d618_SetTotalSkySupply
func callbackBlockchainStatusModel97d618_SetTotalSkySupply(ptr unsafe.Pointer, totalSkySupply C.int) {
	if signal := qt.GetSignal(ptr, "setTotalSkySupply"); signal != nil {
		(*(*func(int))(signal))(int(int32(totalSkySupply)))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).SetTotalSkySupplyDefault(int(int32(totalSkySupply)))
	}
}

func (ptr *BlockchainStatusModel) ConnectSetTotalSkySupply(f func(totalSkySupply int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTotalSkySupply"); signal != nil {
			f := func(totalSkySupply int) {
				(*(*func(int))(signal))(totalSkySupply)
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

func (ptr *BlockchainStatusModel) SetTotalSkySupply(totalSkySupply int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_SetTotalSkySupply(ptr.Pointer(), C.int(int32(totalSkySupply)))
	}
}

func (ptr *BlockchainStatusModel) SetTotalSkySupplyDefault(totalSkySupply int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_SetTotalSkySupplyDefault(ptr.Pointer(), C.int(int32(totalSkySupply)))
	}
}

//export callbackBlockchainStatusModel97d618_TotalSkySupplyChanged
func callbackBlockchainStatusModel97d618_TotalSkySupplyChanged(ptr unsafe.Pointer, totalSkySupply C.int) {
	if signal := qt.GetSignal(ptr, "totalSkySupplyChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(totalSkySupply)))
	}

}

func (ptr *BlockchainStatusModel) ConnectTotalSkySupplyChanged(f func(totalSkySupply int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "totalSkySupplyChanged") {
			C.BlockchainStatusModel97d618_ConnectTotalSkySupplyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "totalSkySupplyChanged"); signal != nil {
			f := func(totalSkySupply int) {
				(*(*func(int))(signal))(totalSkySupply)
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
		C.BlockchainStatusModel97d618_DisconnectTotalSkySupplyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "totalSkySupplyChanged")
	}
}

func (ptr *BlockchainStatusModel) TotalSkySupplyChanged(totalSkySupply int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_TotalSkySupplyChanged(ptr.Pointer(), C.int(int32(totalSkySupply)))
	}
}

//export callbackBlockchainStatusModel97d618_CurrentCoinHoursSupply
func callbackBlockchainStatusModel97d618_CurrentCoinHoursSupply(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "currentCoinHoursSupply"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewBlockchainStatusModelFromPointer(ptr).CurrentCoinHoursSupplyDefault()))
}

func (ptr *BlockchainStatusModel) ConnectCurrentCoinHoursSupply(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "currentCoinHoursSupply"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
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

func (ptr *BlockchainStatusModel) CurrentCoinHoursSupply() int {
	if ptr.Pointer() != nil {
		return int(int32(C.BlockchainStatusModel97d618_CurrentCoinHoursSupply(ptr.Pointer())))
	}
	return 0
}

func (ptr *BlockchainStatusModel) CurrentCoinHoursSupplyDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.BlockchainStatusModel97d618_CurrentCoinHoursSupplyDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackBlockchainStatusModel97d618_SetCurrentCoinHoursSupply
func callbackBlockchainStatusModel97d618_SetCurrentCoinHoursSupply(ptr unsafe.Pointer, currentCoinHoursSupply C.int) {
	if signal := qt.GetSignal(ptr, "setCurrentCoinHoursSupply"); signal != nil {
		(*(*func(int))(signal))(int(int32(currentCoinHoursSupply)))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).SetCurrentCoinHoursSupplyDefault(int(int32(currentCoinHoursSupply)))
	}
}

func (ptr *BlockchainStatusModel) ConnectSetCurrentCoinHoursSupply(f func(currentCoinHoursSupply int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setCurrentCoinHoursSupply"); signal != nil {
			f := func(currentCoinHoursSupply int) {
				(*(*func(int))(signal))(currentCoinHoursSupply)
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

func (ptr *BlockchainStatusModel) SetCurrentCoinHoursSupply(currentCoinHoursSupply int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_SetCurrentCoinHoursSupply(ptr.Pointer(), C.int(int32(currentCoinHoursSupply)))
	}
}

func (ptr *BlockchainStatusModel) SetCurrentCoinHoursSupplyDefault(currentCoinHoursSupply int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_SetCurrentCoinHoursSupplyDefault(ptr.Pointer(), C.int(int32(currentCoinHoursSupply)))
	}
}

//export callbackBlockchainStatusModel97d618_CurrentCoinHoursSupplyChanged
func callbackBlockchainStatusModel97d618_CurrentCoinHoursSupplyChanged(ptr unsafe.Pointer, currentCoinHoursSupply C.int) {
	if signal := qt.GetSignal(ptr, "currentCoinHoursSupplyChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(currentCoinHoursSupply)))
	}

}

func (ptr *BlockchainStatusModel) ConnectCurrentCoinHoursSupplyChanged(f func(currentCoinHoursSupply int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "currentCoinHoursSupplyChanged") {
			C.BlockchainStatusModel97d618_ConnectCurrentCoinHoursSupplyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "currentCoinHoursSupplyChanged"); signal != nil {
			f := func(currentCoinHoursSupply int) {
				(*(*func(int))(signal))(currentCoinHoursSupply)
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
		C.BlockchainStatusModel97d618_DisconnectCurrentCoinHoursSupplyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "currentCoinHoursSupplyChanged")
	}
}

func (ptr *BlockchainStatusModel) CurrentCoinHoursSupplyChanged(currentCoinHoursSupply int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_CurrentCoinHoursSupplyChanged(ptr.Pointer(), C.int(int32(currentCoinHoursSupply)))
	}
}

//export callbackBlockchainStatusModel97d618_TotalCoinHoursSupply
func callbackBlockchainStatusModel97d618_TotalCoinHoursSupply(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "totalCoinHoursSupply"); signal != nil {
		return C.int(int32((*(*func() int)(signal))()))
	}

	return C.int(int32(NewBlockchainStatusModelFromPointer(ptr).TotalCoinHoursSupplyDefault()))
}

func (ptr *BlockchainStatusModel) ConnectTotalCoinHoursSupply(f func() int) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "totalCoinHoursSupply"); signal != nil {
			f := func() int {
				(*(*func() int)(signal))()
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

func (ptr *BlockchainStatusModel) TotalCoinHoursSupply() int {
	if ptr.Pointer() != nil {
		return int(int32(C.BlockchainStatusModel97d618_TotalCoinHoursSupply(ptr.Pointer())))
	}
	return 0
}

func (ptr *BlockchainStatusModel) TotalCoinHoursSupplyDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.BlockchainStatusModel97d618_TotalCoinHoursSupplyDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackBlockchainStatusModel97d618_SetTotalCoinHoursSupply
func callbackBlockchainStatusModel97d618_SetTotalCoinHoursSupply(ptr unsafe.Pointer, totalCoinHoursSupply C.int) {
	if signal := qt.GetSignal(ptr, "setTotalCoinHoursSupply"); signal != nil {
		(*(*func(int))(signal))(int(int32(totalCoinHoursSupply)))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).SetTotalCoinHoursSupplyDefault(int(int32(totalCoinHoursSupply)))
	}
}

func (ptr *BlockchainStatusModel) ConnectSetTotalCoinHoursSupply(f func(totalCoinHoursSupply int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setTotalCoinHoursSupply"); signal != nil {
			f := func(totalCoinHoursSupply int) {
				(*(*func(int))(signal))(totalCoinHoursSupply)
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

func (ptr *BlockchainStatusModel) SetTotalCoinHoursSupply(totalCoinHoursSupply int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_SetTotalCoinHoursSupply(ptr.Pointer(), C.int(int32(totalCoinHoursSupply)))
	}
}

func (ptr *BlockchainStatusModel) SetTotalCoinHoursSupplyDefault(totalCoinHoursSupply int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_SetTotalCoinHoursSupplyDefault(ptr.Pointer(), C.int(int32(totalCoinHoursSupply)))
	}
}

//export callbackBlockchainStatusModel97d618_TotalCoinHoursSupplyChanged
func callbackBlockchainStatusModel97d618_TotalCoinHoursSupplyChanged(ptr unsafe.Pointer, totalCoinHoursSupply C.int) {
	if signal := qt.GetSignal(ptr, "totalCoinHoursSupplyChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(totalCoinHoursSupply)))
	}

}

func (ptr *BlockchainStatusModel) ConnectTotalCoinHoursSupplyChanged(f func(totalCoinHoursSupply int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "totalCoinHoursSupplyChanged") {
			C.BlockchainStatusModel97d618_ConnectTotalCoinHoursSupplyChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "totalCoinHoursSupplyChanged"); signal != nil {
			f := func(totalCoinHoursSupply int) {
				(*(*func(int))(signal))(totalCoinHoursSupply)
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
		C.BlockchainStatusModel97d618_DisconnectTotalCoinHoursSupplyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "totalCoinHoursSupplyChanged")
	}
}

func (ptr *BlockchainStatusModel) TotalCoinHoursSupplyChanged(totalCoinHoursSupply int) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_TotalCoinHoursSupplyChanged(ptr.Pointer(), C.int(int32(totalCoinHoursSupply)))
	}
}

func BlockchainStatusModel_QRegisterMetaType() int {
	return int(int32(C.BlockchainStatusModel97d618_BlockchainStatusModel97d618_QRegisterMetaType()))
}

func (ptr *BlockchainStatusModel) QRegisterMetaType() int {
	return int(int32(C.BlockchainStatusModel97d618_BlockchainStatusModel97d618_QRegisterMetaType()))
}

func BlockchainStatusModel_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.BlockchainStatusModel97d618_BlockchainStatusModel97d618_QRegisterMetaType2(typeNameC)))
}

func (ptr *BlockchainStatusModel) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.BlockchainStatusModel97d618_BlockchainStatusModel97d618_QRegisterMetaType2(typeNameC)))
}

func BlockchainStatusModel_QmlRegisterType() int {
	return int(int32(C.BlockchainStatusModel97d618_BlockchainStatusModel97d618_QmlRegisterType()))
}

func (ptr *BlockchainStatusModel) QmlRegisterType() int {
	return int(int32(C.BlockchainStatusModel97d618_BlockchainStatusModel97d618_QmlRegisterType()))
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
	return int(int32(C.BlockchainStatusModel97d618_BlockchainStatusModel97d618_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
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
	return int(int32(C.BlockchainStatusModel97d618_BlockchainStatusModel97d618_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *BlockchainStatusModel) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.BlockchainStatusModel97d618___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *BlockchainStatusModel) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *BlockchainStatusModel) __children_newList() unsafe.Pointer {
	return C.BlockchainStatusModel97d618___children_newList(ptr.Pointer())
}

func (ptr *BlockchainStatusModel) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.BlockchainStatusModel97d618___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *BlockchainStatusModel) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *BlockchainStatusModel) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.BlockchainStatusModel97d618___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *BlockchainStatusModel) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.BlockchainStatusModel97d618___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *BlockchainStatusModel) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *BlockchainStatusModel) __findChildren_newList() unsafe.Pointer {
	return C.BlockchainStatusModel97d618___findChildren_newList(ptr.Pointer())
}

func (ptr *BlockchainStatusModel) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.BlockchainStatusModel97d618___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *BlockchainStatusModel) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *BlockchainStatusModel) __findChildren_newList3() unsafe.Pointer {
	return C.BlockchainStatusModel97d618___findChildren_newList3(ptr.Pointer())
}

func (ptr *BlockchainStatusModel) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.BlockchainStatusModel97d618___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *BlockchainStatusModel) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *BlockchainStatusModel) __qFindChildren_newList2() unsafe.Pointer {
	return C.BlockchainStatusModel97d618___qFindChildren_newList2(ptr.Pointer())
}

func NewBlockchainStatusModel(parent std_core.QObject_ITF) *BlockchainStatusModel {
	tmpValue := NewBlockchainStatusModelFromPointer(C.BlockchainStatusModel97d618_NewBlockchainStatusModel(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackBlockchainStatusModel97d618_DestroyBlockchainStatusModel
func callbackBlockchainStatusModel97d618_DestroyBlockchainStatusModel(ptr unsafe.Pointer) {
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
		C.BlockchainStatusModel97d618_DestroyBlockchainStatusModel(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *BlockchainStatusModel) DestroyBlockchainStatusModelDefault() {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_DestroyBlockchainStatusModelDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackBlockchainStatusModel97d618_ChildEvent
func callbackBlockchainStatusModel97d618_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *BlockchainStatusModel) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackBlockchainStatusModel97d618_ConnectNotify
func callbackBlockchainStatusModel97d618_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *BlockchainStatusModel) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackBlockchainStatusModel97d618_CustomEvent
func callbackBlockchainStatusModel97d618_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *BlockchainStatusModel) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackBlockchainStatusModel97d618_DeleteLater
func callbackBlockchainStatusModel97d618_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewBlockchainStatusModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *BlockchainStatusModel) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackBlockchainStatusModel97d618_Destroyed
func callbackBlockchainStatusModel97d618_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackBlockchainStatusModel97d618_DisconnectNotify
func callbackBlockchainStatusModel97d618_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *BlockchainStatusModel) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackBlockchainStatusModel97d618_Event
func callbackBlockchainStatusModel97d618_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewBlockchainStatusModelFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *BlockchainStatusModel) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.BlockchainStatusModel97d618_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackBlockchainStatusModel97d618_EventFilter
func callbackBlockchainStatusModel97d618_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewBlockchainStatusModelFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *BlockchainStatusModel) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.BlockchainStatusModel97d618_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackBlockchainStatusModel97d618_ObjectNameChanged
func callbackBlockchainStatusModel97d618_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackBlockchainStatusModel97d618_TimerEvent
func callbackBlockchainStatusModel97d618_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewBlockchainStatusModelFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *BlockchainStatusModel) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.BlockchainStatusModel97d618_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

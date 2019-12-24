package util

import (
	"errors"
	"fmt"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"reflect"
)

var (
	errIteratorEmpty = errors.New("iterator is empty")
	errMustBePointer = errors.New("val must be a pointer")
)

// GenericIterator implement the iterator interface for any type.
type GenericIterator struct {
	data     []interface{}
	index    int
	dataType reflect.Type
}

// NewGenericIterator return a new instance of GenericIterator
func NewGenericIterator(value interface{}) core.Iterator {
	var iteratorData []interface{}
	reflectValue := reflect.ValueOf(value)
	if reflectValue.Kind() == reflect.Slice ||
		reflectValue.Kind() == reflect.Array {
		for i := 0; i < reflectValue.Len(); i++ {
			iteratorData = append(iteratorData, reflectValue.Index(i).Interface())
		}
	}

	return &GenericIterator{
		data:     iteratorData,
		index:    -1,
		dataType: reflect.TypeOf(value).Elem(),
	}
}

// CurrentData write in val parameter the current data of the iterator. If val is not a pointer or is nil
// return an error. If iterator type is not assignable to val type return an error.
func (g *GenericIterator) CurrentData(val interface{}) error {
	reflectVal := reflect.ValueOf(val)
	if len(g.data) == 0 {
		return errIteratorEmpty
	}

	if reflectVal.Kind() != reflect.Ptr || reflectVal.IsNil() {
		return errMustBePointer
	}
	if !g.dataType.AssignableTo(reflect.TypeOf(val).Elem()) {
		return fmt.Errorf("type:%T can't be assignable to type %T", g.dataType, val)
	}
	reflectCurrentData := reflect.ValueOf(g.data[g.index])
	reflectVal.Elem().Set(reflectCurrentData)
	return nil
}

// HasNext ask if next instance of the iterator is not nil.
func (g *GenericIterator) HasNext() bool {
	return (g.index + 1) < len(g.data)
}

// Next assign to the current data the data of the next instance of the iterator and return true.
// If the next instance not exist (is nil) return false.
func (g *GenericIterator) Next() bool {
	if g.HasNext() {
		g.index++
		return true
	}
	return false
}

// func (g *GenericIterator) GetType() string {
// 	if len(g.data) == 0 {
// 		return ""
// 	}
// 	return g.dataType.Name()
// }

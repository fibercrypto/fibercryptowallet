package util

import (
	"errors"
	"fmt"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"reflect"
)

type GenericIterator struct {
	data     []interface{}
	index    int
	dataType reflect.Type
}

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

func (g *GenericIterator) CurrentData(val interface{}) error {
	reflectVal := reflect.ValueOf(val)
	if len(g.data) == 0 {
		return errors.New("iterator is empty")
	}

	if reflectVal.Kind() != reflect.Ptr || reflectVal.IsNil() {
		return errors.New("val must be a pointer")
	}
	if !g.dataType.AssignableTo(reflect.TypeOf(val).Elem()) {
		return fmt.Errorf("type:%T can't be assignable to type %T", g.dataType, val)
	}
	reflectCurrentData := reflect.ValueOf(g.data[g.index])
	reflectVal.Elem().Set(reflectCurrentData)
	return nil
}

func (g *GenericIterator) HasNext() bool {
	return (g.index + 1) < len(g.data)
}

func (g *GenericIterator) Next() bool {
	if g.HasNext() {
		g.index++
		return true
	}
	return false
}

func (g *GenericIterator) GetType() string {
	if len(g.data) == 0 {
		return ""
	}
	return g.dataType.Name()
}

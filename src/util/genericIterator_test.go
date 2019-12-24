package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testInterface interface {
	testParams() string
}

type testStruct1 struct {
	param1 string
}

func (t testStruct1) testParams() string {
	return t.param1
}

type testStruct2 struct {
	param1 bool
	param2 int
}

func (t *testStruct2) testParams() string {
	return "this is a test"
}

func TestGenericIterator_CurrentData(t *testing.T) {
	tests := []struct {
		name    string
		test    func(t *testing.T)
		wantErr bool
	}{
		{name: "int", test: func(t *testing.T) {
			var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
			it := NewGenericIterator(list)
			require.True(t, it.HasNext())
			for it.Next() {
				var i int
				require.NoError(t, it.CurrentData(&i))
				require.Contains(t, list, i)
			}
			require.False(t, it.HasNext())
		}, wantErr: false},
		{name: "struct", test: func(t *testing.T) {
			type Writer struct {
				name string
				Age  int
			}
			var writers = []Writer{{
				name: "Mary Shelly",
				Age:  25,
			}, {
				name: "Walter Scott",
				Age:  49,
			}, {
				name: "John Milton",
				Age:  59,
			}, {
				name: "Oscar Wilde",
				Age:  36,
			}, {
				name: "Abraham Stoker",
				Age:  50,
			}}
			it := NewGenericIterator(writers)
			require.True(t, it.HasNext())
			for it.Next() {
				var i Writer
				require.NoError(t, it.CurrentData(&i))
				require.Contains(t, writers, i)
				require.NotNil(t, i.name)
				require.NotNil(t, i.Age)
			}
			require.False(t, it.HasNext())
		}, wantErr: false},
		{name: "implementing interfaces", test: func(t *testing.T) {
			var list = []testInterface{
				testStruct1{param1: "a1"},
				&testStruct2{
					param1: true,
					param2: 23,
				}, &testStruct2{
					param1: false,
					param2: 32,
				}}
			it := NewGenericIterator(list)
			require.True(t, it.HasNext())

			for it.Next() {
				var i testInterface
				require.NoError(t, it.CurrentData(&i))
				require.Contains(t, list, i)
			}
			require.False(t, it.HasNext())

		}, wantErr: false},
		{name: "wrong", test: func(t *testing.T) {
			var list = []string{"1", "2", "3", "4"}
			it := NewGenericIterator(list)
			require.True(t, it.HasNext())

			for it.Next() {
				var i int
				require.Error(t, it.CurrentData(&i))
			}
			require.False(t, it.HasNext())

		}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

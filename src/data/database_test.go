package data

import (
	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestBoltStorage_InsertValue(t *testing.T) {

	tests := []struct {
		name    string
		args    interface{}
		field   func() *BoltStorage
		want    uint64
		wantErr error
	}{
		{name: "databaseNotOpen",
			args: nil, field: func() *BoltStorage {
				return &BoltStorage{&bolt.DB{}}
			}, want: 0, wantErr: bolt.ErrDatabaseNotOpen},
		{name: "valueNotMatch",
			args: nil, field: func() *BoltStorage {
				b, err := GetBoltStorage(GetFilePath(t))
				require.NoError(t, err)
				return b
			}, want: 0, wantErr: errValueNoMatch(nil, []byte{})},
		{name: "valid",
			args: []byte("TestByteSlice"), field: func() *BoltStorage {
				b, err := GetBoltStorage(GetFilePath(t))
				require.NoError(t, err)
				return b
			}, want: 1, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			b := tt.field()
			require.NoError(t, err)

			if b.Path() != "" {
				defer CloseTest(t, b)
			}

			got, err := b.InsertValue(tt.args)
			if tt.wantErr != nil {
				require.Equal(t, tt.wantErr, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBoltStorage_GetValue(t *testing.T) {

	tests := []struct {
		name    string
		field   func(interface{}) (*BoltStorage, uint64)
		args    []byte
		want    interface{}
		wantErr error
	}{
		{name: "databaseNotOpen", field: func(interface{}) (*BoltStorage, uint64) {
			return &BoltStorage{&bolt.DB{}}, 0
		}, args: nil, want: nil, wantErr: bolt.ErrDatabaseNotOpen},
		{name: "BucketEmpty", field: func(val interface{}) (*BoltStorage, uint64) {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			return b, 0
		}, args: []byte("TestByteSlice"), want: nil, wantErr: errBucketEmpty},
		{name: "ValEmpty", field: func(val interface{}) (*BoltStorage, uint64) {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			id, err := b.InsertValue(val)
			require.NoError(t, err)
			require.NoError(t, b.DeleteValue(id))
			return b, id
		}, args: []byte("TestByteSlice"), want: nil, wantErr: errValEmpty},
		{name: "valid", field: func(val interface{}) (*BoltStorage, uint64) {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			id, err := b.InsertValue(val)
			require.NoError(t, err)
			return b, id
		}, args: []byte("TestByteSlice"), want: []byte("TestByteSlice"), wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			b, id := tt.field(tt.args)
			require.NoError(t, err)

			if b.Path() != "" {
				defer CloseTest(t, b)
			}

			got, err := b.GetValue(id)
			if tt.wantErr != nil {
				require.Equal(t, tt.wantErr, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBoltStorage_ListValues(t *testing.T) {

	tests := []struct {
		name    string
		field   func() *BoltStorage
		want    map[uint64]interface{}
		wantErr error
	}{
		{name: "databaseNotOpen", field: func() *BoltStorage {
			return &BoltStorage{&bolt.DB{}}
		}, want: nil, wantErr: bolt.ErrDatabaseNotOpen},
		{name: "BucketEmpty", field: func() *BoltStorage {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			return b
		}, want: nil, wantErr: errBucketEmpty},
		{name: "emptyStorage", field: func() *BoltStorage {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			id, err := b.InsertValue([]byte("TestCase1"))
			require.NoError(t, err)
			require.NoError(t, b.DeleteValue(id))
			return b
		}, want: map[uint64]interface{}{}, wantErr: nil},
		{name: "validValues", field: func() *BoltStorage {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			for i := 1; i <= 10; i++ {
				_, err := b.InsertValue([]byte("TestCase" + strconv.Itoa(i)))
				require.NoError(t, err)
			}
			return b
		}, want: map[uint64]interface{}{1: []byte("TestCase1"), 2: []byte("TestCase2"), 3: []byte("TestCase3"),
			4: []byte("TestCase4"), 5: []byte("TestCase5"), 6: []byte("TestCase6"), 7: []byte("TestCase7"),
			8: []byte("TestCase8"), 9: []byte("TestCase9"), 10: []byte("TestCase10")}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			b := tt.field()
			require.NoError(t, err)

			if b.Path() != "" {
				defer CloseTest(t, b)
			}

			got, err := b.ListValues()
			if tt.wantErr != nil {
				require.Equal(t, tt.wantErr, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestBoltStorage_DeleteValue(t *testing.T) {

	tests := []struct {
		name    string
		field   func() (*BoltStorage, uint64)
		wantErr error
	}{
		{name: "databaseNotOpen", field: func() (*BoltStorage, uint64) {
			return &BoltStorage{&bolt.DB{}}, 0
		}, wantErr: bolt.ErrDatabaseNotOpen},
		{name: "BucketEmpty", field: func() (*BoltStorage, uint64) {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			return b, 0
		}, wantErr: errBucketEmpty},
		{name: "ValEmpty", field: func() (*BoltStorage, uint64) {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			id, err := b.InsertValue([]byte("foo"))
			require.NoError(t, err)
			require.NoError(t, b.DeleteValue(id))
			return b, id
		}, wantErr: errValEmpty},
		{name: "valid", field: func() (*BoltStorage, uint64) {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			id, err := b.InsertValue([]byte("foo"))
			require.NoError(t, err)
			return b, id
		}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			b, id := tt.field()
			require.NoError(t, err)

			if b.Path() != "" {
				defer CloseTest(t, b)
			}

			err = b.DeleteValue(id)
			if tt.wantErr != nil {
				require.Equal(t, tt.wantErr, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestBoltStorage_UpdateValue(t *testing.T) {
	tests := []struct {
		name    string
		field   func() (*BoltStorage, uint64)
		args    interface{}
		wantErr error
	}{
		{name: "databaseNotOpen", field: func() (*BoltStorage, uint64) {
			return &BoltStorage{&bolt.DB{}}, 0
		}, wantErr: bolt.ErrDatabaseNotOpen},
		{name: "BucketEmpty", field: func() (*BoltStorage, uint64) {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			return b, 0
		}, wantErr: errBucketEmpty},
		{name: "BucketEmpty", field: func() (*BoltStorage, uint64) {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			id, err := b.InsertValue([]byte("foo"))
			require.NoError(t, err)
			return b, id
		}, args: "Bar", wantErr: errValueNoMatch("Bar", []byte{})},
		{name: "EmptyVal", field: func() (*BoltStorage, uint64) {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			id, err := b.InsertValue([]byte("foo"))
			require.NoError(t, err)
			require.NoError(t, b.DeleteValue(id))
			return b, id
		}, args: []byte("Bar"), wantErr: nil},
		{name: "OK", field: func() (*BoltStorage, uint64) {
			b, err := GetBoltStorage(GetFilePath(t))
			require.NoError(t, err)
			id, err := b.InsertValue([]byte("foo"))
			require.NoError(t, err)
			return b, id
		}, args: []byte("Bar"), wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			b, id := tt.field()
			require.NoError(t, err)

			if b.Path() != "" {
				defer CloseTest(t, b)
			}

			err = b.UpdateValue(id, tt.args)
			if tt.wantErr != nil {
				require.Equal(t, tt.wantErr, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

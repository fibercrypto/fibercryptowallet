package util

import (
	"bytes"
	"encoding/gob"
	stdErr "errors"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func GetCoinValue(value string, ticker string) (uint64, error) {
	accuracy, err := AltcoinQuotient(ticker)
	if err != nil {
		return uint64(0), err
	}
	coin, err2 := strconv.ParseFloat(value, 64)
	if err2 != nil {
		return uint64(0), err
	}
	return uint64(coin * float64(accuracy)), nil
}

func FormatUint64(n uint64) string {
	in := strconv.FormatUint(n, 10)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

func FormatCoins(n uint64, quotient uint64) string {
	if n == uint64(0) {
		return "0"
	}

	number := strconv.FormatUint(n, 10)
	lenQ := len(strconv.FormatUint(quotient, 10)) - 1
	nFormatted := FormatUint64(n / quotient)
	if lenQ > len(number) {
		return nFormatted
	}
	reminder := number[len(number)-lenQ:]
	reminder = RemoveZeros(reminder)
	if len(reminder) == 0 {
		return nFormatted
	}
	return nFormatted + "." + reminder
}

func RemoveZeros(s string) string {
	index := 0
	temp := 0
	for {
		temp = index
		for _, c := range s[index:] {
			if string(c) != "0" {
				index++
				break
			}
		}
		if temp == index {
			break
		}
	}
	return s[:index]
}

func StringInList(s string, list []string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

// DeepCopy create a deep copy of the src
func DeepCopy(src interface{}, dst interface{}) error {
	// FIXME performance issue, a reflection alternative could be 10X more faster
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(src); err != nil {
		// FIXME i18n
		logrus.WithError(err).Error("error encoding source")
		return stdErr.New("error cloning source")
	}
	dec := gob.NewDecoder(&buf)
	if err := dec.Decode(dst); err != nil {
		// FIXME i18n
		logrus.WithError(err).Error("error decoding source")
		return stdErr.New("error filling destination")
	}
	return nil
}

// CloneTransactionInputs create a deep copy of the src
func CloneTransactionInputs(src []core.TransactionInput) ([]core.TransactionInput, error) {
	newSrc := make([]core.Clonable, len(src), cap(src))
	for idx, s := range src {
		newSrc[idx] = s
	}
	newDst, err := cloneSlice(newSrc)
	if err != nil {
		return nil, err
	}
	dst := make([]core.TransactionInput, len(newDst), cap(newDst))
	for idx, d := range newDst {
		var ok bool
		if dst[idx], ok = d.(core.TransactionInput); !ok {
			return nil, stdErr.New("unable to get d as a core.TransactionInput")
		}
	}
	return dst, nil
}

// CloneTransactionOutputs create a deep copy of the src
func CloneTransactionOutputs(src []core.TransactionOutput) ([]core.TransactionOutput, error) {
	newSrc := make([]core.Clonable, len(src), cap(src))
	for idx, s := range src {
		newSrc[idx] = s
	}
	newDst, err := cloneSlice(newSrc)
	if err != nil {
		return nil, err
	}
	dst := make([]core.TransactionOutput, len(newDst), cap(newDst))
	for idx, d := range newDst {
		var ok bool
		if dst[idx], ok = d.(core.TransactionOutput); !ok {
			return nil, stdErr.New("unable to get d as a core.TransactionOutput")
		}
	}
	return dst, nil
}

// cloneSlice create a deep copy of the src slice
func cloneSlice(src []core.Clonable) (dst []core.Clonable, err error) {
	dst = make([]core.Clonable, len(src), cap(src))
	for idx, s := range src {
		o, err := s.Clone()
		if err != nil  {
			logrus.WithError(err).Errorln("unable to clone object")
			return nil, errors.ErrDeepCopyFailed
		}
		oTyped, ok := o.(core.Clonable)
		if !ok {
			logrus.Errorln("unable to get o as a core.Clonable object")
			return nil, errors.ErrInvalidTypeAssertion
		}
		dst[idx] = oTyped
	}
	return
}

// StrSlice2IntSlice transform a numbers slices from string type to
// a numbers slice expressed as int data type
func StrSlice2IntSlice(ss []string) ([]int, error) {
	if len(ss) > 0 {
		is := make([]int, len(ss))
		for i, strIdx := range ss {
			var err error
			is[i], err = strconv.Atoi(strIdx)
			if err != nil {
				logrus.WithError(err).Errorf("unable to get %s as integer\n", strIdx)
				return nil, errors.ErrIntegerInputsRequired
			}
		}
		return is, nil
	}
	return nil, nil
}

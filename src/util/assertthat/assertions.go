package assertthat

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/testutil"
	"github.com/stretchr/testify/assert"
)

// ObjectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
func ObjectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

// Fail reports a failure through
func Fail(t testutil.TestingT, failureMessage string, msgAndArgs ...interface{}) bool {
	if h, ok := t.(testutil.THelper); ok {
		h.Helper()
	}
	content := []util.LabeledContent{
		{"Error Trace", strings.Join(assert.CallerInfo(), "\n\t\t\t")},
		{"Error", failureMessage},
	}

	// Add test name if the Go version supports it
	if n, ok := t.(interface {
		Name() string
	}); ok {
		content = append(content, util.LabeledContent{"Test", n.Name()})
	}

	message := util.MessageFromMsgAndArgs(msgAndArgs...)
	if len(message) > 0 {
		content = append(content, util.LabeledContent{"Messages", message})
	}

	t.Errorf("\n%s", ""+util.LabeledOutput(content...))

	return false
}

// FailNow fails test
func FailNow(t testutil.TestingT, failureMessage string, msgAndArgs ...interface{}) bool {
	if h, ok := t.(testutil.THelper); ok {
		h.Helper()
	}
	Fail(t, failureMessage, msgAndArgs...)

	// We cannot extend TestingT with FailNow() and
	// maintain backwards compatibility, so we fallback
	// to panicking when FailNow is not available in
	// TestingT.
	// See issue #263

	if t, ok := t.(testutil.FailNower); ok {
		t.FailNow()
	} else {
		panic("test failed and t is missing `FailNow()`")
	}
	return false
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
func ElementsMatch(t testutil.TestingT, listA, listB interface{}, msgAndArgs ...interface{}) (ok bool) {
	if h, ok := t.(testutil.THelper); ok {
		h.Helper()
	}
	if testutil.IsEmpty(listA) && testutil.IsEmpty(listB) {
		return true
	}

	aKind := reflect.TypeOf(listA).Kind()
	bKind := reflect.TypeOf(listB).Kind()

	if aKind != reflect.Array && aKind != reflect.Slice {
		return Fail(t, fmt.Sprintf("%q has an unsupported type %s", listA, aKind), msgAndArgs...)
	}

	if bKind != reflect.Array && bKind != reflect.Slice {
		return Fail(t, fmt.Sprintf("%q has an unsupported type %s", listB, bKind), msgAndArgs...)
	}

	aValue := reflect.ValueOf(listA)
	bValue := reflect.ValueOf(listB)

	aLen := aValue.Len()
	bLen := bValue.Len()

	if aLen != bLen {
		return Fail(t, fmt.Sprintf("lengths don't match: %d != %d", aLen, bLen), msgAndArgs...)
	}

	// Mark indexes in bValue that we already used
	visited := make([]bool, bLen)
	for i := 0; i < aLen; i++ {
		element := aValue.Index(i).Interface()
		found := false
		for j := 0; j < bLen; j++ {
			if visited[j] {
				continue
			}
			if ObjectsAreEqual(bValue.Index(j).Interface(), element) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			return Fail(t, fmt.Sprintf("element %s appears more times in %s than in %s", element, aValue, bValue), msgAndArgs...)
		}
	}

	return true
}

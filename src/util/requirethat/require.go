package requirethat

import (
	"github.com/fibercrypto/fibercryptowallet/src/util/assertthat"
	"github.com/fibercrypto/fibercryptowallet/src/util/testutil"
)

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// requirethat.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
func ElementsMatch(t testutil.TestingT, listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	if h, ok := t.(testutil.THelper); ok {
		h.Helper()
	}
	if assertthat.ElementsMatch(t, listA, listB, msgAndArgs...) {
		return
	}
	t.FailNow()
}

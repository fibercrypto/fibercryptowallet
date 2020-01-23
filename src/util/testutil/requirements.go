package testutil

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Errorf(format string, args ...interface{})
	FailNow()
}

// THelper test type
type THelper interface {
	Helper()
}

// FailNower test type
type FailNower interface {
	FailNow()
}

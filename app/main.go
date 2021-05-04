package main

import "github.com/MikeAustin71/strmechExamples/testrunners"

// Two tests are shown here for demonstration purposes.
// Type MainTest includes another 18-tests which can be
// invoked directly from main.go
//
func main() {

	tm := TestMain{}

	tm.MainTest002()
}

type TestMain struct {
}

// MainTest001 - This method provides a demonstration of
// StrMech.ExtractDataField()
//
func (tMain TestMain) MainTest001() {
	tRunner := testrunners.MainTest{}

	tRunner.TestExtractDataField001()

}

// MainTest002 - This method provides a demonstration of
// StrMech.ReplaceRunes()
//
func (tMain TestMain) MainTest002() {
	tRunner := testrunners.MainTest{}

	tRunner.TestReplaceRunes002()

}

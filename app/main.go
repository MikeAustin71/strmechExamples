package main

import "github.com/MikeAustin71/strmechExamples/testrunners"

func main() {
	TestMain{}.MainTest002()
}

type TestMain struct {
}

// MainTest001 - This method provides a demonstration of
// StrMech.ExtractDataField()
//
func (tMain TestMain) MainTest001() {
	tRunner := testrunners.TestRunners{}

	tRunner.TestExtractDataField001()

}

// MainTest002 - This method provides a demonstration of
// StrMech.ReplaceRunes()
//
func (tMain TestMain) MainTest002() {
	tRunner := testrunners.TestRunners{}

	tRunner.TestReplaceRunes002()

}

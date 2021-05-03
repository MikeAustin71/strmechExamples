package main

import "github.com/MikeAustin71/strmechExamples/testrunners"

func main() {
	TestMain{}.MainTest001()
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
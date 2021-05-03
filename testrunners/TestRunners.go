package testrunners

import (
	"fmt"
	"github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmech"
)

import "strings"

type TestRunners struct {

}

func (tRunrs TestRunners) TestExtractDataField001() {

	ePrefix := errpref.ErrPrefixDto{}.NewEPrefCtx(
		"TestExtractDataField001()",
		"")

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago\t Link:\t US/Central\t\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedEndOfLineDelimiterIdx := expectedLastGoodIdx
	expectedLastGoodIdx--
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Zone:", "Link:"}
	expectedDataFieldStr := "America/Chicago"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "\t"
	expectedDataFieldTrailingDelimiterType := strmech.DfTrailDelimiter.EndOfField()
	expectedLeadingWordDelimiter := leadingKeyWordDelimiters[0]
	expectedLeadingKeyWordDelimiterIndex := strings.Index(targetStr, expectedLeadingWordDelimiter)
	expectedEndOfLineDelimiter := "\n"
	expectedCommentDelimiter := ""
	expectedCommentDelimiterIndex := -1

	expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

	if expectedNextTargetIdx > expectedLastGoodIdx {
		expectedNextTargetIdx = -1
	}

	sMech := strmech.StrMech{}

	datDto,
	err := sMech.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err != nil {
		fmt.Printf("Error returned by StrMech{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		fmt.Printf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
		return
	}

	if lenTargetStr != datDto.TargetStrLength {
		fmt.Printf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
		return
	}

	if startIdx != datDto.TargetStrStartIndex {
		fmt.Printf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
		return
	}

	if expectedLeadingWordDelimiter != datDto.LeadingKeyWordDelimiter {
		fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			expectedLeadingWordDelimiter, datDto.LeadingKeyWordDelimiter)
		return
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
		return
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		fmt.Printf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
		return
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		fmt.Printf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
		return
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		fmt.Printf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
		return
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		fmt.Printf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			strmech.StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedDataFieldTrailingDelimiter), true),
			strmech.StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.DataFieldTrailingDelimiter), true))
		return
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		fmt.Printf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
		return
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		fmt.Printf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
		return
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		fmt.Printf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
		return
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		fmt.Printf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			strmech.StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedEndOfLineDelimiter), false),
			strmech.StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.EndOfLineDelimiter), false))
		return
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		fmt.Printf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
		return
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		fmt.Printf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			strmech.StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedCommentDelimiter), true),
			strmech.StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.CommentDelimiter), true))
		return
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		fmt.Printf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
		return
	}

	fmt.Printf("     %s\n" +
		"     !!! SUCCESSFUL COMPLETION !!!\n" +
		"Demonstration of StrMech.ExtractDataField()",
		ePrefix.String())

}

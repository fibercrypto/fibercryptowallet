package util

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/fibercrypto/fibercryptowallet/src/core"
)

// EmptyPassword read no password
func EmptyPassword(string, core.KeyValueStore) (string, error) {
	return "", nil
}

// ConstantPassword always return same known password
func ConstantPassword(pwdText string) core.PasswordReader {
	return func(string, core.KeyValueStore) (string, error) {
		return pwdText, nil
	}
}

func MessageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}

// LabeledOutput key value tuple
type LabeledContent struct {
	Label   string
	Content string
}

// Aligns the provided message so that all lines after the first line start at the same location as the first line.
// Assumes that the first line starts at the correct location (after carriage return, tab, label, spacer and tab).
// The longestLabelLen parameter specifies the length of the longest label in the output (required becaues this is the
// basis on which the alignment occurs).
func IndentMessageLines(message string, longestLabelLen int) string {
	outBuf := new(bytes.Buffer)

	for i, scanner := 0, bufio.NewScanner(strings.NewReader(message)); scanner.Scan(); i++ {
		// no need to align first line because it starts at the correct location (after the label)
		if i != 0 {
			// append alignLen+1 spaces to align with "{{longestLabel}}:" before adding tab
			_ = outBuf.WriteString("\n\t" + strings.Repeat(" ", longestLabelLen+1) + "\t")
		}
		_ = outBuf.WriteString(scanner.Text())
	}

	return outBuf.String()
}

// LabeledOutput returns a string consisting of the provided labeledContent. Each labeled output is appended in the following manner:
//
//   \t{{label}}:{{align_spaces}}\t{{content}}\n
//
// The initial carriage return is required to undo/erase any padding added by testing.T.Errorf. The "\t{{label}}:" is for the label.
// If a label is shorter than the longest label provided, padding spaces are added to make all the labels match in length. Once this
// alignment is achieved, "\t{{content}}\n" is added for the output.
//
// If the content of the labeledOutput contains line breaks, the subsequent lines are aligned so that they start at the same location as the first line.
func LabeledOutput(content ...LabeledContent) string {
	longestLabel := 0
	for _, v := range content {
		if len(v.Label) > longestLabel {
			longestLabel = len(v.Label)
		}
	}
	var output string
	for _, v := range content {
		output += "\t" + v.Label + ":" + strings.Repeat(" ", longestLabel-len(v.Label)) + "\t" + IndentMessageLines(v.Content, longestLabel) + "\n"
	}
	return output
}

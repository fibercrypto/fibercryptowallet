package util

import (
	"fmt"
	"testing"

	"github.com/fibercrypto/fibercryptowallet/src/core"

	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/stretchr/testify/require"
)

func TestPasswords(t *testing.T) {
	tests := []struct {
		name          string
		valid         bool
		pwd_generator core.PasswordReader
		expected_pwd  string
		param         string
	}{
		{
			name:          "valid_empty1",
			pwd_generator: EmptyPassword,
			expected_pwd:  "",
			param:         "pwd1",
		},
		{
			name:          "valid_empty2",
			pwd_generator: EmptyPassword,
			expected_pwd:  "",
			param:         "pwd2",
		},
		{
			name:          "valid_constant1",
			pwd_generator: ConstantPassword("pass_number"),
			expected_pwd:  "pass_number",
			param:         "pwd1",
		},
		{
			name:          "valid_constant2",
			pwd_generator: ConstantPassword("pass_number"),
			expected_pwd:  "pass_number",
			param:         "pwd2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pwd, err := tt.pwd_generator(tt.param, &mocks.KeyValueStore{})
			require.NoError(t, err)
			require.Equal(t, tt.expected_pwd, pwd)
		})
	}
}

func TestMessageFromMsgAndArgs(t *testing.T) {
	require.Equal(t, "", MessageFromMsgAndArgs())
	require.Equal(t, "5", MessageFromMsgAndArgs(5))
	require.Equal(t, "skycoin", MessageFromMsgAndArgs("skycoin"))
	require.Equal(t, "hello world", MessageFromMsgAndArgs("%s %s", "hello", "world"))

	intFormat := "%d"
	data := make([]interface{}, 0)
	data = append(data, "")
	expected, format := "", ""
	for i := 0; i < 10; i++ {
		expected += fmt.Sprintf(intFormat, i)
		format += intFormat
		data[0] = format
		data = append(data, i)
		require.Equal(t, expected, MessageFromMsgAndArgs(data...))
	}
}

func TestIndentMessageLines(t *testing.T) {
	s := "def factorial(n):\nif n <= 1:\n\t\treturn 1\nreturn n * factorial(n - 1)"
	sFormated := "def factorial(n):\n\t\tif n <= 1:\n\t\t\t\treturn 1\n\t\treturn n * factorial(n - 1)"
	require.Equal(t, sFormated, IndentMessageLines(s, -1))

	s, sFormated = "firstLine\nsecondLine", "firstLine\n\t       \tsecondLine"
	require.Equal(t, sFormated, IndentMessageLines(s, 6))
}

func TestLabeledOutput(t *testing.T) {
	method := LabeledContent{
		Label:   "[method head]",
		Content: "//this is the method's description and next line is it's declaration\nfunc foo(param1 type1, param2 type2) (returnType1) {",
	}
	body := LabeledContent{
		Label:   "[body]",
		Content: "\t...\n\t//some proccess\n\t...",
	}
	end := LabeledContent{
		Label:   "[method end]",
		Content: "}",
	}
	lines := []interface{}{
		"\t[method head]:\t//this is the method's description and next line is it's declaration\n",
		"\t              \tfunc foo(param1 type1, param2 type2) (returnType1) {\n",
		"\t[body]:       \t\t...\n",
		"\t              \t\t//some proccess\n",
		"\t              \t\t...\n",
		"\t[method end]: \t}\n",
	}
	require.Equal(t, fmt.Sprintf("%s%s%s%s%s%s", lines...), LabeledOutput(method, body, end))
}

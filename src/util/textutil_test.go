package util

import (
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

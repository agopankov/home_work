package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpackAdditional(t *testing.T) {
	additionalTests := []struct {
		input    string
		expected string
		err      error
	}{
		{input: "a1b2c3", expected: "abbccc", err: nil},
		{input: "45", expected: "", err: ErrInvalidString},
		{input: "2abc", expected: "", err: ErrInvalidString},
		{input: "😊2🚀3", expected: "😊😊🚀🚀🚀", err: nil},
		{input: "123", expected: "", err: ErrInvalidString},
		{input: "a0b0c0", expected: "", err: nil},
		{input: `aaab\n`, expected: "", err: ErrInvalidString},
		{input: `aaab\`, expected: "", err: ErrInvalidString},
	}

	for _, tc := range additionalTests {
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			if tc.err != nil {
				require.Truef(t, errors.Is(err, tc.err), "expected error %v, got %v", tc.err, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			}
		})
	}
}

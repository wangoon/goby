package vm

import "testing"

func TestURIParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// Scheme
		{`
		require "uri"

		u = URI.parse("http://example.com")
		u.scheme
		`, "http"},
		{`
		require "uri"

		u = URI.parse("https://example.com")
		u.scheme
		`, "https"},
		// Host
		{`
		require "uri"

		u = URI.parse("http://example.com")
		u.host
		`, "example.com"},
		// Port
		{`
		require "uri"

		u = URI.parse("http://example.com")
		u.port
		`, 80},
		{`
		require "uri"

		u = URI.parse("https://example.com")
		u.port
		`, 443},
		// Path
		{`
		require "uri"

		u = URI.parse("https://example.com/posts/1")
		u.path
		`, "/posts/1"},
		{`
		require "uri"

		u = URI.parse("https://example.com")
		u.path
		`, "/"},
		// Query
		{`
		require "uri"

		u = URI.parse("https://example.com?foo=bar&a=b")
		u.query
		`, "foo=bar&a=b"},
		{`
		require "uri"

		u = URI.parse("https://example.com")
		u.query
		`, nil},
		// User
		{`
		require "uri"

		u = URI.parse("https://example.com?foo=bar&a=b")
		u.user
		`, nil},
		// Password
		{`
		require "uri"

		u = URI.parse("https://example.com")
		u.password
		`, nil},
	}

	for i, tt := range tests {
		v := initTestVM()
		evaluated := v.testEval(t, tt.input, getFilename())
		VerifyExpected(t, i, evaluated, tt.expected)
		v.checkCFP(t, i, 0)
		v.checkSP(t, i, 1)
	}
}

func TestURIParsingFail(t *testing.T) {
	testsFail := []errorTestCase{
		// No argument
		{`
		require "uri"
		URI.parse
		`, "ArgumentError: Expect 1 argument(s). got: 0", 1},
		// Unexcpeted arguments
		{`
		require "uri"
		URI.parse("foo", "bar")
		`, "ArgumentError: Expect 1 argument(s). got: 2", 1},
		// Invalid argument type
		{`
		require "uri"
		URI.parse 1
		`, "TypeError: Expect argument to be String. got: Integer", 1},
	}
	for i, tt := range testsFail {
		v := initTestVM()
		evaluated := v.testEval(t, tt.input, getFilename())
		checkErrorMsg(t, i, evaluated, tt.expected)
		v.checkCFP(t, i, tt.expectedCFP)
		v.checkSP(t, i, 1)
	}
}

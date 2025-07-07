package analyzers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLInjectionAnalyzer_Analyze(t *testing.T) {
	analyzer := NewSQLInjectionAnalyzer()

	tests := []struct {
		name     string
		line     string
		expected     bool
	}{
		{
			name:     "Detects SELECT injection with comment",
			line:     `"SELECT * FROM users WHERE name = %s"`,
			expected:     true,
		},
		{
			name:     "Detects SELECT injection with like expression",
			line:     `"SELECT id, name FROM customers WHERE email LIKE %s"`,
			expected:     true,
		},
		{
			name:     "Detects SELECT injection with other clause after WHERE",
			line:     `"SELECT * FROM logs WHERE timestamp > %s ORDER BY timestamp DESC"`,
			expected:     true,
		},
		{
			name:     "Detects SELECT injection with muliple expression in WHERE",
			line:     `"SELECT COUNT(*) FROM orders WHERE customer_id = %s AND status = 'shipped'"`,
			expected:     true,
		},
		{
			name:     "Detects SELECT injection with nested query",
			line:     `"SELECT * FROM users WHERE id IN (SELECT user_id FROM logins WHERE ip = %s)"`,
			expected:     true,
		},
		{
			name:     "Ignores unrelated SQL keyword",
			line:     `This is just a string with SELECTED text %s.`,
			expected:     false,
		},
		{
			name:     "Ignores query without quotation marks",
			line:     `SELECT * FROM users WHERE name = %s`,
			expected:     false,
		},
	}

	for _, test := range tests {
		match := analyzer.Analyze(test.line)
		assert.Equal(t, test.expected, match, test.name)
	}
}


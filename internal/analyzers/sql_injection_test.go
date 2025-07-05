package analyzers

import (
	"testing"
)

func TestSQLInjectionAnalyzer_Analyze(t *testing.T) {
	analyzer := NewSQLInjectionAnalyzer()

	tests := []struct {
		name     string
		line     string
		want     bool
	}{
		{
			name:     "Detects SELECT injection with comment",
			line:     `"SELECT * FROM users WHERE name = %s"`,
			want:     true,
		},
		{
			name:     "Detects SELECT injection with like expression",
			line:     `"SELECT id, name FROM customers WHERE email LIKE %s"`,
			want:     true,
		},
		{
			name:     "Detects SELECT injection with other clause after WHERE",
			line:     `"SELECT * FROM logs WHERE timestamp > %s ORDER BY timestamp DESC"`,
			want:     true,
		},
		{
			name:     "Detects SELECT injection with muliple expression in WHERE",
			line:     `"SELECT COUNT(*) FROM orders WHERE customer_id = %s AND status = 'shipped'"`,
			want:     true,
		},
		{
			name:     "Detects SELECT injection with nested query",
			line:     `"SELECT * FROM users WHERE id IN (SELECT user_id FROM logins WHERE ip = %s)"`,
			want:     true,
		},
		{
			name:     "Ignores unrelated SQL keyword",
			line:     `This is just a string with SELECTED text %s.`,
			want:     false,
		},
		{
			name:     "Ignores query without quotation marks",
			line:     `SELECT * FROM users WHERE name = %s`,
			want:     false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			found, err := analyzer.Analyze(test.line)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if found != test.want {
				t.Errorf("Expected %v, got %v", test.want, found)
			}
		})
	}
}


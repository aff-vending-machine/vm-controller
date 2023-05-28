package gen

import "testing"

func TestToURLPath(t *testing.T) {
	tests := []struct {
		params   []string
		expected string
	}{
		{[]string{"https://api.example.com", "path1", "path2"}, "https://api.example.com/path1/path2"},
		{[]string{"https://api.example.com"}, "https://api.example.com"},
		{[]string{"https://api.example.com", "path1"}, "https://api.example.com/path1"},
		{[]string{"path1", "path2"}, "path1/path2"},
		{[]string{}, ""},
		{[]string{"path1"}, "path1"},
		{[]string{"https://api.example.com/", "path1", "/path2"}, "https://api.example.com/path1/path2"},
		{[]string{"/path1", "path2"}, "path1/path2"},
	}

	for i, test := range tests {
		path := ToURLPath(test.params...)

		if path != test.expected {
			t.Errorf("Test case %d failed: expected %s, got %s", i, test.expected, path)
		}
	}
}

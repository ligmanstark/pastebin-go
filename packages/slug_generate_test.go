package services

import "testing"

func TestGenerateRandomSlug(t *testing.T) {
	tests := []struct {
		name       string
		lengthSlug int
		wantLength int
	}{
		{"length 6", 6, 6},
		{"length 10", 10, 10},
		{"length 32", 32, 32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateRandomSlug(tt.lengthSlug)
			if err != nil {
				t.Errorf("GenerateRandomSlug(%d) returned error: %v", tt.lengthSlug, err)
				return
			}

			if len(got) != tt.wantLength {
				t.Errorf("GenerateRandomSlug(%d) returned length %d; want %d", tt.lengthSlug, len(got), tt.wantLength)
			}

			for _, ch := range got {
				if !isValidChar(ch) {
					t.Errorf("GenerateRandomSlug(%d) returned invalid char %c", tt.lengthSlug, ch)
				}
			}
		})
	}

}

func isValidChar(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}

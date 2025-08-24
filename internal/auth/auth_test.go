package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		want      string
		wantError error
	}{
		{
			name: "valid header",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-token"},
			},
			want:      "my-secret-token",
			wantError: nil,
		},
		{
			name:      "missing header",
			headers:   http.Header{},
			want:      "",
			wantError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed header - wrong scheme",
			headers: http.Header{
				"Authorization": []string{"Bearer something"},
			},
			want:      "",
			wantError: errors.New("malformed authorization header"),
		},
		{
			name: "malformed header - no token",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			want:      "",
			wantError: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)

			if tt.wantError != nil {
				if err == nil || err.Error() != tt.wantError.Error() {
					t.Fatalf("expected error: %v, got: %v", tt.wantError, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.want {
				t.Fatalf("expected: %v, got: %v", tt.want, got)
			}
		})
	}
}

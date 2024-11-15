package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "valid API key",
			headers: http.Header{"Authorization": {"ApiKey abc123"}},
			want:    "abc123",
			wantErr: nil,
		},
		{
			name:    "No Authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "MAH - Missing ApiKey Prefix",
			headers: http.Header{"Authorization": {"Hello abc123"}},
			want:    "",
			wantErr: ErrMalformedAuthHeader,
		},
		{
			name:    "MAH - Missing token",
			headers: http.Header{"Authorization": {"ApiKey"}},
			want:    "",
			wantErr: ErrMalformedAuthHeader,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)

			if (err != nil || tt.wantErr != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

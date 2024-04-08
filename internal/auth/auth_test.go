package auth

import (
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name    string
		headers map[string]string
		want    string
		wantErr error
	}{
		{
			name: "no auth header",
			headers: map[string]string{
				"Authorization ": "",
			},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := make(map[string][]string)
			for k, v := range tt.headers {
				headers[k] = []string{v}
			}

			got, err := GetAPIKey(headers)
			if err != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

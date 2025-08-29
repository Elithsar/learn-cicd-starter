package auth

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetApiKey(t *testing.T) {
	tests := map[string]struct {
		headers     http.Header
		expectedKey string
		expectError error
	}{
		"no auth header":        {headers: http.Header{}, expectedKey: "", expectError: ErrNoAuthHeaderIncluded},
		"malformed auth header": {headers: http.Header{"Authorization": []string{"Bearer abcd1234"}}, expectedKey: "", expectError: errors.New("malformed authorization header")},
		"valid auth header":     {headers: http.Header{"Authorization": []string{"ApiKey abcd1234"}}, expectedKey: "abcd1234", expectError: nil},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)

			if diff := cmp.Diff(tc.expectedKey, key); diff != "" {
				t.Errorf("key mismatch (-want +got):\n%s", diff)
			}

			var wantErr, gotErr string
			if tc.expectError != nil {
				wantErr = tc.expectError.Error()
			}
			if err != nil {
				gotErr = err.Error()
			}
			if diff := cmp.Diff(wantErr, gotErr); diff != "" {
				t.Errorf("error mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

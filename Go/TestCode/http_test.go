package TestCode

import "testing"

type arg struct {
	name    string
	method  Method
	domain  string
	wantErr error
}

type args []arg

func TestGetRequest(t *testing.T) {
	requests := args{
		arg{
			name:    "GET Method",
			method:  GET,
			domain:  "https://www.google.com",
			wantErr: nil,
		},
		arg{
			name:    "PUT Method",
			method:  PUT,
			domain:  "https://github.com/8luebottle/",
			wantErr: nil,
		},
	}

	for _, tt := range requests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRequest(string(tt.method)); got != tt.wantErr {
				t.Errorf("GetRequest() = %v want %v", got, tt.wantErr)
			}
		})
	}

}

// go test -run=GetRequest
/*
	--- FAIL: TestGetRequest (0.00s)
		--- FAIL: TestGetRequest/PUT_Method (0.00s)
			http_test.go:33: GetRequest() = methods not allowed want <nil>
	FAIL
	exit status 1
	FAIL    github.com/8luebottle/Play-Lab/Go/TestCode      0.100s
*/

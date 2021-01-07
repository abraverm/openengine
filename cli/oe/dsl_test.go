package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type ctrl struct {
	statusCode int
	response   interface{}
}

func (c *ctrl) mockHandler(w http.ResponseWriter, r *http.Request) {
	resp := []byte{}

	rt := reflect.TypeOf(c.response)
	if rt.Kind() == reflect.String {
		resp = []byte(c.response.(string))
	} else if rt.Kind() == reflect.Struct || rt.Kind() == reflect.Ptr {
		resp, _ = json.Marshal(c.response)
	} else {
		resp = []byte("{}")
	}

	w.WriteHeader(c.statusCode)
	w.Write(resp)
}

func HTTPMock(pattern string, statusCode int, response interface{}) *httptest.Server {
	c := &ctrl{statusCode, response}

	handler := http.NewServeMux()
	handler.HandleFunc(pattern, c.mockHandler)

	return httptest.NewServer(handler)
}

func Test_fileExists(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{"empty", "", false},
		{"dir", "testdata", false},
		{"exists", "testdata/empty", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileExists(tt.filename); got != tt.want {
				t.Errorf("fileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSource(t *testing.T) {
	srv := HTTPMock("/empty", http.StatusOK, "")
	defer srv.Close()
	tests := []struct {
		name    string
		uri     string
		want    []byte
		wantErr bool
	}{
		{"empty", "", nil, true},
		{"local file", "testdata/empty", []byte{}, false},
		{"unreachable remote file", "http://asdfasddf.com:9999/file.yaml", nil, true},
		{"unreachable remote file", fmt.Sprintf("%v/empty", srv.URL), []byte{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSource(tt.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSource() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

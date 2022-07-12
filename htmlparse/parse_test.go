package htmlparse

import (
	"reflect"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		name           string
		text           string
		wantHasToolbar bool
		wantLinks      []string
		wantErr        bool
	}{
		{"empty", "", false, nil, false},
		{"empty2", `<html></html>`, false, nil, false},
		{"header", `<html><li id="wp-admin-bar-my-account" class="menupop with-avatar"></html>`, true, nil, false},
		{"links", `<html><a href="aaa"><a href="bbb"></html>`, false, []string{"aaa", "bbb"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHasToolbar, gotLinks, err := Process(strings.NewReader(tt.text))
			if (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHasToolbar != tt.wantHasToolbar {
				t.Errorf("Process() gotHasToolbar = %v, want %v", gotHasToolbar, tt.wantHasToolbar)
			}
			if !reflect.DeepEqual(gotLinks, tt.wantLinks) {
				t.Errorf("Process() gotLinks = %v, want %v", gotLinks, tt.wantLinks)
			}
		})
	}
}

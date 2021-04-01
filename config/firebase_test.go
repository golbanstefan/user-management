package config

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func Test_firebaseStr_SetupFirebase1(t *testing.T) {
	// configure firebase
	f := Init()
	f.SetPath("./keys/firebase-key.json")
	assert.NotPanics(t, func() {
		f.SetupFirebase()
	})
}

func Test_firebaseStr_SetupFirebase(t *testing.T) {
	f := Init()
	f.SetPath("asda")
	assert.Panics(t, func() {
		f.SetupFirebase()
	}, "Get Panic")

}

func TestInit(t *testing.T) {
	f := Init()
	tests := []struct {
		name string
		want Firebase
	}{
		{"1", f},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Init(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firebaseStr_getPath(t *testing.T) {
	want := "firebase-key.json"
	f := firebaseStr{}
	if got := f.GetPath(); !strings.Contains(got, want) {
		t.Errorf("getPath() = %v, want %v", got, want)
	}

}

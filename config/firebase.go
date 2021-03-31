package config

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type Firebase interface {
	SetupFirebase() *auth.Client
	SetPath(path string)
}

func (f *firebaseStr) SetupFirebase() *auth.Client {
	opt := option.WithCredentialsFile(f.GetPath())

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("Firebase load error")
	}

	return auth
}

func (f *firebaseStr) GetPath() string {
	if f.path == "" {
		f.path = "./keys/mvp-h-c588d-firebase-adminsdk-aofl8-c959b37226.json"
	}
	serviceAccountKeyFilePath, err := filepath.Abs(f.path)
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}
	fmt.Println(serviceAccountKeyFilePath)
	return serviceAccountKeyFilePath
}
func (f *firebaseStr) SetPath(path string) {
	f.path = path
}

type firebaseStr struct {
	path string
}

var (
	m          *firebaseStr
	routerOnce sync.Once
)

func Init() Firebase {
	if m == nil {
		routerOnce.Do(func() {
			m = &firebaseStr{}
		})
	}
	return m
}

package config

import (
	"context"
	"path/filepath"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

const KEY = "./keys/firebase-key.json"
//Firebase interface provide method for connection to google firebase
type Firebase interface {
	SetupFirebase() *auth.Client
	SetPath(path string)
}

//SetupFirebase provide admin SDK initialization and
// return firebase Auth Client
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

//GetPath handle an return path to firebase key
func (f *firebaseStr) GetPath() string {
	if f.path == "" {
		f.path = KEY
	}
	serviceAccountKeyFilePath, err := filepath.Abs(f.path)
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}
	return serviceAccountKeyFilePath
}

//SetPath set custom path to  location
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

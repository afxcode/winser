package main

import (
	"context"

	"golang.org/x/sys/windows/svc/mgr"
)

type Service struct {
	Name        string
	Description string
	Command     string
	WorkingDir  string
	StartMode
	Status
}

type StartMode string

const (
	Auto   StartMode = "auto"
	Manual StartMode = "manual"
)

type Status string

const (
	Running Status = "running"
	Paused  Status = "paused"
	Stopped Status = "stopped"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Find(name string) (service Service, err error) {
	m, err := mgr.Connect()
	if err != nil {
		return
	}
	defer m.Disconnect()

	service = Service{
		Name:        name,
		Description: "Description " + name,
		Command:     "cmd",
		WorkingDir:  "workdir",
		StartMode:   Manual,
		Status:      Running,
	}
	return
}

func (a *App) Create(service Service) (err error) {

	return
}

func (a *App) Update(service Service) (err error) {

	return
}

func (a *App) Remove(name string) (err error) {

	return
}

// Control
func (a *App) Start(name string) (err error) {
	return
}

func (a *App) Stop(name string) (err error) {
	return
}

func (a *App) Pause(name string) (err error) {
	return
}

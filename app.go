package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/caarlos0/go-shellwords"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/eventlog"
	"golang.org/x/sys/windows/svc/mgr"
)

type Service struct {
	Name        string
	DisplayName string
	Description string
	Executable  string
	Argument    string
	StartMode
	Status
}

type StartMode string

const (
	Auto             StartMode = "auto"
	Manual           StartMode = "manual"
	Disabled         StartMode = "disabled"
	StartModeUnknown StartMode = "unknown"
)

func (s StartMode) toSvcStartType() uint32 {
	switch s {
	case Auto:
		return mgr.StartAutomatic
	case Manual:
		return mgr.StartManual
	case Disabled:
		return mgr.StartDisabled
	}
	return 0
}

func startModeFromSvc(startType uint32) StartMode {
	switch startType {
	case mgr.StartAutomatic:
		return Auto
	case mgr.StartManual:
		return Manual
	case mgr.StartDisabled:
		return Disabled
	default:
		return StartModeUnknown
	}
}

type Status string

const (
	Running         Status = "running"
	Paused          Status = "paused"
	Stopped         Status = "stopped"
	StartPending    Status = "pending_start"
	PausePending    Status = "pending_pause"
	StopPending     Status = "pending_stop"
	ContinuePending Status = "pending_continue"
	StatusUnknown   Status = "unknown"
)

func statusFromSvc(state svc.State) Status {
	switch state {
	case svc.Running:
		return Running
	case svc.Paused:
		return Paused
	case svc.Stopped:
		return Stopped
	case svc.StartPending:
		return StartPending
	case svc.PausePending:
		return PausePending
	case svc.StopPending:
		return StopPending
	case svc.ContinuePending:
		return ContinuePending
	default:
		return StatusUnknown
	}
}

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

	s, err := m.OpenService(name)
	if err != nil {
		err = fmt.Errorf("Could not access %s: %v", name, err)
		return
	}
	defer s.Close()

	sConf, err := s.Config()
	if err != nil {
		err = fmt.Errorf("Could not access config: %v", err)
		return
	}

	sStatus, err := s.Query()
	if err != nil {
		err = fmt.Errorf("Could not access %s: %v", name, err)
		return
	}

	args, _ := shellwords.Parse(sConf.BinaryPathName)
	executable := ""
	if len(args) > 0 {
		executable = args[0]
	}
	argument := ""
	if len(args) > 1 {
		argument = strings.Join(args[1:], " ")
	}

	service = Service{
		Name:        name,
		DisplayName: sConf.DisplayName,
		Description: sConf.Description,
		Executable:  executable,
		Argument:    argument,
		StartMode:   startModeFromSvc(sConf.StartType),
		Status:      statusFromSvc(sStatus.State),
	}
	return
}

func (a *App) SelectExecutable(currentExecutable string) (selectedExecutable string, err error) {
	dir := filepath.Dir(currentExecutable)
	if dir == "." {
		dir = "c:\\"
	}
	filename := filepath.Base(currentExecutable)
	if filename == "." {
		filename = ""
	}

	selectedExecutable, err = runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: dir,
		DefaultFilename:  filename,
		Title:            "Select Executable",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Windows Executable Files",
				Pattern:     "*.com;*.exe;*.bat;*.cmd;*.vbs;*.vbe;*.js;*.jse;*.wsf;*.wsh;*.msc",
			},
			{
				DisplayName: "All Files",
				Pattern:     "*",
			},
		},
		ShowHiddenFiles:            true,
		CanCreateDirectories:       true,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
	if selectedExecutable == "" {
		selectedExecutable = currentExecutable
	}
	return
}

func (a *App) Install(service Service) error {
	fi, err := os.Stat(service.Executable)
	if err == nil {
		if fi.Mode().IsDir() {
			return fmt.Errorf("Executable is a directory")
		}
	}

	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	s, err := m.OpenService(service.Name)
	if err == nil {
		s.Close()
		return fmt.Errorf("Service %s already exist", service.Name)
	}

	s, err = m.CreateService(
		service.Name,
		service.Executable,
		mgr.Config{
			StartType:   service.StartMode.toSvcStartType(),
			DisplayName: service.DisplayName,
			Description: service.Description,
		},
		strings.Split(service.Argument, " ")...,
	)
	if err != nil {
		return fmt.Errorf("Install service failed: %s", err)
	}
	defer s.Close()

	if err = eventlog.InstallAsEventCreate(service.Name, eventlog.Error|eventlog.Warning|eventlog.Info); err != nil {
		s.Delete()
		return fmt.Errorf("Install event log failed: %s", err)
	}
	return nil
}

func (a *App) Update(service Service) (err error) {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	s, err := m.OpenService(service.Name)
	if err != nil {
		s.Close()
		return fmt.Errorf("Could not access %s: %v", service.Name, err)
	}

	sConf, err := s.Config()
	if err != nil {
		return fmt.Errorf("Could not access config: %v", err)
	}

	sConf.StartType = service.StartMode.toSvcStartType()
	sConf.DisplayName = service.DisplayName
	sConf.Description = service.Description
	if err = s.UpdateConfig(sConf); err != nil {
		return fmt.Errorf("Update service failed: %v", err)
	}
	return
}

func (a *App) Remove(name string) (err error) {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	s, err := m.OpenService(name)
	if err != nil {
		return fmt.Errorf("Could not access %s: %v", name, err)
	}
	defer s.Close()

	if err = s.Delete(); err != nil {
		return fmt.Errorf("Failed to remove %s: %v", name, err)
	}

	if err = eventlog.Remove(name); err != nil {
		return fmt.Errorf("Remove event log failed: %s", err)
	}
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

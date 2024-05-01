package main

import (
	"slices"
)

func (a *App) GetPinedServices() (services []Service, err error) {
	names := getPinedServiceNames()
	if len(names) <= 0 {
		return
	}

	allServices, err := a.GetServices()
	if err != nil {
		return
	}

	for _, service := range allServices {
		if slices.ContainsFunc(names, func(name string) bool {
			return service.Name == name
		}) {
			services = append(services, service)
		}
	}
	return
}

func (a *App) TogglePinService(name string) (pined bool, err error) {
	return setPinedServiceName(name)
}

var pinedServiceNames = make([]string, 0)

func initPinedServiceNames() (err error) {
	conf, err := readConfig()
	pinedServiceNames = conf.PinedServiceNames
	return
}

func getPinedServiceNames() (names []string) {
	if len(pinedServiceNames) <= 0 {
		return
	}
	return pinedServiceNames
}

func setPinedServiceName(name string) (add bool, err error) {
	i := slices.Index(pinedServiceNames, name)
	if i < 0 {
		pinedServiceNames = append(pinedServiceNames, name)
		slices.Sort(pinedServiceNames)
		add = true
	} else {
		slices.Delete(pinedServiceNames, i, i+1)
	}

	updateConfig()
	return
}

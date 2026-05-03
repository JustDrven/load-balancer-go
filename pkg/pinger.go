package pkg

import "github.com/go-ping/ping"

type ServiceStatus int

const (
	Success ServiceStatus = iota
	Failed
)

func MakeDiagnostic(addr string) ServiceStatus {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		return Failed
	}

	pinger.Count = 1
	err = pinger.Run()
	if err != nil {
		return Failed
	}

	return Success
}

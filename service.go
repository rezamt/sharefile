// +build windows

package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"
	"net/http"

	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
)

var Elog debug.Log

type sharefileservice struct{}

func (m *sharefileservice) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue
	changes <- svc.Status{State: svc.StartPending}

	srv := registerHttpServer()

	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

loop:
	for {
		select {
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
				// Testing deadlock from https://code.google.com/p/winsvc/issues/detail?id=4
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				// golang.org/x/sys/windows/svc.TestExample is verifying this output.
				testOutput := strings.Join(args, "-")
				testOutput += fmt.Sprintf("-%d", c.Context)
				Elog.Info(1, testOutput)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()
				err := srv.Shutdown(ctx)
				if err != nil {
					Elog.Error(1, "Failed to Shutdown Shareclient Service Proxy")
				}

				break loop
			case svc.Pause:
				changes <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}
			case svc.Continue:
				changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
			default:
				Elog.Error(1, fmt.Sprintf("unexpected control request #%d", c))
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

func runService(name string, isDebug bool) {
	var err error
	if isDebug {
		Elog = debug.New(name)
	} else {
		Elog, err = eventlog.Open(name)
		if err != nil {
			return
		}
	}
	defer Elog.Close()

	Elog.Info(1, fmt.Sprintf("starting %s service", name))
	run := svc.Run
	if isDebug {
		run = debug.Run
	}
	err = run(name, &sharefileservice{})
	if err != nil {
		Elog.Error(1, fmt.Sprintf("%s service failed: %v", name, err))
		return
	}
	Elog.Info(1, fmt.Sprintf("%s service stopped", name))
}

func registerHttpServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Version)

	mux.HandleFunc("/employees", ListEmployees)
	mux.HandleFunc("/roles", ListRoles)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			println("ListenAndServe(): %s", err)
		}

	}()

	return srv
}

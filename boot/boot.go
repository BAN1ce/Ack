package boot

import (
	"context"
	"errors"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/BAN1ce/Ack/pkg/daemon"
)

type IComponent interface {
	Start(context.Context) error
	Stop() error
	GetString() string
}

type App struct {
	component   map[string]IComponent
	stopTimeout time.Duration
	cancel      context.CancelFunc
}

func NewApp() *App {
	return &App{
		component: make(map[string]IComponent, 100),
	}

}

func (a *App) RegisterProvider(provider IComponent) {
	a.component[provider.GetString()] = provider
}

func (a *App) Run() error {
	var ctx context.Context
	ctx, a.cancel = context.WithCancel(context.TODO())
	for _, v := range a.component {
		err := v.Start(ctx)
		if err != nil {
			return err
		}
	}

	daemon.SetSigHandler(func(sig os.Signal) (err error) {

		log.Println("Stopping Application")
		if err := a.Stop(); err != nil {
			log.Printf("Stop Application Fail , Error = %s\n", err.Error())
		}
		log.Println("Stop Application Success")
		os.Exit(0)
		return nil
	}, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Application Running")
	if err := daemon.ServeSignals(); err != nil {
		panic(err)
	}
	return nil
}
func (a *App) Stop() error {

	a.cancel()
	done := make(chan int, 10)
	go func() {
		for _, v := range a.component {
			if err := v.Stop(); err != nil {
				log.Printf("Stop Component Fail , Error = %s\n", err.Error())
			}
		}
		done <- 1
	}()
	stoptimeout := a.stopTimeout
	if stoptimeout == 0 {
		stoptimeout = 10 * time.Second
	}

	timeAfter := time.After(stoptimeout)
	select {
	case <-done:
		return nil
	case <-timeAfter:
		return errors.New("stop timeout")
	}

}

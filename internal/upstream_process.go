package internal

import (
	"errors"
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

type UpstreamProcess struct {
	Started chan struct{}
	cmd     *exec.Cmd
}

func NewUpstreamProcess(name string, arg ...string) *UpstreamProcess {
	slog.Info("Creating upstream process:", name, arg)
	return &UpstreamProcess{
		Started: make(chan struct{}, 1),
		cmd:     exec.Command(name, arg...),
	}
}

func (p *UpstreamProcess) Run() (int, error) {
	slog.Info("Running upstream process")
	p.cmd.Stdin = os.Stdin
	p.cmd.Stdout = os.Stdout
	p.cmd.Stderr = os.Stderr

	slog.Info("Starting command")
	err := p.cmd.Start()
	if err != nil {
		return 0, err
	}

	slog.Info("Command Started", "command", p.cmd.String(), "process", p.cmd.Process, "pid", p.cmd.Process, "process_state", p.cmd.ProcessState.String())
	p.Started <- struct{}{}

	slog.Info("Handling Signals")
	go p.handleSignals()
	err = p.cmd.Wait()

	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		return exitErr.ExitCode(), nil
	}

	return 0, err
}

func (p *UpstreamProcess) Signal(sig os.Signal) error {
	return p.cmd.Process.Signal(sig)
}

func (p *UpstreamProcess) handleSignals() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	sig := <-ch
	slog.Info("Relaying signal to upstream process", "signal", sig.String())
	p.Signal(sig)
}

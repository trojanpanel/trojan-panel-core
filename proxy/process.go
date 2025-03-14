package proxy

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os/exec"
	"sync"
)

type process struct {
	logger *logrus.Logger
	cmdMap *sync.Map
}

func (p *process) isRunning(key string) bool {
	cmd, ok := p.cmdMap.Load(key)
	if !ok {
		return false
	}
	if cmd == nil {
		return false
	}
	if execCmd, ok := cmd.(*exec.Cmd); ok {
		if execCmd.Process == nil {
			return false
		}
		return execCmd.ProcessState == nil
	}
	return false
}

func (p *process) start(key string, name string, arg ...string) error {
	if p.isRunning(key) {
		return nil
	}

	cmd := exec.Command(name, arg...)
	if cmd.Err != nil {
		return cmd.Err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	p.cmdMap.Store(key, cmd)

	go p.handleLogs(stdout, stderr)

	return nil
}

func (p *process) stop(key string) error {
	if !p.isRunning(key) {
		return nil
	}

	cmd, ok := p.cmdMap.Load(key)
	if !ok {
		return fmt.Errorf("process not found")
	}

	execCmd, ok := cmd.(*exec.Cmd)
	if !ok {
		return fmt.Errorf("process stop err")
	}

	if err := execCmd.Process.Kill(); err != nil {
		return err
	}

	if err := execCmd.Process.Release(); err != nil {
		return err
	}

	p.cmdMap.Delete(key)

	return nil
}

func (p *process) handleLogs(stdout, stderr io.ReadCloser) {
	stdoutChan := make(chan string)
	stderrChan := make(chan string)

	readLog := func(r io.Reader, ch chan string) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			logrus.Errorf("Error reading log: %v", err)
		}
		close(ch)
	}

	go readLog(stdout, stdoutChan)
	go readLog(stderr, stderrChan)

	for stdoutChan != nil || stderrChan != nil {
		select {
		case line, ok := <-stdoutChan:
			if !ok {
				stdoutChan = nil
			} else {
				p.logger.Info(line)
			}
		case line, ok := <-stderrChan:
			if !ok {
				stderrChan = nil
			} else {
				p.logger.Error(line)
			}
		}
	}
}

package worker

import (
	"github.com/l0k18/log"
	"github.com/l0k18/stdconn"
	"os"
	"os/exec"
)

type Worker struct {
	*exec.Cmd
	args    []string
	StdConn stdconn.StdConn
}

// Spawn starts up an arbitrary executable 
// file with given arguments and
// attaches a connection to its stdin/stdout
func Spawn(args ...string) (w *Worker) {
	w = &Worker{
		Cmd:  exec.Command(
			args[0], args[1:]...),
		args: args,
	}
	// pass error prints through
	w.Stderr = os.Stdout
	// attach a pipe to the child process out
	cmdOut, err := w.StdoutPipe()
	if err != nil {
		log.ERROR(err)
		return
	}
	// and in
	cmdIn, err := w.StdinPipe()
	if err != nil {
		log.ERROR(err)
		return
	}
	w.StdConn = stdconn.New(
		cmdOut, cmdIn, make(chan struct{}))
	// start the child in another OS 
	// process, which has cores assigned then
	// by the kernel
	err = w.Start()
	if err != nil {
		log.ERROR(err)
		return nil
	} else {
		return
	}
}

// Kill forces the child process to shut 
// down without cleanup
func (w *Worker) Kill() (err error) {
	return w.Process.Kill()
}

// Stop signals the worker to shut down 
// cleanly.
// Note that the worker must have handlers 
// for os.Signal messages.
// It is possible and neater to put a quit 
// method in the IPC API and use the
// quit channel built into the StdConn
func (w *Worker) Stop() (err error) {
	return w.Process.Signal(os.Interrupt)
}
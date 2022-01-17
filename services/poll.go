package services

import "sync"

type Command struct {
	Args     []interface{}
	Callable func(args []interface{}) ([]interface{}, error)
	Return   []interface{}
	Error    error
	Bell     sync.WaitGroup
}

func NewCommand(args []interface{}, callable func(args []interface{}) ([]interface{}, error)) *Command {
	cmd := &Command{
		Args:     args,
		Callable: callable,
		Return:   nil,
		Error:    nil,
		Bell:     sync.WaitGroup{},
	}
	cmd.Bell.Add(1)
	return cmd
}

type Motor struct {
	Running   bool
	Framerate int
	Conveyor  chan *Command
}

var motor *Motor

func poll(motor *Motor) {
	for {
		cmd, ok := <-motor.Conveyor
		if !ok {
			break
		}

		ret, err := cmd.Callable(cmd.Args)
		cmd.Return = ret
		cmd.Error = err
		cmd.Bell.Done()
	}
}
func Loop(fps int) {
	StopLoop()

	motor = &Motor{
		Running:   false,
		Framerate: fps,
		Conveyor:  make(chan *Command),
	}

	go poll(motor)
}
func SyncOneCommand(cmd *Command) {
	if motor == nil {
		return
	}
	motor.Conveyor <- cmd
	cmd.Bell.Wait()
}
func StopLoop() {
	if motor == nil {
		return
	}
	if motor.Conveyor != nil {
		close(motor.Conveyor)
	}
	motor.Running = false
	motor = nil
}

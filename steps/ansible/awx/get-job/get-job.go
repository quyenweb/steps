package main

// test
import (
	"github.com/stackpulse/steps-sdk-go/step"
	"github.com/stackpulse/steps/steps/ansible/awx/base"
)

type Args struct {
	base.Args
	ID string `env:"JOB_ID,required"`
}

type AWXGetJob struct {
	*base.AwxCommand
	args Args
}

func (a *AWXGetJob) Init() error {
	var args Args
	baseCommand, err := base.NewAwxCommand(&args)
	if err != nil {
		return err
	}

	a.AwxCommand = baseCommand
	a.args = args
	return nil
}

func (a *AWXGetJob) Run() (int, []byte, error) {
	return a.Execute([]string{"jobs", "get", a.args.ID})
}

func main() {
	step.Run(&AWXGetJob{})
}

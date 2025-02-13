package kbd

import (
	"time"

	"github.com/barista-run/barista/bar"
	"github.com/barista-run/barista/base/value"
	"github.com/barista-run/barista/outputs"
	"github.com/barista-run/barista/timing"
)

// Module represents a kbd bar module.
type Module struct {
	outputFunc value.Value // of func(string) bar.Output
	scheduler  *timing.Scheduler
}

// Named constructs an instance of the kbd module.
func New() *Module {
	m := &Module{scheduler: timing.NewScheduler()}

	m.RefreshInterval(5 * time.Second)

	// Default output
	m.Output(func(layout string) bar.Output {
		return outputs.Text(layout)
	})

	return m
}

func (m *Module) RefreshInterval(interval time.Duration) *Module {
	m.scheduler.Every(interval)
	return m
}

// Output configures a module to display the output of a user-defined function.
func (m *Module) Output(outputFunc func(string) bar.Output) *Module {
	m.outputFunc.Set(outputFunc)
	return m
}

func (m *Module) Stream(s bar.Sink) {
	outputFunc := m.outputFunc.Get().(func(string) bar.Output)

	for {
		layout, _ := getKeyboardLayout()
		s.Output(outputFunc(layout))

		<-m.scheduler.C
	}
}

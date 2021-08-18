package core

import (
	"fmt"

	"github.com/Fish-pro/multi-cluster-proxy/pkg/config"
	"github.com/Fish-pro/multi-cluster-proxy/pkg/mlog"
)

// Runnable run components after initialized
type Runnable interface {
	Run(stop <-chan struct{})
	Initialize() error
}

type Core struct {
	*config.Config

	RunnableComponents map[string]Runnable
}

var _ Runnable = &Core{}

func New(c *config.Config) *Core {
	return &Core{Config: c, RunnableComponents: make(map[string]Runnable)}
}

func (c *Core) IntegrateWith(name string, r Runnable) {
	if name == "" {
		mlog.Fatal("component name is empty")
	}

	if r == nil {
		mlog.Fatal("component %s is nil", name)
	}

	if _, dup := c.RunnableComponents[name]; dup {
		mlog.Fatal("component %s already set up", name)
	}

	c.RunnableComponents[name] = r
}

func (c *Core) Run(stop <-chan struct{}) {

	for name, c := range c.RunnableComponents {
		// Run action is non blocking
		go c.Run(stop)
		mlog.Info("component %s already started", name)
	}

	<-stop
}

// Initialize init all components of core, block util initialization
// completed or error occurred.
func (c *Core) Initialize() error {
	if len(c.RunnableComponents) < 1 {
		mlog.Fatal("cube have no components to initialize")
	}

	// initialize runnable components of cube
	for name, c := range c.RunnableComponents {
		if err := c.Initialize(); err != nil {
			return fmt.Errorf("component %s initialize failed: %v", name, err)
		}
		mlog.Info("component %s already initialized", name)
	}

	return nil
}

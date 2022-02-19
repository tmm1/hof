// This package provides a context for Tasks
// and a registry for their usage in flows.
package context

import (
	gocontext "context"
	"io"
	"sync"

	"cuelang.org/go/cue"
)

// A Context provides context for running a task.
type Context struct {
	RootValue cue.Value
	GoContext gocontext.Context

	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	Value   cue.Value
	Error   error

  // debug / internal
  DebugTasks bool
  Verbosity  int

  // Per context, copyable down the stack?
  TaskRegistry *sync.Map

  // BOOKKEEPING
  Tasks *sync.Map

  // Middleware

  // how can the below become middleware, extensions, plugin?

  // Global (for this context, tbd shared) lock around CUE evaluator 
  CUELock  *sync.Mutex

  // map of cue.Values
  ValStore *sync.Map

  // map of chan?
  Mailbox  *sync.Map

  // channels for
  // - stats & progress

}

func New() *Context {
  return &Context{
    GoContext: gocontext.Background(),
    CUELock: new(sync.Mutex),
    ValStore: new(sync.Map),
    Mailbox: new(sync.Map),
    TaskRegistry: new(sync.Map),
    Tasks: new(sync.Map),
  }
}

func Copy(ctx *Context) *Context {
  return &Context{
    RootValue: ctx.RootValue,
    GoContext: ctx.GoContext,

    Stdin:   ctx.Stdin,
    Stdout:  ctx.Stdout,
    Stderr:  ctx.Stderr,

    DebugTasks: ctx.DebugTasks,
    Verbosity: ctx.Verbosity,

    CUELock: ctx.CUELock,
    Mailbox: ctx.Mailbox,
    ValStore: ctx.ValStore,

    TaskRegistry: ctx.TaskRegistry,
    Tasks: ctx.Tasks,
  }
}

// Register registers a task for cue commands.
func (C *Context) Register(key string, f RunnerFunc) {
	C.TaskRegistry.Store(key, f)
}

// Lookup returns the RunnerFunc for a key.
func (C *Context) Lookup(key string) RunnerFunc {
	v, ok := C.TaskRegistry.Load(key)
	if !ok {
		return nil
	}
	return v.(RunnerFunc)
}

// consider adding here... a
// global registry of named channels

// A RunnerFunc creates a Runner.
type RunnerFunc func(v cue.Value) (Runner, error)

// A Runner defines a task type.
type Runner interface {
	// Runner runs given the current value and returns a new value which is to
	// be unified with the original result.
	Run(ctx *Context) (results interface{}, err error)
}

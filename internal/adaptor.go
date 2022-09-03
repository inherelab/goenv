package internal

import (
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/sysutil"
	"github.com/inherelab/goenv"
)

// Adaptor consts
const (
	AdaptorAuto  = "auto"
	AdaptorGoEnv = "goenv"
	AdaptorBrew  = "brew"
	AdaptorScoop = "scoop"
)

type baseAdaptor struct {
	name string
	opts *CallOpts
}

// ApplyOpFunc handle
func (a *baseAdaptor) ApplyOpFunc(fn OpFunc) {
	if fn != nil {
		fn(a.opts)
	}
}

// Name of adaptor
func (a *baseAdaptor) Name() string {
	return a.name
}

// CallOpts struct
type CallOpts struct {
	LibDir string
	Yes    bool
}

// OpFunc define
type OpFunc func(opts *CallOpts)

// Adaptor interface
type Adaptor interface {
	Name() string
	List() error
	ApplyOpFunc(fn OpFunc)
	Switch(ver string) error
	Install(ver string) error
	Update(ver string) error
	Uninstall(ver string) error
}

// MakeAdaptor instance
func MakeAdaptor(adaptor string) (Adaptor, error) {
	if adaptor == goenv.ModeAuto {
		adaptor = autoSelectAdaptor(adaptor)
		cliutil.Cyanln("auto select adaptor:", adaptor)
	}

	switch adaptor {
	case goenv.ModeBrew:
		return NewBrewAdaptor(), nil
	case goenv.ModeScoop:
		return NewScoopAdaptor(), nil
	case goenv.ModeGoEnv:
		return NewGoEnvAdaptor(), nil
	default:
		return nil, errorx.Rawf("unsupported adaptor %q", adaptor)
	}
}

func autoSelectAdaptor(name string) string {
	// is Windows
	if sysutil.IsWindows() {
		if sysutil.HasExecutable(goenv.ModeScoop) {
			return goenv.ModeScoop
		}
		return goenv.ModeGoEnv
	}

	if sysutil.HasExecutable(goenv.ModeBrew) {
		return goenv.ModeBrew
	}
	return goenv.ModeGoEnv
}

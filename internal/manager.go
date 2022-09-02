package internal

import (
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/sysutil"
	"github.com/inherelab/goenv"
)

// MakeAdaptor instance
func MakeAdaptor(adaptor string) (Adaptor, error) {
	switch autoSelectAdaptor(adaptor) {
	case goenv.ModeBrew:
		return NewBrewAdaptor(), nil
	case goenv.ModeGoEnv:
		return &GoEnvAdaptor{}, nil
	default:
		return nil, errorx.Rawf("unsupported adaptor %q", adaptor)
	}
}

func autoSelectAdaptor(name string) string {
	if name != goenv.ModeAuto {
		return name
	}

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

// CallOpts struct
type CallOpts struct {
	LibDir string
	Yes    bool
}

// Adaptor interface
type Adaptor interface {
	List() error
	WithOptions(opts *CallOpts) Adaptor
	Install(ver string) error
	Update(ver string) error
	Uninstall(ver string) error
	Switch(ver string) error
}

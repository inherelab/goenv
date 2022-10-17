package internal

import (
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/sysutil"
)

// Adaptor name consts: scoop, winget
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

func newBaseAdaptor(name string, opts *CallOpts) baseAdaptor {
	return baseAdaptor{name: name, opts: opts}
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
	// Available() error // TODO check available
}

// MakeAdaptor instance
func MakeAdaptor(adaptor string) (Adaptor, error) {
	if adaptor == AdaptorAuto {
		adaptor = autoSelectAdaptor(adaptor)
		cliutil.Cyanln("TIP: auto select a adaptor:", adaptor)
	}

	switch adaptor {
	case AdaptorBrew:
		return NewBrewAdaptor(), nil
	case AdaptorScoop:
		return NewScoopAdaptor(), nil
	case AdaptorGoEnv:
		return NewGoEnvAdaptor(), nil
	case "winget":
		return NewStdAdaptor(adaptor), nil
	default:
		return nil, errorx.Rawf("unsupported adaptor %q", adaptor)
	}
}

func autoSelectAdaptor(name string) string {
	// is Windows
	if sysutil.IsWindows() {
		if sysutil.HasExecutable(AdaptorScoop) {
			return AdaptorScoop
		}
		return AdaptorGoEnv
	}

	if sysutil.HasExecutable(AdaptorBrew) {
		return AdaptorBrew
	}
	return AdaptorGoEnv
}

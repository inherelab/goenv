package internal

import (
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/errorx"
	"github.com/inherelab/goenv"
)

// EnvManager struct
type EnvManager struct {
	adaptor string
}

func NewEnvManager(adaptor string) *EnvManager {
	return &EnvManager{
		adaptor: adaptor,
	}
}

func (m *EnvManager) CreateAdaptor() (Adaptor, error) {
	switch m.adaptor {
	case goenv.ModeBrew:
		return &BrewAdaptor{}, nil
	case goenv.ModeGoEnv:
		return &GoEnvAdaptor{}, nil
	default:
		return nil, errorx.Rawf("unsupported adaptor %q", m.adaptor)
	}
}

func (m *EnvManager) Install(ver string) error {
	a, err := m.CreateAdaptor()
	if err != nil {
		return err
	}

	return a.Install(ver)
}

type CallOpts struct {
	LibDir string
	Yes    bool
}

// Adaptor interface
type Adaptor interface {
	WithOptions(opts *CallOpts) Adaptor
	Install(ver string) error
	Update(ver string) error
	Uninstall(ver string) error
	Switch(ver string) error
}

// GoEnvAdaptor struct
type GoEnvAdaptor struct {
	opts *CallOpts
}

func (a *GoEnvAdaptor) WithOptions(opts *CallOpts) Adaptor {
	a.opts = opts
	return a
}

func (a *GoEnvAdaptor) Switch(ver string) error {
	cliutil.Infoln("TODO un-supported")
	return nil
}

func (a *GoEnvAdaptor) Install(ver string) error {
	cliutil.Infoln("TODO un-supported")
	return nil
}

func (a *GoEnvAdaptor) Update(ver string) error {
	cliutil.Infoln("TODO un-supported")
	return nil
}

func (a *GoEnvAdaptor) Uninstall(ver string) error {
	cliutil.Infoln("TODO un-supported")
	return nil
}

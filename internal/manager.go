package internal

import (
	"github.com/gookit/goutil/errorx"
	"github.com/inherelab/goenv"
)

// EnvManager struct
type EnvManager struct {
	adaptor string
}

// NewEnvManager instance
func NewEnvManager(adaptor string) *EnvManager {
	return &EnvManager{
		adaptor: adaptor,
	}
}

// MakeAdaptor instance
func (m *EnvManager) MakeAdaptor() (Adaptor, error) {
	switch m.adaptor {
	case goenv.ModeBrew:
		return NewBrewAdaptor(), nil
	case goenv.ModeGoEnv:
		return &GoEnvAdaptor{}, nil
	default:
		return nil, errorx.Rawf("unsupported adaptor %q", m.adaptor)
	}
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

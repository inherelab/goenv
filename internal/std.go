package internal

import "github.com/gookit/goutil/cliutil"

// HandleFn define
type HandleFn func(ver string) error

// StdAdaptor struct
type StdAdaptor struct {
	baseAdaptor

	ListFn    func() error
	SwitchFn  HandleFn
	InstallFn HandleFn
	UpdateFn  HandleFn
	// UninstallFn handler
	UninstallFn HandleFn
}

// NewStdAdaptor instance
func NewStdAdaptor(name string) *StdAdaptor {
	return &StdAdaptor{
		baseAdaptor: newBaseAdaptor(name, &CallOpts{}),
	}
}

func (a *StdAdaptor) WithOptions(fns ...OpFunc) *StdAdaptor {
	for _, fn := range fns {
		fn(a.opts)
	}
	return a
}

func (a *StdAdaptor) List() error {
	if a.ListFn == nil {
		cliutil.Infoln("TODO ...")
		return nil
	}
	return a.ListFn()
}

func (a *StdAdaptor) Switch(ver string) error {
	if a.SwitchFn == nil {
		cliutil.Infoln("TODO ...")
		return nil
	}
	return a.SwitchFn(ver)
}

func (a *StdAdaptor) Install(ver string) error {
	if a.InstallFn == nil {
		cliutil.Infoln("TODO ...")
		return nil
	}
	return a.InstallFn(ver)
}

func (a *StdAdaptor) Update(ver string) error {
	if a.UpdateFn == nil {
		cliutil.Infoln("TODO ...")
		return nil
	}
	return a.UpdateFn(ver)
}

func (a *StdAdaptor) Uninstall(ver string) error {
	if a.UninstallFn == nil {
		cliutil.Infoln("TODO ...")
		return nil
	}
	return a.UninstallFn(ver)
}

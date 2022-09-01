package go_pkg_template

import (
	"github.com/gookit/color"
	"github.com/gookit/goutil/stdutil"
	"github.com/gookit/goutil/sysutil"
)

func Example() {
	color.Infoln(sysutil.BinDir())
	color.Infoln(stdutil.GoVersion())
}

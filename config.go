package goenv

var defConf = `
# 
# config for https://github.com/inherelab/goenv
# author: https://github.com/inhere
#

# mode allow: goenv, brew
mode: brew

# on use brew mode
brew_lib_dir: /usr/local/opt

# custom add env map
env_map:
  GO_ROOT: path/to

current:
  version: 1.16
  path: /usr/local/opt/go@1.16

# use goenv mode.
# https://golang.org/dl
dl_host: https://golang.google.cn/dl
# install go dir.
install_dir: /usr/local/go

`

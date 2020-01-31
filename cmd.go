package main

import (
	"bytes"
	"fmt"
)

var (
	cmds = make(map[string]Cmd)
	keys = make([]string, 0)
)

type Cmd struct {
	Args   uint32
	Handle Handler
}

type Handler interface {
	help() string
	handle([]string) (string, error)
}

func addCommand(cmd string, args uint32, handle Handler) {
	if _, ok := cmds[cmd]; !ok {
		cmds[cmd] = Cmd{
			Args:   args,
			Handle: handle,
		}
		keys = append(keys, cmd)
	}
}

func help() string {
	buffer := bytes.NewBufferString(fmt.Sprint("Use:", "\n"))
	for _, key := range keys {
		buffer.WriteString(fmt.Sprint("\te[nvi] ", cmds[key].Handle.help(), "\n"))
	}
	buffer.WriteString(`
Developers:
	Author: Yenole
	Email : Netxy@vip.qq.com
	Github: https://github.com/yenole
	Home  : https://github.com/yenole/envi.git`)
	return buffer.String()
}

type InitCmd struct {
}

func (cmd *InitCmd) handle(args []string) (string, error) {
	return "", nil
}

func (_ *InitCmd) help() string {
	return `init [bash|zsh]`
}

type AddCmd struct{}

func (cmd *AddCmd) handle(args []string) (string, error) {
	return "", nil
}

func (_ *AddCmd) help() string {
	return `add [name] path`
}

type DelCmd struct {
}

func (cmd *DelCmd) handle(args []string) (string, error) {
	return "", nil
}

func (_ *DelCmd) help() string {
	return `del name`
}

type ViewCmd struct {
}

func (cmd *ViewCmd) handle(args []string) (string, error) {
	return "", nil
}

func (_ *ViewCmd) help() string {
	return `view`
}

type AliasCmd struct {
}

func (cmd *AliasCmd) handle(args []string) (string, error) {
	return "", nil
}

func (_ *AliasCmd) help() string {
	return `alias name [command]`
}

func init() {
	addCommand("init", 0, &InitCmd{})
	addCommand("add", 1, &AddCmd{})
	addCommand("del", 1, &DelCmd{})
	addCommand("view", 0, &ViewCmd{})
	addCommand("alias", 2, &AliasCmd{})
}

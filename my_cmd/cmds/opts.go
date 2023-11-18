package cmds

import "sort"

type Cmd struct {
	ID          int
	Description string
	Execute     func()
}

type cmds []Cmd

func (c cmds) Len() int {
	return len(c)
}

func (c cmds) Less(i, j int) bool {
	return c[i].ID < c[j].ID
}

func (c cmds) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

var (
	Cmds cmds
)

func InstallCmd(cmd Cmd) {
	Cmds = append(Cmds, cmd)
	sort.Sort(Cmds)
}

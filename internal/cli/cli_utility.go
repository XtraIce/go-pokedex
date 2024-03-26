package cli

var prevCmds []string
var cmdIdx int = -1

func init() {
	prevCmds = []string{""}
}

// AddPrevCmd adds a command to the prevCmds slice.
// If the command is empty or the same as the last one, it will not be added to the list.
func AddPrevCmd(cmd string) {
	if cmd == "" || cmd == prevCmds[len(prevCmds)-1] {
		return
	}
	prevCmds = append(prevCmds, cmd)
	trimPrevCmds()
}

// GetPrevCmds returns the prevCmds slice.
func GetPrevCmds() []string {
	return prevCmds
}

// trimPrevCmds trims the prevCmds slice to the last 10 commands.
func trimPrevCmds() {
	if len(prevCmds) > 10 {
		prevCmds = prevCmds[len(prevCmds)-10:]
	}
	cmdIdx = len(prevCmds)
}

// TraversePrevCmds traverses the prevCmds slice in the backward direction.
// It decrements the cmdIdx and returns the command.
func TraversePrevCmds() string {
	cmdIdx = max(0, cmdIdx-1)
	return prevCmds[cmdIdx]
}

// TraverseNextCmds traverses the prevCmds slice in the forward direction.
// It increments the cmdIdx and returns the command.
func TraverseNextCmds() string {
	cmdIdx = min(cmdIdx+1, len(prevCmds)-1)
	return prevCmds[cmdIdx]
}

package cli

var prevCmds []string
var cmdIdx int = -1

func init() {
	prevCmds = []string{""}
}

// add a command to the prevCmds slice
func AddPrevCmd(cmd string) {
	if cmd == "" || cmd == prevCmds[len(prevCmds)-1] { //if the command is the same as the last one, don't add it to the list
		return
	}
	prevCmds = append(prevCmds, cmd)
	//log.Default().Println("prevCmds: ", prevCmds)
	trimPrevCmds()
}

// return the prevCmds slice
func GetPrevCmds() []string {
	return prevCmds
}

// trim the prevCmds slice to the last 10 commands
func trimPrevCmds() {
	if len(prevCmds) > 10 {
		prevCmds = prevCmds[len(prevCmds)-10:]
	}
	cmdIdx = len(prevCmds) //reset the cmdIdx
}

// traverse the prevCmds slice in the backward direction
func TraversePrevCmds() string {
	cmdIdx = max(0, cmdIdx-1) //decrement the cmdIdx
	return prevCmds[cmdIdx]   //return the command
}

// traverse the prevCmds slice in the forward direction
func TraverseNextCmds() string {
	cmdIdx = min(cmdIdx+1, len(prevCmds)-1) //increment the cmdIdx
	return prevCmds[cmdIdx]                 //return the command
}

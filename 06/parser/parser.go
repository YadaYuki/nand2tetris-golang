package parser

// CommandType is type of command
type CommandType int

const (
	aCommand CommandType = iota
	cCommand
	lCommand
)

func (command CommandType) String() string {
	switch command {
	case aCommand:
		return "A_COMMAND"
	case cCommand:
		return "C_COMMAND"
	case lCommand:
		return "L_COMMAND"
	default:
		return "Unknown"
	}
}

func GetCommandType(commandStr string) CommandType {
	return aCommand
}

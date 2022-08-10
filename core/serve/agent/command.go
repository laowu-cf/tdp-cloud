package agent

import "errors"

type CommandPayload struct {
	Content          string `binding:"required"`
	Username         string `binding:"required"`
	CommandType      string `binding:"required"`
	WorkingDirectory string `binding:"required"`
	Timeout          uint   `binding:"required"`
}

func RunCommand(addr string, data *CommandPayload) error {

	node, ok := AgentPool[addr]

	if !ok {
		return errors.New("客户端已断开")
	}

	return node.Pod.Write(SocketData{
		Action:  "runCommand",
		Method:  "request",
		Payload: data,
	})

}

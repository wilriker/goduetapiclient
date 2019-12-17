package connection

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/wilriker/goduetapiclient/commands"
	"github.com/wilriker/goduetapiclient/connection/initmessages"
)

const (
	// TaskCanceledException is the name of a remote exception to be checked for
	TaskCanceledException = "TaskCanceledException"
	// SocketDirectory is the default directory in which DSF-related UNIX sockets reside
	SocketDirectory = "/var/run/dsf"
	// SocketFile is the default UNIX socket file for DuetControlServer
	SocketFile = "dcs.sock"
	// FullSocketPath is the default fully-qualified path to the UNIX socket for DuetControlServer
	FullSocketPath = SocketDirectory + "/" + SocketFile
)

// BaseConnection provides common functionalities for more concrete implementations
type BaseConnection struct {
	socket  net.Conn
	decoder *json.Decoder
	id      int64
}

// Connect establishes a connecton to the given UNIX socket file
func (bc *BaseConnection) Connect(initMessage initmessages.ClientInitMessage, socketPath string) error {
	var err error
	bc.socket, err = net.Dial("unix", socketPath)
	if err != nil {
		return err
	}
	bc.decoder = json.NewDecoder(bc.socket)

	sim, err := bc.receiveServerInitMessage()
	if err != nil {
		return nil
	}

	if sim.Version < initmessages.ExpectedServerVersion {
		return errors.New(fmt.Sprintf("Incompatible API version (expected %d got %d)", initmessages.ExpectedServerVersion, sim.Version))
	}

	bc.id = sim.Id

	err = bc.Send(initMessage)
	if err != nil {
		return err
	}

	br, err := bc.ReceiveResponse()
	if err != nil {
		return err
	}
	if !br.IsSuccess() {
		return errors.New(fmt.Sprintf("Could not set connection type %s (%s: %s)", initMessage.GetMode(), br.GetErrorType(), br.GetErrorMessage()))
	}

	return nil
}

// Close the UNIX socket connection
func (bc *BaseConnection) Close() {
	if bc.socket != nil {
		bc.socket.Close()
		bc.socket = nil
	}
}

// PerformCommand performs an arbitrary command
func (bc *BaseConnection) PerformCommand(command commands.Command) (commands.Response, error) {
	err := bc.Send(command)
	if err != nil {
		return nil, err
	}
	br, err := bc.ReceiveResponse()
	if err != nil {
		return nil, err
	}
	if !br.IsSuccess() {
		if br.GetErrorType() == TaskCanceledException {
			return br, errors.New(br.GetErrorMessage())
		}
		return br, errors.New(fmt.Sprintf("InternalServerError: %s, %s, %s", command.GetCommand(), br.GetErrorType(), br.GetErrorMessage()))
	}
	return br, nil
}

// ReceiveResponse receives a deserialized response from the server
func (bc *BaseConnection) ReceiveResponse() (commands.Response, error) {
	br := commands.BaseResponse{}
	err := bc.Receive(&br)
	if err != nil {
		return nil, err
	}
	return &br, nil
}

// receiveServerInitMessage returns the ServerInitMessage
func (bc *BaseConnection) receiveServerInitMessage() (*initmessages.ServerInitMessage, error) {
	sim := initmessages.ServerInitMessage{}
	err := bc.Receive(&sim)
	if err != nil {
		return nil, err
	}
	return &sim, nil
}

// Receive a deserialized object
func (bc *BaseConnection) Receive(responseContainer interface{}) error {
	if err := bc.decoder.Decode(responseContainer); err != nil {
		return err
	}
	return nil
}

// ReceiveJson returns a server response as a JSON string
func (bc *BaseConnection) ReceiveJson() (string, error) {
	var raw json.RawMessage
	err := bc.Receive(&raw)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

// Send arbitrary data
func (bc *BaseConnection) Send(data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// log.Println(string(b))
	_, err = bc.socket.Write(b)
	return err
}

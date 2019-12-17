package commands

import "github.com/wilriker/goduetapiclient/types"

// AddUserSession registers a new user session
type AddUserSession struct {
	BaseCommand
	// AccessLevel of this session
	AccessLevel types.AccessLevel
	// SessionType of this session
	SessionType types.SessionType
	// Origin of this session. For remote sessions this equals the remote IP address
	Origin string
	// OriginPort corresponds to the identifier of the origin.
	// If it is a remote session it is the remote port
	// else it defaults to the PID of the current process
	OriginPort int
}

// NewAddUserSession creates a new instance of AddUserSession
func NewAddUserSession(access types.AccessLevel, t types.SessionType, origin string, op int) *AddUserSession {
	return &AddUserSession{
		BaseCommand: *NewBaseCommand("AddUserSession"),
		AccessLevel: access,
		SessionType: t,
		Origin:      origin,
		OriginPort:  op,
	}
}

// RemoveUserSession to remove an existing user session
type RemoveUserSession struct {
	BaseCommand
	// Id of the user session to remove
	Id int
}

// NewRemoveUserSession to create a correctly initialized instance of RemoveUserSession
func NewRemoveUserSession(id int) *RemoveUserSession {
	return &RemoveUserSession{
		BaseCommand: *NewBaseCommand("RemoveUserSession"),
		Id:          id,
	}
}

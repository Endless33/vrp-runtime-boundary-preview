package boundary

// Session represents the public view of a VRP logical session.
//
// This structure intentionally exposes only observable runtime
// information. Protected runtime state is never exported.
type Session struct {
	ID string

	State SessionState

	ActiveTransport string
}

// SessionState describes the observable session lifecycle.
type SessionState string

const (

	// SessionCreated indicates that the logical session exists.
	SessionCreated SessionState = "CREATED"

	// SessionActive indicates that the session is operational.
	SessionActive SessionState = "ACTIVE"

	// SessionRecovering indicates that transport recovery is in progress.
	SessionRecovering SessionState = "RECOVERING"

	// SessionClosed indicates that the logical session has ended.
	SessionClosed SessionState = "CLOSED"
)

// NewSession creates a new public session representation.
func NewSession(id string) *Session {

	return &Session{
		ID:    id,
		State: SessionCreated,
	}
}

// Activate marks the session as active.
func (s *Session) Activate() {

	s.State = SessionActive
}

// BeginRecovery marks the session as recovering.
func (s *Session) BeginRecovery() {

	s.State = SessionRecovering
}

// AttachTransport updates the active transport.
func (s *Session) AttachTransport(name string) {

	s.ActiveTransport = name
}

// SwitchTransport replaces the active transport while preserving
// the logical session.
func (s *Session) SwitchTransport(name string) {

	s.ActiveTransport = name
}

// Close marks the session as closed.
func (s *Session) Close() {

	s.State = SessionClosed
}

package boundary

type SessionState string

const (
	SessionCreated    SessionState = "CREATED"
	SessionActive     SessionState = "ACTIVE"
	SessionRecovering SessionState = "RECOVERING"
	SessionClosed     SessionState = "CLOSED"
)

type Session struct {
	ID              string
	State           SessionState
	ActiveTransport string
}

func NewSession(id string) *Session {
	return &Session{
		ID:    id,
		State: SessionCreated,
	}
}

func (s *Session) Activate() {
	s.State = SessionActive
}

func (s *Session) BeginRecovery() {
	s.State = SessionRecovering
}

func (s *Session) AttachTransport(name string) {
	s.ActiveTransport = name
}

func (s *Session) Close() {
	s.State = SessionClosed
}

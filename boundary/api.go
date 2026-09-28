package boundary

// Version identifies the public Runtime Boundary API version.
const Version = "v0.1.0"

// API exposes the public Runtime Boundary.
type API struct{}

// New creates a new Runtime Boundary API.
func New() *API {
	return &API{}
}

// Name returns the public API name.
func (a *API) Name() string {
	return "VRP Runtime Boundary"
}

// Version returns the API version.
func (a *API) Version() string {
	return Version
}

// DesignPrinciple returns the public architectural principle.
func (a *API) DesignPrinciple() string {
	return "SESSION ≠ TRANSPORT"
}

// CreateSession creates a new public session.
func (a *API) CreateSession(id string) *Session {
	return NewSession(id)
}

// CreateTransport creates a new public transport.
func (a *API) CreateTransport(
	id string,
	t TransportType,
) *Transport {

	return NewTransport(id, t)
}

// CreateEvidence creates a public evidence record.
func (a *API) CreateEvidence(
	scenario string,
	verdict Verdict,
	message string,
) Evidence {

	return NewEvidence(
		scenario,
		verdict,
		message,
	)
}

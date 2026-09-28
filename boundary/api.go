package boundary

const Version = "v0.1.0"

type API struct{}

func New() *API {
	return &API{}
}

func (a *API) Name() string {
	return "VRP Runtime Boundary"
}

func (a *API) Version() string {
	return Version
}

func (a *API) DesignPrinciple() string {
	return "SESSION ≠ TRANSPORT"
}

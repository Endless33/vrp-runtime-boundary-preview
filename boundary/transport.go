package boundary

type TransportType string

const (
	TransportUDP  TransportType = "UDP"
	TransportTCP  TransportType = "TCP"
	TransportQUIC TransportType = "QUIC"
)

type TransportState string

const (
	TransportDetached TransportState = "DETACHED"
	TransportAttached TransportState = "ATTACHED"
	TransportActive   TransportState = "ACTIVE"
	TransportLost     TransportState = "LOST"
)

type Transport struct {
	ID    string
	Type  TransportType
	State TransportState
}

func NewTransport(id string, t TransportType) *Transport {
	return &Transport{
		ID:    id,
		Type:  t,
		State: TransportDetached,
	}
}

func (t *Transport) Attach() {
	t.State = TransportAttached
}

func (t *Transport) Activate() {
	t.State = TransportActive
}

func (t *Transport) Lose() {
	t.State = TransportLost
}

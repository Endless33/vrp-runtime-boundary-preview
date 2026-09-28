package boundary

import "time"

type Verdict string

const (
	VerdictPass Verdict = "PASS"
	VerdictFail Verdict = "FAIL"
)

type Evidence struct {
	Scenario  string
	Verdict   Verdict
	Message   string
	Timestamp time.Time
}

func NewEvidence(scenario string, verdict Verdict, message string) Evidence {
	return Evidence{
		Scenario:  scenario,
		Verdict:   verdict,
		Message:   message,
		Timestamp: time.Now().UTC(),
	}
}

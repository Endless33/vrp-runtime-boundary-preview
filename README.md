VRP Runtime Boundary Preview

VRP Runtime Boundary Preview is an executable demonstration of the boundary between external runtime events and VRP-style validation decisions.

The goal is not to expose private core implementation details.

The goal is to show how external events can be admitted, rejected, or preserved through an observable decision surface.

---

Purpose

Most networking systems treat transport instability as an application-level recovery event.

VRP explores a different model:

Session identity is canonical.
Transport is replaceable.

This repository demonstrates the public-facing boundary:

External Event
    ↓
Validation Boundary
    ↓
Decision Surface
    ↓
Observable Verdict

---

What This Preview Demonstrates

This preview demonstrates:

- Replay rejection
- Authority rollback rejection
- Runtime recovery preservation
- Transport migration preservation
- Canonical state preservation across boundary events

---

What This Preview Does Not Expose

This repository does not expose:

- Private VRP core logic
- Production runtime internals
- Cryptographic implementation details
- Commercial integration code
- Private pilot adapters

This is a boundary preview, not a core release.

---

Run

go run ./cmd/runtime-boundary-demo

Expected result:

FINAL_VERDICT=BOUNDARY_PRESERVED

---

Example Output

=== VRP RUNTIME BOUNDARY PREVIEW ===
Boundary: external events -> validation decisions -> observable verdicts

EVENT: replay packet
DECISION: rejected
VERDICT=REPLAY_REJECTED

EVENT: authority rollback
DECISION: rejected
VERDICT=AUTHORITY_ROLLBACK_REJECTED

EVENT: runtime recovery
DECISION: preserved
VERDICT=RECOVERY_PRESERVED

EVENT: transport migration
DECISION: preserved
VERDICT=TRANSPORT_MIGRATION_PRESERVED

FINAL_VERDICT=BOUNDARY_PRESERVED

---

Machine Readable Validation Evidence

The runtime boundary preview generates both human-readable and machine-readable validation artifacts.

Run:

go run ./cmd/runtime-boundary-demo

Inspect:

cat validation-report.json

Example:

{
  "project": "vrp-runtime-boundary-preview",
  "final_verdict": "BOUNDARY_PRESERVED"
}

The generated report contains:

- Validation events
- Observable decisions
- Validation verdicts
- Canonical state preservation results
- Final validation outcome

The objective is to make runtime behavior:

- Observable
- Reproducible
- Machine-readable
- Independently inspectable

---

Boundary Principle

The boundary model is designed around a simple idea:

Failures may enter the system.

Invalid execution must not.

The validation boundary determines whether an external runtime event may affect canonical execution state.

---

Validation Artifact Flow

External Event
        ↓
Validation Decision
        ↓
Observable Verdict
        ↓
validation-report.json

The validation report provides a structured artifact that can be inspected by engineers, tooling, CI pipelines, or external evaluators.

---

Related Projects

Main validation repository:

https://github.com/Endless33/vrp-validation-kit

The validation kit contains:

- Executable validation artifacts
- Failure models
- Invariant mappings
- External validation guides
- Pilot documentation

Research repository:

https://github.com/Endless33/jumping-vpn-preview

The research repository contains architecture discussions, runtime models, validation evidence, and continuity-oriented protocol research.

---

Current Status

Current observable scenarios:

- Replay rejection
- Authority rollback rejection
- Runtime recovery preservation
- Transport migration preservation

Current validation artifact:

validation-report.json

Current expected outcome:

FINAL_VERDICT=BOUNDARY_PRESERVED

The repository exists to expose observable runtime behavior without exposing proprietary implementation details.
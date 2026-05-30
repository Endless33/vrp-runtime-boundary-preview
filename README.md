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

Boundary Principle

The boundary model is designed around a simple idea:

Failures may enter the system.
Invalid execution must not.

The validation boundary determines whether an external runtime event may affect canonical execution state.

---

Related Project

Main validation repository:

https://github.com/Endless33/vrp-validation-kit

The validation kit contains executable artifacts, failure models, invariant mappings, external validation guides, and pilot documentation.
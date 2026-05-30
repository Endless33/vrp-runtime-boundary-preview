Decision Surface

Purpose

The decision surface describes how runtime events become observable validation decisions.

The objective is to make boundary behavior understandable without exposing private implementation details.

---

Decision Model

Input Event
    ↓
Classification
    ↓
Boundary Decision
    ↓
Verdict

---

Example Decisions

Event| Decision| Verdict
Replay Packet| Rejected| REPLAY_REJECTED
Authority Rollback| Rejected| AUTHORITY_ROLLBACK_REJECTED
Runtime Recovery| Preserved| RECOVERY_PRESERVED
Transport Migration| Preserved| TRANSPORT_MIGRATION_PRESERVED

---

Design Goal

The design goal is not to reveal internal implementation.

The design goal is to expose enough observable behavior for external engineers to evaluate assumptions.

---

Relationship to Validation Kit

This repository is a boundary preview.

The main validation environment is:

https://github.com/Endless33/vrp-validation-kit

That repository contains executable validation releases, failure models, invariant mappings, and pilot documentation.
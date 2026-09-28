# AGENTS.md — plugin-builder

Standalone plugin repo for the `builder` kind (`kind:builder`). The plugin is a
Go module at `candy/plugin-builder/` (module path
`github.com/opencharly/plugin-builder/candy/plugin-builder`); the root
`charly.yml` only declares `discover: candy` so the repo is a project and its
candy is scanned.

Canonical files:

- `candy/plugin-builder/charly.yml` — the `plugin-builder:` candy entity
  (`plugin:` block, `plan:` check).
- `candy/plugin-builder/plugin.go` — the kind provider (`Invoke(OpLoad)` decode
  → canonical JSON) and `NewProvider()`/`NewMeta()`.
- `candy/plugin-builder/schema/builder.cue` — the self-contained
  `#BuilderInput`.
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.
- `README.md` — user overview only; never agent guidance.

## Load these skills first (R0)

- `/charly-image:image` — box/builder configuration, the `builder:` vocabulary,
  and the box dependency graph. Load before changing the kind's decoded fields.
- `/charly-internals:plugin` — the plugin authoring reference: the `plugin:`
  block, the `kind` provider class, the per-plugin CUE-schema contract. Load
  before touching the provider or schema.
- `/charly-internals:git-workflow` — before any git/PR action.

## Build / validate / test

- `go build ./...` in `candy/plugin-builder/` — compile the plugin module.
- `go test ./...` in `candy/plugin-builder/` — the plugin's Go tests.
- `charly box validate` at the repo root — the structural check (the candy +
  `plugin:` block, CUE schema).
- The merge gate is the **org-wide** `charly/pr-validator` (required check
  `validate / validate`, defined in `opencharly/.github`); this repo has **no**
  per-repo candy gate.
- The changed path is exercised by any box/deploy composing a `builder:` node.

## Modify this repo

- Edit the `plugin-builder:` candy entity, the Go source, and
  `schema/builder.cue` **together** — the schema is the single source for the
  kind's `params/` struct and a faithful reproduction of the core `#Builder`
  wire keys. A wire-key change must keep the host's canonicalisation through
  `spec.Builder` valid.
- The entity name is the node KEY, never a body field.

## Landing

- PR-only. Every change lands through a pull request; the org-required
  `charly/pr-validator` validates the diff and body and arms native auto-merge on
  PASS. Direct pushes to `main` are blocked.
- History lives in `CHANGELOG/` (written by `tag-on-merge` at merge time); the PR
  body IS the changelog.
- The authoritative rulebook is the umbrella `AGENTS.md` in
  `opencharly/opencharly` and `charly/AGENTS.md` in the charly repo. Do not
  restate its rules here.

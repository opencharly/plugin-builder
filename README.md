# plugin-builder

The `builder` kind for OpenCharly — the multi-stage builder vocabulary
(`pixi` / `npm` / `cargo` / `aur` / `bootstrap`) as a plugin kind.

The provider dispatches via the `pb` `Invoke(OpLoad)` envelope: it decodes the
authored `builder:` entity into its core spec type and re-marshals it as
canonical JSON. It serves itself in both placements (compiled-in or
out-of-process) — no kit contract is needed because kinds are `pb`-shape.

## What it provides

| Capability | Surface |
|---|---|
| `kind:builder` | the `builder:` kind entity — a builder definition with its detect rules, cache mounts, phases, install command, and artifacts |

The authored body is validated at runtime against the self-contained
`#BuilderInput` (`schema/builder.cue`), spliced by the host.

## How to use it

Compose the plugin candy in a box's `candy:` list:

```yaml
- '@github.com/opencharly/plugin-builder/candy/plugin-builder:<tag>'
```

Then author a `builder:` entity in a project:

```yaml
my-builder:
  builder:
    kind: layer
    detect_file: package.json
    install_command:
      npm: npm ci
    path_contribution: [node_modules/.bin]
```

## Layout

- `candy/plugin-builder/` — the plugin module: `plugin.go` (the kind provider +
  `NewProvider()`/`NewMeta()`), `schema/builder.cue` (the self-contained
  `#BuilderInput`), `params/cue_types_gen.go`, `cmd/serve/main.go`.
- `candy/plugin-builder/charly.yml` — the `plugin-builder:` candy entity.
- `charly.yml` — the root project manifest (`discover: candy`).
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.

## Related

- Owning skill: `/charly-image:image` — box/builder configuration and the box
  dependency graph. This candy carries no `skill:` entity of its own; the gap is
  tracked in
  [opencharly/opencharly#291](https://github.com/opencharly/opencharly/issues/291).
- `/charly-internals:plugin` — the plugin/provider model, including the `kind`
  provider class.
- [`opencharly/charly`](https://github.com/opencharly/charly) — the charly CLI.

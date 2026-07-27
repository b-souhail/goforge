<h2 align="center">📄 Blueprint</h2>

The **blueprint** (`blueprint.yaml`) is the single source of truth for your project. Every command GoForge runs — `add` and `setup`, and soon `remove` — reads and/or writes this file.

Instead of remembering flags or re-answering the same questions, you configure your project once, and GoForge generates (or regenerates) it from that configuration.

---

<h2 align="center">📌 Where It Comes From</h2>

`blueprint.yaml` is created automatically by `goforge init`, and updated automatically by `goforge add` (and, once finished, `goforge remove`). You're not expected to write it by hand, but you *can* edit it directly before running `goforge setup` if you want finer control.

---

<h2 align="center">🧱 Structure</h2>

```yaml
path: /home/user/projects
name: myapp
modules:
  - order
  - client
layers:
  - entities
  - usecases
  - interfaces_adapters
  - frameworks_drivers
resources:
  - name: http
    params:
      middleware: true
      middlewares: [cors, limiter]
      websocket: true
      hub: true
  - name: mysql
    params: {}
```

<div align="center">

| Field | Type | Description |
| --- | --- | --- |
| `path` | string | Absolute path where the project directory was created |
| `name` | string | Project name — also the generated directory name and Go module name |
| `modules` | list of string | Business modules to scaffold (e.g. `order`, `client`) — each gets its own Entities, and eventually Use Cases |
| `layers` | list of string | Internal architectural layer identifiers used by the generator |
| `resources` | list | Resources added to the project (`http`, `mysql`, ...), each with its own `params` |

</div>

> 🚧 The `layers` identifiers are being aligned with the strict Clean Architecture folder names (`entities`, `usecases`, `frameworks_drivers`) described in [Architecture](architecture.md) — the field itself is stable, its default values are still being finalized.

---

<h2 align="center">🧩 Resources & Params</h2>

Each entry under `resources` stores a resource name and the answers you gave when running `goforge add`. Those `params` are exactly what gets passed back to the resource when `goforge setup` generates its files — so the blueprint fully determines what gets generated, with no hidden state. See [Resources](resources.md) for what each resource stores and generates.

---

<h2 align="center">🔁 Blueprint as Source of Truth</h2>

Because `goforge setup` regenerates the project *from* the blueprint, it stays consistent even if you:

- Manually edit `blueprint.yaml` before running `setup`
- Add a resource or a module, then re-run `setup` to generate what's new
- Share the blueprint with teammates to reproduce the exact same project structure

> 🚧 Removing a resource or module from the blueprint — and the generated files that go with it — is handled by the `goforge remove` command, which is still under active development. See [Commands](commands.md) and [Roadmap](roadmap.md).

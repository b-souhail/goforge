<h2 align="center">📄 Blueprint<h2>

The **blueprint** (`blueprint.yaml`) is the single source of truth for your project. Every command GoForge runs — `add`, `setup`, and eventually `delete` — reads and writes this file.

Instead of remembering flags or re-answering the same questions, you configure your project once, and GoForge generates (or regenerates) it from that configuration.

<h2 align="center">📌 Where It Comes From <h2>

The blueprint is created automatically by `goforge init`, and updated automatically by `goforge add` (and later `goforge --delete / -D <resource/modules>`). You're not expected to write it by hand, but you *can* edit it directly before running `goforge setup` if you want finer control.

<h2 align="center">🧱 Structure <h2>

```yaml
path: /home/user/projects/myapp
name: myapp
modules: "order,client"
layers:
  - name: domain
    directory: []
  - name: application
    directory: []
  - name: infrastructure
    directory: []
  - name: delivery
    directory: []
resources:
  - name: http
    params:
      middleware: false
      websocket: false
  ...

```

<div align="center">

| Field | Type | Description |
| --- | --- | --- |
| `path` | string | Absolute path to the project on disk |
| `name` | string | Project name |
| `modules` | string | Comma-separated list of business modules (e.g. `user,post,comment`) |
| `layers` | list | The architectural layers to generate (`domain`, `application`, `infrastructure`, `delivery`) |
| `resources` | list | Resources added to the project (`http`, `mysql`...), each with its own `params` |

</div>

<h2 align="center">🧩 Resources & Params<h2>

Each entry under `resources` stores a resource name and the answers you gave when running `goforge add`. For example, adding `http` with middlewares and a WebSocket hub produces:

```yaml
resources:
  - name: http
    params:
      middleware: true
      middlewares: [cors, limiter]
      websocket: true
      hub: true
```

These `params` are exactly what gets passed back to the resource when `goforge setup` generates its files — so the blueprint fully determines what gets generated, with no hidden state.

<h2 align="center">🔁 Blueprint as Source of Truth <h2>

Because `goforge setup` regenerates the project *from* the blueprint, it stays consistent even if you:

- Manually edit `blueprint.yaml` before running `setup`
- Add a resource, then re-run `setup` to generate only what's new
- Share the blueprint with teammates to reproduce the exact same project structure

> 🚧 Removing a resource or module from the blueprint (and the generated files that go with it) will be handled by the upcoming `--delete` / `-D` flag — see [Commands](commands.md).

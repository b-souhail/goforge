<h2 align="center">📖 Commands</h2>

<h2 align="center">✅ Available Today</h2>

<div align="center">

| Command | Description |
|---|---|
| `goforge init [name]` | Create a new project directory and its `blueprint.yaml` |
| `goforge add <resource>` | Add a resource (`http`, `mysql`, ...) to the blueprint |
| `goforge add --modules <list>` | Add one or more business modules to the blueprint |
| `goforge setup` | Generate the project from `blueprint.yaml` — modules and resources |

</div>

```bash
goforge init myapp
cd myapp
goforge add --modules order,client
goforge add http
goforge add mysql
goforge setup
```

See [Modules](modules.md) and [Resources](resources.md) for what each of these composes.

---

<h2 align="center">🔭 Stay Tuned — Upcoming Commands</h2>

GoForge's command surface is still growing. Here's what's coming next:

<div align="center">

| Command | Purpose | Priority |
|---|---|:---:|
| `goforge remove <scope>` | Remove a module or resource from the blueprint and its generated files | 🔴 High — in development |
| `goforge make <scope>` | Targeted generation — e.g. add a single use case to one module, instead of a whole module or resource | — |
| `goforge init -a, --archi <clean-strict\|hexa\|onion\|...>` | Choose the architecture at project creation | 🟢 Low — after GoForge's initial release |
| `goforge analyze` | Analyze dependencies and flag Clean Architecture (Dependency Rule) violations | 🟢 Low |

</div>

Follow progress on the [Roadmap](roadmap.md).

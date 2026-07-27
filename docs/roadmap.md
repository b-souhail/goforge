<h2 align="center">🗺️ Roadmap</h2>

GoForge is a work in progress. This page tracks what's implemented, what's actively being built, and what's planned.

---

<h2 align="center">🏛️ Architecture</h2>

<div align="center">

| Item | Status |
|---|:---:|
| Clean Architecture — Use Case oriented (strict, Uncle Bob) — default | 🚧 In Development |
| Clean Architecture — DDD oriented | 🚧 In Development — follows the strict approach |
| Hexagonal Architecture | 💬 Community Interest |
| Onion Architecture | 💬 Community Interest |
| Event-Driven Architecture | 💬 Community Interest |
| Architecture as a configurable choice (`goforge init -a/--archi`) | 🟢 Low priority — after GoForge's initial release |

</div>

See [Architecture](architecture.md).

---

<h2 align="center">🌐 Resources</h2>

<div align="center">

| Resource | Status |
|---|:---:|
| http | ✅ Available |
| mysql | ✅ Available |
| sqlite / postgresql | 💬 Coming with v1 |
| redis | 🚧 Template in progress |
| rabbitmq | 🚧 Template in progress |
| docker | 🚧 Template in progress |

</div>

See [Resources](resources.md).

---

<h2 align="center">⚙️ Commands</h2>

<div align="center">

| Command | Status / Priority |
|---|:---:|
| `goforge init` / `add` / `setup` | ✅ Available |
| `goforge remove <scope>` | 🔴 High — in development |
| `goforge make <scope>` — targeted generation (e.g. one use case in one module) | 🟢 Low — after initial release |
| `goforge init -a/--archi <archi>` | 🟢 Low — after initial release |
| `goforge analyze` — dependency / Dependency Rule violation analysis | 🟢 Low |

</div>

See [Commands](commands.md).

---

<h2 align="center">💬 Community Feedback</h2>

Architectures marked *Community Interest* are shaped by what developers ask for. If there's a pattern you'd like to see supported, open an issue or a discussion on the repository.

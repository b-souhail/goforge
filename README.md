<div align="center">

<img src="docs/assets/go_at_forge.png" width="430" alt="GoForge">

# 🔨 GoForge

### Forge production-ready Go applications from a declarative Blueprint.

<p>

![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)
![Open Source](https://img.shields.io/badge/Open%20Source-Yes-success?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Work%20In%20Progress-orange?style=for-the-badge)

</p>

**GoForge generates Go applications from a single declarative Blueprint — following a strict, literal implementation of Robert C. Martin's Clean Architecture.**

[⚙️ Installation](/docs/installation.md) •
[🚀 Quick Start](/docs/getting-started.md) •
[📖 Commands](/docs/commands.md) •
[📄 Blueprint](/docs/blueprint.md) •
[🧩 Modules](/docs/modules.md) •
[🌐 Resources](/docs/resources.md) •
[🏛️ Architecture](/docs/architecture.md) •
[💡 Examples](/docs/examples.md) •
[🗺️ Roadmap](/docs/roadmap.md) •
[❓ FAQ](/docs/faq.md)

</div>

---

## What is GoForge?

GoForge is **not a framework**. It doesn't execute your application, doesn't replace the Go standard library, and doesn't lock you into a specific ecosystem.

Instead, it generates a project from a declarative **Blueprint** (`blueprint.yaml`), so you own every line of generated code from the start.

```bash
goforge init myapp
cd myapp
goforge add --modules order,client
goforge add http
goforge add mysql
goforge setup
```

> **Generate architecture. Write business logic yourself.**

The goal: spend less time on boilerplate, more time on business rules, and keep the project maintainable as it grows.

---

## 🏛️ Architecture

By default, GoForge generates a **strict, Use Case oriented Clean Architecture** — built directly from Robert C. Martin's (Uncle Bob) own layer and component names: **Entities**, **Use Cases** (Interactors, Boundaries, Request/Response Models), **Interface Adapters** (Controllers, Presenters, Gateways), **Frameworks & Drivers**. Not a `domain/application/infrastructure/delivery` reinterpretation — the terminology maps directly back to [the source material](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

<div align="center">

| Architecture | Status |
|--------------|:---:|
| Clean Architecture (Use Case oriented, strict) | 🚧 In Development — default |
| Clean Architecture (DDD oriented) | 🚧 In Development — follows the strict approach |
| Hexagonal Architecture | 💬 Community Interest |
| Onion Architecture | 💬 Community Interest |
| Event-Driven Architecture | 💬 Community Interest |

</div>

> The long-term vision is to make architecture a configurable choice instead of a fixed decision.

➡️ **Read more:** [/docs/architecture.md](/docs/architecture.md) · [/docs/roadmap.md](/docs/roadmap.md)

---

## 🌐 Resources & 🧩 Modules

**Modules** and **resources** are the two parts that compose your system with GoForge — modules are your business concepts (`user`, `order`, ...), resources are the technical capabilities they run on.

<div align="center">

| Resource | Status |
|---|:---:|
| 🌐 http | ✅ |
| 🗄️ mysql | ✅ |
| 🐘 sqlite / postgresql | 💬 coming with v1 |
| ⚡ redis | 🚧 template in progress |
| 🐇 rabbitmq | 🚧 template in progress |
| 🐳 docker | 🚧 template in progress |

</div>

➡️ **Read more:** [/docs/modules.md](/docs/modules.md) · [/docs/resources.md](/docs/resources.md)

---

## 📄 Blueprint

Every GoForge project is driven by a single file, `blueprint.yaml` — the single source of truth for modules, resources, and architecture. It's written automatically by `goforge init` and `goforge add`, but can be edited directly for finer control before running `goforge setup`.

```text
Blueprint
      │
      ▼
Modules + Resources
      │
      ▼
Templates
      │
      ▼
Generated Project
```

➡️ **Read more:** [/docs/blueprint.md](/docs/blueprint.md)

---

## 📖 Commands

<div align="center">

| Command | Description |
|---|---|
| `goforge init [name]` | Create a project and its `blueprint.yaml` |
| `goforge add <resource>` \| `--modules <list>` | Add a resource or module(s) |
| `goforge setup` | Generate the project from the blueprint |

</div>

More is on the way: `goforge remove` (🔴 in development), `goforge make`, `goforge init -a/--archi`, `goforge analyze` — full breakdown in [Commands](/docs/commands.md) and [Roadmap](/docs/roadmap.md).

---

## Who is GoForge for?

- **🎓 Students** — learn Clean Architecture by exploring generated projects instead of starting from an empty directory.
- **👨‍💻 Freelancers** — bootstrap MVPs quickly while keeping a maintainable architecture.
- **🚀 Startups** — cut project setup time and focus on delivering product features.
- **🏢 Teams** — standardize project structures and reduce onboarding time.
- **❤️ Go enthusiasts** — experiment with architectures, templates, and reusable resources.

---

## Philosophy

GoForge believes developers should own their code. The generator automates repetitive setup — it doesn't hide your application behind another framework. Everything it generates is designed to be maintainable, extensible, deterministic, and easy to understand.

**GoForge generates architecture, not business logic.**

---

## Status

GoForge is under active, work-in-progress development and hasn't had its initial release yet. Entities generation and the `http`/`mysql` resources are usable today; the Use Case layer, additional resources, and additional architectures are actively being built — see the full [Roadmap](/docs/roadmap.md).

---

<div align="center">

Licensed under [MIT](LICENSE).

</div>

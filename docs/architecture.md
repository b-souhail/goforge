<h2 align="center">🏛️ Architecture</h2>

GoForge scaffolds every project around **Clean Architecture**: a layered design where dependencies always point inward, toward the domain.

``` yaml

myapp/
├── blueprint.yaml
├── cmd/
│   └── main.go
└── internal/
    ├── domain/
    │   ├── entity/
    │   └── repository/
    ├── application/
    │   ├── DTOs/
    │   └── usecases/
    ├── infrastructure/
    │   ├── repository/
    │   └── database/
    └── delivery/
        │── http/
        └── websocket/

```

---
<h2 align="center"> 🧭 The Four Layers </h2>

<div align="center"> 

| Layer | Role | Depends on |
| --- | --- | --- |
| 🧬 **Domain** | Business entities and repository *interfaces*. The core of the application. | Nothing |
| ⚙️ **Application** | Use cases and DTOs. Orchestrates domain logic for the outside world. | Domain |
| 🔧 **Infrastructure** | Concrete implementations: database repositories, external services. | Domain, Application |
| 🌐 **Delivery** | Entry points: HTTP handlers, routes, WebSocket, middlewares. | Application |

The rule is simple: **inner layers never know about outer layers.** The domain has no idea an HTTP server or a MySQL database exists — it only exposes interfaces that the infrastructure layer implements.

</div>

---

<h2 align="center"> 🧬 Domain Layer </h2>

Contains the business entities and the repository interfaces they need, without any framework or database detail.

- `domain/entity/` — plain Go structs representing business objects
- `domain/repository/` — interfaces describing how entities are persisted, with no implementation detail

---

<h2 align="center"> ⚙️ Application Layer </h2>

Contains the use cases: the actual business operations of the application, plus the DTOs used to move data in and out of them.

- `application/usecases/` — one use case per business operation, built on top of domain repository interfaces
- `application/DTOs/` — request/response shapes exposed to the delivery layer

---

<h2 align="center">🔧 Infrastructure Layer </h2>

Implements the interfaces defined in the domain layer, and wires up anything external: databases, caches, third-party services.

- `infrastructure/repository/` — concrete implementations of the domain repository interfaces
- `infrastructure/database/` — database connection and configuration (MySQL, PostgreSQL...)

---

<h2 align="center">🌐 Delivery Layer </h2>

The entry point of the application — where the outside world talks to your use cases.

- `delivery/http/` — routes and HTTP handlers ...
- `delivery/websocket/` — optional WebSocket hub and connections

---

<h2 align="center"> 🔄 Why This Structure </h2>

- **Testability** — the domain and application layers can be tested without a database or an HTTP server.
- **Replaceability** — swapping MySQL for PostgreSQL, or the HTTP framework for another one, only touches the infrastructure/delivery layers.
- **Clarity** — every file has one obvious place to live, which keeps the codebase predictable as it grows.

> 🌱 The four layers (`domain`, `application`, `infrastructure`, `delivery`) are currently fixed, but the architecture is designed to support additional or alternative architectures in the future without breaking existing projects.

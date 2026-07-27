<h2 align="center">🏛️ Architecture</h2>

<p align="center">
By default, GoForge generates a strict, <strong>Use Case oriented Clean Architecture</strong> —
built directly from Robert C. Martin's (Uncle Bob) own layer and component names:
<strong>Entities</strong>, <strong>Use Cases</strong> (Interactors, Input/Output Boundaries,
Request/Response Models), <strong>Interface Adapters</strong> (Controllers, Presenters, Gateways),
<strong>Frameworks &amp; Drivers</strong>.
</p>

> 📖 **Sources**
> - [*The Clean Architecture*](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) — blog.cleancoder.com, 2012
> - *Clean Architecture: A Craftsman's Guide to Software Structure and Design* — Robert C. Martin, 2017
> - [*Architecture: The Lost Years*](https://www.youtube.com/watch?v=WpkDN78P884) — Ruby Midwest 2011 keynote, Robert C. Martin
> - [*conferences By Uncle Bob*]

---

<h2 align="center">📌 Status</h2>

<div align="center">

| Architecture | Status |
|--------------|:---:|
| Clean Architecture — Use Case oriented (strict, Uncle Bob) | 🚧 In Development — default |
| Clean Architecture — DDD oriented | 🚧 In Development — already started, will follow the strict approach |
| Hexagonal Architecture | 💬 Community Interest |
| Onion Architecture | 💬 Community Interest |
| Event-Driven Architecture | 💬 Community Interest |

</div>

The strict Use Case oriented approach is what `goforge init` generates today. The DDD-oriented Clean Architecture is already underway and will land after the strict approach is complete. Hexagonal, Onion, and Event-Driven are being considered based on community interest — the long-term goal is for architecture to be a configurable choice, not a fixed decision.

---

📝 *A detailed breakdown of the architecture (layers, generated structure, use case wiring) will be added here directly.*

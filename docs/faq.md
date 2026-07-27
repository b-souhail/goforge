<h2 align="center">❓ FAQ</h2>

<h3>Is GoForge a framework?</h3>

No. GoForge doesn't run your application, doesn't replace the Go standard library, and doesn't lock you into an ecosystem. It generates a project once — after that, it's your code.

<h3>What architecture does GoForge generate by default?</h3>

A strict, Use Case oriented **Clean Architecture**, following Robert C. Martin's own terminology (Entities, Use Cases/Interactors, Boundaries, Gateways). See [Architecture](architecture.md).

<h3>Why not the common `domain/application/infrastructure/delivery` layout?</h3>

That layout is a popular interpretation of Clean Architecture, but it isn't Uncle Bob's own naming. GoForge generates code that maps directly back to the source material, so there's no ambiguity about what each piece is for or where it came from.

<h3>Will GoForge support other architectures?</h3>

Yes — a DDD-oriented Clean Architecture is already in development and will follow the strict approach. Hexagonal, Onion, and Event-Driven are being considered based on community interest. See [Roadmap](roadmap.md).

<h3>What happens if I edit `blueprint.yaml` by hand?</h3>

It's supported. `blueprint.yaml` is the single source of truth — `goforge setup` regenerates the project from whatever is in the file. See [Blueprint](blueprint.md).

<h3>Can I remove a resource or module once it's generated?</h3>

Not yet. `goforge remove` is in active development — it's a high priority. See [Commands](commands.md) and [Roadmap](roadmap.md).

<h3>Is GoForge production-ready?</h3>

Not yet — it hasn't had its initial release. Expect command names, blueprint fields, and generated file layouts to evolve. Installation and getting-started guides will be written once that release lands.

<h3>How can I contribute or request a feature?</h3>

Open an issue or a discussion on the repository — feedback directly shapes what's built next, especially [Roadmap](roadmap.md) items like additional architectures and resources.

<h2 align="center">🧩 Modules</h2>

**Modules** and **[Resources](resources.md)** are the two parts that compose your system in GoForge.

A module is a business concept in your application — `user`, `order`, `invoice`. Declaring your modules is how you tell the forger (the tool) what your application is actually made of, so it can generate and later manage the pieces that depend on that knowledge — routes, controllers, and so on.

```bash
goforge add --modules user,order
```

This registers `user` and `order` in `blueprint.yaml`. Running `goforge setup` then generates each module's Entities — see [Architecture](architecture.md).

More module-level generation (Use Cases, targeted additions via the upcoming `goforge make`) is on the way — see [Roadmap](roadmap.md).

<h2 align="center"> 🌐 Resources</h2>

**Resources** are optional building blocks you attach to a project: an HTTP server, a database, a cache, a WebSocket layer. Unlike the four core architecture layers, resources are added on demand with `goforge add <resource>`.

---

<h2 align="center">⚙️ How Resources Work</h2>

Every resource implements the same small contract, which is what lets GoForge treat `http`, `mysql`, or any future resource identically:

- **Name** — the identifier used in `goforge add <name>` and in the blueprint
- **Questions** — an optional set of interactive questions asked when the resource is added
- **Config** — turns your answers into `params` stored in the blueprint
- **Files** — the list of template files to generate, based on those `params`

Because resources are self-contained, adding a new one to GoForge doesn't require touching the core generation workflow — it's registered once and immediately available through `goforge add`.

---

<h2 align="center"> 📦 Available Resources</h2>

| Resource | Description | Status |
|---|---|:---:|
| 🌐 **http** | HTTP server with routes, optional middlewares (CORS, rate limiter) and optional WebSocket support | ✅ |
| 🗄️ **mysql** | MySQL database connection and configuration | ✅ |
| 🐘 **postgresql** | PostgreSQL database connection and configuration | 🚧 |
| ⚡ **redis** | Redis integration for caching | 🚧 |
| 🐳 **docker** | Docker configuration based on the project's resources | 🚧 |

---

## 🌐 The `http` Resource

Adding `http` walks you through a short interactive setup:

```bash
goforge add http
```

```text
? Do you want customized middlewares?  Yes
  ↳ Select middlewares: cors, limiter

? Do you want websocket?  Yes
  ↳ Enable websocket hub?  Yes
```

Depending on your answers, `goforge setup` will generate:

- `internal/delivery/http/routes.go` — always generated
- `internal/delivery/http/middlewares/*.go` — one file per selected middleware
- `internal/delivery/websocket/ws.go` — if WebSocket is enabled
- `internal/delivery/websocket/hub.go` and `connection.go` — if the WebSocket hub is enabled

---

## 🗄️ The `mysql` Resource

Adding `mysql` — it generates a ready-to-configure database connection:

```bash
goforge add mysql
```

Generates:

- `internal/infrastructure/database/mysql.go`

---

## ➕ Adding a Resource

```bash
goforge add <resource>
```

- Fails if the resource doesn't exist in the registry
- Fails if the resource was already added to the blueprint
- On success, stores the resource and its answers in `goforge.yaml`

## ➖ Removing a Resource

```bash
goforge -D <resource>
# or
goforge --delete <resource> 
```

Removes the resource from the blueprint and cleans up the files it generated, without touching the rest of the project.

> 🚧 This flag is planned and will be availble in the mvp beta .

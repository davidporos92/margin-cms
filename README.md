# MarginCMS

A small, modular CMS with a Go API and a React admin. It's a pet project built in the open: one admin user, one content type (markdown posts), with revisions, an activity log and optimistic concurrency.

The API is designed contract-first. [`openapi/`](openapi/) is the single source of truth. The Go server and the TypeScript types are generated from it.

I'm writing up the build as a blog series: [Building MarginCMS](https://davidporos92.github.io/posts/building-margincms-part-1-contract-first-code-later/).

> **Status:** early work in progress. The API serves `GET /api/v1/healthz`. Every other operation returns `501 Not Implemented` until it's built.

## Stack

| Area | Tools |
| --- | --- |
| API contract | OpenAPI 3, split into small files and bundled with [Redocly CLI](https://redocly.com/docs/cli/) |
| API | Go, [chi](https://github.com/go-chi/chi), [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) strict server, [go-envconfig](https://github.com/sethvargo/go-envconfig) |
| Dev loop | Docker Compose, [Air](https://github.com/air-verse/air) hot reload, [Prism](https://stoplight.io/open-source/prism) mock server, Redocly preview docs |
| Web types | [openapi-typescript](https://openapi-ts.dev/) |
| Planned | PostgreSQL with pgx, [Ent](https://entgo.io/), [Atlas](https://atlasgo.io/) migrations, [testcontainers-go](https://golang.testcontainers.org/); a React + TypeScript admin |

## Requirements

- [Docker](https://docs.docker.com/get-docker/) with Docker Compose v2
- [Go](https://go.dev/dl/) 1.27.1 or newer, to run the API or regenerate code outside Docker
- [Node.js](https://nodejs.org/) 22 or newer with npm, for the spec tooling
- `make`

oapi-codegen and Air are pinned as Go tools in [`api/go.mod`](api/go.mod), so there's nothing extra to install. Run them with `go tool oapi-codegen` and `go tool air`.

## Getting started

```sh
git clone https://github.com/davidporos92/margin-cms.git
cd margin-cms
docker compose up --build
```

The first start takes a while, because Go downloads modules and compiles Air. Later starts reuse the cached volumes.

| Service | URL | What it is |
| --- | --- | --- |
| `api` | http://localhost:8080/api/v1 | The Go API, rebuilt on every save |
| `apidocs` | http://localhost:8081 | API docs with a "Try it" console |
| `apimock` | http://localhost:8082 | Prism mock server, responses generated from the spec |

Check that the API is up:

```sh
curl http://localhost:8080/api/v1/healthz
```

The mock supports Prism's `Prefer` header, so you can request a specific response, for example `Prefer: code=412`.

## Common tasks

Install the npm tooling first with `make install-tools`.

| Command | What it does |
| --- | --- |
| `make openapi-lint` | Lint the spec |
| `make openapi-bundle` | Bundle the split spec into `openapi/dist/openapi.bundled.yaml` |
| `make openapi-generate` | Bundle, then regenerate the TypeScript types and the Go server |
| `make openapi-docs` | Run the docs preview without Docker |
| `make openapi-mock` | Run the Prism mock without Docker |
| `cd api && make test` | Run the Go tests |
| `cd api && make build` | Build a static binary to `api/build/api` |

Generated code is committed. After changing the spec, run `make openapi-generate` and commit the spec and the generated files together.

## Configuration

The API reads its configuration from environment variables. Every setting has a default that works for local development. See [`api/.env.example`](api/.env.example) for the full list: listen address, timeouts and CORS.

## Project layout

```
api/                Go API
  cmd/margincms/    entry point
  internal/api/     generated code (do not edit)
  internal/config/  config loading
  internal/server/  handlers and server options
infra/              Dockerfiles for the dev environment
openapi/            API contract, see openapi/README.md
  oapi-codegen/     Go code generation configs
web/                React admin (generated API types so far)
compose.yaml        dev environment
```

## License

[MIT](LICENSE).

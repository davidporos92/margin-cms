# MarginCMS API contract

The source of truth for the API, split into small files. Edit these, never the bundle.

```
openapi/
  openapi.yaml              root: info, servers, tags, security, and a $ref index of everything below
  redocly.yaml              lint rules
  paths/                    one file per URL path (all methods for that path)
  components/
    schemas/                data shapes: Post, PostPage, Problem, ...
    parameters/             reusable query, path and header params: Limit, Cursor, IfMatch, ...
    headers/                ETag
    requestBodies/          PostCreateRequest, PostUpdateRequest, ...
    responses/              named responses: PostResponse, NotFoundResponse, ...
  dist/openapi.bundled.yaml generated single file (git-ignored)
```

## Commands

```sh
npx @redocly/cli lint openapi/openapi.yaml
npx @redocly/cli bundle openapi/openapi.yaml -o openapi/dist/openapi.bundled.yaml
```

Tools that follow multi-file `$ref`s can read `openapi.yaml` directly (`openapi-typescript`, Redocly, Prism). Go's `oapi-codegen` is easier to point at the bundle, so run `bundle` first and generate from `dist/openapi.bundled.yaml`.

## Conventions

- File name equals component name: `components/schemas/Post.yaml` defines `Post`.
- Reference with relative paths (`../components/schemas/Post.yaml` from `paths/`, `./Post.yaml` from another schema). Never use `#/components/...` inside split files.
- New thing used by one endpoint only? Still give it a file if it is a schema, parameter, request body or response. Inline only tiny one-off fragments.
- A new component also needs a line in the matching section of the root `openapi.yaml`. Bundling only pulls in what the root lists or what a path reaches.
- Error responses are named after what happened (`NotFoundResponse`, `PreconditionFailedResponse`), not after the status code, and all use the `Problem` schema.
- Change the spec in its own PR before changing code. CI should lint, bundle, regenerate Go and TS, and fail on a diff.

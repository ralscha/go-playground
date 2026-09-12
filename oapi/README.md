# OpenAPI client

The generator is pinned as a Go tool dependency in `go.mod`. Regenerate the client
after changing `api.yml` or `cfg.yml`:

```sh
go generate ./...
go test ./...
```

Run `go run .` to fetch a random fact from the public API.

`github.com/dprotaso/go-yit` remains pinned to the revision required by
`vmware-labs/yaml-jsonpath`: newer revisions change their node types to YAML v4,
while `yaml-jsonpath` v0.3.2 still requires YAML v3. A blanket `go get -u` of the
generator currently breaks generation; recheck this constraint when upgrading.

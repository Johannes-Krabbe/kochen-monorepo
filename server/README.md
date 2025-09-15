# Kochen Server

## API Documentation

This project uses [swaggo/swag v2](https://github.com/swaggo/swag) to generate OpenAPI v3.1 documentation.

### Generating Documentation

```bash
swag init --dir cmd/server,internal --output docs --parseDependency --parseInternal --parseDepth 1 --ot go,json --generatedTime --v3.0
```

**Prerequisites:**
- Install swag v2: `go install github.com/swaggo/swag/v2/cmd/swag@latest`

**Note:** If you encounter missing go.sum entries during documentation generation, run:
```bash
go mod download golang.org/x/text
go mod tidy
```

### Generated Files

- `docs/swagger.json` - OpenAPI v3.1 specification
- `docs/docs.go` - Go package for programmatic access

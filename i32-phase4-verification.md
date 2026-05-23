# Issue 32 Phase 4 Verification

Date verified: 2026-05-23

Commands:

```sh
goimports -w .
go build -v ./...
go test ./...
PORT=18080 go run ./cmd/aklapi
curl 'http://localhost:18080/api/v1/addr?addr=500%20Queen%20Street'
curl 'http://localhost:18080/api/v1/rr?addr=500%20Queen%20Street'
curl 'http://localhost:18080/api/v1/rrext?addr=500%20Queen%20Street'
```

Results:

- `go build -v ./...`: passed.
- `go test ./...`: passed.
- `/api/v1/addr?addr=500%20Queen%20Street`: `200 OK`.

Sanitized address response:

```json
{"items":[{"id":"12342478585","address":"500 Queen Street, Auckland Central"}]}
```

- `/api/v1/rr?addr=500%20Queen%20Street`: address lookup succeeded, then the
  upstream collection-detail page returned `406`.
- `/api/v1/rrext?addr=500%20Queen%20Street`: address lookup succeeded, then the
  upstream collection-detail page returned `406`.

Sanitized rubbish response:

```json
{"error":"collection API returned status code: 406"}
```

The browser HAR provided for the same workflow shows the same upstream behavior:

- Address search calls to `/nextapi/property?...` returned `200`.
- The selected address detail request returned `406`:

```text
GET /en/rubbish-recycling/rubbish-recycling-collections/rubbish-recycling-collection-days/12342478585.html?_rsc=<redacted>
```

Because the upstream collection-detail page returns `406` even for the real
browser session, the address-token fix restores `/api/v1/addr` but cannot fully
restore `/api/v1/rr` or `/api/v1/rrext` from this environment while the upstream
detail page rejects the request.

No HAR, cookie, bearer token, or captured secret is committed.

# Issue 32 Phase 1 Findings

Date verified: 2026-05-22

The reusable token source is the public Auckland Council collection-day page:

```text
GET https://www.aucklandcouncil.govt.nz/en/rubbish-recycling/rubbish-recycling-collections/rubbish-recycling-collection-days.html
```

The token-source request uses no request body. It works with the browser-style
document headers already represented by `setBrowserDocumentHeaders`:

```text
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
Accept-Language: en-NZ,en;q=0.9,en-US;q=0.8
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
Sec-Fetch-Site: none
Sec-Fetch-Mode: navigate
Sec-Fetch-User: ?1
Sec-Fetch-Dest: document
Sec-CH-UA: "Chromium";v="135", "Not.A/Brand";v="8"
Sec-CH-UA-Mobile: ?0
Sec-CH-UA-Platform: "macOS"
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36
```

The response content type is:

```text
text/html; charset=utf-8
```

The page HTML contains Next Flight script chunks in `self.__next_f.push(...)`.
One chunk includes `SessionTokenProvider` state with an escaped JSON field:

```text
initialToken\":\"<JWT bearer token>\"
```

The response parse is to scan the HTML/Next Flight stream for the first escaped
`initialToken` string and unescape that token value for use as a bearer token.

The token is reusable for the address endpoint:

```text
GET https://experience.aucklandcouncil.govt.nz/nextapi/property?query=500%20Queen%20Street&pageSize=10
Authorization: Bearer <token from initialToken>
```

The address request uses no request body.

Verification result:

- Without `Authorization`: `401 Unauthorized`.
- With `Authorization: Bearer <initialToken>`: `200 OK`.
- Sanitized response body shape:

```json
{"items":[{"id":"12342478585","address":"500 Queen Street, Auckland Central"}]}
```

No cookie, captured bearer token, or private secret is required or stored.

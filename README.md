# Fizz-Buzz REST Server

A REST API written in Go (standard library only, zero dependencies) implementing
a generalized fizz-buzz: numbers from 1 to `limit`, where multiples of `int1`
are replaced by `str1`, multiples of `int2` by `str2`, and multiples of both by
`str1str2`. A statistics endpoint reports the most frequently used parameters.

## Requirements

- Go 1.24+

## Run

```bash
go run ./cmd/server
```

The server listens on port 8080 by default; override with the `PORT`
environment variable.

## API

### GET /fizzbuzz

| Parameter | Type   | Constraints  |
|-----------|--------|--------------|
| int1      | int    | >= 1         |
| int2      | int    | >= 1         |
| limit     | int    | 1 to 100000  |
| str1      | string | non-empty    |
| str2      | string | non-empty    |

```bash
curl "http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=15&str1=Fizz&str2=Buzz"
```

```json
{"result":["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]}
```

Invalid or missing parameters return HTTP 400 with every problem listed:

```json
{"errors":["int1 must be a positive integer","str1 must not be empty"]}
```

### GET /stats

Returns the most frequently requested parameter set and its hit count.
Only valid requests are counted.

```bash
curl "http://localhost:8080/stats"
```

```json
{"hits":2,"most_frequent_request":{"Int1":3,"Int2":5,"Limit":15,"Str1":"Fizz","Str2":"Buzz"}}
```

## Tests

```bash
go test ./...
```

Tests cover the fizz-buzz generation logic, parameter validation, the
statistics store, and the HTTP endpoints (status codes, JSON bodies,
error responses).

## Project structure

- `cmd/server` — entry point: configuration and HTTP server with timeouts.
- `internal/fizzbuzz` — domain logic: parameters, validation, sequence generation.
- `internal/stats` — in-memory, mutex-protected request counter.
- `internal/api` — HTTP layer: routing, query parsing, JSON responses.

## Design notes

- **Zero dependencies**: routing uses `net/http`'s method-aware `ServeMux`
  (Go 1.22+); no framework needed at this scale.
- **Statistics are in memory**: simple and dependency-free, at the cost of
  being reset on restart and being per-instance. To scale horizontally or
  survive restarts, `stats.Store` is the single component to swap for a
  shared store such as Redis.
- **Input hardening**: all parameters are validated and every problem is
  reported in one response; `limit` is capped at 100 000 to protect the
  server from abusive requests.
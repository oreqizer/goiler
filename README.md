# Goiler

The best boilerplate for a **Go** backend — running on `gqlgen` and `sqlboiler`!

## Features

* `go generate` — **GQLgen**, **SQLboiler**, **Dataloaden** and **RelayGen**
* Postgres
* Migrations
* Firebase auth
* Relay
* Query complexity
* Persisted queries
* CI
* Heroku
* Sentry

## Setup

You need a `.env` file with the following variables:

```
DATABASE_URL=...
SENTRY_URL=...
SENTRY_ENVIRONMENT=...
FIREBASE_KEY_FILE=...

# Optional
PORT=...
CORS=...
STAGING=...
```

## Development

Required files:
* `sqlboiler.yaml` with `psql` field filled in, see [docs](https://github.com/volatiletech/sqlboiler)
* `.env` as described above
* `firebase-private-key.json` if you don't have it in `FIREBASE_KEY_FILE` env var

Run `go generate` to regenerate files.

## License

MIT

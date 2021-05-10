# Horror

[![Go Report Card](https://goreportcard.com/badge/github.com/thinkofher/horror)](https://goreportcard.com/report/github.com/thinkofher/horror)
[![Go Reference](https://pkg.go.dev/badge/github.com/thinkofher/horror.svg)](https://pkg.go.dev/github.com/thinkofher/horror)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/thinkofher/horror)](https://github.com/thinkofher/horror/tags)

Package __horror__ implements interfaces for http error handling.

Horror was created to simplify error handling with standard go `http.Handler`
and give developers tools to encapsulate common error scenarios.

Thanks to `WithError` function and `Adapter` type you can use `horror.Handler`
with existing go code based on `http.Handler` interface.

If you're interested in using this module, please see examples that are
placed in [examples](./examples) directory.

See [github.com/thinkofher/horror/status](./status/doc.go) for convenient http errors that
satisfy Error interface from this module.

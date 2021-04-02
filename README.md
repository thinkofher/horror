# Horror

Package __horror__ implements interfaces for http error handling.

Horror was created to simplify error handling with standard go http.Handlers
and give developers tools to encapsulate common error scenarios.

Thanks to `WithError` function and `Adapter` type you can use horror.Handlers
with existing go code based on http.Handler interface.

If you're interested how to use this module, please see examples that are
placed in [examples](./examples) directory.

See [github.com/thinkofher/horror/status](./status/doc.go) for convenient http errors that
satisfy Error interface from this module.

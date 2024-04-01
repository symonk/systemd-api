## Systemd-api

> [!CAUTION]
> This project is completely over-engineered.  It was a learning exercise for new technologies!
> Cobra for a single command app, FX etc is completely excessive!

`systemd-api` is a simple service written in `go` for controlling systemd services
over the wire.  It uses:

 * `Golang` for the server backend.
 * `Gin` as the rest server.
 * `Cobra` for commandline management.
 * `Viper` for configuration management.
 * `fx` for dependency injection.
 * `testify` for unit testing.
 * `zap` for structured logging.
 * `prometheus` & `grafana` for metrics and tracking.

 -----

 ## Quick start

## Systemd-api

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

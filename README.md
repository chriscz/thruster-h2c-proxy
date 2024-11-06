# Thruster H2C Proxy

Fork of [Thruster](https://github.com/basecamp/thruster) but with H2C support.

## Installation

Thruster is distributed as a Ruby gem. Because Thruster is written in Go, we
provide several pre-built platform-specific binaries. Installing the gem will
automatically fetch the appropriate binary for your platform.

To install it, add it to your application's Gemfile:

```ruby
gem 'thruster'
```

Or install it globally:

```sh
$ gem install thruster
```


## Usage

To run your Puma application inside Thruster, prefix your usual command string
with `thrust`. For example:

```sh
$ thrust bin/rails server
```

Or with automatic TLS:

```sh
$ TLS_DOMAIN=myapp.example.com thrust bin/rails server
```


## Custom configuration

In most cases, Thruster should work out of the box with no additional
configuration. But if you need to customize its behavior, there are a few
environment variables that you can set.

| Variable Name         | Description                                             | Default Value |
|-----------------------|---------------------------------------------------------|---------------|
| `TLS_DOMAIN`          | Comma-separated list of domain names to use for TLS provisioning. If not set, TLS will be disabled. | None |
| `TARGET_PORT`         | The port that your Puma server should run on. Thruster will set `PORT` to this value when starting your server. | 3000 |
| `CACHE_SIZE`          | The size of the HTTP cache in bytes. | 64MB |
| `MAX_CACHE_ITEM_SIZE` | The maximum size of a single item in the HTTP cache in bytes. | 1MB |
| `X_SENDFILE_ENABLED`  | Whether to enable X-Sendfile support. Set to `0` or `false` to disable. | Enabled |
| `MAX_REQUEST_BODY`    | The maximum size of a request body in bytes. Requests larger than this size will be refused; `0` means no maximum size is enforced. | `0` |
| `STORAGE_PATH`        | The path to store Thruster's internal state. Provisioned TLS certificates will be stored here, so that they will not need to be requested every time your application is started. | `./storage/thruster` |
| `BAD_GATEWAY_PAGE`    | Path to an HTML file to serve when the backend server returns a 502 Bad Gateway error. If there is no file at the specific path, Thruster will serve an empty 502 response instead. Because Thruster boots very quickly, a custom page can be a useful way to show that your application is starting up. | `./public/502.html` |
| `HTTP_PORT`           | The port to listen on for HTTP traffic. | 80 |
| `HTTPS_PORT`          | The port to listen on for HTTPS traffic. | 443 |
| `HTTP_IDLE_TIMEOUT`   | The maximum time in seconds that a client can be idle before the connection is closed. | 60 |
| `HTTP_READ_TIMEOUT`   | The maximum time in seconds that a client can take to send the request headers and body. | 30 |
| `HTTP_WRITE_TIMEOUT`  | The maximum time in seconds during which the client must read the response. | 30 |
| `ACME_DIRECTORY`      | The URL of the ACME directory to use for TLS certificate provisioning. | `https://acme-v02.api.letsencrypt.org/directory` (Let's Encrypt production) |
| `EAB_KID`             | The EAB key identifier to use when provisioning TLS certificates, if required. | None |
| `EAB_HMAC_KEY`        | The Base64-encoded EAB HMAC key to use when provisioning TLS certificates, if required. | None |
| `FORWARD_HEADERS`     | Whether to forward X-Forwarded-* headers from the client. | Disabled when running with TLS; enabled otherwise |
| `DEBUG`               | Set to `1` or `true` to enable debug logging. | Disabled |

To prevent naming clashes with your application's own environment variables,
Thruster's environment variables can optionally be prefixed with `THRUSTER_`.
For example, `TLS_DOMAIN` can also be written as `THRUSTER_TLS_DOMAIN`. Whenever
a prefixed variable is set, it will take precedence over the unprefixed version.

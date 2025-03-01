---
slug: 'authentication'
---

## Authentication

PMM Server's user authentication is built on top of and is compatible with [Grafana authentication](https://grafana.com/docs/grafana/latest/auth/grafana/).

In this section we will talk about two main authentication mechanisms:

- Basic HTTP authentication
- Bearer Authentication (Authentication with an API key)

### Basic HTTP Authentication

Basic authentication is a very simple way to authenticate a user. An API request must contain a header field in the form of `Authorization: Basic <credentials>`, where `<credentials>` is the Base64 encoding of ID (most often username or login) and password joined by a single colon `:`.

```bash
echo -n admin:admin | base64
```

Let's assume the username is `admin` and the password is also `admin`. Then the API call to get the PMM Server version info would be as follows:

```bash
curl -X GET -H 'Authorization: Basic YWRtaW46YWRtaW4=' \
  -H 'Content-Type: application/json' https://127.0.0.1/v1/version
```

If you use `curl` to make API calls, a simple equivalent to the command above is:

```bash
curl -X GET -u admin:admin -H 'Content-Type: application/json' https://127.0.0.1/v1/version
```

### Bearer Authentication

Bearer authentication (also called token authentication) is an HTTP authentication scheme that involves security tokens called bearer tokens. The bearer token is a cryptic API key, which can be generated by the server admin from the Settings UI or via a respective API call (<a href="https://docs.percona.com/percona-monitoring-and-management/details/api.html?h=generate+api+keys#generate-api-keys">read more about how to generate an API key</a>). The client must send the API key in the `Authorization` header when making requests to protected resources:

```bash
curl -X GET -H 'Authorization: Bearer eyJrIjoiUXRkeDNMS1g1bFVyY0tUj1o0SmhBc3g4QUdTRVAwekoiLCJuIjoicG1tLXRlc3QiLCJpZCI6MX0=' \
  -H 'Content-Type: application/json' https://127.0.0.1/v1/version
```

You can use the API key in basic authentication as well:

```bash
curl -X GET -H 'Content-Type: application/json' \
  https://api_key:eyJrIjoiUXRkeDNMS1g1bFVyY0tUj1o0SmhBc3g4QUdTRVAwekoiLCJuIjoicG1tLXRlc3QiLCJpZCI6MX0=@127.0.0.1/v1/version
```

### Protecting Credentials

In the previous examples, the credentials can be gleaned from the shell history or processlist, which is undesirable.

#### Disable History

It is possible to hide from the shell history:

bash

```bash
set +o history
```

zsh

```zsh
SAVEHIST=0
```

#### Using --netrc

When using `curl` you also have the option of using `--netrc`. Here is an example `~/.netrc`:

```
machine 127.0.0.1
login admin
password admin
```

This can then be used as follows:

```bash
curl --netrc -X GET -H 'Content-Type: application/json' https://127.0.0.1/v1/version
```

Should you wish to use a different file then the `--netrc-file` option needs to be used. If we have the credentials stored in `~/.netrc-pmm` then the command would become:

```bash
curl --netrc --netrc-file ~/.netrc-pmm -X GET -H 'Content-Type: application/json' https://127.0.0.1/v1/version
```

You can use API keys in this way too, for example:

```
machine 127.0.0.1
login api_key
password eyJrIjoiUXRkeDNMS1g1bFVyY0tUj1o0SmhBc3g4QUdTRVAwekoiLCJuIjoicG1tLXRlc3QiLCJpZCI6MX0=
```

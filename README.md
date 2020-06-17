# catchcert

Connect to a server using TLS 🔒 and get PEM-encoded server certificate 📜.

It's a pretty trivial tool, yet it rocks! 🎸

## Usage examples

```
$ catchcert wikipedia.org
-----BEGIN CERTIFICATE-----
[...]

$ catchcert github.com > github.pem

# port is optional (defaults to 443)
> catchcert dc.contoso.com 636 > contoso_dc.crt
```

## Get it!
*catchcert* is written in [Go](https://golang.org/) and runs on Windows, GNU/Linux and macOS.

You can download the latest pre-compiled binaries from the [releases page](https://github.com/marcobellaccini/catchcert/releases).

## Want to build it?

```
$ git clone https://github.com/marcobellaccini/catchcert.git
$ cd catchcert
$ go build
```

## Why?
I wrote this tool because:

- I needed it
- I know I can [do the same with *openssl*](https://superuser.com/a/641396) but it is not so straightforward
- I [decided to try Go](https://twitter.com/lasagnasec/status/1268856625533779968) programming language
- I [wanted to try remote developement in a container](https://twitter.com/lasagnasec/status/1271399524179861504) with VS Code


## Contributors

[Marco Bellaccini](https://github.com/marcobellaccini) - creator, maintainer and cool guy

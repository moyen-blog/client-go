# Moyen CLI Client

A CLI client for the [Moyen publishing platform](https://moyen.blog). This is a bare-bones client accessible only via the command line.

## Installing the Client

There are two methods to install the client: via Go and directly using a binary.

### Installing Using Go

Having the Go programming language (Golang) installed is a prerequisite for this method of installation. Please refer to the [official documentation](https://golang.org) for information relating to installing Go.

Additionally, the [Go binary directory](https://golang.org/doc/gopath_code#GOPATH) (`$GOPATH/bin/`) should be in your PATH environment variable. On Linux and macOS, you can accomplish this by running `export PATH=$PATH:$(go env GOPATH)/bin`. To ensure your GOPATH remains in your PATH for new terminal sessions, add the previous command to your [shell startup file](https://www.gnu.org/software/bash/manual/html_node/Bash-Startup-Files.html).

Finally, to install the CLI client, run the following.

```
go install github.com/moyen-blog/client-go/moyen-cli@latest
```

### Installing Using the Binary

Download the appropriate binary for your system from the [releases page](https://github.com/moyen-blog/client-go/releases). Add the binary to a directory in your PATH environment variable e.g. `/usr/local/bin`.

## Using the Client

Please refer to the [Moyen Getting Started Guide](https://moyen.blog/Getting%20Started.md).

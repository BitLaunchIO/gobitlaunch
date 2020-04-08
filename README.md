# GoBitLaunch

GoBitLaunch is a Go client library for accessing the BitLaunch API.

You can view the client API docs here: [http://godoc.org/github.com/bitlaunchio/gobitlaunch](http://godoc.org/github.com/bitlaunchio/gobitlaunch)

You can view BitLaunch API docs here: [https://developers.bitlaunch.io/](https://developers.bitlaunch.io/)

## Install
```sh
go get -u github.com/bitlaunchio/gobitlaunch
```

## Usage

```go
import "github.com/bitlaunchio/gobitlaunch"
```

### Authentication

You must use your API Token to authenticate with BitLaunch API.
You can (re)generate your access token on the BitLaunch [API Page](https://app.bitlaunch.io/account/api).

You can then use your token to create a new client.

```go
client := gobitlaunch.NewClient(token)
```

## Documentation

For a comprehensive list of examples, check out the [API documentation](https://developers.bitlaunch.io/).

For details on all the functionality in this library, see the [GoDoc](http://godoc.org/github.com/bitlaunchio/gobitlaunch)

## Contributing

Please do!
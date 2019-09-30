# `kukushka`

[![GoDoc](https://img.shields.io/badge/api-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/recoilme/kukushka)

A Cuckoo Filter server with memcached protocol for Go.


# Getting Started

## Installing

To start using `kukushka`, install Go and run `go get`:

```sh
$ go get -u github.com/recoilme/kukushka
```

This will retrieve the library.

## Usage

```
telnet 127.0.0.1 11212
set some_userid-blockid 0 0 1
1
STORED
get some_userid-blockid
VALUE some_userid-blockid 0 1
1
END
get miss
VALUE miss 0 1
0
END
close
Connection closed by foreign host.
```


## Contact

Vadim Kulibaba [@recoilme](https://github.com/recoilme)

## License

`kukushka` source code is available under the MIT [License](/LICENSE).
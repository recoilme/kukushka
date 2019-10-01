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

## How it work

`Kukushka` implement SET,GET,GETS and CLOSE memcache comand

Value - ignored. Then you send SET yourkey - kukushka will add this key at cuckoo filter

If value in filter - GET return string "1". If not - "0".

Server started at 11212 on localhost (hardcoded)

## Usage (telnet example)

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

## map vs cuckoo vs bloom

```
1_000_000
map[uint64]uint64 memory usage: Alloc = 76 MiB
map[uint64]uint8 memory usage:  Alloc = 60 MiB 
cuckoo memory usage:            Alloc = 39 MiB
bloom memory usage:             Alloc = 31 MiB

100_000_000
map[uint64]uint64 - 4174 MiB   4.648s
bloom -   Alloc = 237 MiB  43.037s
cuckoo - Alloc = 289 MiB  29.405s

```

## Contact

Vadim Kulibaba [@recoilme](https://github.com/recoilme)

## License

`kukushka` source code is available under the MIT [License](/LICENSE).
# grpcrgen

[![Build Status](https://travis-ci.org/dictav/go-grpcrgen.svg?branch=master)](https://travis-ci.org/dictav/go-grpcrgen)
[![Build status](https://ci.appveyor.com/api/projects/status/oat9q5j05dqnrir3/branch/master?svg=true)](https://ci.appveyor.com/project/dictav/go-grpcrgen/branch/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/dictav/go-grpcrgen)](https://goreportcard.com/report/github.com/dictav/go-grpcrgen)

grpcrgen reads generated grpc codes created by protoc or flatc, and generates a reverse-proxy router.

It helps you to provide your gRPC APIs to Web browser.

![](http://g.gravizo.com/svg?
digraph G {
  rankdir="LR";
  node[shape=box];
  client[label="API Client"];
  proxy[label="Reverse Proxy"];
  server[label="gRPC Service"];
  fbs[label="service.fbs"];
  node[shape=oval];
  flatc;
  fbproxyc;
;
  subgraph flow {
    rank=same;
    // ???: back is required;
    client -> proxy[dir=back,label="POST"];
    proxy -> server[dir=back,label="gRPC"];
  }
;
  subgraph gen {
    fbs -> flatc;
    flatc -> client[label="generate stubs"];
    flatc -> server[label="generate stubs"];
    flatc -> fbproxyc[label="generate client"];
    fbproxyc -> proxy[label="generate router"];
  }
})

## Instration

```sh
go install github.com/dictav/go-grpcrgen/cmd/grpcrgen
```

## Usage

```sh
grpcrgen -o <output_dir> <flatc_generated_dir>
```

## Additional reading

- [PROTOCOL-WEB](https://github.com/grpc/grpc/blob/9d59861bc6dd969c9d42e0d8f83c4a77ba6ddc84/doc/PROTOCOL-WEB.md)

## Alternative

- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)

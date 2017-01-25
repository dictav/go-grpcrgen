# grpcrgen

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
go install github.com/dictav/go-grpcrgen
```

## Usage

```sh
grpcrgen --framework httprouter <generated_pkg_dir>
```

## Alternative

- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)

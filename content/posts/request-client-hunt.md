+++
Categories = ["development", "golang", "github"]
Tags = ["open-source"]
Description = "Extensible HTTP Request client wrapping lower & higher level API's"
date = "2015-04-04T06:36:43+03:00"
title = "Request client hunt"

+++

It's been a while I am on a hunt for elibile & configurable HTTP request client.
So I decided to write my own with blackjack & hookers, extensible, idiomatic,
exposing low/high level of API.

Package: https://github.com/ernestas-poskus/requestclient

Main features:

- Exposure of low/high API
- Connection and response timeouts
- Extensibility
- Connection Pooling
- Options struct for: TLS, Dialer, Transport & Client
- Retries
- Following redirects
- Delayed requests
- Multiple/Single request timeout

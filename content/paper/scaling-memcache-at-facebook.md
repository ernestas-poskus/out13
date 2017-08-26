+++
tags = ["paper","learning","research","review","scientific","operating systems"]
description = "Take away from scientific papers, discussion & summary"
link = "https://www.usenix.org/system/files/conference/nsdi13/nsdi13-final170_update.pdf"
name = "Scaling Memcache at Facebook"
author = "Rajesh Nishtala, Hans Fugal, Steven Grimm, Marc Kwiatkowski, Herman Lee, Harry C. Li, Ryan McElroy, Mike Paleczny, Daniel Peek, Paul Saab, David Stafford, Tony Tung, Venkateshwaran Venkataramani"
date = "2017-07-27T19:47:58+03:00"
title = "Scaling Memcache at Facebook"

+++

## Memcache at Facebook

Largest memcached installation in the world, processing over a billion requests per second and storing trillions of items.

Items are distributed across the memcached servers through consistent hashing.

`all` web servers communicate with every memcached server in a short period of time.
This all-to-all communication pattern can cause incast congestion or allow a single server to become the bottleneck for many web servers.
Reduce latency mainly by focusing on the memcache client, which runs on each web server.

Constructing a directed acyclic graph (DAG) representing the dependencies between data.
A web server uses this DAG to maximize the number of items that can be fetched concurrently.

Clients treat get errors as cache misses, but web servers will skip inserting entries
into memcached after querying for data to avoid putting additional load on
a possibly overloaded network or server.

Clients therefore use a sliding window mechanism to control the number of outstanding requests.
When the client receives a response, the next request can be sent.
Similar to TCP’s congestion control, the size of this sliding window grows
slowly upon a successful request and shrinks when a request goes unanswered.

Dedicate a small set of machines, named Gutter, to take over the responsibilities of a few
failed servers. Gutter accounts for approximately 1% of the memcached servers in a cluster.

When a memcached client receives no response to its get request,
the client assumes the server has failed and issues the request again to a special Gutter pool.

If this second request misses, the client will insert the appropriate key-value
pair into the Gutter machine after querying the database.
Entries in Gutter expire quickly to obviate Gutter invalidations.
Gutter limits the load on backend services at the cost of slightly stale data.

Deploy invalidation daemons (named mcsqueal) on every database.
Each daemon inspects the SQL statements that its database commits, extracts any deletes, and broad-
casts these deletes to the memcache deployment in every frontend cluster in that region.

### Performance optimizations

Began with a single-threaded memcached which used a fixed-size hash table. The first major optimizations
were to: (1) allow automatic expansion of the hash table to avoid look-up times drifting to O(n), make the
server multi-threaded using a global lock to protect multiple data structures, and (3) giving each thread
its own UDP port to reduce contention when sending replies and later spreading interrupt processing overhead.

#### Notes

> DAG - directed acyclic graph

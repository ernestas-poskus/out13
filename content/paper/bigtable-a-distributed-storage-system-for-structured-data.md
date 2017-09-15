+++
description = "Take away from scientific papers, discussion & summary"
tags = [
  "paper",
  "learning",
  "research",
  "review",
  "scientific",
]
author = "Fay Chang, Jeffrey Dean, Sanjay Ghemawat, Wilson C. Hsieh, Deborah A. Wallach, Mike Burrows, Tushar Chandra, Andrew Fikes, Robert E. Gruber"
name = "Bigtable: A Distributed Storage System for Structured Data"
link = "http://static.googleusercontent.com/media/research.google.com/en//archive/bigtable-osdi06.pdf"
date = "2016-11-03T19:54:45+02:00"
title = "Bigtable: A Distributed Storage System for Structured Data"

+++

## Bigtable

Bigtable is a distributed storage system for managing structured data that is
designed to scale to a very large size: petabytes of data across thousands of commodity servers.

Bigtable does not support a full relational data model; instead, it provides
clients with a simple data model that supports dynamic control over data layout
and format, and allows clients to reason about the locality properties of the data
represented in the underlying storage.

### Data model

A Bigtable is a sparse, distributed, persistent multidimensional sorted map.
The map is indexed by a row key, column key, and a timestamp; each value in the map
is an uninterpreted array of bytes.

> (row:string, column:string, time:int64) → string

Bigtable maintains data in lexicographic order by row key. The row range for a table is dynamically partitioned.
Each row range is called a tablet, which is the unit of distribution and load balancing.

### Architecture

File format to store data: SSTable provides a persistent, ordered immutable map from keys to values, where both keys and values are arbitrary byte strings.

First find the appropriate block by performing a binary search in the in-memory index, and then reading the appropriate block from disk.

Bigtable relies on a highly-available and persistent distributed lock service called Chubby.
Chubby service consists of five active replicas, one of which is elected to be the master and actively serve requests.

Chubby uses the Paxos algorithm to keep its replicas consistent in the face of failure

### Client

The client library caches tablet locations. 
If the client does not know the location of a tablet, or if it discovers that cached 
location information is incorrect, then it recursively moves up the tablet location hierarchy.

### Caching

To improve read performance, tablet servers use two levels of caching. 

 - Scan Cache is a higher-level cache that caches the key-value pairs returned by the SSTable interface to the tablet server code. 
 - Block Cache is a lower-level cache that caches SSTables blocks that were read from GFS.
 
Bloom filter allows us to ask whether an SSTable might contain any data for a specified row/column pair.

#### Notes

> GFS - Google File System

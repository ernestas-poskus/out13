+++
description = "Take away from scientific papers, discussion & summary"
title = "TAO: Facebook’s Distributed Data Store for the Social Graph"
date = "2016-12-15T19:36:32+02:00"
link = "https://cs.uwaterloo.ca/~brecht/courses/854-Emerging-2014/readings/data-store/tao-facebook-distributed-datastore-atc-2013.pdf"
name = "TAO: Facebook’s Distributed Data Store for the Social Graph"
author = "Nathan Bronson, Zach Amsden, George Cabrera, Prasad Chakka, Peter Dimov, Hui Ding, Jack Ferris, Anthony Giardullo, Sachin Kulkarni, Harry Li, Mark Marchukov, Dmitri Petrov, Lovro Puzar, Yee Jiun Song, Venkat Venkataramani"
tags = [
  "paper",
  "learning",
  "research",
  "review",
  "scientific",
]

+++

## Distributed data store for social graph

TAO is geographically distributed data store that provides efficient and timely
access to the social graph using a fixed set of queries.
Read optimized, persisted in MySQL.

Inefficient edge lists: A key-value cache is not a good
semantic fit for lists of edges; queries must always fetch
the entire edge list and changes to a single edge require
the entire list to be reloaded.

Distributed control logic: In a lookaside cache architecture 
the control logic is run on clients that don’t communicate 
with each other. This increases the number of
failure modes, and makes it difficult to avoid thundering herds.

Expensive read-after-write consistency: Facebook
uses asynchronous master/slave replication for MySQL,
which poses a problem for caches in data centers using a
replica. Writes are forwarded to the master, but some
time will elapse before they are reflected in the local
replica. By restricting the data model
to objects and associations we can update the replica’s
cache at write time, then use graph semantics to interpret
cache maintenance messages from concurrent updates.


### Data model and API

Facebook focuses on people, actions, and relationships.
We model these entities and connections as nodes and
edges in a graph. This representation is very flexible;
it directly models real-life objects, and can also be used
to store an application’s internal implementation-specific
data.

### Architecture

TAO needs to handle a far larger volume of data than can be stored on a 
single MySQL server, therefore data is divided into logical shards.

### MySQL mapping

Each shard is assigned to a logical MySQL database 
that has a table for objects and a table
for associations. All of the fields of an object are serialized into a 
single ‘data‘ column. This approach allows
us to store objects of different types within the same table, 
Objects that benefit from separate data management
polices are stored in separate custom tables.
Associations are stored similarly to objects, but to support 
range queries, their tables have an additional index
based on id1, atype, and time. To avoid potentially expensive 
SELECT COUNT queries, association counts
are stored in a separate table.

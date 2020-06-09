+++
tags = ["paper"]
description = "Take away from scientific papers, discussion & summary"
link = ""
title = "Aerospike: Architecture of a Real-Time Operational DBMS"
date = "2017-01-29T13:47:18+02:00"
name = "Aerospike: Architecture of a Real-Time Operational DBMS"
author = "V. Srinivasan, Sunil Sayyaparaju, Ashish Shinde, Brian Bulkowski, Andrew Gooding, Thomas Lopatic, Wei-Ling Chu, Rajkumar Iyer"

+++

## Aerospike architecture

Modeled on the classic shared-nothing database architecture

Objectives of the cluster management subsystem:

  - Arrive at a single consistent view of current cluster members across all nodes in the cluster.
  - Automatically detect new node arrival/departure and seamless cluster reconfiguration.
  - Detect network faults and be resilient to such network flakiness.
  - Minimize time to detect and adapt to cluster membership changes.

### Discovery

Node arrival or departure is detected via heartbeat messages
exchanged periodically between nodes.

### Surrogate heartbeats

In addition to regular heartbeat messages, nodes use other messages that are regularly exchanged
between nodes as an alternative secondary heartbeat mechanism.

### Node Health Score

Every node in the cluster evaluates the health score of each of its
neighboring nodes by computing the average message loss, which
is an estimate of how many incoming messages from that node are lost.

### Data Distribution

A record’s primary key is hashed into a 160-byte digest using the RipeMD160 algorithm.

Colocated indexes and data to avoid any cross-node traffic when running read operations or queries.

A partition assignment algorithm generates a replication list for every
partition. The replication list is a permutation of the cluster succession list.

Reads can also be uniformly spread across all the
replicas via a runtime configuration setting.

### Master Partition Without Data

An empty node newly added to a running cluster will be master
for a proportional fraction of the partitions and have no data for
those partitions.

### Migration Ordering

#### Smallest Partition First

Migration is coordinated in such a manner as to let nodes with the
fewest records in their partition versions start migration first. This
strategy quickly reduces the number of different copies of a
specific partition, and does this faster than any other strategy.

#### Hottest Partition First

At times, client accesses are skewed to a very small number of
keys from the key space. Therefore the latency on these accesses
could be improved quickly by migrating these hot partitions
before other partitions.

### Defragmentation

Aerospike uses a log-structured file system with a copy-on-write
mechanism. Hence, it needs to reclaim space by continuously
running a background defragmentation process. Each device
stores a MAP of block and information relating to the fill-factor of
each block. The fill-factor of the block is the block fraction
utilized by valid records. At boot time, this information is loaded
and kept up-to-date on every write. When the fill-factor of a block
falls below a certain threshold, the block becomes a candidate for
defragmentation and is then queued up for the defragmentation
process.

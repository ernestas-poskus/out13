+++
date = "2017-05-25T19:53:19+03:00"
title = "Spanner: Google’s Globally-Distributed Database"
link = "https://static.googleusercontent.com/media/research.google.com/en//archive/spanner-osdi2012.pdf"
name = "Spanner: Google’s Globally-Distributed Database"
author = "James C. Corbett, Jeffrey Dean, Michael Epstein, Andrew Fikes, Christopher Frost, JJ Furman, Sanjay Ghemawat, Andrey Gubarev, Christopher Heiser, Peter Hochschild, Wilson Hsieh, Sebastian Kanthak, Eugene Kogan, Hongyi Li, Alexander Lloyd, Sergey Melnik, David Mwaura, David Nagle, Sean Quinlan, Rajesh Rao, Lindsay Rolig, Yasushi Saito, Michal Szymaniak, Christopher Taylor, Ruth Wang, Dale Woodford"
tags = ["paper"]
description = "Take away from scientific papers, discussion & summary"

+++

## Spanner

Spanner is a scalable, globally-distributed database designed, built, and deployed at Google.

At the highest level of abstraction, it is a database that shards data across many sets of Paxos state machines.

Replication is used for global availability and geographic locality.

Spanner is designed to scale up to millions of machines across hundreds of datacenters and trillions of database rows.

Data is stored in schematized semi-relational tables; data is versioned,
and each version is automatically timestamped with its commit time; old versions of
data are subject to configurable garbage-collection policies; and applications can read data at old timestamps.

Major features:

  - externally consistent reads and writes
  - globally-consistent reads across the database at a timestamp

### Implementation

A Spanner deployment is called a universe. Given that Spanner manages data globally,
there will be only a handful of running universes.

Spanner is organized as a set of zones, where each zone is the rough analog of a deployment of Bigtable.

Zones are the unit of administrative deployment.

Zones are also the unit of physical isolation: there may be one or more zones in a datacenter, for example,
if different applications’ data must be partitioned across different sets of servers in the same datacenter.

Each spanserver is responsible for between 100 and 1000 instances of a data structure called a tablet.
A tablet is similar to Bigtable’s tablet abstraction, in that it implements a bag of the following mappings:

```
(key:string, timestamp:int64) → string
```

Unlike Bigtable, Spanner assigns timestamps to data, which is an important way in which Spanner is more
like a multi-version database than a key-value store.

Tablet state is stored in set of B-tree-like files and a write-ahead log,
all on a distributed file system called Colossus.

To support replication, each spanserver implements a
single Paxos state machine on top of each tablet.

### TrueTime

TrueTime explicitly represents time as a TTinterval, which is an interval with bounded time
uncertainty (unlike standard time interfaces that give clients no notion of uncertainty).

The time epoch is anal-
ogous to UNIX time with leap-second smearing. De-
fine the instantaneous error bound as ϵ, which is half of
the interval’s width, and the average error bound as ϵ.

TrueTime is implemented by a set of time master machines per datacenter and a timeslave daemon per machine.
The majority of masters have GPS receivers with dedicated antennas;
these masters are separated physically to reduce the effects of antenna failures, radio interference, and spoofing.

The remaining masters (which we refer to as Armageddon masters) are equipped with
atomic clocks. An atomic clock is not that expensive: the cost of an Armageddon master
is of the same order as that of a GPS master.

Between synchronizations, a daemon advertises a slowly increasing time uncertainty.

#### Notes

> Colossus - the successor of GFS

> GFS - Google File System

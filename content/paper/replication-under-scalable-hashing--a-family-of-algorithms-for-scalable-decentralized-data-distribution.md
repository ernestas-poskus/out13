+++
description = "Take away from scientific papers, discussion & summary"
tags = [
  "paper",
  "learning",
  "research",
  "review",
  "scientific",
  "operating systems",
]
author = "R. J. Honicky and Ethan L. Miller"
date = "2016-11-10T22:27:23+02:00"
title = "Replication Under Scalable Hashing: A Family of Algorithms for Scalable Decentralized Data Distribution"
name = "Replication Under Scalable Hashing: A Family of Algorithms for Scalable Decentralized Data Distribution"
link = "http://www.ssrc.ucsc.edu/media/papers/honicky-ipdps04.pdf"

+++

## Replication Under Scalable Hashing


Typical algorithms for decentralized data distribution work best in a system that is fully built before it first used; 
adding or removing components results in either extensive reorganization of data or load imbalance in the system.

RUSH variants also support weighting, allowing disks of different vintages to be added to a system.

RUSH variants is optimal or near-optimal reorganization. When new disks are added to the system, 
or old disks are retired, RUSH variants minimize the number of objects that need to 
be moved in order to bring the system back into balance.

RUSH variants can perform reorganization online without locking the filesystem for a long time to relocate data.

### Algorithm

Subcluster in a system managed by RUSH t must have at least as many disks as an object has replicas.

RUSH t is the best algorithms for distributing data over very large clusters of disks.

RUSH r may be the best option for systems which need to remove disks one at a time from the system.

RUSH p may be the best option for smaller systems where storage space is at a premium.

#### Notes

> RUSH t - RUSH tree

> RUSH r - RUSH support for removal

> PUSH p - RUSH placement using prime numbers

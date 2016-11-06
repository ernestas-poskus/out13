+++
name = "Dynamo: Amazon’s Highly Available Key-value Store"
link = "http://s3.amazonaws.com/AllThingsDistributed/sosp/amazon-dynamo-sosp2007.pdf"
date = "2016-11-06T12:32:44+02:00"
title = "Dynamo: Amazon’s Highly Available Key-value Store"
description = "Take away from scientific papers, discussion & summary"
tags = [
  "paper",
  "learning",
  "research",
  "review",
  "scientific",
  "operating systems",
  "performance",
]
author = "Giuseppe DeCandia, Deniz Hastorun, Madan Jampani, Gunavardhan Kakulapati, Avinash Lakshman, Alex Pilchin, Swaminathan Sivasubramanian, Peter Vosshall and Werner Vogels"

+++ 

## Dynamo 

Dynamo sacrifices Consistency for Availability under certain failure scenarios. 
It makes extensive use of object versioning and application-assisted conflict resolution in a manner that provides a novel interface for developers to use.

Gossip based distributed failure detection and membership protocol.

### Query Model

Read & Write operations to data item that is uniquely identified by a key. 
State is stored as blobs.
Targets application that store objects up to 1MB.

### ACID

Dynamo targets applications that operate with weaker consistency (the “C” in ACID) if this results in high availability.

No isolation guarantees. Permits only single key updates.

### Design

Incremental scalability: Dynamo should be able to scale out one storage host (henceforth, referred to as “node”) at a time, 
with minimal impact on both operators of the system and the system itself.

Symmetry: Every node in Dynamo should have the same set of responsibilities as its peers; there should be no distinguished node
or nodes that take special roles or extra set of responsibilities. In our experience, symmetry simplifies the process of system
provisioning and maintenance.

Decentralization: An extension of symmetry, the design should favor decentralized peer-to-peer techniques over centralized
control. In the past, centralized control has resulted in outages and the goal is to avoid it as much as possible. This leads to a simpler,
more scalable, and more available system.

Heterogeneity: The system needs to be able to exploit heterogeneity in the infrastructure it runs on. e.g. the work
distribution must be proportional to the capabilities of the individual servers. This is essential in adding new nodes with
higher capacity without having to upgrade all hosts at once.

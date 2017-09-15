+++
description = "Take away from scientific papers, discussion & summary"
tags = [
  "paper",
  "learning",
  "research",
  "review",
  "scientific",
]
author = "Edsger W. Dijkstra, Leslie Lamport, A.J. Martin, C.S. Scholten, E.F.M. Steffens"
name = "On-the-Fly Garbage Collection: An Exercise in Cooperation "
link = "http://research.microsoft.com/en-us/um/people/lamport/pubs/garbage.pdf"
date = "2016-08-25T19:13:56+03:00"
title = "On the fly garbage collection"

+++

In our abstract form of the problem, we consider a
directed graph of varying structure but with a fixed
number of nodes, in which each node has at most two
outgoing edges. More precisely, each node may have a
left-hand outgoing edge and may have a right-hand
outgoing edge, but either of them or both may be missing.
In this graph a fixed set of nodes exists, called "the
roots." A node is called "reachable" if it is reachable
from at least one root via a directed path along the edges.

The subgraph consists of all reachable nodes and their interconnections is
called 'the data structure'; nonreachable nodes that do not belong to the
data structure are called garbage.

Data structure can modified:
 - Redirecting an outgoing edge of a reachable node towards an already reachable one.
 - Redirecting an outgoing edge of a reachable node towards a not yet reachable one without outgoing edges.
 - Adding--where an outgoing edge was missing an edge pointing from a reachable node towards an already reachable one.
 - Adding--where an outgoing edge was missing an edge pointing from a reachable node towards a not yet reachable one without outgoing edges.
 - Removing an outgoing edge of a reachable node


Mutator: redirect an outgoing edge of reachable node towards an already reachable one.

Collector:
 - marking phase: mark all reachable nodes
 - appending phase: append all unmarked nodes to the free list and remove the markings from all marked nodes


#### Notes

> Free list - collection of nodes that have been identified as garbage.

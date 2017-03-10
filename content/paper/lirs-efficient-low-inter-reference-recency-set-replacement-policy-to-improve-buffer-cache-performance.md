+++
name = "LIRS: An Efficient Low Inter-reference Recency Set Replacement Policy to Improve Buffer Cache Performance"
author = "Song Jiang, Xiaodong Zhang"
tags = ["paper","learning","research","review","scientific","operating systems","performance"]
description = "Take away from scientific papers, discussion & summary"
date = "2017-03-09T19:34:52+02:00"
title = "LIRS: An Efficient Low Inter-reference Recency Set Replacement Policy to Improve Buffer Cache Performance"
link = "http://web.cse.ohio-state.edu/hpcs/WWW/HTML/publications/papers/TR-02-6.pdf"

+++

### LIRS

LRU replacement policy has been commonly used in the buffer cache management,
it is well known for its inability to cope with access patterns with weak locality.

LIRS effectively addresses the limits of LRU by using recency to evaluate Inter-Reference
Recency (IRR) for making a replacement decision.


#### LRU inefficiency

 - Under the LRU policy, a burst of references to infrequently used blocks such
   as “sequential scans” through a large file, may cause
   replacement of commonly referenced blocks in the cache.

 - For a cyclic (loop-like) pattern of accesses to a file that is only slightly
   larger than the cache size, LRU always mistakenly evicts the blocks that will
   be accessed soonest, because these blocks have not been accessed for the longest time.

The reason for LRU to behave poorly in these situations is
that LRU makes a bold assumption – a block that has not
been accessed the longest would wait for relatively longest
time to be accessed again.

#### Implementation

IRR as the recorded history information of each block, where IRR of a block
refers to the number of other blocks accessed between two consecutive references
to the block.

Specifically, the recency refers to the number of other blocks accessed from
last reference to the current time.

It is assumed that if the IRR of a block is large,
the next IRR of the block is likely to be large again.
Following this assumption, we select the blocks with large IRRs
for replacement, because these blocks are highly possible to
be evicted later by LRU before being referenced again under our assumption.

#### Notes

> LIRS - low inter reference set

> IRR - inter reference recency

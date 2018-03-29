+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper", "learning", "research", "review", "scientific", "rust"]
author = "Colin S. Gordon, Matthew J. Parkinson, Jared Parsons, Aleks Bromfield, Joe Duffy"
name = "Uniqueness and Reference Immutability for Safe Parallelism"
link = "https://www.microsoft.com/en-us/research/publication/uniqueness-and-reference-immutability-for-safe-parallelism/"
title = "Uniqueness and Reference Immutability for Safe Parallelism"
date = 2018-03-29T20:00:23+03:00
+++

Key challenge for concurrent programming is that side-effects (memory operations) in one thread can affect the behavior of another thread.

We wish to restrict, or tame, side-effects to make programs easier to maintain and understand.
To do so, we build on reference immutability.

To achieve this we give two novel typing
rules, which allow recovering isolated or immutable references from arbitrary code checked in environments containing only isolated or immutable inputs.
We provide two forms of parallelism:

 - Symmetric: Assuming that at most one thread may hold
   writable references to an object at a given point in time,
   then while all writable references in a context are temporarily forgotten (framed away, in the separation logic
   sense [28, 32]), it becomes safe to share all read-only or immutable references among multiple threads,
   in addition to partitioning externally-unique clusters between threads.

 - Asymmetric If all data accessible to a new thread is immutable
   or from externally-unique clusters which are
   made inaccessible to the spawning thread, then the new
   and old threads may run in parallel without interference.

### Reference Immutability, Uniqueness, and Parallelism

 - writable: An “ordinary” object reference, which allows mutation of its referent.

 - readable: A read-only reference, which allows no mutation of its referent. Furthermore, no heap traversal
   through a read-only reference produces a writable reference (writable references to the same objects may exist
   and be reachable elsewhere, just not through a readable reference). A readable reference may also refer to an immutable object.

 - immutable: A read-only reference which additionally notes
   that its referent can never be mutated through any reference.
   Immutable references may be aliased by read-only
   or immutable references, but no other kind of reference.
   All objects reachable from an immutable reference are
   also immutable.

 - isolated: An external reference to an externally-unique
   object cluster. External uniqueness naturally captures
   thread locality of data. An externally-unique aggregate
   is a cluster of objects that freely reference each other,
   but for which only one external reference into the aggregate exists.
   We define isolation slightly differently
   from most work on external uniqueness because we also
   have immutable objects: all paths to non-immutable objects reachable from the isolated reference pass through
   the isolated reference. We allow references out of the
   externally-unique aggregate to immutable data because it
   adds flexibility without compromising our uses for isolation:
   converting clusters to immutable, and supporting non-interference among threads. This
   change in definition does limit some traditional uses of
   externally-unique references that are not our focus, such
   as resource management tasks.

The most obvious use for reference immutability is to control where heap modification may occur in a program,
similar to the owner-as-modifier discipline in ownership and universe type systems.

### Safe Symmetric Parallelism

Fork-join concurrency is deterministic when neither forked
thread interferes with the other by writing to shared memory.

### Polymorphism

Any practical application of this sort of system naturally re-
quires support for polymorphism over type qualifiers.
Otherwise code must be duplicated, for example, for each
possible permission of a collection and each possible
permission for the objects contained within a collection.


#### Notes

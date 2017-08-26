+++
date = "2016-12-29T20:25:26+02:00"
author = "Graham Cormode, S. Muthukrishnan"
tags = [
  "paper",
  "learning",
  "review",
  "scientific",
]
description = "Take away from scientific papers, discussion & summary"
link = "http://dimacs.rutgers.edu/~graham/pubs/papers/cmsoft.pdf"
name = "Approximating Data with the Count-Min Data Structure"
title = "Approximating Data with the Count-Min Data Structure"

+++

## Count-Min Data Structure

Algorithmic problems such as tracking the contents of a set arise frequently in the course of building
systems. Given the variety of possible solutions, the choice of appropriate data structures for
such tasks is at the heart of building efficient and effective software.

The Count-Min sketch provides a different kind of solution to count tracking.
It allocates a fixed amount of space to store count information, which does not vary over time even
as more and more counts are updated.

### Implementation

With all data structures, it is important to understand the data organization
and algorithms for updating the structure, to make clear the relative merits of different choices of
structure for a given task. The Count-Min Sketch data structure primarily consists of a fixed array
of counters, of width w and depth d. The counters are initialized to all zeros. Each row of counters
is associated with a different hash function.

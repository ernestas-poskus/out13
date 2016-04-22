+++
author = "Gil Einziger, Roy Friedman"
date = "2016-04-22T21:26:15+03:00"
description = "Take away from scientific papers, discussion & summarize"
link = "http://arxiv.org/pdf/1512.00727.pdf"
name = "TinyLFU: A Highly Efficient Cache Admission Policy"
tags = ["paper", "learning", "research", "review", "scientific", "software engineering", "operating systems", "container", "performance"]
title = "Tiny LFU highly efficient cache admission policy"

+++

## Frequency based cache admission policy

> Approximate LFU structure called TinyLFU, which maintains an approximate representation of the access frequency of a large sample of recently accessed items.

TinyLFU is very compact and light-weight as it builds upon Bloom filter theory.

### Usage

The intuitive reason why caching works is that data accesses in many
domains of computer science exhibit a considerable degree of “locality”.

When a data item is accessed, if it already appears in the cache, we say that there is a cache hit; otherwise, it is a cache miss. The ratio between the number of cache hits and the total number of data accesses is known as the cache hit-ratio.

Admission policy - caching architecture in which an accessed item is only inserted into the cache if an admission policy decides that the cache hit ratio is likely to benefit from replacing it with the cache victim (as chosen by the cache’s replacement policy).

### Architecture

The cache eviction policy picks a cache victim, while TinyLFU decides if replacing the cache victim with the new item is expected to increase the hit-ratio.
To do so, TinyLFU maintains statistics of items frequency over a sizable recent history. Storing these statistics is considered prohibitively expensive for practical implementation and therefore TinyLFU approximates them in a highly efficient manner. To keep the history fresh an aging process is performed periodically or incrementally to halve all of the counters.


#### Notes

> Time locality - access pattern, and consequently the corresponding probability distribution, change over time

> WLFU - Window Least Frequently Used, access frequency for a window, needs to keep track order of requests. Samples of the request stream (called window).

> PLFU - Perfect LFU, popularity based has metadata with counters

> In-memory LFU, outperformed by WLFU at the cost of larger meta-data

> SLRU - Segmented Least Recently Used, policy captures recent popularity by distinguishing between tem-porally popular items that are accessed at least twice in a short window vs. items accessed only once during that period

> LRU-K - combination of LRU & LFU the last K occurrences of each element are remembered. Using this data, LRU-K statistically estimates the momentary frequency of items in order to keep the most frequent pages in memory.

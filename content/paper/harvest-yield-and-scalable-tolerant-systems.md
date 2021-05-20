+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper"]
author = "Armando Fox, Eric A. Brewer"
name = "Harvest, Yield, and Scalable Tolerant Systems"
link = "http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.24.3690&rep=rep1&type=pdf"
title = "Harvest Yield and Scalable Tolerant Systems"
date = 2021-05-07T18:40:52+03:00
+++

## Scalable tolerant systems

Degradation in terms of harvest and yield, and map it directly onto engineering mechanisms that enhance availability by improving fault isolation, and in some cases also simplify programming.

CA without P: Databases that provide distributed transactional semantics can only do so in the absence of a network partition separating server peers.

CP without A: In the event of a partition, further transactions to an ACID database may be blocked until the partition heals, to avoid the risk of introducing merge conflicts (and thus inconsistency).

AP without C: HTTP Web caching provides client-server partition resilience by replicating documents, but a client-server partition prevents verification of the freshness of an expired replica.
In general, any distributed database problem can be solved with either expiration-based caching to get AP, or replicas and majority voting to get PC (the minority is unavailable).

Many applications are best described in terms of reduced consistency or availability.

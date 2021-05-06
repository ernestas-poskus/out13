+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper"]
author = "Daniel Ford, François Labelle, Florentina I. Popovici, Murray Stokely, Van-Anh Truong, Luiz Barroso, Carrie Grimes, and Sean Quinlan"
name = "Globally distributed storage systems"
link = "https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/36737.pdf"
title = "Availability in Globally Distributed Storage Systems"
date = 2021-05-06T20:21:50+03:00
+++

Highly available cloud storage is often implemented with complex, multi-tiered distributed systems built on top of clusters of commodity servers and disk drives.
Sophisticated management, load balancing and recovery techniques are needed to achieve high performance and availability amidst an abundance of failure sources that include software, hardware, network connectivity, and power issues.

## Component availability

- Compare mean time to failure for system components at different granularities, including disks, machines and racks of machines.
- Classify the failure causes for storage nodes, their characteristics and contribution to overall unavailability.
- Apply a clustering heuristic for grouping failures which occurs almost simultaneously and show that a large fraction of failures happen in bursts.
- Quantify how likely a failure burst is associated with a given failure domain. We find that most large bursts of failures are associated with rack- or multi-rack level events.

## Data availability

 - Demonstrate the importance of modeling correlated failures when predicting availability, and show their impact under a variety of replication schemes and placement policies.
 - Formulate a Markov model for data availability, that can scale to arbitrary cell sizes, and captures the interaction of failures with replication policies and recovery times.
 - Introduce multi-cell replication schemes and compare the availability and bandwidth trade-offs against single-cell schemes.
 - Show the impact of hardware failure on our cells is significantly smaller than the impact of effectively tuning recovery and replication parameters.

#### Notes

> MTTF (mean time to failure) = uptime / number failures

> ARR - annual replacement rate

> AFR - annual failure rate

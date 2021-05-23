+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper"]
author = "Yogeshwer Sharma, Philippe Ajoux, Petchean Ang, David Callies, Abhishek Choudhary,Laurent Demailly, Thomas Fersch, Liat Atsmon Guz, Andrzej Kotulski, Sachin Kulkarni, Sanjeev Kumar, Harry Li, Jun Li, Evgeniy Makeev, and Kowshik Prakasam, Robbert van Renesse, Cornell University; Sabyasachi Roy, Pratyush Seth, Yee Jiun Song, Benjamin Wester, Kaushik Veeraraghavan, and Peter Xie"
name = "Wormhole"
link = "https://research.fb.com/publications/wormhole-reliable-pub-sub-to-support-geo-replicated-internet-services/"
title = "Wormhole: Reliable Pub-Sub to Support Geo-replicated Internet Services"
date = 2021-05-23T07:19:49+03:00
+++

## Wormhole

Wormhole is a publish-subscribe (pub-sub) system developed for use within Facebook’s geographically replicated datacenters.
It is used to reliably replicate changes among several Facebook services including TAO, Graph Search and Memcache.

Facebook production deployment of Wormhole transfers over 35 GBytes/sec in steady state (50 millions messages/sec or 5 trillion messages/day) across all deployments with bursts up to 200 GBytes/sec during failure recovery.

Wormhole, a pub-sub system that identifies new writes and publishes updates to all interested applications.
Publishers directly read the transaction logs maintained by the data storage systems to learn of new writes committed by producers.

Wormhole provides multiple-copy reliable delivery where it allows applications to configure a primary source and many secondary sources they can receive updates from.

1. Different consumption speeds: Applications consume updates at different speeds.
   A slow application that synchronously processes updates should not hold up data delivery to a fast one.

2. At least once delivery: All updates are delivered at least once.
   This ensures that applications can trust that they have received all updates that they are interested in.

3. In-order delivery of new updates: When an update is received, the application should be confident that all updates prior to the received one have also been received earlier.

4. Fault tolerance: The system must be resilient to frequent hardware and software failure both on the datastore as well as the application end.

Subscribers receive the stream of updates for every shard, which we call a flow.
Publishers periodically track data-markers per flow after the subscribers acknowledge that they have processed the updates up to new datamarker.

Subscribers are assumed to be stateless. In particular, they don’t need to keep track of the state of the flow.

Wormhole supports two types of datasets: single-copy datasets and multiple-copy datasets. The latter indicates a georeplicated dataset.
Accordingly, Wormhole supports both single-copy reliable delivery (SCRD) and multiple-copy reliable delivery (MCRD). For SCRD, Wormhole guarantees that when an application is subscribed to the single copy of a dataset,
its subscribers receive at least once all updates contained in that single copy of the dataset.
The updates for any shard are delivered to the application in order that they were stored in the transaction logs: delivery of an update means all prior updates for that shard have already been delivered.

#### Notes

> Dataset: is a collection of related data, for example, user generated data in Facebook.

+++
author = "Jim Gray"
date = "2016-08-12T16:57:55+03:00"
description = "Take away from scientific papers, discussion & summary"
link = "https://arxiv.org/pdf/cs/0701158.pdf"
name = "Queues Are Databases"
tags = ["paper"]
title = "Queues Are Databases"

+++

## Queued transaction processing over pure client-server transaction processing.

Queued systems are build on top of direct systems.

TP systems offer both queued and direct transaction processing. They offer both client-server and P2P direct processing.

Queue manager is best built as a naive resource manager atop an object-relational database system.
That system must have good concurrency control, recovery, triggers, security, operations interfaces, and utilities.

Queues pose difficult problems when implemented atop a database:

 - Performance: An enqueue transaction is an insert followed by a commit. This places
extreme performance demands on the concurrency control and recovery components
of a database -- it exposes hotspots and high-overhead code.

 - Concurrency control: The dequeue transaction typically involves deleting a record from
the queue, processing the request, enqueuing results in other queues, and then
committing. Serializable isolation requires that there can be at most one dequeue
executing at a time against each queue. This suggests that queues need lower, indeed specialized, isolation levels.

 - Read past: locks allow a program to skip over dirty (uncommitted records) to find the
first committed record. This is what a dequeue() operation wants.

 - Read through: locks allow a program to examine records that have not yet been
committed. This is useful in polling the status of a queued request that is currently
being processed.

 - Notify: allow a program to wait for a state change in a lock. This allows a
dequeue() operation to wait for one or more queues to become non-empty.


#### Notes

> MOM - message oriented middleware

> TP - transaction processing

> P2P - peer to peer

> ORB - object request broker

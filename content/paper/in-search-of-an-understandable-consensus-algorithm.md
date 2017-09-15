+++
date = "2017-04-20T19:13:57+03:00"
title = "In Search of an Understandable Consensus Algorithm"
link = "https://raft.github.io/raft.pdf"
name = "In Search of an Understandable Consensus Algorithm"
author = "Diego Ongaro and John Ousterhout"
tags = ["paper","learning","research","review","scientific"]
description = "Take away from scientific papers, discussion & summary"

+++


## Raft

Consensus algorithm for managing a replicated log.

Raft separates the key elements of consensus, such as leader election, log replication, and safety, and it enforces
a stronger degree of coherency to reduce the number of states that must be considered.

Paxos first defines a protocol capable of reaching agreement on a single decision,
such as a single replicated log entry.

Raft implements consensus by first electing a distinguished leader,
then giving the leader complete responsibility for managing the replicated log.
The leader accepts log entries from clients, replicates them on other servers,
and tells servers when it is safe to apply log entries to
their state machines.

A leader can fail or become disconnected from the other servers, in which case
a new leader is elected.

 - Election Safety: at most one leader can be elected in a given term.
 - Leader Append-Only: a leader never overwrites or deletes entries in its log; it only appends new entries.
 - Log Matching: if two logs contain an entry with the same index and term, then the logs are identical in all entries up through the given index.
 - Leader Completeness: if a log entry is committed in a given term, then that entry will be present in the logs of the leaders for all higher-numbered terms.
 - State Machine Safety: if a server has applied a log entry at a given index to its state machine, no other server will ever apply a different log entry for the same index.

### Basics

A Raft cluster contains several servers; five is a typical number, which allows the system to tolerate two failures.
At any given time each server is in one of three states: leader, follower, or candidate.

To prevent split votes in the first place, election timeouts are chosen randomly from a fixed interval (e.g., 150–300ms).

Raft guarantees that committed entries are durable and will eventually be executed by all of the available state machines.

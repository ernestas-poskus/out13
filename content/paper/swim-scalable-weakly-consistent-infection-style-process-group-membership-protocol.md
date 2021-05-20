+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper"]
author = "Abhinandan Das, Indranil Gupta, Ashish Motivala"
name = "Scalable Weakly-consistent Infection-style Process Group Membership Protocol"
link = "https://www.cs.cornell.edu/projects/Quicksilver/public_pdfs/SWIM.pdf"
title = "Swim Scalable Weakly Consistent Infection Style Process Group Membership Protocol"
date = 2021-05-14T20:01:12+03:00
+++

## SWIM

Distributed peer-to-peer applications require weakly-consistent knowledge of process group membership information at all participating processes.

SWIM separates the failure detection and membership update dissemination functionalities of the membership protocol.

Swim focus on a weaker variant of group membership, where membership lists at different members need not be consistent across the group at the same (causal) point in time.

The design of a distributed membership algorithm has traditionally been approached through the technique of heart-beating.

Popular class of all-to-all heart-beating protocols arises from the implicit decision therein to fuse the two principal functions of the membership problem specification:

1) Membership update Dissemination: propagating membership updates arising from processes joining, leaving or failing
2) Failure detection: detecting failures of existing members.

SWIM, provides a membership substrate that:

  (1) imposes a constant message load per group member;
  (2) detects a process failure in an (expected) constant time at some non-faulty process in the group;
  (3) provides a deterministic bound (as a function of group size) on the local time that a non-faulty process takes to detect failure of another process;
  (4) propagates membership updates, including information about failures, in infection-style (also gossip-style or epidemic-style);
      the dissemination latency in the group grows slowly (logarithmically) with the number of members;
  (5) provides a mechanism to reduce the rate of false positives by “suspecting” a process before “declaring” it as failed within the group.

#### Notes

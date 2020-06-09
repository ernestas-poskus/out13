+++
link = "https://www.cs.cornell.edu/home/rvr/papers/flowgossip.pdf"
name = "Efficient Reconciliation and Flow Control for Anti-Entropy Protocols"
author = "Robbert van Renesse, Dan Dumitriu, Valient Gough, Chris Thomas"
title = "Efficient Reconciliation and Flow Control for Anti-Entropy Protocols"
date = "2016-12-01T16:05:39+02:00"
tags = ["paper"]
description = "Take away from scientific papers, discussion & summary"

+++

## Flow Gossip

Anti-entropy, or gossip, is an attractive way of replicating state that does not have strong consistency requirements.
With few limitations, updates spread in expected time that grows logarithmic in the number of participating hosts, even in the face of host failures and message loss.
The behavior of update propagation is easily modeled with well-known epidemic analysis techniques.

### Gossip basics

There are two classes of gossip: anti-entropy and rumor mongering protocols.
Anti-entropy protocols gossip information until it is made obsolete by newer information,
and are useful for reliably sharing information among a group of participants.
Rumor-mongering has participants gossip information for some amount of time chosen sufficiently
high so that with high likelihood all participants receive the information.

3 Gossip styles:

 - push: push everything and apply everything
 - pull: sends its state with values removed, leaving only keys and version numbers, then returns only necessary updates
 - push-pull: like pull but sends a list of participant-key pairs for which if has outdated entries (most efficient)

### Precise reconciliation

The two participants in a gossip exchange send exactly those mappings that are more recent
than those of the peer. Thus, if the participants are p and q, p sends to q the set of deltas.

### Scuttlebutt reconciliation

A gossiper never transmits updates that were already known at the receiver.
If gossip messages were unlimited in size, then the sets contains the exact differences, just like with precise reconciliation.
If a set does not fit in the gossip message, then it is not allowed to use an arbitrary subsetas in precise reconciliation.

### Flow control

The objective of a flow control mechanism for gossip is to determine, adaptively,
the maximum rate at which a participant can submit updates without creating a backlog of updates.
A flow control mechanism should be fair, and under high load afford each participant that wants to submit updates the same update rate.
As there is no global oversight, the flow control mechanism has to be decentralized,
where the desired behavior emerges from participants responding to local events.

### Local adaptation

For local adaptation, we use an approach inspired by TCP flow control.
In TCP, the send window adapts according to a strategy called Additive Increase Multiplicative decrease.

In this strategy, window size grows linearly with each successful transmission,
but is decreased by a certain factor whenever overflow occurs.
In the case of TCP, the overflow signal is the absence of an acknowledgment.

#### Notes

> Anti-entropy - gossip information until it is made obsolete.

> Rumor-mongering - gossip information for some of high amount of time with high likelihood all participants received the information.

> AIMD - additive increase multiplicative decrease

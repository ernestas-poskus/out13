+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper"]
author = "Satoshi Nakamoto"
name = "Bitcoin: A Peer-to-Peer Electronic Cash System"
link = "https://bitcoin.org/bitcoin.pdf"
title = "Bitcoin a Peer to Peer Electronic Cash System"
date = 2017-10-19T20:23:22+03:00
+++

## P2P Electronic Cash System

A purely peer-to-peer version of electronic cash would allow online
payments to be sent directly from one party to another without going through a
financial institution.

Prevent double-spending problem using a peer-to-peer network.

As long as a majority of CPU power is controlled by nodes that are not cooperating to
attack the network, they'll generate the longest chain and outpace attackers.

The system is secure as long as honest nodes collectively control more CPU power than any
cooperating group of attacker nodes.

### Transactions

We define an electronic coin as a chain of digital signatures. Each owner transfers the coin to the
next by digitally signing a hash of the previous transaction and the public key of the next owner
and adding these to the end of the coin.

### Timestamp server

The timestamp proves that the data must have existed at the
time, obviously, in order to get into the hash. Each timestamp includes the previous timestamp in
its hash, forming a chain, with each additional timestamp reinforcing the ones before it.

### Network

1) New transactions are broadcast to all nodes.

2) Each node collects new transactions into a block.

3) Each node works on finding a difficult proof-of-work for its block.

4) When a node finds a proof-of-work, it broadcasts the block to all nodes.

5) Nodes accept the block only if all transactions in it are valid and not already spent.

6) Nodes express their acceptance of the block by working on creating the next block in the chain, using the hash of the accepted block as the previous hash.

Nodes always consider the longest chain to be the correct one and will keep working on extending it.

### Privacy

Additional firewall, a new key pair should be used for each transaction to keep them
from being linked to a common owner. Some linking is still unavoidable with multi-input transactions,
which necessarily reveal that their inputs were owned by the same owner.

#### Notes

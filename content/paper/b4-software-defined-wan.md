+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper", "learning", "research", "review", "scientific"]
author = "Sushant Jain, Alok Kumar, Subhasree Mandal, Joon Ong, Leon Poutievski, Arjun Singh, Subbaiah Venkata, Jim Wanderer, Junlan Zhou, Min Zhu, Jonathan Zolla, Urs Hölzle, Stephen Stuart, Amin Vahdat"
name = "B4: Experience with a Globally-Deployed Software Defined WAN"
link = "http://cseweb.ucsd.edu/~vahdat/papers/b4-sigcomm13.pdf"
title = "B4: Experience with a Globally-Deployed Software Defined WAN"
date = 2018-01-07T11:32:25+02:00
+++

## Software defined WAN

B4, a private WAN connecting Google’s data centers across the planet.

i) Massive bandwidth requirements deployed to a modest number of sites.

ii) Elastic traffic demand that seeks to maximize average bandwidth.

iii) Full control over the edge servers and network, which enables rate limiting and demand measurement at the edge.

### Design

Within each B4 site, the switch hardware layer primarily forwards traffic and does not run complex control
software, and the site controller layer consists of Network Control Servers (NCS) hosting both OpenFlow controllers (OFC)
and Network Control Applications (NCAs).

 - B4 routers built from merchant switch silicon
 - Drive links to 100% utilization
 - Centralized traffic engineering
 - Separate hardware from software

### Centralized TE Architecture

 - The Network Topology graph represents sites as vertices and site to site connectivity as edges. The SDN Gateway consolidates topology events from multiple sites and individual switches to TE. TE aggregates trunks to compute site-site edges. This abstraction significantly reduces the size of the graph input to the TE Optimization Algorithm.
 - Flow Group (FG): For scalability, TE cannot operate at the granularity of individual applications. Therefore, we aggregate applications to a Flow Group defined as {source site, dest site, QoS} tuple.
 - A Tunnel (T) represents a site-level path in the network. B4 implements tunnels using IP in IP encapsulation.
 - A Tunnel Group (TG) maps FGs to a set of tunnels and corresponding weights.
   The weight specifies the fraction of FG traffic to be forwarded along each tunnel.

#### Notes

> TE - traffic engineering

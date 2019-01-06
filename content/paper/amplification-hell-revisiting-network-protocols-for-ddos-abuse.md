+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper", "learning", "research", "review", "scientific"]
author = "Christian Rossow"
name = "Amplification Hell: Revisiting Network Protocols for DDoS Abuse"
link = "https://christian-rossow.de/publications/amplification-ndss2014.pdf"
title = "Amplification Hell: Revisiting Network Protocols for DDoS Abuse"
date = 2019-01-06T11:35:06+02:00
+++

## Amplification attacks

Adversaries send requests to public servers (e.g., open
recursive DNS resolvers) and spoof the IP address of a victim.
These servers, in turn, flood the victim with valid responses
and – unknowingly – exhaust its bandwidth.

Attackers can abuse these protocols to multiply their
attack bandwidth by factors from 3.8 (BitTorrent, NetBios) up to 4670 (NTP).

Exclude all TCP-based protocols from our analysis, as IP address spoofing is
restricted to the start of the TCP handshake. Although the TCP
handshake fulfils the reflection criterion, it does not allow for
easy amplification (since a TCP ACK is not larger than a TCP
SYN packet). However, we show that 14 popular UDP-based
network protocols are suitable candidates for DRDoS attacks.

SNMP v2, NTP, DNS, NetBios, SSDP, CharGen, QOTD, BitTorrent, Kad, Quake 3,
Steam, ZAv2, Sality, Gameover

An attacker A aims to consume all available bandwidth of a victim V . Reflective means
that A does not directly send traffic to V , but instead uses
systems that reflect the attack traffic to V (so called amplifiers).
Distributed accounts for the fact that A abuses thousands of
amplifiers and V thus faces thousands of attack sources.

BAF = len(UDP payload) amplifier to victim / len(UDP payload) attacker to amplifier

PAF = number or packets amplifier to victim / number of packets attacker to amplifier

#### Notes

> DRDoS - distributed reflective denial-of-service

> BAF - bandwidth amplification factor

> PAF - packet amplification factor

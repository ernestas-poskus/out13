+++
author = "Nandita Dukkipati, Tiziana Refice, Yuchung Cheng, Jerry Chu, Tom Herbert, Amit Agarwal, Arvind Jain and Natalia Sutin"
date = "2016-08-04T22:02:54+03:00"
description = "Take away from scientific papers, discussion & summary"
link = "http://research.google.com/pubs/pub36640.html"
name = "An Argument for Increasing TCP’s Initial Congestion Window"
tags = ["paper"]
title = "An Argument for Increasing TCP’s Initial Congestion Window"

+++

## TCP congestion window

> TCP flows start with initial congestion window of 4 segments (4KB of data).

Window if critical for how quickly flows can finish.

Increase in 15KB congestion window improves average HTTP latency by 10%, mostly benefits RTT and BDP.

Slow start increases congestion window by the number of data segments acknowledged for each received ACK.

TCP latency is dominated by the number of round-trip times in slow-start phase.

Increasing init_cwnd enables transfers to finish in fewer RTT.

#### Notes

> BDP - bandwidth delay product.

> RTT - round trip delay time.

> Wep page average size - 384KB.

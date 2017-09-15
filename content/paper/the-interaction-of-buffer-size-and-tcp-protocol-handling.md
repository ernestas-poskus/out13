+++
author = "Miercom"
name = "Speeding Applications in Data Center Networks The Interaction of Buffer Size and TCP Protocol Handling and its Impact on Data-Mining and Large Enterprise IT Traffic Flows"
link = "http://miercom.com/pdf/reports/20160210.pdf"
date = "2016-11-17T19:23:07+02:00"
title = "The Interaction of Buffer Size and TCP Protocol Handling and its Impact"
description = "Take away from scientific papers, discussion & summary"
tags = [
  "paper",
  "learning",
  "research",
  "review",
  "scientific",
]

+++

### Abstract

Miercom was engaged by Cisco Systems to conduct independent testing of two vendors’ top of the line, 
data-center switch-routers, including the Cisco Nexus 92160YC-X and Nexus 9272Q switches and the Arista 7280SE-72 switch.

#### TCP Congestion Control versus System Buffer Management

TCP congestion control. The Transmission Control Protocol (TCP) is the Layer-4 control
protocol (atop IP at Layer 3) that ensures a block of data that’s sent is received intact.
Invented 35 years ago, TCP handles how blocks of data are broken up, sequenced, sent,
reconstructed and verified at the recipient’s end. The congestion-control mechanism
was added to TCP in 1988 to avoid network congestion meltdown. It makes sure data
transfers are accelerated or slowed down, exploiting the bandwidth that’s available,
depending on network conditions.

System buffer management. Every network device that transports data has buffers,
usually statically allocated on a per-port basis or dynamically shared by multiple ports, so
that periodic data bursts can be accommodated without having to drop packets.
Network systems such as switch-routers are architected differently, however, and can
vary significantly in the size of their buffers and how they manage different traffic flows.

#### Deep buffer vs Intelligent buffer

A common practice is to put in as much buffer as possible. However, since the
buffer space is a common resource shared by the inevitable mixture of elephant and mice flows,
how to use this shared resource can significantly impact applications’ performance.

The deeper the buffer, the longer the queue and the longer the latency. So more buffer does
not necessarily guarantee better small-flow performance, it often leads to longer queuing delay
and hence longer flow completion time.

Therefore, no one benefits from simple deep buffering: mice flows aren’t guaranteed buffer
resources and can suffer from long queuing delays and bandwidth hungry elephant flows suffer
because large buffers do not create more link bandwidth.

#### Conclusion

Since mice flows are often mission critical (including, for example, control and alarm messages,
Hadoop application communications, etc.), giving these flows a priority buffer pathway enables
them to complete faster and their applications to perform better overall. The above test results
show that expediting mice flows and regulating the elephant flows early under the intelligent
buffer architecture on the Cisco Nexus 92160YC-X and 9272Q switches can bring orders of
magnitude better performance for mission critical flows without causing elephant flows to slow
down.

Intelligent buffering allows the elephant and mice flows to share network buffers gracefully:
there is enough buffer space for the bursts of mice flows while the elephant flows are properly
regulated to fully utilize the link capacity. Simple, deep buffering can lead to collateral damage
in the form of longer queuing latency, and hence longer flow completion time for all flow types.


#### Notes

> Elephant - big flows

> Mice - small flows

> FCT - flow completion time

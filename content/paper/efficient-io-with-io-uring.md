+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper", "kernel"]
author = "Kernel developers"
name = "Efficient IO with io_uring"
link = "https://kernel.dk/io_uring.pdf"
title = "IO with io_uring"
date = 2020-06-09T20:55:21+03:00
+++

## IO history

There are many ways to do file based IO in Linux. The oldest and most basic are the read(2) and write(2) system
calls. These were later augmented with pread(2) and pwrite(2) versions which allow passing in of an offset, and later
still we got preadv(2) and pwritev(2) which are vector-based versions of the former. Because that still wasn't quite
enough, Linux also has preadv2(2) and pwritev2(2) system calls, which further extend the API to allow modifier flags.

### Linux native async IO interface: aio

Supports only O_DIRECT (unbuffered access).
Blocks if metadata is required to perform IO.
Blocks if requests slots are unavailable.
104 bytes for each memory copy.

### io_uring

Kernel produces completion events and the
application consumes them. Hence, we need a pair of rings to provide an effective communication channel between
an application and the kernel. That pair of rings is at the core of the new interface, io_uring. They are suitably named
submission queue (SQ), and completion queue (CQ), and form the foundation of the new interface.

### Polled IO

Applications chasing the very lowest of latencies, io_uring offers support for polled IO for files. In this context,
polling refers to performing IO without relying on hardware interrupts to signal a completion event. When IO is polled,
the application will repeatedly ask the hardware driver for status on a submitted IO request.
This is different than non-polled IO, where an application would typically go to sleep waiting for the hardware interrupt as its wakeup source. For
very low latency devices, polling can significantly increase the performance.
The same is true for very high IOPS
applications as well, where high interrupt rates makes a non-polled load have a much higher overhead. The boundary
numbers for when polling makes sense, either in terms of latency or overall IOPS rates, vary depending on the
application, IO device(s), and capability of the machine.

#### Notes

> SQ - submission queue

> CQ - completion queue

> CQE - completion queue entry

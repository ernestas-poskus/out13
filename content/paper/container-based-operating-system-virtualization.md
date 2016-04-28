+++
author = "Stephen Soltesz, Herbert Pötzl, Marc E. Fiuczynski, Andy Bavier, Larry Peterson"
date = "2016-04-19T19:30:48+03:00"
description = "Take away from scientific papers, discussion & summary"
link = "http://dl.acm.org/citation.cfm?id=1273025"
name = "Container-based Operating System Virtualization: A Scalable, High-performance Alternative to Hypervisors"
tags = ["paper", "learning", "research", "review", "scientific", "software engineering", "operating systems", "container", "performance"]
title = "Container based operating system virtualization"

+++

## Alternative to hypervisors.

> Workload requirements for a given system will direct users to the point in the design space that
> requires the least trade-off.

### Sharing over isolation? 

Hypervisors often deployed to let a single machine host multiple, unrelated
applications, which may run on behalf of independent organizations, as is common when a data center
consolidates multiple physical servers. Hypervisors favor full isolation over sharing.
However, when each virtual machine is running the same kernel and similar operating system 
distributions, the degree of isolation offered by hypervisors comes at the cost of efficiency 
relative to running all applications on a single kernel.

### Usage

Software configuration problems incompatibilities between specific OS distributions.

Resource isolation corresponds to the ability to account for and enforce the resource consumption of one VM such that guarantees and fair shares are preserved for other VM's.

Many hybrid approaches are also possible: for instance, a system may enforce fair sharing of resources between classes of VMs, which lets one overbook available resources while preventing starvation in overload scenarios. 

The key point is that both hypervisors and COS's incorporate sophisticated resource schedulers to avoid or minimize crosstalk.

### Security isolation

Configuration independence - cannot conflict with other VM's
Safety - global namespace shared

### Fair share and Reservations

Vserver implements CPU isolation by overlaying a token TBF on top of standard O(1) Linux CPU scheduler.

For memory storage one can specify the following limits: 
 * a) the maximum resident set size (RSS)
 * b) number of anonymous memory pages have (ANON)
 * c) number of pages that may be pinned into memory using mlock() and mlockall() that processes may have within a VM (MEMLOCK).
 
### Conclusion

Xen is able to support multiple kernels while by design VServer cannot. 
Xen also has greater support for virtualizing the network stack and allows for the possibility of VM migration, a feature that is possible for a COS design, but not yet available in VServer. VServer, in turn, maintains a small kernel footprint and performs equally with native Linux kernels in most cases.

#### Notes

> Undesired interactions between VMs are sometimes called cross-talk.

> COS - Container based Operating System

> TBF - token bucker filter

> HTB - Hierarchical Token Bucket

> RSS - maximum resident set size

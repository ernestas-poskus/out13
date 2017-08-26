+++
author = "Benjamin Hindman, Andy Konwinski, Matei Zaharia, Ali Ghodsi, Anthony D. Joseph, Randy Katz, Scott Shenker, Ion Stoica"
date = "2016-04-28T19:50:29+03:00"
description = "Take away from scientific papers, discussion & summary"
link = "https://www.cs.berkeley.edu/~alig/papers/mesos.pdf"
name = "Mesos: A Platform for Fine-Grained Resource Sharing in the Data Center"
tags = ["paper", "learning", "research", "review", "scientific"]
title = "Mesos: A Platform for Fine-Grained Resource Sharing in the Data Center"

+++

## Platform for resource sharing

> Sharing improves cluster utilization and avoids per-framework data repli-cation
> Organizations will want to run multiple frameworks in the same cluster, picking the best one for each application. 
> Sharing a cluster between frameworks improves utilization and allows applications to share access to large datasets that may be too costly to replicate

### Architecture

Mesos decides how many resources to offer each framework, based on an organizational policy such as fair sharing, while frameworks decide which resources to accept and which tasks to run on them.

> Design philosophy - define a minimal interface that enables efficient resource sharing across frameworks, and otherwise push control of task scheduling and execution to the frameworks.

The master decides how many resources to offer to each framework according to a given organizational policy, such as fair sharing, or strict priority.

A framework running on top of Mesos consists of two components: a scheduler that registers with the master to be offered resources, and an executor process that is launched on slave nodes to run the framework’s tasks.

Master determines how many resources are offered to each framework, the frameworks’ schedulers select which of the offered resources to use.

When a frameworks accepts offered resources, it passes to Mesos a description of the tasks it wants to run on them.

Frameworks achieve data locality by rejecting offers.

Mesos can reallocate resources if cluster becomes filled with long tasks by revoking (killing) tasks with grace period.

Isolation through existing OS isolation mechanisms usually system containers. These technologies can limit the CPU, memory, network bandwidth and I/O usage of a process tree.

Mesos lets them short-circuit the rejection process and avoid communication by providing filters to the master. We support two types of filters: “only offer nodes from list L” and “only offer nodes with at least R resources free”.

Two types of resources: mandatory and preferred

 - A resource is mandatory if a framework must acquire it in order to run.
 - Preferred if a framework performs “better” using it, but can also run using another equivalent resource.


#### Notes

> Two-level scheduling mechanism called resource offers

> Delegating control over scheduling to the framework

> Resource offer - encapsulates a bundle of resources that a framework can allocate on a cluster node to run tasks

> Framework ramp-up time - time it takes a new framework to achieve its allocation

> Job completion time - time it takes a job to complete, assuming one job per framework;

> System utilization - total cluster utilization.

> Scale up - frameworks can elastically increase their allocation to take advantage of free resources.

> Scale down - frameworks can relinquish resources without significantly impacting their performance.

> Minimum allocation - frameworks require a certain minimum number of slots before they can start using their slots.

> Task distribution - distribution of the task durations. We consider both homogeneous and heterogeneous distributions.

+++
name = "Large-scale cluster management at Google with Borg"
author = "Abhishek Verma, Luis Pedrosa, Madhukar R. Korupolu, David Oppenheimer, Eric Tune, John Wilkes"
tags = ["paper"]
description = "Take away from scientific papers, discussion & summary"
link = "https://research.google.com/pubs/pub43438.html"
title = "Large-scale cluster management at Google with Borg"
date = "2017-02-09T20:27:52+02:00"

+++

## Borg

Cluster manager that runs hundreds of thousands of jobs, from many thousands of
different applications, across a number of clusters each with up to tens of thousands of machines.

3 main benefits:

  - hides the details of resource management and failure handling so its users can
    focus on application development instead
  - operates with very high reliability and availability, and supports applications that do the same
  - lets us run workloads across tens of thousands of machines effectively

A key design feature in Borg is that already-running tasks
continue to run even if the Borgmaster or a task’s Borglet
goes down. But keeping the master up is still important
because when it is down new jobs cannot be submitted
or existing ones updated, and tasks from failed machines
cannot be rescheduled.

Each job runs in one Borg cell, a set of machines that are managed as a unit.

The machines in a cell belong to a single cluster. A cluster lives inside a single datacenter
building, and a collection of buildings makes up a site.

Median cell size is about 10 k machines after excluding test cells; some are
much larger.
The machines in a cell are heterogeneous in many dimensions: sizes (CPU, RAM,
disk, network), processor type, performance, and capabilities such as an
external IP address or flash storage.
Borg isolates users from most of these differences by determining where in a
cell to run tasks, allocating their resources, installing their programs and
other dependencies, monitoring their health, and restarting them if they fail.

### Jobs

A Borg alloc (short for allocation) is a reserved set of resources on a machine
in which one or more tasks can be run; the resources remain assigned whether or
not they are used.

Quota is used to decide which jobs to admit for scheduling.
Quota is expressed as a vector of resource quantities (CPU, RAM, disk, etc.)
at a given priority, for a period of time (typically months).

Every job has a priority, a small positive integer. A high priority task
can obtain resources at the expense of a lower priority one,
even if that involves preempting (killing) the latter.

### Naming and monitoring

BNS (DNS) for Borg jobs for each task that includes the cell name, job name, and task number.
Borg writes the task’s hostname and port into a consistent,
highly-available file in Chubby with this name, which
is used by our RPC system to find the task endpoint.

Borg also writes job size and task health information into
Chubby whenever it changes, so load balancers can see
where to route requests to.

Borg monitors the health-check URL and restarts
tasks that do not respond promptly or return an HTTP error code.

### Architecture

A Borg cell consists of a set of machines, a logically centralized
controller called the Borgmaster, and an agent process
called the Borglet that runs on each machine in a cell.

Borgmaster process handles client RPCs that either
mutate state (e.g., create job) or provide read-only access
to data (e.g., lookup job).

## Scheduling

The scheduling algorithm has two parts: feasibility checking, to find
machines on which the task could run, and scoring, which picks
one of the feasible machines.

To reduce task startup time, the scheduler prefers to assign
tasks to machines that already have the necessary packages.

Borg distributes packages to machines in parallel using tree-
and torrent-like protocols.

## Borglet

Borglet is a local Borg agent that is present on every
machine in a cell. It starts and stops tasks; restarts them if
they fail; manages local resources by manipulating OS kernel settings;
rolls over debug logs; and reports the state of the
machine to the Borgmaster and other monitoring systems.

#### Notes

> BNS - Borg name system

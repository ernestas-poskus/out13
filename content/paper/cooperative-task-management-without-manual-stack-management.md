+++
link = "https://pdfs.semanticscholar.org/8ed9/eb8f7897701726bc4771aa283dd45c143db1.pdf"
name = "Cooperative Task Management without Manual Stack Management"
author = "Atul Adya, Jon Howell, Marvin Theimer, William J. Bolosky, John R. Douceur"
tags = ["paper","learning","research","review","scientific"]
date = "2017-08-10T21:48:19+03:00"
title = "Cooperative Task Management without Manual Stack Management"
description = "Take away from scientific papers, discussion & summary"

+++

## Or, Event-driven Programming is Not the Opposite of Threaded Programming

Two programming styles as a conflation of two concepts: task
management and stack management.

Those two concerns define a two-axis space in which 'multithreaded' and 'event-driven'
programming are diagonally opposite; there is a third 'sweet spot'
in the space that combines the advantages of both programming styles.

Different task management approaches offer different granularities
of atomicity on shared state.
Conflict management considers how to convert available atomicity
to a meaningful mechanism for avoiding resource conflicts.

High-performance programs are often written with preemptive task management,
wherein execution of tasks can interleave on uniprocessors
or overlap on multiprocessors.

The opposite approach, serial task management, runs each task to completion
before starting the next task. Its advantage is that there is no conflict
of access to the shared state; one can define inter-task invariants on
the shared state and be assured that, while the present task is running,
no other tasks can violate the invariants.

A compromise approach is cooperative task management.
In this approach, a task’s code only yields control to other tasks
at well-defined points in its execution; usually only when the task
must wait for long-running I/O. The approach is valuable when tasks
must interleave to avoid waiting on each other’s I/O, but multiprocessor
parallelism is not crucial for good application performance.

#### Notes

> Preemptive - wherein execution of tasks can interleave on uniprocessors or overlap on multiprocessors.

> Serial task management, runs each task to completion before starting the next task.

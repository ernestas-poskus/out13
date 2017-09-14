+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper", "learning", "research", "review", "scientific", "kernel"]
author = "Dmitry Evtyushkin, Dmitry Ponomarev, Nael Abu-Ghazaleh"
name = "Jump Over ASLR: Attacking Branch Predictors to Bypass ASLR"
link = "http://www.cs.ucr.edu/~nael/pubs/micro16.pdf"
title = "Attacking Branch Predictors to Bypass ASLR"
date = "2017-09-14T19:48:19+03:00"
+++

## Address Space Layout Randomization

ASLR is widely used technique that protects systems against range of attacks.

ASLR works by randomizing the offset of key program segments
in virtual memory, making it difficult for an attacker to derive
the addresses of specific code objects and consequently redirect
the control flow to this code.

Purpose of ASLR is to make it difficult, if not impossible, for the
attacker to know the location of specific code pages in program address space.

Exploited collisions in shared BTBs to create BTB
side-channels and allow the attacker process to recover the
memory layout of both the kernel and user-level applications.

#### Notes

> ASLR - address space layout randomization

> BTB - branch target buffer

> ROP - return oriented programming

> KASLR - Kernel ASLR

> JOP - jump oriented programming

> SDC - same domain collisions

> PDE - page directory entry

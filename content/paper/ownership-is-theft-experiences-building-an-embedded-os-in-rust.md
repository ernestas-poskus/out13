+++
date = "2016-08-25T20:39:03+03:00"
title = "Ownership is theft experiences building an embedded os in rust"
description = "Take away from scientific papers, discussion & summary"
tags = ["paper", "rust"]
author = "Amit Levy, Michael P Andersen, Bradford Campbell, David Culler,Prabal Dutta, Branden Ghena, Philip Levis, Pat Pannuto"
name = "Ownership is Theft: Experiences Building an Embedded OS in Rust"
link = "https://sing.stanford.edu/site/publications/levy-plos15-tock.pdf"

+++

## Embedded OS in Rust

Embedded systems:

 - lack hardware protection mechanism
 - less tolerant to crashes
 - no easy way for debugging
 - GC introduces non-deterministic delay

### Rust

Rust, a new systems programming language, provides compile-time memory safety checks to help eliminate runtime bugs that manifest from improper memory management.

Rust’s ownership model prevents otherwise safe resource sharing common in the embedded domain, conflicts with the reality of hardware resources, and hinders using closures for programming asynchronously.

Rust achieves memory and type safety without garbage collection by using mechanism, derived from affine type and unique pointers, called ownership.

Preserved type safety without relying on a runtime GC for memory management.

Allows the programmer to explicitly separate code which is strictly bound to the type system from code which may subvert it.

> Borrowing

 - Value can only be mutably borrowed if there are no other borrows of the value.
 - Borrows cannot outlive the value they borrow. This prevents dangling pointer bugs.

### Execution context (extension for Rust)

Reflects the thread of a value's owner in its type.

Allows multiple borrows of a value from within same thread, but not across threads.

The goal of execution context is to allow program mutably borrow values multiple times as long as those borrows are never shared between threads.

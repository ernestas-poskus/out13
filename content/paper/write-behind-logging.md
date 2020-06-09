+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper"]
author = "Joy Arulraj, Matthew Perron, Andrew Pavlo"
name = "Write Behind Logging"
link = "http://www.vldb.org/pvldb/vol10/p337-arulraj.pdf"
title = "Write Behind Logging"
date = 2018-01-11T19:27:57+02:00
+++

## WBL vs WAL

Design of the logging and recovery components of database
management systems (DBMSs) has always been influenced by the
difference in the performance characteristics of volatile (DRAM)
and non-volatile storage devices (HDD/SSDs). The key assumption
has been that non-volatile storage is much slower than DRAM and
only supports block-oriented read/writes. But the arrival of new nonvolatile
memory (NVM) storage that is almost as fast as DRAM with
fine-grained read/writes invalidates these previous design choices.

### Recovery Principles

A DBMS guarantees the integrity of a database by ensuring

 - (1) that all of the changes made by committed transactions are durable
 - (2) that none of the changes made by aborted transactions or
   transactions that were running at the point of a failure are visible after recovering from the failure.

These two constraints are referred to as durability of updates and failure atomicity, respectively.

DBMS must protect against:

 - (1) transaction failure
 - (2) system failure
 - (3) media failure

### WAL

The most well-known recovery method based on WAL is the
ARIES protocol developed by IBM in the 1990s. ARIES
is a physiological logging protocol where the DBMS combines a
physical redo process with a logical undo process. During
normal operations, the DBMS records transactions’ modifications
in a durable log that it uses to restore the database after a crash.

A disk-oriented DBMS maintains two meta-data tables at runtime
that it uses for recovery. The first is the dirty page table (DPT) that
contains the modified pages that are in DRAM but have not been
propagated to durable storage.

The second table is the active transaction table (ATT) that tracks the status
of the running transactions.

With an in-memory DBMS, transactions access tuples through
pointers without indirection through a buffer pool.

#### Runtime operations

 - Execute the operation.
 - Write changes to table heap on DRAM.
 - Construct a log record based on operation.
 - Append log record to log entry buffer.

#### Commit processing

 - Collect log entries from log entry buffers.
 - Sync the collected entries on durable storage.
 - Mark all the transactions as committed.
 - Inform workers about group commit.

#### Checkpointing

 - Construct checkpoint containing after-images of visible tuples.
 - Write out transactionally consistent checkpoint to durable storage.
 - Truncate unnecessary log records.

### WBL

Write-behind logging leverages fast, byte-addressable
NVM to reduce the amount of data that the DBMS records in the
log when a transaction modifies the database.
Individual bytes in NVM can be accessed by the processor, and hence
there is no need to organize tuples into pages or go through the I/O subsystem.

WBL reduces data duplication by flushing changes to the database in
NVM during regular transaction processing.

Log is always (slightly) behind the contents of the database.

#### Runtime operations

 - Execute the operation.
 - Write changes to table heap on DRAM.
 - Add an entry to the DTT for that modification.

#### Commit processing

 - Determine dirty tuples using the DTT.
 - Compute changes for this group commit.
 - Sync dirty blocks to durable storage.
 - Sync a log entry containing changes.
 - Inform workers about group commit.

#### Checkpointing

 - Construct a checkpoint containing only the active commit identifier gaps (no after-images).
 - Write out transactionally consistent checkpoint to durable storage.
 - Truncate unnecessary log records.

#### Notes

> WBL - write behind logging

> NVM - non volatile memory

> DBMS - database management system

> DPT - dirty page table

> ATT - active transaction table

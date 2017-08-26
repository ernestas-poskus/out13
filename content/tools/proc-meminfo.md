+++
title = "/proc/meminfo - memory usage statistics"
tags = ["linux","operating systems","tools"]
description = "Review of performance tools and ways to use them"
link = "http://man7.org/linux/man-pages/man5/proc.5.html"
name = "/proc/meminfo"
date = "2017-07-31T19:25:47+03:00"
cmd = "cat /proc/meminfo"

+++

Reports statistics about memory usage on the system.
Useful for inspecting more granular memory usage.

> Sample output below.

```
$ cat /proc/meminfo
MemTotal:       12189912 kB
MemFree:          231992 kB
MemAvailable:    4174992 kB
Buffers:          430884 kB
Cached:          4515856 kB
SwapCached:           60 kB
Active:          8019760 kB
Inactive:        3120804 kB
Active(anon):    6121448 kB
Inactive(anon):  1099620 kB
Active(file):    1898312 kB
Inactive(file):  2021184 kB
Unevictable:        3088 kB
Mlocked:            3088 kB
SwapTotal:      12467708 kB
SwapFree:       12467352 kB
Dirty:              1568 kB
Writeback:             0 kB
AnonPages:       6196916 kB
Mapped:          1339276 kB
Shmem:           1027248 kB
Slab:             584576 kB
SReclaimable:     349700 kB
SUnreclaim:       234876 kB
KernelStack:       20800 kB
PageTables:        93864 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:    18562664 kB
Committed_AS:   19186116 kB
VmallocTotal:   34359738367 kB
VmallocUsed:           0 kB
VmallocChunk:          0 kB
HardwareCorrupted:     0 kB
AnonHugePages:   1019904 kB
CmaTotal:              0 kB
CmaFree:               0 kB
HugePages_Total:       0
HugePages_Free:        0
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
DirectMap4k:      386324 kB
DirectMap2M:    12083200 kB
DirectMap1G:           0 kB
```

### Interpreting each one

> MemTotal:       12189912 kB

Total usable RAM memory minus reserved bits and kernel binary.

> MemFree:          231992 kB

Sum of `LowFree` and `HighFree`.

> MemAvailable:    4174992 kB

Self explanatory.

> Buffers:          430884 kB

Temporary storage for raw disk blocks.

> Cached:          4515856 kB

Does not include `SwapCached` in memory cache of files read from disk.

> SwapCached:           60 kB

Memory once swapped out and swapped back in but still also in the swap file.
In event of memory pressure swapped pages don't need to be swapped out again
since they are already present.

> Active:          8019760 kB

Memory used more recently.

> Inactive:        3120804 kB

Memory used less recently. More eligible to be reclaimed.

> Active(anon):    6121448 kB

Anonymous memory used more recently.

> Inactive(anon):  1099620 kB

Anonymous memory used less recently, can be swapped out.

> Active(file):    1898312 kB

Page cache (file cache) memory used more recently.

> Inactive(file):  2021184 kB

Page cache (file cache) memory used less recently, can be reclaimed.

> Unevictable:        3088 kB

Self explanatory.

> Mlocked:            3088 kB

Locked memory pages using `mlock()` syscall, unevictable.

> SwapTotal:      12467708 kB

Self explanatory.

> SwapFree:       12467352 kB

Remaining swap space available.

> Dirty:              1568 kB

Memory waiting to be written to disk.

> Writeback:             0 kB

Memory being actively written to disk.

> AnonPages:       6196916 kB

Non-file backed pages mapped in userland.

> Mapped:          1339276 kB

Files mapped to memory using `mmapp()` syscall.

> Shmem:           1027248 kB

Amount of memory consumed by `tmpfs` file system.

> Slab:             584576 kB

Kernel data structures cache. More details in `/proc/slabinfo` || slabtop.
Sum of `SReclaimable` +  `SUnreclaim`.

> SReclaimable:     349700 kB

Part of `Slab` cache that can be reclaimed.

> SUnreclaim:       234876 kB

Part of `Slab` cache, unevictable.

> KernelStack:       20800 kB

Amount of memory allocated for Kernel stacks.

> PageTables:        93864 kB

Memory allocated to lowest levels of page tables.

> NFS_Unstable:          0 kB

NFS file system pages sent to server but not yet committed.

> Bounce:                0 kB

Memory used in block device **bounce** buffers.

> WritebackTmp:          0 kB

Memory used in `FUSE` (file system in user space) for temporary
write-back buffers.

> CommitLimit:    18562664 kB

Over commit memory limit adhered only if if strict over commit
accounting is enabled in `/proc/sys/vm/overcommit_memory`.

> Committed_AS:   19186116 kB

Amount of memory presently allocated on the system.
Committed memory is the sum of all memory allocated by the process,
even if it has not been 'used'.

> VmallocTotal:   34359738367 kB

Total sizes of `vmalloc` memory area.

> VmallocUsed:           0 kB

Amount of `vmalloc` area which is used.

> VmallocChunk:          0 kB

Largest contiguous block of `vmalloc` memory which is being used.

> HardwareCorrupted:     0 kB

Amount of memory Kernel identified as corrupted.

> AnonHugePages:   1019904 kB

Non-file backed huge pages mapped into userland page tables.

> CmaTotal:              0 kB

Total pages allocated by contiguous memory allocator `CMA`.

> CmaFree:               0 kB

Free contiguous memory allocator pages.

> HugePages_Total:       0

Size of the huge page tables.

> HugePages_Free:        0

Number of huge page tables that are free.

> HugePages_Rsvd:        0

Number of huge page tables where commitment to allocate has been made,
but actual allocation is not yet completed.

> HugePages_Surp:        0

Number of huge page tables above allowed value.

> Hugepagesize:       2048 kB

Size of huge page.

> DirectMap4k:      386324 kB

Number of bytes linearly mapped by Kernel in 4KB pages.

> DirectMap2M:    12083200 kB

Number of bytes linearly mapped by Kernel in 2M pages.

> DirectMap1G:           0 kB

Number of bytes linearly mapped by Kernel in 1G pages.

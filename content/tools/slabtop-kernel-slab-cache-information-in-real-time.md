+++
tags = ["linux","operating systems","performance","tools"]
description = "Review of performance tools and ways to use them"
date = "2017-06-28T18:57:58+03:00"
title = "slabtop - kernel slab cache information in real time"
author = "Chris Rivera and Robert Love"
cmd = "slabtop -osc"
params = "-o = display once, -sc = sort by cache size column"

+++

Displays detailed kernel slab cache information by aggregating `/proc/slabinfo`.

Tool shows a glimpse into kernel data structures.


> Sample output below.

```
root@ow:~# slabtop -osc | head -n 20
 Active / Total Objects (% used)    : 4649227 / 4694474 (99.0%)
 Active / Total Slabs (% used)      : 153429 / 153429 (100.0%)
 Active / Total Caches (% used)     : 82 / 118 (69.5%)
 Active / Total Size (% used)       : 1259115.61K / 1273939.45K (98.8%)
 Minimum / Average / Maximum Object : 0.01K / 0.27K / 18.50K

  OBJS ACTIVE   USE OBJ SIZE  SLABS OBJ/SLAB CACHE SIZE NAME
327090 325687   99%    1.05K  10903       30    348896K ext4_inode_cache
1639344 1639083 99%    0.19K  78064       21    312256K dentry
217504 216070   99%    1.00K   6797       32    217504K ecryptfs_inode_cache
196352 192576   98%    0.61K   7552       26    120832K proc_inode_cache
519792 519564   99%    0.10K  13328       39     53312K buffer_head
410976 401655   97%    0.12K  12843       32     51372K kmalloc-128
 76020  72290   95%    0.57K   2715       28     43440K radix_tree_node
 19808  18066   91%    1.00K    619       32     19808K kmalloc-1024
 24668  24322   98%    0.55K    881       28     14096K inode_cache
432640 432640  100%    0.03K   3380      128     13520K kmalloc-32
337416 335756   99%    0.04K   3308      102     13232K ext4_extent_status
 56220  51332   91%    0.20K   2811       20     11244K vm_area_struct
122368 116329   95%    0.06K   1912       64      7648K kmalloc-64
```

Interesting object are `ext4_inode_cache` and `dentry`.
These are fs cache objects they speed up fs file/directory access.
The `ext4_inode_cache` is underlying fs `kmem_cache` structure cache.

> Initialized in `fs/ext4/super.c`.

```c
static int __init init_inodecache(void)
{
	ext4_inode_cachep = kmem_cache_create("ext4_inode_cache",
					     sizeof(struct ext4_inode_info),
					     0, (SLAB_RECLAIM_ACCOUNT|
						SLAB_MEM_SPREAD|SLAB_ACCOUNT),
					     init_once);
	if (ext4_inode_cachep == NULL)
		return -ENOMEM;
	return 0;
}
```

The `dentry` is `kmem_cache` structure cache.

> Initialized in `kernel/fs/dcache.c`

```c
static void __init dcache_init(void)
{
	unsigned int loop;

	/*
	 * A constructor could be added for stable state like the lists,
	 * but it is probably not worth it because of the cache nature
	 * of the dcache.
	 */
	dentry_cache = KMEM_CACHE(dentry,
		SLAB_RECLAIM_ACCOUNT|SLAB_PANIC|SLAB_MEM_SPREAD|SLAB_ACCOUNT);

	/* Hash may have been set up in dcache_init_early */
	if (!hashdist)
		return;

	dentry_hashtable =
		alloc_large_system_hash("Dentry cache",
					sizeof(struct hlist_bl_head),
					dhash_entries,
					13,
					0,
					&d_hash_shift,
					&d_hash_mask,
					0,
					0);

	for (loop = 0; loop < (1U << d_hash_shift); loop++)
		INIT_HLIST_BL_HEAD(dentry_hashtable + loop);
}
```

These objects are freed automatically by kernel if there is memory pressure.


> To forcefully clean system slab cache.

```
echo 3 > /proc/sys/vm/drop_caches # free pagecache, dentries and inodes
```

+++
Description = "Using mnemonics in software engineering to remember arguments & commands"
Tags = ["personal", "mnemonics"]
title = "Socket statistics"
date = 2017-11-14T20:35:44+02:00
man = "https://linux.die.net/man/8/ss"
cmd = "ss"
+++

-n, --numeric | Do now try to resolve service names.

-a, --all | Display all sockets.

-e, --extended | Show detailed socket information

-p, --processes | Show process using socket.

-t, --tcp | Display only TCP sockets.

-u, --udp | Display only UDP sockets.

-l, --listening | Display only listening sockets (these are omitted by default).

-4, --ipv4 | Display only IP version 4 sockets (alias for -f inet).

-6, --ipv6 | Display only IP version 6 sockets (alias for -f inet6).

---

These are the main command arguments, it may be hard to remember them all.
However if arguments are combined into known words or combinations it becomes
cheerfully easy to recall them.

 - ss -nepal: do not resolve, extended, process pids, all and listening
 - ss -puta: process pids, UDP & TCP, all
 - ss -tuna: TCP & UDP, do not resolve, all

Furthermore one can use images to aid visualization of commands.

![](/images/mnemonics/nepal.png)
![](/images/mnemonics/puta.jpg)
![](/images/mnemonics/tuna.png)

 - Nepal - nice country
 - Puta - whore
 - Tuna - simply popular fish

Imagine whore eating whole tuna in Nepal on mountain Everest.

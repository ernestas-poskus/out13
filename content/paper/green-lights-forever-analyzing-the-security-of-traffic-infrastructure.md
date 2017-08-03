+++
name = "Green Lights Forever: Analyzing the Security of Traffic Infrastructure"
author = "Branden Ghena, William Beyer, Allen Hillaker, Jonathan Pevarnek, and J. Alex Halderman"
tags = ["paper","learning","research","review","scientific"]
date = "2017-08-03T19:10:33+03:00"
title = "Analyzing the Security of Traffic Infrastructure"
description = "Take away from scientific papers, discussion & summary"
link = "https://jhalderm.com/pub/papers/traffic-woot14.pdf"

+++

## Green Lights Forever

Safety critical nature of traffic infrastructure requires that
it be secure against computer-based attacks.

Traffic signals were originally designed as standalone hardware,
each running on fixed timing schedules, but have evolved into more
complex, networked systems.

Traffic controllers now store multiple timing plans, integrate varied
sensor data, and even communicate with other intersections in order
to better coordinate traffic.

Wireless networking has helped to mitigate
these costs, and many areas now use intelligent wireless
traffic management systems.

### Controllers

Traffic controllers read sensor inputs and control light
states. The controller is typically placed in a metal cabinet
by the roadside along with relays to activate the traffic lights.

### Communications

Radios commonly operate in the ISM band at 900 MHz or 5.8 GHz,
or in the 4.9 GHz band allocated for public safety.

### MMU

Malfunction management unit, also referred to as conflict management
units, are hardware-level safety mechanisms

Valid configurations are stored on a circuit board rather than
in software, with safe configurations literally wired together.

If an unsafe configuration (e.g. conflicting green lights) is
detected, the MMU overrides the controller and forces
the lights into a known-safe configuration.

The MMU also ensures that durations of lights are long enough.
Too short of a yellow or red light duration will trigger a fault.

### Network

One intersection acts as a root node and connects back to a
management server under the control of the road agency.
Intersections often have two radios, one slave radio to transmit
to the next intersection towards the root and one master
radio to receive from one or more
child nodes beyond it. All devices form a single private
network and belong to the same IP subnet.

The proprietary protocol is similar to 802.11 and
broadcasts an SSID which is visible from standard laptops
and smartphones but cannot be connected to.

The wireless connections are unencrypted and the radios
use factory default usernames and passwords.

### Controller

A single controller at each intersection reads sensor data
and controls the traffic lights and pedestrian signs. Many
settings on the controller are programmable, including
light timing parameters.

All of the settings on the controller may be configured
via the physical interface on the controller, but they may
also be modified though the network. An FTP connection
to the device allows access to a writable configuration
database. This requires a username and password, but
they are fixed to default values which are published online
by the manufacturer. It is not possible for a user to modify
the FTP username or password.

The controller runs the VxWorks 5.5 real-time operating system.

### Findings

 - The network is accessible to attackers due to the lack of encryption.
 - Devices on the network lack secure authentication due to the use of default usernames and passwords.
 - The traffic controller is vulnerable to known exploits.

### Types of attacks

DOS - A denial of service attack in this context refers to stopping normal light functionality.
  The most obvious way to cause a loss of service is to set all lights to red.

Traffic Congestion - more subtly, attacks could be made against the entire traffic infrastructure
  of a city which would manipulate the timings of an intersection relative to its neighbors.

Light Control - An attacker can also control lights for personal gain.
  Lights could be changed to be green along the route the attacker is driving.

#### Notes

> MMU - Malfunction management unit

> Denial of service

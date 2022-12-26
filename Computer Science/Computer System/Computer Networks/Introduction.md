# Introduction

[toc]

## Introduction

- Network Edge: hosts(clients and servers), Access networks, physical media
- Network Core: interconnected routers, network of networks

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1gzzhrvvcavj20u0170wjj.jpg" alt="image-20220305101321732" style="zoom:50%;" />

## Network Edge

### Access Networks

Access network is the network that physically connects an end system to the first router (also known as the “edge router”) on a path from the end system to any other distant end system.

- Home Access: 
  - Digital Subscriber Line(DSL): A residence typically obtains DSL Internet access from the same local telephone company (telco) that provides its wired local phone access. The residential telephone line carries both data and traditional telephone signals simultaneously, which are encoded at different frequencies
  - Cable Internet Access: **cable Internet access** makes use of the cable television company’s existing cable television infrastructure.
  - Fiber to the Home(FTTH): provide an optical fiber path from the CO directly to the home.
- Enterprise Access:
  - Ethernet: Ethernet users use twisted-pair copper wire to connect to an Ethernet switch
  - WIFI: Wireless LAN access based on IEEE 802.11 technology, more colloquially known as WiFi
- Wider area Wireless Access: 3G, LTE 4G and 5G

### Physical Media

- Twisted-Pair Copper Wire: The least expensive and most commonly used guided transmission medium is twisted-pair copper wire.
- Coaxial Cable:  coaxial cable consists of two copper conductors, but the two con- ductors are concentric rather than parallel.
- Fiber Optics: An optical fiber is a thin, flexible medium that conducts pulses of light, with each pulse representing a bit.

## Network Core

### Packet Switching

- Hosts break application-layer messages into packets, and network forwards packets from one router to the next, across links on path from source to destination.

- Store-and-Forward Transmission: Store-and-forward transmission means that the packet switch must receive the entire packet before it can begin to transmit the first bit of the packet onto the outbound link.

- Queueing Delays and Packet Loss: For each attached link, the packet switch has an **output buffer** (also called an **output queue**), which stores packets that the router is about to send into that link. If the ouput buffer is full, the incoming packet will be dropped, thus packet loss will occur.

- Forwarding Tables: Each router has a **forwarding table** that maps destination addresses (or portions of the destination addresses) to that router’s outbound links. A routing protocol may, for example, determine the shortest path from each router to each destination and use the shortest path results to configure the forwarding tables in the routers.

### Circuit Switching

- In circuit-switched networks, the resources needed along a path (buffers, link transmission rate) to provide for communication between the end systems are *reserved* for the duration of the communication session between the end systems.
- Frequency Division multiplexing(FDM): the frequency spectrum of a link is divided up among the connections established across the link.
- Time Division Multiplexing(TDM): time is divided into frames of fixed duration, and each frame is divided into a fixed number of time slots.
- Compared to packet switching, circuit switching may have silent periods, thus have low efficiency.

### Network of Networks

- Network Structure 1: interconnects all of the access ISPs with a single global transit ISP.
- Network Structure 2: a two-tier hierarchy with global transit providers residing at the top tier and access ISPs at the bottom tier.
- Network Structure 3: add points of presence (PoPs), multi-homing, peering, and Internet exchange points (IXPs)
- Network Structure 4: consisting of access ISPs, regional ISPs, tier-1 ISPs, PoPs, multi-homing, peering, and IXPs
- Network Structure 5: builds on top of Network Structure 4 by adding content-provider networks. 

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1gzzhrst11kj21900mgn0e.jpg" alt="image-20220305114116158" style="zoom:50%;" />

## Performance

### Packet Delay

$$
d_{nodal}=d_{proc}+d_{queue}+d_{trans}+d_{prop}
$$

- $d_{proc}$: nodal processing, check bit errors, determine output link, etc. typically smaller than microsecs
- $d_{queue}$: queueing delay, time waiting at output link for transmission, depends on congestion level of router
- $d_{trans}$: transmission delay, transmit packet to outbound link, $d_{trans}=\frac{L}{R}$
- $d_{prop}$: propagation delay, propagate the packet to next router, $d_{prop}=\frac{d}{s}$

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1gzzhrojyh3j21580bkabx.jpg" alt="image-20220305115340568" style="zoom:50%;" />

queueing delay: a is average packet arrival rate, L is packet length, and R is link bandwidth

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1gzzhtkuxhyj20hq0eqgm7.jpg" alt="image-20220305115533360" style="zoom:33%;" />

### Throughput

Throughput: rate (bits/time unit) at which bits are being sent from sender to receiver

- Instantaneous: rate at given point in time
- Average: rate over longer period of time

Bottleneck link: link on end-end path that constrains end-end throughput

## Protocol Layers

- Application Layer: The application layer is where network applications and their application-layer protocols reside.
- Transport Layer: The Internet’s transport layer transports application-layer messages between application endpoints.
- Network Layer: The Internet’s network layer is responsible for moving network-layer packets known as **datagrams** from one host to another.
- Link Layer: The Internet’s network layer routes a datagram through a series of routers between the source and destination.
- Physical Layer: the job of the physical layer is to move the *individual bits* within the frame from one node to the next.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1gzzi4k7jz2j21bo0u0n2e.jpg" alt="image-20220305120613299" style="zoom:50%;" />
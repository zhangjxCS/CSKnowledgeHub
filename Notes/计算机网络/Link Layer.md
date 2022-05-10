# Link Layer

[toc]

## Introduction

Link layer has responsibility of transferring datagram from one node to physically adjacent node over a link

- *Framing.* Almost all link-layer protocols encapsulate each network-layer data- gram within a link-layer frame before transmission over the link.
- *Link access.* A medium access control (MAC) protocol specifies the rules by which a frame is transmitted onto the link.

- *Reliable delivery.* When a link-layer protocol provides reliable delivery service, it guarantees to move each network-layer datagram across the link without error.
- *Error detection and correction.* The link-layer hardware in a receiving node can incorrectly decide that a bit in a frame is zero when it was transmitted as a one, and vice versa.

## Error Detecting and Correction

### Parity Checks

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h21fei50aij20u00v840e.jpg" alt="image-20220508114346501" style="zoom:50%;" />

### Cyclic Redundancy Check CRC

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h21iovoqfaj210d0u0402.jpg" alt="image-20220508133732593" style="zoom:50%;" />

## Multiple Access Protocols

- **point-to-point link**: consists of a single sender at one end of the link and a single receiver at the other end of the link.

- **broadcast link**: can have multiple sending and receiving nodes all connected to the same, single, shared broadcast channel.

### Channel Partitioning Protocols

- TDM divides time into **time frames** and further divides each time frame into *N* **time slots**.
- FDM divides the *R* bps channel into different frequencies (each with a bandwidth of *R*/*N*) and assigns each frequency to one of the *N* nodes.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h21izklc8aj213d0u0gnr.jpg" alt="image-20220508134748212" style="zoom:33%;" />

### Random Access Protocols

In a random access protocol, a transmitting node always transmits at the full rate of the channel, namely, *R* bps. When there is a collision, each node involved in the collision repeatedly retransmits its frame (that is, packet) until its frame gets through without a collision.

Slotted ALOHA

- All frames consist of exactly *L* bits.

- Time is divided into slots of size *L*/*R* seconds (that is, a slot equals the time to

  transmit one frame).

- Nodes start to transmit frames only at the beginnings of slots.

- The nodes are synchronized so that each node knows when the slots begin.

- If two or more frames collide in a slot, then all the nodes detect the collision event before the slot ends.

- When the node has a fresh frame to send, it waits until the beginning of the next slot and transmits the entire frame in the slot.

- If there isn’t a collision, the node has successfully transmitted its frame and thus need not consider retransmitting the frame. (The node can prepare a new frame for transmission, if it has one.)

- If there is a collision, the node detects the collision before the end of the slot. The node retransmits its frame in each subsequent slot with probability *p* until the frame is transmitted without a collision.

ALOHA

- when a frame first arrives, the node immediately transmits the frame in its entirety into the broadcast channel.

- If a transmitted frame experiences a collision with one or more other transmissions, the node will then immediately retransmit the frame with probability *p*. Otherwise, the node waits for a frame transmission time.

Carrier Sense Multiple Access CSMA

- *Listen before speaking.* If someone else is speaking, wait until they are finished, which called **carrier sensing**—a node listens to the channel before transmitting.
- *If someone else begins talking at the same time, stop talking.* **collision detection**—a transmitting node listens to the channel while it is transmitting.

CSMA/CD Efficiency
$$
Efficiency=\frac{1}{1+5d_{prop}/d_{trans}}
$$

### Taking Turns

- polling: master node invites other nodes to transmit in turn, typically used with dumb devices

- token passing: control token passed from one node to next sequentially

### DOCSIS

data over cable serrvice interface specification

- FDM over upstream, downstream frequency channels
- TDM upstream: some slots assigned, some have cotention

## Local Area Networks

### Link Layer Address

MAC Address:

- Each interface on LAN has unique 48-bit MAC address and a locally unique 32-bit IP address
- Mac address allocation by IEEE, and manufacturer buys portin of MAC address space

Address resolution protocol ARP

ARP table: each IP node on LAN has table

- IP/MAC address mapping: <IP address, MAC address, TTL(Time to live)>
- MAC address not in ARP table, use ARP to find address

### Ethernet

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h21l5q4elwj21oc0cita7.jpg" alt="image-20220508150255505" style="zoom: 33%;" />

- Bus: all nodes in same collision domain
- Switched: active link-layer 2 switch in center

Ethernet is unreliable connectionless with unslotted CSMA/CD protocol

### Switches

- **Filtering** is the switch function that determines whether a frame should be for- warded to some interface or should just be dropped. 

- **Forwarding** is the switch function that determines the interfaces to which a frame should be directed, and then moves the frame to those interfaces.

- **Self-learning:** its table is built automatically, dynamically, and autonomously— without any intervention from a network administrator or from a configuration pro- tocol.

### VLANs

a switch that sup- ports VLANs allows multiple *virtual* local area networks to be defined over a sin- gle *physical* local area network infrastructure.

In the VLAN trunking approach shown, a special port on each switch is configured as a trunk port to interconnect the two VLAN switches.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h21lou9gv2j21lm0tmq75.jpg" alt="image-20220508152118176" style="zoom:33%;" />

## Link Virtualization

Multiprotocol label switching MPLS

MPLS: path to destination can be based on source and destination address


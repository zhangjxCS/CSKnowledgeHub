# Network Layer

[toc]

## Overview

- Forwarding: the router-local action of transferring a packet from an input link interface to the appropriate output link interface.
- Routing: the network-wide process that determines the end-to-end paths that packets take from source to destination.

### Network Service Model

Best effort service model

- No guarantees on successful datagram delivery to destination
- No guarantees on timing or order of delivery
- No guarantees on bandwidth available to end-end flow

## Router

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1x7ojdzf4j21nb0u042a.jpg" alt="image-20220504201418685" style="zoom: 33%;" />

- Input ports: get the incoming data and perform link layer and physical layer function
- Switching fabric: The switching fabric connects the router’s input ports to its output ports.
- Output ports: An **output port** stores packets received from the switching fabric and transmits these packets on the outgoing link by performing the necessary link-layer and physical-layer functions.
- Routing processor: The routing processor performs control-plane functions.

### Forwarding

- Destination-based forwarding: forward based only on destination IP address
- Generalized forwarding: forward based on any set of header field values

longest prefix matching: when looking for forwarding table entry for given destination address, use longest address prefix that matches destination address.

### Switching

- Switching via memory: traditional computers with switching under CPU. Packet copied to system's memory and speed limited by memory bandwidth
- Switching via bus: datagram from input port memory to output port memory via a shared bus; switching speed limited by bus bandwidth
- Switching via interconnection network: crossbar nets initially developed to connect processors in multiprocessor

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1xathjq0rj218y0u042h.jpg" alt="image-20220504220255007" style="zoom:33%;" />

### Queuing

- Input port queuing: if switch fabric slower than input ports combined, queueing may occur at input queues; queued datagram at front of queue prevents others in queue from moving forward(Head of Line blocking)
- Output port queuing: Buffering required when datagrams arrive from fabric faster than link transmission rate.

Link capacity C, flows N, and Round Trip Time RTT, the buffering equal to
$$
RTT*C/\sqrt N
$$

### Packet Scheduling

- FIFO: The FIFO (also known as first-come-first-served, or FCFS) scheduling discipline selects packets for link transmission in the same order in which they arrived at the output link queue.
- Priority Queuing: Under priority queuing, packets arriving at the output link are classified into priority classes upon arrival at the queue
- Round Robin and Weighted Fair Queuing: Under the round robin queuing discipline, packets are sorted into classes as with priority queuing. A generalized form of round robin queuing that has been widely implemented in routers is the so-called **weighted fair queuing (WFQ) discipline**, which gets weighted amount of service in each cycle in round robin.

## Internet protocol IP

### IP Datagram

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1ydmatqnlj215f0u0whs.jpg" alt="image-20220505202518959" style="zoom:33%;" />

- Version number: These 4 bits specify the IP protocol version of the datagram.
- Header length: these 4 bits are needed to determine where in the IP datagram the payload actually begins.
- Type of service: The type of service (TOS) bits were included in the IPv4 header to allow different types of IP datagrams to be distinguished from each other.
- Datagram length: total length of the IP datagram

- Identifier, flags, fragmentation offset: These three fields have to do with so-called IP fragmentation, when a large IP datagram is broken into several smaller IP datagrams which are then forwarded independently to the destination

- Time-to-live: The time-to-live (TTL) field is included to ensure that datagrams do not circulate forever

- Protocol: This field is typically used only when an IP datagram reaches its final destination.

- Header checksum: The header checksum aids a router in detecting bit errors in a received IP datagram.

- Source and destination IP addresses: When a source creates a datagram, it inserts its IP address into the source IP address field and inserts the address of the ulti- mate destination into the destination IP address field.

- Options: The options fields allow an IP header to be extended.

- Data: the data field of the IP datagram contains the transport-layer segment (TCP or UDP) to be deliv- ered to the destination.

### IPV4 Addressing

- subnet: device interfaces that can physically reach each other without passing through an intervening router

- Classless Interdomain Routing(CIDR): CIDR generalizes the notion of subnet addressing. As with subnet addressing, the 32-bit IP address is divided into two parts and again has the dotted-decimal form *a.b.c.d/x*, where *x* indicates the number of bits in the first part of the address.
- DHCP Dynamic Host Configuration Protocol: dynamically get address from as server
  - host broadcasts DHCP discover msg
  - DHCP server responds with DHCP offer msg
  - host requests IP address
  - DHCP server sends address: DHCP ack msg

### Network Address Translation NAT

All devices in local network share just one IPv4 address as far as outside world is concerned

- Outgoing datagrams: replace (sourceIP, Port) of every outgoing datagram to (NAT IP address, new port)
- Remember every (sourceIP, port) to (NAT IP address, new port) pair in NAT translation table
- Incoming datagrams: replace (NAT IP address, new port) in destination fields of every incoming datagram with corresponding (sourceIP, Port) stored in NAT table

### IPv6

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1z2sky5xuj21kv0u0q69.jpg" alt="image-20220506105620648" style="zoom:33%;" />

Tunneling: IPv6 datagram carried as payload in IPv4 datagram among IPv4 routers

### Generalized Forwarding

Each router contains a forwarding table (flow table)

- Match: pattern values in packet header fields
- actions: for matched packet, drop forward, modify, matched packet or send matched packet to controller
- Priority: disambiguate overlapping patterns
- counters: bytes and packets

Openflow

- Forward packet to ports
- Drop packet
- Modify fields in headers
- Encapsulate and forward to controller

## Middlebox

Middlebox: any intermediary box performing functions apart from normal, standard func- tions of an IP router on the data path between a source host and destination host

- NAT Translation: NAT boxes implement private network addressing, rewriting datagram header IP addresses and port numbers.

- Security Services: Firewalls block traffic based on header-field values or redirect packets for additional processing, such as deep packet inspection (DPI).

- Performance Enhancement: These middleboxes perform services such as com- pression, content caching, and load balancing of service requests to one of a set of servers that can provide the desired service.

## Routing Protocols

Determine good paths from sending hosts to receiving host, through network of routers.

- A **centralized routing algorithm** computes the least-cost path between a source and destination using complete, global knowledge about the network.

- In a **decentralized routing algorithm**, the calculation of the least-cost path is carried out in an iterative, distributed manner by the routers.

- In **static routing algorithms**, routes change very slowly over time, often as a result of human intervention

- In a **load-sensitive algorithm**, link costs vary dynami- cally to reflect the current level of congestion in the underlying link.

### Link State

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1z49ou9abj21iw0u0wih.jpg" alt="image-20220506114720321" style="zoom: 33%;" />

### Distance Vector

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1z4amxodyj21iy0u0jvz.jpg" alt="image-20220506114814488" style="zoom:33%;" />

## Intra-ISP Routing

- Scale: As the number of routers becomes large, the overhead involved in communi- cating, computing, and storing routing information becomes prohibitive.
- Administrative autonomy: the Internet is a network of ISPs, with each ISP consisting of its own network of routers. An ISP generally desires to operate its network as it pleases or to hide aspects of its network’s internal organization from the outside.

Both of these problems can be solved by organizing routers into **autonomous systems (ASs)**, with each AS consisting of a group of routers that are under the same administrative control.

- Intra-AS routing: routing among within same AS
- Inter-AS routing: routing among AS

### Open Shortest Path First OSPF

- each router floods OSPF link state advertisements to all other routers in entire AS
- multiple linkcosts metrics possible: bandwidth, delay
- each router has full topology, use Dijkstra algorihtm to compute forwarding table

Hierarchical OSPF

- Link-state advertisements flooded only in area, or backbone
- Each node has detailed area topology; only knows direction to reach other destinations

## Routing among ISPs

### Border Gateway Protocol BGP

- Obtain prefix reachability information from neighboring ASs.
- Determine the “best” routes to the prefixes.

### Route Advertisement

BGP advertised route: prefix + attributes

- prefix: destination being advertised
- two important attributes
  - AS-PATH: list of CASes through which prefix advertisement has passed
  - NEXT-HOP: indicates specific internal-AS router to next-hop AS

Policy-based routing:

- Gateway receiving route advertisement uses import policy to accept/decline path
- AS policy also determines whether to advertise path to other neighboring ASes

### BGP messages

- OPEN: opens TCP connection to remote BGP peer and authenticates sending BGP peer
- UPDATE: advertises new path
- KEEPALIVE: keeps connection alive in absence of UPDATES; also ACKs OPEN request
- NOTIFICATION: reports errors inprevious msg; also used to close connection

### Route Selection

- Local preference value attribute: policy decision
- Shortest AS-PATH
- Closest NEXT-HOP router: hot potato routing (choose local gateway that has least intra-domain cost)
- additional criteria: ISP only wants to route traffic to/from its customer networks

## Software Defined Networking

- *Flow-based forwarding.* Packet forwarding by SDN-controlled switches can be based on any number of header field values in the transport-layer, network-layer, or link-layer header.
- *Separation of data plane and control plane.* this software exe- cutes on servers that are both distinct and remote from the network’s switches.
- *A programmable network.* The network is programmable through the network- control applications running in the control plane.

 <img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h21dlyfnl5j20y30u0whr.jpg" alt="image-20220508104141325" style="zoom: 50%;" />

- Data plane switchers: implementing generalized data-plane forwarding in hardware; flow table computed under controller supervision; API for table-based switch control
- SDN controller: maintain network state information; interacts with network control applications; interacts with network switches; implemented as distributed system for performance
- Network-control apps: implement control functions using lower-level services; can be provided by 3rd party

### Openflow

- *Configuration.* This message allows the controller to query and set a switch’s configuration parameters.
- *Modify-State.* This message is used by a controller to add/delete or modify entries in the switch’s flow table, and to set switch port properties.

- *Read-State.* This message is used by a controller to collect statistics and counter values from the switch’s flow table and ports.

- *Send-Packet.* This message is used by the controller to send a specific packet out of a specified port at the controlled switch. The message itself contains the packet to be sent in its payload.
- *Flow-Removed.* This message informs the controller that a flow table entry has been removed, for example by a timeout or as the result of a received *modify-state* message.
- *Port-status.* This message is used by a switch to inform the controller of a change in port status.

- *Packet-in.* a packet arriving at a switch port and not matching any flow table entry is sent to the controller for additional processing.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h21e2wf2kzj20u40u0q5s.jpg" alt="image-20220508105801254" style="zoom:50%;" />

## ICMP

The Internet Control Message Protocol (ICMP), specified in [RFC 792], is used by hosts and routers to communicate network-layer information to each other.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h21e78wbx3j21260u041v.jpg" alt="image-20220508110212264" style="zoom:50%;" />

## Network Management

Network management includes the deployment, integration, and coordination of the hardware, software, and human elements to monitor, test, poll, configure, ana- lyze, evaluate, and control the network and element resources to meet the real-time, operational performance, and Quality of Service requirements at a reasonable cost.

- *Managing server.* The managing server is an application, typically with **network managers** (humans) in the loop, running in a centralized network management station in the network operations center (NOC).

- *Managed device.* A managed device is a piece of network equipment (including its software) that resides on a managed network.

- *Data.* Each managed device will have data, also known as “state,” associated with it.

- *Network management agent.* The network management agent is a software pro- cess running in the managed device that communicates with the managing server, taking local actions at the managed device under the command and control of the managing server.

- *Network management protocol.* This protocol runs between the managing server and the managed devices, allowing the managing server to query the status of managed devices and take actions at these devices via its agents.


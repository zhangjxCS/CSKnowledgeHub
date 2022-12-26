# Transport Layer

[toc]

## Transport Layer Services

A transport-layer protocol provides for **logical communication** between application processes running on different hosts.

Transport-layer protocols live in the end systems. Within an end system, a transport protocol moves messages from application processes to the network edge (that is, the network layer) and vice versa

Sender:

- Pass an application layer message
- Determine segment header filed values
- Create segment and pass segment to IP

Receiver:

- Receive segment from IP
- Check header values
- Extracts application layer message
- Demultiplex message up to application via socket

## Multiplexing and Demultiplexing

This job of delivering the data in a transport-layer segment to the correct socket is called **demultiplexing**. 

The job of gathering data chunks at the source host from different sockets, encapsulating each data chunk with header information (that will later be used in demultiplexing) to create segments, and passing the segments to the network layer is called **multiplexing**.

- Connectionless: When a UDP socket is created in this manner, the transport layer automatically assigns a port number to the socket. UDP socket is fully identified by a two-tuple consisting of a destination IP address and a destination port number. As a consequence, if two UDP segments have different source IP addresses and/or source port numbers, but have the same *destination* IP address and *destination* port number, then the two segments will be directed to the same destination process via the same destination socket.
- Connection-Oriented: TCP socket is identified by a four-tuple: (source IP address, source port number, destination IP address, destination port number). When a TCP segment arrives from the network to a host, the host uses all four values to direct (demultiplex) the segment to the appropriate socket. 

## Connectionless Transport UDP

- no connection establishment (which can add RTT delay)

- simple: no connection state at sender, receiver

- small header size

- no congestion control

### UDP Segment Header

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h02wk418yzj20o80k6t9r.jpg" alt="image-20220306151824460" style="zoom:33%;" />

### UDP Checksum

Sender:

- Treat contents of UDP segment as sequence of 16-bit integers
- Checksum is addition of segment content
- Checksum value put into UDP checksum field

Receiver:

- Compute checksum of received segment
- Check if computed checksum equals checksum filed value

## Principles of Reliable Data Transfer

### rdt1.0 

reliable transfer over a reliable channel

- underlying channel not bit errors, and no loss of packets
- Separate Finite State Machines for sender, receiver

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h02wk1nax2j211g0tcmzn.jpg" alt="image-20220306172217999" style="zoom:50%;" />

### rdt 2.0 

channel with bit errors

- Error Detection: a mechanism is needed to allow the receiver to detect when bit errors have occurred.
- Receiver Feedback: the receiver provides explicit feedback to the sender
  - Acknowledgements(ACKs): receiver explicitly tells sender that pkt received OK
  - Negative acknowledgements(NAKs): receiver explicitly tells sender that pkt had errors
- Retransmission: A packet that is received in error at the receiver will be retrans- mitted by the sender.

Stop and wait: sender sends one packet, then waits for receiver response

- sender retransmits current pkt if ACK/NAK corrupted

- sender adds *sequence number* to each pkt

- receiver discards (doesn’t deliver up) duplicate pkt

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h02wjz02tdj20wg0u0ju1.jpg" alt="image-20220306172139385" style="zoom: 50%;" />

### rdt 2.1

Protocol rdt2.1 uses both positive and negative acknowledgments from the receiver to the sender. When an out-of-order packet is received, the receiver sends a positive acknowledgment for the packet it has received. When a corrupted packet is received, the receiver sends a negative acknowledgment.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h00x9dzs9ej216e0u0787.jpg" alt="image-20220306173524860" style="zoom: 50%;" />

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h02wjtq04xj21e10u043v.jpg" alt="image-20220306173542417" style="zoom: 33%;" />

### rdt 2.2

- same functionality as rdt2.1, using ACKs only
- instead of NAK, receiver sends ACK for last pkt received OK; receiver must *explicitly* include seq # of pkt being ACKed 
- duplicate ACK at sender results in same action as NAK: *retransmit current pkt*

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h00xcnbq24j21700okaev.jpg" alt="image-20220306173832320" style="zoom:50%;" />

### rdt 3.0

pipelining: sender allows multiple, “in-flight”, yet-to-be-acknowledged packets

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h036byrbrdj20qa0k0q57.jpg" alt="image-20220308162024860" style="zoom: 67%;" />

### Go-Back-N

- sender of “window” of up to N, consecutive transmitted but unACKed pkts, cumulative ACKs all packets up to, including seq # *n* 

  on receiving ACK(*n*); timeout(n) then retransmit packet n and all higher seq # packets in window

- Receiver always send ACK for correctly-received packet so far, with highest *in-order* seq # and delete out-of-order packet

  <img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h02x2gjbemj218w09u3zz.jpg" alt="image-20220308105953318" style="zoom:50%;" />

  <img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h02wjoi9b3j21660qkwle.jpg" alt="image-20220308104140101" style="zoom:50%;" />

### Selective Repeat

- receiver individually acknowledges all correctly received packets, buffers packets, as needed, for eventual in-order delivery to upper layer

- sender times-out/retransmits individually for unACKed packets, sender maintains timer for each unACKed pkt

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h02wlzenznj213c0ck405.jpg" alt="image-20220308104403314" style="zoom:50%;" />

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h02wmudfy5j21640o244n.jpg" alt="image-20220308104452149" style="zoom:50%;" />

## Connection-Oriented Transport: TCP

- Connection-Oriented: before one application process can begin to send data to another, the two processes must first “handshake” with each other, which is they first establish connection between each other.
- Full-deuplex service: application-layer data can flow from Process A to Process B at the same time as application-layer data flows from Process B to Process A
- Point-to-point: one sender, one receiver
- Reliable: reliable data transfer service ensures that the data stream that a process reads out of its TCP receive buffer is uncorrupted, without gaps, without duplication, and in sequence
- pipelining: TCP congestion and flow control

### TCP Segment

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1ws75nn5dj21490u0acs.jpg" alt="image-20220504111837396" style="zoom:50%;" />

- Sequence number: The sequence number for a segment is the byte-stream number of the first byte in the segment.

- Acknowledgment number: The acknowledgment number that Host A puts in its segment is the sequence number of the next byte Host A is expecting from Host B. TCP only acknowledges bytes up to the first missing byte in the stream, TCP is said to provide cumulative acknowledgments

- Receive window: used to indicate the number of bytes that a receiver is willing to accept

- Header length field: the length of the TCP header in 32-bit words

- Options: used when a sender and receiver negotiate the maximum segment size (MSS) or as a window scaling factor for use in high-speed networks.
- Flag filed: **ACK** is used to indicate that the value carried in the acknowledgment field is valid. **RST**, **SYN**, and **FIN** bits are used for connection setup and teardown. **PSH** bit indicates that the receiver should pass the data to the upper layer immediately. **URG** bit is used to indicate that there is data in this segment that the sending-side upper- layer entity has marked as “urgent.”

### RTT and Timeout

- TCP, like our rdt protocol in Section 3.4, uses a timeout/retransmit mechanism to recover from lost segments. 
- The timeout should be larger than the connection’s round-trip time (RTT)
- Exponential weighted moving average: the new value of **EstimatedRTT** is a weighted combination of the previous value of **EstimatedRTT** and the new value for **SampleRTT**

$$
EstimatedRTT=(1-\alpha)*EstimatedRTT+\alpha*SampleRTT
$$

$$
DevRTT=(1-\beta)*DevRTT+\beta*|SampleRTT-EstimatedRTT|
$$

$$
TimeoutInterval=EstimatedRTT+4*DevRTT
$$

### Reliable Data Transfer

TCP’s reliable data transfer service ensures that the data stream that a process reads out of its TCP receive buffer is uncorrupted, without gaps, without duplication, and in sequence

- TCP timer management procedures use only a *single* retransmission timer, even if there are multiple transmitted but not yet acknowledged segments
- TCP fast retransmit: if sender receives 3 additional ACKs for same data, resend unACKed segment with smallest seq #

### Flow Control

- TCP provides a **flow-control service** to its applications to eliminate the possibility of the sender overflowing the receiver’s buffer.

- TCP provides flow control by having the *sender* maintain a variable called the **receive window**.

$$
LastByteRcvd-LastByteRead\le RcvBuffer
$$

$$
rwnd=RcvBuffer-[LastByteRcvd-LastByteRead]
$$

$$
LastByteSent-LastByteAcked\le rwnd
$$

### Connection Management

TCP 3-way handshake

- The client-side TCP first sends a special TCP segment to the server-side TCP. Set the SYN segment to 1, and the client randomly chooses an initial sequence number and puts the number in the sequence number field
- Once received TCP SYN segment, server allocates the TCP buffers and variables to the connection, and sends a connection-granted segment to the client TCP. Set the SYN bit to 1; the ack field is set to client seqnum + 1; Server chooses its own sequence number
- Upon receiving the SYNACK segment, the client also allocates buffers and variables to the connection.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1wxpysp3qj210h0u0gob.jpg" alt="image-20220504142944636" style="zoom:50%;" />

TCP close connection

- The client application process issues a close command. This causes the client TCP to send a special TCP segment to the server process. The FIN bit is set to 1.
- When the server receives this segment, it sends the client an acknowledgment segment in return.
- The server then sends its own shutdown segment, which has the FIN bit set to 1.
- Finally, the client acknowledges the server’s shutdown segment.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1wxy63cumj21240u0766.jpg" alt="image-20220504143738224" style="zoom:50%;" />

## Principles of Congestion Control

### Congestion Scenario

- Congestion scenario 1: Two senders, a router with infinite buffers; The highest throughput is R/2. The average delay becomes larger and larger to infinity when the sending rate approaches R/2
- Congestion scenario 2: Two senders and a router with finite buffers; The highest throughput is less than R/2 because of the following reasons. Packet will be dropped when arriving to an full buffer. The sender must pperform retransmissions to compensate for dropped packets. Unneeded retransmission by the sender cause a router to use its link bandwidth. 
- Congestion scenario 3: Four senders, routers with finite buffers and multihop paths. The end-to-end throughput goes to zero in the limit of heavy traffic. When a packet is dropped along a path, the transmission capacity used before is wasted.

### Approaches

- End-to-end congestion control: TCP segment loss (as indicated by a timeout or the receipt of three duplicate acknowledgments) is taken as an indication of network congestion, and TCP decreases its window size accordingly.
- Network-assisted congestion control: routers provide explicit feedback to the sender and/or receiver regarding the con- gestion state of the network.

## TCP Congestion Control

- A lost segment implies congestion, and hence, the TCP sender’s rate should be decreased when a segment is lost.

- An acknowledged segment indicates that the network is delivering the sender’s segments to the receiver, and hence, the sender’s rate can be increased when an ACK arrives for a previously unacknowledged segment.
- *Bandwidth probing*. Given ACKs indicating a congestion-free source-to-destina- tion path and loss events indicating a congested path, TCP’s strategy for adjusting its transmission rate is to increase its rate in response to arriving ACKs until a loss event occurs, at which point, the transmission rate is decreased.

### TCP AIMD

Approach: senders increase sending rate until packet loss occurs, then decrease sending rate on loss event.

- Additive Increase: increase sending rate by 1 maximum segment size every RTT until loss detected
- Multiplicative Decrease: cut sending rate in half at each loss event
  - TCP Reno: cut in half on losss detected by triple duplicate ACK
  - TCP Tahoe: cut to 1 maximum segment size MSS when loss detected by timeout

$$
LastByteSent - LastByteAcked\le cwnd
$$

$$
TCP\ rate=\frac{cwnd}{RTT}
$$

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1x1nhnqqgj21mf0u0dih.jpg" alt="image-20220504164545969" style="zoom: 33%;" />

### TCP slow start

- When connection begins, increase rate exponentially until first loss event: Initial cwnd to 1 MSS and double cwnd every RTT
- Variable ssthresh, on loss event, ssthrest is set to 1/2 of condo, when pkt length gets to ssthresh, exponential increase switch to linear

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1x23tpgnbj21g00u0ads.jpg" alt="image-20220504170127860" style="zoom:33%;" />

### TCP Cubic

- Wmax: sending rate at which congestion loss was detected
- CUBIC increases the congestion window as a function of *cube* of the distance between the current time, *t*, and *K*
- After cutting rate in half on loss, initially ramp to Wmax faster, but then approach Wmax more slowly

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h1x1o1nsfnj21bk0u0dim.jpg" alt="image-20220504164618382" style="zoom:33%;" />

### Delay-based congestion control

- RTTmin: minimum observed RTT
- uncongested throughput with congestion window cwnd is cwnd/RTTmin


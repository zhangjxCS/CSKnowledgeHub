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

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20220306151824460.png" alt="image-20220306151824460" style="zoom:33%;" />

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

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20220306172217999.png" alt="image-20220306172217999" style="zoom:50%;" />

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

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20220306172139385.png" alt="image-20220306172139385" style="zoom: 50%;" />

### rdt 2.1

Protocol rdt2.1 uses both positive and negative acknowledgments from the receiver to the sender. When an out-of-order packet is received, the receiver sends a positive acknowledgment for the packet it has received. When a corrupted packet is received, the receiver sends a negative acknowledgment.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h00x9dzs9ej216e0u0787.jpg" alt="image-20220306173524860" style="zoom: 50%;" />

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20220306173542417.png" alt="image-20220306173542417" style="zoom: 33%;" />

### rdt 2.2

- same functionality as rdt2.1, using ACKs only
- instead of NAK, receiver sends ACK for last pkt received OK; receiver must *explicitly* include seq # of pkt being ACKed 
- duplicate ACK at sender results in same action as NAK: *retransmit current pkt*

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1h00xcnbq24j21700okaev.jpg" alt="image-20220306173832320" style="zoom:50%;" />

### rdt 3.0

pipelining: sender allows multiple, “in-flight”, yet-to-be-acknowledged packets

Go-Back-N: sender: “window” of up to N, consecutive transmitted but unACKed pkts 


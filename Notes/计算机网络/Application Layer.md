# Application Layer

[toc]

## Principles of Network Application

At the core of network application development is writing programs that run on different end systems and communicate with each other over the network.

### Architechture

- Client-server: there is an always-on host, called the *server*, which services requests from many other hosts, called *clients*.
- P2P: there is minimal (or no) reliance on dedicated servers in data centers. Instead the application exploits direct communication between pairs of intermittently connected hosts, called *peers*.

### Process Communicating

Process: program running within a host

- within same host, two processes communicate using inter-process communication (defined by OS)
- processes in different hosts communicate by exchanging messages

Sockets: process sends/receives messages to/from its socket

Addressing process: To identify the receiving process, two pieces of information need to be specified: (1) the address of the host and (2) an identifier that specifies the receiving process in the destination host. Identifier includes both IP address and port numbers associated with process on host

### Transport Service

- data integrity: One important service that a transport-layer protocol can potentially provide to an applica- tion is process-to-process reliable data transfer.
- throughput: With such a service, the application could request a guaranteed throughput of *r* bits/sec, and the transport protocol would then ensure that the avail- able throughput is always at least *r* bits/sec.
- timing: A transport-layer protocol can also provide timing guarantees.
- Security: a transport protocol can provide an application with one or more security services.

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20220305124311303.png" alt="image-20220305124311303" style="zoom:50%;" />

## Web

HyperText Transfer Protocol (HTTP): HTTP is implemented in two programs: a client program and a server program. The client program and server program, executing on different end systems, talk to each other by exchanging HTTP messages. HTTP defines how Web clients request Web pages from Web servers and how servers transfer Web pages to clients.

### HTTP 1.1

Non-persistent HTTP:

- Requires 2 RTTs per object
- OS overhead for each TCP connection
- browsers often open multiple parallel TCP connections to fetch referenced objects in parallel

Persistent HTTP:

- server leaves connection open after sending response
- subsequent HTTP messages between same client/server sent over open connection
- Client sends requests as son as it encounters a referenced object
- As little as one RTT for all the referenced objects

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20220305150629948.png" alt="image-20220305150629948" style="zoom:33%;" />

### HTTP Request

- POST: An HTTP client often uses the POST method when the user fills out a form
- GET:  The GET method is used when the browser requests an object, with the requested object identified in the URL field.
- HEAD: When a server receives a request with the HEAD method, it responds with an HTTP message but it leaves out the requested object.
- PUT: It allows a user to upload an object to a specific path (directory) on a specific Web server.
- DELETE: The DELETE method allows a user, or an application, to delete an object on a Web server.

### HTTP Response

- STATUS code: 200(OK); 301(Moved Permanently); 400(Bad Request); 404(Not Found); 505(HTTP Version Not Supported)
- Header: date, server, etc.
- Entity body: response message

### Cookies

HTTP is stateless, cookies is used to keep track of users

- At protocol endpoints: maintain state at sender/receiver over multiple transactions
- In messages: cookies in HTTP messages carry state

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20220305152405197.png" alt="image-20220305152405197" style="zoom: 33%;" />

### Web caches

A network entity that satisfies HTTP requests on the behalf of an origin Web server. The Web cache has its own disk storage and keeps copies of recently requested objects in this storage.

- A cache is both a server and a client at the same time. When it receives requests from and sends responses to a browser, it is a server. When it sends requests to and receives responses from an origin server, it is a client.

- Typically a Web cache is purchased and installed by an ISP. 
- Web cache can substantially reduce the response time for a client request, particularly if the bottleneck bandwidth between the client and the origin server is much less than the bottleneck bandwidth between the client and the cache.
- Web caches can substantially reduce traffic on an institution’s access link to the Internet.

conditional GET: HTTP has a mechanism that allows a cache to verify that its objects are up to date. This mechanism is called the **conditional GET**

- Client: specify date of cached copy in HTTP request
- Server: response contains no object if cached copy is up-to-date

### HTTP 2

Decreased delay in multi-object HTTP requests, increased flexibility at *server* in sending objects to client

- methods, status codes, most header fields unchanged from HTTP 1.1
- transmission order of requested objects based on client-specified object priority
- push unrequested objects to client
- divide objects into frames, schedule frames to mitigate HOL blocking

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20220305153830056.png" alt="image-20220305153830056" style="zoom: 33%;" />

## E-mail

Three major components: user agents; mail servers; simple mail transfer protocol SMTP

### SMTP

- Use TCP to reliably transfer email message from client to server, port 25
- Three phases of transfer: SMTP handshake(Persistent connection), SMTP transfer of messages, SMTP closure
- Command/response interaction

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20220305161459418.png" alt="image-20220305161459418" style="zoom: 50%;" />

### Mail Access Control

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1gzzpdtf6dkj21gm0cy418.jpg" alt="image-20220305161718549" style="zoom:50%;" />

- Alice’s user agent uses SMTP or HTTP to deliver the e-mail message into her mail server 

- Alice’s mail server uses SMTP (as an SMTP cli- ent) to relay the e-mail message to Bob’s mail server.

- If Bob is using Web-based e-mail or a smartphone app (such as Gmail), then the user agent will use HTTP to retrieve Bob’s e-mail. The alternative method, typically used with mail clients such as Microsoft Outlook, is to use the Internet Mail Access Protocol (IMAP)

## Domain Name System

DNS: translate host name to its IP address

- a distributed database implemented in a hierarchy of DNS servers
- an application-layer protocol that allows hosts to query the distributed database

### DNS Services

- Host Aliasing: DNS can be invoked by an application to obtain the canonical hostname for a supplied alias hostname as well as the IP address of the host.
- Mail Server Aliasing: DNS can be invoked by a mail application to obtain the canonical hostname for a supplied alias hostname as well as the IP address of the host.
- Load Distribution: DNS is also used to perform load distribution among repli- cated servers, such as replicated Web servers.

### Distributed Hierarchical Database

- Root DNS Server: Root name servers provide the IP addresses of the TLD servers
- Top-level domain name system: For each of the top-level domains—top-level domains such as com, org, net, edu, and gov, and all of the country top-level domains such as uk, fr, ca, and jp—there is TLD server (or server cluster).
- Authortative DNS servers: Every organization with publicly accessible hosts (such as Web servers and mail servers) on the Internet must provide publicly accessible DNS records that map the names of those hosts to IP addresses.
- Local DNS servers: Local DNS server does not strictly belong to hierarchy. Each ISP—such as a residential ISP or an institutional ISP—has a local DNS server (also called a default name server). When a host connects to an ISP, the ISP provides the host with the IP addresses of one or more of its local DNS servers

A registrar is a commercial entity that verifies the uniqueness of the domain name, enters the domain name into the DNS database, and collects a small fee from you for its services.

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1gzzpxm0r3qj217y0jw76i.jpg" alt="image-20220305163620187" style="zoom:50%;" />

### DNS Caching

- once (any) name server learns mapping, it *caches* mapping, and i*mmediately* returns a cached mapping in response to a query
- cached entries may be *out-of-date*

### DNS Records

- Type A: name is hostname, value is IP address
- Type CNAME: name is alias name for some canonical name, value is canonical name
- Type NS: name is domain, value is hostname of authoritative name server for this domain
- Type MX: value is name of SMTP mail server associated with name

### DNS Messages

## P2P Application

### Scalability

$u_s$ is the server's upload rate, d is the peer download rate

- Client-Server: $D\ge max\{\frac{NF}{u_s},\frac{F}{d_{min}}\}$
- P2P: $D\ge \{\frac{F}{u_s}, \frac{F}{d_{min}}, \frac{NF}{u_s+\sum_{i=1}^Nu_i}\}$

<img src="https://tva1.sinaimg.cn/large/e6c9d24egy1gzzvso0nhuj21520qgwgu.jpg" alt="image-20220305195910361" style="zoom: 33%;" />

### BitTorrent

- File divided into 256Kb chunks
- Peers in torrent send/receive file chunks

BitTorrent file distribution:

- When peer join torrent, it has no chunks, registers with tracker to get list of peers, and connects to subset of peers
- While downloading, peer uploads chunks to other peers
- Peer may change peers with whom it exchanges chunks
- Once peer has entire file, it may leave or remain in torrent

When requesting chunks:

- At any given time, different have different subsets of file chunks
- Periodically, one peer asks each peer for list of chunks that they have
- The peer requests missing chunks from peers, the rule is **rarest first**

When sending chunks:

- One peer sends chunks to those four peers currently sending her chunks at highest rate
- Every 30 secs, ranomly select another peer, starts sending chunks

<img src="/Users/shawnzhang/Library/Application Support/typora-user-images/image-20220305201150385.png" alt="image-20220305201150385" style="zoom:50%;" />

## Video Streaming and CDN

### Video Streaming

- Playout Buffer: compensate for network-added delay, delay jitter.
- Dynamic, Adaptive Streaming over HTTP(DASH): divide video file into multiple chuncks, each encoded at different rates, clients periodically estimate server-to-client bandwidth, and request one chunk at a time

Streaming video = encoding + DASH + playout buffering

### Content Distribution Network

Store/serve multiple copies of videos at multiple geographically distributed sites

- Enter deep: push CDN servers deep into many access networks
- Bring home: smaller number of larger clusters in POPs near access nets


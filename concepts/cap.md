## The Three Properties

### **Consistency (C)**

- All nodes in the distributed system see the same data at the same time.
- Every read operation receives the most recent write or an error.
- When data is updated on one node, all other nodes must be updated before any read operations can proceed.
- Examples: Strong consistency in databases like traditional RDBMS with ACID properties.

### **Availability (A)**

- The system remains operational and responsive to requests at all times.
- Every request (read or write) receives a response, even if some nodes are down.
- The system continues to function despite failures in individual components.
- Examples: DNS systems, CDNs that serve cached content even when origin servers are unavailable.

### **Partition Tolerance (P)**

- The system continues to operate despite network failures or communication breakdowns between nodes.
- The system can handle arbitrary message loss or failure of part of the system.
- Network partitions are inevitable in distributed systems, so this property is typically considered mandatory.
- Examples: Systems that can split into multiple groups and continue operating independently.

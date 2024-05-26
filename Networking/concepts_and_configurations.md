# Switching
- Switch is a device (physical or virtual) that creates a network to connect two or more devices such as two computers, VMs, etc.

## Connection and communication within a single network
- To connect them to a switch - an interface is required on each host (physical or virutal).
- `ip link`: to see the interfaces that will be used to connect to the switch.
- Once, the ip address of the switch is found, it is used to connect the devices to the switch using:
    - `ip addr add <ip address> <ip link>`
In this case, messages and packets are sent from one host of the network to another host of the same network via switch.

## Connection and communication with another network.
Router - connects two or more separate networks.
- If it connects 2 networks, it gets 2 IP addresses assigned.
Assume devices A & B on network N1 and devices C & D on network N2. If B wants to pass a packet to C, how will B know where the router is on the network?
<t/>Done by configuring systems with a gateway or a route.
<t/> ### Configuration systems with a gateway/route


# Routing
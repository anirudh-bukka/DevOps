# Switching
- Switch is a device (physical or virtual) that creates a network to connect two or more devices such as two computers, VMs, etc.
- To connect them to a switch - an interface is required on each host (physical or virutal).
- `ip link`: to see the interfaces that will be used to connect to the switch.
- Once, the ip address of the switch is found, it is used to connect the devices to the switch using:
    - `ip addr add <ip address> <ip link>`

# Routing
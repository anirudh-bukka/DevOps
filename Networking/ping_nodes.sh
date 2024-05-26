# `jump host`` is able to access all four apps. But if you try to ping app03 or app04 from app01 or app02 or vice versa you will see ping is not working. So now we want to use jump host as a router so that app01 and app02 can access app03 and app04 and vice versa, lets add some routing table entries on these hosts to make it work.

# a. Add a routing table entry in app01 and app02 hosts so that these hosts can reach app03 and app04 hosts via jump host.

# b. Add a routing table entry in app03 and app04 hosts so that these hosts can reach app01 and app02 hosts via jump host.

# c. Now try to ping app03 and app04 from app01 and app02 and vice versa, every app should be able to ping each other.



# On app01 and app02:
ssh app01
sudo ip route add 172.16.239.0/24 via 172.16.238.10
exit

ssh app02
sudo ip route add 172.16.239.0/24 via 172.16.238.10
exit

# On app03 and app04:
ssh app03
sudo ip route add 172.16.238.0/24 via 172.16.239.10
exit

ssh app04
sudo ip route add 172.16.238.0/24 via 172.16.239.10
exit
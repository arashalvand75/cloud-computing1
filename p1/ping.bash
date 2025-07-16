#problem1
#ping
#by a switch case map input parameteres to corresponding ip 
#with netns ping the other ones

case $2 in
	node1) ip_dst=172.0.0.2;;
	node2) ip_dst=172.0.0.3;;
	node3) ip_dst=10.10.0.2;;
	node4) ip_dst=10.10.0.3;;
	router) ip_dst=172.0.0.1;;
esac

ip netns exec $1 ping $ip_dst


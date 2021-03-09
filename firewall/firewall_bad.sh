#!/bin/bash

BAN_IP=/tmp/ban_ip.txt

for j in $(cat ${BAN_IP}); do
    firewall-cmd --remove-source=${j} --zone=drop --permanent
    firewall-cmd --remove-source=${j} --zone=drop
done

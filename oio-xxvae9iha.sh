#!/bin/bash

# Get the public IP address
public_ip=$(curl -s ifconfig.me)

# Get the local IP address
local_ip=$(ip route get 8.8.8.8 | awk '{print $7}' | tr -d '\n')

echo "Public IP: $public_ip"
echo "Local IP: $local_ip"
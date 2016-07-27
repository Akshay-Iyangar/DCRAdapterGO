#!/bin/sh

# find the entry of the hostname

HOST_IP=$(curl http://rancher-metadata/2015-12-19/self/host/agent_ip)


url=$(echo "http://$HOST_IP:12285/v1/dc/logs/ecomm/logs")

/gopath/bin/app $url

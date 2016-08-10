#!/bin/sh

# find the entry of the hostname

HOST_IP=$(curl http://rancher-metadata/2015-12-19/self/host/agent_ip)

RESTurl="http://dashboard.ecom.int.godaddy.com/"

DCRurl=$(echo "http://$HOST_IP:12285/v1/dc/logs/ecomm/logs")

/root/dcradapter/dcr $RESTurl $DCRurl

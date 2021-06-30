#!/usr/bash

mkdir -p /home/moon-street/gospace
cd /home/moon-street/gospace

wget https://golang.org/dl/go1.16.5.linux-amd64.tar.gz

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz

go version


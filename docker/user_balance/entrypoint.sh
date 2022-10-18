#!/bin/bash

tar -C /usr/local -xzf /tmp/go1.18.6.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

cd /var/www/user_balance/
/usr/local/go/bin/go run service/cmd/service/main.go &

while true ; do sleep 5; done;

#!/bin/bash

# 启动service
run_service() {
    nohup ./etcd --listen-peer-urls  http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2380 --listen-client-urls http://0.0.0.0:2380 >> /data/etcd/run.log 2>&1  &
    sleep 1
    echo -e "\033[32m 启动成功: \033[0m"
}
cd /usr/etcd-v3.4.13-linux-amd64
run_service

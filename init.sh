#!/bin/bash
path=$(cd `dirname $0`; pwd)
mkdir -p $path/docker/data/edcd;
mkdir -p $path/docker/data/mysql;
mkdir -p $path/docker/data/mysql-files;
mkdir -p $path/docker/data/redis;
mkdir -p $path/docker/data/mongodb;
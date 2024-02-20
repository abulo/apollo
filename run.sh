#!/bin/bash
mkdir -pv /usr/local/openresty/nginx/logs;
mkdir -pv /var/log/supervisor;
/usr/local/openresty/nginx/sbin/nginx -c /usr/local/openresty/nginx/conf/nginx.conf -g "daemon off;"
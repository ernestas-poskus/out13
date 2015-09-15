#!/bin/sh
ip=$(ip addr show | grep inet  | tail -n 1 | cut -d ' ' -f 8)
hugo server --bind=$ip --port=1313 --baseUrl=$ip --watch

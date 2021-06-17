# Makefile features contains.
# 1)、Install commands program for features.
# 2)、Build commands program for features.
# 3)、Copy binary file into the docker container.
# 4)、Copy binary file into the linux server.
# Author: helloshaohua
# Date: 2021-06-13

# Install cli application with the features.
.PHONY: install
install:
	go install ./cmd/speedtestx

# Build cli application with the features.
.PHONY: build
build:
	/bin/bash ./build.sh

# Copy binary file into the docker container.
#
# Usage:
# 	make linux-copy CONTAINER_NAME=rocky
.PHONY: docker-copy
docker-copy:
	docker cp ./compile/speedtest.linux.amd64 $(CONTAINER_NAME):/usr/bin/speedtest

# Copy binary file into the linux server.
#
# Makefile use extras variables. Ref: https://www.codenong.com/2826029/
# Usage:
# 	make linux-copy LINUX_REMOTE_IP=192.168.1.100
.PHONY: linux-copy
linux-copy:
	scp ./compile/speedtest.linux.amd64 root@$(LINUX_REMOTE_IP):/usr/local/bin/speedtest
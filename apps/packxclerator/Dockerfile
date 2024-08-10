FROM ubuntu:22.04

# Install DPDK and build tools
RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y \
    build-essential \
    cmake \
    libdpdk-dev \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /apps/packxclerator
COPY . .

RUN mkdir build && cd build && cmake .. && make

CMD ["./build/packxclerator"]
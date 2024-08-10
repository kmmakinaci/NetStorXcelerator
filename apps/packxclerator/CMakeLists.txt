cmake_minimum_required(VERSION 3.10)
project(packxclerator)

set(CMAKE_CXX_STANDARD 17)

#find_package(DPDK REQUIRED)

add_executable(packxclerator main.cpp)
#target_link_libraries(packxclerator DPDK::DPDK)
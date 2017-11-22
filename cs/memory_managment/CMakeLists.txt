cmake_minimum_required(VERSION 3.7)
project(memory_managment)

set(CMAKE_C_STANDARD 99)

set(SOURCE_FILES main.c main.h)
add_executable(memory_managment ${SOURCE_FILES})
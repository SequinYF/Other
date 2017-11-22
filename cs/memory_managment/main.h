//
// Created by Sequin_YF on 2017/11/22.
//

#ifndef MEMORY_MANAGMENT_MAIN_H
#define MEMORY_MANAGMENT_MAIN_H

#include <stdio.h>
#include <stdlib.h>



#define PROCESS_NAME_LEN 32        /*进程名称的最大长度*/
#define MIN_SLICE    10            /*最小碎片的大小*/
#define DEFAULT_MEM_SIZE 1024       /*默认内存的大小*/
#define DEFAULT_MEM_START 0        /*默认内存的起始位置*/

/* 内存分配算法 */
#define MA_FF 1
#define MA_BF 2
#define MA_WF 3

int mem_size=DEFAULT_MEM_SIZE;        /*内存大小*/
int ma_algorithm = MA_FF;            /*当前分配算法*/
static int pid = 0;                    /*初始pid*/
int flag = 0;                        /*设置内存大小标志*/

struct free_block_type{
    int size;
    int start_addr;
    struct free_block_type *next;
};

struct free_block_type *free_block;

struct allocated_block{
    int pid;
    int size;
    int start_addr;
    char process_name[PROCESS_NAME_LEN];
    struct allocated_block *next;
};

struct allocated_block *allocated_block_head = NULL;



struct free_block_type* init_free_block(int mem_size);
void display_menu();
int set_mem_size();
void set_algorithm();
void rearrange(int algorithm);
int display_mem_usage();
int dispose(struct allocated_block *free_ab);
int free_mem(struct allocated_block *ab);
void kill_process();
int allocate_mem(struct allocated_block *ab);
int new_process();
rearrange_WF();
rearrange_BF();
void rearrange_FF();
void rearrange(int algorithm);
struct allocated_block *find_process(int pid);
void swap(int *pInt, int *pInt1);
#endif //MEMORY_MANAGMENT_MAIN_H

## GoPRM 

GoPRM is a process resource manager written by golang.

Created by coldriver, 2017.05.02.

 struct:
 - 就绪（ready）队列是关于PCB的单向链表PCBLinkList，就绪池（PCBPool）是关于就绪队列的一维三元数组
 - 每个资源（resource）中有该资源的阻塞（blocked）队列，是关于PCB的一维切片/数组
 - 资源池（RCBPool）是关于RCB的一维切片/数组
 - 进行（running）块，是关于正在运行的进程的数据结构
 - 完成（finish）队列是关于已完成PCB的一维切片/数组
 - 撤销（destroy）队列是关于已撤销PCB的一维切片/数组

 Init：
 - 系统初始化Init进程PCB，优先级设为0
 - 系统初始化四种资源R1，R2，R3，R4的四个RCB
 - 系统初始化PCB池（PCBPool），并将Init进程PCB块放入PCB池
 - 系统初始化RCB池，并将四个RCB放入RCB池

 Create:
 - 父进程可创建子进程
 - 除去Init进程，所有进程均有其父进程
 - 所有进程均可创建最多5个子进程
 - 子进程可以继承父进程的资源

 Request:
 - 已创建的进程有权request资源
 - 按顺序申请一次所请求的资源
 - 若某资源可用，将该资源的准备状态OK置为true，将该资源RCB的Available减1
 - 若某资源当前不可用，将该PCB插入该RCB的阻塞队列

 destory:
 - 销毁进程，其子进程也要被全部销毁
 - 释放该进程和其子进程的全部资源
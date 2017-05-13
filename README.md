## GoPRM 

GoPRM is a process resource manager written by golang.

Created by coldriver, 2017.05.02.

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
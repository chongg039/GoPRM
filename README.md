## GoPRM 

GoPRM is a process resource manager written by golang.

Created by coldriver, 2017.05.02.

 Init：
 - 创建Init进程PCB
 - 创建四种资源R1，R2，R3，R4的四个RCB
 - 创建PCB池（PCBPool），并将Init进程PCB块放入PCB池
 - 创建RCB池，并将四个RCB放入RCB池

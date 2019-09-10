

# Introduction

This is an adaptation of [GCP](https://github.com/gcp)'s [Leela Zero](https://github.com/gcp/leela-zero/) repository to chess, using Stockfish's position representation and move generation. (No heuristics or prior knowledge are carried over from Stockfish.)

The goal is to build a strong UCT chess AI following the same type of techniques as AlphaZero, as described in [Mastering Chess and Shogi by Self-Play with a General Reinforcement Learning Algorithm](https://arxiv.org/abs/1712.01815).

We will need to do this with a distributed project, as it requires a huge amount of computations.

# GGzero 项目官网地址：

http://www.ggzero.cn

# GGzero 项目论坛地址：

http://bbs.ggzero.cn

# GGzero 项目进展情况：

  自评 ELO 达到了3700 分，弈天桃花岛9星（大概1060的显卡）！ 
  
  ![](https://github.com/leedavid/leela-chess-to-Chinese-Chess/blob/master/lc0/20190904.png)
  
  GGzero 参加 第三届楚河汉界象棋人工智能第三名 
  
  机器配置： 
  
  GGzero    2070super X 4 无开局库，无残局库
  
  其它引擎   2496V4 44核心 有开局库，有残局库
  
  用时：10分+3秒 
  

# 谷歌 colab 跑谱共享笔记本地址

https://colab.research.google.com/drive/18yj5p1vF50_U4-X0iMf4eq-hzFWUg78z


只要用上面的colab jupyter笔记本，就能跑谱参加训练！

# 程序使用说明

  ## 注意：所有的文件都可到 GGzero 官方群号：779375937 的 "GGzero训练客户端" 文件夹下载

1. 安装visual studio 2019 运行库（在群共享GGzero训练客户端下载）。

2. 下载 “GGzeroALL.rar”，解压到英文目录下，注意整个目录不能有中文和空格！！

3. 下载最新的显卡驱动，可到 nvidia网站上下载，或到群共享 GGzero训练客户端 目录下载

4. 运行客户端训练程序 lc0_main.exe, 输入用户名和密码，就可以正常训练了。

5. 训练时，客户端下载的权重保存在 client-cache 目录下面，如果你想用 GGzero 下棋，请将 client-cache 目录下的最新权重文件复制到引擎文件lc0.exe目录下面。然后用支持UCI协议的象棋界面加载 lc0.exe， 就可以正常下棋了。

6. GGzere 目前暂时不开源，但会在合适的时间开源。 GGzero 团队承诺，公开训练的权重，永久给大家免费使用。

# 常见问题 Q/A:

1.  问：不能正常训练，怎么办？
    答：点击：lc0.exe, 在lc0.exe的界面里，输入 go nodes 800 回车，如果lc0.exe正常显示棋步了。就表明程序安装准确了。
    
2.  问：如何用兵河加载引擎？
    答：先要能正常训练，然后用兵河加载 lc0.exe.  

如您还有其它问题，请加入QQ 讨论群号：779375937 


# 帮助学习后有什么好处？
a. 首先感谢您对 	佳佳象棋 GGzero 团队的支持。
b. 你的学习记录将永久记录在网站上。当您达到一定的学习时间时，会以您提供的名字出现在引擎,及网站的赞助名单上。
c. 我们团队还在讨论其它的公平，公开的合适的奖励方案，一旦讨论通过，将会在第一时间公布并实施。

# 特别感谢以下棋友对 GGzero 项目所做出的贡献

 

1. 東檸 QQ: 1540899650
2. 痴情麦片 QQ: 81660127
3. Nooby QQ: 392620089
4. 冲锋 QQ: 272556328
5. 永宅依寍 QQ: 395466412
6. 扈 QQ: 549556859
7. 唐尸三摆手 QQ: 87103966
8. 灵玄子 QQ: 263140837
9. 王较瘦 QQ: 11559911
10. ❀小雪❀ QQ: 695044453
11. 木野狐 QQ: 122349668
12. 兵九进一 QQ: 29289180
13. 沙漠雄鹰 QQ: 529213121
14. 剑 QQ: 81060870
15. netsea2018
16. 卧薪尝胆 QQ: 40305749
17. 此生又何劳
18. 青雨菲扬 QQ: 63479139
19. gakki QQ: 450673130 香港
20. 竹海听风 QQ: 363036507
21. 小兵过河 QQ: 270909139
22. 飞鱼
23. 神说要有光 QQ: 472495748
24. cai1975
25. apollo
26. 若愚（lxc）QQ: 1176222948
27. 神龟 QQ: 1033008052
28. 小呆呆 QQ: 1067806890
29. 天涯 QQ: 38819386
30. 晴天 QQ: 1053558212



## 以上排名不分先后，名称多数是网络昵称

# QQ 讨论群号：779375937
# 官网： http://www.ggzero.cn

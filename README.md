[![Linux Build Status](https://travis-ci.org/glinscott/leela-chess.svg?branch=master)](https://travis-ci.org/glinscott/leela-chess)
[![Windows Build status](https://ci.appveyor.com/api/projects/status/w2nymx3wpd0d1da1/branch/master?svg=true)](https://ci.appveyor.com/project/glinscott/leela-chess/branch/master)

# Introduction

This is an adaptation of [GCP](https://github.com/gcp)'s [Leela Zero](https://github.com/gcp/leela-zero/) repository to chess, using Stockfish's position representation and move generation. (No heuristics or prior knowledge are carried over from Stockfish.)

The goal is to build a strong UCT chess AI following the same type of techniques as AlphaZero, as described in [Mastering Chess and Shogi by Self-Play with a General Reinforcement Learning Algorithm](https://arxiv.org/abs/1712.01815).

We will need to do this with a distributed project, as it requires a huge amount of computations.

Please visit the LCZero forum to discuss: https://groups.google.com/forum/#!forum/lczero, or the github issues.

# Contributing

For precompiled binaries, see:
* [wiki](https://github.com/glinscott/leela-chess/wiki)
* [wiki/Getting-Started](https://github.com/glinscott/leela-chess/wiki/Getting-Started)

For live status: http://lczero.org

The rest of this page is for users who want to compile the code themselves.
Of course, we also appreciate code reviews, pull requests and Windows testers!

程序使用说明
下载软件后，解压，如果目录内有设置文件，settings.json，则要删除掉。
安装visual studio 2017 运行库，下载地址：https://github.com/leedavid/leela-chess-to-Chinese-Chess/releases/download/v0.71/VC_redist.x64.rar
如果机器有多个显卡，就运行mainMultGpu.exe，如果机器只有一个显卡，就运行mainClient.exe
根据提示输入 Enter username:
用户名（请输入英文用户） Enter password:
口令 也是英文 Enter train num per time befor upload:
一般输入数字 5， 表示训练5局才上传到服务器 Use GpU？1 use or 0 not use : 如果有支持CUDA运算的显卡，就输入1，如果显卡比较老，就输入 0 How many Gpu do u want to use: 如果你想用几块显卡一起运算，就输入数字几，只有一块显卡，就输入1
QQ 讨论群号：779375937，网站：http://www.lcchess.com/
加群：点击链接加入群聊【LCCZero】：https://jq.qq.com/?_wv=1027&k=5Epwv5H
文件说明： GGengineCPU.exe CPU版本的 UCI引擎，可用兵河等界面程序加载 GGengineGPU.exe GPU版本的 UCI引擎，可用兵河等界面程序加载 mainClient.exe 不支持多个显卡的训练程序 mainMultiGpu.exe 支持多个显卡的训练程序 weights.txt 这个是UCI 引擎下棋时会调用的缺省权重文件， 可到 http://www.lcchess.com/networks下载最新的权重文件，并改名成 weights.txt. networks目录，这个目录在界面训练时，会自动从网站上下载最新权重文件去训练，但有时由于网络问题，会出现文件没有下载完成的情况，造成引擎闪退，那么我们可以到上面的网站去下载对应的权重文件，放到这个目录下。主要是看一下权重文件的大小，如果太少，肯定是没有下载完整。
常见问题 A/Q
1.	可不可以多开？
答：可以，如果你的电脑运算能力足够，可以多开，但要看一下任务管理器，最好留一些余量。

2.	训练程序闪退怎么？
a.	系统最好是win10, win7没有测试过 b.	要安装 visual studio 2017 运行库, 然后重启 c.	因为引擎使用了 SSE4.2指令，太旧的CPU可能也不支持 d.	检查引擎目录下没有 weigths.txt文件，如果没有，可到群里下载 e.	界面在训练时，需要调用相应的权利文件，请见上述的 networks目录的说明 因为引擎使用CUDA指令，所以使用GPU显卡运行时，最好要安装最新的显卡驱动。N卡驱动下载地址：https://www.geforce.cn/drivers

3. 核显和独显，不算多个显卡吧？
是的

4. 帮助学习后有什么好处？
a. 首先感谢您对 中国象棋 alphazero 团队的支持。 b. 你的学习记录将永久记录在网站上。当您达到一定的学习时间时，会以您提供的名字出现在引擎,及网站的赞助名单上。 c. 我们团队还在讨论其它的公平，公开的合适的奖励方案，一旦讨论通过，将会在第一时间公布并实施。

QQ 讨论群号：779375937，网站：http://www.lcchess.com

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

http://lcc.ya.cn
http://www.ggzero.cn

# 程序使用说明


1. 安装visual studio 2017 运行库，下载地址：https://github.com/leedavid/leela-chess-to-Chinese-Chess/releases/download/v0.71/VC_redist.x64.rar

2. 将下载的训练程序解压到英文目录下，注意，英文路径不能有空格，否则不能正常上传训练数据。
   解压后共二个程序文件。lc0_main.exe 训练程序， lc0.exe 是兼容 UCI 协议的引擎文件。

3. 如果系统没有安装cuda9.2运行库，则需要将三个cuda dll文件放到训练程序目录下，请注意，dll对应的不同系统，分win7,和win10版本。
   程序目录下，应该共有5个文件。3个dll,2个exe文件。  

4. 点击 lc0.exe, 在lc0.exe的界面里，输入 go nodes 800 回车，如果lc0.exe正常显示棋步了。就表明程序安装准确了。
5. 运行 lc0_main.exe，根据提示，输入您想注册的用户和密码，就可以正常训练棋谱，自动上传训练数据了。

如您还有其它问题，请加入QQ 讨论群号：779375937 


## 4. 帮助学习后有什么好处？
a. 首先感谢您对 	佳佳象棋 GGzero 团队的支持。
b. 你的学习记录将永久记录在网站上。当您达到一定的学习时间时，会以您提供的名字出现在引擎,及网站的赞助名单上。
c. 我们团队还在讨论其它的公平，公开的合适的奖励方案，一旦讨论通过，将会在第一时间公布并实施。


QQ 讨论群号：779375937，网站：http://lcc.ya.cn http://www.ggzero.cn

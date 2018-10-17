#!/bin/bash
apt-get install protobuf-compiler
#wget http://developer.download.nvidia.com/compute/redist/cudnn/v7.1.4/cudnn-9.0-linux-x64-v7.1.tgz
#tar -xzvf cudnn-9.0-linux-x64-v7.1.tgz
#cp -P cuda/include/cudnn.h /usr/local/cuda/include/cudnn.h
#cp -P cuda/lib64/libcudnn* /usr/local/cuda/lib64/
#chmod a+r /usr/local/cuda/lib64/libcudnn*
wget -c https://github.com/leedavid/leela-chess-to-Chinese-Chess/raw/master/lc0/ggzero_linux
wget -c https://github.com/leedavid/leela-chess-to-Chinese-Chess/raw/master/lc0/lc0
chmod 777 lc0
chmod 777 ggzero_linux
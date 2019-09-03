#!/bin/bash
# apt-get install protobuf-compiler 20190904
#wget http://developer.download.nvidia.com/compute/redist/cudnn/v7.1.4/cudnn-9.0-linux-x64-v7.1.tgz
#tar -xzvf cudnn-9.0-linux-x64-v7.1.tgz
#cp -P cuda/include/cudnn.h /usr/local/cuda/include/cudnn.h
#cp -P cuda/lib64/libcudnn* /usr/local/cuda/lib64/
#chmod a+r /usr/local/cuda/lib64/libcudnn*

#apt-get update  20190904
#apt-get install -y --fix-missing --no-install-recommends cuda-compiler-10-0 cuda-cublas-dev-10-0 cuda-cudart-dev-10-0
# apt-get install -y --fix-missing --no-install-recommends nvidia-opencl-dev libopenblas-dev 20190904

#wget http://developer.download.nvidia.com/compute/redist/cudnn/v7.3.1/cudnn-10.0-linux-x64-v7.3.1.20.tgz
#cd /usr/local && tar -xzvf /content/cudnn-10.0-linux-x64-v7.3.1.20.tgz
#chmod a+r /usr/local/cuda/lib64/libcudnn*

cd /content
rm -f ggzero_linux
rm -f lc0

wget -c https://github.com/leedavid/leela-chess-to-Chinese-Chess/raw/master/lc0/ggzero_linux
wget -c https://github.com/leedavid/leela-chess-to-Chinese-Chess/raw/master/lc0/lc0
chmod 777 lc0
chmod 777 ggzero_linux




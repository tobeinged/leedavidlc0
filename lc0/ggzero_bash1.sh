#!/bin/bash

rm -rf /etc/apt/sources.list.d
apt-get install protobuf-compiler
wget http://developer.download.nvidia.com/compute/cuda/repos/ubuntu1704/x86_64/cuda-repo-ubuntu1704_9.0.176-1_amd64.deb
apt-get install -y --fix-missing --no-install-recommends dirmngr
dpkg -i cuda-repo-ubuntu1704_9.0.176-1_amd64.deb
apt-key adv --fetch-keys https://developer.download.nvidia.com/compute/cuda/repos/ubuntu1704/x86_64/7fa2af80.pub
apt-get update
mkdir /usr/lib/nvidia
apt-get install -y --fix-missing --no-install-recommends nvidia-384=384.111-0ubuntu1 libcuda1-384=384.111-0ubuntu1 nvidia-384-dev=384.111-0ubuntu1
apt-get install -y --fix-missing --no-install-recommends cuda-cudart-9-0 cuda-cublas-9-0  cuda-core-9-0  cuda-cublas-dev-9-0 cuda-cudart-dev-9-0
cd /usr/local/ &&  ln -s cuda-9.0 cuda
[ -f /usr/local/cuda/bin/nvcc ] && echo Using CUDA.
[ ! -f /usr/local/cuda/bin/nvcc ] && apt-get install -y --fix-missing --no-install-recommends nvidia-opencl-dev nvidia-opencl-icd-384=384.111-0ubuntu1 opencl-headers libopenblas-dev



#!/bin/bash

TEXT_RESET='\e[0m'
TEXT_YELLOW='\e[1;33m'

# https://developer.nvidia.com/compute/cuda/9.2/Prod2/local_installers/cuda-repo-ubuntu1604-9-2-local_9.2.148-1_amd64
wget https://developer.nvidia.com/compute/cuda/9.2/Prod2/local_installers/cuda-repo-ubuntu1604-9-2-local_9.2.148-1_amd64
echo -e $TEXT_YELLOW
echo 'WEBGET finished..'
echo -e $TEXT_RESET

dpkg --install cuda-repo-ubuntu1604-9-2-local_9.2.148-1_amd64
echo -e $TEXT_YELLOW
echo 'DPKG finished..'
echo -e $TEXT_RESET

apt-key add /var/cuda-repo-9-2-local/7fa2af80.pub
echo -e $TEXT_YELLOW
echo 'APT added key..'
echo -e $TEXT_RESET

apt-get update
echo -e $TEXT_YELLOW
echo 'APT update finished..'
echo -e $TEXT_RESET

apt-get install cuda
echo -e $TEXT_YELLOW
echo 'APT finished installing cuda..'

echo 'The CUDA version is: '
cat /usr/local/cuda/version.txt
echo -e $TEXT_RESET

# install cudnn runtime
# https://developer.nvidia.com/compute/machine-learning/cudnn/secure/v7.2.1/prod/9.2_20180806/Ubuntu16_04-x64/libcudnn7_7.2.1.38-1_cuda9.2_amd64
wget https://developer.nvidia.com/compute/machine-learning/cudnn/secure/v7.2.1/prod/9.2_20180806/Ubuntu16_04-x64/libcudnn7_7.2.1.38-1_cuda9.2_amd64
echo -e $TEXT_YELLOW
echo 'WEBGET cudnn runtime finished..'
echo -e $TEXT_RESET

dpkg --install libcudnn7_7.2.1.38-1_cuda9.2_amd64
echo -e $TEXT_YELLOW
echo 'DPKG cudnn runtime finished..'
echo -e $TEXT_RESET


# https://developer.nvidia.com/compute/machine-learning/cudnn/secure/v7.2.1/prod/9.2_20180806/Ubuntu16_04-x64/libcudnn7-dev_7.2.1.38-1_cuda9.2_amd64
#install cudnn devloper
wget https://developer.nvidia.com/compute/machine-learning/cudnn/secure/v7.2.1/prod/9.2_20180806/Ubuntu16_04-x64/libcudnn7-dev_7.2.1.38-1_cuda9.2_amd64
echo -e $TEXT_YELLOW
echo 'WEBGET cudnn devloper finished..'
echo -e $TEXT_RESET

dpkg --install Ubuntu16_04-x64/libcudnn7-dev_7.2.1.38-1_cuda9.2_amd64
echo -e $TEXT_YELLOW
echo 'DPKG cudnn devloper finished..'
echo -e $TEXT_RESET



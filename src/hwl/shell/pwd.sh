#!/bin/sh

#两种方法
script_path=$(cd $(dirname $0) && pwd)
script_path=$(dirname $(readlink -f $0))

echo $script_path  #无论在哪里执行脚本都能打印当前脚本路径

echo $PWD   #打印执行脚本时候所在的路径
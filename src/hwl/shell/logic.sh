#!/bin/sh

# -eq 是否相等
# if的中括号前后必须加空格 ]后面有分号再接着then。 then也可以另起一行，缩进不影响
if [ 1 -eq 1 ]; then
    echo "1 is equal 1"
fi

# -ne 不相等
if [ 1 -ne 2 ]; then
    echo "1 is not equal 2"
fi

# -a 并且
if [ 2 -eq 2 -a 1 -eq 1 ]; then
    echo "2 equal 2 and 1 equal 1"
fi

# -o 或
if [ 1 -eq 2 -o 1 -eq 1 ]; then
    echo "1 eq 2 or 1 eq 1"
fi
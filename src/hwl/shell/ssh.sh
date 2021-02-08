#!/usr/bin/expect


spawn ssh hwl@192.168.117.240
expect {
    "*yes/no" { send "yes\r"; exp_continue }
    "*password:" { send "l\r" }
}


interact

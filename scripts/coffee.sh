#!/bin/bash
#这里可替换为你自己的执行程序，其他代码无需更改
APP_NAME=coffee-cli

#使用说明，用来提示输入参数
usage() {
 echo "Usage: sh 脚本名.sh [start|stop|restart|status]"
 exit 1
}
  
#检查程序是否在运行
is_exist(){
 pid=`ps -ef|grep $APP_NAME|grep -v grep|awk '{print $2}' `
 #如果不存在返回1，存在返回0 
 if [ -z "${pid}" ]; then
 return 1
 else
 return 0
 fi
}
  
#启动方法
start(){
 is_exist
 if [ $? -eq "0" ]; then
 echo "${APP_NAME} is already running. pid=${pid} ."
 else
  nohup ./${APP_NAME} > ./logs/coffee.log 2>&1 &
 echo "${APP_NAME} start success"
 fi
}
  
#停止方法
stop(){
 is_exist
 if [ $? -eq "0" ]; then
 kill $pid
 else
 echo "${APP_NAME} is not running"
 fi
}
  
#输出运行状态
status(){
 is_exist
 if [ $? -eq "0" ]; then
 echo "${APP_NAME} is running. Pid is ${pid}"
 else
 echo "${APP_NAME} is NOT running."
 fi
}

log(){
 tail -f logs/coffee.log
}

build(){
  version=""
  if [ -f "VERSION" ]; then
      version=`cat VERSION`
  fi

  if [[ -z $version ]]; then
      if [ -d ".git" ]; then
          version=`git symbolic-ref HEAD | cut -b 12-`-`git rev-parse HEAD`
      else
          version="unknown"
      fi
  fi
  echo "$version"
  go build -ldflags "-X main.Version=$version" -o coffee-cli
}
  
#重启
restart(){
 stop
 start
}
  
#根据输入参数，选择执行对应方法，不输入则执行使用说明
case "$1" in
 "start")
 start
 ;;
 "stop")
 stop
 ;;
 "status")
 status
 ;;
 "restart")
 restart
 ;;
 "log")
 log
 ;;
 "build")
 build
 ;;
 *)
 usage
 ;;
esac


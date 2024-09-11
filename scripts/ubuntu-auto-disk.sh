#!/bin/bash

#使用说明，用来提示输入参数
usage() {
 echo "Usage: bash auto-disk.sh [batchMount]"
 exit 1
}

batchMount(){
  batchSuccess=true
  for disk in $(parted -l |grep 'Disk /dev/sd' |grep Disk |grep TB |awk '{split($2,s,":"); print s[1]}')
  do
    dirName=${disk:5}
    #4.对磁盘/dev/sd*进行分区
    parted -s "$disk" mklabel gpt
    #5.对磁盘/dev/sd* 指定分区类型和容量占比
    parted -s "$disk" mkpart primary 1 100%
    #6.格式化磁盘/dev/sd*
    mkfs.xfs -f "${disk}"1
    echo "/n/n****************$disk parted was Finished! Waiting For 2 second****/n/n"
    sleep 2s
    #7.创建对应磁盘个数的目录，/hadoop*,创建挂载点
    mkdir -p /mnt/"${dirName}"
    #9.通过blk id命令查看磁盘的uuid，获取uuid值
    uuid=$(blkid "${disk}"1 |awk '{print $2}' |sed s#\"##g)
     if [ -z "${uuid}" ]; then
       batchSuccess=false
       echo "$disk uuid not found, batch mount fail!"
       break
     fi
    #10.设置开机自动挂载磁盘，追加uuid信息到 /etc/fstab中。
    echo "$uuid     /mnt/${dirName}   xfs     noatime,nodiratime 0       0">>/etc/fstab
  done

  if [ "$batchSuccess" == true ]; then
    mount -a
    echo "Batch mount success!"
  else
    echo "Batch mount fail!"
  fi

}

#根据输入参数，选择执行对应方法，不输入则执行使用说明
case "$1" in
 "batchMount")
 batchMount
 ;;
 *)
 usage
 ;;
esac
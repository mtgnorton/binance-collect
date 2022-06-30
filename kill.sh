#! /bin/bash
# collect-monitor.sh
process="binance-collect"
pid=$(ps -ef | grep $process | grep -v grep | awk '{print $2}')
if [ $pid ]; then
  echo "该进程已存在,已存在进程的pid如下,将杀死该进程"
  echo $pid
  kill -9 $pid
fi

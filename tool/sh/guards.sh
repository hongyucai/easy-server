#! /bin/sh
PRO_NAME=job
while true ; do
    #用ps获取$PRO_NAME进程数量
    NUM=`ps aux | grep ${PRO_NAME} | grep -v grep |wc -l`
    echo "${PRO_NAME} num ${NUM} ..."
    #少于1，重启进程
    if [ "${NUM}" -lt "1" ];then
        echo "${PRO_NAME} reload ..."
        nohup node /data/web/go-xm/build/job &
        #启动后沉睡10s
        sleep 10
    #大于1，杀掉所有进程，重启
    elif [ "${NUM}" -gt "1" ];then
        echo "${PRO_NAME} killall ..."
        killall -9 $PRO_NAME
    fi
    #kill僵尸进程
    NUM_STAT=`ps aux | grep ${PRO_NAME} | grep T | grep -v grep | wc -l`
    if [ "${NUM_STAT}" -gt "0" ];then
        echo "${PRO_NAME} killall zombie ..."
        killall -9 ${PRO_NAME}
    fi
    sleep 10
done
exit 0
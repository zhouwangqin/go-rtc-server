#!/bin/bash

APP_DIR=$(cd `dirname $0`/../; pwd)
echo "$APP_DIR"
cd $APP_DIR

mkdir -p $APP_DIR/logs

BIZ=biz
ISLB=islb
SFU=sfu

BIZ_CFG=$APP_DIR/configs/biz.toml
ISLB_CFG=$APP_DIR/configs/islb.toml
SFU_CFG=$APP_DIR/configs/sfu.toml

BIZ_LOG=$APP_DIR/logs/$BIZ.log
ISLB_LOG=$APP_DIR/logs/$ISLB.log
SFU_LOG=$APP_DIR/logs/$SFU.log

BUILD_PATH1=$APP_DIR/bin/$BIZ
BUILD_PATH2=$APP_DIR/bin/$ISLB
BUILD_PATH3=$APP_DIR/bin/$SFU

# echo "------------------start $BIZ------------------"
# echo "nohup $BUILD_PATH1 -c $BIZ_CFG >>$BIZ_LOG 2>&1 &"
# nohup $BUILD_PATH1 -c $BIZ_CFG >>$BIZ_LOG 2>&1 &
# sleep 1s
# echo "------------------start $ISLB------------------"
# echo "nohup $BUILD_PATH2 -c $ISLB_CFG >>$ISLB_LOG 2>&1 &"
# nohup $BUILD_PATH2 -c $ISLB_CFG >>$ISLB_LOG 2>&1 &
# sleep 1s
echo "------------------start $SFU------------------"
echo "nohup $BUILD_PATH3 -c $SFU_CFG >>$SFU_LOG 2>&1 &"
nohup $BUILD_PATH3 -c $SFU_CFG >>$SFU_LOG 2>&1 &
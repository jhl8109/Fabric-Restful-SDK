#!/bin/bash

start=`date +%s.%N`

# run something...
echo 'package chaincode'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS

peer lifecycle chaincode package simplecc.tar.gz --path ./chaincode/ --label simplecc

finish=`date +%s.%N`
diff=$( echo "$finish - $start" | bc -l )
echo '실행시간 : ' $diff


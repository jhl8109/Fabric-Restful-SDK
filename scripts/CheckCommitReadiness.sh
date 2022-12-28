#!/bin/bash

start=`date +%s.%N`

# run something...
echo 'checkcommitreadiness'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name simplecc --version 1.0 --sequence 1  --output json


finish=`date +%s.%N`
diff=$( echo "$finish - $start" | bc -l )
echo '실행시간 : ' $diff


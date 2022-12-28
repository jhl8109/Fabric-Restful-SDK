#!/bin/bash

start=`date +%s.%N`

# run something...
echo 'install chaincode'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode install simplecc.tar.gz

#peer 1
export CORE_PEER_ADDRESS=peer1.org1.example.com:9051
export CORE_PEER_TLS_ROOTCERT_FILE=/home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt

echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode install simplecc.tar.gz

finish=`date +%s.%N`
diff=$( echo "$finish - $start" | bc -l )
echo '실행시간 : ' $diff


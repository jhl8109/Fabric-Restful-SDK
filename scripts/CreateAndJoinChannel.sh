#!/bin/bash

start=`date +%s.%N`

# run something...
export FABRIC_CFG_PATH=/home/jeho/go/src/fabric-sdk-go-modulization/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=/home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/
export ORDERER_CA=/home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

echo 'create channel'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
#peer 0
export CORE_PEER_ADDRESS=peer0.org1.example.com:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

peer channel create -o orderer.example.com:7050 -c mychannel -f ./fixtures/channel-artifacts/channel.tx --tls --cafile /home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

echo 'join channel'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer channel join -b mychannel.block

#peer 1
export CORE_PEER_ADDRESS=peer1.org1.example.com:9051
export CORE_PEER_TLS_ROOTCERT_FILE=/home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt

echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer channel join -b mychannel.block

finish=`date +%s.%N`
diff=$( echo "$finish - $start" | bc -l )
echo '실행시간 : ' $diff


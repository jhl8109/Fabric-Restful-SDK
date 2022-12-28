#!/bin/bash

start=`date +%s.%N`

# run something...
# package cc
echo 'package chaincode'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode package simplecc.tar.gz --path ./chaincode/ --label simplecc
# install cc
echo 'install chaincode'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode install simplecc.tar.gz

#peer1
export CORE_PEER_ADDRESS=peer1.org1.example.com:9051
export CORE_PEER_TLS_ROOTCERT_FILE=/home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt

#peer1
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode install simplecc.tar.gz


# queryinstalled

#peer1
echo 'query installed chaincode'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode queryinstalled

#peer0
export CORE_PEER_ADDRESS=peer0.org1.example.com:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

#peer0
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode queryinstalled

# approve org
echo 'approve org'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode approveformyorg -o orderer.example.com:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name simplecc --version 1.0 --package-id simplecc:6e8003394d3dc351f158953b406d0c41223b99eb879d476d1ecd0fe591ce2f53 --sequence 1 --tls --cafile "${PWD}/fixtures/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"

# check approve
echo 'check approveformyorg'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode checkcommitreadiness --channelID mychannel --name simplecc --version 1.0 --sequence 1  --output json

# commit cc
#peer0
echo 'commit chaincode'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode commit -o orderer.example.com:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name simplecc --version 1.0 --sequence 1 --tls --cafile /home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles "/home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses peer1.org1.example.com:9051 --tlsRootCertFiles "/home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt"

#peer1
export CORE_PEER_ADDRESS=peer1.org1.example.com:9051
export CORE_PEER_TLS_ROOTCERT_FILE=/home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt

# querycommitted cc
echo 'query committed chaincode'
#peer1
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode querycommitted -C mychannel

#peer0
export CORE_PEER_ADDRESS=peer0.org1.example.com:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

#peer0
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode querycommitted -C mychannel

finish=`date +%s.%N`
diff=$( echo "$finish - $start" | bc -l )
echo '실행시간 : ' $diff


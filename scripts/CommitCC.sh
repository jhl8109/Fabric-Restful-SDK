#!/bin/bash

start=`date +%s.%N`

# run something...
echo 'commit chaincode'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode commit -o orderer.example.com:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name simplecc --version 1.0 --sequence 1 --init-required --tls --cafile /home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer1.org1.example.com:9051 --tlsRootCertFiles /home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt
finish=`date +%s.%N`
diff=$( echo "$finish - $start" | bc -l )
echo '실행시간 : ' $diff


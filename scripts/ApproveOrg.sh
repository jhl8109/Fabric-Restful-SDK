#!/bin/bash

start=`date +%s.%N`

# run something...
echo 'approve org'
echo 'CORE_PEER_ADDRESS : ' $CORE_PEER_ADDRESS
peer lifecycle chaincode approveformyorg -o orderer.example.com:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name simplecc --version 1.0 --package-id simplecc:6e8003394d3dc351f158953b406d0c41223b99eb879d476d1ecd0fe591ce2f53 --sequence 1 --tls --cafile /home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

finish=`date +%s.%N`
diff=$( echo "$finish - $start" | bc -l )
echo '실행시간 : ' $diff


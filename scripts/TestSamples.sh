n/bash

start=`date +%s.%N`

# run something...
peer chaincode invoke -o peer0.org1.example.com:7050 --ordererTLSHostnameOverride orderer.example.com --tls true --cafile $ORDERER_CA -C mychannel -n simplecc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"Args":["set","ID1","123"]}'
peer chaincode invoke -o peer0.org1.example.com:7050 --ordererTLSHostnameOverride orderer.example.com --tls true --cafile $ORDERER_CA -C mychannel -n simplecc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"Args":["set","ID2","456"]}'
peer chaincode invoke -o peer0.org1.example.com:7050 --ordererTLSHostnameOverride orderer.example.com --tls true --cafile $ORDERER_CA -C mychannel -n simplecc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /home/jeho/go/src/fabric-sdk-go-modulization/fixtures/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt -c '{"Args":["set","ID3","789"]}'

echo 'wait for 3 seconds...'
sleep 3
peer chaincode query -C mychannel -n simplecc -c '{"Args":["get","ID1"]}'
peer chaincode query -C mychannel -n simplecc -c '{"Args":["get","ID2"]}'
peer chaincode query -C mychannel -n simplecc -c '{"Args":["get","ID3"]}'

finish=`date +%s.%N`
diff=$( echo "$finish - $start" | bc -l )
echo '실행시간 : ' $diff


export CORE_PEER_TLS_ENABLED=true
export ORDERER_CA=${PWD}/../../artifacts/channel/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

export PEER0_ORG1_CA=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/fiber/peers/peer0.fiber/tls/ca.crt
export PEER1_ORG1_CA=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/fiber/peers/peer1.fiber/tls/ca.crt

export PEER0_ORG2_CA=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/yarn/peers/peer0.yarn/tls/ca.crt
export PEER1_ORG2_CA=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/yarn/peers/peer1.yarn/tls/ca.crt

export PEER0_ORG3_CA=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/fabric/peers/peer0.fabric/tls/ca.crt
export PEER1_ORG3_CA=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/fabric/peers/peer1.fabric/tls/ca.crt

export PEER0_ORG4_CA=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/apparel/peers/peer0.apparel/tls/ca.crt
export PEER1_ORG4_CA=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/apparel/peers/peer1.apparel/tls/ca.crt

export PEER0_ORG5_CA=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/retailer/peers/peer0.retailer/tls/ca.crt
export PEER1_ORG5_CA=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/retailer/peers/peer1.retailer/tls/ca.crt

# export PEER0_ORG6_CA=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/retailor/peers/peer0.retailor/tls/ca.crt
# export PEER1_ORG6_CA=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/retailor/peers/peer1.retailor/tls/ca.crt

export FABRIC_CFG_PATH=${PWD}/../../artifacts/channel/config/

export CHANNEL_NAME=mychannel

setGlobalsForPeer0Org1() {
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/fiber/users/Admin@fiber/msp
    export CORE_PEER_ADDRESS=localhost:7051
}

setGlobalsForPeer1Org1() {
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER1_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/fiber/users/Admin@fiber/msp
    export CORE_PEER_ADDRESS=localhost:8051
}

setGlobalsForPeer0Org2() {
    export CORE_PEER_LOCALMSPID="Org2MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/yarn/users/Admin@yarn/msp
    export CORE_PEER_ADDRESS=localhost:9051
    export ANCHOR_NODE="yarnMSP"
}

setGlobalsForPeer1Org2() {
    export CORE_PEER_LOCALMSPID="Org2MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/yarn/users/Admin@yarn/msp
    export CORE_PEER_ADDRESS=localhost:10051
}

setGlobalsForPeer0Org3() {
    export CORE_PEER_LOCALMSPID="Org3MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/fabric/users/Admin@fabric/msp
    export CORE_PEER_ADDRESS=localhost:7061
    export ANCHOR_NODE="fabricMSP"
}

setGlobalsForPeer1Org3() {
    export CORE_PEER_LOCALMSPID="Org3MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/fabric/users/Admin@fabric/msp
    export CORE_PEER_ADDRESS=localhost:7071
}

setGlobalsForPeer0Org4() {
    export CORE_PEER_LOCALMSPID="Org4MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG4_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/apparel/users/Admin@apparel/msp
    export CORE_PEER_ADDRESS=localhost:7081
    export ANCHOR_NODE="apparelMSP"
}

setGlobalsForPeer1Org4() {
    export CORE_PEER_LOCALMSPID="Org4MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG4_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/apparel/users/Admin@apparel/msp
    export CORE_PEER_ADDRESS=localhost:7091
}

setGlobalsForPeer0Org5() {
    export CORE_PEER_LOCALMSPID="Org5MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG5_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/retailer/users/Admin@retailer/msp
    export CORE_PEER_ADDRESS=localhost:8001
    export ANCHOR_NODE="retailerMSP"
}

setGlobalsForPeer1Org5() {
    export CORE_PEER_LOCALMSPID="Org5MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG5_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/../../artifacts/channel/crypto-config/peerOrganizations/retailer/users/Admin@retailer/msp
    export CORE_PEER_ADDRESS=localhost:8011
}

presetup() {
    echo Vendoring Go dependencies ...
    pushd ./../../artifacts/src/github.com/fabcar/go
    GO111MODULE=on go mod vendor
    popd
    echo Finished vendoring Go dependencies
}
# presetup

CHANNEL_NAME="mychannel"
CC_RUNTIME_LANGUAGE="golang"
VERSION="1"
CC_SRC_PATH="./../../artifacts/src/github.com/fabcar/go"
CC_NAME="fabcar"

packageChaincode() {
    rm -rf ${CC_NAME}.tar.gz
    setGlobalsForPeer0Org1
    peer lifecycle chaincode package ${CC_NAME}.tar.gz \
        --path ${CC_SRC_PATH} --lang ${CC_RUNTIME_LANGUAGE} \
        --label ${CC_NAME}_${VERSION}
    echo "===================== Chaincode is packaged on peer0.org1 ===================== "

    setGlobalsForPeer1Org1
    peer lifecycle chaincode package ${CC_NAME}.tar.gz \
        --path ${CC_SRC_PATH} --lang ${CC_RUNTIME_LANGUAGE} \
        --label ${CC_NAME}_${VERSION}
    echo "===================== Chaincode is packaged on peer1.org1 ===================== "
}
# packageChaincode

installChaincode() {
    setGlobalsForPeer0Org1
    peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer0.org1 ===================== "

    setGlobalsForPeer1Org1
    peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer1.org1 ===================== "
}
# installChaincode

queryInstalled() {
    setGlobalsForPeer0Org1
    peer lifecycle chaincode queryinstalled >&log.txt
    cat log.txt
    PACKAGE_ID=$(sed -n "/${CC_NAME}_${VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    echo PackageID is ${PACKAGE_ID}
    echo "===================== Query installed successful on peer0.org1 on channel ===================== "

    setGlobalsForPeer1Org1
    peer lifecycle chaincode queryinstalled
    echo "===================== Query installed successful on peer1.org1 on channel ===================== "
}
# queryInstalled

approveForMyOrg1() {
    setGlobalsForPeer0Org1
    # set -x
    # Replace localhost with your orderer's vm IP address
    peer lifecycle chaincode approveformyorg -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com --tls \
        --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name ${CC_NAME} --version ${VERSION} \
        --init-required --package-id ${PACKAGE_ID} \
        --sequence ${VERSION}
    # set +x

    echo "===================== chaincode approved from org 1 ===================== "
}
# queryInstalled
# approveForMyOrg1

checkCommitReadyness() {
    setGlobalsForPeer0Org1
    peer lifecycle chaincode checkcommitreadiness \
        --channelID $CHANNEL_NAME --name ${CC_NAME} --version ${VERSION} \
        --sequence ${VERSION} --output json --init-required
    echo "===================== checking commit readyness from org 1 ===================== "
}
# checkCommitReadyness

commitChaincodeDefination() {
    setGlobalsForPeer0Org1
    peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        --channelID $CHANNEL_NAME --name ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:8051 --tlsRootCertFiles $PEER1_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:10051 --tlsRootCertFiles $PEER1_ORG2_CA \
        --peerAddresses localhost:7061 --tlsRootCertFiles $PEER0_ORG3_CA \
        --peerAddresses localhost:7071 --tlsRootCertFiles $PEER1_ORG3_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        --peerAddresses localhost:7091 --tlsRootCertFiles $PEER1_ORG4_CA \
        --peerAddresses localhost:8001 --tlsRootCertFiles $PEER0_ORG5_CA \
        --peerAddresses localhost:8011 --tlsRootCertFiles $PEER1_ORG5_CA \
        --version ${VERSION} --sequence ${VERSION} --init-required
}
# commitChaincodeDefination

queryCommitted() {
    setGlobalsForPeer0Org1
    peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name ${CC_NAME}

    setGlobalsForPeer1Org1
    peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name ${CC_NAME}
}
# queryCommitted

chaincodeQuery() {
    setGlobalsForPeer0Org1
    peer chaincode query -C $CHANNEL_NAME -n ${CC_NAME} -c '{"Args":["GetBatchData","b1"]}'
}
# chaincodeQuery

# Run this function if you add any new dependency in chaincode
# presetup

# packageChaincode
# installChaincode
# queryInstalled
# approveForMyOrg1
# checkCommitReadyness

# commitChaincodeDefination
# queryCommitted
# chaincodeInvokeInit
# sleep 5
# chaincodeInvoke
# sleep 3
# chaincodeQuery

chaincodeInvokeInit() {
    setGlobalsForPeer0Org1
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        --peerAddresses localhost:7061 --tlsRootCertFiles $PEER0_ORG3_CA \
        --isInit -c '{"function":"Init","Args":[]}'
}
# chaincodeInvokeInit

chaincodeInvokeCreateBatch() {
    setGlobalsForPeer0Org1
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        --peerAddresses localhost:7061 --tlsRootCertFiles $PEER0_ORG3_CA \
        -c '{"function":"CreateBatch","Args":["b1"]}'
}
# chaincodeInvokeCreateBatch

chaincodeInvokeCreateFiberProducer() {
    setGlobalsForPeer0Org1
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        --peerAddresses localhost:7061 --tlsRootCertFiles $PEER0_ORG3_CA \
        -c '{"function":"CreateFiberProducer","Args":["fiber1","Ali","Swat"]}'
}
# chaincodeInvokeCreateFiberProducer

chaincodeInvokeAddFiberProducerData() {
    setGlobalsForPeer0Org1
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        --peerAddresses localhost:7061 --tlsRootCertFiles $PEER0_ORG3_CA \
        -c '{"Args":["AddFiberProducerData", "fiber1", "b1", "Ali", "Cotton", "Swat Barikot", "10 march 2024", "30", "100", "Mchines", "Good", "90", "Posphorous", "40", "Organic Compost", "40", "Neem Oil", "40", "200", "500", "12 March 2024", "[{\"transactionDate\":\"13 March 2024\",\"transactionRecipient\":\"Yarn1\",\"transactionQuantity\":15.0,\"transactionFiberType\":\"Cotton\"}]"]}'
}
# chaincodeInvokeAddFiberProducerData

chaincodeInvokeCreateYarnManufacturer() {
    setGlobalsForPeer0Org2
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        -c '{"Args":["CreateYarnManufacturer", "yarn1", "Qaiser", "Islamabad"]}'
}
# chaincodeInvokeCreateYarnManufacturer

chaincodeInvokeAddYarnManufacturerData() {
    setGlobalsForPeer0Org2
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        -c '{"Args":["AddYarnManufacturerData", "yarn1", "b1", "Bareeze", "Linen", "Cotton-Linen", "Combed lInen yarn", "40s", "1000", "Reactive Dyeing", "3.5", "Soft finish", "High", "[\"White\", \"blue\"]", "Medium", "2.5", "Ring Spinning Machines", "950", "14 March 2024", "[{\"transactionDate\":\"16 March 2024\",\"transactionRecipient\":\"Fabric1\",\"transactionQuantity\":20,\"transactionFiberType\":\"Linen\"}]"]}'
}
# chaincodeInvokeAddYarnManufacturerData

chaincodeInvokeCreateFabricManufacturer() {
    setGlobalsForPeer0Org3
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        -c '{"Args":["CreateFabricManufacturer", "fabric1", "Bareeze", "Islamabad street 2"]}'
}
# chaincodeInvokeCreateFabricManufacturer

chaincodeInvokeAddFabricManufacturerData() {
    setGlobalsForPeer0Org3
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        --peerAddresses localhost:7061 --tlsRootCertFiles $PEER0_ORG3_CA \
        -c '{"Args":["AddFabricManufacturerData", "fabric1", "b1", "Bareeze Inc.", "Combed Linen", "Denim", "[{\"fiberType\":\"cotton\",\"percentage\":50},{\"fiberType\":\"polyester\",\"percentage\":50}]", "[\"Softening\", \"Wrinkle-free\"]", "0.2", "White", "Soft", "Highly automated", "2.5", "[\"Shearing\", \"Bleaching\"]", "2024-03-14", "1000", "[{\"transactionDate\":\"2024-03-15\",\"transactionRecipient\":\"Apparel1\",\"transactionQuantity\":20,\"transactionFiberType\":\"Denim\"},{\"transactionDate\":\"2024-03-16\",\"transactionRecipient\":\"Apparel1\",\"transactionQuantity\":25,\"transactionFiberType\":\"Cotton\"}]"]}'
}
# chaincodeInvokeAddFabricManufacturerData

chaincodeInvokeCreateApparelManufacturer() {
    setGlobalsForPeer0Org4
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        --peerAddresses localhost:7061 --tlsRootCertFiles $PEER0_ORG3_CA \
        -c '{"Args":["CreateApparelManufacturer", "apparel1", "Bareeze apparel department", "Islamabad blue area"]}'
}
# chaincodeInvokeCreateApparelManufacturer

chaincodeInvokeAddApparelManufacturerData() {
    setGlobalsForPeer0Org4
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        --peerAddresses localhost:7061 --tlsRootCertFiles $PEER0_ORG3_CA \
        -c '{"Args":["AddApparelManufacturerData", "apparel1", "b1", "Bareeze apparel department", "Denim", "T-shirt", "1000", "[\"Cutting\", \"Sewing\", \"Finishing\"]", "0.2", "Small", "[\"Red\", \"Blue\", \"Green\"]", "Casual Style", "High-speed Sewing Machines", "500", "2024-03-14", "[{\"transactionDate\":\"2024-03-19\",\"transactionRecipient\":\"retailer1\",\"transactionQuantity\":50,\"transactionFiberType\":\"T-shirt\"},{\"transactionDate\":\"2024-03-20\",\"transactionRecipient\":\"retailer2\",\"transactionQuantity\":30,\"transactionFiberType\":\"T-shirt\"}]"]}'
}
# chaincodeInvokeAddApparelManufacturerData

chaincodeInvokeCreateRetailer() {
    setGlobalsForPeer0Org5
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        --peerAddresses localhost:7061 --tlsRootCertFiles $PEER0_ORG3_CA \
        -c '{"Args":["CreateRetailer", "retailer2", "Soor qamees botique", "Main bazar barikot"]}'
}
# chaincodeInvokeCreateRetailer

chaincodeInvokeAddRetailerData1() {
    setGlobalsForPeer0Org5
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        --peerAddresses localhost:7061 --tlsRootCertFiles $PEER0_ORG3_CA \
        -c '{"Args":["AddRetailerData", "retailer1", "b1", "Soor qamees botique", "Main bazar barikot swat", "[\"T-shirt\", \"T-shirt\"]", "[{\"productID\":\"p1\",\"quantity\":10},{\"productID\":\"p2\",\"quantity\":20}]"]}'
}
chaincodeInvokeAddRetailerData2() {
    setGlobalsForPeer0Org5
    peer chaincode invoke -o localhost:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
        -C $CHANNEL_NAME -n ${CC_NAME} \
        --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
        --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
        --peerAddresses localhost:7081 --tlsRootCertFiles $PEER0_ORG4_CA \
        --peerAddresses localhost:7061 --tlsRootCertFiles $PEER0_ORG3_CA \
        -c '{"Args":["AddRetailerData", "retailer2", "b1", "Soor qamees botique", "Main bazar barikot swat", "[\"T-shirt\", \"Jeans\"]", "[{\"type\":\"Shirt\",\"color\":\"Red\",\"brand\":\"Bareeze\",\"size\":\"M\",\"quantity\":10},{\"type\":\"Jeans\",\"color\":\"Blue\",\"brand\":\"My shirt\",\"size\":\"L\",\"quantity\":20}]"]}'
}
# chaincodeInvokeAddRetailerData1
# chaincodeInvokeAddRetailerData2

# chaincodeInvokeCreateBatch

# chaincodeInvokeCreateFiberProducer
# chaincodeInvokeAddFiberProducerData

# chaincodeInvokeCreateYarnManufacturer
# chaincodeInvokeAddYarnManufacturerData

# chaincodeInvokeCreateFabricManufacturer
# chaincodeInvokeAddFabricManufacturerData

# chaincodeInvokeCreateApparelManufacturer
# chaincodeInvokeAddApparelManufacturerData

# chaincodeInvokeCreateRetailer
# chaincodeInvokeAddRetailerData2

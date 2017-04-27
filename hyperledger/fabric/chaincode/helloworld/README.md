Using govendor to manage dependency
===================================

1. govendor init
2. govendor add +external
3. govendor add -tree github.com/hyperledger/fabric/peer

**Note**: `govendor fetch github.com/hyperledger/fabric` didn't work
although it is the recommended way according to the following guide:

https://github.com/hyperledger-archives/fabric/blob/master/docs/Setup/NodeSDK-setup.md

So instead, `github.com/hyperledger/fabric` is download before hand and put into
$GOPATH/src and then be copied into vendor file using `govendor add`.

`github.com/hyperledger/fabric/peer` is also required. Without it, chaincode
deploying will fail.

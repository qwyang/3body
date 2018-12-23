#  以太坊RPC调用
## geth启动时需要增加参数--rpc
```
API AND CONSOLE OPTIONS:
  --rpc                  Enable the HTTP-RPC server
  --rpcaddr value        HTTP-RPC server listening interface (default: "localhost")
  --rpcport value        HTTP-RPC server listening port (default: 8545)
  --rpcapi value         API's offered over the HTTP-RPC interface
```
## curl命令进行测试
注意加参数:-H 'content-type: application/json'
### 等同于eth.getBalance("0xa282a3ba1df2fa0591bfee36d972acb3c756bb57")
```
qwyang@ubuntu:~/blockchain-repo$ curl -X POST  -H 'content-type: application/json' --data '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0xa282a3ba1df2fa0591bfee36d972acb3c756bb57","latest"],"id":1}' http://localhost:8545
{"jsonrpc":"2.0","id":1,"result":"0x140ec80fa7ee880000"}
```
### 等同于eth.accounts
```
qwyang@ubuntu:~/blockchain-repo$ curl -X POST -H 'content-type: application/json' --data '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' http://localhost:8545
{"jsonrpc":"2.0","id":1,"result":["0xa282a3ba1df2fa0591bfee36d972acb3c756bb57"]}
```

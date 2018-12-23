# Operating a private network
## 参考
https://github.com/ethereum/go-ethereum 
Operating a private network

## genesis.json:
***难度值不要太低，不然挖矿太快***
```
{
  "config": {
        "chainId": 1414,
        "homesteadBlock": 0,
        "eip155Block": 0,
        "eip158Block": 0
    },
  "alloc": {
        "0xa282a3ba1df2fa0591bfee36d972acb3c756bb57":{"balance":"100000000000000000000"},
        "0x6a01a193f80969c29be0e1f07441a157ddd6a910":{"balance":"100000000000000000000"}
  },
  "coinbase"   : "0x0000000000000000000000000000000000000000",
  "difficulty" : "0x4000",
  "extraData"  : "",
  "gasLimit"   : "0x2fefd8",
  "nonce"      : "0x0000000000000056",
  "mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
  "parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp"  : "0x00"
}
```
## 目录结构：
```
qwyang@ubuntu:~/m/eth_cluster$ tree -L 2
.
├── genesis.json
└── node1
└── node2
```
## linux虚拟机节点：
```
geth --datadir ~/node1 account new
geth --datadir ~/node1 init ./genesis.json
geth --datadir ~/node1 --networkid 1414 console
admin.addPeer("enode://0e7060d15d108d9577bf20d1e7f65f6c20f5f31fcc918f082b3c1b8defc1595547eac49714cfbcf68a2b8ef91f5df7834bd41e2318358e81dc7e83ee3394caac@192.168.1.3:30303")
```
## windows host机节点：
```
geth --datadir D:\\m\\eth_cluster\\node2 account new
geth --datadir D:\\m\\eth_cluster\\node2 init ./genesis.json
geth --datadir D:\\m\\eth_cluster\\node2 --networkid 1414 console
> admin.nodeInfo.enode
`"enode://0e7060d15d108d9577bf20d1e7f65f6c20f5f31fcc918f082b3c1b8defc1595547eac49714cfbcf68a2b8ef91f5df7834bd41e2318358e81dc7e83ee3394caac@100.64.129.67:30303"`
```
## 测试
linux虚拟机发交易，windows主机挖矿

### linux
```
> personal.unlockAccount(eth.coinbase)
> eth.sendTransaction({from:eth.accounts[0],to:"0xa282a3ba1df2fa0591bfee36d972acb3c756bb57",value:web3.toWei(1,"ether")})
> eth.getTransaction("0x94ad665752afd5e56cf38d40daaf91cb16cd1701e80561ed80d3aa1b9ef8ae49")
```
### windows
```
> miner.start(1)
> miner.stop()
> eth.getBalance(eth.coinbase)
```

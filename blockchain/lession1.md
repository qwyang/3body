# first try out  
## 1.start a node  
`geth --datadir testNet --dev console 2>> test.log`  
## 2.prepare accounts  
```
eth.accounts  
personal.listAccounts  
eth.getBalance(eth.accounts[0])  
personal.newAccount("x")  
personal.newAccount("y")  
personal.newAccount("z")  
```
## 3.transaction  
```
eth.sendTransaction({from: eth.coinbase, to: eth.accounts[1], value: web3.toWei(1, "ether")})  
eth.sendTransaction({from: eth.coinbase, to: eth.accounts[2], value: web3.toWei(1, "ether")})  
eth.sendTransaction({from: eth.coinbase, to: eth.accounts[3], value: web3.toWei(1, "ether")})  
```
## 4.solidity  
```
pragma solidity ^0.4.18;  
contract hello {  
    string v;  
      
    function set(string v) public {  
        v = _v;  
    }  
  
    function get() public returns (string) {  
        return v;  
    }  
}  
```
## 5.compile in remix-ide  
  ...  
## 6.run the contract  
```
eth.defaultAccount=eth.accounts[1]  
personal.unlockAccount(eth.accounts[1],"x")  
hello.get()  
hello.set("abc")  
eth.getBalance(eth.accounts[1])  
```

## 2018.12.5
完成阅读bitcoin白皮书  
完成ubuntu solidity开发环境安装  
理解难点：双花问题
解决方法:签名,公开交易,最长链

## 2018.12.6
openssl命令行使用  
难点：拜占庭将军问题  
在p2p网络中如何实现对一个状态达成共识。
只有一个账本，最长的账本才被承认。
假设大多数人是诚实的。诚实节点大于恶意节点。
状态由多数人决定（pow基于算力/股权)。
增加作恶成本，奖励贡献。  

## 2018.12.7
Go语言实现区块链DEMO V1.0 V2.0  
难点：难度值如何自动调整  
各个节点独立计算，每隔2016个区块调整一次HASH难度，difficuty=difficuty*过去2016块计算花费时间/2016*10*60  
难点：nonce值溢出  
CoinBase交易中输入交易脚本可以任意值作为extraNonce  

## 2018.12.11
Go语言实现区块链DEMO V3.0（CLI界面/bolt数据库） V4.0(支持transaction)
Go语言条件变量：sync.Cond:wait()/signal()/broadcast()

## 2018.12.13
搭建私链
mist客户端部署合约
geth console部署合约
Solidity语言学习，重点：存储。
官网:Solidity ABI Specification.

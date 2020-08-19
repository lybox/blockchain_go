# Blockchain in Go

A blockchain implementation in Go, as described in these articles:

1. [Basic Prototype](https://jeiwan.cc/posts/building-blockchain-in-go-part-1/)
2. [Proof-of-Work](https://jeiwan.cc/posts/building-blockchain-in-go-part-2/)
3. [Persistence and CLI](https://jeiwan.cc/posts/building-blockchain-in-go-part-3/)
4. [Transactions 1](https://jeiwan.cc/posts/building-blockchain-in-go-part-4/)
5. [Addresses](https://jeiwan.cc/posts/building-blockchain-in-go-part-5/)
6. [Transactions 2](https://jeiwan.cc/posts/building-blockchain-in-go-part-6/)
7. [Network](https://jeiwan.cc/posts/building-blockchain-in-go-part-7/)


https://studygolang.com/topics/6010/comment/19468


```js
go build main.go

./main createwallet  // 创建钱包
./main createblockchain -address xxxxxx // 创建区块，并将奖励发放到钱包地址 
./main printchain // 打印现有区块
./main send -from xxxx -to xxxx -amount 2 -mine // 转账
./main getbalance -address xxx // 查看余额
```
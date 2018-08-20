## Plasma implement in go

###  Related items

- https://ethresear.ch/t/minimal-viable-plasma/426
- https://github.com/ethereum-plasma/plasma
- https://github.com/omisego/plasma-mvp

### Related tools

- golang
- solidity
- remix
- abigen

### Create your own private ethereum test network

#### setup genesis.json

```
{
  "config": {
    "chainId": 15,
    "homesteadBlock": 0,
    "eip155Block": 0,
    "eip158Block": 0
  },
  "alloc"      : {},
  "coinbase"   : "0x0000000000000000000000000000000000000000",
  "difficulty" : "0x20000",
  "extraData"  : "",
  "gasLimit"   : "0x2fefd8",
  "nonce"      : "0x0000000000000042",
  "mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
  "parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp"  : "0x00"
}

```

#### create genesis block

```
geth --datadir "./" init genesis.json
```

#### crete private network

```
geth --datadir "./" --nodiscover console 2>>geth.log
```

#### create account

```
personal.newAccount("password")
```

#### run console with rpc and ws

```
geth --identity "TestNode" --rpc --rpcport "8545" --datadir private-network --port "30303"  --rpccorsdomain "*" --rpcaddr "0.0.0.0" --networkid 10  --ws --wsport 8546 console
```

### Deploy contract with truffle

#### compile

```
truffle compile
```

#### migrate

```
truffle migrate --reset
```


### More needs to be added
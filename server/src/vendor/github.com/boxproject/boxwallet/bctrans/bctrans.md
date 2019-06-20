### 概念
bctrans主要包括币种的转账能主要体现于用户操作层面的直观逻辑

### 目录介绍
1. clientseries 各币种的基础client单元，这里目前btc,ltc,usdt分成了3次，其实可以fork出btcd库后自己封装，能实现3个统一，真正实现series的概念
> 如何一统三国呢，btc和usdt的区别就在于sendcmd的二次封装，直接添加新的cmd调用就能完成，ltc和btc的区别在于chaincfg的配置不同，理论上实现各自的包就可以了，对于btc和ltc的合并其实理解简单，修改起来可能并没有那么容易
2. client   币种客户端1:1实现，代币分支也抽象为一种币种
3. token  代币主要信息获取相关，目前只实现了eth的erc20，没有做过多的抽象，另外这里获取到代币详情后会调用bccoin缓存器币种的信息，以便代币计算使用，所以boxwallet中如果没有在这里获取过代币信息的币种是无法调用bccoin操作成功的，因为精度等都没有，这类算术计算是不被项目认可的
4. trans.go 以上3个目录的工厂封装
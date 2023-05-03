# go-fpdecimal

一个简单的精度有限的decimal库（或许不应该叫fixed-point）。代码基于[fpdecimal](https://github.com/nikolaydubina/fpdecimal)，但这个库精度写死了，我又不需要[decimal](https://github.com/shopspring/decimal)如此大的开销，所以在前者基础上做了些许改动。

目前支持与string,float的相互转换，以及加减、整数乘除。测试的不是很充分，可能会有bug。


## install

```bash
go get github.com/hesic73/go-fpdecimal
```
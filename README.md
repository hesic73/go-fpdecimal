# go-fpdecimal

a very simple fixed-point decimal library which only supports conversion to/from string and float. The capacity is limited and no arithmetic operation is supported. This is only a toy project and the functionality is enough for me.

Codes are mainly copied from [fpdecimal](https://github.com/nikolaydubina/fpdecimal). I change ParseFixedPointDecimal so that the precision is inferer from the string itself.
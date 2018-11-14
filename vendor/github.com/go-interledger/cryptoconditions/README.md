# cryptoconditions

This package provides a Go (golang) implementation of the 
[Crypto-Conditions RFC specification](https://github.com/rfcs/crypto-conditions)
intended for the Interledger protocol.

All the  test vectors from the RFC specs are included in the unit tests.

**!! For now, the `THRESHOLD-SHA-256` condition is not yet supported because 
of limitation from the ASN.1 library that we use. We're looking towards a 
solution for that. !!**


## Licensing
 
This implementation is part of the public domain. 
More information can be found in the `UNLICENSE` file.


Base64URL
===============
Implementation of base64url removing the traling = and also changing +,/ to the more url friendly -,_ characters.

The decode implementation is reversing the process before running DecodeString from the base64 package.
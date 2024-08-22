# IS-IS NET Generator

This is a simple Golang script that generates a Network Entity Title (NET) in the format X...X.XXXX.XXXX.XXXX.00, where:

- Area ID: A random hexadecimal string of 1 to 13 bytes (2 to 26 hex characters).
- System ID: A random 6-byte (12 hex characters) string.
- SEL: A fixed 1-byte value, always 00.

## Usage

`make build`

`make test`

`make run`

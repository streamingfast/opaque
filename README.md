# StreamingFast Opaque library

[![reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://pkg.go.dev/github.com/streamingfast/opaque)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Encrypt plaintexts for simple obfuscation, like database cursors.
It is used in **[StramingFast](https://github.com/streamingfast/streamingfast)**.

## Usage

Encode a proto message into an opaque element:

```
	cursor := pb.Cursor{Id: "123"}
	payload, err := proto.Marshal(&cursor)
	if err != nil { panic(err) }

	out := opaque.Encode(payload)
    // out == "xE77fJ_z6Z7UXwyx-e0kWqTtdc4yRR8vAEnvaUQSx96y"
```

Decode opaque element into a proto message:

```
	payload, err := opaque.Decode("xE77fJ_z6Z7UXwyx-e0kWqTtdc4yRR8vAEnvaUQSx96y")
	if err != nil { panic(fmt.Errorf("invalid cursor: %w", err)) }

	cursor := pb.Cursor{}
    err := proto.Unmarshal(payload, &cursor)
	if err != nil { panic(err) }

    // out == pb.Cursor{Id: "123"}
```

## Contributing

**Issues and PR in this repo related strictly to the opaque library.**

Report any protocol-specific issues in their
[respective repositories](https://github.com/streamingfast/streamingfast#protocols)

**Please first refer to the general
[StreamingFast contribution guide](https://github.com/streamingfast/streamingfast/blob/master/CONTRIBUTING.md)**,
if you wish to contribute to this code base.

This codebase uses unit tests extensively, please write and run tests.


## License

[Apache 2.0](LICENSE)

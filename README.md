# dfuse Opaque library

[![reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://pkg.go.dev/github.com/dfuse-io/opaque)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Encrypt plaintexts for simple obfuscation, like database cursors.
It is used in **[dfuse](https://github.com/dfuse-io/dfuse)**.

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
[respective repositories](https://github.com/dfuse-io/dfuse#protocols)

**Please first refer to the general
[dfuse contribution guide](https://github.com/dfuse-io/dfuse/blob/master/CONTRIBUTING.md)**,
if you wish to contribute to this code base.

## License

[Apache 2.0](LICENSE)

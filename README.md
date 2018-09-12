# ipkiss
[WIP] This command masks confidential/sensitive information (e.g. customer's mail address, name, etc...) from text files.

This command is a implementation of following bash function using `awk`.

```bash
function ipkiss() {
    awk 'BEGIN {FS=",";OFS=","} {$1="***"} {$5="***"} 1' "$filename"
}
```

## Installation

```
go get github.com/abetar0/ipkiss
```

## Author
Taro Abe ([@abetar0](https://github.com/abetar0))

## License
MIT

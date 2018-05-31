# ipkiss
[WIP] file masking command

this command is a implementation of following bash function using`awk`.

```bash
function ipkiss() {
    awk 'BEGIN {FS=",";OFS=","} {$1="***"} {$5="***"} 1' "$filename"
}
```

## Author
Taro Abe ([@abetar0](https://github.com/abetar0))

## License
MIT

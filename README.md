# ipkiss
[WIP] file masking command

this command is a implementation of following `awk` script.

```bash
awk 'BEGIN {FS=",";OFS=","} {$1="***"} {$5="***"} 1' "$filename"
```

## Author
Taro Abe ([@abetar0](https://github.com/abetar0))

## License
MIT

# go-sum

Maybe there is already a tool to sum numbers on STDIN? I couldn't find it so this one is ours.

## Install

You will need to have both `Go` and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Usage

### sum

```
cat /usr/local/cooperhewitt/tiles-list-201701.txt | awk '{ print $3 }' | .bin/sum
280977234516

cat /usr/local/cooperhewitt/tiles-list-201701.txt | awk '{ print $3 }' | ./bin/sum -humanize
368 GB
```

## See also

* https://github.com/dustin/go-humanize
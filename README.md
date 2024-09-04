# Unix Coreutils
## Description
This repo implements the Unix coreutils like `head` `tail` `cat` `echo` `wc` `tree` `yes` `false` `true` `env` in Golang

## Table of Contents

- <a href ="#head">head</a>
- <a href ="#tail">tail</a>
- <a href ="#cat">cat</a>
- <a href ="#echo">echo</a>
- <a href ="#wc">wc</a>
- <a href ="#tree">tree</a>

<hr style="background-color: #4b4c60"></hr>

 <a id = "head"></a>
 
## 1- head
### Description
  head outputs the first part of the file 
### Flags
`-n=num` : print the first NUM lines instead of the first 10
### How to use 
```
go run cmd/head/main.go [OPTIONS] [FILE]
```

 <a id = "tail"></a>
 
## 2- tail
### Description
  tail outputs the last part of the file 
### Flags
`-n=NUM` : print the last NUM lines instead of the last 10
### How to use 
```
go run cmd/tail/main.go [OPTIONS] [FILE]
```

 <a id = "cat"></a>
 
## 3- cat
### Description
  cat concatenate files and print on the standard output
  With no FILE, or when FILE is -, read standard input.
### Flags
`-n` : number all the output lines
### How to use 
```
go run cmd/cat/main.go [OPTIONS] [FILE]
```

 <a id = "echo"></a>

## 4- echo
### Description
   echo echoes the STRING(s) to standard output.
### Flags
`-n` : don't output the trailing newline
### How to use 

```
go run cmd/echo/main.go [OPTIONS] [FILE]
```

  <a id = "wc"></a>

## 5- wc
 
### Description
   wc print newline, word, and byte counts for each file
   
### Flags
- `-c` : print the bytes counts
- `-m` : print the character counts
- `-l` : print the newline counts

### How to use 

```
go run cmd/wc/main.go [OPTIONS] [FILE]
```

  <a id = "tree"></a>

## 6- tree
 
### Description
  tree prints the heirarchy of a given directory
  If no directory given , prints that of the current directory
### Flags
`-L=LEVEL` : print the heirarchy down to LEVEL number of levels

### How to use 

```
go run cmd/tree/main.go [OPTIONS] [FILE]
```

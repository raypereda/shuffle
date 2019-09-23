# shuffle
This shuffle command reads an input text file, and writes shuffled lines.
Output is set to standard output.

Example
```
Usage of C:\Users\rpereda\src\shuffle\shuffle.exe:
    shuffle [-seed N] <file>
    shuffle [-h] [-help]
<file> is a required input file.
This command reads an input text file, and writes shuffled lines.
Output is set to standard output.
  -seed int
        Optional random number seed (default 1)
$ cat temp.txt
1
2
3
4
5
6
7
8
9
10
$ .\shuffle temp.txt
2
8
5
1
10
3
4
6
9
7
```

Window 64-bit build is shuffle.exe.
Linux 64-bit build is shuffle.

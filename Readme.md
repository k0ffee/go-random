Generate random numbers for fun and profit
==========================================

Build it
--------

On Linux get a Go compiler:

Assuming a Debian-style distribution (we'll use `zsh`
later in an example):

```
sudo apt-get update
sudo apt-get install golang make zsh
```

On MacOS get a Go compiler using Homebrew:

Visit https://brew.sh/ if you don't have it already.

```
brew install golang
```

On FreeBSD get a Go compiler:

```
sudo pkg install go zsh
```

After that, clone the github-repo and build the program using `make`:

```
git clone https://github.com/k0ffee/go-random
cd go-random
make
make test
```

Run it
------

In the directory where you've just build the programs, run them like this:

```
./jar
```

or

```
./shuffle
```

or

```
./map
```

They should output a ten digit list like:

```
4 3 7 6 8 10 5 1 2 9
```

Run it multiple times to get more numbers:

If you're using the Z-Shell:

```
repeat 10 ./jar
```

Output:

```
1 5 9 2 6 3 8 7 10 4
10 3 6 5 1 8 7 4 2 9
9 10 4 1 7 5 6 3 2 8
8 4 1 2 3 7 9 6 5 10
8 7 3 5 9 1 2 10 6 4
1 10 5 9 7 6 2 8 3 4
8 2 1 9 7 4 5 10 3 6
7 5 10 8 9 1 2 6 3 4
7 9 3 10 4 6 2 8 5 1
3 2 9 8 10 1 5 7 6 4
```

If your're using the Bash shell, try:

```
for i in {1..10}; do ./jar; done
```

Alternatively run it on the Golang playground:

https://play.golang.org/p/S-Xrd_cSiay

Notice that on the playground any randomness is supressed,
it will output the same result for all runs.

I like `jar` most, then `shuffle`, then `map`.

About pseudo-random numbers
---------------------------

`map` and `shuffle` are limited in the randomness of their
output data. It is considered "good enough" for a toy program though.
Consider using `jar` for less predictable random numbers.

See https://en.wikipedia.org/wiki/Pseudorandom_number_generator

# CSC 431 Project
A compiler for the Mini language written in Go. Includes targets for LLVM and
64-bit ARMv8.  Created by [Ellis Ruckman](https://ellisruckman.com) and [Jaxon
Haws](https://www.jphaws.com). See the [Mini grammar
file](parser/mantlr/Mini.g4) for details on the compiled language and the [mini
directory](mini) for example code and benchmarks.

This is also the main project for the Compiler Construction course at Cal Poly.

  - Course: CSC 431, spring 2023
  - Professors: [Aaron Keen](http://users.csc.calpoly.edu/~akeen) and [Bruce
    DeBruhl](http://users.csc.calpoly.edu/~bdebruhl)

## Compiling programs
Build the compiler itself by running `go build`.
```sh
$ go get    # To get dependencies
$ go build    # To build the compiler
```

The compiler binary is statically linked and can be run directly.
```sh
$ ./compiler-project-c --help    # To get usage details
```

Use `-o <outfile>` to direct output to a file, as shown:
```sh
$ ./compiler-project-c -o asm.s source.mini    # To compile Mini to assembly
```

Setting `--stack` directs the compiler to use stack-based IR (instead of
register-based IR). In the same manner, passing `--llvm` tells the compiler to
output the LLVM intermediate representation instead of translating fully to ARM
assembly. See the help output for more command-line flags.
```sh
$ ./compiler-project-c --stack -o asm.s source.mini    # To use stack-based IR

$ ./compiler-project-c --llvm -o llvm.ll source.mini    # To output LLVM instead of assembly

$ ./compiler-project-c --const-prop=false --trivial-mov=true    # To toggle optimizations
```

## Dependencies
The compiler requires Go version 1.20 or higher to build.

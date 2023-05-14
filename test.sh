#!/bin/sh
# Test Mini compiler against all Mini benchmarks

set -eu
trap cleanup HUP INT QUIT ABRT

MC="./compiler-project-c"
BENCHMARKS="mini/benchmarks"
IN="input"
IN_LONGER="input.longer"
EXPECTED="output.expected"
EXPECTED_LONGER="output.longer.expected"
ACTUAL="output.actual"
ACTUAL_LONGER="output.longer.actual"
CLEANUP=true
RET=0
LOG="test.log"
ALL=true

# Remove build and testing artifacts
cleanup() {

   go clean
   
   wait

   for dir in $BENCHMARKS/*
   do
      minillvm="$dir/$(basename $dir).ll"
      bin="$dir/$(basename $dir)"
      actual="$dir/$ACTUAL"
      actual_longer="$dir/$ACTUAL_LONGER"

      set +e
      rm -f "$minillvm"
      rm -f "$bin"
      rm -f "$actual"
      rm -f "$actual_longer"
      set -e

   done

   exit "$RET"
}

print_error() {
   printf '=== Benchmark: %s ===\n%s\n\n' "$1" "$2"
}

# Compile an individual Mini benchmark
compile_benchmark() {
   dir="$BENCHMARKS/$1"
   mini="$dir/$1.mini"
   minillvm="$dir/$1.ll"

   # Check if benchmark exists
   if [ ! -f "$mini" ]; then
      print_error $1 "benchmark not found: $mini"
      RET="$((RET + 1))"
      return
   fi

   touch "$minillvm"

   set +e
   # Compile mini source code to LLVM
   $MC $mini > $minillvm 2>> $LOG
   if [ "$?" -ne 0 ]; then
      print_error $1 "mini->llvm compilation failed"
      RET=$((RET + 1))
      set -e
      return
   fi

   # Compile LLVM with clang
   clang -o $dir/$1 $minillvm 2>> $LOG
   if [ "$?" -ne 0 ]; then
      print_error $1 "llvm->binary compilation failed"
      RET=$((RET + 1))
   fi
   set -e
}

# Run an individual binary
run_binary() {
   dir="$BENCHMARKS/$1"
   actual="$BENCHMARKS/$1/$ACTUAL"
   actual_longer="$BENCHMARKS/$1/$ACTUAL_LONGER"
   bin="$dir/$1"

   # Check if benchmark has been compiled
   if [ ! -f "$bin" ]; then
      # printf 'benchmark not compiled: %s\n' $mini
      return
   fi

   set +e
   # Run benchmark with inputs
   $bin < $dir/$IN > $actual
   $bin < $dir/$IN_LONGER > $actual_longer

   # Diff benchmark actual output with expected output
   act=$(diff "$actual" "$dir/$EXPECTED")
   if [ $? -ne 0 ]; then
      RET=$((RET + 1))
      print_error "$1" "$act"
   fi

   act_long=$(diff "$actual_longer" "$dir/$EXPECTED_LONGER")
   if [ $? -ne 0 ]; then
      RET=$((RET + 1))
      print_error "$1" "$act_long"
   fi
   set -e
}

# Run an individual benchmark
run_benchmark() {
   compile_benchmark "$1"
   run_binary "$1"
}

# Run all benchmarks
run_benchmarks() {
   for b in $BENCHMARKS/*
   do
      run_benchmark $(basename "$b") &

   done
}

usage() {
   printf 'usage: %s [OPTION] [BENCHMARK]...\n' $0 >&2
   printf 'Optional arguments\n'
   printf -- '  -h, \t--help \tdisplay this help and exit\n' $0 >&2
   printf -- '  -n, \t--no-cleanup \tleave compiled files and output files\n' $0 >&2
}

# If there are aguments, display usage or run individual benchmarks
if [ "$#" -gt 0 ]; then


   # build mini compiler
   go build
   rm -f "$LOG"

   # check args
   for arg in "$@"
   do
      case "$arg" in
         '-h' | '--help')
            usage
            exit 0
            ;;
         '-n' | '--no-cleanup')
            CLEANUP=false
            ;;
         '-'*)
            printf 'invalid option: %s\n' "$arg"
            usage
            exit 1
            ;;
         *)
            ALL=false
            run_benchmark "$arg" &
            ;;
      esac
   done
fi

if [ "$ALL" = true ]; then
   go build
   rm -f "$LOG"

   # Compile the benchmarks
   run_benchmarks
fi

wait
# Cleanup all build artifacts
if [ $CLEANUP = true ]; then
   cleanup
fi

exit "$RET"

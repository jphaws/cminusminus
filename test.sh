#!/bin/sh
# Test Mini compiler against all Mini benchmarks

set -eu
trap cleanup HUP INT QUIT ABRT

MC='./compiler-project-c'
BENCHMARKS_DIR='mini/benchmarks'
INPUT='input'
INPUT_LONGER='input.longer'
EXPECTED='output.expected'
EXPECTED_LONGER='output.longer.expected'
ACTUAL='output.actual'
ACTUAL_LONGER='output.longer.actual'
LOG_FILE='test.log'

# Edit above here!

cleanup='true'
compile='compile_via_arm'
stack=''
const_prop=''
trivial=''
useless=''
ret=0
all='true'
to_run=''
pids=''

# Wait for children
wait_for_children() {
	for pid in $pids; do
		if ! wait "$pid"; then
			ret=$((ret + 1))
		fi
	done
}

# Remove build and testing artifacts
cleanup() {
	wait_for_children

    go clean

    for dir in "$BENCHMARKS_DIR"/*; do
        rm -f "$dir/$(basename $dir)" # Binary
        rm -f "$dir/$(basename $dir).ll" # LLVM
        rm -f "$dir/$(basename $dir).s" # Assembly
        rm -f "$dir/$(basename $dir).o" # Objects
        rm -f "$dir/$ACTUAL" # Actual output
        rm -f "$dir/$ACTUAL_LONGER" # Actual longer output
    done

    exit "$ret"
}

print_error() {
    printf '=== Benchmark: %s ===\n%s\n\n' "$1" "$2"
}

compile_mini() {
    dir="$BENCHMARKS_DIR/$1"
    mini="$dir/$1.mini"
	output="$2"

	llvm_arg=''
	if [ "$3" = 'llvm' ]; then
		llvm_arg='--llvm'
	fi

    # Check if benchmark exists
    if [ ! -f "$mini" ]; then
        print_error "$1" "benchmark not found: $mini"
        return 1
    fi

    # Compile Mini
    set +e
    "$MC" -o "$output" $llvm_arg $stack $const_prop $trivial $useless "$mini" 2>> "$LOG_FILE"
	if [ "$?" -ne 0 ]; then
        print_error "$1" "Mini compilation failed"
        set -e
        return 1
    fi

	set -e
	return 0
}

compile_binary() {
    dir="$BENCHMARKS_DIR/$1"
    bin="$dir/$1"
	input="$2"
    object="$dir/$1.o"

    # Compile to an object file with clang
    clang --target=aarch64 -c -o "$object" "$input" 2>> $LOG_FILE
    if [ "$?" -ne 0 ]; then
        print_error $1 "LLVM->binary compilation failed"
		set -e
		return 1
    fi

	# Compile to a binary with gcc
    aarch64-linux-gnu-gcc -no-pie -o "$bin" "$object" 2>> $LOG_FILE
    if [ "$?" -ne 0 ]; then
        print_error $1 "ARM->binary compilation failed"
		set -e
		return 1
    fi

	set -e
	return 0
}

# Compile an individual Mini benchmark (via LLVM)
compile_via_llvm() {
    dir="$BENCHMARKS_DIR/$1"
    mini="$dir/$1.mini"
    llvm="$dir/$1.ll"
    object="$dir/$1.o"

	# Compile Mini to LLVM
	compile_mini "$1" "$llvm" 'llvm' || return 1

	# Compile LLVM to binary
	compile_binary "$1" "$llvm" || return 1

	return 0
}

compile_via_arm() {
    dir="$BENCHMARKS_DIR/$1"
    bin="$dir/$1"
    mini="$dir/$1.mini"
    asm="$dir/$1.s"
    object="$dir/$1.o"

	# Compile Mini to assembly
	compile_mini "$1" "$asm" 'asm' || return 1

	# Compile assembly to binary
	compile_binary "$1" "$asm" || return 1

	return 0
}

# Run an individual binary
run_binary() {
    dir="$BENCHMARKS_DIR/$1"
    bin="$dir/$1"
    input="$dir/$INPUT"
    input_longer="$dir/$INPUT_LONGER"
    expected="$dir/$EXPECTED"
    expected_longer="$dir/$EXPECTED_LONGER"
    actual="$dir/$ACTUAL"
    actual_longer="$dir/$ACTUAL_LONGER"

    # Run benchmark with inputs
    set +e
    qemu-aarch64 -L /usr/aarch64-linux-gnu "$bin" < "$input" > "$actual"
    qemu-aarch64 -L /usr/aarch64-linux-gnu "$bin" < "$input_longer" > "$actual_longer"

    # Diff benchmark output with expected output
    diff_out=$(diff "$actual" "$expected")
    if [ $? -ne 0 ]; then
        print_error "$1" "$diff_out"
		return 1
    fi

    diff_out_longer=$(diff "$actual_longer" "$expected_longer")
    if [ $? -ne 0 ]; then
        print_error "$1" "$diff_out_longer"
		return 1
    fi

    set -e
	return 0
}

# Run an individual benchmark
run_benchmark() {
    if ! "$compile" "$1"; then
		return 1
	fi

    if ! run_binary "$1"; then
		return 1
	fi
}

usage() {
    printf 'usage: %s [OPTION] [BENCHMARK]...\n' $0 >&2
    printf 'Optional arguments\n'
    printf -- '  %-24s %s\n' '-h, --help' 'display this help and exit' >&2
    printf -- '  %-24s %s\n' '-n, --no-cleanup' 'leave compiled files and output files' >&2
    printf -- '  %-24s %s\n' '-l, --llvm' 'compile Mini to LLVM (instead of ARM)' >&2
    printf -- '  %-24s %s\n' '-s, --stack' 'compile using stack-based IR' >&2
    printf -- '  %-24s %s\n' '-c, --no-const-prop' 'disable constant propagation' >&2
    printf -- '  %-24s %s\n' '-u, --no-useless-elim' 'disable useless code elimination' >&2
    printf -- '  %-24s %s\n' '-t, --no-trivial-phi' 'disable trivial phi removal' >&2
}

# Handle any arguments
for arg in "$@"; do
	case "$arg" in
		'-h' | '--help')
			usage
			exit 0
			;;

		'-n' | '--no-cleanup')
			cleanup='false'
			;;

		'-l' | '--llvm')
			compile='compile_via_llvm'
			;;

		'-s' | '--stack')
			stack='--stack'
			;;

		'-c' | '--no-const-prop')
			const_prop='--const-prop=false'
			;;

		'-u' | '--no-useless-elim')
			useless='--useless-elim=false'
			;;

		'-t' | '--no-trivial-phi')
			trivial='--trivial-phi=false'
			;;

		'-'*)
			printf '%s: invalid option: %s\n' "$0" "$arg"
			usage
			exit 1
			;;

		*)
			all='false'
			to_run="$to_run $arg"
			;;
	esac
done

# Build Mini compiler
go build

# Remove previous log file
rm -f "$LOG_FILE"

# Run all benchmarks or specified benchmarks
if [ "$all" = 'true' ]; then
    for b in "$BENCHMARKS_DIR"/*; do
        run_benchmark "$(basename "$b")" &
		pids="$pids $!"
    done
else
	for b in $to_run; do
		run_benchmark "$b" &
		pids="$pids $!"
	done
fi

# Wait for subprocesses
wait_for_children

# Cleanup all build artifacts
[ "$cleanup" = 'true' ] && cleanup

exit "$ret"

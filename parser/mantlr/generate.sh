#!/bin/sh

alias antlr4='java -Xmx512M -cp "./antlr-4.12.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -no-listener -package mantlr Mini.g4

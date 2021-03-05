*=$0800

start:
LDX #$0F
loop:
STX $d020
DEX
BNE loop
jmp start
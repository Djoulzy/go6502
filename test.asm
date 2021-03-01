*=$0600

zp =    $22
val1 =  $01

start:
  LDA #val1
  CMP #$02
  BNE notequal
  STA zp
notequal: BEQ start

; 0600: a9 01 c9 02 d0 02 85 22 f0 f6 
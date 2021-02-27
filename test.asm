*=$c000

; var1 = $12
; var2 = $13

Label1: LDA ($50,X)  ; test ok
       LDA (#$0000),Y
       INX
Label2: CLC
       BNE Label1    ; titi
       JMP (Label); toto

fin:
       LDA toto,X
.END
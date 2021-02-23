*=$c000
Label  LDA_IM #$50
       ADC_IM #$10
       SHW
       LDA_IM #$50
       ADC_IM #$50
       SHW
       LDA_IM #$50
       ADC_IM #$90
       SHW
       LDA_IM #$50
       ADC_IM #$d0
       SHW
test   NOP
       JMP_ABS test
.END
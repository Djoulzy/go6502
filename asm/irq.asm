!source "basic-boot.asm"
+start_at $0816

init_screen      ldx #$00     ; set X to zero (black color code)
                 stx $d021    ; set background color
                 stx $d020    ; set border color
            jsr clear
            jsr init_text
;============================================================
;    some initialization and interrupt redirect setup
;============================================================

           sei         ; set interrupt disable flag

           ldy #$7f    ; $7f = %01111111
           sty $dc0d   ; Turn off CIAs Timer interrupts
           sty $dd0d   ; Turn off CIAs Timer interrupts
           lda $dc0d   ; cancel all CIA-IRQs in queue/unprocessed
           lda $dd0d   ; cancel all CIA-IRQs in queue/unprocessed
          
           lda #$01    ; Set Interrupt Request Mask...
           sta $d01a   ; ...we want IRQ by Rasterbeam

           lda #<irq   ; point IRQ Vector to our custom irq routine
           ldx #>irq 
           sta $314    ; store in $314/$315
           stx $315   

           lda #$01    ; trigger first interrupt at row zero
           sta $d012

           lda $d011   ; Bit#0 of $d011 is basically...
           and #$7f    ; ...the 9th Bit for $d012
           sta $d011   ; we need to make sure it is set to zero 

           cli         ; clear interrupt disable flag
           jmp *       ; infinite loop


;============================================================
;    custom interrupt routine
;============================================================

irq        dec $d019        ; acknowledge IRQ
           jsr colwash 
           jmp $ea81        ; return to kernel interrupt routine

color        !byte $09,$09,$02,$02,$08 
             !byte $08,$0a,$0a,$0f,$0f 
             !byte $07,$07,$01,$01,$01 
             !byte $01,$01,$01,$01,$01 
             !byte $01,$01,$01,$01,$01 
             !byte $01,$01,$01,$07,$07 
             !byte $0f,$0f,$0a,$0a,$08 
             !byte $08,$02,$02,$09,$09 

color2       !byte $09,$09,$02,$02,$08 
             !byte $08,$0a,$0a,$0f,$0f 
             !byte $07,$07,$01,$01,$01 
             !byte $01,$01,$01,$01,$01 
             !byte $01,$01,$01,$01,$01 
             !byte $01,$01,$01,$07,$07 
             !byte $0f,$0f,$0a,$0a,$08 
             !byte $08,$02,$02,$09,$09 

; the two lines of text for color washer effect

line1            !scr "    actraiser in 2013 presents...        "
line2            !scr " example effect for dustlayer tutorials  " 

clear            lda #$20     ; #$20 is the spacebar Screen Code
                 sta $0400,x  ; fill four areas with 256 spacebar characters
                 sta $0500,x 
                 sta $0600,x 
                 sta $06e8,x 
                 lda #$00     ; set foreground to black in Color Ram 
                 sta $d800,x  
                 sta $d900,x
                 sta $da00,x
                 sta $dae8,x
                 inx           ; increment X
                 bne clear     ; did X turn to zero yet?
                               ; if not, continue with the loop
                 rts           ; return from this subroutine

init_text  ldx #$00         ; init X register with $00
loop_text  lda line1,x      ; read characters from line1 table of text...
           sta $0590,x      ; ...and store in screen ram near the center
           lda line2,x      ; read characters from line2 table of text...
           sta $05e0,x      ; ...and put 2 rows below line1

           inx 
           cpx #$28         ; finished when all 40 cols of a line are processed
           bne loop_text    ; loop if we are not done yet
           rts

colwash   ldx #$27        ; load x-register with #$27 to work through 0-39 iterations
          lda color+$27   ; init accumulator with the last color from first color table

cycle1    ldy color-1,x   ; remember the current color in color table in this iteration
          sta color-1,x   ; overwrite that location with color from accumulator
          sta $d990,x     ; put it into Color Ram into column x
          tya             ; transfer our remembered color back to accumulator
          dex             ; decrement x-register to go to next iteration
          bne cycle1      ; repeat if there are iterations left
          sta color+$27   ; otherwise store te last color from accu into color table
          sta $d990       ; ... and into Color Ram
                          
colwash2  ldx #$00        ; load x-register with #$00
          lda color2+$27  ; load the last color from the second color table

cycle2    ldy color2,x    ; remember color at currently looked color2 table location
          sta color2,x    ; overwrite location with color from accumulator
          sta $d9e0,x     ; ... and write it to Color Ram
          tya             ; transfer our remembered color back to accumulator 
          inx             ; increment x-register to go to next iteraton
          cpx #$26        ; have we gone through 39 iterations yet?
          bne cycle2      ; if no, repeat
          sta color2+$27  ; if yes, store the final color from accu into color2 table
          sta $d9e0+$27   ; and write it into Color Ram
 
          rts             ; return from subroutine
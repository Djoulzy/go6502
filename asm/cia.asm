!source "basic-boot.asm"
+start_at $0810

    lda #1
    sta $d800
    lda #21
    sta $0400

    sei
    ; lda #%01010010 ; 
    ; sta $dc0f

    lda #$03    ; set timer A count value: $03FC
    sta $dc04
    lda #$FC
    sta $dc05

    lda #$ff    ; set timer B count value: $FFFF
    sta $dc06
    sta $dc07

    lda #<timer
    sta $0314
    lda #>timer
    sta $0315
    lda #%00010001
    ; lda #%11111111
    sta $dc0e
    cli

loop:
    jmp loop
    rts

timer:
    inc $0400
    lda #%11111111
    sta $dc0d
    lda $dc0d
    jmp $ea31
    ;rti
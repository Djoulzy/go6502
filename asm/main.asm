!source "basic-boot.asm"

+start_at $0816

inc $D020
jmp $0816

; jmp start

; start:
; LDX #$0F
; loop:
; STX $d020
; STX $D021
; DEX
; BNE loop
; jmp start
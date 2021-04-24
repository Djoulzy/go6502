# go6502

A basic C64 emulator for educationnal purpose.

The main goal of this developement was for me to explore the undelaying mecanism of emulation. This piece of code doesn't pretend to be as accurate
as VICE or Virtual64 but I tried to stay as close as possible to the C64 architechture.
Moreover, it's still work in progress, so ...

I use SDL2 for the display, and I'm an obsolute beginner in the domain, but the needs ar very basic.

Issues I've faced :
- At first, I tried to use GORoutines for each chipsets (CPU/VIC/CIA), but Channels communication and Mutex were so slow that I ended to stick to a single
sequential process (except from CIA).
- Moreover, my Macbook became so hot ! Fans at 100% ! I did'nt understand why... Problems with SDL2 display and multiple routines.

# Dev status:
- 6502 : 95% (some opcode are missing ... but nobody seems to use it :)
- VIC-II : 50% (no sprites / no lightpen / only text mode) - only PAL
- CIA : 10% - No keyboard input / no drive / no cassettes
- SID : 0% (I don't think I would try to emulate this chip as it's beyound my purpose)


# What could you expect:
- Read and execute KERNALS (vanilla / JiffyDOS / Dolphin)
- Use CharROM
- Read ans execute BasicROM
- Can load prg file and run them (almost)

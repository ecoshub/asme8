; ----------------------------------------------------------
; Segment Name.: SEG_ECHO
; Includes.....: 'ECHO'
; Author.......: eco
; Date.........: 28.01.2025
; Description..: Provides basic functionality for reading a character from
;                standard input and echoing it back to the screen.
; ----------------------------------------------------------
; Subroutines:
; ----------------------------------------------------------
;    Name........: ECHO
;    Description.: Continuously reads a character from the user and echoes it
;                  back to the screen.
;    Input.......: None (input comes from standard input via __GET_CHAR__)
;    Output......: Characters are echoed to the screen via __PUT_CHAR__
;    Modified....: None
; ----------------------------------------------------------

.segment "SEG_ECHO"

    extern __GET_CHAR__
    extern __PUT_CHAR__
    extern __RAM_START__
    global ECHO

ECHO:
    call __GET_CHAR__                ; read from stdin
    call __PUT_CHAR__                ; print char to screen
    jmp ECHO                         ; start over

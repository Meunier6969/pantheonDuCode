section .data
	playerone dw 0
	playertwo dw 0

	playprompt db "> "
	playpromptlen equ 2

	debugstring db "iamherehiyesitme :D", 0xa
	debugstringlen equ $ - debugstring

	playeronewonstring db "Player 1 (O) won !", 0xa
	playeronewonstringlen equ $ - playeronewonstring

	playertwowonstring db "Player 2 (X) won !", 0xa
	playertwowonstringlen equ $ - playertwowonstring


	; buffer to write/read single characters
	buffer db 0
	bufferlen equ 1

section .text
	global _start

_start:

	init_game:

	mov r10, 0

	main_game_loop:

	call print_board

	loop_input:

	; print prompt
	mov rax, 1
    mov rdi, 1
    mov rsi, playprompt
    mov rdx, playpromptlen
    syscall

	; read 
    mov rax, 0
    mov rdi, 0
    mov rsi, buffer
    mov rdx, bufferlen
    syscall

	; convert ascii -> number
	mov cl, byte [buffer] 
	sub cl, 0x30

	; bound checking
	cmp cl, 9
	jge loop_input
	cmp cl, -1
	jle loop_input
	
	; create mask
	mov rbx, 1
	shl	rbx, cl

	; check if played
	xor rax, rax
	or  rax, [playerone]
	or  rax, [playertwo]
	and rax, rbx
	cmp rax, 1
	jge loop_input

	; check which turn is it
	xor rax, rax
	mov rax, r10
	and rax, 1
	cmp rax, 1
	jge playerone_make_move
	jmp playertwo_make_move 

	; play the move
	playerone_make_move:
		or  word [playerone], bx
		jmp after_move_played

	playertwo_make_move:
		or  word [playertwo], bx
		jmp after_move_played

	after_move_played:

	; checking who won (oh god)
	mov rax, [playerone]
	and rax, 0x1c0
	cmp rax, 0x1c0
	je  playerone_won
	mov rax, [playerone]
	and rax, 0x038
	cmp rax, 0x038
	je  playerone_won
	mov rax, [playerone]
	and rax, 0x007
	cmp rax, 0x007
	je  playerone_won
	mov rax, [playerone]
	and rax, 0x124
	cmp rax, 0x124
	je  playerone_won
	mov rax, [playerone]
	and rax, 0x092
	cmp rax, 0x092
	je  playerone_won
	mov rax, [playerone]
	and rax, 0x049
	cmp rax, 0x049
	je  playerone_won
	mov rax, [playerone]
	and rax, 0x111
	cmp rax, 0x111
	je  playerone_won
	mov rax, [playerone]
	and rax, 0x054
	cmp rax, 0x054
	je  playerone_won

	mov rax, [playertwo]
	and rax, 0x1c0
	cmp rax, 0x1c0
	je  playertwo_won
	mov rax, [playertwo]
	and rax, 0x038
	cmp rax, 0x038
	je  playertwo_won
	mov rax, [playertwo]
	and rax, 0x007
	cmp rax, 0x007
	je  playertwo_won
	mov rax, [playertwo]
	and rax, 0x124
	cmp rax, 0x124
	je  playertwo_won
	mov rax, [playertwo]
	and rax, 0x092
	cmp rax, 0x092
	je  playertwo_won
	mov rax, [playertwo]
	and rax, 0x049
	cmp rax, 0x049
	je  playertwo_won
	mov rax, [playertwo]
	and rax, 0x111
	cmp rax, 0x111
	je  playertwo_won
	mov rax, [playertwo]
	and rax, 0x054
	cmp rax, 0x054
	je  playertwo_won


	turncounter:
	inc r10
	cmp r10, 9
	jge end_of_game

	incorrect_move:
	call empty_stdin
	jmp main_game_loop

	playerone_won:
	mov rax, 1
    mov rdi, 1
    mov rsi, playeronewonstring
    mov rdx, playeronewonstringlen
    syscall
	jmp end_of_game

	playertwo_won:
	mov rax, 1
    mov rdi, 1
    mov rsi, playertwowonstring
    mov rdx, playertwowonstringlen
    syscall
	jmp end_of_game

	end_of_game:
	call print_board
	call empty_stdin

	quit_game:
    mov rax, 60
    xor rdi, rdi
    syscall

print_board:
	mov rbx, 8
	call print_cell
	mov rbx, 7
	call print_cell
	mov rbx, 6
	call print_cell
	call print_new_line

	mov rbx, 5
	call print_cell
	mov rbx, 4
	call print_cell
	mov rbx, 3
	call print_cell
	call print_new_line
	
	mov rbx, 2
	call print_cell
	mov rbx, 1
	call print_cell
	mov rbx, 0
	call print_cell
	call print_new_line

	ret

print_cell: ; get position in rbx
	; get mask
	xor rcx, rcx
	mov cl, bl
	mov r8, 1	; r8 is the mask
	shl r8, cl

	mov r9, [playerone]
	and r9, r8
	cmp r9, 1
	jge .print_cell_print_o

	mov r9, [playertwo]
	and r9, r8
	cmp r9, 1
	jge .print_cell_print_x

	.print_cell_print_empty:
		mov byte [buffer], bl
		add byte [buffer], 0x30
		jmp .print_cell_print

	.print_cell_print_o:
		mov byte [buffer], 'O'
		jmp .print_cell_print

	.print_cell_print_x:
		mov byte [buffer], 'X'
		jmp .print_cell_print

	.print_cell_print:

	mov rax, 1
    mov rdi, 1
    mov rsi, buffer
    mov rdx, 1
    syscall

	mov byte [buffer], ' '
	mov rax, 1
    mov rdi, 1
    mov rsi, buffer
    mov rdx, 1
    syscall

	ret

print_new_line:
	mov byte [buffer], 0xa
	
	mov rax, 1
    mov rdi, 1
    mov rsi, buffer
    mov rdx, 1
    syscall

	ret

empty_stdin:
	;;; EMPTY STDIN START
	empty:          ; chars remain in stdin unread
	    mov rax, 0      ; read 1-char from stdin into choice
	    mov rdi, 0
	    mov rsi, buffer
	    mov rdx, 1
	    syscall

		cmp byte [buffer], 0xa
    	jne empty
	
	ret

print_debug:
	mov rax, 1
    mov rdi, 1
    mov rsi, debugstring
    mov rdx, debugstringlen
    syscall

	ret

section .data
	num db 69

section .bss
    buffer resb 16

section .text
	global _start

_start:
    mov eax, 4    ; sys_write
    mov ebx, 1    ; 1 => stdout
    mov ecx, [num] ; msg
    mov edx, 4    ; length of output
	int 0x80       ; interupt

	mov eax, 1     ; sys_exit
	mov ebx, 0     ; return code
	int 0x80       ; interupt

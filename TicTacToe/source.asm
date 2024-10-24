section .data
	

section .bss
    buffer resb 16

section .text
	global _start

_start:
    push    '!'         ; push qword 0x21
	mov     rax, 1      ; write call number, __NR_write from unistd_64.h
	mov     edi, 1      ; write to stdout (int fd=1)
	mov     rsi, rsp    ; use char on stack
	mov     rdx, 1      ; size_t len = 1 char to write.
	syscall            ; call the kernel, it looks at registers to decide what to do
	add     rsp, 8      ; restore stack pointer


	mov eax, 1     ; sys_exit
	mov ebx, 0     ; return code
	int 0x80       ; interupt

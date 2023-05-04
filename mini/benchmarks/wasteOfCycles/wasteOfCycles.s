	.arch armv7-a
	.text
	.align 2
	.global function
function:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r4, r0
	mov r0, #0
	cmp r4, #0
	movle r0, #1
	cmp r0, #1
	beq .LU2
	b .LU3
.LU2:
	movw r0, #0
	b .LU0
.LU3:
	b .LU4
.LU4:
	mul r2, r4, r4
	movw r0, #0
	mov r1, #0
	movw r3, #0
	cmp r3, r2
	movlt r1, #1
	cmp r1, #1
	beq .LU5
	b .LU6
.LU5:
	mov r5, r0
	add r0, r5, r4
	mov r1, r0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	add r1, r5, #1
	mul r3, r4, r4
	mov r0, r1
	mov r2, #0
	cmp r1, r3
	movlt r2, #1
	cmp r2, #1
	beq .LU5
	b .LU6
.LU6:
	sub r0, r4, #1
	bl function
	b .LU0
.LU0:
	pop {r4, r5}
	pop {fp, pc}
	.size function, .-function
	.align 2
	.global main
main:
.LU8:
	push {fp, lr}
	add fp, sp, #4
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r0, #:lower16:.read_scratch
	movt r0, #:upper16:.read_scratch
	ldr r0, [r0]
	bl function
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU7
.LU7:
	movw r0, #0
	pop {fp, pc}
	.size main, .-main
	.section	.rodata
	.align	2
.PRINTLN_FMT:
	.asciz	"%ld\n"
	.align	2
.PRINT_FMT:
	.asciz	"%ld "
	.align	2
.READ_FMT:
	.asciz	"%ld"
	.comm	.read_scratch,4,4
	.global	__aeabi_idiv

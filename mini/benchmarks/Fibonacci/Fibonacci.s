	.arch armv7-a
	.text
	.align 2
	.global computeFib
computeFib:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r4, r0
	mov r0, #0
	cmp r4, #0
	moveq r0, #1
	cmp r0, #1
	beq .LU2
	b .LU3
.LU2:
	movw r0, #0
	b .LU0
.LU3:
	mov r0, #0
	cmp r4, #2
	movle r0, #1
	cmp r0, #1
	beq .LU4
	b .LU5
.LU4:
	movw r0, #1
	b .LU0
.LU5:
	sub r0, r4, #1
	bl computeFib
	mov r5, r0
	sub r0, r4, #2
	bl computeFib
	add r0, r5, r0
	b .LU0
.LU0:
	pop {r4, r5}
	pop {fp, pc}
	.size computeFib, .-computeFib
	.align 2
	.global main
main:
.LU7:
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
	bl computeFib
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU6
.LU6:
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

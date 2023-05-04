	.arch armv7-a
	.text
	.align 2
	.global mod
mod:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r5, r0
	mov r4, r1
	mov r1, r4
	mov r0, r5
	bl __aeabi_idiv
	mul r0, r0, r4
	sub r0, r5, r0
	b .LU0
.LU0:
	pop {r4, r5}
	pop {fp, pc}
	.size mod, .-mod
	.align 2
	.global hailstone
hailstone:
.LU3:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	b .LU4
.LU4:
	mov r4, r0
	mov r1, r4
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	movw r1, #2
	mov r0, r4
	bl mod
	mov r1, r0
	mov r0, #0
	cmp r1, #1
	moveq r0, #1
	cmp r0, #1
	beq .LU6
	b .LU7
.LU6:
	movw r0, #3
	mul r0, r0, r4
	add r0, r0, #1
	b .LU8
.LU7:
	movw r1, #2
	mov r0, r4
	bl __aeabi_idiv
	b .LU8
.LU8:
	mov r1, r0
	mov r0, #0
	cmp r1, #1
	movle r0, #1
	cmp r0, #1
	beq .LU9
	b .LU10
.LU9:
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU2
.LU2:
	pop {r4}
	pop {fp, pc}
.LU10:
	b .LU11
.LU11:
	mov r0, r1
	b .LU4
	.size hailstone, .-hailstone
	.align 2
	.global main
main:
.LU13:
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
	bl hailstone
	b .LU12
.LU12:
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

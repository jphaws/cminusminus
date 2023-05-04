	.arch armv7-a
	.text
	.align 2
	.global isqrt
isqrt:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	movw r1, #3
	movw ip, #1
	movw r3, #3
	mov r2, #0
	movw lr, #1
	cmp lr, r0
	movle r2, #1
	cmp r2, #1
	beq .LU2
	b .LU3
.LU2:
	mov r2, ip
	add lr, r2, r1
	add r2, r1, #2
	mov r1, r2
	mov ip, lr
	mov r3, r2
	mov r2, #0
	cmp lr, r0
	movle r2, #1
	cmp r2, #1
	beq .LU2
	b .LU3
.LU3:
	mov r0, r3
	movw r1, #2
	bl __aeabi_idiv
	sub r0, r0, #1
	b .LU0
.LU0:
	pop {fp, pc}
	.size isqrt, .-isqrt
	.align 2
	.global prime
prime:
.LU5:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6}
	mov r5, r0
	mov r0, #0
	cmp r5, #2
	movlt r0, #1
	cmp r0, #1
	beq .LU6
	b .LU7
.LU6:
	movw r0, #0
	b .LU4
.LU7:
	mov r0, r5
	bl isqrt
	mov r4, r0
	movw r1, #2
	mov r0, #0
	movw r2, #2
	cmp r2, r4
	movle r0, #1
	cmp r0, #1
	beq .LU8
	b .LU9
.LU8:
	mov r6, r1
	mov r1, r6
	mov r0, r5
	bl __aeabi_idiv
	mul r0, r0, r6
	sub r1, r5, r0
	mov r0, #0
	cmp r1, #0
	moveq r0, #1
	cmp r0, #1
	beq .LU10
	b .LU11
.LU10:
	movw r0, #0
	b .LU4
.LU11:
	b .LU12
.LU12:
	add r2, r6, #1
	mov r1, r2
	mov r0, #0
	cmp r2, r4
	movle r0, #1
	cmp r0, #1
	beq .LU8
	b .LU9
.LU9:
	movw r0, #1
	b .LU4
.LU4:
	pop {r4, r5, r6}
	pop {fp, pc}
	.size prime, .-prime
	.align 2
	.global main
main:
.LU14:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r4, #:lower16:.read_scratch
	movt r4, #:upper16:.read_scratch
	ldr r4, [r4]
	movw r1, #0
	mov r0, #0
	movw r2, #0
	cmp r2, r4
	movle r0, #1
	cmp r0, #1
	beq .LU15
	b .LU16
.LU15:
	mov r5, r1
	mov r0, r5
	bl prime
	cmp r0, #1
	beq .LU17
	b .LU18
.LU17:
	mov r1, r5
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU19
.LU18:
	b .LU19
.LU19:
	add r0, r5, #1
	mov r1, r0
	mov r2, #0
	cmp r0, r4
	movle r2, #1
	cmp r2, #1
	beq .LU15
	b .LU16
.LU16:
	b .LU13
.LU13:
	movw r0, #0
	pop {r4, r5}
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

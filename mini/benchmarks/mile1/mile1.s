	.arch armv7-a
	.text
	.align 2
	.global calcPower
calcPower:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	mov lr, r0
	mov r0, r1
	movw r3, #1
	movw r2, #1
	mov ip, #0
	cmp r1, #0
	movgt ip, #1
	cmp ip, #1
	beq .LU2
	b .LU3
.LU2:
	mov r2, r0
	mov r0, r3
	mul r1, r0, lr
	sub ip, r2, #1
	mov r0, ip
	mov r3, r1
	mov r2, r1
	mov r1, #0
	cmp ip, #0
	movgt r1, #1
	cmp r1, #1
	beq .LU2
	b .LU3
.LU3:
	mov r0, r2
	b .LU0
.LU0:
	pop {fp, pc}
	.size calcPower, .-calcPower
	.align 2
	.global main
main:
.LU5:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	movw r0, #8
	bl malloc
	mov r4, r0
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	ldr r1, [r1]
	add r0, r4, #4
	str r1, [r0]
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r0, #:lower16:.read_scratch
	movt r0, #:upper16:.read_scratch
	ldr r0, [r0]
	mov r1, #0
	cmp r0, #0
	movlt r1, #1
	cmp r1, #1
	beq .LU6
	b .LU7
.LU6:
	movw r0, #65535
	movt r0, #65535
	b .LU4
.LU7:
	b .LU8
.LU8:
	mov r1, r4
	str r0, [r1]
	movw r2, #0
	b .LU9
.LU9:
	mov r0, r2
	add r5, r0, #1
	add r0, r4, #4
	ldr r2, [r0]
	mov r0, r4
	ldr r0, [r0]
	mov r1, r0
	mov r0, r2
	bl calcPower
	mov r1, r0
	mov r2, r5
	mov r0, #0
	movw r3, #16960
	movt r3, #15
	cmp r5, r3
	movlt r0, #1
	cmp r0, #1
	beq .LU9
	b .LU10
.LU10:
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r0, #0
	b .LU4
.LU4:
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

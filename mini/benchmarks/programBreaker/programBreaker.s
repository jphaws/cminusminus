	.arch armv7-a
	.comm	GLOBAL,4,4
	.comm	count,4,4

	.text
	.align 2
	.global fun2
fun2:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	mov r2, r0
	mov r0, r1
	mov r1, #0
	cmp r2, #0
	moveq r1, #1
	cmp r1, #1
	beq .LU2
	b .LU3
.LU2:
	b .LU0
.LU3:
	sub r2, r2, #1
	mov r1, r0
	mov r0, r2
	bl fun2
	b .LU0
.LU0:
	pop {fp, pc}
	.size fun2, .-fun2
	.align 2
	.global fun1
fun1:
.LU5:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6, r7}
	mov r5, r0
	mov r4, r1
	mov r6, r2
	movw r0, #2
	mul r0, r5, r0
	rsb r7, r0, #11
	mov r1, r4
	movw r0, #4
	bl __aeabi_idiv
	add r0, r7, r0
	add r0, r0, r6
	mov r1, #0
	cmp r0, r4
	movgt r1, #1
	cmp r1, #1
	beq .LU6
	b .LU7
.LU6:
	mov r1, r5
	bl fun2
	b .LU4
.LU7:
	mov r1, #0
	cmp r0, r4
	movle r1, #1
	and r1, r1, #1
	cmp r1, #1
	beq .LU8
	b .LU9
.LU8:
	mov r1, r4
	bl fun2
	b .LU4
.LU9:
	b .LU10
.LU10:
	b .LU11
.LU11:
	b .LU4
.LU4:
	pop {r4, r5, r6, r7}
	pop {fp, pc}
	.size fun1, .-fun1
	.align 2
	.global main
main:
.LU13:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	ldr r1, [r1]
	mov r2, r1
	mov r0, #0
	movw r3, #10000
	cmp r1, r3
	movlt r0, #1
	cmp r0, #1
	beq .LU14
	b .LU15
.LU14:
	mov r4, r2
	movw r2, #5
	mov r1, r4
	movw r0, #3
	bl fun1
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	add r1, r4, #1
	mov r2, r1
	mov r0, #0
	movw r3, #10000
	cmp r1, r3
	movlt r0, #1
	cmp r0, #1
	beq .LU14
	b .LU15
.LU15:
	b .LU12
.LU12:
	movw r0, #0
	pop {r4}
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

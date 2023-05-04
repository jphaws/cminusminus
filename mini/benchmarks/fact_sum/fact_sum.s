	.arch armv7-a
	.text
	.align 2
	.global sum
sum:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	mov r2, r0
	mov r0, r1
	add r0, r2, r0
	b .LU0
.LU0:
	pop {fp, pc}
	.size sum, .-sum
	.align 2
	.global fact
fact:
.LU3:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r4, r0
	mov r1, #0
	cmp r4, #1
	moveq r1, #1
	mov r0, #0
	cmp r4, #0
	moveq r0, #1
	orr r0, r1, r0
	cmp r0, #1
	beq .LU4
	b .LU5
.LU4:
	movw r0, #1
	b .LU2
.LU5:
	b .LU6
.LU6:
	mov r0, #0
	cmp r4, #1
	movle r0, #1
	cmp r0, #1
	beq .LU7
	b .LU8
.LU7:
	movw r0, #65535
	movt r0, #65535
	mul r0, r0, r4
	bl fact
	b .LU2
.LU8:
	b .LU9
.LU9:
	sub r0, r4, #1
	bl fact
	mul r0, r4, r0
	b .LU2
.LU2:
	pop {r4}
	pop {fp, pc}
	.size fact, .-fact
	.align 2
	.global main
main:
.LU11:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	b .LU12
.LU12:
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r4, #:lower16:.read_scratch
	movt r4, #:upper16:.read_scratch
	ldr r4, [r4]
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r5, #:lower16:.read_scratch
	movt r5, #:upper16:.read_scratch
	ldr r5, [r5]
	mov r0, r4
	bl fact
	mov r4, r0
	mov r0, r5
	bl fact
	mov r1, r0
	mov r0, r4
	bl sum
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r0, #:lower16:.read_scratch
	movt r0, #:upper16:.read_scratch
	ldr r0, [r0]
	mov r1, #0
	movw r2, #65535
	movt r2, #65535
	cmp r0, r2
	movne r1, #1
	cmp r1, #1
	beq .LU12
	b .LU13
.LU13:
	b .LU10
.LU10:
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

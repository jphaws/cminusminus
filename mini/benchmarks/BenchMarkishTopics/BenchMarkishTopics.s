	.arch armv7-a
	.comm	intList,4,4

	.text
	.align 2
	.global length
length:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	mov r1, #0
	cmp r0, #0
	moveq r1, #1
	cmp r1, #1
	beq .LU2
	b .LU3
.LU2:
	movw r0, #0
	b .LU0
.LU3:
	b .LU4
.LU4:
	ldr r0, [r0]
	bl length
	add r0, r0, #1
	b .LU0
.LU0:
	pop {fp, pc}
	.size length, .-length
	.align 2
	.global addToFront
addToFront:
.LU6:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r5, r0
	mov r4, r1
	mov r0, #0
	cmp r5, #0
	moveq r0, #1
	cmp r0, #1
	beq .LU7
	b .LU8
.LU7:
	movw r0, #8
	bl malloc
	mov r1, r0
	add r0, r1, #4
	str r4, [r0]
	mov r0, r1
	movw r2, #0
	str r2, [r0]
	mov r0, r1
	b .LU5
.LU8:
	b .LU9
.LU9:
	movw r0, #8
	bl malloc
	add r1, r0, #4
	str r4, [r1]
	mov r1, r0
	str r5, [r1]
	b .LU5
.LU5:
	pop {r4, r5}
	pop {fp, pc}
	.size addToFront, .-addToFront
	.align 2
	.global deleteFirst
deleteFirst:
.LU11:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r1, #0
	cmp r0, #0
	moveq r1, #1
	cmp r1, #1
	beq .LU12
	b .LU13
.LU12:
	movw r0, #0
	b .LU10
.LU13:
	b .LU14
.LU14:
	mov r1, r0
	ldr r4, [r1]
	bl free
	mov r0, r4
	b .LU10
.LU10:
	pop {r4}
	pop {fp, pc}
	.size deleteFirst, .-deleteFirst
	.align 2
	.global main
main:
.LU16:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r0, #:lower16:.read_scratch
	movt r0, #:upper16:.read_scratch
	ldr r0, [r0]
	movw r1, #:lower16:intList
	movt r1, #:upper16:intList
	str r0, [r1]
	movw r0, #:lower16:intList
	movt r0, #:upper16:intList
	ldr r0, [r0]
	movw r3, #0
	movw r2, #0
	mov r1, #0
	cmp r0, #0
	movgt r1, #1
	cmp r1, #1
	beq .LU17
	b .LU18
.LU17:
	mov r2, r3
	movw r0, #:lower16:intList
	movt r0, #:upper16:intList
	ldr r0, [r0]
	mov r1, r0
	mov r0, r2
	bl addToFront
	mov r4, r0
	add r0, r4, #4
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	movw r0, #:lower16:intList
	movt r0, #:upper16:intList
	ldr r0, [r0]
	sub r0, r0, #1
	movw r1, #:lower16:intList
	movt r1, #:upper16:intList
	str r0, [r1]
	movw r1, #:lower16:intList
	movt r1, #:upper16:intList
	ldr r1, [r1]
	mov r3, r4
	mov r2, r4
	mov r0, #0
	cmp r1, #0
	movgt r0, #1
	cmp r0, #1
	beq .LU17
	b .LU18
.LU18:
	mov r4, r2
	mov r0, r4
	bl length
	mov r1, r0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	mov r0, r4
	bl length
	mov ip, r0
	mov r2, r4
	movw r1, #0
	movw r3, #0
	mov r0, #0
	cmp ip, #0
	movgt r0, #1
	cmp r0, #1
	beq .LU19
	b .LU20
.LU19:
	mov r5, r2
	mov r0, r1
	add r1, r5, #4
	ldr r1, [r1]
	add r4, r0, r1
	mov r0, r5
	bl length
	mov r1, r0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	mov r0, r5
	bl deleteFirst
	mov r5, r0
	mov r0, r5
	bl length
	mov ip, r0
	mov r2, r5
	mov r1, r4
	mov r3, r4
	mov r0, #0
	cmp ip, #0
	movgt r0, #1
	cmp r0, #1
	beq .LU19
	b .LU20
.LU20:
	mov r0, r3
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU15
.LU15:
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

	.arch armv7-a
	.text
	.align 2
	.global buildList
buildList:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6, r7, r8, r9}
	movw r0, #8
	bl malloc
	mov r9, r0
	movw r0, #8
	bl malloc
	mov r7, r0
	movw r0, #8
	bl malloc
	mov r8, r0
	movw r0, #8
	bl malloc
	mov r4, r0
	movw r0, #8
	bl malloc
	mov r6, r0
	movw r0, #8
	bl malloc
	mov r5, r0
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	ldr r1, [r1]
	add r0, r9, #4
	str r1, [r0]
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	ldr r1, [r1]
	add r0, r7, #4
	str r1, [r0]
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	ldr r1, [r1]
	add r0, r8, #4
	str r1, [r0]
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
	add r1, r6, #4
	str r0, [r1]
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r0, #:lower16:.read_scratch
	movt r0, #:upper16:.read_scratch
	ldr r0, [r0]
	add r1, r5, #4
	str r0, [r1]
	mov r0, r9
	str r7, [r0]
	mov r0, r7
	str r8, [r0]
	mov r0, r8
	str r4, [r0]
	mov r0, r4
	str r6, [r0]
	mov r0, r6
	str r5, [r0]
	mov r1, r5
	movw r0, #0
	str r0, [r1]
	b .LU0
.LU0:
	mov r0, r9
	pop {r4, r5, r6, r7, r8, r9}
	pop {fp, pc}
	.size buildList, .-buildList
	.align 2
	.global multiple
multiple:
.LU3:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6}
	add r1, r0, #4
	ldr r1, [r1]
	ldr r0, [r0]
	movw r2, #0
	mov r3, r0
	b .LU4
.LU4:
	mov r5, r2
	mov r0, r3
	mov r2, r1
	add r1, r0, #4
	ldr r1, [r1]
	mul r4, r2, r1
	ldr r6, [r0]
	mov r1, r4
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	add ip, r5, #1
	mov r2, ip
	mov r3, r6
	mov r1, r4
	mov r0, #0
	cmp ip, #5
	movlt r0, #1
	cmp r0, #1
	beq .LU4
	b .LU5
.LU5:
	b .LU2
.LU2:
	mov r0, r4
	pop {r4, r5, r6}
	pop {fp, pc}
	.size multiple, .-multiple
	.align 2
	.global add
add:
.LU7:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6}
	add r1, r0, #4
	ldr r1, [r1]
	ldr r0, [r0]
	movw r2, #0
	mov ip, r1
	b .LU8
.LU8:
	mov r5, r2
	mov r1, r0
	mov r0, ip
	add r2, r1, #4
	ldr r2, [r2]
	add r4, r0, r2
	mov r0, r1
	ldr r6, [r0]
	mov r1, r4
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	add r3, r5, #1
	mov r2, r3
	mov r0, r6
	mov ip, r4
	mov r1, #0
	cmp r3, #5
	movlt r1, #1
	cmp r1, #1
	beq .LU8
	b .LU9
.LU9:
	b .LU6
.LU6:
	mov r0, r4
	pop {r4, r5, r6}
	pop {fp, pc}
	.size add, .-add
	.align 2
	.global recurseList
recurseList:
.LU11:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r4, r0
	mov r0, r4
	ldr r0, [r0]
	mov r1, #0
	cmp r0, #0
	moveq r1, #1
	cmp r1, #1
	beq .LU12
	b .LU13
.LU12:
	add r0, r4, #4
	ldr r0, [r0]
	b .LU10
.LU13:
	mov r0, r4
	ldr r0, [r0]
	bl recurseList
	add r1, r4, #4
	ldr r1, [r1]
	mul r0, r1, r0
	b .LU10
.LU10:
	pop {r4}
	pop {fp, pc}
	.size recurseList, .-recurseList
	.align 2
	.global main
main:
.LU15:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6, r7}
	bl buildList
	mov r4, r0
	mov r0, r4
	bl multiple
	mov r5, r0
	mov r0, r4
	bl add
	movw r1, #2
	bl __aeabi_idiv
	sub r5, r5, r0
	movw r1, #0
	movw r0, #0
	b .LU16
.LU16:
	mov r7, r1
	mov r6, r0
	mov r0, r4
	bl recurseList
	add r6, r6, r0
	add r3, r7, #1
	mov r1, r3
	mov r0, r6
	mov r2, #0
	cmp r3, #2
	movlt r2, #1
	cmp r2, #1
	beq .LU16
	b .LU17
.LU17:
	mov r1, r6
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, r6
	mov r3, r6
	mov r1, #0
	cmp r6, #0
	movne r1, #1
	cmp r1, #1
	beq .LU18
	b .LU19
.LU18:
	sub r2, r0, #1
	mov r0, r2
	mov r3, r2
	mov r1, #0
	cmp r2, #0
	movne r1, #1
	cmp r1, #1
	beq .LU18
	b .LU19
.LU19:
	mov r4, r3
	mov r1, r5
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r1, r4
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU14
.LU14:
	movw r0, #0
	pop {r4, r5, r6, r7}
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

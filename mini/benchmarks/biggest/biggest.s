	.arch armv7-a
	.text
	.align 2
	.global getIntList
getIntList:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
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
	mov r0, #0
	movw r2, #65535
	movt r2, #65535
	cmp r1, r2
	moveq r0, #1
	cmp r0, #1
	beq .LU2
	b .LU3
.LU2:
	mov r0, r4
	str r1, [r0]
	add r0, r4, #4
	movw r1, #0
	str r1, [r0]
	b .LU0
.LU3:
	mov r0, r4
	str r1, [r0]
	bl getIntList
	add r1, r4, #4
	str r0, [r1]
	b .LU0
.LU0:
	mov r0, r4
	pop {r4}
	pop {fp, pc}
	.size getIntList, .-getIntList
	.align 2
	.global biggest
biggest:
.LU5:
	push {fp, lr}
	add fp, sp, #4
	mov r2, r1
	mov r1, #0
	cmp r0, r2
	movgt r1, #1
	cmp r1, #1
	beq .LU6
	b .LU7
.LU6:
	b .LU4
.LU7:
	mov r0, r2
	b .LU4
.LU4:
	pop {fp, pc}
	.size biggest, .-biggest
	.align 2
	.global biggestInList
biggestInList:
.LU9:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r1, r0
	ldr r2, [r1]
	add r1, r0, #4
	ldr lr, [r1]
	mov ip, r0
	mov r3, r2
	mov r0, r2
	mov r1, #0
	cmp lr, #0
	movne r1, #1
	cmp r1, #1
	beq .LU10
	b .LU11
.LU10:
	mov r4, ip
	mov r2, r3
	mov r0, r4
	ldr r0, [r0]
	mov r1, r0
	mov r0, r2
	bl biggest
	mov r1, r0
	add r0, r4, #4
	ldr r0, [r0]
	add r2, r0, #4
	ldr r2, [r2]
	mov ip, r0
	mov r3, r1
	mov r0, r1
	mov r1, #0
	cmp r2, #0
	movne r1, #1
	cmp r1, #1
	beq .LU10
	b .LU11
.LU11:
	b .LU8
.LU8:
	pop {r4}
	pop {fp, pc}
	.size biggestInList, .-biggestInList
	.align 2
	.global main
main:
.LU13:
	push {fp, lr}
	add fp, sp, #4
	bl getIntList
	bl biggestInList
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
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

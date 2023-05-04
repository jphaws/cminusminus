	.arch armv7-a
	.text
	.align 2
	.global getRands
getRands:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6, r7}
	mov r4, r0
	mov r6, r1
	mul r5, r4, r4
	movw r0, #8
	bl malloc
	mov r2, r0
	add r0, r2, #4
	str r5, [r0]
	mov r0, r2
	movw r1, #0
	str r1, [r0]
	sub r1, r6, #1
	mov r3, r1
	mov lr, r2
	mov ip, r5
	movw r2, #0
	mov r0, #0
	cmp r1, #0
	movgt r0, #1
	cmp r0, #1
	beq .LU2
	b .LU3
.LU2:
	mov r7, r3
	mov r6, lr
	mov r0, ip
	mul r0, r0, r0
	mov r1, r4
	bl __aeabi_idiv
	mov r5, r0
	movw r1, #2
	mov r0, r4
	bl __aeabi_idiv
	mul r0, r5, r0
	add r5, r0, #1
	movw r1, #51712
	movt r1, #15258
	mov r0, r5
	bl __aeabi_idiv
	mov r1, r0
	movw r0, #51712
	movt r0, #15258
	mul r0, r1, r0
	sub r5, r5, r0
	movw r0, #8
	bl malloc
	mov r1, r0
	add r0, r1, #4
	str r5, [r0]
	mov r0, r1
	str r6, [r0]
	sub r0, r7, #1
	mov r3, r0
	mov lr, r1
	mov ip, r5
	mov r2, r1
	mov r1, #0
	cmp r0, #0
	movgt r1, #1
	cmp r1, #1
	beq .LU2
	b .LU3
.LU3:
	mov r0, r2
	b .LU0
.LU0:
	pop {r4, r5, r6, r7}
	pop {fp, pc}
	.size getRands, .-getRands
	.align 2
	.global calcMean
calcMean:
.LU5:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r1, r0
	mov r0, r1
	movw ip, #0
	movw r2, #0
	movw r5, #0
	movw r4, #0
	mov r3, #0
	cmp r1, #0
	movne r3, #1
	cmp r3, #1
	beq .LU6
	b .LU7
.LU6:
	mov r3, ip
	mov r1, r2
	add r1, r1, #1
	add r2, r0, #4
	ldr r2, [r2]
	add r3, r3, r2
	ldr lr, [r0]
	mov r0, lr
	mov ip, r3
	mov r2, r1
	mov r5, r3
	mov r4, r1
	mov r1, #0
	cmp lr, #0
	movne r1, #1
	cmp r1, #1
	beq .LU6
	b .LU7
.LU7:
	mov r2, r5
	mov r0, r4
	mov r1, #0
	cmp r0, #0
	movne r1, #1
	cmp r1, #1
	beq .LU8
	b .LU9
.LU8:
	mov r1, r0
	mov r0, r2
	bl __aeabi_idiv
	b .LU10
.LU9:
	movw r0, #0
	b .LU10
.LU10:
	b .LU4
.LU4:
	pop {r4, r5}
	pop {fp, pc}
	.size calcMean, .-calcMean
	.align 2
	.global approxSqrt
approxSqrt:
.LU12:
	push {fp, lr}
	add fp, sp, #4
	mov ip, r0
	movw r3, #1
	movw r2, #1
	mov r0, #0
	movw r1, #0
	cmp r1, ip
	movlt r0, #1
	cmp r0, #1
	beq .LU13
	b .LU14
.LU13:
	mov r2, r3
	mul r1, r2, r2
	add r0, r2, #1
	mov r3, r0
	mov r0, #0
	cmp r1, ip
	movlt r0, #1
	cmp r0, #1
	beq .LU13
	b .LU14
.LU14:
	mov r0, r2
	b .LU11
.LU11:
	pop {fp, pc}
	.size approxSqrt, .-approxSqrt
	.align 2
	.global approxSqrtAll
approxSqrtAll:
.LU16:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r2, r0
	mov r1, #0
	cmp r0, #0
	movne r1, #1
	cmp r1, #1
	beq .LU17
	b .LU18
.LU17:
	mov r4, r2
	add r0, r4, #4
	ldr r0, [r0]
	bl approxSqrt
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, r4
	ldr r1, [r0]
	mov r2, r1
	mov r0, #0
	cmp r1, #0
	movne r0, #1
	cmp r0, #1
	beq .LU17
	b .LU18
.LU18:
	b .LU15
.LU15:
	pop {r4}
	pop {fp, pc}
	.size approxSqrtAll, .-approxSqrtAll
	.align 2
	.global range
range:
.LU20:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r2, r0
	movw r5, #0
	movw r3, #0
	mov r0, r2
	movw r1, #1
	movw lr, #0
	movw ip, #0
	mov r4, #0
	cmp r2, #0
	movne r4, #1
	cmp r4, #1
	beq .LU21
	b .LU22
.LU21:
	mov r2, r5
	mov ip, r1
	cmp ip, #1
	beq .LU23
	b .LU24
.LU23:
	add r1, r0, #4
	ldr r2, [r1]
	add r1, r0, #4
	ldr r3, [r1]
	movw r1, #0
	b .LU31
.LU24:
	add ip, r0, #4
	ldr lr, [ip]
	mov ip, #0
	cmp lr, r3
	movlt ip, #1
	cmp ip, #1
	beq .LU25
	b .LU26
.LU25:
	add r3, r0, #4
	ldr r3, [r3]
	mov ip, r2
	mov r2, r3
	b .LU30
.LU26:
	add ip, r0, #4
	ldr lr, [ip]
	mov ip, #0
	cmp lr, r2
	movgt ip, #1
	cmp ip, #1
	beq .LU27
	b .LU28
.LU27:
	add r2, r0, #4
	ldr r2, [r2]
	b .LU29
.LU28:
	b .LU29
.LU29:
	mov ip, r2
	mov r2, r3
	b .LU30
.LU30:
	mov r3, ip
	b .LU31
.LU31:
	mov ip, r3
	ldr r4, [r0]
	mov r5, ip
	mov r3, r2
	mov r0, r4
	mov lr, ip
	mov ip, r2
	mov r2, #0
	cmp r4, #0
	movne r2, #1
	cmp r2, #1
	beq .LU21
	b .LU22
.LU22:
	mov r4, lr
	mov r0, ip
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r1, r4
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU19
.LU19:
	pop {r4, r5}
	pop {fp, pc}
	.size range, .-range
	.align 2
	.global main
main:
.LU33:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
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
	movw r0, #:lower16:.read_scratch
	movt r0, #:upper16:.read_scratch
	ldr r0, [r0]
	mov r1, r0
	mov r0, r4
	bl getRands
	mov r4, r0
	mov r0, r4
	bl calcMean
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, r4
	bl range
	mov r0, r4
	bl approxSqrtAll
	b .LU32
.LU32:
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

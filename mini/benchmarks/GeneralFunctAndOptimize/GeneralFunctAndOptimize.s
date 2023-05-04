	.arch armv7-a
	.comm	interval,4,4
	.comm	end,4,4

	.text
	.align 2
	.global multBy4xTimes
multBy4xTimes:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r4, r0
	mov r0, r1
	mov r1, #0
	cmp r0, #0
	movle r1, #1
	cmp r1, #1
	beq .LU2
	b .LU3
.LU2:
	mov r0, r4
	ldr r0, [r0]
	b .LU0
.LU3:
	b .LU4
.LU4:
	mov r1, r4
	ldr r1, [r1]
	movw r2, #4
	mul r2, r2, r1
	mov r1, r4
	str r2, [r1]
	sub r0, r0, #1
	mov r1, r0
	mov r0, r4
	bl multBy4xTimes
	mov r0, r4
	ldr r0, [r0]
	b .LU0
.LU0:
	pop {r4}
	pop {fp, pc}
	.size multBy4xTimes, .-multBy4xTimes
	.align 2
	.global divideBy8
divideBy8:
.LU6:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r4, r0
	mov r0, r4
	ldr r0, [r0]
	movw r1, #2
	bl __aeabi_idiv
	mov r1, r4
	str r0, [r1]
	mov r0, r4
	ldr r0, [r0]
	movw r1, #2
	bl __aeabi_idiv
	mov r1, r4
	str r0, [r1]
	mov r0, r4
	ldr r0, [r0]
	movw r1, #2
	bl __aeabi_idiv
	mov r1, r0
	mov r0, r4
	str r1, [r0]
	b .LU5
.LU5:
	pop {r4}
	pop {fp, pc}
	.size divideBy8, .-divideBy8
	.align 2
	.global main
main:
.LU8:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6}
	movw r0, #4
	bl malloc
	mov r4, r0
	movw r1, #16960
	movt r1, #15
	movw r0, #:lower16:end
	movt r0, #:upper16:end
	str r1, [r0]
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r5, #:lower16:.read_scratch
	movt r5, #:upper16:.read_scratch
	ldr r5, [r5]
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r0, #:lower16:.read_scratch
	movt r0, #:upper16:.read_scratch
	ldr r0, [r0]
	movw r1, #:lower16:interval
	movt r1, #:upper16:interval
	str r0, [r1]
	mov r1, r5
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r0, #:lower16:interval
	movt r0, #:upper16:interval
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r0, #0
	movw r1, #0
	b .LU9
.LU9:
	mov r5, r1
	movw ip, #:lower16:end
	movt ip, #:upper16:end
	ldr ip, [ip]
	movw r1, #0
	mov r2, r0
	movw r0, #0
	mov lr, #0
	movw r3, #0
	cmp r3, ip
	movle lr, #1
	cmp lr, #1
	beq .LU11
	b .LU12
.LU11:
	mov r0, r1
	add r6, r0, #1
	mov r0, r4
	str r6, [r0]
	movw r1, #2
	mov r0, r4
	bl multBy4xTimes
	mov r0, r4
	bl divideBy8
	movw r0, #:lower16:interval
	movt r0, #:upper16:interval
	ldr r0, [r0]
	sub r0, r0, #1
	mov r1, #0
	cmp r0, #0
	movle r1, #1
	cmp r1, #1
	beq .LU13
	b .LU14
.LU13:
	movw r0, #1
	b .LU15
.LU14:
	b .LU15
.LU15:
	add ip, r6, r0
	movw r3, #:lower16:end
	movt r3, #:upper16:end
	ldr r3, [r3]
	mov r1, ip
	movw r2, #5376
	movt r2, #609
	mov r0, ip
	mov lr, #0
	cmp ip, r3
	movle lr, #1
	cmp lr, #1
	beq .LU11
	b .LU12
.LU12:
	mov r6, r2
	mov r2, r0
	add r3, r5, #1
	mov r0, r6
	mov r1, r3
	mov ip, #0
	cmp r3, #50
	movlt ip, #1
	cmp ip, #1
	beq .LU9
	b .LU10
.LU10:
	mov r1, r2
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r1, r6
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU7
.LU7:
	movw r0, #0
	pop {r4, r5, r6}
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

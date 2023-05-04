	.arch armv7-a
	.text
	.align 2
	.global wait
wait:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	mov r2, r0
	mov r0, r2
	mov r1, #0
	cmp r2, #0
	movgt r1, #1
	cmp r1, #1
	beq .LU2
	b .LU3
.LU2:
	sub r2, r0, #1
	mov r0, r2
	mov r1, #0
	cmp r2, #0
	movgt r1, #1
	cmp r1, #1
	beq .LU2
	b .LU3
.LU3:
	b .LU0
.LU0:
	movw r0, #0
	pop {fp, pc}
	.size wait, .-wait
	.align 2
	.global power
power:
.LU5:
	push {fp, lr}
	add fp, sp, #4
	mov lr, r0
	mov ip, r1
	mov r3, ip
	movw r2, #1
	movw r0, #1
	mov r1, #0
	cmp ip, #0
	movgt r1, #1
	cmp r1, #1
	beq .LU6
	b .LU7
.LU6:
	mov r0, r3
	mov r1, r2
	mul r1, r1, lr
	sub ip, r0, #1
	mov r3, ip
	mov r2, r1
	mov r0, r1
	mov r1, #0
	cmp ip, #0
	movgt r1, #1
	cmp r1, #1
	beq .LU6
	b .LU7
.LU7:
	b .LU4
.LU4:
	pop {fp, pc}
	.size power, .-power
	.align 2
	.global recursiveDecimalSum
recursiveDecimalSum:
.LU9:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6}
	mov r5, r0
	mov r6, r1
	mov r4, r2
	mov r0, #0
	cmp r5, #0
	movgt r0, #1
	cmp r0, #1
	beq .LU10
	b .LU11
.LU10:
	movw r1, #10
	mov r0, r5
	bl __aeabi_idiv
	movw r1, #10
	mul r0, r0, r1
	sub r1, r5, r0
	mov r0, #0
	cmp r1, #1
	moveq r0, #1
	cmp r0, #1
	beq .LU12
	b .LU13
.LU12:
	mov r1, r4
	movw r0, #2
	bl power
	add r0, r6, r0
	b .LU14
.LU13:
	mov r0, r6
	b .LU14
.LU14:
	mov r6, r0
	movw r1, #10
	mov r0, r5
	bl __aeabi_idiv
	mov r3, r0
	add r0, r4, #1
	mov r2, r0
	mov r1, r6
	mov r0, r3
	bl recursiveDecimalSum
	b .LU8
.LU11:
	b .LU15
.LU15:
	mov r0, r6
	b .LU8
.LU8:
	pop {r4, r5, r6}
	pop {fp, pc}
	.size recursiveDecimalSum, .-recursiveDecimalSum
	.align 2
	.global convertToDecimal
convertToDecimal:
.LU17:
	push {fp, lr}
	add fp, sp, #4
	movw r2, #0
	movw r1, #0
	bl recursiveDecimalSum
	b .LU16
.LU16:
	pop {fp, pc}
	.size convertToDecimal, .-convertToDecimal
	.align 2
	.global main
main:
.LU19:
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
	bl convertToDecimal
	mov r4, r0
	mul r0, r4, r4
	mov r1, r0
	mov r2, #0
	cmp r0, #0
	movgt r2, #1
	cmp r2, #1
	beq .LU20
	b .LU21
.LU20:
	mov r5, r1
	mov r0, r5
	bl wait
	sub r2, r5, #1
	mov r1, r2
	mov r0, #0
	cmp r2, #0
	movgt r0, #1
	cmp r0, #1
	beq .LU20
	b .LU21
.LU21:
	mov r1, r4
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU18
.LU18:
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

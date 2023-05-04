	.arch armv7-a
	.comm	global3,4,4
	.comm	global2,4,4
	.comm	global1,4,4

	.text
	.align 2
	.global constantFolding
constantFolding:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	b .LU0
.LU0:
	movw r0, #226
	pop {fp, pc}
	.size constantFolding, .-constantFolding
	.align 2
	.global constantPropagation
constantPropagation:
.LU3:
	push {fp, lr}
	add fp, sp, #4
	b .LU2
.LU2:
	movw r0, #35615
	movt r0, #65147
	pop {fp, pc}
	.size constantPropagation, .-constantPropagation
	.align 2
	.global deadCodeElimination
deadCodeElimination:
.LU5:
	push {fp, lr}
	add fp, sp, #4
	movw r1, #11
	movw r0, #:lower16:global1
	movt r0, #:upper16:global1
	str r1, [r0]
	movw r1, #5
	movw r0, #:lower16:global1
	movt r0, #:upper16:global1
	str r1, [r0]
	movw r1, #9
	movw r0, #:lower16:global1
	movt r0, #:upper16:global1
	str r1, [r0]
	b .LU4
.LU4:
	movw r0, #38
	pop {fp, pc}
	.size deadCodeElimination, .-deadCodeElimination
	.align 2
	.global sum
sum:
.LU7:
	push {fp, lr}
	add fp, sp, #4
	mov r1, r0
	mov r2, r1
	movw ip, #0
	movw r3, #0
	mov r0, #0
	cmp r1, #0
	movgt r0, #1
	cmp r0, #1
	beq .LU8
	b .LU9
.LU8:
	mov r0, r2
	mov r1, ip
	add r1, r1, r0
	sub r0, r0, #1
	mov r2, r0
	mov ip, r1
	mov r3, r1
	mov r1, #0
	cmp r0, #0
	movgt r1, #1
	cmp r1, #1
	beq .LU8
	b .LU9
.LU9:
	mov r0, r3
	b .LU6
.LU6:
	pop {fp, pc}
	.size sum, .-sum
	.align 2
	.global doesntModifyGlobals
doesntModifyGlobals:
.LU11:
	push {fp, lr}
	add fp, sp, #4
	b .LU10
.LU10:
	movw r0, #3
	pop {fp, pc}
	.size doesntModifyGlobals, .-doesntModifyGlobals
	.align 2
	.global interProceduralOptimization
interProceduralOptimization:
.LU13:
	push {fp, lr}
	add fp, sp, #4
	movw r1, #1
	movw r0, #:lower16:global1
	movt r0, #:upper16:global1
	str r1, [r0]
	movw r1, #0
	movw r0, #:lower16:global2
	movt r0, #:upper16:global2
	str r1, [r0]
	movw r1, #0
	movw r0, #:lower16:global3
	movt r0, #:upper16:global3
	str r1, [r0]
	movw r0, #100
	bl sum
	mov r2, r0
	movw r1, #:lower16:global1
	movt r1, #:upper16:global1
	ldr r1, [r1]
	mov r0, #0
	cmp r1, #1
	moveq r0, #1
	cmp r0, #1
	beq .LU14
	b .LU15
.LU14:
	movw r0, #10000
	bl sum
	b .LU22
.LU15:
	movw r0, #:lower16:global2
	movt r0, #:upper16:global2
	ldr r0, [r0]
	mov r1, #0
	cmp r0, #2
	moveq r1, #1
	cmp r1, #1
	beq .LU16
	b .LU17
.LU16:
	movw r0, #20000
	bl sum
	b .LU18
.LU17:
	mov r0, r2
	b .LU18
.LU18:
	mov r2, r0
	movw r1, #:lower16:global3
	movt r1, #:upper16:global3
	ldr r1, [r1]
	mov r0, #0
	cmp r1, #3
	moveq r0, #1
	cmp r0, #1
	beq .LU19
	b .LU20
.LU19:
	movw r0, #30000
	bl sum
	b .LU21
.LU20:
	mov r0, r2
	b .LU21
.LU21:
	b .LU22
.LU22:
	b .LU12
.LU12:
	pop {fp, pc}
	.size interProceduralOptimization, .-interProceduralOptimization
	.align 2
	.global commonSubexpressionElimination
commonSubexpressionElimination:
.LU24:
	push {fp, lr}
	add fp, sp, #4
	b .LU23
.LU23:
	movw r0, #16740
	movt r0, #65535
	pop {fp, pc}
	.size commonSubexpressionElimination, .-commonSubexpressionElimination
	.align 2
	.global hoisting
hoisting:
.LU26:
	push {fp, lr}
	add fp, sp, #4
	movw r1, #0
	b .LU27
.LU27:
	mov r0, r1
	add r0, r0, #1
	mov r1, r0
	mov r2, #0
	movw r3, #16960
	movt r3, #15
	cmp r0, r3
	movlt r2, #1
	cmp r2, #1
	beq .LU27
	b .LU28
.LU28:
	b .LU25
.LU25:
	movw r0, #2
	pop {fp, pc}
	.size hoisting, .-hoisting
	.align 2
	.global doubleIf
doubleIf:
.LU30:
	push {fp, lr}
	add fp, sp, #4
	b .LU31
.LU31:
	b .LU33
.LU33:
	b .LU35
.LU35:
	b .LU36
.LU36:
	b .LU29
.LU29:
	movw r0, #50
	pop {fp, pc}
	.size doubleIf, .-doubleIf
	.align 2
	.global integerDivide
integerDivide:
.LU38:
	push {fp, lr}
	add fp, sp, #4
	b .LU37
.LU37:
	movw r0, #736
	pop {fp, pc}
	.size integerDivide, .-integerDivide
	.align 2
	.global association
association:
.LU40:
	push {fp, lr}
	add fp, sp, #4
	b .LU39
.LU39:
	movw r0, #10
	pop {fp, pc}
	.size association, .-association
	.align 2
	.global tailRecursionHelper
tailRecursionHelper:
.LU42:
	push {fp, lr}
	add fp, sp, #4
	mov r2, r0
	mov r0, r1
	mov r1, #0
	cmp r2, #0
	moveq r1, #1
	cmp r1, #1
	beq .LU43
	b .LU44
.LU43:
	b .LU41
.LU44:
	sub r3, r2, #1
	add r0, r0, r2
	mov r1, r0
	mov r0, r3
	bl tailRecursionHelper
	b .LU41
.LU41:
	pop {fp, pc}
	.size tailRecursionHelper, .-tailRecursionHelper
	.align 2
	.global tailRecursion
tailRecursion:
.LU46:
	push {fp, lr}
	add fp, sp, #4
	movw r1, #0
	bl tailRecursionHelper
	b .LU45
.LU45:
	pop {fp, pc}
	.size tailRecursion, .-tailRecursion
	.align 2
	.global unswitching
unswitching:
.LU48:
	push {fp, lr}
	add fp, sp, #4
	movw r2, #1
	b .LU49
.LU49:
	mov r0, r2
	b .LU51
.LU51:
	add r0, r0, #1
	b .LU53
.LU53:
	mov r2, r0
	mov r1, #0
	movw r3, #16960
	movt r3, #15
	cmp r0, r3
	movlt r1, #1
	cmp r1, #1
	beq .LU49
	b .LU50
.LU50:
	b .LU47
.LU47:
	pop {fp, pc}
	.size unswitching, .-unswitching
	.align 2
	.global randomCalculation
randomCalculation:
.LU55:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r4, r0
	movw r1, #0
	movw r3, #0
	movw r2, #0
	mov r0, #0
	movw ip, #0
	cmp ip, r4
	movlt r0, #1
	cmp r0, #1
	beq .LU56
	b .LU57
.LU56:
	mov r0, r3
	add r5, r0, #19
	movw r0, #2
	mul r0, r1, r0
	movw r1, #2
	bl __aeabi_idiv
	movw r1, #3
	mul r0, r1, r0
	movw r1, #3
	bl __aeabi_idiv
	movw r1, #4
	mul r0, r0, r1
	movw r1, #4
	bl __aeabi_idiv
	add ip, r0, #1
	mov r1, ip
	mov r3, r5
	mov r2, r5
	mov r0, #0
	cmp ip, r4
	movlt r0, #1
	cmp r0, #1
	beq .LU56
	b .LU57
.LU57:
	mov r0, r2
	b .LU54
.LU54:
	pop {r4, r5}
	pop {fp, pc}
	.size randomCalculation, .-randomCalculation
	.align 2
	.global iterativeFibonacci
iterativeFibonacci:
.LU59:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r2, r0
	movw r3, #0
	movw r0, #65535
	movt r0, #65535
	movw r4, #1
	movw r1, #1
	mov lr, #0
	movw ip, #0
	cmp ip, r2
	movlt lr, #1
	cmp lr, #1
	beq .LU60
	b .LU61
.LU60:
	mov r1, r0
	mov r0, r4
	add r1, r0, r1
	add ip, r3, #1
	mov r3, ip
	mov r4, r1
	mov lr, #0
	cmp ip, r2
	movlt lr, #1
	cmp lr, #1
	beq .LU60
	b .LU61
.LU61:
	mov r0, r1
	b .LU58
.LU58:
	pop {r4}
	pop {fp, pc}
	.size iterativeFibonacci, .-iterativeFibonacci
	.align 2
	.global recursiveFibonacci
recursiveFibonacci:
.LU63:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r4, r0
	mov r1, #0
	cmp r4, #0
	movle r1, #1
	mov r0, #0
	cmp r4, #1
	moveq r0, #1
	orr r0, r1, r0
	cmp r0, #1
	beq .LU64
	b .LU65
.LU64:
	mov r0, r4
	b .LU62
.LU65:
	sub r0, r4, #1
	bl recursiveFibonacci
	mov r5, r0
	sub r0, r4, #2
	bl recursiveFibonacci
	add r0, r5, r0
	b .LU62
.LU62:
	pop {r4, r5}
	pop {fp, pc}
	.size recursiveFibonacci, .-recursiveFibonacci
	.align 2
	.global main
main:
.LU67:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r5, #:lower16:.read_scratch
	movt r5, #:upper16:.read_scratch
	ldr r5, [r5]
	movw r2, #1
	mov r0, #0
	movw r1, #1
	cmp r1, r5
	movlt r0, #1
	cmp r0, #1
	beq .LU68
	b .LU69
.LU68:
	mov r4, r2
	bl constantFolding
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	bl constantPropagation
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	bl deadCodeElimination
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	bl interProceduralOptimization
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	bl commonSubexpressionElimination
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	bl hoisting
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	bl doubleIf
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	bl integerDivide
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	bl association
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #1000
	mov r0, r5
	bl __aeabi_idiv
	bl tailRecursion
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	bl unswitching
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, r5
	bl randomCalculation
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #5
	mov r0, r5
	bl __aeabi_idiv
	bl iterativeFibonacci
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #1000
	mov r0, r5
	bl __aeabi_idiv
	bl recursiveFibonacci
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	add r0, r4, #1
	mov r2, r0
	mov r1, #0
	cmp r0, r5
	movlt r1, #1
	cmp r1, #1
	beq .LU68
	b .LU69
.LU69:
	movw r1, #9999
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU66
.LU66:
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

	.arch armv7-a
	.text
	.align 2
	.global cleanBoard
cleanBoard:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	mov r1, r0
	movw r2, #0
	str r2, [r1]
	add r1, r0, #4
	movw r2, #0
	str r2, [r1]
	add r1, r0, #8
	movw r2, #0
	str r2, [r1]
	add r1, r0, #12
	movw r2, #0
	str r2, [r1]
	add r1, r0, #16
	movw r2, #0
	str r2, [r1]
	add r1, r0, #20
	movw r2, #0
	str r2, [r1]
	add r1, r0, #24
	movw r2, #0
	str r2, [r1]
	add r1, r0, #28
	movw r2, #0
	str r2, [r1]
	add r0, r0, #32
	movw r1, #0
	str r1, [r0]
	b .LU0
.LU0:
	pop {fp, pc}
	.size cleanBoard, .-cleanBoard
	.align 2
	.global printBoard
printBoard:
.LU3:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r4, r0
	mov r0, r4
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	add r0, r4, #4
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	add r0, r4, #8
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	add r0, r4, #12
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	add r0, r4, #16
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	add r0, r4, #20
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	add r0, r4, #24
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	add r0, r4, #28
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	add r0, r4, #32
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU2
.LU2:
	pop {r4}
	pop {fp, pc}
	.size printBoard, .-printBoard
	.align 2
	.global printMoveBoard
printMoveBoard:
.LU5:
	push {fp, lr}
	add fp, sp, #4
	movw r1, #123
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #456
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #789
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU4
.LU4:
	pop {fp, pc}
	.size printMoveBoard, .-printMoveBoard
	.align 2
	.global placePiece
placePiece:
.LU7:
	push {fp, lr}
	add fp, sp, #4
	mov r3, r0
	mov r0, r1
	mov r1, r2
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU8
	b .LU9
.LU8:
	mov r1, r3
	str r0, [r1]
	b .LU34
.LU9:
	mov r2, #0
	cmp r1, #2
	moveq r2, #1
	cmp r2, #1
	beq .LU10
	b .LU11
.LU10:
	add r1, r3, #4
	str r0, [r1]
	b .LU33
.LU11:
	mov r2, #0
	cmp r1, #3
	moveq r2, #1
	cmp r2, #1
	beq .LU12
	b .LU13
.LU12:
	add r1, r3, #8
	str r0, [r1]
	b .LU32
.LU13:
	mov r2, #0
	cmp r1, #4
	moveq r2, #1
	cmp r2, #1
	beq .LU14
	b .LU15
.LU14:
	add r1, r3, #12
	str r0, [r1]
	b .LU31
.LU15:
	mov r2, #0
	cmp r1, #5
	moveq r2, #1
	cmp r2, #1
	beq .LU16
	b .LU17
.LU16:
	add r1, r3, #16
	str r0, [r1]
	b .LU30
.LU17:
	mov r2, #0
	cmp r1, #6
	moveq r2, #1
	cmp r2, #1
	beq .LU18
	b .LU19
.LU18:
	add r1, r3, #20
	str r0, [r1]
	b .LU29
.LU19:
	mov r2, #0
	cmp r1, #7
	moveq r2, #1
	cmp r2, #1
	beq .LU20
	b .LU21
.LU20:
	add r1, r3, #24
	str r0, [r1]
	b .LU28
.LU21:
	mov r2, #0
	cmp r1, #8
	moveq r2, #1
	cmp r2, #1
	beq .LU22
	b .LU23
.LU22:
	add r1, r3, #28
	str r0, [r1]
	b .LU27
.LU23:
	mov r2, #0
	cmp r1, #9
	moveq r2, #1
	cmp r2, #1
	beq .LU24
	b .LU25
.LU24:
	add r1, r3, #32
	str r0, [r1]
	b .LU26
.LU25:
	b .LU26
.LU26:
	b .LU27
.LU27:
	b .LU28
.LU28:
	b .LU29
.LU29:
	b .LU30
.LU30:
	b .LU31
.LU31:
	b .LU32
.LU32:
	b .LU33
.LU33:
	b .LU34
.LU34:
	b .LU6
.LU6:
	pop {fp, pc}
	.size placePiece, .-placePiece
	.align 2
	.global checkWinner
checkWinner:
.LU36:
	push {fp, lr}
	add fp, sp, #4
	mov r1, r0
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU37
	b .LU38
.LU37:
	add r1, r0, #4
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU39
	b .LU40
.LU39:
	add r1, r0, #8
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #1
	moveq r1, #1
	cmp r1, #1
	beq .LU41
	b .LU42
.LU41:
	movw r0, #0
	b .LU35
.LU42:
	b .LU43
.LU43:
	b .LU44
.LU40:
	b .LU44
.LU44:
	b .LU45
.LU38:
	b .LU45
.LU45:
	mov r1, r0
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #2
	moveq r2, #1
	cmp r2, #1
	beq .LU46
	b .LU47
.LU46:
	add r1, r0, #4
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #2
	moveq r1, #1
	cmp r1, #1
	beq .LU48
	b .LU49
.LU48:
	add r1, r0, #8
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #2
	moveq r1, #1
	cmp r1, #1
	beq .LU50
	b .LU51
.LU50:
	movw r0, #1
	b .LU35
.LU51:
	b .LU52
.LU52:
	b .LU53
.LU49:
	b .LU53
.LU53:
	b .LU54
.LU47:
	b .LU54
.LU54:
	add r1, r0, #12
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU55
	b .LU56
.LU55:
	add r1, r0, #16
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #1
	moveq r1, #1
	cmp r1, #1
	beq .LU57
	b .LU58
.LU57:
	add r1, r0, #20
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #1
	moveq r1, #1
	cmp r1, #1
	beq .LU59
	b .LU60
.LU59:
	movw r0, #0
	b .LU35
.LU60:
	b .LU61
.LU61:
	b .LU62
.LU58:
	b .LU62
.LU62:
	b .LU63
.LU56:
	b .LU63
.LU63:
	add r1, r0, #12
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #2
	moveq r1, #1
	cmp r1, #1
	beq .LU64
	b .LU65
.LU64:
	add r1, r0, #16
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #2
	moveq r1, #1
	cmp r1, #1
	beq .LU66
	b .LU67
.LU66:
	add r1, r0, #20
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #2
	moveq r2, #1
	cmp r2, #1
	beq .LU68
	b .LU69
.LU68:
	movw r0, #1
	b .LU35
.LU69:
	b .LU70
.LU70:
	b .LU71
.LU67:
	b .LU71
.LU71:
	b .LU72
.LU65:
	b .LU72
.LU72:
	add r1, r0, #24
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU73
	b .LU74
.LU73:
	add r1, r0, #28
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #1
	moveq r1, #1
	cmp r1, #1
	beq .LU75
	b .LU76
.LU75:
	add r1, r0, #32
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU77
	b .LU78
.LU77:
	movw r0, #0
	b .LU35
.LU78:
	b .LU79
.LU79:
	b .LU80
.LU76:
	b .LU80
.LU80:
	b .LU81
.LU74:
	b .LU81
.LU81:
	add r1, r0, #24
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #2
	moveq r2, #1
	cmp r2, #1
	beq .LU82
	b .LU83
.LU82:
	add r1, r0, #28
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #2
	moveq r1, #1
	cmp r1, #1
	beq .LU84
	b .LU85
.LU84:
	add r1, r0, #32
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #2
	moveq r2, #1
	cmp r2, #1
	beq .LU86
	b .LU87
.LU86:
	movw r0, #1
	b .LU35
.LU87:
	b .LU88
.LU88:
	b .LU89
.LU85:
	b .LU89
.LU89:
	b .LU90
.LU83:
	b .LU90
.LU90:
	mov r1, r0
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU91
	b .LU92
.LU91:
	add r1, r0, #12
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU93
	b .LU94
.LU93:
	add r1, r0, #24
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU95
	b .LU96
.LU95:
	movw r0, #0
	b .LU35
.LU96:
	b .LU97
.LU97:
	b .LU98
.LU94:
	b .LU98
.LU98:
	b .LU99
.LU92:
	b .LU99
.LU99:
	mov r1, r0
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #2
	moveq r1, #1
	cmp r1, #1
	beq .LU100
	b .LU101
.LU100:
	add r1, r0, #12
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #2
	moveq r1, #1
	cmp r1, #1
	beq .LU102
	b .LU103
.LU102:
	add r1, r0, #24
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #2
	moveq r2, #1
	cmp r2, #1
	beq .LU104
	b .LU105
.LU104:
	movw r0, #1
	b .LU35
.LU105:
	b .LU106
.LU106:
	b .LU107
.LU103:
	b .LU107
.LU107:
	b .LU108
.LU101:
	b .LU108
.LU108:
	add r1, r0, #4
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU109
	b .LU110
.LU109:
	add r1, r0, #16
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU111
	b .LU112
.LU111:
	add r1, r0, #28
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU113
	b .LU114
.LU113:
	movw r0, #0
	b .LU35
.LU114:
	b .LU115
.LU115:
	b .LU116
.LU112:
	b .LU116
.LU116:
	b .LU117
.LU110:
	b .LU117
.LU117:
	add r1, r0, #4
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #2
	moveq r1, #1
	cmp r1, #1
	beq .LU118
	b .LU119
.LU118:
	add r1, r0, #16
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #2
	moveq r1, #1
	cmp r1, #1
	beq .LU120
	b .LU121
.LU120:
	add r1, r0, #28
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #2
	moveq r2, #1
	cmp r2, #1
	beq .LU122
	b .LU123
.LU122:
	movw r0, #1
	b .LU35
.LU123:
	b .LU124
.LU124:
	b .LU125
.LU121:
	b .LU125
.LU125:
	b .LU126
.LU119:
	b .LU126
.LU126:
	add r1, r0, #8
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU127
	b .LU128
.LU127:
	add r1, r0, #20
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #1
	moveq r1, #1
	cmp r1, #1
	beq .LU129
	b .LU130
.LU129:
	add r1, r0, #32
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #1
	moveq r2, #1
	cmp r2, #1
	beq .LU131
	b .LU132
.LU131:
	movw r0, #0
	b .LU35
.LU132:
	b .LU133
.LU133:
	b .LU134
.LU130:
	b .LU134
.LU134:
	b .LU135
.LU128:
	b .LU135
.LU135:
	add r1, r0, #8
	ldr r2, [r1]
	mov r1, #0
	cmp r2, #2
	moveq r1, #1
	cmp r1, #1
	beq .LU136
	b .LU137
.LU136:
	add r1, r0, #20
	ldr r1, [r1]
	mov r2, #0
	cmp r1, #2
	moveq r2, #1
	cmp r2, #1
	beq .LU138
	b .LU139
.LU138:
	add r0, r0, #32
	ldr r0, [r0]
	mov r1, #0
	cmp r0, #2
	moveq r1, #1
	cmp r1, #1
	beq .LU140
	b .LU141
.LU140:
	movw r0, #1
	b .LU35
.LU141:
	b .LU142
.LU142:
	b .LU143
.LU139:
	b .LU143
.LU143:
	b .LU144
.LU137:
	b .LU144
.LU144:
	movw r0, #65535
	movt r0, #65535
	b .LU35
.LU35:
	pop {fp, pc}
	.size checkWinner, .-checkWinner
	.align 2
	.global main
main:
.LU146:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6}
	movw r0, #36
	bl malloc
	mov r4, r0
	mov r0, r4
	bl cleanBoard
	movw r2, #0
	movw r1, #0
	b .LU147
.LU147:
	mov r5, r2
	mov r6, r1
	mov r0, r4
	bl printBoard
	mov r0, #0
	cmp r6, #0
	moveq r0, #1
	cmp r0, #1
	beq .LU149
	b .LU150
.LU149:
	add r6, r6, #1
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r0, #:lower16:.read_scratch
	movt r0, #:upper16:.read_scratch
	ldr r0, [r0]
	mov r2, r0
	movw r1, #1
	mov r0, r4
	bl placePiece
	mov r0, r6
	b .LU151
.LU150:
	sub r6, r6, #1
	movw r1, #:lower16:.read_scratch
	movt r1, #:upper16:.read_scratch
	movw r0, #:lower16:.READ_FMT
	movt r0, #:upper16:.READ_FMT
	bl scanf
	movw r0, #:lower16:.read_scratch
	movt r0, #:upper16:.read_scratch
	ldr r0, [r0]
	mov r2, r0
	movw r1, #2
	mov r0, r4
	bl placePiece
	mov r0, r6
	b .LU151
.LU151:
	mov r6, r0
	mov r0, r4
	bl checkWinner
	add r1, r5, #1
	mov r3, #0
	cmp r0, #0
	movlt r3, #1
	mov ip, #0
	cmp r1, #8
	movne ip, #1
	mov r2, r1
	mov r1, r6
	and r3, r3, ip
	cmp r3, #1
	beq .LU147
	b .LU148
.LU148:
	add r0, r0, #1
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU145
.LU145:
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

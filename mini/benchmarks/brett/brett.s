	.arch armv7-a
	.comm	gs1,4,4
	.comm	gb1,4,4
	.comm	counter,4,4
	.comm	gi1,4,4

	.text
	.align 2
	.global printgroup
printgroup:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r4, r0
	movw r1, #1
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	movw r1, #1
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	movw r1, #1
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	mov r1, r4
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU0
.LU0:
	pop {r4}
	pop {fp, pc}
	.size printgroup, .-printgroup
	.align 2
	.global setcounter
setcounter:
.LU3:
	push {fp, lr}
	add fp, sp, #4
	movw r1, #:lower16:counter
	movt r1, #:upper16:counter
	str r0, [r1]
	b .LU2
.LU2:
	movw r0, #1
	pop {fp, pc}
	.size setcounter, .-setcounter
	.align 2
	.global takealltypes
takealltypes:
.LU5:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r3, r0
	mov r4, r1
	mov r5, r2
	mov r0, #0
	cmp r3, #3
	moveq r0, #1
	cmp r0, #1
	beq .LU6
	b .LU7
.LU6:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU8
.LU7:
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU8
.LU8:
	mov r0, r4
	cmp r0, #1
	beq .LU9
	b .LU10
.LU9:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU11
.LU10:
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU11
.LU11:
	mov r0, r5
	ldr r0, [r0]
	cmp r0, #1
	beq .LU12
	b .LU13
.LU12:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU14
.LU13:
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU14
.LU14:
	b .LU4
.LU4:
	pop {r4, r5}
	pop {fp, pc}
	.size takealltypes, .-takealltypes
	.align 2
	.global tonofargs
tonofargs:
.LU16:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6, r7}
	mov r0, r1
	mov r0, r2
	mov r0, r3
	ldr r7, [fp,#4]
	ldr r6, [fp,#8]
	ldr r5, [fp,#12]
	ldr r4, [fp,#16]
	mov r0, #0
	cmp r7, #5
	moveq r0, #1
	cmp r0, #1
	beq .LU17
	b .LU18
.LU17:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU19
.LU18:
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	mov r1, r7
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU19
.LU19:
	mov r0, #0
	cmp r6, #6
	moveq r0, #1
	cmp r0, #1
	beq .LU20
	b .LU21
.LU20:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU22
.LU21:
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	mov r1, r6
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU22
.LU22:
	mov r0, #0
	cmp r5, #7
	moveq r0, #1
	cmp r0, #1
	beq .LU23
	b .LU24
.LU23:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU25
.LU24:
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	mov r1, r5
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU25
.LU25:
	mov r0, #0
	cmp r4, #8
	moveq r0, #1
	cmp r0, #1
	beq .LU26
	b .LU27
.LU26:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU28
.LU27:
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	mov r1, r4
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU28
.LU28:
	b .LU15
.LU15:
	pop {r4, r5, r6, r7}
	pop {fp, pc}
	.size tonofargs, .-tonofargs
	.align 2
	.global returnint
returnint:
.LU30:
	push {fp, lr}
	add fp, sp, #4
	b .LU29
.LU29:
	pop {fp, pc}
	.size returnint, .-returnint
	.align 2
	.global returnbool
returnbool:
.LU32:
	push {fp, lr}
	add fp, sp, #4
	b .LU31
.LU31:
	pop {fp, pc}
	.size returnbool, .-returnbool
	.align 2
	.global returnstruct
returnstruct:
.LU34:
	push {fp, lr}
	add fp, sp, #4
	b .LU33
.LU33:
	pop {fp, pc}
	.size returnstruct, .-returnstruct
	.align 2
	.global main
main:
.LU36:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	sub sp, sp, #16
	movw r0, #0
	movw r1, #:lower16:counter
	movt r1, #:upper16:counter
	str r0, [r1]
	movw r0, #1
	bl printgroup
	b .LU38
.LU38:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU39
.LU39:
	b .LU41
.LU41:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU42
.LU42:
	b .LU44
.LU44:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU45
.LU45:
	b .LU46
.LU46:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU48
.LU48:
	movw r1, #0
	movw r0, #:lower16:counter
	movt r0, #:upper16:counter
	str r1, [r0]
	movw r0, #2
	bl printgroup
	b .LU49
.LU49:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU51
.LU51:
	b .LU52
.LU52:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU54
.LU54:
	b .LU55
.LU55:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU57
.LU57:
	b .LU59
.LU59:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU60
.LU60:
	movw r0, #3
	bl printgroup
	b .LU61
.LU61:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU63
.LU63:
	b .LU64
.LU64:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU66
.LU66:
	b .LU68
.LU68:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU69
.LU69:
	b .LU71
.LU71:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU72
.LU72:
	b .LU74
.LU74:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU75
.LU75:
	b .LU76
.LU76:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU78
.LU78:
	b .LU79
.LU79:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU81
.LU81:
	b .LU83
.LU83:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU84
.LU84:
	b .LU86
.LU86:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU87
.LU87:
	b .LU88
.LU88:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU90
.LU90:
	b .LU91
.LU91:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU93
.LU93:
	movw r0, #4
	bl printgroup
	b .LU94
.LU94:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU96
.LU96:
	b .LU97
.LU97:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU99
.LU99:
	b .LU100
.LU100:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU102
.LU102:
	b .LU103
.LU103:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU105
.LU105:
	b .LU106
.LU106:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU108
.LU108:
	movw r0, #5
	bl printgroup
	b .LU109
.LU109:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU111
.LU111:
	b .LU112
.LU112:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU114
.LU114:
	b .LU115
.LU115:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU117
.LU117:
	b .LU119
.LU119:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU120
.LU120:
	b .LU122
.LU122:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU123
.LU123:
	b .LU124
.LU124:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU126
.LU126:
	b .LU128
.LU128:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU129
.LU129:
	movw r0, #6
	bl printgroup
	b .LU130
.LU130:
	b .LU133
.LU133:
	b .LU134
.LU134:
	b .LU131
.LU131:
	b .LU135
.LU135:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU137
.LU137:
	movw r0, #7
	bl printgroup
	movw r0, #12
	bl malloc
	mov r4, r0
	add r0, r4, #8
	movw r1, #42
	str r1, [r0]
	mov r0, r4
	movw r1, #1
	str r1, [r0]
	add r0, r4, #8
	ldr r1, [r0]
	mov r0, #0
	cmp r1, #42
	moveq r0, #1
	cmp r0, #1
	beq .LU138
	b .LU139
.LU138:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU140
.LU139:
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	add r0, r4, #8
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU140
.LU140:
	mov r0, r4
	ldr r0, [r0]
	cmp r0, #1
	beq .LU141
	b .LU142
.LU141:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU143
.LU142:
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU143
.LU143:
	movw r0, #12
	bl malloc
	add r1, r4, #4
	str r0, [r1]
	add r0, r4, #4
	ldr r0, [r0]
	add r0, r0, #8
	movw r1, #13
	str r1, [r0]
	add r0, r4, #4
	ldr r0, [r0]
	movw r1, #0
	str r1, [r0]
	add r0, r4, #4
	ldr r0, [r0]
	add r0, r0, #8
	ldr r0, [r0]
	mov r1, #0
	cmp r0, #13
	moveq r1, #1
	cmp r1, #1
	beq .LU144
	b .LU145
.LU144:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU146
.LU145:
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	add r0, r4, #4
	ldr r0, [r0]
	add r0, r0, #8
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU146
.LU146:
	add r0, r4, #4
	ldr r0, [r0]
	ldr r0, [r0]
	eor r0, r0, #1
	cmp r0, #1
	beq .LU147
	b .LU148
.LU147:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU149
.LU148:
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU149
.LU149:
	mov r0, #0
	cmp r4, r4
	moveq r0, #1
	cmp r0, #1
	beq .LU150
	b .LU151
.LU150:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU152
.LU151:
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU152
.LU152:
	add r0, r4, #4
	ldr r1, [r0]
	mov r0, #0
	cmp r4, r1
	movne r0, #1
	cmp r0, #1
	beq .LU153
	b .LU154
.LU153:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU155
.LU154:
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU155
.LU155:
	add r0, r4, #4
	ldr r0, [r0]
	bl free
	mov r0, r4
	bl free
	movw r0, #8
	bl printgroup
	movw r1, #7
	movw r0, #:lower16:gi1
	movt r0, #:upper16:gi1
	str r1, [r0]
	movw r0, #:lower16:gi1
	movt r0, #:upper16:gi1
	ldr r0, [r0]
	mov r1, #0
	cmp r0, #7
	moveq r1, #1
	cmp r1, #1
	beq .LU156
	b .LU157
.LU156:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU158
.LU157:
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	movw r0, #:lower16:gi1
	movt r0, #:upper16:gi1
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU158
.LU158:
	movw r0, #1
	movw r1, #:lower16:gb1
	movt r1, #:upper16:gb1
	str r0, [r1]
	movw r0, #:lower16:gb1
	movt r0, #:upper16:gb1
	ldr r0, [r0]
	cmp r0, #1
	beq .LU159
	b .LU160
.LU159:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU161
.LU160:
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU161
.LU161:
	movw r0, #12
	bl malloc
	movw r1, #:lower16:gs1
	movt r1, #:upper16:gs1
	str r0, [r1]
	movw r0, #:lower16:gs1
	movt r0, #:upper16:gs1
	ldr r0, [r0]
	add r1, r0, #8
	movw r0, #34
	str r0, [r1]
	movw r0, #:lower16:gs1
	movt r0, #:upper16:gs1
	ldr r0, [r0]
	movw r1, #0
	str r1, [r0]
	movw r0, #:lower16:gs1
	movt r0, #:upper16:gs1
	ldr r0, [r0]
	add r0, r0, #8
	ldr r0, [r0]
	mov r1, #0
	cmp r0, #34
	moveq r1, #1
	cmp r1, #1
	beq .LU162
	b .LU163
.LU162:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU164
.LU163:
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	movw r0, #:lower16:gs1
	movt r0, #:upper16:gs1
	ldr r0, [r0]
	add r0, r0, #8
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU164
.LU164:
	movw r0, #:lower16:gs1
	movt r0, #:upper16:gs1
	ldr r0, [r0]
	ldr r0, [r0]
	eor r0, r0, #1
	cmp r0, #1
	beq .LU165
	b .LU166
.LU165:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU167
.LU166:
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU167
.LU167:
	movw r0, #12
	bl malloc
	movw r1, #:lower16:gs1
	movt r1, #:upper16:gs1
	ldr r1, [r1]
	add r1, r1, #4
	str r0, [r1]
	movw r0, #:lower16:gs1
	movt r0, #:upper16:gs1
	ldr r0, [r0]
	add r0, r0, #4
	ldr r0, [r0]
	add r0, r0, #8
	movw r1, #16
	str r1, [r0]
	movw r0, #:lower16:gs1
	movt r0, #:upper16:gs1
	ldr r0, [r0]
	add r0, r0, #4
	ldr r0, [r0]
	movw r1, #1
	str r1, [r0]
	movw r0, #:lower16:gs1
	movt r0, #:upper16:gs1
	ldr r0, [r0]
	add r0, r0, #4
	ldr r0, [r0]
	add r0, r0, #8
	ldr r0, [r0]
	mov r1, #0
	cmp r0, #16
	moveq r1, #1
	cmp r1, #1
	beq .LU168
	b .LU169
.LU168:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU170
.LU169:
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	movw r0, #:lower16:gs1
	movt r0, #:upper16:gs1
	ldr r0, [r0]
	add r0, r0, #4
	ldr r0, [r0]
	add r0, r0, #8
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU170
.LU170:
	movw r0, #:lower16:gs1
	movt r0, #:upper16:gs1
	ldr r0, [r0]
	add r0, r0, #4
	ldr r0, [r0]
	ldr r0, [r0]
	cmp r0, #1
	beq .LU171
	b .LU172
.LU171:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU173
.LU172:
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU173
.LU173:
	movw r0, #:lower16:gs1
	movt r0, #:upper16:gs1
	ldr r0, [r0]
	add r0, r0, #4
	ldr r0, [r0]
	bl free
	movw r0, #:lower16:gs1
	movt r0, #:upper16:gs1
	ldr r0, [r0]
	bl free
	movw r0, #9
	bl printgroup
	movw r0, #12
	bl malloc
	mov r1, r0
	movw r2, #1
	str r2, [r1]
	mov r2, r0
	movw r1, #1
	movw r0, #3
	bl takealltypes
	movw r1, #2
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r0, #8
	str r0, [sp,#12]
	movw r0, #7
	str r0, [sp,#8]
	movw r0, #6
	str r0, [sp,#4]
	movw r0, #5
	str r0, [sp]
	movw r3, #4
	movw r2, #3
	movw r1, #2
	movw r0, #1
	bl tonofargs
	movw r1, #3
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r0, #3
	bl returnint
	mov r4, r0
	mov r0, #0
	cmp r4, #3
	moveq r0, #1
	cmp r0, #1
	beq .LU174
	b .LU175
.LU174:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU176
.LU175:
	movw r1, #0
	movw r0, #:lower16:.PRINT_FMT
	movt r0, #:upper16:.PRINT_FMT
	bl printf
	mov r1, r4
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU176
.LU176:
	movw r0, #1
	bl returnbool
	cmp r0, #1
	beq .LU177
	b .LU178
.LU177:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU179
.LU178:
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU179
.LU179:
	movw r0, #12
	bl malloc
	mov r4, r0
	mov r0, r4
	bl returnstruct
	mov r1, #0
	cmp r4, r0
	moveq r1, #1
	cmp r1, #1
	beq .LU180
	b .LU181
.LU180:
	movw r1, #1
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU182
.LU181:
	movw r1, #0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU182
.LU182:
	movw r0, #10
	bl printgroup
	b .LU35
.LU35:
	movw r0, #0
	add sp, sp, #16
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

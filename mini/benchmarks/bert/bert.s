	.arch armv7-a
	.comm	a,4,4
	.comm	b,4,4
	.comm	i,4,4

	.text
	.align 2
	.global concatLists
concatLists:
.LU1:
	push {fp, lr}
	add fp, sp, #4
	mov r2, #0
	cmp r0, #0
	moveq r2, #1
	cmp r2, #1
	beq .LU2
	b .LU3
.LU2:
	mov r0, r1
	b .LU0
.LU3:
	b .LU4
.LU4:
	mov r2, r0
	ldr r3, [r2]
	mov ip, r0
	mov lr, r0
	mov r2, #0
	cmp r3, #0
	movne r2, #1
	cmp r2, #1
	beq .LU5
	b .LU6
.LU5:
	mov r2, ip
	ldr lr, [r2]
	mov r2, lr
	ldr r3, [r2]
	mov ip, lr
	mov r2, #0
	cmp r3, #0
	movne r2, #1
	cmp r2, #1
	beq .LU5
	b .LU6
.LU6:
	mov r2, lr
	str r1, [r2]
	b .LU0
.LU0:
	pop {fp, pc}
	.size concatLists, .-concatLists
	.align 2
	.global add
add:
.LU8:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r4, r0
	mov r5, r1
	movw r0, #8
	bl malloc
	add r1, r0, #4
	str r5, [r1]
	mov r1, r0
	str r4, [r1]
	b .LU7
.LU7:
	pop {r4, r5}
	pop {fp, pc}
	.size add, .-add
	.align 2
	.global size
size:
.LU10:
	push {fp, lr}
	add fp, sp, #4
	mov r1, #0
	cmp r0, #0
	moveq r1, #1
	cmp r1, #1
	beq .LU11
	b .LU12
.LU11:
	movw r0, #0
	b .LU9
.LU12:
	b .LU13
.LU13:
	ldr r0, [r0]
	bl size
	add r0, r0, #1
	b .LU9
.LU9:
	pop {fp, pc}
	.size size, .-size
	.align 2
	.global get
get:
.LU15:
	push {fp, lr}
	add fp, sp, #4
	mov r2, r1
	mov r1, #0
	cmp r2, #0
	moveq r1, #1
	cmp r1, #1
	beq .LU16
	b .LU17
.LU16:
	add r0, r0, #4
	ldr r0, [r0]
	b .LU14
.LU17:
	b .LU18
.LU18:
	ldr r3, [r0]
	sub r0, r2, #1
	mov r1, r0
	mov r0, r3
	bl get
	b .LU14
.LU14:
	pop {fp, pc}
	.size get, .-get
	.align 2
	.global pop
pop:
.LU20:
	push {fp, lr}
	add fp, sp, #4
	ldr r0, [r0]
	b .LU19
.LU19:
	pop {fp, pc}
	.size pop, .-pop
	.align 2
	.global printList
printList:
.LU22:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r4, r0
	mov r0, #0
	cmp r4, #0
	movne r0, #1
	cmp r0, #1
	beq .LU23
	b .LU24
.LU23:
	add r0, r4, #4
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, r4
	ldr r0, [r0]
	bl printList
	b .LU25
.LU24:
	b .LU25
.LU25:
	b .LU21
.LU21:
	pop {r4}
	pop {fp, pc}
	.size printList, .-printList
	.align 2
	.global treeprint
treeprint:
.LU27:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r4, r0
	mov r0, #0
	cmp r4, #0
	movne r0, #1
	cmp r0, #1
	beq .LU28
	b .LU29
.LU28:
	add r0, r4, #4
	ldr r0, [r0]
	bl treeprint
	mov r0, r4
	ldr r0, [r0]
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	add r0, r4, #8
	ldr r0, [r0]
	bl treeprint
	b .LU30
.LU29:
	b .LU30
.LU30:
	b .LU26
.LU26:
	pop {r4}
	pop {fp, pc}
	.size treeprint, .-treeprint
	.align 2
	.global freeList
freeList:
.LU32:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r4, r0
	mov r0, #0
	cmp r4, #0
	movne r0, #1
	cmp r0, #1
	beq .LU33
	b .LU34
.LU33:
	mov r0, r4
	ldr r0, [r0]
	bl freeList
	mov r0, r4
	bl free
	b .LU35
.LU34:
	b .LU35
.LU35:
	b .LU31
.LU31:
	pop {r4}
	pop {fp, pc}
	.size freeList, .-freeList
	.align 2
	.global freeTree
freeTree:
.LU37:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r4, r0
	mov r0, #0
	cmp r4, #0
	moveq r0, #1
	eor r0, r0, #1
	cmp r0, #1
	beq .LU38
	b .LU39
.LU38:
	add r0, r4, #4
	ldr r0, [r0]
	bl freeTree
	add r0, r4, #8
	ldr r0, [r0]
	bl freeTree
	mov r0, r4
	bl free
	b .LU40
.LU39:
	b .LU40
.LU40:
	b .LU36
.LU36:
	pop {r4}
	pop {fp, pc}
	.size freeTree, .-freeTree
	.align 2
	.global postOrder
postOrder:
.LU42:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6}
	mov r5, r0
	mov r0, #0
	cmp r5, #0
	movne r0, #1
	cmp r0, #1
	beq .LU43
	b .LU44
.LU43:
	movw r0, #8
	bl malloc
	mov r4, r0
	mov r0, r5
	ldr r1, [r0]
	add r0, r4, #4
	str r1, [r0]
	mov r0, r4
	movw r1, #0
	str r1, [r0]
	add r0, r5, #4
	ldr r0, [r0]
	bl postOrder
	mov r6, r0
	add r0, r5, #8
	ldr r0, [r0]
	bl postOrder
	mov r1, r0
	mov r0, r6
	bl concatLists
	mov r1, r4
	bl concatLists
	b .LU41
.LU44:
	b .LU45
.LU45:
	movw r0, #0
	b .LU41
.LU41:
	pop {r4, r5, r6}
	pop {fp, pc}
	.size postOrder, .-postOrder
	.align 2
	.global treeadd
treeadd:
.LU47:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r4, r0
	mov r5, r1
	mov r0, #0
	cmp r4, #0
	moveq r0, #1
	cmp r0, #1
	beq .LU48
	b .LU49
.LU48:
	movw r0, #12
	bl malloc
	mov r1, r0
	mov r0, r1
	str r5, [r0]
	add r2, r1, #4
	movw r0, #0
	str r0, [r2]
	add r2, r1, #8
	movw r0, #0
	str r0, [r2]
	mov r0, r1
	b .LU46
.LU49:
	b .LU50
.LU50:
	mov r0, r4
	ldr r1, [r0]
	mov r0, #0
	cmp r5, r1
	movlt r0, #1
	cmp r0, #1
	beq .LU51
	b .LU52
.LU51:
	add r0, r4, #4
	ldr r0, [r0]
	mov r1, r5
	bl treeadd
	mov r1, r0
	add r0, r4, #4
	str r1, [r0]
	b .LU53
.LU52:
	add r0, r4, #8
	ldr r0, [r0]
	mov r1, r5
	bl treeadd
	mov r1, r0
	add r0, r4, #8
	str r1, [r0]
	b .LU53
.LU53:
	mov r0, r4
	b .LU46
.LU46:
	pop {r4, r5}
	pop {fp, pc}
	.size treeadd, .-treeadd
	.align 2
	.global quickSort
quickSort:
.LU55:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6, r7, r8, r9}
	mov r4, r0
	mov r0, r4
	bl size
	mov r1, #0
	cmp r0, #1
	movle r1, #1
	cmp r1, #1
	beq .LU56
	b .LU57
.LU56:
	mov r0, r4
	b .LU54
.LU57:
	b .LU58
.LU58:
	movw r1, #0
	mov r0, r4
	bl get
	mov r5, r0
	mov r0, r4
	bl size
	sub r0, r0, #1
	mov r1, r0
	mov r0, r4
	bl get
	add r0, r5, r0
	movw r1, #2
	bl __aeabi_idiv
	mov r6, r0
	mov r5, r4
	movw r3, #0
	movw r1, #0
	movw r0, #0
	movw ip, #0
	movw r7, #0
	mov r2, #0
	cmp r4, #0
	movne r2, #1
	cmp r2, #1
	beq .LU59
	b .LU60
.LU59:
	mov r7, r5
	mov r5, r3
	mov r9, r1
	mov r8, r0
	mov r1, r8
	mov r0, r4
	bl get
	mov r1, r0
	mov r0, #0
	cmp r1, r6
	movgt r0, #1
	cmp r0, #1
	beq .LU61
	b .LU62
.LU61:
	mov r1, r8
	mov r0, r4
	bl get
	mov r1, r0
	mov r0, r9
	bl add
	mov r1, r0
	mov r0, r5
	b .LU63
.LU62:
	mov r1, r8
	mov r0, r4
	bl get
	mov r1, r0
	mov r0, r5
	bl add
	mov r1, r9
	b .LU63
.LU63:
	mov ip, r1
	mov r2, r0
	mov r0, r7
	ldr lr, [r0]
	add r0, r8, #1
	mov r5, lr
	mov r3, r2
	mov r1, ip
	mov r7, r2
	mov r2, #0
	cmp lr, #0
	movne r2, #1
	cmp r2, #1
	beq .LU59
	b .LU60
.LU60:
	mov r5, ip
	mov r6, r7
	mov r0, r4
	bl freeList
	mov r0, r6
	bl quickSort
	mov r4, r0
	mov r0, r5
	bl quickSort
	mov r1, r0
	mov r0, r4
	bl concatLists
	b .LU54
.LU54:
	pop {r4, r5, r6, r7, r8, r9}
	pop {fp, pc}
	.size quickSort, .-quickSort
	.align 2
	.global quickSortMain
quickSortMain:
.LU65:
	push {fp, lr}
	add fp, sp, #4
	push {r4}
	mov r4, r0
	mov r0, r4
	bl printList
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, r4
	bl printList
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, r4
	bl printList
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	b .LU64
.LU64:
	movw r0, #0
	pop {r4}
	pop {fp, pc}
	.size quickSortMain, .-quickSortMain
	.align 2
	.global treesearch
treesearch:
.LU67:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r5, r0
	mov r4, r1
	movw r1, #65535
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, #0
	cmp r5, #0
	movne r0, #1
	cmp r0, #1
	beq .LU68
	b .LU69
.LU68:
	mov r0, r5
	ldr r0, [r0]
	mov r1, #0
	cmp r0, r4
	moveq r1, #1
	cmp r1, #1
	beq .LU70
	b .LU71
.LU70:
	movw r0, #1
	b .LU66
.LU71:
	b .LU72
.LU72:
	add r0, r5, #4
	ldr r0, [r0]
	mov r1, r4
	bl treesearch
	mov r1, #0
	cmp r0, #1
	moveq r1, #1
	cmp r1, #1
	beq .LU73
	b .LU74
.LU73:
	movw r0, #1
	b .LU66
.LU74:
	b .LU75
.LU75:
	add r0, r5, #8
	ldr r0, [r0]
	mov r1, r4
	bl treesearch
	mov r1, #0
	cmp r0, #1
	moveq r1, #1
	cmp r1, #1
	beq .LU76
	b .LU77
.LU76:
	movw r0, #1
	b .LU66
.LU77:
	movw r0, #0
	b .LU66
.LU69:
	b .LU78
.LU78:
	movw r0, #0
	b .LU66
.LU66:
	pop {r4, r5}
	pop {fp, pc}
	.size treesearch, .-treesearch
	.align 2
	.global inOrder
inOrder:
.LU80:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6}
	mov r5, r0
	mov r0, #0
	cmp r5, #0
	movne r0, #1
	cmp r0, #1
	beq .LU81
	b .LU82
.LU81:
	movw r0, #8
	bl malloc
	mov r4, r0
	mov r0, r5
	ldr r0, [r0]
	add r1, r4, #4
	str r0, [r1]
	mov r1, r4
	movw r0, #0
	str r0, [r1]
	add r0, r5, #4
	ldr r0, [r0]
	bl inOrder
	mov r6, r0
	add r0, r5, #8
	ldr r0, [r0]
	bl inOrder
	mov r1, r0
	mov r0, r4
	bl concatLists
	mov r1, r0
	mov r0, r6
	bl concatLists
	b .LU79
.LU82:
	movw r0, #0
	b .LU79
.LU79:
	pop {r4, r5, r6}
	pop {fp, pc}
	.size inOrder, .-inOrder
	.align 2
	.global bintreesearch
bintreesearch:
.LU84:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r5, r0
	mov r4, r1
	movw r1, #65535
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, #0
	cmp r5, #0
	movne r0, #1
	cmp r0, #1
	beq .LU85
	b .LU86
.LU85:
	mov r0, r5
	ldr r0, [r0]
	mov r1, #0
	cmp r0, r4
	moveq r1, #1
	cmp r1, #1
	beq .LU87
	b .LU88
.LU87:
	movw r0, #1
	b .LU83
.LU88:
	b .LU89
.LU89:
	mov r0, r5
	ldr r1, [r0]
	mov r0, #0
	cmp r4, r1
	movlt r0, #1
	cmp r0, #1
	beq .LU90
	b .LU91
.LU90:
	add r0, r5, #4
	ldr r0, [r0]
	mov r1, r4
	bl bintreesearch
	b .LU83
.LU91:
	add r0, r5, #8
	ldr r0, [r0]
	mov r1, r4
	bl bintreesearch
	b .LU83
.LU86:
	b .LU92
.LU92:
	movw r0, #0
	b .LU83
.LU83:
	pop {r4, r5}
	pop {fp, pc}
	.size bintreesearch, .-bintreesearch
	.align 2
	.global buildTree
buildTree:
.LU94:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6, r7}
	mov r4, r0
	mov r0, r4
	bl size
	mov r3, r0
	movw lr, #0
	movw r0, #0
	movw r1, #0
	mov ip, #0
	movw r2, #0
	cmp r2, r3
	movlt ip, #1
	cmp ip, #1
	beq .LU95
	b .LU96
.LU95:
	mov r6, lr
	mov r5, r0
	mov r1, r6
	mov r0, r4
	bl get
	mov r1, r0
	mov r0, r5
	bl treeadd
	mov r7, r0
	add r5, r6, #1
	mov r0, r4
	bl size
	mov r2, r0
	mov lr, r5
	mov r0, r7
	mov r1, r7
	mov r3, #0
	cmp r5, r2
	movlt r3, #1
	cmp r3, #1
	beq .LU95
	b .LU96
.LU96:
	mov r0, r1
	b .LU93
.LU93:
	pop {r4, r5, r6, r7}
	pop {fp, pc}
	.size buildTree, .-buildTree
	.align 2
	.global treeMain
treeMain:
.LU98:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	bl buildTree
	mov r4, r0
	mov r0, r4
	bl treeprint
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, r4
	bl inOrder
	mov r5, r0
	mov r0, r5
	bl printList
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, r5
	bl freeList
	mov r0, r4
	bl postOrder
	mov r5, r0
	mov r0, r5
	bl printList
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, r5
	bl freeList
	movw r1, #0
	mov r0, r4
	bl treesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #10
	mov r0, r4
	bl treesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #65534
	movt r1, #65535
	mov r0, r4
	bl treesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #2
	mov r0, r4
	bl treesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #3
	mov r0, r4
	bl treesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #9
	mov r0, r4
	bl treesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #1
	mov r0, r4
	bl treesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #0
	mov r0, r4
	bl bintreesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #10
	mov r0, r4
	bl bintreesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #65534
	movt r1, #65535
	mov r0, r4
	bl bintreesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #2
	mov r0, r4
	bl bintreesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #3
	mov r0, r4
	bl bintreesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #9
	mov r0, r4
	bl bintreesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #1
	mov r0, r4
	bl bintreesearch
	mov r1, r0
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	movw r1, #64537
	movt r1, #65535
	movw r0, #:lower16:.PRINTLN_FMT
	movt r0, #:upper16:.PRINTLN_FMT
	bl printf
	mov r0, r4
	bl freeTree
	b .LU97
.LU97:
	pop {r4, r5}
	pop {fp, pc}
	.size treeMain, .-treeMain
	.align 2
	.global myCopy
myCopy:
.LU100:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5}
	mov r4, r0
	mov r0, #0
	cmp r4, #0
	moveq r0, #1
	cmp r0, #1
	beq .LU101
	b .LU102
.LU101:
	movw r0, #0
	b .LU99
.LU102:
	b .LU103
.LU103:
	mov r0, r4
	ldr r0, [r0]
	bl myCopy
	mov r5, r0
	add r0, r4, #4
	ldr r0, [r0]
	mov r1, r0
	movw r0, #0
	bl add
	mov r1, r5
	bl concatLists
	b .LU99
.LU99:
	pop {r4, r5}
	pop {fp, pc}
	.size myCopy, .-myCopy
	.align 2
	.global main
main:
.LU105:
	push {fp, lr}
	add fp, sp, #4
	push {r4, r5, r6, r7}
	movw r3, #0
	movw r2, #0
	b .LU106
.LU106:
	mov r7, r3
	mov r4, r2
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
	bl add
	mov r4, r0
	mov r0, r4
	bl myCopy
	mov r5, r0
	mov r0, r4
	bl myCopy
	mov r6, r0
	mov r0, r5
	bl quickSortMain
	bl freeList
	mov r0, r6
	bl treeMain
	add r0, r7, #1
	mov r3, r0
	mov r2, r4
	mov r1, #0
	cmp r0, #10
	movlt r1, #1
	cmp r1, #1
	beq .LU106
	b .LU107
.LU107:
	mov r0, r4
	bl freeList
	mov r0, r5
	bl freeList
	mov r0, r6
	bl freeList
	b .LU104
.LU104:
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

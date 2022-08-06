// rbtree 实现泛型红黑树
// 红黑树是一个二叉搜索平衡树，需要满足特性:
// 1.每个节点是红色或黑色
// 2.根节点是黑色
// 3.叶子结点Nil是黑色
// 4.红节点的两个孩子均为黑节点
// 5.每条简单路径上的黑节点数相同

package rbtree

import (
	"fmt"
	"github.com/royleo35/go-generics/tools"
	"github.com/royleo35/go-generics/types"
)

func NumberCompare[T types.Number](v1, v2 T) int {
	if v1 > v2 {
		return 1
	} else if v1 == v2 {
		return 0
	}
	return -1
}

type color int8

const (
	colorRed   color = 1
	colorBlack color = 2
)

type TreeNode[T any] struct {
	val                 T
	left, right, parent *TreeNode[T]
	color               color
	tree                *RBTree[T] // the tree which tree node belongs to
}
type RBTree[T any] struct {
	root    *TreeNode[T] // 根节点
	dummy   TreeNode[T]  // 哑结点，作为所有叶子节点节点的子节点，作为根节点的父节点, 为黑色；简化边界问题而设置
	len     int
	compare func(v1, v2 T) int
}

// NewRBTree 创建一颗空的红黑树，并传入比较函数,该比较函数需满足: v1 > v2 时返回1， v1=v2时返回0, v1 < v2 时返回 -1
func NewRBTree[T any](compare func(v1, v2 T) int) *RBTree[T] {
	t := &RBTree[T]{
		dummy: TreeNode[T]{
			color: colorBlack,
		},
		len:     0,
		compare: compare,
	}
	t.dummy.tree = t
	t.root = &t.dummy
	return t
}

func NewRBTreeWitchNumber[T types.Number]() *RBTree[T] {
	return NewRBTree[T](NumberCompare[T])
}

func NewWithNodes[T any](nodes []T, compare func(v1, v2 T) int) *RBTree[T] {
	t := NewRBTree(compare)
	for _, v := range nodes {
		t.Insert(v)
	}
	return t
}

func NewWithNumberNodes[T types.Number](nodes []T) *RBTree[T] {
	t := NewRBTreeWitchNumber[T]()
	for _, v := range nodes {
		t.Insert(v)
	}
	return t
}

// leftRotate rotate from node x by left side
//  	  	|      <--leftRotate(T, x)  	       		|
// 	   		y									   		x
//		 /     \	rightRotate(T, y) -->		      /    \
//		x       n3					 				 n1     y
//    /  \												   /  \
//   n1  n2												  n2   n3
// 左旋的时候，将n1->x->y的结构看成一个定滑轮，从n1点往后拉，这样x替换到n1的位置，y替换到x的位置，n2成为x的右孩子
func (t *RBTree[T]) leftRotate(x *TreeNode[T]) {
	if x.tree != t {
		return
	}
	// 假设x的右孩子是y
	y := x.right
	// 1.先将y的左孩子设置为x的右孩子
	x.right = y.left
	// 如果y的左孩子不是哑结点，即y存在左孩子，则设置左孩子的父节点为x
	if y.left != &t.dummy {
		y.left.parent = x
	}
	// 2.将y换到x的位置
	y.parent = x.parent
	if x.parent == &t.dummy {
		// 如果x就是根节点，则y变为根节点
		t.root = y
	} else if x.parent.left == x {
		// 如果x是父节点的左孩子
		x.parent.left = y
	} else {
		// 如果x是父节点的右孩子
		x.parent.right = y
	}
	// 3. 将x设置为y的左孩子
	x.parent = y
	y.left = x

}

func (t *RBTree[T]) rightRotate(y *TreeNode[T]) {
	if y.tree != t {
		return
	}
	x := y.left
	y.left = x.right
	if x.right != &t.dummy {
		x.right.parent = y
	}

	x.parent = y.parent
	if y.parent == &t.dummy {
		t.root = x
	} else if y.parent.left == y {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	x.right = y
	y.parent = x
}

func (t *RBTree[T]) Len() int {
	return t.len
}

func (t *TreeNode[T]) Value() T {
	return t.val
}

func (t *RBTree[T]) newNode(val T) *TreeNode[T] {
	return &TreeNode[T]{
		val:   val,
		color: colorRed,
		left:  &t.dummy,
		right: &t.dummy,
		tree:  t,
	}
}

// First 返回红黑树的第一个节点，如果是空树， 返回nil
func (t *RBTree[T]) First() *TreeNode[T] {
	if t.root == &t.dummy {
		return nil
	}
	return t.root.minNode()
}

// Next 返回当前节点的后继节点，如果不存在后继节点，返回nil
func (n *TreeNode[T]) Next() *TreeNode[T] {
	next := n.successor()
	if next == &n.tree.dummy {
		return nil
	}
	return next
}

// Contain tells if val in rbtree
func (t *RBTree[T]) Contain(val T) bool {
	return t.Find(val) != nil
}

// Find 查找红黑树中的一个节点，如果不存在返回nil
func (t *RBTree[T]) Find(val T) *TreeNode[T] {
	if t.len == 0 {
		return nil
	}
	node := t.root
	for node != &t.dummy {
		c := t.compare(val, node.val)
		if c == 0 {
			return node
		} else if c < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}
	return nil
}

func (t *RBTree[T]) ToSlice() []T {
	res := make([]T, 0, t.len)
	toSlice(t.root, &res)
	return res
}

func toSlice[T any](node *TreeNode[T], s *[]T) {
	if node.left != &node.tree.dummy {
		toSlice(node.left, s)
	}
	*s = append(*s, node.val)
	if node.right != &node.tree.dummy {
		toSlice(node.right, s)
	}
}

func (n *TreeNode[T]) colorStr() string {
	return tools.IfThen(n.color == colorRed, "red", "black")
}

func printSpace(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
}

func printElem(e string, length int) {
	//l := (length - len(e)) >> 1
	//printSpace(l)
	fmt.Printf(e)
	//printSpace(l)
}

func isPowerOf2(num int) bool {
	cnt := 0
	for num != 0 {
		if num&1 == 1 {
			cnt++
			if cnt > 1 {
				return false
			}
		}
		num >>= 1
	}
	return !(cnt > 1)
}

func powerOf2(num int) int {
	bit := 0
	for num != 0 {
		if num&1 == 1 {
			return bit
		}
		num >>= 1
		bit++
	}
	panic("ops")
}

func repeat(times int, f func()) {
	for i := 0; i < times; i++ {
		f()
	}
}

func printLine(s []string, level int, elemLen int) {
	tools.Assert(isPowerOf2(len(s)))
	currLev := powerOf2(len(s))
	side := 1 << (level - currLev - 1)
	inner := tools.IfThen(currLev == 0, 0, (1<<(level-currLev))-1)
	// side
	repeat(side, func() {
		printSpace(elemLen)
	})
	printElem(s[0], elemLen)
	// elem + inner
	for _, v := range s[1:] {
		repeat(inner, func() {
			printSpace(elemLen)
		})
		printElem(v, elemLen)
	}
	// side
	repeat(side, func() {
		printSpace(elemLen)
	})
	fmt.Print("\n\n")
}

func (t *RBTree[T]) Print() {
	if t.len == 0 {
		fmt.Println("empty rb-tree")
	}
	s := []*TreeNode[T]{t.root}
	var ps = []string{fmt.Sprintf("%v(%s)", s[0].val, s[0].colorStr())}
	level := 0
	var lastLevel int
	for len(s) > 0 {
		l := len(s)
		// push next level
		for _, v := range s {
			if v.left != &t.dummy {
				s = append(s, v.left)
				ps = append(ps, fmt.Sprintf("%v(%s)", v.left.val, v.left.colorStr()))
			} else {
				ps = append(ps, "Nil(black)")
			}
			if v.right != &t.dummy {
				s = append(s, v.right)
				ps = append(ps, fmt.Sprintf("%v(%s)", v.right.val, v.right.colorStr()))
			} else {
				ps = append(ps, "Nil(black)")
			}
		}
		lastLevel = len(s) << 1
		// cut last level
		s = s[l:]
		level++
	}
	ps = ps[:len(ps)-lastLevel]
	// print
	const elemLen = 10
	for i := 0; i < level; i++ {
		st := (1 << i) - 1
		end := (1 << (i + 1)) - 1
		currLe := ps[st:end]
		printLine(currLe, level, elemLen)
	}
}

// Insert 插入一个节点到红黑树中，即使存在节点，也覆盖
func (t *RBTree[T]) Insert(val T) {
	z := t.newNode(val)
	// 1.找到z需要插入的位置
	y := &t.dummy // 保持y一直是x的parent
	x := t.root   //
	for x != &t.dummy {
		y = x
		c := t.compare(val, x.val)
		if c == 0 { // 元素存在，覆盖之后，直接返回
			x.val = val
			return
		} else if c < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	// 2.将要插入的节点作为y的子节点
	z.parent = y
	if y == &t.dummy { // 空树
		t.root = z
	} else if t.compare(z.val, y.val) < 0 {
		y.left = z
	} else {
		y.right = z
	}

	// 3.调整红黑树
	t.insertFixUp(z)
	t.len++
}

// insertFixUp 以插入节点z为起点，调整红黑树，使其满足红黑树的所有性质
// 因为插入的节点一开始作为叶子节点，并且是红色的，因此只有可能破坏性质2或性质4
// 破坏性质2只可能是插入前为空树，插入后z为根节点，这种情况只需要将z改为黑色即可
// 如果破坏了性质4，一定是由于z和其父节点均为红色造成的；此种情况z的父节点可能是左节点或者右节点，两种情况是对称，只需要研究其中一种即可
// 已z的父节点z.parent为右节点为例，一共分为三种情况：
// 1.z 的叔叔节点z.parent.left为红色
//   将z.parent与z的叔叔设置为黑色，z的祖父设置为红色，这样不为违反别的性质;然后另z=z.p.p进行迭代，因为此时z.p.p.p有可能也是红色
// 2.z是右子节点，此时需要z与z.parent都是红色
//   z = z.parent 之后围绕z左旋，即转入情况3
// 3.z是左子节点,z.parent依然是红色, z.parent.parent是黑色
//  此时将z.parent 和z.parent.parent调换颜色，不影响z.parent.parent左侧的黑高，但是由于p.parent.parent反转成了红色，右侧黑高减少了1；
// 只需要绕z.parent.parent右旋即可
func (t *RBTree[T]) insertFixUp(z *TreeNode[T]) {
	for z.parent.color == colorRed {
		// z.parent是右孩子
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right // z的叔叔
			if y.color == colorRed {   // case 1
				z.parent.color = colorBlack
				y.color = colorBlack
				z.parent.parent.color = colorRed
				z = z.parent.parent
			} else if z == z.parent.right { // case2
				z = z.parent
				t.leftRotate(z)
			} else if z == z.parent.left { // case3
				z.parent.color = colorBlack
				z.parent.parent.color = colorRed
				t.rightRotate(z.parent.parent)
			}
		} else { // 另外一共情况完全对称，只需要文本left right互换即可
			y := z.parent.parent.left // z的叔叔
			if y.color == colorRed {  // case 1
				z.parent.color = colorBlack
				y.color = colorBlack
				z.parent.parent.color = colorRed
				z = z.parent.parent
			} else if z == z.parent.left { // case2
				z = z.parent
				t.rightRotate(z)
			} else if z == z.parent.right { // case3
				z.parent.color = colorBlack
				z.parent.parent.color = colorRed
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = colorBlack
}

// maxNode 返回以当前节点出发的最大节点
func (n *TreeNode[T]) maxNode() *TreeNode[T] {
	node := n
	for node.right != &n.tree.dummy {
		node = node.right
	}
	return node
}

// minNode 返回以当前节点出发的最小节点
func (n *TreeNode[T]) minNode() *TreeNode[T] {
	node := n
	for node.left != &n.tree.dummy {
		node = node.left
	}
	return node
}

// successor 查找当前节点的后继节点
func (n *TreeNode[T]) successor() *TreeNode[T] {
	// 如果节点有右子树，则后继节点为右子树的最左节点
	if n.right != &n.tree.dummy {
		return n.right.minNode() // 右子树的最小节点
	}
	// 如果没有右子树，则为其一个祖先节点，并且当前节点处于祖先节点的左子树中(因为后继节点一定大于当前节点)
	node := n
	p := node.parent
	for p != &n.tree.dummy && node == p.right {
		node = p
		p = p.parent
	}
	return p
}

// preSuccessor 查找当前节点的前驱节点
func (n *TreeNode[T]) preSuccessor() *TreeNode[T] {
	// 如果节点有左子树，则前驱节点为左子树的最右节点
	if n.left != &n.tree.dummy {
		return n.left.maxNode()
	}
	// 如果没有左子树，则为一个祖先节点，并且当前节点处于祖先节点的右子树中
	node := n
	p := node.parent
	for p != &n.tree.dummy && node == p.left {
		node = p
		p = p.parent
	}
	return p
}

// transplant 用v节点替换u节点的位置
// 该函数不设置v节点的left和right，由调用者维护
func transplant[T any](u, v *TreeNode[T]) {
	if u.tree != v.tree {
		return
	}
	null := &u.tree.dummy
	if u.parent == null { // u 是根节点
		u.tree.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

// Delete delete a val for rbtree, when node not exists, it dose nothing
// 删除红黑树节点分为三种情况
// 1.需要删除的节点z为叶子节点
//	 直接删除
// 2.需要删除的节点z只有一个孩子
//	 直接删除z: 即将z的父节点y和子节点x连接起来就行
// 3. 需要删除的节点z有两个孩子
//   则找到z节点的后继节点y，用y的染色和值覆盖z，然后删除y即可
// 不管属于以上哪种情况，如果最终被删除的节点是红色，不用管；如果是黑色，则破坏了红黑树的性质，需要调整
func (t *RBTree[T]) Delete(val T) {
	// 查找要删除的节点z，如果查不到就直接返回
	z := t.Find(val)
	if z == nil {
		return
	}
	null := &t.dummy
	y := z              // y 为真正需要删除的节点
	yColor := y.color   // 暂存y的颜色，后续要用
	x := null           // 节点x为null或者z的左孩子或右孩子
	if z.left == null { // 如果z没有左孩子
		x = z.right
		transplant(z, z.right)
	} else if z.right == null {
		x = z.left
		transplant(z, z.left)
	} else { // z有两个孩子
		y = z.right.minNode() // z的后继节点
		yColor = y.color
		x = y.right        // y为z的后继节点，所以y没有左孩子
		if y.parent == z { // y为z的右孩子，即y没有叶子z.right没有叶子节点
			x.parent = y
		} else {
			transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		transplant(z, y) // 用子树y代替子树z
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}
	if yColor == colorBlack {
		t.deleteFixUp(x)
	}
	t.len--
}

func (t *RBTree[T]) deleteFixUp(x *TreeNode[T]) {
	for x != t.root && x.color == colorBlack {
		if x == x.parent.left { // x为左孩子
			w := x.parent.right
			if w.color == colorRed {
				w.color = colorBlack
				x.parent.color = colorRed
				t.leftRotate(x.parent)
				w = x.parent.right
			}
			if w.left.color == colorBlack && w.right.color == colorBlack {
				w.color = colorRed
				x = x.parent
			} else if w.right.color == colorBlack {
				w.left.color = colorBlack
				w.color = colorRed
				t.rightRotate(w)
				w = x.parent.right
			}
			w.color = x.parent.color
			x.parent.color = colorBlack
			w.right.color = colorBlack
			t.leftRotate(x.parent)
			x = t.root
		} else {
			w := x.parent.left
			if w.color == colorRed {
				w.color = colorBlack
				x.parent.color = colorRed
				t.rightRotate(x.parent)
				w = x.parent.left
			}
			if w.right.color == colorBlack && w.left.color == colorBlack {
				w.color = colorRed
				x = x.parent
			} else if w.left.color == colorBlack {
				w.right.color = colorBlack
				w.color = colorRed
				t.leftRotate(w)
				w = x.parent.left
			}
			w.color = x.parent.color
			x.parent.color = colorBlack
			w.left.color = colorBlack
			t.rightRotate(x.parent)
			x = t.root
		}
	}
	x.color = colorBlack
}

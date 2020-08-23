/*
红黑树简单介绍
    红黑树是一种不严格平衡二叉树，通过**保持其性质**就能保持整个树的平衡。红黑树的允许左右节点之差大于1。
    红黑树的性质如下：

每一个结点要么是红色，要么是黑色。
根结点是黑色的。
所有叶子结点都是黑色的（NIL）。叶子结点不包含任何关键字信息，所有查询关键字都在非终结点上。
每个红色结点的两个子节点必须是黑色的。换句话说：从每个叶子到根的所有路径上不能有两个连续的红色结点
从任一结点到其每个叶子的所有路径都包含相同数目的黑色结点

   红黑树再具体实现的时候，查找方法和基本的二叉树没有区别，重点和难点就是在于**通过左旋右旋和置换颜色**来保持性质，从而实现平衡。
*/
package rbtree

//RBTree 红黑树本身为一个结构体，可以直接使用
type RBTree struct {
	root *node
}

//node 节点为一个结构体
type node struct {
	color      bool
	leftNode   *node
	rightNode  *node
	fatherNode *node
	value      int
}

//红黑通过bool类型来存储，并设置常量
const (
	RBTRed   = false
	RBTBlack = true
)

//find 查找
func (t *RBTree) find(v int) *node {
	n := t.root
	for n != nil {
		if v < n.value {
			//小于当前节点的话，往左节点找
			n = n.leftNode
		} else if v > n.value {
			//大于当前节点的话，往右节点找
			n = n.rightNode
		} else {
			//等于的话表示找到，返回
			return n
		}
	}
	//循环结束没找到，返回
	return nil
}

/*
 定义了结构体和查找方法之后我们来实现插入，在插入操作上采用二叉树的插入，插入了之后再对结构进行修正。
插入的新节点应该对当前红黑树的结构进行最小的破坏(最好修正)。插入的时候新的节点颜色设置为红节点，
如果插入的是黑节点，则根节点到新插入节点的黑路径比其他所有叶子节点的黑路径都大1，
问题扩散开来。如果插入的是红节点，那么被破坏的性质只剩下性质2和性质4。
性质2被破坏可以直接设置根节点颜色，性质4被破坏可以通过旋转变色(将两个连续的红色变成一红一黑)来解决。
*/
func (t *RBTree) insert(v int) {
	//如果根节点为nil，则先插入新的根节点。
	if t.root == nil {
		t.root = &node{value: v, color: RBTBlack}
		return
	}

	n := t.root
	//新插入的节点为红色
	insertNode := &node{value: v, color: RBTRed}
	//标记父节点
	var nf *node
	//一下代码找到插入位置的父节点
	for n != nil {

		nf = n
		if v < n.value {
			n = n.leftNode
		} else if v > n.value {
			n = n.rightNode
		} else {
			//TODO fix the condition that replace value
			//已经存在，返回
			return
		}
	}
	//设置新插入节点的父节点
	insertNode.fatherNode = nf
	//将新的节点挂到父节点上
	if v < nf.value {
		nf.leftNode = insertNode
	} else {
		nf.rightNode = insertNode
	}
	t.insertFixUp(insertNode)
}

// 插入调整：
func (t *RBTree) insertFixUp(n *node) {
	//父节点是黑色的话终止，不是的话进入调整
	for !isBlack(n.fatherNode) {
		//fmt.Printf("%v\t", n)
		//grandpa's color is black
		//case1  uncle's color is red then set grandpa's red color and his child black
		// n -> n's grandpa
		//if n is the same side of its father
		//exchange its father and grandpa by rotate
		//else make its side by rotate
		uncleNode := findBroNode(n.fatherNode)
		if !isBlack(uncleNode) {
			//叔节点是红色的话，将爷爷节点变红，爹和叔变黑，调整节点上移指向爷爷
			n.fatherNode.color = RBTBlack
			uncleNode.color = RBTBlack
			uncleNode.fatherNode.color = RBTRed
			n = n.fatherNode.fatherNode
			//	fmt.Println("condition1")
		} else if n.fatherNode == n.fatherNode.fatherNode.leftNode {

			//fmt.Println("condition2")
			if n == n.fatherNode.leftNode {
				//父节点和自己在一条直线上，旋转变色后满足循环终止条件
				n.fatherNode.fatherNode.color = RBTRed
				n.fatherNode.color = RBTBlack
				n = n.fatherNode.fatherNode
				t.rightRotate(n)

			} else {
				//父节点和自己不在一条直线上，旋转到一条直线

				n = n.fatherNode
				t.leftRotate(n)
			}

		} else {
			//fmt.Println("condition2")

			if n == n.fatherNode.rightNode {
				n.fatherNode.fatherNode.color = RBTRed
				n.fatherNode.color = RBTBlack
				n = n.fatherNode.fatherNode
				t.leftRotate(n)

			} else {
				n = n.fatherNode
				t.rightRotate(n)
			}
		}
		t.root.color = RBTBlack
	}
}

// 左旋右旋：
func (t *RBTree) leftRotate(n *node) {
	rn := n.rightNode
	//first give n's father to rn's father
	rn.fatherNode = n.fatherNode
	if n.fatherNode != nil {
		if n.fatherNode.leftNode == n {
			n.fatherNode.leftNode = rn
		} else {
			n.fatherNode.rightNode = rn
		}
	} else {
		t.root = rn
	}
	n.fatherNode = rn
	n.rightNode = rn.leftNode
	if n.rightNode != nil {
		n.rightNode.fatherNode = n
	}
	rn.leftNode = n
}
func (t *RBTree) rightRotate(n *node) {

	ln := n.leftNode
	ln.fatherNode = n.fatherNode
	if n.fatherNode != nil {
		if n.fatherNode.leftNode == n {
			n.fatherNode.leftNode = ln
		} else {
			n.fatherNode.rightNode = ln
		}
	} else {
		t.root = ln
	}
	n.fatherNode = ln

	n.leftNode = ln.rightNode
	if n.leftNode != nil {

		n.leftNode.fatherNode = n
	}
	ln.rightNode = n

}

// 其他一些辅助函数:
//判断是否为黑，空为黑
func isBlack(n *node) bool {
	if n == nil {
		return true
	} else {
		return n.color == RBTBlack
	}
}

//设置节点颜色
func setColor(n *node, color bool) {
	if n == nil {
		return
	}
	n.color = color
}

//寻找兄弟节点
func findBroNode(n *node) (bro *node) {
	if n.fatherNode == nil {
		return nil
	}

	if n.fatherNode.leftNode == n {
		bro = n.fatherNode.rightNode
	} else if n.fatherNode.rightNode == n {
		bro = n.fatherNode.leftNode
	} else {
		if n.fatherNode.leftNode == nil {
			bro = n.fatherNode.rightNode
		} else {
			bro = n.fatherNode.leftNode

		}
	}
	return bro
}

/*  删除的话要分种情况，当删除的是红节点的时候不影响性质

删除节点没有只有一个子节点的话，直接删除，将子节点补位
删除节点没有子节点的话，直接删除
删除节点有两个子节点的话，删除后将后继节点补位，并设置颜色(后继节点与删除节点交换位置和颜色后删除) ，
后继节点的子节点补位后继节点的位置，其实删除位置和颜色相当于后继节点
    删除节点的话，需要记录删除的颜色，第三种情况删除的颜色是后继节点。
	如果删除的是黑节点的话，删除后从删除节点(后继节点)的补位节点开始修复。

    如果删除的节点没有补位节点的话还面临补位节点为空的问题，可以用fixNode来解决，
	一个虚拟的补位节点，设置父节点，但是不挂到树上。
*/
// 删除代码
func (t *RBTree) delete(v int) {
	n := t.find(v)
	if n == nil {
		return
	}
	// if n == t.root {
	// 	fmt.Println("delete root")
	// 	t.printGra()
	// }
	//copy color of fixNode
	var fixColor = n.color
	//if fixNode == nil copy node of start fix node
	//set it's father and set color black
	var fixNode = &node{fatherNode: n.fatherNode, color: RBTBlack}

	if n.leftNode == nil {
		t.transplant(n, n.rightNode)
		if n.rightNode != nil {
			fixNode = n.rightNode
		}
	} else if n.rightNode == nil {
		t.transplant(n, n.leftNode)
		if n.leftNode != nil {
			fixNode = n.leftNode
		}
	} else {
		succNode := t.miniNum(n.rightNode)
		fixColor = succNode.color
		if succNode.rightNode == nil {
			if succNode.fatherNode != n {
				fixNode = &node{fatherNode: succNode.fatherNode, color: RBTBlack}
			} else {
				fixNode = &node{fatherNode: succNode, color: RBTBlack}
			}
		} else {
			fixNode = succNode.rightNode
		}
		if succNode.fatherNode != n {
			t.transplant(succNode, succNode.rightNode)
			succNode.rightNode = n.rightNode
			succNode.rightNode.fatherNode = succNode
		} else {

		}
		t.transplant(n, succNode)
		succNode.leftNode = n.leftNode
		succNode.leftNode.fatherNode = succNode
		succNode.color = n.color
	}

	if fixColor == RBTBlack {
		t.deleteFixUp(fixNode)
	}

}

// 删除调整代码:
func (t *RBTree) deleteFixUp(n *node) {

	if t.root == nil {
		return
	}
	//n为红色，或者是根，结束调整，置为黑色
	for n != t.root && isBlack(n) {
		bro := findBroNode(n)

		if bro != n.fatherNode.leftNode {
			//case 1 如果兄弟节点是红色，旋转变色给自己弄个黑色的兄弟，转移到情况234
			if !isBlack(bro) {

				n.fatherNode.color = RBTRed
				bro.color = RBTBlack
				t.leftRotate(n.fatherNode)
				bro = findBroNode(n)
				// now new bro is black
			}

			//if bro is black its children perhaps be nil
			//if bro's children are black
			// n up
			if isBlack(bro.leftNode) && isBlack(bro.rightNode) {
				//case2 两个侄子都是黑色的，设置兄弟节点为红色，这样自己缺1黑高，兄弟缺1黑高，调整节点上移
				setColor(bro, RBTRed)
				n = n.fatherNode
				//fmt.Printf("%v\n", n)

			} else {
				//case3 红色侄子和兄弟在一条直线上，旋转变色，达到终止条件
				if !isBlack(bro.rightNode) {

					bro.color = n.fatherNode.color
					bro.rightNode.color = RBTBlack
					n.fatherNode.color = RBTBlack
					t.leftRotate(n.fatherNode)
					n = t.root
				} else {
					//case4 ，旋转变色，变成case3

					bro.color = RBTRed
					bro.leftNode.color = RBTBlack
					t.rightRotate(bro)
					bro = findBroNode(n)
				}
			}

		} else {
			//case 1
			if !isBlack(bro) {

				n.fatherNode.color = RBTRed
				bro.color = RBTBlack
				t.rightRotate(n.fatherNode)
				bro = findBroNode(n)
				// now new bro is black
			}

			//if bro is black its children perhaps be nil
			//if bro's children are black
			// n up

			if isBlack(bro.leftNode) && isBlack(bro.rightNode) {
				//case2

				setColor(bro, RBTRed)
				n = n.fatherNode
			} else {
				//case3
				if !isBlack(bro.leftNode) {

					bro.color = n.fatherNode.color
					bro.leftNode.color = RBTBlack
					n.fatherNode.color = RBTBlack
					t.rightRotate(n.fatherNode)
					break

				} else {
					//case4

					bro.color = RBTRed
					bro.rightNode.color = RBTBlack
					t.leftRotate(bro)
				}
			}

		}
	}
	n.color = RBTBlack
}

//其他辅助代码：
//
//后继节点，右边最大的数
func (t *RBTree) miniNum(n *node) *node {
	for n.leftNode != nil {
		n = n.leftNode
	}
	return n
}

//替换代码，颜色的情况比较复杂在具体删除的时候写
func (t *RBTree) transplant(u, v *node) {

	if u.fatherNode == nil {
		t.root = v
		if v != nil {
			v.fatherNode = nil
		}

	} else if u == u.fatherNode.leftNode {
		u.fatherNode.leftNode = v

	} else {
		u.fatherNode.rightNode = v

	}
	if v != nil {
		v.fatherNode = u.fatherNode
	}

}

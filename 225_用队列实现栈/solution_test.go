package leetcode

import (
	"container/list"
	"testing"
)

/**
请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。

 实现 MyStack 类：


 void push(int x) 将元素 x 压入栈顶。
 int pop() 移除并返回栈顶元素。
 int top() 返回栈顶元素。
 boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。




 注意：


 你只能使用队列的标准操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
 你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。




 示例：


输入：
["MyStack", "push", "push", "top", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 2, 2, false]

解释：
MyStack myStack = new MyStack();
myStack.push(1);
myStack.push(2);
myStack.top(); // 返回 2
myStack.pop(); // 返回 2
myStack.empty(); // 返回 False




 提示：


 1 <= x <= 9
 最多调用100 次 push、pop、top 和 empty
 每次调用 pop 和 top 都保证栈不为空




 进阶：你能否仅用一个队列来实现栈。

 Related Topics 栈 设计 队列 👍 903 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	q        *list.List
	top_elem int
}

// MyStack构造器
func Constructor() MyStack {
	return MyStack{q: list.New()}
}

// 添加元素到栈顶
func (this *MyStack) Push(x int) {
	this.q.PushBack(x)
	this.top_elem = x
}

// 返回栈顶元素
func (this *MyStack) Top() int {
	return this.top_elem
}

// 检查栈是否为空
func (this *MyStack) Empty() bool {
	return this.q.Len() == 0
}

// 删除栈顶的元素并返回
func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0 // 或者可以考虑用 error 返回
	}

	size := this.q.Len()
	for i := 0; i < size-1; i++ {
		elem := this.q.Remove(this.q.Front())
		this.q.PushBack(elem)
		this.top_elem = elem.(int) // 更新top_elem为最新的顶端元素
	}

	// Remove and return the last element, which is the stack top
	backElem := this.q.Remove(this.q.Front())
	return backElem.(int)
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestImplementStackUsingQueues(t *testing.T) {

}

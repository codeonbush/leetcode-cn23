package leetcode

import(
    "testing"
)

/**
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）： 

 实现 MyQueue 类： 

 
 void push(int x) 将元素 x 推到队列的末尾 
 int pop() 从队列的开头移除并返回元素 
 int peek() 返回队列开头的元素 
 boolean empty() 如果队列为空，返回 true ；否则，返回 false 
 

 说明： 

 
 你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
 
 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。 
 

 

 示例 1： 

 
输入：
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 1, 1, false]

解释：
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false
 

 
 

 

 提示： 

 
 1 <= x <= 9 
 最多调用 100 次 push、pop、peek 和 empty 
 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作） 
 

 

 进阶： 

 
 你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。 
 

 Related Topics 栈 设计 队列 👍 1137 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
type MyQueue struct {
    s1, s2 []int
}

//建构子
func Constructor() MyQueue {
    return MyQueue{
        s1: make([]int, 0),
        s2: make([]int, 0),
    }
}
// 添加元素到队尾
func (this *MyQueue) Push(x int) {
    this.s1 = append(this.s1, x)
}

// 返回队头元素
func (this *MyQueue) Peek() int {
    if len(this.s2) == 0 {
        // 把 s1 元素压入 s2
        for len(this.s1) > 0 {
            this.s2 = append(this.s2, this.s1[len(this.s1)-1])
            this.s1 = this.s1[:len(this.s1)-1]
        }
    }
    return this.s2[len(this.s2)-1]
}
// 删除队头元素并返回
func (this *MyQueue) Pop() int {
    // 先调用 peek 保证 s2 非空
    this.Peek()
    retval := this.s2[len(this.s2)-1]
    this.s2 = this.s2[:len(this.s2)-1]
    return retval
}

// 判断队列是否为空
// 两个栈都为空才说明队列为空
func (this *MyQueue) Empty() bool {
    return len(this.s1) == 0 && len(this.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)


func TestImplementQueueUsingStacks(t *testing.T){
    
}

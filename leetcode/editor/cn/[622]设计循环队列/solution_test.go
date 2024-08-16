package leetcode

import (
	"testing"
)

/**
设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。


 循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列
，我们能使用这些空间去存储新的值。

 你的实现应该支持如下操作：


 MyCircularQueue(k): 构造器，设置队列长度为 k 。
 Front: 从队首获取元素。如果队列为空，返回 -1 。
 Rear: 获取队尾元素。如果队列为空，返回 -1 。
 enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
 deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
 isEmpty(): 检查循环队列是否为空。
 isFull(): 检查循环队列是否已满。




 示例：

 MyCircularQueue circularQueue = new MyCircularQueue(3); // 设置长度为 3
circularQueue.enQueue(1);  // 返回 true
circularQueue.enQueue(2);  // 返回 true
circularQueue.enQueue(3);  // 返回 true
circularQueue.enQueue(4);  // 返回 false，队列已满
circularQueue.Rear();  // 返回 3
circularQueue.isFull();  // 返回 true
circularQueue.deQueue();  // 返回 true
circularQueue.enQueue(4);  // 返回 true
circularQueue.Rear();  // 返回 4



 提示：


 所有的值都在 0 至 1000 的范围内；
 操作数将在 1 至 1000 的范围内；
 请不要使用内置的队列库。


 Related Topics 设计 队列 数组 链表 👍 528 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// ArrayQueue 用数组实现队列
type ArrayQueue struct {
	size        int
	data        []interface{}
	INIT_CAP    int //初始容量
	first, last int //队首，队尾 为什么要记录队首队尾
}

// NewArrayQueue 创建一个新的 ArrayQueue
func NewArrayQueue(initCap int) *ArrayQueue {
	if initCap < 1 {
		initCap = 1
	}
	return &ArrayQueue{
		size:     0,
		data:     make([]interface{}, initCap),
		INIT_CAP: 2, //初始容量为什么是2
		first:    0,
		last:     0,
	}
}

// Enqueue 入队
func (q *ArrayQueue) Enqueue(e interface{}) {
	if q.size == len(q.data) {
		q.resize(len(q.data) * 2)
	}
	q.data[q.last] = e
	q.last++
	if q.last == len(q.data) {
		q.last = 0
	}
	q.size++
}

// Dequeue 出队
func (q *ArrayQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	if q.size == len(q.data)/4 {
		q.resize(len(q.data) / 2)
	}
	oldVal := q.data[q.first]
	q.data[q.first] = nil
	q.first++
	if q.first == len(q.data) {
		q.first = 0
	}
	q.size--
	return oldVal
}

// PeekFirst 获取队列的首部元素
func (q *ArrayQueue) PeekFirst() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.data[q.first]
}

// PeekLast 获取队列的尾部元素
func (q *ArrayQueue) PeekLast() interface{} {
	if q.IsEmpty() {
		return nil
	}
	if q.last == 0 {
		return q.data[len(q.data)-1]
	}
	return q.data[q.last-1]
}

// Size 获取队列的元素数量
func (q *ArrayQueue) Size() int {
	return q.size
}

// IsEmpty 判断队列是否为空
func (q *ArrayQueue) IsEmpty() bool {
	return q.size == 0
}

// resize 调整队列的容量大小
func (q *ArrayQueue) resize(newCap int) {
	temp := make([]interface{}, newCap)

	for i := 0; i < q.size; i++ {
		temp[i] = q.data[(q.first+i)%len(q.data)]
	}
	q.first = 0
	q.last = q.size
	q.data = temp
}

/**
  先用数组实现队列，再组合为环形队列
*/
// MyCircularQueue 顺时针旋转队列
type MyCircularQueue struct {
	q      *ArrayQueue
	maxCap int //最大容量
}

// Constructor 初始化时创建 MyCircularQueue
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		q:      NewArrayQueue(k),
		maxCap: k,
	}
}

// EnQueue 入队
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.q.Size() == this.maxCap {
		return false
	}
	this.q.Enqueue(value)
	return true
}

// DeQueue 出队
func (this *MyCircularQueue) DeQueue() bool {
	if this.q.IsEmpty() {
		return false
	}
	this.q.Dequeue()
	return true
}

// Front 获取队列的首部元素
func (this *MyCircularQueue) Front() int {
	if this.q.IsEmpty() {
		return -1
	}
	return this.q.PeekFirst().(int)
}

// Rear 获取队列的尾部元素
func (this *MyCircularQueue) Rear() int {
	if this.q.IsEmpty() {
		return -1
	}
	return this.q.PeekLast().(int)
}

// IsEmpty 判断队列是否为空
func (this *MyCircularQueue) IsEmpty() bool {
	return this.q.IsEmpty()
}

// IsFull 判断队列是否已满
func (this *MyCircularQueue) IsFull() bool {
	return this.q.Size() == this.maxCap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestDesignCircularQueue(t *testing.T) {

}

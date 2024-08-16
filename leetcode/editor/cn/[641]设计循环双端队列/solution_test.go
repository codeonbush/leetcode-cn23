package leetcode

import (
	"testing"
)

/**
设计实现双端队列。

 实现 MyCircularDeque 类:


 MyCircularDeque(int k) ：构造函数,双端队列最大为 k 。
 boolean insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true ，否则返回 false 。
 boolean insertLast() ：将一个元素添加到双端队列尾部。如果操作成功返回 true ，否则返回 false 。
 boolean deleteFront() ：从双端队列头部删除一个元素。 如果操作成功返回 true ，否则返回 false 。
 boolean deleteLast() ：从双端队列尾部删除一个元素。如果操作成功返回 true ，否则返回 false 。
 int getFront() )：从双端队列头部获得一个元素。如果双端队列为空，返回 -1 。
 int getRear() ：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1 。
 boolean isEmpty() ：若双端队列为空，则返回 true ，否则返回 false 。
 boolean isFull() ：若双端队列满了，则返回 true ，否则返回 false 。




 示例 1：


输入
["MyCircularDeque", "insertLast", "insertLast", "insertFront", "insertFront",
"getRear", "isFull", "deleteLast", "insertFront", "getFront"]
[[3], [1], [2], [3], [4], [], [], [], [4], []]
输出
[null, true, true, true, false, 2, true, true, true, 4]

解释
MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
circularDeque.insertLast(1);			        // 返回 true
circularDeque.insertLast(2);			        // 返回 true
circularDeque.insertFront(3);			        // 返回 true
circularDeque.insertFront(4);			        // 已经满了，返回 false
circularDeque.getRear();  				// 返回 2
circularDeque.isFull();				        // 返回 true
circularDeque.deleteLast();			        // 返回 true
circularDeque.insertFront(4);			        // 返回 true
circularDeque.getFront();				// 返回 4




 提示：


 1 <= k <= 1000
 0 <= value <= 1000
 insertFront, insertLast, deleteFront, deleteLast, getFront, getRear, isEmpty,
isFull 调用次数不大于 2000 次


 Related Topics 设计 队列 数组 链表 👍 237 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
type MyArrayDeque struct {
	size  int
	data  []interface{} // 存储数据
	first int           // 头索引
	last  int           // 尾索引
}

const INIT_CAP = 2 // 初始容量

func newMyArrayDeque(initCap int) *MyArrayDeque {
	return &MyArrayDeque{
		size:  0,
		data:  make([]interface{}, initCap),
		first: 0,
		last:  0,
	}
}

// 从头部插入元素
func (d *MyArrayDeque) AddFirst(e interface{}) {
	if d.size == len(d.data) {
		d.resize(len(d.data) * 2)
	}

	if d.first == 0 {
		d.first = len(d.data) - 1
	} else {
		d.first--
	}
	d.data[d.first] = e

	d.size++
}

// 从尾部插入元素
func (d *MyArrayDeque) AddLast(e interface{}) {
	if d.size == len(d.data) {
		d.resize(len(d.data) * 2)
	}
	d.data[d.last] = e
	d.last++
	if d.last == len(d.data) {
		d.last = 0
	}

	d.size++
}

// 从头部删除元素
func (d *MyArrayDeque) RemoveFirst() interface{} {
	if d.IsEmpty() {
		panic("removeFirst() failed")
	}

	if d.size == len(d.data)/4 {
		d.resize(len(d.data) / 2)
	}

	e := d.data[d.first]
	d.data[d.first] = nil

	d.first++
	if d.first == len(d.data) {
		d.first = 0
	}

	d.size--
	return e
}

// 从尾部删除元素
func (d *MyArrayDeque) RemoveLast() interface{} {
	if d.IsEmpty() {
		panic("removeLast() failed")
	}

	if d.size == len(d.data)/4 {
		d.resize(len(d.data) / 2)
	}

	if d.last == 0 {
		d.last = len(d.data) - 1
	} else {
		d.last--
	}
	e := d.data[d.last]
	d.data[d.last] = nil

	d.size--
	return e
}

// 从头部获取元素
func (d *MyArrayDeque) GetFirst() interface{} {
	if d.IsEmpty() {
		panic("getFirst() failed")
	}
	return d.data[d.first]
}

// 从尾部获取元素
func (d *MyArrayDeque) GetLast() interface{} {
	if d.IsEmpty() {
		panic("getLast() failed")
	}

	if d.last == 0 {
		return d.data[len(d.data)-1]
	}
	return d.data[d.last-1]
}

// 是否为空
func (d *MyArrayDeque) IsEmpty() bool {
	return d.size == 0
}

// 动态扩容
func (d *MyArrayDeque) resize(newCap int) {
	temp := make([]interface{}, newCap)

	for i := 0; i < d.size; i++ {
		temp[i] = d.data[(d.first+i)%len(d.data)]
	}

	d.first = 0
	d.last = d.size
	d.data = temp
}

type MyCircularDeque struct {
	cap  int
	list *MyArrayDeque
}

// 初始化双端队列，设置容量
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		cap:  k,
		list: newMyArrayDeque(k),
	}
}

// 从头部插入元素
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.list.size == this.cap {
		return false
	}

	this.list.AddFirst(value)
	return true
}

// 从尾部插入元素
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.list.size == this.cap {
		return false
	}

	this.list.AddLast(value)
	return true
}

// 从头部删除元素
func (this *MyCircularDeque) DeleteFront() bool {
	if this.list.IsEmpty() {
		return false
	}

	this.list.RemoveFirst()
	return true
}

// 从尾部删除元素
func (this *MyCircularDeque) DeleteLast() bool {
	if this.list.IsEmpty() {
		return false
	}

	this.list.RemoveLast()
	return true
}

// 获取头元素
func (this *MyCircularDeque) GetFront() int {
	if this.list.IsEmpty() {
		return -1
	}

	return this.list.GetFirst().(int)
}

// 获取尾元素
func (this *MyCircularDeque) GetRear() int {
	if this.list.IsEmpty() {
		return -1
	}

	return this.list.GetLast().(int)
}

// 判断是否为空
func (this *MyCircularDeque) IsEmpty() bool {
	return this.list.IsEmpty()
}

// 判断是否已满
func (this *MyCircularDeque) IsFull() bool {
	return this.list.size == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestDesignCircularDeque(t *testing.T) {

}

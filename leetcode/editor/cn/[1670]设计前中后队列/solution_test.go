package leetcode

import (
	"container/list"
	"testing"
)

/**
请你设计一个队列，支持在前，中，后三个位置的 push 和 pop 操作。

 请你完成 FrontMiddleBack 类：


 FrontMiddleBack() 初始化队列。
 void pushFront(int val) 将 val 添加到队列的 最前面 。
 void pushMiddle(int val) 将 val 添加到队列的 正中间 。
 void pushBack(int val) 将 val 添加到队里的 最后面 。
 int popFront() 将 最前面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
 int popMiddle() 将 正中间 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
 int popBack() 将 最后面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。


 请注意当有 两个 中间位置的时候，选择靠前面的位置进行操作。比方说：


 将 6 添加到 [1, 2, 3, 4, 5] 的中间位置，结果数组为 [1, 2, 6, 3, 4, 5] 。
 从 [1, 2, 3, 4, 5, 6] 的中间位置弹出元素，返回 3 ，数组变为 [1, 2, 4, 5, 6] 。




 示例 1：


输入：
["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle",
"popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
[[], [1], [2], [3], [4], [], [], [], [], []]
输出：
[null, null, null, null, null, 1, 3, 4, 2, -1]

解释：
FrontMiddleBackQueue q = new FrontMiddleBackQueue();
q.pushFront(1);   // [1]
q.pushBack(2);    // [1, 2]
q.pushMiddle(3);  // [1, 3, 2]
q.pushMiddle(4);  // [1, 4, 3, 2]
q.popFront();     // 返回 1 -> [4, 3, 2]
q.popMiddle();    // 返回 3 -> [4, 2]
q.popMiddle();    // 返回 4 -> [2]
q.popBack();      // 返回 2 -> []
q.popFront();     // 返回 -1 -> [] （队列为空）




 提示：


 1 <= val <= 10⁹
 最多调用 1000 次 pushFront， pushMiddle， pushBack， popFront， popMiddle 和 popBack 。


 Related Topics 设计 队列 数组 链表 数据流 👍 94 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
type FrontMiddleBackQueue struct {
	left  *list.List
	right *list.List
}

// 用两个列表表示队列的左右两部分，一遍从中间操作元素
// 如果是奇数个元素，维护左边少右边多，所以：
// 1、如果有偶数个元素时，pushMiddle 优先向右边添加
// 2、如果有奇数个元素时，popMiddle 优先从右边删除
// 3、如果只有 1 个元素，popFront 的时候，要去右边删除
// 要把以上三个特点写到代码里，才能保证细节不出错

// 维护左边少右边多的状态，每次增删元素之后都要执行一次
func (q *FrontMiddleBackQueue) balance() {
	// 右边最多比左边多一个元素
	if q.right.Len() > q.left.Len()+1 {
		// 右边多，匀一个给左边
		q.left.PushBack(q.right.Remove(q.right.Front()))
	}
	if q.left.Len() > q.right.Len() {
		// 左边多，匀一个给右边
		q.right.PushFront(q.left.Remove(q.left.Back()))
	}
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		left:  list.New(),
		right: list.New(),
	}
}

func (q *FrontMiddleBackQueue) PushFront(val int) {
	q.left.PushFront(val)
	q.balance()
}

func (q *FrontMiddleBackQueue) PushMiddle(val int) {
	if (q.left.Len()+q.right.Len())%2 == 0 {
		// 如果有偶数个元素时，pushMiddle 优先向右边添加
		q.right.PushFront(val)
	} else {
		q.left.PushBack(val)
	}
	q.balance()
}

func (q *FrontMiddleBackQueue) PushBack(val int) {
	q.right.PushBack(val)
	q.balance()
}

func (q *FrontMiddleBackQueue) PopFront() int {
	if q.left.Len()+q.right.Len() == 0 {
		return -1
	}
	if q.left.Len()+q.right.Len() == 1 {
		// 如果只有 1 个元素，popFront 的时候，要去右边删除
		return q.right.Remove(q.right.Front()).(int)
	}
	e := q.left.Remove(q.left.Front()).(int)
	q.balance()
	return e
}

func (q *FrontMiddleBackQueue) PopMiddle() int {
	if q.left.Len()+q.right.Len() == 0 {
		return -1
	}
	var e int
	if (q.left.Len()+q.right.Len())%2 == 0 {
		e = q.left.Remove(q.left.Back()).(int)
	} else {
		// 如果有奇数个元素时，popMiddle 优先从右边删除
		e = q.right.Remove(q.right.Front()).(int)
	}
	q.balance()
	return e
}

func (q *FrontMiddleBackQueue) PopBack() int {
	if q.left.Len()+q.right.Len() == 0 {
		return -1
	}
	e := q.right.Remove(q.right.Back()).(int)
	q.balance()
	return e
}

func (q *FrontMiddleBackQueue) Size() int {
	return q.left.Len() + q.right.Len()
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestDesignFrontMiddleBackQueue(t *testing.T) {

}

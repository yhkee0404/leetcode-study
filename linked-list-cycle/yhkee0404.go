/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    u := head
    v := head
    for u != nil && v != nil { // O(n) https://hyp.is/gLsBIOHBEe6E_9M9T5MJNw/yhkee0404.github.io/posts/algorithms/leetcode/linked-list-cycle/
        u = u.Next
        if (u == nil) {
            break
        }
        u = u.Next
        v = v.Next
        if u == v {
            return true
        }
    }
    return false
}

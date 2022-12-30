# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        if not root:
            return 0

        depth_left = self.maxDepth(root.left)
        depth_right = self.maxDepth(root.right)

        return depth_left + 1 if depth_left > depth_right else depth_right + 1

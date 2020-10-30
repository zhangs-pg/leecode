class TreeNode:
    def __init__(self, val, left, right):
        self.val = val
        self.left = left
        self.right = right

def helper(node, l):
    if not l:
        return
    if node:
        left = TreeNode(l.pop(0), None, None)
        node.left = left
        if not l:
            return
        right = TreeNode(l.pop(0), None, None)
        node.right = right
        helper(left, l)
        helper(right, l)

def create(l):
    root = TreeNode(l.pop(0), None, None)
    helper(root, l)
    return root

def get_node(node):
    if node:
        if node.left:
            get_node(node.left)
        print(node.val)
        if node.right:
            get_node(node.right)

def head_append(l, v):
    tmp = l[::-1]
    tmp.append(v)
    return tmp[::-1]

def iter_node(root):
    stack = []
    res = []

    while root or stack:
    	while root:
    	    stack.append(root)
    	    root = root.left
    	
    	root = stack.pop(-1)	
    	res.append(root.val)
    	root = root.right	
    print(res)
    return res

#BST二叉搜索树
def helper_BST(node, l, index):
    if node:
        left_l = l[0:index]
        right_l = l[index+1:]
        if left_l:
            l_index = len(left_l)//2
            left = TreeNode(left_l[l_index], None, None)
            node.left = left
            helper_BST(left, left_l, l_index)

        if right_l:
            r_index = len(right_l)//2
            right = TreeNode(right_l[r_index], None, None)
            node.right = right
            helper_BST(right, right_l, r_index)

def create_BST(l):
    index = len(l)//2
    if not index:
        return None
    root = TreeNode(l[index], None, None)
    
    helper_BST(root, l, index)
    return root

#镜像反转
def get_node_list(node, l):
    if node:
        if node.left:
            get_node_list(node.left, l)
        l.append(node)
        if node.right:
            get_node_list(node.right, l)

def reverse(l):
    for node in l:
        if node.left or node.right:
            node.right, node.left = node.left, node.right


#叶子节点相同
def get_tree_lef(node, l):
    if node:
        if node.left:
            get_tree_lef(node.left, l)
        if (not node.left) and (not node.right):
            print(node.val)
            l.append(node.val)
        if node.right:
            get_tree_lef(node.right, l)

#create right tree
def helper_right(node, l):
    if not l:
        return
    if node:
        right = TreeNode(l.pop(0), None, None)
        node.right = right

        if not l:
            return
        helper(right, l)

def create_right(l):
    root = TreeNode(l.pop(0), None, None)
    helper_right(root, l)
    return root

if __name__=="__main__":
    #root = create([1,2,3,4])
    #get_node(root)
    #print("---------")
    #iter_node(root)
    
    #BST
    #root = create_BST([-10,-3,0,5,9])
    #get_node(root)
    
    #root = create_BST([1,2,4,3,6,7,9])
    #l = []
    #get_node_list(root, l)
    #reverse(l)
    #get_node(root)
    
    #root1 = create_BST([1,2,4,3,6,7,9])
    #l1 = []
    #get_tree_lef(root1, l1)
    #print(l1)
    #l2 = []
    #get_tree_lef(root2, l2)
 
    root = create([5,3,6,2,4,None,8,1,None,None,None,7,9])
    l = []
    get_node_list(root, l)
    

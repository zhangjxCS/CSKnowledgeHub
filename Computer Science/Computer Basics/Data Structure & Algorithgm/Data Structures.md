## Tree

### Binary Tree

Construct binary tree from two of preorder traverse, inorder traverse, postorder traverse

```python
# Construct from preorder and postorder
class Solution:
    def constructFromPrePost(self, preorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        if not preorder or not postorder:
            return None
        root = TreeNode(preorder[0])
        preorder.pop(0)
        postorder.pop()
        if not preorder or not postorder:
            return root
        if preorder[0] == postorder[-1]:
            root.left = self.constructFromPrePost(preorder, postorder)
        else:
            index = 0
            while preorder[index] != postorder[-1]:
                index += 1
            root.left = self.constructFromPrePost(preorder[:index], postorder[:index])
            root.right = self.constructFromPrePost(preorder[index:], postorder[index:])
        return root

# Construct from preorder and inorder
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if len(preorder) == 0:
            return None
        rootval = preorder[0]
        index = 0
        while index < len(inorder):
            if rootval == inorder[index]:
                break
            index += 1
        root = TreeNode(rootval)
        leftlength = index
        rightlength = len(inorder) - index - 1
        root.left = self.buildTree(preorder[1:leftlength+1], inorder[:leftlength])
        root.right = self.buildTree(preorder[leftlength+1:], inorder[index+1:index+1+rightlength])
        return root

# Construct from inorder and postorder
class Solution:
    def buildTree(self, inorder: List[int], postorder: List[int]) -> Optional[TreeNode]:
        if not inorder or not postorder:
            return None
        root = TreeNode(postorder[-1])
        index = 0
        while index < len(inorder):
            if inorder[index] == postorder[-1]:
                break
            index += 1
        root.left = self.buildTree(inorder[:index], postorder[:index])
        root.right = self.buildTree(inorder[index+1:], postorder[index:-1])
        return root
```

Searialize and Desearilize: Transfer a data structure to string for web communication

```python
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
        
class Codec:
    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        res = []
        def helper(root):
            if not root:
                res.append('#')
                return
            res.append(str(root.val))
            helper(root.left)
            helper(root.right)
        helper(root)
        return ','.join(res) 

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        :type data: str
        :rtype: TreeNode
        """
        nodes = data.split(',')
        def helper():
            if len(nodes) == 0:
                return None
            if nodes[0] == '#':
                nodes.pop(0)
                return None
            node = TreeNode(nodes.pop(0))
            node.left = helper()
            node.right = helper()
            return node
        return helper()
```

### Binary Search Tree

Delete Node

```python
class Solution:
    def deleteNode(self, root: Optional[TreeNode], key: int) -> Optional[TreeNode]:
        if not root:
            return None
        elif root.val < key:
            root.right = self.deleteNode(root.right, key)
        elif root.val > key:
            root.left = self.deleteNode(root.left, key)
        else:
            if root.left and root.right:
                node = self.findleftmax(root)
                root.val = node.val
                root.left = self.deleteNode(root.left, node.val)
            elif root.right:
                root = root.right
            elif root.left:
                root = root.left
            else:
                root = None
        return root
                
        
    def findleftmax(self, root):
        node = root.left
        while node.right:
            node = node.right
        return node
```

Insert Node

```python
class Solution:
    def insertIntoBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        if not root:
            return TreeNode(val)
        if root.val < val:
            root.right = self.insertIntoBST(root.right, val)
        else:
            root.left = self.insertIntoBST(root.left, val)
        return root
```

### Prefix Tree (Trie)

A **trie** or **prefix tree** is a tree data structure used to efficiently store and retrieve keys in a dataset of strings.

```python
class TrieNode:
  
    def __init__(self):
        self.child = collections.defaultdict(TrieNode)
        self.isword = False

class Trie:

    def __init__(self):
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        curr = self.root
        for w in word:
            curr = curr.child[w]
        curr.isword = True

    def search(self, word: str) -> bool:
        curr = self.root
        for w in word:
            curr = curr.child.get(w)
            if not curr:
                return False
        return curr.isword
        
    def startsWith(self, prefix: str) -> bool:
        curr = self.root
        for w in prefix:
            curr = curr.child.get(w)
            if not curr:
                return False
        return True
```

### Huffman Tree

```python
import heapq

class node:
    def __init__(self, freq, symbol):
        self.freq = freq
        self.symbol = symbol
        self.left = None
        self.right = None
        self.huff = ''

class HuffmanTree():
    def __init__(self, A, P):
        nodes = [node(P[i], A[i]) for i in range(len(A))]
        while len(nodes) > 1:
            nodes.sort(key=lambda x:x.freq, reverse=True)
            parent = nodes(nodes[-1].freq+nodes[-2].freq, nodes[-1].symbol+nodes[-2].symbol)
            parent.right = nodes.pop()
            parent.left = nodes.pop()
        self.root = parent

    def Huffman(self, root, val=''):
        if not root:
            return
        root.huff = val
        self.Huffman(root.left, val+'0')
        self.Huffman(root.right, val+'1')
```



## Priority Queue

### Heap

- Every level is filled except the leaf nodes
- Every node has a maximum of 2 children
- All the nodes are as far left as possible, this means that every child is to the left of his parent
- Max(Min) property: the key at the parent node is always greater(smaller) than the key at both child nodes

Heap Operations:

- heapify: rearranges the elements in the heap to maintain the heap property, takes O(n) time
- insert: adds an item to a heap while maintaining its heap property, takes O(log n) time
- delete: removes top item in a heap, takes O(log n) time
- getMax, getMin: returns the maximum of minimum value in a heap

```python
class MinHeap:
  	def __init__(self, maxsize):
      	self.maxsize = maxsize
        self.size = 0
        self.heap = [0] * (self.maxsize + 1)
    
    def isLeaf(self, pos):
      	return 2 * pos > self.size
      
    def helper(self, pos)
        if self.isLeaf(pos):
          	return
        if self.heap[pos] < self.heap[2*pos] and self.heap[pos] < self.heap[2*pos+1]:
          	return
        if self.heap[2*pos] > self.heap[2*pos+1]:
          	self.heap[pos], self.heap[2*pos] = self.heap[2*pos], self.heap[pos]
          	self.helper(2*pos)
        else:
          	self.heap[pos], self.heap[2*pos+1] = self.heap[2*pos+1], self.heap[pos]
          	self.helper(2*pos+1)

    def heapify(self):
        for pos in range(self.size//2, 0, -1):
          	self.helper(pos)
    
    def insert(self, element):
      	if self.size >= self.maxsize:
          	return
        self.size += 1
        self.heap[self.size] = element
        curr = self.size
        while self.heap[curr] < self.heap[curr//2]:
          	self.heap[curr], self.heap[curr//2] = self.heap[curr//2], self.heap[curr]
            curr = curr // 2
    
    def delete(self):
      	self.heap[1], self.heap[self.size] = self.heap[self.size], self.heap[1]
        self.size -= 1
        self.helper(1)
    
    def getMin(self):
      	return self.heap[1]
				
```

## Design Data Structure

### LRU Cache

- LRUCache(capacity): Initialize the LRU cache with positive size capacity
- get(key): return the value of the key if the key exists, otherwise return -1
- put(key, value): update the value of the key if key exists, otherwise add key-value pair to the cache. If the number of keys exceeds, evict the least recently used key

```python
class Node:
    def __init__(self, key=0, value=0, prev=None, next=None):
        self.key = key
        self.value = value
        self.prev = prev
        self.next = next
        
class LinkedList:
    def __init__(self):
        self.head = Node()
        self.tail = Node()
        self.head.next = self.tail
        self.tail.prev = self.head
        self.size = 0
    
    def _add(self, node):
        node.next = self.head.next
        node.prev = self.head
        self.head.next.prev = node
        self.head.next = node
        self.size += 1
    
    def _remove(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev
        self.size -= 1
    
    def _pop(self):
        node = self.tail.prev
        self._remove(node)
        return node
    
class LRUCache:

    def __init__(self, capacity: int):
        self.cache = {}
        self.capacity = capacity
        self.lst = LinkedList()

    def get(self, key: int) -> int:
        node = self.cache.get(key, None)
        if not node:
            return -1
        self.lst._remove(node)
        self.lst._add(node)
        return node.value
        
    def put(self, key: int, value: int) -> None:
        node = self.cache.get(key, None)
        if not node:
            newnode = Node(key, value)
            self.cache[key] = newnode
            self.lst._add(newnode)
            if self.lst.size > self.capacity:
                tail = self.lst._pop()
                del self.cache[tail.key]
        else:
            node.value = value
            self.lst._remove(node)
            self.lst._add(node)
```

### LFU Cache

- LFUCache(capacity): initialize the object with capacity of the data structure
- get(key): gets the value of the key if the key exists in the cache, otherwise return -1
- put(key, value): update the value of the key if present, or insert the key if not present. When reach the capacity, remove the least frequently used key,  if there is a tie, return the least recently used one

```python
class Node:
    def __init__(self, key=0, value=0, prev=None, next=None):
        self.key = key
        self.value = value
        self.prev = prev
        self.next = next
        self.freq = 1
        
class LinkedList:
    def __init__(self):
        self.head = Node()
        self.tail = Node()
        self.head.next = self.tail
        self.tail.prev = self.head
        self.size = 0
    
    def _add(self, node):
        node.next = self.head.next
        node.prev = self.head
        self.head.next.prev = node
        self.head.next = node
        self.size += 1
    
    def _remove(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev
        self.size -= 1
    
    def _pop(self):
        node = self.tail.prev
        self._remove(node)
        return node
    
class LFUCache:

    def __init__(self, capacity: int):
        self.cache = {}
        self.freqTable = collections.defaultdict(LinkedList)
        self.capacity = capacity
        self.minFreq = 0
    
    def _update(self, node, key, value):
        freq = node.freq
        node.value = value
        node.freq += 1
        self.freqTable[freq]._remove(node)
        self.freqTable[node.freq]._add(node)
        if freq == self.minFreq and self.freqTable[freq].size == 0:
            self.minFreq += 1

    def get(self, key: int) -> int:
        node = self.cache.get(key, None)
        if not node:
            return -1
        self._update(node, node.key, node.value)
        return node.value

    def put(self, key: int, value: int) -> None:
        if not self.capacity:
            return
        node = self.cache.get(key, None)
        if node:
            self._update(node, key, value)
        else:
            newnode = Node(key, value)
            if len(self.cache) == self.capacity:
                removenode = self.freqTable[self.minFreq]._pop()
                del self.cache[removenode.key]
            self.freqTable[1]._add(newnode)
            self.cache[key] = newnode
            self.minFreq = 1
```

### Max Stack

```python
class MaxStack:
  	'''
  	Maintain two stack, one for the item, and the other for the max item so far
  	For popmax, first pop the item until the maxitem, and then reload the item in the stack
  	'''

    def __init__(self):
        self.stack = []

    def push(self, x: int) -> None:
        m = self.stack[-1][1] if self.stack else float('-inf')
        self.stack.append([x, max(x, m)])

    def pop(self) -> int:
        return self.stack.pop()[0]


    def top(self) -> int:
        return self.stack[-1][0]

    def peekMax(self) -> int:
        return self.stack[-1][1]

    def popMax(self) -> int:
        lst = []
        m = self.stack[-1][1]
        while self.stack and self.top() != m:
            lst.append(self.pop())
        self.pop()
        for item in reversed(lst):
            self.push(item)
        return m
```

### Calculator


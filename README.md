# Welcome to AlgorithmsGo
This contains sample collection of data structures and algorithms in Go language. These are implemented as tutorial lessons.
It is work in progress. 

# Fundamental Algorithms
---
## [Mathematical Preliminaries](/mathematics)
  - [Euclid's Algorithm](/mathematics/euclidean.go)
  - [Matrix Manipulations](/mathematics/matrix.go)
  Matrix addition, subtraction, multiplication, transpose operations
## Information Structures
### [Data Structures](/datastructures)
  - [Binary Tree](/datastructures/binaryTree/binaryTree.go)
  - [Doubly Linked List](/datastructures/linkedList/doubleLinkedList.go)
  - [Priority Queue](/datastructures/priorityQueue.go)
  - [Queue](/datastructures/queue.go)
  - [Red Black Tree](/datastructures/redBlackTree/redBlackTree.go)
  - [Singly Linked List](/datastructures/linkedList/linkedList.go)
  - [Stack](/datastructures/stack.go)
## Concurrency - multi core processing
### [Safe collections](/safecollections)
  - [SafeList](/safecollections/safeList.go)
  - [SafeMap](/safecollections/safeMap.go)
### [Concurrent execution](/concurrency)
  - [WorkerPool](/concurrency/workerPool.go)
# Seminumerical Algorithms
---
## Random Numbers
### Generating Random Numbers
- [Linear Congruential Method](/random/randomgenerator.go)
- [Additive Congruential Method](/random/randomgenerator.go)
### Statistical Distribution Tests
- [Chi-square distribution](/statistics/discreteSampleSpace.go)
- [The Kolmogorov-Smirnov's Distribution](/statistics/continuousSampleSpace.go)

# Sorting and Searching
---
## [Sorting](/sorting)
Sorting algorithms demonstrated using integer lists, because of no generic support is in golang.
- [Heap Sort](/sorting/heapSort.go)
- [Insertion Sort](/sorting/insertionSort.go)
- [Merge Sort](/sorting/mergeSort.go)
- [Quick Sort](/sorting/quickSort.go)

## [Searching](/searching)
  Searching algorithms are basically information retrieval. Basic search problem is locating the record from
the storage. In addition this section includes some interesting algorithms to search information for pattern
matches like matching subsequence or sequence with maximum value.
- [Find Maximum Subarray](/searching/findMaxSubarray.go)
- [Sorted List - Sequential search with skip](/searching/sortedArraySearch.go)
- [Sorted List - Binary search](/searching/sortedArraySearch.go)
- [Find longest common subsequence](/searching/commonSubsequence.go)


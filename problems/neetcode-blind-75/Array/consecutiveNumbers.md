# Longest Consecutive Sequence (LeetCode 128) - Step by Step Walkthrough

## Problem Statement

Given an **unsorted** array of integers `nums`, return the length of the longest consecutive elements sequence.

**You must write an algorithm that runs in O(n) time.**

---

## Input Example
```go
nums = [100, 4, 200, 1, 3, 2]
```



# 🧠 Key Clues to Understand the Problem

## ❗ Important Observations
- The order of numbers in the array doesn't matter.
- Duplicates can exist, but we care about unique numbers only.
- Need to find **sequences** of numbers, NOT subsequences or subarrays.

## 🧠 Brainstorm to Write Pseudocode
If you are stuck, here’s how a top-level DSA expert thinks 👇:

### 🎯 CLUE 1:
You need O(n) solution →
You **CANNOT sort** the array (sorting takes O(n log n)).

### 🎯 CLUE 2:
You need fast lookup (check if a number exists in O(1) time).
→ Use a HashSet (like a map without values).

### 🎯 CLUE 3:
How to find **start of a sequence**?

A number `x` is the start of a sequence **if `x-1` is not in the set**.

Then you can try to **expand from `x`**, checking `x+1`, `x+2`, `x+3`, etc.



---

## Step 1: Store all numbers in a Set
```go
set := {100, 4, 200, 1, 3, 2}
```

- O(1) lookup time for checking if a number exists.

---

## Step 2: Iterate through each number

### 1. Check `100`:
- Is `100-1 (99)` in set? ❌ No
- Start of a new sequence!
- Initialize:
```go
currentNum = 100
currentStreak = 1
```
- Check `set[100+1]` (101)? ❌ No.
- Update longest streak:
```go
longest = max(0, 1) = 1
```

---

### 2. Check `4`:
- Is `4-1 (3)` in set? ✅ Yes.
- Not start of a sequence → Skip.

---

### 3. Check `200`:
- Is `200-1 (199)` in set? ❌ No
- Start of a new sequence!
- Initialize:
```go
currentNum = 200
currentStreak = 1
```
- Check `set[200+1]` (201)? ❌ No.
- Update longest streak:
```go
longest = max(1, 1) = 1
```

---

### 4. Check `1`:
- Is `1-1 (0)` in set? ❌ No
- Start of a new sequence!
- Initialize:
```go
currentNum = 1
currentStreak = 1
```
- Expand sequence:
    - `2` in set ✅ → `currentNum = 2`, `currentStreak = 2`
    - `3` in set ✅ → `currentNum = 3`, `currentStreak = 3`
    - `4` in set ✅ → `currentNum = 4`, `currentStreak = 4`
    - `5` in set ❌ Stop.
- Update longest streak:
```go
longest = max(1, 4) = 4
```

---

### 5. Check `3`:
- Is `3-1 (2)` in set? ✅ Yes.
- Not start of sequence → Skip.

### 6. Check `2`:
- Is `2-1 (1)` in set? ✅ Yes.
- Not start of sequence → Skip.

---

# Final Answer

```go
return longest = 4
```
(Sequence: `1 → 2 → 3 → 4`)

---

## Why check `num-1 not in set`?
- If `num-1` exists, that means this number is **already part of** an earlier sequence.
- Only start expanding from the **smallest** number of the sequence.

---

## Quick Summary Table

| Step | Check | Action | Streak Updated? | Longest Updated? |
|-----|------|--------|-----------------|------------------|
| 100 | 99 not in set | Start sequence | 1 | 1 |
| 4 | 3 in set | Skip | - | - |
| 200 | 199 not in set | Start sequence | 1 | - |
| 1 | 0 not in set | Start + expand (2,3,4) | 4 | 4 |
| 3 | 2 in set | Skip | - | - |
| 2 | 1 in set | Skip | - | - |

---

# Golang Code

```go
func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num] = true
    }

    longest := 0

    for _, num := range nums {
        if !numSet[num-1] { // start of sequence
            currentNum := num
            currentStreak := 1

            for numSet[currentNum+1] {
                currentNum++
                currentStreak++
            }

            if currentStreak > longest {
                longest = currentStreak
            }
        }
    }

    return longest
}
```

---

# Time and Space Complexity
- **Time:** O(n)
- **Space:** O(n)

---

# Cheat Sheet to Remember
- Use a Set for O(1) lookups.
- Only start counting a sequence if `num-1` is NOT in set.
- Expand forward to `num+1`, `num+2`, etc.
- Update longest streak as you find longer sequences.
- Works even if array is unsorted!

---


# 🧩 Leetcode 153: Find Minimum in Rotated Sorted Array

---

## ✅ Problem Statement (Beginner Friendly)

You're given an array of **unique integers**, `nums`, which was **originally sorted in ascending order**, but then **rotated at an unknown pivot index**.

### 🔁 Rotation Example

```
Original: [1, 2, 3, 4, 5, 6, 7]  
Rotated : [4, 5, 6, 7, 1, 2, 3]
```

- Every number still appears exactly once
- The array is still **partially sorted**

### 🎯 Goal

> Find and return the **minimum element** in this rotated array.

### 📌 Constraints

- `1 <= nums.length <= 5000`
- `-5000 <= nums[i] <= 5000`
- All values are **unique**
- The array was sorted and rotated exactly once

---

## 🔢 Examples

### Example 1
```
Input:  nums = [3, 4, 5, 1, 2]
Output: 1
```

### Example 2
```
Input:  nums = [4, 5, 6, 7, 0, 1, 2]
Output: 0
```

### Example 3
```
Input:  nums = [1]
Output: 1
```

---

## 💡 What Kind of Problem?

This is a **Binary Search on Answer** problem.

- The array was originally sorted → leverage that.
- Use binary search to find the **"valley"** — the smallest value after rotation.

---

## 🌄 Visual Intuition: “Find the Valley in a Rotated Mountain”

Think of the array like a **mountain range** that was rotated:

```
[4, 5, 6, 7, 0, 1, 2]
            ↓
         🕳 Valley
```

You need to find that **dip** — the smallest point — efficiently.

---

## 🐌 Brute-Force Solution

Loop through all elements and track the minimum.

```go
func findMin(nums []int) int {
    minVal := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] < minVal {
            minVal = nums[i]
        }
    }
    return minVal
}
```

### Time: `O(n)`
### Space: `O(1)`

---

## 🚀 Optimal Plan: Binary Search

### 🧠 Insight:

Even though the array is rotated, it still contains **sorted parts**.

### 🔍 Rule of Thumb:

| Condition                      | Direction                     |
|-------------------------------|-------------------------------|
| `nums[mid] > nums[right]`     | Minimum is in the **right**   |
| `nums[mid] <= nums[right]`    | Minimum is in the **left**    |

---

## 🔁 Dry Run (Visualized)

Using: `[4, 5, 6, 7, 0, 1, 2]`

### Step 1

```
left = 0, right = 6
mid = 3 → nums[mid] = 7, nums[right] = 2

7 > 2 → Go Right → left = mid + 1 = 4
```

### Step 2

```
left = 4, right = 6
mid = 5 → nums[mid] = 1, nums[right] = 2

1 < 2 → Go Left → right = mid = 5
```

### Step 3

```
left = 4, right = 5
mid = 4 → nums[mid] = 0, nums[right] = 1

0 < 1 → Go Left → right = mid = 4
```

Now `left == right` → Return `nums[left] = 0`

---

## ✅ Final Golang Code

```go
func findMin(nums []int) int {
    left := 0
    right := len(nums) - 1

    for left < right {
        mid := (left + right) / 2

        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return nums[left]
}
```

---

## 📊 Time & Space Complexity

| Metric         | Value         |
|----------------|---------------|
| Time           | `O(log n)`    |
| Space          | `O(1)`        |

---

## 🧠 DSA Pattern: Binary Search on Answer

| Concept         | Explanation                                       |
|----------------|---------------------------------------------------|
| Pattern         | Binary Search on Sorted / Rotated Array          |
| Use-case        | Find dip/inflection point                        |
| Tip             | Compare `nums[mid]` with `nums[right]`           |
| Metaphor        | Like finding a valley in a rotated mountain 🌄    |

---

## 🔁 Practice Variations

1. **Leetcode 154** — Find Minimum in Rotated Sorted Array II (with duplicates)
2. **Leetcode 33** — Search in Rotated Sorted Array
3. **Leetcode 81** — Search in Rotated Sorted Array II
4. **Leetcode 162** — Find Peak Element

---

## 🧾 Summary Cheat Sheet

| Step                         | Description                                 |
|------------------------------|---------------------------------------------|
| Check if mid > right         | Go right: `left = mid + 1`                  |
| Else                         | Go left: `right = mid`                      |
| Stop when left == right      | Return `nums[left]`                         |
| Time Complexity              | `O(log n)`                                  |
| Space Complexity             | `O(1)`                                      |

---
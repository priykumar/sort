In-place: An algorithm that sorts the array using only a constant amount of extra memory (O(1) space). Recursion stack space is not considered here  
Stable: An algorithm that preserves the relative order of equal elements  

Sorting Algorithms Comparison Table
| Algorithm       | Best Case    | Average Case | Worst Case   | Space Complexity         | In-place | Stable |
|-----------------|--------------|--------------|--------------|--------------------------|----------|--------|
| Quick Sort      | O(n log n)   | O(n log n)   | O(n²)        | O(log n) avg, O(n) worst | Yes*     | No     |
| Merge Sort      | O(n log n)   | O(n log n)   | O(n log n)   | O(n)                     | No**     | Yes    |
| Insertion Sort  | O(n)         | O(n²)        | O(n²)        | O(1)                     | Yes      | Yes    |
| Bubble Sort     | O(n²)***     | O(n²)        | O(n²)        | O(1)                     | Yes      | Yes    |
| Selection Sort  | O(n²)        | O(n²)        | O(n²)        | O(1)                     | Yes      | No     |

*Quick Sort is "in-place" by practical definition (sorts within original array) but uses O(log n) space for recursion stack  
**Merge Sort can be made in-place but becomes very complex with worse time complexity  
***Bubble Sort can be optimized to O(n) best case with early termination  

| Use Case                     | Recommended Algorithm    | Reason                         |
|------------------------------|--------------------------|--------------------------------|
| Large datasets               | Quick Sort or Merge Sort | O(n log n) average/guaranteed  |
| Nearly sorted data           | Insertion Sort           | O(n) best case, adaptive       |
| Stability required           | Merge Sort               | Stable + O(n log n)            |
| Memory constrained           | Quick Sort               | In-place + fast average case   |
| Simple implementation        | Selection Sort           | Predictable,easy to understand |
| Small arrays (n < 10-50)     | Insertion Sort           | Low overhead, adaptive         |

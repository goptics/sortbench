# SortBench

A project to experiment with benchmark visualization CI/CD pipeline using [Vizb](https://github.com/goptics/vizb) on sorting algorithms.

The `Sort` function in `sort.go` is tagged with different algorithms across the git history:

| Version | Algorithm |
|---------|-----------|
| `v1.0.0` | Bubble sort |
| `v1.1.0` | Insertion sort |
| `v1.2.0` | Shell sort |
| `v1.3.0` | Quicksort |

Each tag represents a different sorting implementation, allowing benchmark comparisons to be run and visualized across algorithm variations.

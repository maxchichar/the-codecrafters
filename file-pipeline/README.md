# CodeCrafters: Operation Gopher Protocol  
## Module: File Pipeline  
**Squad:** Goroutines  



## Overview

This program is part of **SENTINEL’s archive processing system**.  
It reads a messy field report (`input.txt`), applies a sequence of transformation rules, and writes a clean, structured report to `output.txt`.

This pipeline is deterministic given the same input, every squad member’s program must produce **identical output**.


## How to Run

```bash
go run . input.txt output.txt
```
Good topic. Most people *think* they understand this and then design systems that collapse under load. Let’s be precise and ruthless.

---

## 1. Concurrency vs Parallelism — What They Actually Mean

### Concurrency

**Concurrency is about structure.**
It means dealing with **multiple tasks at the same time**, making progress by switching between them.

* Tasks *overlap* in time
* Doesn’t require multiple CPU cores
* Focused on coordination, responsiveness, and waiting (I/O, locks, channels)

In Go:

* Goroutines + channels enable concurrency
* Even on a single core, Go can run many goroutines concurrently

Use concurrency when:

* Tasks spend time waiting (network, disk, user input)
* You want responsiveness and clean task separation

---

### Parallelism

**Parallelism is about execution.**
It means doing **multiple tasks at the exact same time** on different CPU cores.

* Tasks run simultaneously
* Requires multiple cores
* Focused on throughput and CPU utilization

In Go:

* Parallelism depends on `GOMAXPROCS`
* Goroutines run in parallel only if the runtime schedules them on multiple OS threads

Use parallelism when:

* Tasks are CPU-bound
* You want faster computation, not just better structure

---

### The Key Difference (Burn This In)

> Concurrency is about **dealing with many things**.
> Parallelism is about **doing many things**.

You can have:

* Concurrency without parallelism (single-core)
* Parallelism without good concurrency (badly structured code)
* The best systems have both

---

## 2. Simple Code Example

### Concurrent (not necessarily parallel)

```go
package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go task("A")
	go task("B")

	time.Sleep(1 * time.Second)
}
```

What this shows:

* Tasks interleave
* Progress overlaps
* Might run on one core

---

### Parallel (if multiple cores available)

```go
package main

import (
	"fmt"
	"runtime"
)

func work(id int) {
	sum := 0
	for i := 0; i < 1e7; i++ {
		sum += i
	}
	fmt.Println("Worker", id, "done")
}

func main() {
	runtime.GOMAXPROCS(2)

	go work(1)
	go work(2)

	select {}
}
```

What this shows:

* CPU-bound work
* Can run simultaneously on multiple cores
* Throughput increases with cores

---

## 3. Common Mistakes (These Cause Bad Systems)

### Mistake 1: Assuming goroutines = parallelism

Beginners think spawning goroutines automatically uses all CPUs.

Reality:

* Goroutines give **concurrency**
* Parallelism depends on available cores and `GOMAXPROCS`

Avoidance:

* Identify whether the task is I/O-bound or CPU-bound
* Tune parallelism intentionally

---

### Mistake 2: Using concurrency to “speed up” CPU-bound work

Concurrency alone doesn’t make CPU-heavy code faster.

Why:

* Context switching adds overhead
* Too many goroutines can slow things down

Avoidance:

* Limit goroutines for CPU work
* Match worker count to CPU cores

---

### Mistake 3: Ignoring synchronization costs

People parallelize everything and forget shared state.

Why it fails:

* Locks, contention, false sharing
* Parallel code can be slower than sequential

Avoidance:

* Measure before optimizing
* Prefer message passing over shared memory

---

## 4. Real-World Applications

### Scenario 1: Web Servers

* Concurrency: handling thousands of requests without blocking
* Parallelism: processing CPU-heavy requests on multiple cores

Without concurrency → server freezes
Without parallelism → server underutilizes hardware

---

### Scenario 2: Data Pipelines

* Concurrency: reading, parsing, and writing simultaneously
* Parallelism: transforming data chunks in parallel

This is where Go shines if designed correctly.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Write a program that:

* Runs two independent tasks concurrently
* Logs when each task starts and finishes
* Observes interleaving behavior

Goal: understand overlapping execution.

---

### Exercise 2 (Medium)

Create a worker pool that:

* Processes CPU-heavy jobs
* Limits workers to the number of CPU cores
* Measures execution time with different worker counts

Goal: observe parallelism effects.

---

### Exercise 3 (Hard)

Design a pipeline with:

* One stage doing I/O-like work
* One stage doing CPU-heavy work
* Controlled concurrency and parallelism per stage

Goal: separate structure (concurrency) from execution (parallelism).

---

## Thought-Provoking Question

**If adding more goroutines makes your program slower, how do you determine whether the bottleneck is scheduling, synchronization, or poor task decomposition—and what does that say about your system’s design?**

If you can’t answer that confidently, you don’t actually control concurrency yet—you’re just hoping it works.

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func PrintNumbers() {
	for i := range 5 {
		fmt.Println(time.Now())
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func PrintLetters() {
	fmt.Println(time.Now())
	for _, v := range "ABCDE" {
		fmt.Println(string(v))
		time.Sleep(500 * time.Millisecond)
	}
}

func HeavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d is starting \n", id)
	for range 100_000_000 { // underscores can be used to improve readability for big ints
	}
	fmt.Println(time.Now())
	fmt.Printf("Task %d is finished \n", id)
}

func main() {
	// Without wait groups :
	// go PrintNumbers()
	// go PrintLetters()
	// time.Sleep(3 * time.Second)

	// // With wait groups :
	// var wg sync.WaitGroup

	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	PrintNumbers()
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	PrintLetters()
	// }()

	// wg.Wait()

	// Heavy task simulation:
	numthreads := 4
	runtime.GOMAXPROCS(numthreads)
	var wg sync.WaitGroup

	for i := range numthreads {
		wg.Add(1)
		go HeavyTask(i, &wg)
	}
	wg.Wait()
}

/*
	The difference boils down to **what’s happening conceptually** vs. **what’s happening physically**.

---

### **1. Concurrency** – *"Managing multiple things at once"*

* **Definition:** Structuring a program so that multiple tasks *appear* to be running at the same time, but they might actually take turns sharing the same CPU.
* **Key point:** It’s about **dealing with** many tasks at once, not necessarily doing them *literally* at the same time.
* **Example:** A single chef cooking three dishes by rapidly switching between them — stir pasta, flip steak, chop salad — but only using **one pair of hands**.
* **Implementation:** Often involves **context switching**, cooperative scheduling, or asynchronous programming.

---

### **2. Parallelism** – *"Actually doing multiple things at once"*

* **Definition:** Running multiple tasks **at the same time** on different processors or cores.
* **Key point:** It’s about **speed** through simultaneous execution.
* **Example:** Three chefs each working on a separate dish at the same time — pasta, steak, salad — all progressing independently.
* **Implementation:** Requires multiple CPU cores or machines, often using threads or processes that run simultaneously.

---

### **Side-by-side table:**

| Feature           | Concurrency                          | Parallelism                                |
| ----------------- | ------------------------------------ | ------------------------------------------ |
| **Goal**          | Manage multiple tasks                | Execute multiple tasks simultaneously      |
| **Time**          | Tasks *overlap* in time              | Tasks *occur* at the same time             |
| **Hardware need** | Can run on 1 core                    | Needs multiple cores/CPUs                  |
| **Analogy**       | One person juggling many jobs        | Many people doing jobs at the same time    |
| **Example tech**  | async/await, goroutines, event loops | SIMD, multi-threading, distributed systems |

---

💡 **Key takeaway:**

* **Concurrency** is about *structure* — making progress on many things at once.
* **Parallelism** is about *execution* — making progress on many things at the same instant.

If you want, I can sketch a **diagram showing concurrency vs. parallelism timelines** — that tends to make the difference instantly click.

*/

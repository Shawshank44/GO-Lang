Good. Rate limiting is one of those topics where **people memorize algorithms but don’t understand trade-offs**, then ship something that fails under real traffic. I’ll keep this tight and honest.

---

## 1. Rate Limiters — What They Are and Why They Exist

**Purpose:**
Rate limiters control **how many actions are allowed over time** to protect systems from overload, abuse, or unfair usage.

They are commonly used when:

* You expose APIs
* You consume third-party services
* You protect shared resources (DBs, queues, CPUs)

### The three you mentioned — brutally summarized:

### 1️⃣ Fixed Window Counter

* Count requests in a fixed time window (e.g., 100 req/min)
* Simple, fast, **but inaccurate at boundaries**

Use when:

* Precision is not critical
* You want simplicity over correctness

---

### 2️⃣ Token Bucket

* Tokens are added at a steady rate
* Each request consumes a token
* Allows short bursts

Use when:

* You want smooth average rate **with burst tolerance**
* Most real-world APIs use this

---

### 3️⃣ Leaky Bucket

* Requests enter a bucket
* Processed at a fixed rate (leaks steadily)
* Excess requests are dropped or delayed

Use when:

* You want **constant output rate**
* Bursts must be flattened completely

---

## 2. Simple Code Examples (Minimal, Not Production-Ready)

### Fixed Window Counter (naive)

```go
var count int
var windowStart = time.Now()

func allow() bool {
	if time.Since(windowStart) > time.Minute {
		windowStart = time.Now()
		count = 0
	}
	if count >= 5 {
		return false
	}
	count++
	return true
}
```

---

### Token Bucket (using time.Ticker)

```go
tokens := make(chan struct{}, 5)

go func() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		select {
		case tokens <- struct{}{}:
		default:
		}
	}
}()

func allow() bool {
	select {
	case <-tokens:
		return true
	default:
		return false
	}
}
```

---

### Leaky Bucket (queue + fixed drain)

```go
queue := make(chan struct{}, 5)

go func() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		select {
		case <-queue:
		default:
		}
	}
}()

func allow() bool {
	select {
	case queue <- struct{}{}:
		return true
	default:
		return false
	}
}
```

These are **concept demos**, not production-safe. If you treat them as final, you will create bugs.

---

## 3. Common Mistakes (Where People Screw This Up)

### Mistake 1: Choosing the wrong algorithm

Beginners pick randomly.

Reality:

* Fixed window → unfair bursts
* Token bucket → burst-friendly
* Leaky bucket → strict smoothing

Avoidance:
Choose based on **traffic pattern**, not tutorial popularity.

---

### Mistake 2: Ignoring concurrency

Rate limiters are shared state.

Bad:

* No mutex
* No atomics
* Race conditions everywhere

Avoidance:
Protect counters properly or use channels intentionally.

---

### Mistake 3: Assuming rate limiting = security

It’s not.

Rate limiting:

* Reduces load
* Slows abuse

It does NOT:

* Prevent attacks
* Replace authentication
* Stop distributed abuse

Avoidance:
Treat rate limiting as **one layer**, not a solution.

---

## 4. Real-World Applications

### Scenario 1: API Gateway

* Per-user or per-IP limits
* Prevents one client from starving others

Token bucket is the usual choice.

---

### Scenario 2: Downstream Dependency Protection

* Protect databases
* Protect third-party APIs

Leaky bucket works well to smooth spikes.

---

## 5. Practice Exercises (No Solutions)

### Exercise 1 (Easy)

Implement a fixed window rate limiter that allows N requests per second for a single goroutine.

---

### Exercise 2 (Medium)

Build a token bucket rate limiter that:

* Allows bursts
* Is safe under multiple goroutines

---

### Exercise 3 (Hard)

Design a rate limiter that:

* Supports per-user limits
* Works concurrently
* Can switch between token bucket and leaky bucket strategies

If you immediately think “easy”, you haven’t thought about synchronization or memory usage yet.

---

## Thought-Provoking Question

**If your rate limiter drops requests under load, how do you decide whether that failure should be silent, delayed, or visible to the client—and who pays the cost of that decision: your system or your users?**

That answer determines whether your system feels *reliable* or *hostile* under pressure.

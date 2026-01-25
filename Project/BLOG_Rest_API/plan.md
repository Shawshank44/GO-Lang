Good — you’re asking the *right* question. Folder names are cheap; **clear responsibility is hard**. I’ll be blunt and precise. If any description sounds fuzzy, that’s a red flag you should fix *before* writing code.

---

## `/cmd/server/main.go` — **Process bootstrapper**

**Purpose:** Start the application. Nothing else.

What belongs here:

* Load configuration
* Initialize dependencies (DB, router, services)
* Wire everything together
* Start the HTTP server
* Handle graceful shutdown

What **never** belongs here:

* Business logic
* SQL
* Auth rules
* Request handling logic

Mental model:

> `main.go` is the *composition root*, not a brain.

If this file grows beyond ~100 lines, you’re doing too much here.

---

## `/internal/config` — **Source of truth for runtime behavior**

**Purpose:** Decide *how* the app runs in different environments.

Typical responsibilities:

* Reading env vars / config files
* Validating required config
* Providing typed config structs

Key rule:

> No other package should read environment variables directly.

If handlers start calling `os.Getenv`, your architecture is already leaking.

---

## `/internal/db` — **Database lifecycle & connectivity**

**Purpose:** Own *how* you connect to the database, not *what* you do with it.

Responsibilities:

* Open/close DB connections
* Connection pooling
* Migrations (optional but recommended)
* Exposing a DB handle

Not allowed here:

* SQL queries for business logic
* Table-specific logic

Mental model:

> This package knows *how to talk to the database engine*, not *what to say*.

---

## `/internal/models` — **Domain data shapes**

**Purpose:** Represent the core entities of your system.

Contains:

* User
* BlogPost
* Role
* Possibly DTO-like structs shared across layers

Rules:

* No SQL
* No HTTP knowledge
* Minimal tags (avoid coupling to frameworks)

These structs should make sense even if the API didn’t exist.

If a struct only exists to satisfy JSON output, question its existence.

---

## `/internal/repositories` — **Persistence logic (SQL lives here)**

**Purpose:** Translate domain intent into database operations.

Responsibilities:

* CRUD queries
* Transactions
* Mapping DB rows → models

Rules:

* No HTTP concerns
* No auth rules
* No business decisions

Repositories answer questions like:

> “How do I fetch published posts by this author?”

They should **not** decide whether the caller *is allowed* to do so.

---

## `/internal/services` — **Business rules & decision-making**

**Purpose:** Enforce *what is allowed* in your system.

Responsibilities:

* Authorization checks
* Ownership validation
* Workflow rules (draft → published)
* Orchestration of multiple repositories

This is where your system’s **intelligence** lives.

Rule of thumb:

> If you removed HTTP and SQL, services should still make sense.

If services feel thin, your logic is leaking elsewhere.

---

## `/internal/handlers` — **HTTP translation layer**

**Purpose:** Convert HTTP into domain calls and back.

Responsibilities:

* Parse request data
* Validate inputs (shape, not rules)
* Call services
* Return proper HTTP responses

Strict limitations:

* No SQL
* No business rules
* Minimal logic

Handlers should read like:

> “Take request → ask service → return response”

If a handler needs a comment to explain itself, it’s doing too much.

---

## `/internal/middleware` — **Cross-cutting HTTP concerns**

**Purpose:** Apply rules that affect *many* routes uniformly.

Examples:

* Authentication extraction
* Request logging
* Panic recovery
* Rate limiting

Key principle:

> Middleware modifies **context**, not business behavior.

If middleware starts making authorization decisions, you’re mixing layers.

---

## `/internal/auth` — **Identity & trust mechanics**

**Purpose:** Handle *how* identity is established and verified.

Responsibilities:

* Token generation/validation
* Password hashing
* Session invalidation logic
* Role extraction

What doesn’t belong:

* HTTP responses
* Permission rules (that’s service logic)

Auth answers:

> “Who is this request acting as?”

Not:

> “What are they allowed to do?”

---

## `/pkg/utils` — **Reusable, boring helpers**

**Purpose:** Pure utilities with no business meaning.

Examples:

* String helpers
* Time helpers
* Generic validation helpers

Hard rule:

> Nothing in `pkg` should depend on `internal`.

If your app breaks without `utils`, they’re not utilities — they’re core logic misplaced.

---

## Brutal sanity check (don’t skip this)

For **every file**, be able to answer:

1. What layer is this?
2. What problem does it solve?
3. What *must never* be added here?

If you can’t answer #3, the boundary isn’t real.

---


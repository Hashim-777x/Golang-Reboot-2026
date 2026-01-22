# Golang Journey 2026: From Scrap PC to Backend Architect

> **Current Status:** Day 15 of the Repo!✨
> **Location:** Hyderabad, India  
> **Objective:** Senior Architect Sequence (Node.js -> High-Perf Go)

##  The Story: Why This Repo Exists
I am a self-taught Backend Engineer with a hunger for high-performance systems. 

I didn't start with a MacBook. I couldn't afford one. 
**I built my current rig from scrap parts**—an i3 4th Gen wired into a wall-mounted cardboard slot with a custom cooling solution. 

This repository documents my **Non-Negotiable** path to mastering Golang. No skipping steps. No tutorial hell. Just engineering protocols, logic, and distributed systems.

**No fluff. No paid courses. Just shipping.**

---

## 🛠️ The Tech Stack
* **Language:** Golang (Expert Focus), Node.js (Legacy) , rust (learning) 
* **Core:** TCP, WebSockets, gRPC, Concurrency (Goroutines/Channels) , DSA 
* **Database:** PostgreSQL (sqlc), MongoDB, Redis
-+
  * **Infrastructure:** Docker, Kubernetes, AWS, RabbitMQ
* **Observability:** Prometheus, Grafana

## 🏆 Proof of Work
* **2nd Runner Up @ DATANYX Hackathon:** Built "Sporacle" – A real-time IoT mushroom monitoring system.
    * *Role:* Backend dev & (System design) and DB design & setup.

---

##  The Roadmap

### Phase 1: The "Theory & Syntax" Foundation 
*Goal: Understand the "Go Way" (Memory, Pointers, Interfaces). No AI coding.*
- [✅] **A Tour of Go:** Finish every exercise.
- [✅] **Effective Go:** Deep dive into "Pointers vs. Values" and "Interfaces".
- [✅] **Go Spec:** Understand Slices internals.

### Phase 2: Networking & The "Socket" Level
*Goal: Understand how data moves before HTTP abstracts it.*
- [ ] **Project:** Raw TCP-based Key-Value Store (Mini-Redis).
    * *Resource:* Simple Redis Server From Scratch In Golang (Anthony GG).
    * *Focus:* TCP handshakes, buffers, raw bytes.

### Phase 3: Standard Library & Clean Architecture
*Goal: Build production-grade APIs without "magic" frameworks.*
- [ ] **Project:** REST API using only `net/http`.
    * *Resource:* Complete Backend Engineering Course in Go (Tiago).
    * *Focus:* Repository Pattern, Dependency Injection, Std Lib routing.

### Phase 3.1: The "Business Logic" Project
*Goal: Master complex DB relationships and project organization.*
- [ ] **Project:** Restaurant Management Backend.
    * *Resource:* Golang Restaurant Management Backend Project (Akhil Sharma).
    * *Focus:* JWT Auth, Relational Schema Design, Massive Project Structure.

### Phase 3.2: The "Real-Time" Project
*Goal: Master WebSockets and persistent connections.*
- [ ] **Project:** Full-Stack Real-Time Chat App.
    * *Resource:* Building a Chat with WebSockets and HTMX in Golang (Tiago).
    * *Focus:* "Rooms," Broadcasting, persistent connections.

### Phase 3.3: The "Concurrency" Project
*Goal: Master Goroutines, Proxies, and Crawling.*
- [ ] **Project:** Distributed Web Scraper/Crawler.
    * *Resource:* How to Build a Web Scraper in Golang (Boot.dev / Anthony GG).
    * *Focus:* Worker Pools, Mutexes, IP Rotation, Rate Limiting.

### Phase 4: The "Production API" (The Shift to sqlc)
*Goal: 12-Factor Apps & Professional DB Management.*
- [ ] **Project:** E-commerce Backend API.
    * *Resource:* Building a Production API in Golang from Scratch (Tiago).
    * *Focus:* sqlc (Type-safe SQL), Migrations, Structured Logging (slog).

### Phase 5: Microservices & Distributed Systems
*Goal: Communication at Scale.*
- [ ] **Project:** Massive Microservices System.
    * *Resource:* Building a Massive Golang Microservices Project (Akhil Sharma).
    * *Focus:* gRPC (Protobufs), Service Discovery, RabbitMQ.

### Phase 6: Observability & The Performance Pipeline
*Goal: "Fiber + Grafana" vision. Monitoring what you build.*
- [ ] **Project:** High-Performance Data Pipeline.
    * *Resource:* Master Microservices & Event-Driven Architecture (Akhil Sharma).
    * *Focus:* Prometheus metrics, Grafana Dashboards, pprof Profiling.

--- 

## 📝 Daily Log
*Detailed updates posted daily on [X (Twitter)](https://x.com/KillerPand34973).*

| Day | Date | Focus | Key Learning/Build |
| :--- | :--- | :--- | :--- |
| **01** | Jan 01 | **Repo Setup** | Did a research for learning Golang with best quality of projects and made a Master Roadmap. |
| **02** | Jan 02 | ReadMe file | implemented roadmap by defining my step by step plan in ReadMe file and also made a proper folder structure!  |
| **03** | Jan 03 | First.go | i learned and implemented (from **Tour of Go**) - Vars, Types, Funcs, Returns |
| **04** | Jan 04 | second.go |  Mastered Go's for loops (including while-style), if with short statements, and switch logic |
| **05** | Jan 05 | third.go |  Mastered Pointers and Fixed-size Arrays |
| **06** | Jan 06 | Forth.go | Slices, Len vs Cap, and the make function |
| **07** | Jan 08 | fifth.go | Maps, Closures, and Restarting the Grind |
| **08** | Jan 09 | sixth.go | Attaching logic to structs via Receivers. Implemented the Stringer interface for custom formatting.|
| **09** | Jan 10 | seventh.go | Mastered implicit interface implementation and custom error types. The foundation for Phase 1 is 90% complete |
| **10** | Jan 11 | eighth.go |  Finished Tour of Go - Mastered Goroutines and Channels |
| **11** | Jan 12 | effective_go.go |  Effective Go - Allocation, Literals, and Defer |
| **12** | Jan 13 | language_spec.go | Deep dive into Slice internals and memory layout. Understanding the 24-byte SliceHeader. |
| **13** | Jan 14 | language_spec.go | updated and completed the language specs |
| **14** | Jan 15 | tcp_client.go | implemented tcp connection in go |
| **15** | Jan 17 | tcp-client.go | updated the code and optimized |
| **16** | Jan 18 | tcp-client.go | Refactored TCP client with production patterns: separated I/O functions, per-operation timeouts, buffered writes, and proper error context. |


---

## 🤝 Connect
I am documenting this entirely in public. If you are a Recruiter, CTO, or Senior Engineer in Hyderabad, look at the code. It speaks for itself.

* **X (Twitter):** [@KillerPand34973](https://x.com/KillerPand34973)
* **LinkedIn:** www.linkedin.com/in/syed-hashim721

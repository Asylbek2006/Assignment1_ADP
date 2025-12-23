

# Assignment 1: Go Basics

### Packages, Methods & Interfaces | Advanced Programming 1

This project implements four distinct modules demonstrating the core concepts of the **Go (Golang)** programming language, including package management, interfaces, structs, methods, and collection types like maps and slices.

---

## 👤 Student Information

* **Full Name:** Asylbek Abdibay
* **Group:** SE2427
* **Course:** Advanced Programming 1
* **Date:** 23.12.2025

---

## 📂 Project Structure

The project follows a modular package structure as required by the assignment instructions:

```text
Assignment1/
├── go.mod            # Module definition
├── main.go           # Entry point (Menu and Demo)
├── Hotel/            # Task 1: Hotel Room Reservation System
│   └── hotel.go
├── Employee/         # Task 2: Employee Management & Interfaces
│   └── employee.go
├── Gym/              # Task 3: Membership Management System
│   └── gym.go
└── Wallet/           # Task 4: Digital Wallet Simulation
    └── wallet.go

```

---

## 🚀 How to Run

To compile and run the entire project, ensure you are in the root directory (`Assignment1/`) and execute:

```bash
go run .

```

---

## 🛠 Features & Tasks

### 1. Hotel Room Reservation System

* **Package:** `Hotel`
* **Concepts:** Structs, Maps, Pointer Receivers.
* **Logic:** Manages room bookings using a `map[int]Room`. Includes features to add rooms, check-in guests, check-out, and filter vacant rooms.

### 2. Employee & Interfaces

* **Package:** `Employee`
* **Concepts:** Interfaces, Polymorphism, Slices.
* **Logic:** Implements an `Employee` interface with `CalculateSalary()` and `GetWorkHours()` methods. Supports multiple employee types: `FullTime`, `PartTime`, `Contractor`, and `Intern`.

### 3. Membership Management System

* **Package:** `Gym`
* **Concepts:** Interfaces, Map storage.
* **Logic:** Manages gym memberships (Basic and Premium) through a unified `Member` interface. Members are stored and retrieved using a `map[uint64]Member`.

### 4. Digital Wallet Simulation

* **Package:** `Wallet`
* **Concepts:** Methods, Slices, Loops.
* **Logic:** Simulates a digital wallet that tracks a balance and a history of transactions. Includes error checking for insufficient funds during spending.

---

## 🎓 Key Go Concepts Demonstrated

* **Interfaces:** Used to achieve polymorphism, allowing different types to be handled by the same logic (e.g., iterating over different employee types).
* **Maps vs. Slices:** Maps are utilized for high-performance lookups (Rooms, Gym Members), while Slices are used for ordered data (Transaction history, Employee lists).
* **Method Receivers:** Both value and pointer receivers are used appropriately to either read data or modify the internal state of a struct.
* **Encapsulation:** Logic is separated into packages to maintain a clean and maintainable codebase.

---

## 📝 Compliance

* **Project Structure:** Follows the mandatory layout.
* **Naming Conventions:** Adheres to standard Go naming conventions (CamelCase).
* **Clean Code:** No unused variables, clear logic, and meaningful variable naming.

---

*Submitted as part of the Advanced Programming 1 curriculum.*

---

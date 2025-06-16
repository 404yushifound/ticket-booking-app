# ticket-booking-app

# 🎟️ Go Conference Ticket Booking App

Welcome to the **Go Conference Booking Application** – a beginner-friendly console program written in **Golang**. This app allows users to book tickets for a fictional tech conference while demonstrating core concepts of Go like:

- Functions and packages
- Structs and slices
- Maps (for storing ticket info)
- Concurrency using goroutines & wait groups
- Input validation with a separate helper package

---

## 🚀 Features

- Take user input for name, email, and number of tickets
- Validate the input before booking
- Book the tickets and show remaining count
- Send ticket confirmation (simulated with delay using goroutine)
- Use a map to associate first names with ticket counts
- Graceful exit once tickets are sold out

---

## 🧠 Tech Stack

- [x] Golang (Go)
- [x] Go routines for concurrent ticket sending
- [x] `sync.WaitGroup` for graceful waiting
- [x] Modular design (`valid` package for input checking)

---

## 💻 How to Run

1. Clone the repo:

   ```bash
   git clone https://github.com/your-username/go-conference-booking.git
   cd go-conference-booking

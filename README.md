# Nail Appointment Scheduler (Go)

A lightweight, high-performance RESTful API built with **Go (Golang)** for managing salon appointments. This project serves as a practical application of Go concurrency patterns and backend architecture.

## 🚀 Features
* **RESTful API**: 4 core endpoints for managing the appointment lifecycle.
* **Modular Architecture**: Separated into Handlers, Models, and Database layers for maintainability.
* **PostgreSQL Integration**: Persistent data storage for client and appointment records.
* **Go Modules**: Dependency management using `go.mod` and integrity verification via `go.sum`.

## 🛠 Tech Stack
* **Language**: Go (Golang)
* **Framework**: (e.g., Gin-Gonic or Standard Library)
* **Database**: PostgreSQL
* **Tools**: Git, Go Modules

## 📡 API Endpoints
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| GET | `/appointments` | Retrieve all scheduled appointments |
| POST | `/appointments` | Schedule a new nail appointment |
| GET | `/appointments/:id` | Get details of a specific appointment |
| DELETE | `/appointments/:id` | Cancel an existing appointment |

## 🏗 Project Structure
- `main.go`: Application entry point and router initialization.
- `handlers/`: Logic for processing HTTP requests and responses.
- `models/`: Data structures representing the Appointment and Client entities.
- `database/`: Database connection logic and SQL queries.

## ⚙️ Installation & Setup
1. Clone the repository:
   ```bash
   git clone [https://github.com/KonstantinosTheofanousg13/nail-appointment-scheduler.git](https://github.com/KonstantinosTheofanousg13/nail-appointment-scheduler.git)

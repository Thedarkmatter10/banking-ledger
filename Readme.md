# 🏦 Event-Sourced Banking Ledger System

This project is a modular, event-sourced ledger service using:

- **Golang + Gin** for API framework
- **Kafka** for event streaming
- **PostgreSQL** for event store (Write model)
- **Redis** for caching (Read model)
- **CQRS + Event Sourcing** architecture

---

## 🧠 Architecture

### High-Level Flow

             +-----------------+        Write Model (Commands)     +---------------------+
             |                 | --------------------------------> |                     |
  Client --->| Command REST API|                                   |  Command Handler     |
             |                 | <-------------------------------- |                     |
             +-----------------+       Response                    +---------------------+
                       |                          |
                       | Account Exists Middleware|
                       v                          v
                 +----------------+          +---------------+
                 | Input Validator|          | Event Store   | (PostgreSQL)
                 +----------------+          +---------------+
                                                   |
                                                   | Publish to
                                                   v
                                              +----------+
                                              | Kafka     |
                                              +----------+
                                                   |
                                                   v
                                          +------------------+
                                          | Kafka Consumer   |
                                          +------------------+
                                                   |
                                                   v
                                          +---------------------+
                                          | Redis Read Model    |
                                          +---------------------+
                                                   ^
                                                   |
                                             +-----------+
                                             | Query API |
                                             +-----------+
                                                   ^
                                                   |
                                                Client




### ✅ Command Responsibility

Handles incoming requests:

- Validate request
- Create domain event
- Save event to DB
- Publish to Kafka

### 🔎 Query Responsibility

Serves client queries (e.g., balance/history) using fast Redis-based views.

---

## 🧩 Folder Structure

/github.com/Thedarkmatter10/ledger-service
│
├── command/ → Command handlers (CreateAccount, Deposit, Withdraw)
├── query/ → Query handlers (Balance, History)
├── event/ → Kafka producer + consumer
├── model/ → Domain models and event struct
├── repository/ → PostgreSQL event store access
├── cache/ → Redis wrapper
├── middleware/ → Reusable middlewares like AccountExistence check
├── routes/ → Modular route grouping (command/query)
├── projection/ → Event processors to build read models in Redis
├── main.go → Entry point


---

## 🔗 API Endpoints

### Commands

| Method | Endpoint                       | Description           |
|--------|--------------------------------|-----------------------|
| POST   | `/api/v1/accounts`             | Create new account    |
| POST   | `/api/v1/accounts/:id/deposit` | Deposit to account    |
| POST   | `/api/v1/accounts/:id/withdraw`| Withdraw from account |

### Queries

| Method | Endpoint                       | Description           |
|--------|--------------------------------|-----------------------|
| GET    | `/api/v1/accounts/:id/balance` | Get current balance   |
| GET    | `/api/v1/accounts/:id/history` | Transaction history   |

---

## 🧪 Event Types

- `AccountCreated`
- `Deposited`
- `Withdrawn`

---

## 🗃️ Event Model (PostgreSQL Table: `events`)

```sql
CREATE TABLE events (
  id UUID PRIMARY KEY,
  aggregate_id TEXT NOT NULL,
  type TEXT NOT NULL,
  payload JSONB,
  timestamp TIMESTAMP
);

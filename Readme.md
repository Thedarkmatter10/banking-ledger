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
![banking-ledger-flow](assets/<img width="1440" height="1488" alt="architecture" src="https://github.com/user-attachments/assets/5cda5348-652a-4bfc-bd66-5b1971964ebf" />
)



### ✅ Command Responsibility

Handles incoming requests:

- Validate request
- Create domain event
- Save event to DB
- Publish to Kafka

### 🔎 Query Responsibility

Serves client queries (e.g., balance/history) using fast Redis-based views.

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
```
## ❌ Limitations
Due to hardware constraints:
- I did not run Kafka/Redis/PostgreSQL locally
- Instead, I focused on writing clean, testable components

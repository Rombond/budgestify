# Budgestify

Budgestify is an open-source, self-hostable budget management application. It helps users track their income, expenses, and spending patterns with categorized statistics, providing real-time insights into remaining funds after accounting for recurring expenses.

## Features

- **Income Tracking**: Record your salary and its frequency (e.g., monthly, bi-weekly).
- **Expense Management**: Log expenses with their frequency and categorize them (e.g., groceries, utilities, entertainment).
- **Real-Time Budgeting**: Calculate remaining funds by deducting upcoming expenses throughout the month.
- **Statistics**: Visualize spending patterns with categorized breakdowns.
- **Tech Stack**:
  - **Backend**: Go with a MySQL database.
  - **Frontend**: Will be Svelte with shadcn UI components.
  - **Future Plans**: Native mobile apps for Android and iOS post v1.0.

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started) (v20.10 or later)
- [Docker Compose](https://docs.docker.com/compose/) (v2.0 or later)
- [Git](https://git-scm.com/downloads)

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Rombond/budgestify.git
   cd budgestify
   ```

2. **Run the application in development mode**:
   ```bash
   docker compose up
   ```

3. **Run the application in production mode**:
   ```bash
   docker compose -f docker-compose-prod.yml up --build
   ```

4. **Access the application**:
   - Services will be available at the ports defined in the used Docker Compose file.
   - By default:
     - Backend: http://localhost:8080
     - Frontend: not available yet

5. **Stop the application**:
   ```bash
   docker compose down
   # or for production
   docker compose -f docker-compose-prod.yml down
   ```

> **Note**: All dependencies (Go, MySQL, ...) are managed by Docker. No additional local installation is required.
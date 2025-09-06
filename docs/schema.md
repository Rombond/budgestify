# Database Schema

This document describes the database schema for the budget management application, implemented in MySQL. The schema supports tracking salaries, expenses, and remaining budgets with categorization and recurrence. The tables and relationships are defined below, based on the provided table creation code.

## Tables

### User
- **id**: `INT` NOT NULL AUTO_INCREMENT, PRIMARY KEY
- **name**: `VARCHAR(255)` NOT NULL
- **login**: `VARCHAR(255)` NOT NULL
- **hash**: `BINARY(64)` NOT NULL

### House
- **id**: `INT` NOT NULL AUTO_INCREMENT, PRIMARY KEY
- **name**: `VARCHAR(255)` NOT NULL
- **account**: `INT`, FOREIGN KEY REFERENCES `Account(id)`

### Category
- **id**: `INT` NOT NULL AUTO_INCREMENT, PRIMARY KEY
- **name**: `VARCHAR(255)` NOT NULL
- **icons**: `VARCHAR(255)`
- **parent**: `INT`, FOREIGN KEY REFERENCES `Category(id)`
- **house**: `INT` NOT NULL, FOREIGN KEY REFERENCES `House(id)`

### House_User
- **id**: `INT` NOT NULL AUTO_INCREMENT, PRIMARY KEY
- **house**: `INT` NOT NULL, FOREIGN KEY REFERENCES `House(id)`
- **user**: `INT` NOT NULL, FOREIGN KEY REFERENCES `User(id)`
- **admin**: `BOOLEAN` NOT NULL DEFAULT `false`

### Account
- **id**: `INT` NOT NULL AUTO_INCREMENT, PRIMARY KEY
- **name**: `VARCHAR(255)` NOT NULL
- **house_user**: `INT` NOT NULL, FOREIGN KEY REFERENCES `House_User(id)`
- **amount**: `FLOAT` NOT NULL
- **currency**: `CHAR(3)` NOT NULL
- **theoreticalAmount**: `FLOAT`

### Transaction
- **id**: `INT` NOT NULL AUTO_INCREMENT, PRIMARY KEY
- **name**: `VARCHAR(255)` NOT NULL
- **category**: `INT`, FOREIGN KEY REFERENCES `Category(id)`
- **amount**: `FLOAT` NOT NULL
- **payer**: `INT` NOT NULL, FOREIGN KEY REFERENCES `House_User(id)`
- **payer_account**: `INT`, FOREIGN KEY REFERENCES `Account(id)`
- **pay_date**: `DATE` NOT NULL
- **currency**: `CHAR(3)` NOT NULL
- **conversion_rate**: `FLOAT` NOT NULL DEFAULT `1`

### Recurrence
- **id**: `INT` NOT NULL AUTO_INCREMENT, PRIMARY KEY
- **name**: `VARCHAR(255)` NOT NULL
- **house_user**: `INT` NOT NULL, FOREIGN KEY REFERENCES `House_User(id)`
- **payer_account**: `INT` NOT NULL, FOREIGN KEY REFERENCES `Account(id)`
- **category**: `INT`, FOREIGN KEY REFERENCES `Category(id)`
- **amount**: `FLOAT` NOT NULL
- **currency**: `CHAR(3)` NOT NULL
- **conversion_rate**: `FLOAT` NOT NULL DEFAULT `1`
- **pay_date**: `DATE` NOT NULL
- **day_cycle**: `INT` NOT NULL

## Relationships

- **Category.parent** REFERENCES **Category.id**: Supports hierarchical categories (self-referential).
- **Category.house** REFERENCES **House.id**: Associates categories with a house.
- **House.account** REFERENCES **Account.id**: Links a house to an account (optional).
- **House_User.house** REFERENCES **House.id**: Links users to houses.
- **House_User.user** REFERENCES **User.id**: Links users to houses.
- **Account.house_user** REFERENCES **House_User.id**: Accounts are owned by house users.
- **Transaction.category** REFERENCES **Category.id**: Transactions are categorized (optional).
- **Transaction.payer** REFERENCES **House_User.id**: The payer of a transaction.
- **Transaction.payer_account** REFERENCES **Account.id**: The account used for payment (optional).
- **Recurrence.house_user** REFERENCES **House_User.id**: Recurring transactions are linked to a house user.
- **Recurrence.payer_account** REFERENCES **Account.id**: Account used for recurring payments.
- **Recurrence.category** REFERENCES **Category.id**: Recurring transactions are categorized (optional).

## Notes

- The `theoreticalAmount` field in `Account` could benefit from a clearer definition in the documentation (e.g., projected vs. actual balance).

## Recommendations

### Database Design
1. **Currency Validation**: Add a check constraint on `currency` fields to ensure valid ISO 4217 codes (e.g., `CHECK (currency IN ('USD', 'EUR', 'GBP', ...))`).
2. **Amount Precision**: Consider using `DECIMAL(10,2)` instead of `FLOAT` for `amount`, `conversion_rate`, and `theoreticalAmount` to avoid floating-point precision issues in financial calculations.
3. **Indexes**: Create indexes on frequently queried columns (e.g., `Transaction.pay_date`, `Category.house`) to improve performance.

# Contributing to Budgestify

Thank you for your interest in contributing to **Budgestify**, an open-source, self-hostable budget management application! Budgestify helps users track income, expenses, and spending patterns with categorized statistics. This guide outlines how you can contribute to the project.

## Ways to Contribute

We welcome contributions in various forms:
- **Code**: Add new features (e.g., new budgeting tools or UI components) or fix bugs.
- **Documentation**: Improve the `README.md`, code comments, or create guides.
- **Testing**: Write unit tests for the Go backend or Svelte frontend.
- **Bug Reports**: Report issues with clear reproduction steps.
- **Feature Suggestions**: Propose ideas for improving the app (e.g., new statistics visualizations).
- **Translations**: Help make the app multilingual (planned for post-v1.0).

## Setting Up the Development Environment

### Prerequisites
- [Docker](https://www.docker.com/get-started) (v20.10 or later)
- [Docker Compose](https://docs.docker.com/compose/) (v2.0 or later)
- [Git](https://git-scm.com/downloads)

### Steps
1. **Clone the repository**:
   ```bash
   git clone https://github.com/Rombond/budgestify.git
   cd budgestify
   ```

2. **Start the application**:
   ```bash
   docker compose up
   ```
   This sets up the Go backend, MySQL database, and (eventually) the Svelte frontend.

3. **Access the services**:
   - Backend: `http://localhost:8080`
   - Frontend: Not yet available (check `frontend` directory for setup once implemented).

4. **Stop the application**:
   ```bash
   docker compose down
   ```

> **Note**: All dependencies are managed by Docker, so no local installation of Go, MySQL, or Node.js is required.

## Contribution Process

1. **Fork the repository** to your GitHub account.
2. **Create a branch** for your changes:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. **Make changes** following the coding standards below.
4. **Test locally**:
   - Ensure the backend (`http://localhost:8080`) and database work as expected.
   - For frontend contributions (once available), verify the UI renders correctly.
5. **Commit changes** with clear messages:
   ```bash
   git commit -m "Add expense categorization endpoint to Go backend"
   ```
6. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```
7. **Open a Pull Request (PR)**:
   - Use a clear title and description.
   - Reference any related issues (e.g., `Fixes #123`).
   - Include screenshots for UI changes (e.g., shadcn components).

## Coding Standards

- **General**:
  - Write clear, concise code with meaningful comments.
  - Ensure backward compatibility for database migrations.

## Reporting Bugs

- Check existing issues to avoid duplicates.
- Create a new issue with:
  - A clear title and description.
  - Steps to reproduce the bug.
  - Environment details (e.g., Docker version, OS).
  - Screenshots or logs, if applicable.

## Suggesting Features

- Open an issue with the `enhancement` label.
- Describe the feature, its benefits, and potential implementation ideas.
- For example: "Add a pie chart for expense categories using shadcn components."

## Communication

- Use GitHub Issues and Discussions for questions, ideas, or feedback.
- Be respectful and inclusive in all interactions.

## License

By contributing, you agree that your contributions will be licensed under the [GPL-3.0](LICENSE).

Thank you for helping make Budgestify better!


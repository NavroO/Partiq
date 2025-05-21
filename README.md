# 🗳️ Partiq

**Partiq** is a modern, lightweight, and developer-friendly platform for participatory democracy.  
Built with **Go** and **React**, it lets you easily run civic processes such as proposals, voting, and collaborative decision-making — without the legacy complexity.

---

## 📦 Monorepo structure

```
partiq/
├── apps/
│   ├── api/         # Backend – Go + PostgreSQL
│   └── client/      # Frontend – React + Vite + Tailwind
│
├── packages/
│   └── shared/      # Shared types, API contracts, i18n, constants
│
├── docker-compose.yml
├── Makefile
├── .env.example
└── README.md
```

---

## 🚀 Quick Start

### 1. Clone the repo
```bash
git clone https://github.com/yourname/partiq.git
cd partic
```

### 2. Setup environment
```bash
cp .env.example .env
```

### 3. Start the stack
```bash
docker-compose up --build
```

---

## 🌍 Features

- 🔐 Authentication system (JWT)
- 🗳️ Proposals and voting per process
- 📋 Admin panel for managing processes
- 🌐 Fully API-driven architecture
- 🌱 Seed data for fast prototyping
- ⚡ Hot reload (Go + React)
- 📦 Modern DX-first monorepo layout

---

## 🧪 Development commands

```bash
make dev      # Start backend + frontend with live reload
make seed     # Seed the database with example data
make db       # Recreate local database
```

---

## 📄 License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.

# 🫐 BlueBerry Task Scheduler

**BlueBerry** is a robust, feature-rich task scheduler written in **Go**. It provides both a powerful scheduler engine and an integrated platform for management, complete with a modern **Web GUI** and a fully documented **RESTful API**.

## ✨ Key Features

* **Integrated Web GUI:** Manage tasks, schedules, and view execution logs via a user-friendly web interface (supports light/dark mode).
* **Flexible Scheduling:** Supports standard cron expressions, predefined intervals (e.g., `@every 5m`), and graceful shutdown.
* **Dynamic Task Management:** Tasks are Go functions registered with a schema, allowing execution via API or GUI with dynamic parameters.
* **Multi-DB Support:** Store task runs and logs in **SQLite**, **PostgreSQL**, or **MongoDB**.
* **Authentication:** Native support for multi-user authentication via cookies (Web GUI) and API keys (REST API).

---

## 📚 Full Documentation

For comprehensive guides, installation instructions, API reference, and a full feature breakdown, please visit our official documentation site:

➡️ **[BlueBerry Task Scheduler Documentation](https://blueberry-go.github.io/Scheduler/)** ⬅️

---

## 🚀 Quick Start (Installation)

To get BlueBerry running immediately, simply fetch the necessary Go modules:

```bash
# Get the core scheduler engine
go get github.com/blueberry-go/scheduler/core

# Get the SQLite storage module (recommended for quick start)
go get github.com/blueberry-go/scheduler/store/sqlite
````

The **[Getting Started guide](https://blueberry-go.github.io/Scheduler/getting-started/)** in the docs provides a complete, runnable `main.go` example.

## 🤝 Contributing
We welcome contributions\! Please see our documentation on **[Contributing](https://blueberry-go.github.io/Scheduler/)** (link will point to the relevant doc section once created) for more details.
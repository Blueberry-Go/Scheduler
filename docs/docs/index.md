# Welcome to BlueBerry Task Scheduler

![](assets/logo/logo_trans.png)

BlueBerry is a task scheduler written in Go, complete with a web GUI and a RESTful API, designed to make scheduling and managing tasks easy and efficient.

## Features

- **Web GUI** for task management (with light and dark modes).
- **Native Authentication** support (multi-user) via cookie (for web) and API key (for API).
- **Multi-DB Support**:
  - SQLite
  - PostgreSQL
  - MongoDB
- **RESTful API** for integrating task management into other applications.
- **Flexible Scheduling** with support for common cron intervals and custom cron expressions.
- **Graceful Shutdown** handling.
- **Logging** for task execution and status.
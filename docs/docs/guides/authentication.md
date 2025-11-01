# How to Set Up Authentication

BlueBerry provides two independent authentication methods: **Web GUI Auth** (cookie-based) and **API Auth** (API key-based).

If you do not add any users, the GUI and API will be open to everyone.

## 1. Web GUI Authentication

This method uses a username and password, which sets a browser cookie for session management.

```go
// rb is your blueberry.NewBlueBerryInstance(db)
rb.AddWebOnlyPasswordAuth("admin", "password")

// You can add multiple web users
rb.AddWebOnlyPasswordAuth("admin1", "password123")
```

## 2\. API Key Authentication

This method uses an API key (bearer token) for programmatic access to the REST API.

```go
// rb is your blueberry.NewBlueBerryInstance(db)
rb.AddAPIOnlyKeyAuth("your-api-key", "Main API Key")

// You can add multiple API keys
rb.AddAPIOnlyKeyAuth("your-api-key-1", "2nd API Key")
rb.AddAPIOnlyKeyAuth("service-account-key", "Service Account")
```

## 3\. Using Both Methods

Since the methods are independent, you typically need to add credentials for both if you want to secure both the GUI and the API.

```go
// Secure the Web GUI
rb.AddWebOnlyPasswordAuth("admin", "password")

// Secure the API
rb.AddAPIOnlyKeyAuth("my-secret-key", "Admin Key")
```
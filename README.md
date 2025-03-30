# Dynamic Configuration Web Application

## Brief Description

This web application allows dynamic configuration of static HTML files using a Go backend and a JavaScript frontend. The backend provides a RESTful API for managing configurations, which can be applied to specific pages, URLs, or hosts. The frontend fetches these configurations and applies them to the DOM, allowing for actions like inserting, replacing, altering, or removing elements dynamically.

## Features

- **Dynamic Configuration Handling**: Manage configuration rules for pages, URLs, and hosts.
- **Action Types**: Supports actions such as `insert`, `replace`, `remove`, and `alter` to manipulate the DOM.
- **Static File Serving**: Serves static HTML, CSS, and JavaScript files, applying configurations dynamically.
- **Multiple Configuration Handling**: Handle multiple configurations with priority logic to resolve conflicts.
- **RESTful API**: Exposes API endpoints to interact with configurations.

## Setup

### Prerequisites

Ensure you have the following installed:

- **Go** (1.18+)
- **Git** (for version control)

### Cloning the Repository

Clone the repository and navigate to the project folder:

git clone https://github.com/turanmehmetb/dynamicWeb.git

cd dynamicWeb

### Backend Setup

1. **Install Go dependencies**:

go mod tidy

2. **Run the Go server**:

go run main.go

This will start the server at `http://127.0.0.1:8080`.

### Testing the Application

To test the application, visit the following pages in your browser:

- **Home Page**: `http://127.0.0.1:8080/`
- **Orders Page**: `http://127.0.0.1:8080/orders`

## Endpoints

### **Configuration Endpoints**

- **POST `/api/configuration/`**
  - **Description**: Creates a new configuration.
  - **Request Body**:
```json  
  {
    "actions": [
      {
        "type": "insert",
        "position": "after",
        "target": "body",
        "element": "<footer>Footer Content</footer>"
      },
      {
        "type": "alter",
        "oldValue": "Machine Learning",
        "newValue": "AI"
      }
    ]
  }
```
  - **Response**:
```json
{
  "id": "23bc1248-3f03-4990-856c-de1cce83dbdd",
  "actions": [
    {
      "type": "insert",
      "element": "<footer>Footer Content</footer>",
      "position": "after",
      "target": "body"
    },
    {
      "type": "alter",
      "oldValue": "Machine Learning",
      "newValue": "AI"
    }
  ]
}
```
- **GET `/api/configuration/:id`**
  - **Description**: Retrieves the configuration for the specified `id`.

- **GET `/api/configuration/all`**
  - **Description**: Retrieves all configurations.

- **PUT `/api/configuration/:id`**
  - **Description**: Updates the configuration for the given `id`.
  - **Request Body**:
```json  
  {
    "actions": [
      {
        "type": "remove",
        "selector": ".ad-banner"
      }
    ]
  }
```
- **DELETE `/api/configuration/:id`**
  - **Description**: Deletes the configuration for the specified `id`.

---

### **Specific Configuration Endpoints**

- **POST `/api/specific/`**
  - **Description**: Creates a new specific configuration.
  - **Request Body**:
```json
{
  "datasource": {
    "pages": {
      "cart": ["4d508895-70d9-4006-a850-ff14595b6538.yaml"],
      "details": ["123e4567-e89b-12d3-a456-426614174002.yaml"]
    }
  }
}
```
  - **Response**:
```json
{
  "id": "bccabbc4-1497-4c7c-aa23-05cd2ad82c37",
  "datasource": {
    "pages": {
      "cart": [
        "4d508895-70d9-4006-a850-ff14595b6538.yaml"
      ],
      "details": [
        "123e4567-e89b-12d3-a456-426614174002.yaml"
      ]
    },
    "urls": null,
    "hosts": null
  }
}
```

- **GET `/api/specific/:id`**
  - **Description**: Retrieves a specific configuration for the given `id`. Can include query parameters like `host`, `url`, or `page`.

- **GET `/api/specific/all`**
  - **Description**: Retrieves all specific configurations.

- **PUT `/api/specific/:id`**
  - **Description**: Updates the specific configuration for the given `id`.

- **DELETE `/api/specific/:id`**
  - **Description**: Deletes the specific configuration for the given `id`.

## Conclusion

This project provides a flexible solution for dynamically configuring HTML pages using a backend-driven approach. The frontend fetches and applies configurations using actions like insertion, replacement, and alteration. The backend supports full CRUD operations for configurations, allowing for powerful dynamic content management.


---

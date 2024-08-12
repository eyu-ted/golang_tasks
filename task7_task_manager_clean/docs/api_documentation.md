# Task Management API Documentation

## MongoDB Integration

Ensure MongoDB is running and accessible. Update the connection string in `data/mongo.go` if necessary.

## Endpoints

### GET /tasks
Retrieve a list of all tasks.

**Response:**
```json
[
  {
    "id": "60d5ec49edff204b1c67b0b1",
    "title": "Sample Task",
    "description": "This is a sample task",
    "due_date": "2024-08-01T12:00:00Z",
    "status": "pending"
  }
]

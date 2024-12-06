# Medication API

This repository contains a RESTful API for managing medication records. Built with Golang, it provides full CRUD functionality and adheres to best practices in RESTful architecture.

## Features
- **Create, Read, Update, Delete** operations for managing medications.
- Lightweight and scalable design.
- Easy-to-use endpoints for managing medication data.

## Medication Object
A medication record includes:
- **ID**: A unique identifier for the medication.
- **Name**: Name of the medication (e.g., "Paracetamol").
- **Dosage**: Prescribed dosage amount (e.g., "500" in mg).
- **Form**: Form of the medication (e.g., "Tablet", "Capsule").

## Prerequisites
- Docker
- Docker Compose
- Make

## Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/aborilov/hippo
   cd hippo
   ```
2. Build the Docker image and start the development environment:
   ```bash
   make dev-up
   ```
3. Stop the development environment:
   ```bash
   make dev-down
   ```

## API Endpoints

### Get All Medications
```bash
curl -X GET http://localhost:6000/medication/
```

### Get Medication by ID
```bash
curl -X GET http://localhost:6000/medication/<id>
```

### Create a New Medication
```bash
curl -X POST http://localhost:6000/medication/ \
-H "Content-Type: application/json" \
-d '{"name": "red pill", "dosage": 1, "form": "tablet"}'
```

### Update a Medication
```bash
curl -X PUT http://localhost:6000/medication/<id> \
-H "Content-Type: application/json" \
-d '{"name": "blue pill", "dosage": 1, "form": "tablet"}'
```

### Delete a Medication
```bash
curl -X DELETE http://localhost:6000/medication/<id>
```

## Example Usage
1. **Get all medications**:
   ```bash
   curl -X GET http://localhost:6000/medication/
   ```
   Response:
   ```json
   [
       {
           "id": "5cf37266-3473-4006-984f-9325122678b7",
           "name": "magic pill",
           "dosage": 1,
           "form": "tablet"
       }
   ]
   ```

2. **Create a new medication**:
   ```bash
   curl -X POST http://localhost:6000/medication/ \
   -H "Content-Type: application/json" \
   -d '{"name": "red pill", "dosage": 1, "form": "tablet"}'
   ```
   Response:
   ```json
   {
       "id": "8d020735-eaa5-4e0b-86d3-1de8188b615c",
       "name": "red pill",
       "dosage": 1,
       "form": "tablet"
   }
   ```

3. **Update an existing medication**:
   ```bash
   curl -X PUT http://localhost:6000/medication/8d020735-eaa5-4e0b-86d3-1de8188b615c \
   -H "Content-Type: application/json" \
   -d '{"name": "blue pill", "dosage": 1, "form": "tablet"}'
   ```
   Response:
   ```json
   {
       "id": "8d020735-eaa5-4e0b-86d3-1de8188b615c",
       "name": "blue pill",
       "dosage": 1,
       "form": "tablet"
   }
   ```

4. **Delete a medication**:
   ```bash
   curl -X DELETE http://localhost:6000/medication/8d020735-eaa5-4e0b-86d3-1de8188b615c
   ```
   Response:
   ```json
   {}
   ```

## Cleanup
Stop and remove the containers and volumes:
```bash
make dev-down
```

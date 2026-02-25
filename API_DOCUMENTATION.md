# Backend API Documentation

## Overview
This is the complete API documentation for the Palap backend system, including User, Admin, and Freelancer modules.

## Base URL
```
http://localhost:8080/api/v1
```

---

## 1. USER ENDPOINTS

### 1.1 Process Payment
**POST** `/payment`

#### Headers
```
Authorization: <user_token>
Content-Type: application/json
```

#### Request Body (Card Payment)
```json
{
  "service_id": 1,
  "address_id": 1,
  "payment_method": "card",
  "payment_card": {
    "card_number": "4111111111111111",
    "card_holder": "John Doe",
    "expiry_date": "12/25",
    "cvv": "123"
  }
}
```

#### Request Body (QR Code Payment)
```json
{
  "service_id": 1,
  "address_id": 1,
  "payment_method": "qr_code",
  "payment_card": null
}
```

#### Response (Card Payment - Success)
```json
{
  "status": "OK"
}
```

#### Response (QR Code Payment - Success)
```json
{
  "status": "OK",
  "qr_code": "https://api.qr-server.com/v1/qr?size=200x200&data=order_..."
}
```

#### Response (Error)
```json
{
  "error": "Invalid payment method"
}
```

---

### 1.2 Get Orders
**GET** `/get_orders`

#### Headers
```
Authorization: <user_token>
```

#### Response
```json
[
  {
    "id": 1,
    "user_id": 1,
    "service_id": 1,
    "address_id": 1,
    "status": "Completed",
    "payment_method": "card",
    "qr_code": null,
    "created_at": "2026-02-25"
  },
  {
    "id": 2,
    "user_id": 1,
    "service_id": 2,
    "address_id": 2,
    "status": "Pending",
    "payment_method": "qr_code",
    "qr_code": "https://example.com/qr/2",
    "created_at": "2026-02-25"
  }
]
```

---

### 1.3 Get Order Details
**GET** `/get_order_details?order_id=1`

#### Query Parameters
- `order_id` (required): Order ID

#### Response
```json
[
  {
    "id": 1,
    "order_id": 1,
    "description": "Service Description 1",
    "amount": 1000
  },
  {
    "id": 2,
    "order_id": 1,
    "description": "Service Description 2",
    "amount": 500
  }
]
```

---

## 2. ADMIN ENDPOINTS

### 2.1 Admin Management (CRUD)

#### 2.1.1 Get All Admins
**GET** `/manage_admin`

**Response**
```json
[
  {
    "id": 1,
    "username": "admin1",
    "email": "admin1@example.com",
    "role": "Super Admin",
    "status": "active"
  }
]
```

---

#### 2.1.2 Get Admin by ID
**GET** `/manage_admin/:id`

**Response**
```json
{
  "id": 1,
  "username": "admin1",
  "email": "admin1@example.com",
  "role": "Super Admin",
  "status": "active"
}
```

---

#### 2.1.3 Create Admin
**POST** `/manage_admin`

**Request Body**
```json
{
  "username": "admin3",
  "email": "admin3@example.com",
  "role": "Admin",
  "status": "active"
}
```

**Response**
```json
{
  "id": 3,
  "username": "admin3",
  "email": "admin3@example.com",
  "role": "Admin",
  "status": "active"
}
```

---

#### 2.1.4 Update Admin
**PUT** `/manage_admin/:id`

**Request Body**
```json
{
  "username": "admin1_updated",
  "email": "admin1_updated@example.com",
  "role": "Super Admin",
  "status": "inactive"
}
```

**Response**
```json
{
  "id": 1,
  "username": "admin1_updated",
  "email": "admin1_updated@example.com",
  "role": "Super Admin",
  "status": "inactive"
}
```

---

#### 2.1.5 Delete Admin
**DELETE** `/manage_admin/:id`

**Response** (No Content)
```
HTTP 204
```

---

### 2.2 Freelancer Management (CRUD)

#### 2.2.1 Get All Freelancers
**GET** `/manage_freelance`

**Response**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "bio": "Expert in web development",
    "skills": "Go, React, Node.js",
    "rating": 5,
    "status": "active",
    "created_at": "2026-01-15"
  }
]
```

---

#### 2.2.2 Get Freelancer by ID
**GET** `/manage_freelance/:id`

**Response**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "bio": "Expert in web development",
  "skills": "Go, React, Node.js",
  "rating": 5,
  "status": "active",
  "created_at": "2026-01-15"
}
```

---

#### 2.2.3 Create Freelancer
**POST** `/manage_freelance`

**Request Body**
```json
{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "bio": "UI/UX Designer",
  "skills": "Figma, Adobe XD",
  "rating": 4,
  "status": "active"
}
```

**Response**
```json
{
  "id": 3,
  "name": "Jane Doe",
  "email": "jane@example.com",
  "bio": "UI/UX Designer",
  "skills": "Figma, Adobe XD",
  "rating": 4,
  "status": "active",
  "created_at": "2026-02-25"
}
```

---

#### 2.2.4 Update Freelancer
**PUT** `/manage_freelance/:id`

**Request Body**
```json
{
  "name": "John Doe Updated",
  "email": "john.updated@example.com",
  "bio": "Senior web developer",
  "skills": "Go, React, Node.js, Python",
  "rating": 5,
  "status": "active"
}
```

**Response**
```json
{
  "id": 1,
  "name": "John Doe Updated",
  "email": "john.updated@example.com",
  "bio": "Senior web developer",
  "skills": "Go, React, Node.js, Python",
  "rating": 5,
  "status": "active",
  "created_at": "2026-01-15"
}
```

---

#### 2.2.5 Delete Freelancer
**DELETE** `/manage_freelance/:id`

**Response** (No Content)
```
HTTP 204
```

---

### 2.3 Product and Service Management (CRUD)

#### 2.3.1 Get All Products/Services
**GET** `/manage_product_and_service`

**Response**
```json
[
  {
    "id": 1,
    "name": "Web Development",
    "description": "Custom web development services",
    "price": 10000,
    "category": "Software",
    "status": "active"
  }
]
```

---

#### 2.3.2 Get Product/Service by ID
**GET** `/manage_product_and_service/:id`

**Response**
```json
{
  "id": 1,
  "name": "Web Development",
  "description": "Custom web development services",
  "price": 10000,
  "category": "Software",
  "status": "active"
}
```

---

#### 2.3.3 Create Product/Service
**POST** `/manage_product_and_service`

**Request Body**
```json
{
  "name": "Graphic Design",
  "description": "Professional graphic design services",
  "price": 5000,
  "category": "Design",
  "status": "active"
}
```

**Response**
```json
{
  "id": 3,
  "name": "Graphic Design",
  "description": "Professional graphic design services",
  "price": 5000,
  "category": "Design",
  "status": "active"
}
```

---

#### 2.3.4 Update Product/Service
**PUT** `/manage_product_and_service/:id`

**Request Body**
```json
{
  "name": "Web Development Pro",
  "description": "Advanced custom web development services",
  "price": 15000,
  "category": "Software",
  "status": "active"
}
```

**Response**
```json
{
  "id": 1,
  "name": "Web Development Pro",
  "description": "Advanced custom web development services",
  "price": 15000,
  "category": "Software",
  "status": "active"
}
```

---

#### 2.3.5 Delete Product/Service
**DELETE** `/manage_product_and_service/:id`

**Response** (No Content)
```
HTTP 204
```

---

### 2.4 Shop Location Management (CRUD)

#### 2.4.1 Get All Shop Locations
**GET** `/manage_shop_location`

**Response**
```json
[
  {
    "id": 1,
    "name": "Shop 1",
    "address": "123 Main St",
    "phone": "08x-xxx-xxxx",
    "city": "Bangkok",
    "country": "Thailand",
    "status": "active"
  }
]
```

---

#### 2.4.2 Get Shop Location by ID
**GET** `/manage_shop_location/:id`

**Response**
```json
{
  "id": 1,
  "name": "Shop 1",
  "address": "123 Main St",
  "phone": "08x-xxx-xxxx",
  "city": "Bangkok",
  "country": "Thailand",
  "status": "active"
}
```

---

#### 2.4.3 Create Shop Location
**POST** `/manage_shop_location`

**Request Body**
```json
{
  "name": "Shop 3",
  "address": "789 Rama St",
  "phone": "08x-xxx-xxxx",
  "city": "Chiang Mai",
  "country": "Thailand",
  "status": "active"
}
```

**Response**
```json
{
  "id": 3,
  "name": "Shop 3",
  "address": "789 Rama St",
  "phone": "08x-xxx-xxxx",
  "city": "Chiang Mai",
  "country": "Thailand",
  "status": "active"
}
```

---

#### 2.4.4 Update Shop Location
**PUT** `/manage_shop_location/:id`

**Request Body**
```json
{
  "name": "Shop 1 Updated",
  "address": "456 Sukhumvit Rd",
  "phone": "08x-yyy-yyyy",
  "city": "Bangkok",
  "country": "Thailand",
  "status": "active"
}
```

**Response**
```json
{
  "id": 1,
  "name": "Shop 1 Updated",
  "address": "456 Sukhumvit Rd",
  "phone": "08x-yyy-yyyy",
  "city": "Bangkok",
  "country": "Thailand",
  "status": "active"
}
```

---

#### 2.4.5 Delete Shop Location
**DELETE** `/manage_shop_location/:id`

**Response** (No Content)
```
HTTP 204
```

---

## 3. FREELANCER ENDPOINTS

### 3.1 Freelancer Registration
**POST** `/registration`

#### Request Body
```json
{
  "name": "Alice Johnson",
  "email": "alice@example.com",
  "bio": "Frontend Developer",
  "skills": "React, Vue, Angular",
  "rating": 0,
  "status": "active"
}
```

#### Response
```json
{
  "id": 3,
  "name": "Alice Johnson",
  "email": "alice@example.com",
  "bio": "Frontend Developer",
  "skills": "React, Vue, Angular",
  "rating": 0,
  "status": "active",
  "created_at": "2026-02-25"
}
```

---

### 3.2 Get Freelance Jobs
**GET** `/get_freelance_job`

#### Response
```json
[
  {
    "id": 1,
    "title": "Build mobile app",
    "description": "Need to build an iOS app",
    "budget": 50000,
    "status": "open",
    "created_at": "2026-02-20"
  },
  {
    "id": 2,
    "title": "Web design",
    "description": "Design new website",
    "budget": 30000,
    "status": "open",
    "created_at": "2026-02-22"
  }
]
```

---

### 3.3 Get Order Details (Same as User)
**GET** `/get_order_details?order_id=1`

#### Query Parameters
- `order_id` (required): Order ID

#### Response
```json
[
  {
    "id": 1,
    "order_id": 1,
    "description": "Service Description 1",
    "amount": 1000
  }
]
```

---

## Error Responses

All error responses follow this format:

```json
{
  "error": "Error message describing what went wrong"
}
```

### Common Error Cases

| Status Code | Error Message | Cause |
|-----------|---------------|-------|
| 400 | Invalid request body | Malformed JSON or missing required fields |
| 400 | Invalid admin ID | The provided ID cannot be parsed as an integer |
| 400 | Missing authorization token | Authorization header not provided |
| 401 | Missing authorization token | Not authorized to access the resource |
| 404 | Admin not found | Admin with the given ID doesn't exist |
| 404 | Freelancer not found | Freelancer with the given ID doesn't exist |
| 404 | Product/Service not found | Product/Service with the given ID doesn't exist |
| 404 | Shop location not found | Shop location with the given ID doesn't exist |
| 404 | Order not found | Order with the given ID doesn't exist |

---

## Summary of All Endpoints

### User Module (3 endpoints)
- POST `/payment` - Process payment
- GET `/get_orders` - Get user orders
- GET `/get_order_details` - Get order details

### Admin Module (20 endpoints)
- Admin Management (5): GET, GET/:id, POST, PUT/:id, DELETE/:id
- Freelancer Management (5): GET, GET/:id, POST, PUT/:id, DELETE/:id
- Product/Service Management (5): GET, GET/:id, POST, PUT/:id, DELETE/:id
- Shop Location Management (5): GET, GET/:id, POST, PUT/:id, DELETE/:id

### Freelancer Module (3 endpoints)
- POST `/registration` - Register as freelancer
- GET `/get_freelance_job` - Get available jobs
- GET `/get_order_details` - Get order details (shared)

**Total: 26 endpoints**

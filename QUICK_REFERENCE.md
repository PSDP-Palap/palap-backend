# API Quick Reference

## Base URL
```
http://localhost:8080/api/v1
```

## Quick Examples

### USER ENDPOINTS

**1. Payment (Card)**
```bash
POST /payment
Header: Authorization: user_token_here
Body: {
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
Response: { "status": "OK" }
```

**2. Payment (QR Code)**
```bash
POST /payment
Header: Authorization: user_token_here
Body: {
  "service_id": 1,
  "address_id": 1,
  "payment_method": "qr_code",
  "payment_card": null
}
Response: { "status": "OK", "qr_code": "https://..." }
```

**3. Get Orders**
```bash
GET /get_orders
Header: Authorization: user_token_here
Response: [{ id, user_id, service_id, ... }]
```

**4. Get Order Details**
```bash
GET /get_order_details?order_id=1
Response: [{ id, order_id, description, amount }]
```

---

## ADMIN MANAGEMENT

**List All**
```bash
GET /manage_admin
```

**Get One**
```bash
GET /manage_admin/1
```

**Create**
```bash
POST /manage_admin
Body: { "username": "admin3", "email": "admin3@example.com", "role": "Admin", "status": "active" }
```

**Update**
```bash
PUT /manage_admin/1
Body: { "username": "admin1_updated", "email": "admin1_updated@example.com", "role": "Super Admin", "status": "inactive" }
```

**Delete**
```bash
DELETE /manage_admin/1
```

---

## FREELANCER MANAGEMENT

**List All**
```bash
GET /manage_freelance
```

**Get One**
```bash
GET /manage_freelance/1
```

**Create**
```bash
POST /manage_freelance
Body: { "name": "Jane Doe", "email": "jane@example.com", "bio": "UI/UX Designer", "skills": "Figma, Adobe XD", "rating": 4, "status": "active" }
```

**Update**
```bash
PUT /manage_freelance/1
Body: { "name": "John Updated", ... }
```

**Delete**
```bash
DELETE /manage_freelance/1
```

---

## FREELANCER ENDPOINTS

**Register**
```bash
POST /registration
Body: { "name": "Alice", "email": "alice@example.com", "bio": "Developer", "skills": "Go, React", "rating": 0, "status": "active" }
```

**Get Jobs**
```bash
GET /get_freelance_job
```

**Get Order Details** (shared with user)
```bash
GET /get_order_details?order_id=1
```

---

## PRODUCTS & SERVICES

**List All**
```bash
GET /manage_product_and_service
```

**Get One**
```bash
GET /manage_product_and_service/1
```

**Create**
```bash
POST /manage_product_and_service
Body: { "name": "Graphic Design", "description": "...", "price": 5000, "category": "Design", "status": "active" }
```

**Update**
```bash
PUT /manage_product_and_service/1
Body: { "name": "Graphic Design Pro", ... }
```

**Delete**
```bash
DELETE /manage_product_and_service/1
```

---

## SHOP LOCATIONS

**List All**
```bash
GET /manage_shop_location
```

**Get One**
```bash
GET /manage_shop_location/1
```

**Create**
```bash
POST /manage_shop_location
Body: { "name": "Shop 3", "address": "789 Rama St", "phone": "08x-xxx-xxxx", "city": "Chiang Mai", "country": "Thailand", "status": "active" }
```

**Update**
```bash
PUT /manage_shop_location/1
Body: { "name": "Shop 1 Updated", ... }
```

**Delete**
```bash
DELETE /manage_shop_location/1
```

---

## HTTP Status Codes
- `200` OK - Success (GET/PUT)
- `201` Created - Resource created (POST)
- `204` No Content - Success delete (DELETE)
- `400` Bad Request - Invalid data
- `401` Unauthorized - Missing token
- `404` Not Found - Resource doesn't exist

---

## Sample Data Available

### Admins
- ID 1: admin1@example.com (Super Admin)
- ID 2: admin2@example.com (Admin)

### Freelancers
- ID 1: John Doe (Web Developer, Rating: 5)
- ID 2: Jane Smith (UI/UX Designer, Rating: 4)

### Freelance Jobs
- ID 1: Build mobile app (Budget: 50000)
- ID 2: Web design (Budget: 30000)

### Products/Services
- ID 1: Web Development (Price: 10000)
- ID 2: Mobile App Development (Price: 50000)

### Shop Locations
- ID 1: Shop 1 (Bangkok)
- ID 2: Shop 2 (Bangkok)

### Orders
- ID 1: User 1, Service 1, Status: Completed, Payment: card
- ID 2: User 1, Service 2, Status: Pending, Payment: qr_code

---

## Common Patterns

All CRUD endpoints follow REST conventions:
| Operation | Method | Endpoint | Status |
|-----------|--------|----------|--------|
| List | GET | `/resource` | 200 |
| Read | GET | `/resource/:id` | 200 |
| Create | POST | `/resource` | 201 |
| Update | PUT | `/resource/:id` | 200 |
| Delete | DELETE | `/resource/:id` | 204 |

All responses are JSON formatted.
All errors include an `"error"` field with description.

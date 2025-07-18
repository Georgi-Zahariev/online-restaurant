# Simple Role-Based Access Control

## Overview
This is a simplified role-based access control system without authentication. It uses HTTP headers to specify user roles temporarily until your IDP (Identity Provider) is integrated later.

## Roles Defined

### Customer (`customer`)
- **Access**: Home, Menu, Shopping cart, Profile, Orders
- **Endpoints**:
  - `GET /api/home`
  - `GET /api/menu` 
  - `GET /api/dishes`
  - `GET|POST /api/cart`
  - `GET|POST /api/orders`
  - `GET|PUT /api/profile`

### Kitchen (`kitchen`)
- **Access**: Profile, Order dashboard
- **Endpoints**:
  - `GET /api/order-dashboard`
  - `PUT /api/order-items/complete`
  - `GET|PUT /api/profile`

### Delivery (`delivery`)
- **Access**: Profile, Delivery dashboard
- **Endpoints**:
  - `GET /api/delivery-dashboard`
  - `PUT /api/orders/deliver`
  - `GET|PUT /api/profile`

### Owner (`owner`)
- **Access**: Everything (full access)
- **Endpoints**: All endpoints above plus:
  - `GET /api/admin`
  - `GET /api/analytics`

## How It Works

1. **Header-Based Role**: Send role via `X-User-Role` header
2. **Middleware Check**: `RequireRole()` middleware validates access
3. **Owner Override**: Owner role can access any endpoint
4. **Access Denied**: Returns 403 for insufficient permissions

## Example Usage

```bash
# Customer accessing menu
curl -H "X-User-Role: customer" http://localhost:8080/api/menu

# Kitchen accessing order dashboard  
curl -H "X-User-Role: kitchen" http://localhost:8080/api/order-dashboard

# Owner accessing admin panel
curl -H "X-User-Role: owner" http://localhost:8080/api/admin

# Missing role (will fail)
curl http://localhost:8080/api/profile
```

## Database Changes

Added `role` column to User table:
```sql
ALTER TABLE "User" ADD COLUMN role VARCHAR(20) DEFAULT 'customer' CHECK (role IN ('customer', 'kitchen', 'delivery', 'owner'));
```

## Files Modified

- `backend/models/structures.go` - Added UserRole type and constants
- `backend/middlewares/middlewares.go` - Added RequireRole middleware
- `backend/handlers/role_handlers.go` - Created role-aware handlers
- `backend/routers/router.go` - Set up role-protected routes
- `db/migrations/20250718000000_add_user_roles.up.sql` - Database migration

## Testing

Run the test script to verify role access:
```bash
chmod +x test_roles.sh
./test_roles.sh
```

## Next Steps

When you're ready to add authentication:
1. Integrate with your chosen IDP
2. Replace `X-User-Role` header with JWT token parsing
3. Extract user role from token claims
4. The existing role middleware will work unchanged

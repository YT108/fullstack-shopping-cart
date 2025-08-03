# Fullstack Shopping Cart

This is a simple full-stack shopping cart application built using:
- Backend: Golang (Gin + GORM + SQLite)
- Frontend: React.js
- Styling: Basic CSS
- Authentication: Simple login system with token (placeholder)

---

## 🚀 Features

- User login
- List available items
- Add to cart
- View cart
- Checkout (clear cart)
- View past orders

---

## 📦 Project Structure


---

## ✅ 3. Create Postman Collection

1. Open [Postman](https://www.postman.com/)
2. Create a new collection called **Shopping Cart API**
3. Add requests for:
   - `POST /users`
   - `POST /users/login`
   - `GET /items`
   - `POST /items`
   - `POST /carts`
   - `GET /carts`
   - `DELETE /carts/user/:user_id`
   - `POST /orders`
   - `GET /orders`
4. Save it
5. Click the 3 dots → **Export Collection** → Save as `postman_collection.json`
6. Move that file into the project root

---

## ✅ 4. Final Git Commands

```bash
git add .
git commit -m "Added README and Postman collection"
git push


## 🛠 How to Run the Project

### Backend (Golang + GORM + SQLite)

1. Navigate to the backend folder:
   ```bash
   cd backend

go run main.go



### Frontend(react.js + vite)
1.  cd frontend
    npm install
    npm start

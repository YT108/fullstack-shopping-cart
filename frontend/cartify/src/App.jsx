import React, { useState } from "react";
import axios from "axios";
import "./App.css";

function App() {
  const [token, setToken] = useState("");
  const [userId, setUserId] = useState(null);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [items, setItems] = useState([]);

  const handleLogin = async () => {
    try {
      const res = await axios.post("http://localhost:8080/users/login", {
        username,
        password,
      });
      setToken(res.data.token);
      setUserId(res.data.user_id);
      fetchItems();
    } catch (err) {
      alert("Invalid username/password");
    }
  };

  const fetchItems = async () => {
    try {
      const res = await axios.get("http://localhost:8080/items");
      const uniqueItems = Array.from(
        new Map(res.data.map((item) => [item.ID, item])).values()
      );
      setItems(uniqueItems);
    } catch (err) {
      console.error("Error fetching items:", err);
    }
  };

  const addToCart = async (itemId) => {
    try {
      await axios.post("http://localhost:8080/carts", {
        user_id: userId,
        item_id: itemId,
      });
      alert("Item added to cart");
    } catch (err) {
      alert("Failed to add to cart");
    }
  };

  const viewCart = async () => {
    try {
      const res = await axios.get(`http://localhost:8080/carts?user_id=${userId}`);
      const cart = res.data;

      if (!cart.Items || cart.Items.length === 0)
        return alert("Your cart is empty.");

      const itemList = cart.Items
        .map((i) => `• ${i.Item.Name} - ₹${i.Item.Price}`)
        .join("\n");

      alert(`Your Cart:\n${itemList}`);
    } catch (err) {
      alert("Failed to load cart");
    }
  };

  const handleCheckout = async () => {
    try {
      await axios.post("http://localhost:8080/orders", { user_id: userId });
      await axios.delete(`http://localhost:8080/carts/user/${userId}`);
      alert("Order successful. Cart cleared.");
      fetchItems();
    } catch (err) {
      alert("Checkout failed");
    }
  };

  const viewOrders = async () => {
    try {
      const res = await axios.get("http://localhost:8080/orders");
      const userOrders = res.data.filter((o) => o.UserID === userId);

      if (userOrders.length === 0) return alert("No orders found.");

      const orderList = userOrders
        .map((o) => `Order ID: ${o.ID} (Cart ID: ${o.CartID})`)
        .join("\n");

      alert(`Your Orders:\n${orderList}`);
    } catch (err) {
      alert("Failed to load orders");
    }
  };

  const handleLogout = () => {
    setToken("");
    setUserId(null);
    setUsername("");
    setPassword("");
    setItems([]);
  };

  return (
    <div className={token ? "app-container" : "login-background"}>
      {!token ? (
        <div className="login-box">
          <h2>Welcome Back!</h2>
          <input
            className="input-field"
            placeholder="Username"
            onChange={(e) => setUsername(e.target.value)}
          />
          <input
            className="input-field"
            placeholder="Password"
            type="password"
            onChange={(e) => setPassword(e.target.value)}
          />
          <button onClick={handleLogin} className="btn btn-primary">
            Login
          </button>
        </div>
      ) : (
        <div className="dashboard">
          <div className="header">
            <h2>Hello, {username}</h2>
            <button onClick={handleLogout} className="btn btn-danger">
              Logout
            </button>
          </div>

          <div className="actions">
            <button onClick={viewCart} className="btn btn-secondary">View Cart</button>
            <button onClick={viewOrders} className="btn btn-secondary">Order History</button>
            <button onClick={handleCheckout} className="btn btn-success">Checkout</button>
          </div>

          <h3 className="section-title">Shop Items</h3>
          <div className="items-grid">
            {items.map((item) => (
              <div key={item.ID} className="item-card">
                <img
                  src={
                    item.Name.toLowerCase().includes("iphone")
                      ? "/images/iphone 15.jpeg"
                      : item.Name.toLowerCase().includes("airpod")
                      ? "/images/airpod pro.jpeg"
                      : "https://via.placeholder.com/200x150?text=No+Image"
                  }
                  alt={item.Name}
                  className="item-image"
                />
                <p className="item-name">{item.Name}</p>
                <p className="item-price">Price: ₹{item.Price}</p>
                <button onClick={() => addToCart(item.ID)} className="btn btn-cart">
                  Add to Cart
                </button>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}

export default App;

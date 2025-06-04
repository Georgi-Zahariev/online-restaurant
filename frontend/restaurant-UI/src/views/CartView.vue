<template>
  <main>
    <h1>Shopping cart</h1>
    <div class="cart-list">
      <div class="cart-item">
        <div class="cart-item-header">
          <span class="meal-name">Meal 1</span>
        </div>
        <div class="cart-item-details">
          <span>Count</span>
          <span>Price</span>
          <button class="cart-delete-btn">DELETE</button>
        </div>
      </div>
      <div class="cart-item">
        <div class="cart-item-header">
          <span class="meal-name">Meal 2</span>
        </div>
        <div class="cart-item-details">
          <span>Count</span>
          <span>Price</span>
          <button class="cart-delete-btn">DELETE</button>
        </div>
      </div>
    </div>
    <div class="cart-actions">
      <button class="cart-action-btn">UPDATE</button>
      <button class="cart-action-btn">Order</button>
    </div>
  </main>
</template>

<style>
main {
  max-width: 700px;
  margin: 0 auto;
}

.cart-list {
  margin-bottom: 32px;
}

.cart-item {
  border: 2px solid #222;
  border-radius: 8px;
  padding: 18px 20px;
  margin-bottom: 24px;
  background: #fff;
}

.cart-item-header {
  font-size: 1.2rem;
  font-weight: 600;
  margin-bottom: 12px;
  text-align: left;
}

.cart-item-details {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.cart-delete-btn {
  padding: 8px 18px;
  background: #e0f2e9;
  border: 2px solid #7bb37d;
  border-radius: 4px;
  color: #2f5f31;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s, border-color 0.2s;
}

.cart-delete-btn:hover {
  background: #c8e6c9;
  border-color: #388e3c;
}

.cart-actions {
  display: flex;
  justify-content: flex-end;
  gap: 24px;
}

.cart-action-btn {
  padding: 10px 28px;
  border: 2px solid #222;
  border-radius: 4px;
  background: #fff;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s, border-color 0.2s;
}

.cart-action-btn:hover {
  background: #e8f5e9;
  border-color: #4caf50;
}
</style>
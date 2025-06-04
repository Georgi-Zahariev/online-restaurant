<template>
  <main>
    <h1>My Orders</h1>
    <section>
      <h2>Current</h2>
      <div class="order-card">
        <div class="order-header">
          <span>Restaurant 1</span>
          <span>Price: $24.99</span>
        </div>
        <div>Status: <strong>EAT</strong></div>
        <button class="see-more-btn">See more</button>
      </div>
    </section>
    <section>
      <h2>Old</h2>
      <div class="order-card">
        <div class="order-header">
          <span>Restaurant 1</span>
          <span>Price: $18.50</span>
        </div>
        <div>Status: <strong>Delivered</strong></div>
        <div>Arrived at: 2025-06-03 19:30</div>
        <button class="see-more-btn">See more</button>
      </div>
    </section>
  </main>
</template>

<style>
main {
  max-width: 700px;
  margin: 0 auto;
}

.order-card {
  border: 2px solid #222;
  border-radius: 8px;
  padding: 18px 20px;
  margin-bottom: 24px;
  background: #fff;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.order-header {
  display: flex;
  justify-content: space-between;
  font-size: 1.1rem;
  font-weight: 600;
  margin-bottom: 8px;
}

.see-more-btn {
  align-self: flex-end;
  padding: 8px 18px;
  background: #e0f2e9;
  border: 2px solid #7bb37d;
  border-radius: 4px;
  color: #2f5f31;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s, border-color 0.2s;
}

.see-more-btn:hover {
  background: #c8e6c9;
  border-color: #388e3c;
}
</style>
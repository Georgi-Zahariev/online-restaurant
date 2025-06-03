<template>
  <ul class="nav-list">
    <li>
      <router-link to="/menu" class="nav-link" :class="{ active: $route.path.startsWith('/menu') }">Menu</router-link>
    </li>
    <li>
      <router-link to="/orders" class="nav-link" :class="{ active: $route.path.startsWith('/orders') }">Orders</router-link>
    </li>
    <li>
      <router-link to="/cart" class="nav-link" :class="{ active: $route.path.startsWith('/cart') }">Shopping Cart</router-link>
    </li>
    <li>
      <router-link to="/orders-dashboard" class="nav-link" :class="{ active: $route.path.startsWith('/orders-dashboard') }">Orders Dashboard</router-link>
    </li>
    <li>
      <router-link to="/delivery-dashboard" class="nav-link" :class="{ active: $route.path.startsWith('/delivery-dashboard') }">Delivery Dashboard</router-link>
    </li>
    <li>
      <router-link to="/owner-dashboard" class="nav-link" :class="{ active: $route.path.startsWith('/owner-dashboard') }">Owner Dashboard</router-link>
    </li>
    <li>
      <router-link to="/profile" class="nav-link" :class="{ active: $route.path.startsWith('/profile') }">Profile</router-link>
    </li>
  </ul>
</template>
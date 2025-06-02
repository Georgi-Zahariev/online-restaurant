<template>
  <div class="app">
    <!-- Sidebar Navigation -->
    <aside class="sidebar">
      <ul class="nav-list">
        <!--
          If you have Router installed, replace `href="#"` with `to="/menu"` (etc.)
          and import `<router-link>`. For now, these are plain buttons/links.
        -->
        <li>
          <router-link
            to="/menu"
            class="nav-link"
            :class="{ active: activeRoute === '/menu' }"
            @click.native="activeRoute = '/menu'"
          >
            Menu
          </router-link>
        </li>
        <li>
          <router-link
            to="/orders"
            class="nav-link"
            :class="{ active: activeRoute === '/orders' }"
            @click.native="activeRoute = '/orders'"
          >
            Orders
          </router-link>
        </li>
        <li>
          <router-link
            to="/cart"
            class="nav-link"
            :class="{ active: activeRoute === '/cart' }"
            @click.native="activeRoute = '/cart'"
          >
            Shopping Cart
          </router-link>
        </li>
        <li>
          <router-link
            to="/orders-dashboard"
            class="nav-link"
            :class="{ active: activeRoute === '/orders-dashboard' }"
            @click.native="activeRoute = '/orders-dashboard'"
          >
            Orders Dashboard
          </router-link>
        </li>
        <li>
          <router-link
            to="/delivery-dashboard"
            class="nav-link"
            :class="{ active: activeRoute === '/delivery-dashboard' }"
            @click.native="activeRoute = '/delivery-dashboard'"
          >
            Delivery Dashboard
          </router-link>
        </li>
        <li>
          <router-link
            to="/profile"
            class="nav-link"
            :class="{ active: activeRoute === '/profile' }"
            @click.native="activeRoute = '/profile'"
          >
            Profile
          </router-link>
        </li>
      </ul>
    </aside>

    <!-- Main Content Area -->
    <main class="main-content">
      <!-- This is where each routed view will render -->
      <router-view />
    </main>
  </div>
</template>

<script>
export default {
  name: "App", // or MainLayout if you moved this into src/layouts/MainLayout.vue
  data() {
    return {
      // We track the “active” route here so we can add a CSS class to highlight the currently selected link.
      // If you prefer, you can also rely on <router-link v-slot="{ isActive }"> and skip this data altogether.
      activeRoute: "/menu", 
    };
  },
  mounted() {
    // When the app first loads, sync `activeRoute` with whatever the current URL is.
    this.activeRoute = this.$route.path;
    // watch `$route` so that if the user navigates via browser back/forward, `activeRoute` stays in sync:
    this.$watch(
      () => this.$route.path,
      (newPath) => {
        this.activeRoute = newPath;
      }
    );
  },
};
</script>

<style scoped>
/* Make the app container fill 100% height of viewport, split into sidebar + content */
.app {
  display: flex;
  height: 100vh;
  margin: 0;
}

/* Sidebar styles */
.sidebar {
  width: 220px;              /* adjust as needed */
  background-color: #f4f4f4; /* light gray background */
  padding: 24px 16px;
  box-shadow: 2px 0 6px rgba(0, 0, 0, 0.1);
}

/* Remove default list styling, and space out items */
.nav-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.nav-list li + li {
  margin-top: 12px;
}

/* Base style for each navigation link */
.nav-link {
  display: block;
  width: 100%;
  padding: 10px 14px;
  border: 1px solid #bbb;
  border-radius: 4px;
  text-decoration: none;
  color: #333;
  font-weight: 500;
  background-color: #fff;
  transition: background-color 0.2s ease, border-color 0.2s ease;
}

/* Hover/focus for better UX */
.nav-link:hover,
.nav-link:focus {
  background-color: #eef2f3;
  border-color: #999;
}

/* Active (selected) link styling */
.nav-link.active {
  background-color: #d5e8d4; /* pale green */
  border-color: #7bb37d;     /* darker green border */
  color: #2f5f31;            /* dark text so it’s readable */
}

/* Main content area to the right */
.main-content {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
}

main {
  width: 100%;
}
</style>

<style>
html, body, #app {
  height: 100vh;
  width: 100vw;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  overflow: hidden;
}
</style>
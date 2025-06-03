<template>
  <div class="app">
    <!-- Sidebar Navigation -->
    <aside class="sidebar">
      <component :is="sidebarComponent" />
    </aside>

    <!-- Main Content Area -->
    <main class="main-content">
      <!-- This is where each routed view will render -->
      <router-view />
    </main>
  </div>
</template>

<script>
import MainSidebar from './components/MainSidebar.vue'
import MenuSidebar from './components/MenuSidebar.vue'

export default {
  name: "App", // or MainLayout if you moved this into src/layouts/MainLayout.vue
  
  computed: {
    sidebarComponent() {
      // Show MenuSidebar only on /menu and its subroutes
      if (this.$route.path.startsWith('/menu')) {
        return MenuSidebar
      }
      return MainSidebar
    }
  },
};
</script>

<style>
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
  padding: 18px 20px;
  margin-bottom: 14px;
  border: 2px solid #888;
  box-sizing: border-box;
  border-radius: 4px;
  text-decoration: none;
  color: #222;
  font-size: 1.25rem;
  font-weight: 600;
  background-color: #fff;
  transition: background-color 0.2s ease, border-color 0.2s ease;
}

/* Hover/focus for better UX */
.nav-link:hover,
.nav-link:focus {
  background-color: #eef2f3;
  border-color: #999;
  box-shadow: 0 4px 16px rgba(76,175,80,0.08);

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

html, body, #app {
  height: 100vh;
  width: 100vw;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  overflow: hidden;
}
</style>
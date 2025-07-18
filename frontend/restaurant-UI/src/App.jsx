import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Header from './components/Header'
import Footer from './components/Footer'
import HomePage from './pages/HomePage'
import MenuPage from './pages/MenuPage'
import CartPage from './pages/CartPage'
import ContactPage from './pages/ContactPage'
import ProfilePage from './pages/ProfilePage'
import OrderDashboard from './pages/OrderDashboard'
import DeliveryDashboard from './pages/DeliveryDashboard'
import { CartProvider } from './utils/CartContext'
import './App.css'

function App() {
  return (
    <CartProvider>
      <Router>
        <div className="App">
          <Header />
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/menu" element={<MenuPage />} />
            <Route path="/cart" element={<CartPage />} />
            <Route path="/contact" element={<ContactPage />} />
            <Route path="/profile" element={<ProfilePage />} />
            <Route path="/dashboard" element={<OrderDashboard />} />
            <Route path="/delivery" element={<DeliveryDashboard />} />
          </Routes>
          <Footer />
        </div>
      </Router>
    </CartProvider>
  )
}

export default App

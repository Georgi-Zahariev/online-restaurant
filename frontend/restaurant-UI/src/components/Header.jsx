import { useState } from 'react'
import { Link, useLocation } from 'react-router-dom'
import { Menu, X, ShoppingCart, User } from 'lucide-react'
import { useCart } from '../utils/CartContext'

const Header = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false)
  const { cartItems } = useCart()
  const location = useLocation()

  const navItems = [
    { path: '/', label: 'Home' },
    { path: '/menu', label: 'Menu' },
    { path: '/contact', label: 'Contact' },
    { path: '/profile', label: 'Profile' },
    { path: '/dashboard', label: 'Kitchen', className: 'kitchen-link' }, // Will be hidden for non-kitchen staff later
    { path: '/delivery', label: 'Delivery', className: 'delivery-link' } // Will be hidden for non-delivery staff later
  ]

  const cartItemCount = cartItems.reduce((total, item) => total + item.quantity, 0)

  return (
    <header className="header">
      <div className="header-content">
        <Link to="/" className="logo">
          Georgi's Grillhouse
        </Link>
        
        <nav className="nav-desktop">
          <ul className="nav-links">
            {navItems.map((item) => (
              <li key={item.path}>
                <Link 
                  to={item.path}
                  className={`${location.pathname === item.path ? 'active' : ''} ${item.className || ''}`}
                >
                  {item.label}
                </Link>
              </li>
            ))}
            <li>
              <Link to="/cart" className="cart-link">
                <ShoppingCart size={20} />
                {cartItemCount > 0 && (
                  <span className="cart-badge">{cartItemCount}</span>
                )}
              </Link>
            </li>
          </ul>
        </nav>

        <button 
          className="mobile-menu-btn"
          onClick={() => setIsMenuOpen(!isMenuOpen)}
        >
          {isMenuOpen ? <X size={24} /> : <Menu size={24} />}
        </button>

        {isMenuOpen && (
          <nav className="nav-mobile">
            <ul className="nav-links-mobile">
              {navItems.map((item) => (
                <li key={item.path}>
                  <Link 
                    to={item.path}
                    onClick={() => setIsMenuOpen(false)}
                    className={`${location.pathname === item.path ? 'active' : ''} ${item.className || ''}`}
                  >
                    {item.label}
                  </Link>
                </li>
              ))}
              <li>
                <Link 
                  to="/cart" 
                  onClick={() => setIsMenuOpen(false)}
                  className="cart-link-mobile"
                >
                  <ShoppingCart size={20} />
                  Cart ({cartItemCount})
                </Link>
              </li>
            </ul>
          </nav>
        )}
      </div>
    </header>
  )
}

export default Header

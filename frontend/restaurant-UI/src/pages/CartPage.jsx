import { useState } from 'react'
import { useCart } from '../utils/CartContext'
import { Plus, Minus, Trash2, ShoppingBag } from 'lucide-react'
import { Link } from 'react-router-dom'

const CartPage = () => {
  const { cartItems, updateQuantity, removeItem, clearCart, getTotalPrice } = useCart()
  const [isCheckingOut, setIsCheckingOut] = useState(false)

  const handleQuantityChange = (id, newQuantity) => {
    updateQuantity(id, newQuantity)
  }

  const handleCheckout = () => {
    setIsCheckingOut(true)
    // Simulate checkout process
    setTimeout(() => {
      alert('Order placed successfully!')
      clearCart()
      setIsCheckingOut(false)
    }, 2000)
  }

  if (cartItems.length === 0) {
    return (
      <div className="cart-page">
        <div className="section">
          <div className="empty-cart">
            <ShoppingBag size={64} color="#ccc" />
            <h2>Your cart is empty</h2>
            <p>Add some delicious items from our menu!</p>
            <Link to="/menu" className="cta-button">Browse Menu</Link>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="cart-page">
      <div className="section">
        <h1 className="section-title">Your Cart</h1>
        
        <div className="cart-content">
          <div className="cart-items">
            {cartItems.map((item) => (
              <div key={item.id} className="cart-item">
                <div className="cart-item-image">
                  <img 
                    src={item.image} 
                    alt={item.name}
                    onError={(e) => {
                      e.target.style.display = 'none';
                      e.target.nextSibling.style.display = 'flex';
                    }}
                  />
                  <div className="image-placeholder" style={{ display: 'none' }}>
                    {item.name}
                  </div>
                </div>
                <div className="cart-item-details">
                  <h3>{item.name}</h3>
                  <p>{item.description}</p>
                  <div className="cart-item-price">${item.price}</div>
                </div>
                <div className="cart-item-controls">
                  <div className="quantity-controls">
                    <button 
                      onClick={() => handleQuantityChange(item.id, item.quantity - 1)}
                      className="quantity-btn"
                    >
                      <Minus size={16} />
                    </button>
                    <span className="quantity">{item.quantity}</span>
                    <button 
                      onClick={() => handleQuantityChange(item.id, item.quantity + 1)}
                      className="quantity-btn"
                    >
                      <Plus size={16} />
                    </button>
                  </div>
                  <button 
                    onClick={() => removeItem(item.id)}
                    className="remove-btn"
                  >
                    <Trash2 size={16} />
                  </button>
                </div>
                <div className="cart-item-total">
                  ${(item.price * item.quantity).toFixed(2)}
                </div>
              </div>
            ))}
          </div>

          <div className="cart-summary">
            <div className="summary-card">
              <h3>Order Summary</h3>
              <div className="summary-row">
                <span>Subtotal:</span>
                <span>${getTotalPrice().toFixed(2)}</span>
              </div>
              <div className="summary-row">
                <span>Tax (8.5%):</span>
                <span>${(getTotalPrice() * 0.085).toFixed(2)}</span>
              </div>
              <div className="summary-row">
                <span>Delivery Fee:</span>
                <span>$3.99</span>
              </div>
              <div className="summary-row total">
                <span>Total:</span>
                <span>${(getTotalPrice() + getTotalPrice() * 0.085 + 3.99).toFixed(2)}</span>
              </div>
              
              <button 
                className="checkout-btn"
                onClick={handleCheckout}
                disabled={isCheckingOut}
              >
                {isCheckingOut ? 'Processing...' : 'Checkout'}
              </button>
              
              <Link to="/menu" className="continue-shopping">
                Continue Shopping
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default CartPage

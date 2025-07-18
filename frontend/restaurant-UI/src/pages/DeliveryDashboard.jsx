import { useState, useEffect } from 'react'
import { Package, MapPin, Phone, User, Clock, CheckCircle, Truck, AlertCircle, Timer, Navigation } from 'lucide-react'

const DeliveryDashboard = () => {
  const [orders, setOrders] = useState([
    {
      id: '123e4567-e89b-12d3-a456-426614174000',
      customerName: 'John Doe',
      customerPhone: '+1 (555) 123-4567',
      address: '123 Main Street, Downtown, 12345',
      items: [
        { name: 'Grilled Chicken Breast', quantity: 2 },
        { name: 'Caesar Salad', quantity: 1 },
        { name: 'Grilled Vegetables', quantity: 1 }
      ],
      status: 'ready-for-pickup', // ready-for-pickup, out-for-delivery, delivered
      orderTime: '2025-07-17T14:30:00Z',
      readyTime: '2025-07-17T15:00:00Z',
      totalPrice: 45.99,
      deliveryNotes: 'Ring doorbell twice. Leave at door if no answer.',
      assignedDriver: null,
      estimatedDeliveryTime: 20
    },
    {
      id: '456e7890-e89b-12d3-a456-426614174001',
      customerName: 'Jane Smith',
      customerPhone: '+1 (555) 987-6543',
      address: '456 Oak Avenue, Uptown, 67890',
      items: [
        { name: 'BBQ Ribs', quantity: 1 },
        { name: 'Coleslaw', quantity: 1 },
        { name: 'French Fries', quantity: 2 }
      ],
      status: 'out-for-delivery',
      orderTime: '2025-07-17T14:45:00Z',
      readyTime: '2025-07-17T15:10:00Z',
      totalPrice: 32.50,
      deliveryNotes: 'Apartment 3B. Call when arrived.',
      assignedDriver: 'Mike Wilson',
      estimatedDeliveryTime: 15,
      pickupTime: '2025-07-17T15:15:00Z'
    },
    {
      id: '789e0123-e89b-12d3-a456-426614174002',
      customerName: 'Mike Johnson',
      customerPhone: '+1 (555) 456-7890',
      address: '789 Pine Street, Midtown, 54321',
      items: [
        { name: 'Fish and Chips', quantity: 1 },
        { name: 'Garden Salad', quantity: 1 }
      ],
      status: 'ready-for-pickup',
      orderTime: '2025-07-17T14:20:00Z',
      readyTime: '2025-07-17T14:55:00Z',
      totalPrice: 28.75,
      deliveryNotes: 'House with red door. Please be quiet, baby sleeping.',
      assignedDriver: null,
      estimatedDeliveryTime: 25
    },
    {
      id: '321e6547-e89b-12d3-a456-426614174003',
      customerName: 'Sarah Davis',
      customerPhone: '+1 (555) 234-5678',
      address: '321 Elm Street, Westside, 98765',
      items: [
        { name: 'Pasta Carbonara', quantity: 1 },
        { name: 'Garlic Bread', quantity: 2 }
      ],
      status: 'out-for-delivery',
      orderTime: '2025-07-17T13:30:00Z',
      readyTime: '2025-07-17T14:15:00Z',
      totalPrice: 22.99,
      deliveryNotes: 'Business address. Ask for Sarah at reception.',
      assignedDriver: 'Alex Johnson',
      estimatedDeliveryTime: 10,
      pickupTime: '2025-07-17T14:20:00Z'
    }
  ])

  const [filter, setFilter] = useState('all')
  const [currentDriver] = useState('Current User') // This would come from auth later

  const filteredOrders = orders.filter(order => {
    if (filter === 'all') return true
    if (filter === 'available') return order.status === 'ready-for-pickup'
    if (filter === 'my-deliveries') return order.assignedDriver === currentDriver
    return order.status === filter
  })

  const takeOrder = (orderId) => {
    setOrders(prevOrders => 
      prevOrders.map(order => 
        order.id === orderId 
          ? { 
              ...order, 
              status: 'out-for-delivery', 
              assignedDriver: currentDriver,
              pickupTime: new Date().toISOString()
            }
          : order
      )
    )
  }

  const completeDelivery = (orderId) => {
    setOrders(prevOrders => 
      prevOrders.map(order => 
        order.id === orderId 
          ? { 
              ...order, 
              status: 'delivered',
              deliveryTime: new Date().toISOString()
            }
          : order
      )
    )
  }

  const formatTime = (timeString) => {
    return new Date(timeString).toLocaleTimeString('en-US', {
      hour: '2-digit',
      minute: '2-digit'
    })
  }

  const getTimeElapsed = (timeString) => {
    const now = new Date()
    const orderTime = new Date(timeString)
    return Math.floor((now - orderTime) / (1000 * 60))
  }

  const getDeliveryTime = (pickupTime) => {
    const now = new Date()
    const pickup = new Date(pickupTime)
    return Math.floor((now - pickup) / (1000 * 60))
  }

  const readyOrdersCount = orders.filter(o => o.status === 'ready-for-pickup').length
  const myActiveDeliveries = orders.filter(o => o.assignedDriver === currentDriver && o.status === 'out-for-delivery').length
  const totalActiveDeliveries = orders.filter(o => o.status === 'out-for-delivery').length

  return (
    <div className="delivery-dashboard-page">
      <div className="delivery-dashboard-container">
        <div className="dashboard-header">
          <h1>🚚 Delivery Dashboard</h1>
          <p>Manage pickups and deliveries</p>
        </div>

        {/* Stats */}
        <div className="delivery-stats">
          <div className="stat-card ready-stat">
            <Package size={32} />
            <div>
              <h3>{readyOrdersCount}</h3>
              <p>Ready for Pickup</p>
            </div>
          </div>
          <div className="stat-card my-deliveries-stat">
            <Truck size={32} />
            <div>
              <h3>{myActiveDeliveries}</h3>
              <p>My Active Deliveries</p>
            </div>
          </div>
          <div className="stat-card total-stat">
            <Navigation size={32} />
            <div>
              <h3>{totalActiveDeliveries}</h3>
              <p>Total Out for Delivery</p>
            </div>
          </div>
        </div>

        {/* Filter Buttons */}
        <div className="delivery-filter-buttons">
          <button 
            className={`filter-btn ${filter === 'all' ? 'active' : ''}`}
            onClick={() => setFilter('all')}
          >
            All Orders ({orders.length})
          </button>
          <button 
            className={`filter-btn ${filter === 'available' ? 'active' : ''}`}
            onClick={() => setFilter('available')}
          >
            Available ({readyOrdersCount})
          </button>
          <button 
            className={`filter-btn ${filter === 'my-deliveries' ? 'active' : ''}`}
            onClick={() => setFilter('my-deliveries')}
          >
            My Deliveries ({orders.filter(o => o.assignedDriver === currentDriver).length})
          </button>
          <button 
            className={`filter-btn ${filter === 'out-for-delivery' ? 'active' : ''}`}
            onClick={() => setFilter('out-for-delivery')}
          >
            Out for Delivery ({totalActiveDeliveries})
          </button>
        </div>

        {/* Orders Grid */}
        <div className="delivery-orders-grid">
          {filteredOrders.map(order => (
            <div key={order.id} className={`delivery-order-card ${order.status}`}>
              {/* Order Header */}
              <div className="delivery-order-header">
                <div className="order-info">
                  <h3>Order #{order.id.slice(-8)}</h3>
                  <span className={`delivery-status-badge status-${order.status}`}>
                    {order.status === 'ready-for-pickup' && <Package size={16} />}
                    {order.status === 'out-for-delivery' && <Truck size={16} />}
                    {order.status === 'delivered' && <CheckCircle size={16} />}
                    {order.status.replace('-', ' ').toUpperCase()}
                  </span>
                </div>
                <div className="order-timing">
                  <p>Ordered: {formatTime(order.orderTime)}</p>
                  <p>Ready: {formatTime(order.readyTime)}</p>
                  {order.pickupTime && (
                    <p className="pickup-time">
                      <Timer size={14} />
                      Picked up: {formatTime(order.pickupTime)}
                    </p>
                  )}
                </div>
              </div>

              {/* Customer Info */}
              <div className="delivery-customer-info">
                <div className="customer-detail">
                  <User size={16} />
                  <span className="customer-name">{order.customerName}</span>
                  <a href={`tel:${order.customerPhone}`} className="customer-phone">
                    <Phone size={16} />
                    {order.customerPhone}
                  </a>
                </div>
                <div className="delivery-address">
                  <MapPin size={16} />
                  <span>{order.address}</span>
                  <button 
                    className="maps-btn"
                    onClick={() => window.open(`https://maps.google.com/?q=${encodeURIComponent(order.address)}`, '_blank')}
                  >
                    Open in Maps
                  </button>
                </div>
                {order.deliveryNotes && (
                  <div className="delivery-notes">
                    <AlertCircle size={16} />
                    <span><strong>Notes:</strong> {order.deliveryNotes}</span>
                  </div>
                )}
              </div>

              {/* Order Items */}
              <div className="delivery-order-items">
                <h4>Items ({order.items.reduce((sum, item) => sum + item.quantity, 0)}):</h4>
                <div className="items-list">
                  {order.items.map((item, index) => (
                    <span key={index} className="item-tag">
                      {item.quantity}x {item.name}
                    </span>
                  ))}
                </div>
              </div>

              {/* Driver Assignment */}
              {order.assignedDriver && (
                <div className="driver-assignment">
                  <Truck size={16} />
                  <span>Driver: <strong>{order.assignedDriver}</strong></span>
                  {order.pickupTime && (
                    <span className="delivery-duration">
                      ({getDeliveryTime(order.pickupTime)} min on route)
                    </span>
                  )}
                </div>
              )}

              {/* Order Footer */}
              <div className="delivery-order-footer">
                <div className="order-total">
                  <strong>Total: ${order.totalPrice.toFixed(2)}</strong>
                  <span>Est. delivery: {order.estimatedDeliveryTime} min</span>
                </div>
                <div className="delivery-actions">
                  {order.status === 'ready-for-pickup' && (
                    <button 
                      className="delivery-action-btn take-btn"
                      onClick={() => takeOrder(order.id)}
                    >
                      <Package size={16} />
                      Take Order
                    </button>
                  )}
                  
                  {order.status === 'out-for-delivery' && order.assignedDriver === currentDriver && (
                    <button 
                      className="delivery-action-btn complete-btn"
                      onClick={() => completeDelivery(order.id)}
                    >
                      <CheckCircle size={16} />
                      Mark Delivered
                    </button>
                  )}

                  {order.status === 'delivered' && (
                    <div className="delivered-badge">
                      <CheckCircle size={16} />
                      Delivered
                    </div>
                  )}
                </div>
              </div>
            </div>
          ))}
        </div>

        {/* Empty State */}
        {filteredOrders.length === 0 && (
          <div className="delivery-empty-state">
            <Package size={48} />
            <h3>No orders to display</h3>
            <p>
              {filter === 'available' && 'No orders ready for pickup at the moment.'}
              {filter === 'my-deliveries' && 'You have no assigned deliveries.'}
              {filter === 'out-for-delivery' && 'No orders are currently out for delivery.'}
              {filter === 'all' && 'No orders in the system right now.'}
            </p>
          </div>
        )}
      </div>
    </div>
  )
}

export default DeliveryDashboard

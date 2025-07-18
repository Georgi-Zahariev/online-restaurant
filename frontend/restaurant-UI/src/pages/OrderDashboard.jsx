import { useState, useEffect } from 'react'
import { Clock, CheckCircle, AlertCircle, User, MapPin, Bell, Search, Printer, RefreshCw, Timer } from 'lucide-react'

const OrderDashboard = () => {
  const [orders, setOrders] = useState([
    {
      id: '123e4567-e89b-12d3-a456-426614174000',
      customerName: 'John Doe',
      customerPhone: '+1 (555) 123-4567',
      address: '123 Main Street, Downtown, 12345',
      items: [
        { name: 'Grilled Chicken Breast', quantity: 2, comments: 'Medium-well, no sauce' },
        { name: 'Caesar Salad', quantity: 1, comments: 'Extra dressing on the side' },
        { name: 'Grilled Vegetables', quantity: 1, comments: '' }
      ],
      status: 'pending',
      priority: 'normal',
      orderTime: '2025-07-17T14:30:00Z',
      estimatedTime: 25,
      totalPrice: 45.99,
      cookingStartTime: null
    },
    {
      id: '456e7890-e89b-12d3-a456-426614174001',
      customerName: 'Jane Smith',
      customerPhone: '+1 (555) 987-6543',
      address: '456 Oak Avenue, Uptown, 67890',
      items: [
        { name: 'BBQ Ribs', quantity: 1, comments: 'Extra BBQ sauce' },
        { name: 'Coleslaw', quantity: 1, comments: '' },
        { name: 'French Fries', quantity: 2, comments: 'Crispy' }
      ],
      status: 'in-progress',
      priority: 'high',
      orderTime: '2025-07-17T14:45:00Z',
      estimatedTime: 35,
      totalPrice: 32.50,
      cookingStartTime: '2025-07-17T14:50:00Z'
    },
    {
      id: '789e0123-e89b-12d3-a456-426614174002',
      customerName: 'Mike Johnson',
      customerPhone: '+1 (555) 456-7890',
      address: '789 Pine Street, Midtown, 54321',
      items: [
        { name: 'Fish and Chips', quantity: 1, comments: 'Malt vinegar on the side' },
        { name: 'Garden Salad', quantity: 1, comments: 'No onions' }
      ],
      status: 'pending',
      priority: 'urgent',
      orderTime: '2025-07-17T15:00:00Z',
      estimatedTime: 20,
      totalPrice: 28.75,
      cookingStartTime: null
    }
  ])

  const [filter, setFilter] = useState('all')
  const [searchTerm, setSearchTerm] = useState('')
  const [soundEnabled, setSoundEnabled] = useState(true)

  const filteredOrders = orders.filter(order => {
    const matchesFilter = filter === 'all' || order.status === filter
    const matchesSearch = searchTerm === '' || 
      order.customerName.toLowerCase().includes(searchTerm.toLowerCase()) ||
      order.id.includes(searchTerm) ||
      order.items.some(item => item.name.toLowerCase().includes(searchTerm.toLowerCase()))
    
    return matchesFilter && matchesSearch
  }).sort((a, b) => {
    const priorityOrder = { urgent: 3, high: 2, normal: 1 }
    if (priorityOrder[a.priority] !== priorityOrder[b.priority]) {
      return priorityOrder[b.priority] - priorityOrder[a.priority]
    }
    return new Date(a.orderTime) - new Date(b.orderTime)
  })

  const getTimeElapsed = (orderTime) => {
    const now = new Date()
    const orderDate = new Date(orderTime)
    const diffMinutes = Math.floor((now - orderDate) / (1000 * 60))
    return diffMinutes
  }

  const getCookingTime = (cookingStartTime) => {
    if (!cookingStartTime) return null
    const now = new Date()
    const startTime = new Date(cookingStartTime)
    const diffMinutes = Math.floor((now - startTime) / (1000 * 60))
    return diffMinutes
  }

  const updateOrderStatus = (orderId, newStatus) => {
    setOrders(prevOrders => 
      prevOrders.map(order => {
        if (order.id === orderId) {
          const updatedOrder = { ...order, status: newStatus }
          if (newStatus === 'in-progress' && !order.cookingStartTime) {
            updatedOrder.cookingStartTime = new Date().toISOString()
          }
          return updatedOrder
        }
        return order
      })
    )
  }

  const completeOrder = (orderId) => {
    setOrders(prevOrders => 
      prevOrders.filter(order => order.id !== orderId)
    )
    console.log('Order completed:', orderId)
  }

  const formatTime = (timeString) => {
    return new Date(timeString).toLocaleTimeString('en-US', {
      hour: '2-digit',
      minute: '2-digit'
    })
  }

  return (
    <div className="dashboard-page">
      <div className="dashboard-container">
        <div className="dashboard-header">
          <h1>Kitchen Dashboard</h1>
          <p>Manage incoming orders and track preparation progress</p>
        </div>

        {/* Search and Controls */}
        <div className="dashboard-controls">
          <div className="search-container">
            <Search size={20} />
            <input
              type="text"
              placeholder="Search orders, customers, or items..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="search-input"
            />
          </div>
          
          <button 
            className={`control-btn ${soundEnabled ? 'active' : ''}`}
            onClick={() => setSoundEnabled(!soundEnabled)}
            title="Toggle sound notifications"
          >
            <Bell size={18} />
          </button>
        </div>

        {/* Statistics */}
        <div className="dashboard-stats">
          <div className="stat-card">
            <h3>{orders.filter(o => o.status === 'pending').length}</h3>
            <p>Pending Orders</p>
          </div>
          <div className="stat-card">
            <h3>{orders.filter(o => o.status === 'in-progress').length}</h3>
            <p>In Progress</p>
          </div>
          <div className="stat-card">
            <h3>{orders.filter(o => o.priority === 'urgent').length}</h3>
            <p>Urgent Priority</p>
          </div>
        </div>

        {/* Filters */}
        <div className="dashboard-filters">
          <button 
            className={`filter-btn ${filter === 'all' ? 'active' : ''}`}
            onClick={() => setFilter('all')}
          >
            All Orders ({orders.length})
          </button>
          <button 
            className={`filter-btn ${filter === 'pending' ? 'active' : ''}`}
            onClick={() => setFilter('pending')}
          >
            Pending ({orders.filter(o => o.status === 'pending').length})
          </button>
          <button 
            className={`filter-btn ${filter === 'in-progress' ? 'active' : ''}`}
            onClick={() => setFilter('in-progress')}
          >
            In Progress ({orders.filter(o => o.status === 'in-progress').length})
          </button>
        </div>

        {/* Orders Grid */}
        <div className="orders-grid">
          {filteredOrders.map(order => (
            <div key={order.id} className="order-card">
              {/* Order Header */}
              <div className="order-header">
                <div className="order-info">
                  <h3>Order #{order.id.slice(-8)}</h3>
                  <div className="badges">
                    <span className={`status-badge status-${order.status}`}>
                      {order.status === 'pending' && <AlertCircle size={16} />}
                      {order.status === 'in-progress' && <Clock size={16} />}
                      {order.status.replace('-', ' ').toUpperCase()}
                    </span>
                    <span className={`priority-badge priority-${order.priority}`}>
                      {order.priority.toUpperCase()}
                    </span>
                  </div>
                </div>
                <div className="order-time">
                  <p>Ordered: {formatTime(order.orderTime)}</p>
                  <p className="elapsed-time">
                    {getTimeElapsed(order.orderTime)} min ago
                  </p>
                  {order.cookingStartTime && (
                    <p className="cooking-time">
                      <Timer size={14} />
                      Cooking: {getCookingTime(order.cookingStartTime)} min
                    </p>
                  )}
                </div>
              </div>

              {/* Customer Info */}
              <div className="customer-info">
                <div className="customer-detail">
                  <User size={16} />
                  <span>{order.customerName}</span>
                  <span className="phone">{order.customerPhone}</span>
                </div>
                <div className="customer-detail">
                  <MapPin size={16} />
                  <span>{order.address}</span>
                </div>
              </div>

              {/* Order Items */}
              <div className="order-items">
                <h4>Items to Prepare:</h4>
                {order.items.map((item, index) => (
                  <div key={index} className="order-item">
                    <div className="item-info">
                      <span className="item-name">
                        {item.quantity}x {item.name}
                      </span>
                      {item.comments && (
                        <span className="item-comments">
                          Note: {item.comments}
                        </span>
                      )}
                    </div>
                  </div>
                ))}
              </div>

              {/* Order Footer */}
              <div className="order-footer">
                <div className="order-total">
                  <strong>Total: ${order.totalPrice.toFixed(2)}</strong>
                  <span>Est. {order.estimatedTime} min</span>
                </div>
                <div className="order-actions">
                  {order.status === 'pending' && (
                    <button 
                      className="action-btn start-btn"
                      onClick={() => updateOrderStatus(order.id, 'in-progress')}
                    >
                      Start Cooking
                    </button>
                  )}
                  
                  {order.status === 'in-progress' && (
                    <button 
                      className="action-btn complete-btn"
                      onClick={() => completeOrder(order.id)}
                    >
                      <CheckCircle size={16} />
                      Mark Complete
                    </button>
                  )}
                </div>
              </div>
            </div>
          ))}
        </div>

        {/* Empty State */}
        {filteredOrders.length === 0 && (
          <div className="empty-state">
            <CheckCircle size={48} />
            <h3>No orders to display</h3>
            <p>
              {searchTerm 
                ? `No orders found matching "${searchTerm}"`
                : filter === 'all' 
                  ? 'No orders in the system right now.' 
                  : `No ${filter.replace('-', ' ')} orders at the moment.`
              }
            </p>
          </div>
        )}
      </div>
    </div>
  )
}

export default OrderDashboard

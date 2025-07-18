import { useState } from 'react'
import { useCart } from '../utils/CartContext'
import { Plus, Star } from 'lucide-react'

const MenuPage = () => {
  const { addItem } = useCart()
  const [selectedCategory, setSelectedCategory] = useState('all')

  // Sample menu data based on your database structure
  const menuItems = [
    {
      id: 1,
      name: "Grilled Salmon",
      description: "Fresh Atlantic salmon grilled to perfection with herbs and lemon",
      price: 24.99,
      category: "main",
      image: "/images/dishes/salmon-grilled.jpg", // Public folder path
      rating: 4.8
    },
    {
      id: 2,
      name: "Beef Tenderloin",
      description: "Premium beef tenderloin with garlic mashed potatoes and seasonal vegetables",
      price: 32.99,
      category: "main",
      image: "/images/dishes/beef-tenderloin.jpg", // Public folder path
      rating: 4.9
    },
    {
      id: 3,
      name: "Vegetarian Pasta",
      description: "Homemade pasta with seasonal vegetables and parmesan cheese",
      price: 18.99,
      category: "main",
      image: "/images/dishes/pasta-vegetarian.jpg", // Public folder path
      rating: 4.6
    },
    {
      id: 4,
      name: "Caesar Salad",
      description: "Fresh romaine lettuce with parmesan, croutons, and Caesar dressing",
      price: 12.99,
      category: "appetizer",
      image: "/images/dishes/caesar-salad.jpg", // Public folder path
      rating: 4.4
    },
    {
      id: 5,
      name: "Bruschetta",
      description: "Toasted bread topped with tomatoes, basil, and balsamic glaze",
      price: 9.99,
      category: "appetizer",
      image: "/images/dishes/bruschetta.jpg", // Public folder path
      rating: 4.5
    },
    {
      id: 6,
      name: "Chocolate Lava Cake",
      description: "Warm chocolate cake with molten center, served with vanilla ice cream",
      price: 8.99,
      category: "dessert",
      image: "/images/dishes/chocolate-lava-cake.jpg", // Public folder path
      rating: 4.7
    },
    {
      id: 7,
      name: "Tiramisu",
      description: "Classic Italian dessert with coffee-soaked ladyfingers and mascarpone",
      price: 7.99,
      category: "dessert",
      image: "/images/dishes/tiramisu.jpg", // Public folder path
      rating: 4.6
    },
    {
      id: 8,
      name: "Fresh Juice",
      description: "Freshly squeezed orange juice",
      price: 4.99,
      category: "beverage",
      image: "/images/dishes/orange-juice.jpg", // Public folder path
      rating: 4.3
    },
    {
      id: 9,
      name: "Craft Beer",
      description: "Local craft beer selection",
      price: 6.99,
      category: "beverage",
      image: "/images/dishes/craft-beer.jpg", // Public folder path
      rating: 4.4
    }
  ]

  const categories = [
    { id: 'all', name: 'All Items' },
    { id: 'appetizer', name: 'Appetizers' },
    { id: 'main', name: 'Main Courses' },
    { id: 'dessert', name: 'Desserts' },
    { id: 'beverage', name: 'Beverages' }
  ]

  const filteredItems = selectedCategory === 'all' 
    ? menuItems 
    : menuItems.filter(item => item.category === selectedCategory)

  const handleAddToCart = (item) => {
    addItem(item)
  }

  return (
    <div className="menu-page">
      <div className="section">
        <h1 className="section-title">Our Menu</h1>
        
        {/* Category Filter */}
        <div className="category-filter">
          {categories.map((category) => (
            <button
              key={category.id}
              className={`category-btn ${selectedCategory === category.id ? 'active' : ''}`}
              onClick={() => setSelectedCategory(category.id)}
            >
              {category.name}
            </button>
          ))}
        </div>

        {/* Menu Items */}
        <div className="menu-grid">
          {filteredItems.map((item) => (
            <div key={item.id} className="menu-card">
              <div className="menu-card-image">
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
              <div className="menu-card-content">
                <h3>{item.name}</h3>
                <p>{item.description}</p>
                <div className="menu-card-rating">
                  <div className="rating">
                    {[...Array(5)].map((_, i) => (
                      <Star 
                        key={i} 
                        size={16} 
                        fill={i < Math.floor(item.rating) ? "#ffd700" : "none"}
                        color="#ffd700" 
                      />
                    ))}
                  </div>
                  <span className="rating-text">({item.rating})</span>
                </div>
                <div className="menu-card-price">${item.price}</div>
                <button 
                  className="add-to-cart-btn"
                  onClick={() => handleAddToCart(item)}
                >
                  <Plus size={16} />
                  Add to Cart
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

export default MenuPage

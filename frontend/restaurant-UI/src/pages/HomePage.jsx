import { Link } from 'react-router-dom'
import { Star, ChefHat, Clock, Heart } from 'lucide-react'

const HomePage = () => {
  const featuredDishes = [
    {
      id: 1,
      name: "Grilled Salmon",
      description: "Fresh Atlantic salmon with herbs and lemon",
      price: 24.99,
      image: "/images/dishes/salmon-grilled.jpg"
    },
    {
      id: 2,
      name: "Beef Tenderloin",
      description: "Premium cut with garlic mashed potatoes",
      price: 32.99,
      image: "/images/dishes/beef-tenderloin.jpg"
    },
    {
      id: 3,
      name: "Vegetarian Pasta",
      description: "Homemade pasta with seasonal vegetables",
      price: 18.99,
      image: "/images/dishes/pasta-vegetarian.jpg"
    }
  ]

  const features = [
    {
      icon: <ChefHat size={32} />,
      title: "Expert Chefs",
      description: "Our skilled chefs create culinary masterpieces with passion and expertise"
    },
    {
      icon: <Clock size={32} />,
      title: "Fast Service",
      description: "Quick and efficient service without compromising on quality"
    },
    {
      icon: <Heart size={32} />,
      title: "Fresh Ingredients",
      description: "We use only the freshest, locally-sourced ingredients in all our dishes"
    }
  ]

  return (
    <div className="homepage">
      {/* Hero Section */}
      <section className="hero">
        <div className="hero-content">
          <h1>Welcome to Our Restaurant</h1>
          <p>Experience exquisite flavors and exceptional dining in a warm, welcoming atmosphere</p>
          <Link to="/menu" className="cta-button">View Our Menu</Link>
        </div>
      </section>

      {/* Featured Dishes */}
      <section className="section">
        <h2 className="section-title">Featured Dishes</h2>
        <div className="menu-grid">
          {featuredDishes.map((dish) => (
            <div key={dish.id} className="menu-card">
              <div className="menu-card-image">
                <img 
                  src={dish.image} 
                  alt={dish.name}
                  onError={(e) => {
                    e.target.style.display = 'none';
                    e.target.nextSibling.style.display = 'flex';
                  }}
                />
                <div className="image-placeholder" style={{ display: 'none' }}>
                  {dish.name}
                </div>
              </div>
              <div className="menu-card-content">
                <h3>{dish.name}</h3>
                <p>{dish.description}</p>
                <div className="menu-card-price">${dish.price}</div>
                <div className="rating">
                  {[...Array(5)].map((_, i) => (
                    <Star key={i} size={16} fill="#ffd700" color="#ffd700" />
                  ))}
                </div>
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* Features */}
      <section className="section features-section">
        <h2 className="section-title">Why Choose Us</h2>
        <div className="features-grid">
          {features.map((feature, index) => (
            <div key={index} className="feature-card">
              <div className="feature-icon">{feature.icon}</div>
              <h3>{feature.title}</h3>
              <p>{feature.description}</p>
            </div>
          ))}
        </div>
      </section>

      {/* Call to Action */}
      <section className="section cta-section">
        <div className="cta-content">
          <h2>Ready to Order?</h2>
          <p>Browse our full menu and place your order today!</p>
          <Link to="/menu" className="cta-button">Order Now</Link>
        </div>
      </section>
    </div>
  )
}

export default HomePage

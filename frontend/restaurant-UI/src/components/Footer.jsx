import { Phone, MapPin, Clock, Mail } from 'lucide-react'

const Footer = () => {
  return (
    <footer className="footer">
      <div className="footer-content">
        <div className="footer-grid">
          <div className="footer-section">
            <h3>Contact Info</h3>
            <div className="footer-item">
              <Phone size={16} />
              <span>+359 894******</span>
            </div>
            <div className="footer-item">
              <Mail size={16} />
              <span>info@restaurant.com</span>
            </div>
            <div className="footer-item">
              <MapPin size={16} />
              <span>112 **** Street, Sofia, Bulgaria, 1000</span>
            </div>
          </div>

          <div className="footer-section">
            <h3>Hours</h3>
            <div className="footer-item">
              <Clock size={16} />
              <div>
                <p>Monday - Thursday: 9:00 AM - 11:00 PM</p>
                <p>Friday - Saturday: 9:00 AM - 12:00 PM</p>
                <p>Sunday: 12:00 PM - 9:00 PM</p>
              </div>
            </div>
          </div>

          <div className="footer-section">
            <h3>About</h3>
            <p>
              We serve delicious, fresh food made with love and the finest ingredients. 
              Come experience our warm atmosphere and exceptional service.
            </p>
          </div>
        </div>

        <div className="footer-bottom">
          <p>&copy; 2025 Georgi's Grillhouse. All rights reserved.</p>
        </div>
      </div>
    </footer>
  )
}

export default Footer

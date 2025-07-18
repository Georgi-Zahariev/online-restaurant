import { useState } from 'react'
import { Phone, MapPin, Clock, Mail, Send } from 'lucide-react'

const ContactPage = () => {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    phone: '',
    message: ''
  })

  const handleInputChange = (e) => {
    const { name, value } = e.target
    setFormData(prev => ({
      ...prev,
      [name]: value
    }))
  }

  const handleSubmit = (e) => {
    e.preventDefault()
    // Handle form submission
    alert('Thank you for your message! We\'ll get back to you soon.')
    setFormData({ name: '', email: '', phone: '', message: '' })
  }

  return (
    <div className="contact-page">
      <div className="section">
        <h1 className="section-title">Contact Us</h1>
        
        <div className="contact-content">
          <div className="contact-info">
            <h2>Get in Touch</h2>
            <p>We'd love to hear from you! Reach out with any questions, comments, or reservations.</p>
            
            <div className="contact-item">
              <div className="contact-icon">
                <Phone size={24} />
              </div>
              <div className="contact-details">
                <h3>Phone</h3>
                <p>+1 (555) 123-4567</p>
                <p>+1 (555) 123-4568</p>
              </div>
            </div>

            <div className="contact-item">
              <div className="contact-icon">
                <Mail size={24} />
              </div>
              <div className="contact-details">
                <h3>Email</h3>
                <p>info@restaurant.com</p>
                <p>reservations@restaurant.com</p>
              </div>
            </div>

            <div className="contact-item">
              <div className="contact-icon">
                <MapPin size={24} />
              </div>
              <div className="contact-details">
                <h3>Address</h3>
                <p>123 Food Street</p>
                <p>City, State 12345</p>
              </div>
            </div>

            <div className="contact-item">
              <div className="contact-icon">
                <Clock size={24} />
              </div>
              <div className="contact-details">
                <h3>Hours</h3>
                <p>Monday - Thursday: 11:00 AM - 9:00 PM</p>
                <p>Friday - Saturday: 11:00 AM - 10:00 PM</p>
                <p>Sunday: 12:00 PM - 8:00 PM</p>
              </div>
            </div>
          </div>

          <div className="contact-form">
            <h2>Send us a Message</h2>
            <form onSubmit={handleSubmit}>
              <div className="form-group">
                <label htmlFor="name">Name *</label>
                <input
                  type="text"
                  id="name"
                  name="name"
                  value={formData.name}
                  onChange={handleInputChange}
                  required
                />
              </div>

              <div className="form-group">
                <label htmlFor="email">Email *</label>
                <input
                  type="email"
                  id="email"
                  name="email"
                  value={formData.email}
                  onChange={handleInputChange}
                  required
                />
              </div>

              <div className="form-group">
                <label htmlFor="phone">Phone</label>
                <input
                  type="tel"
                  id="phone"
                  name="phone"
                  value={formData.phone}
                  onChange={handleInputChange}
                />
              </div>

              <div className="form-group">
                <label htmlFor="message">Message *</label>
                <textarea
                  id="message"
                  name="message"
                  value={formData.message}
                  onChange={handleInputChange}
                  rows="6"
                  required
                ></textarea>
              </div>

              <button type="submit" className="submit-btn">
                <Send size={16} />
                Send Message
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  )
}

export default ContactPage

import { useState } from 'react'
import { User, Mail, Phone, MapPin, Edit2, Save, X } from 'lucide-react'

const ProfilePage = () => {
  const [isEditing, setIsEditing] = useState(false)
  const [profile, setProfile] = useState({
    name: 'John Doe',
    email: 'john.doe@example.com',
    phone: '+1 (555) 123-4567',
    address: {
      street: '123 Main Street',
      city: 'Downtown',
      code: '12345',
      country: 'United States'
    }
  })

  const [editForm, setEditForm] = useState({ ...profile })

  const handleEditToggle = () => {
    if (isEditing) {
      // Reset form to original values if canceling
      setEditForm({ ...profile })
    }
    setIsEditing(!isEditing)
  }

  const handleSave = () => {
    setProfile({ ...editForm })
    setIsEditing(false)
    // TODO: Send update to backend API
    console.log('Profile updated:', editForm)
  }

  const handleInputChange = (field, value) => {
    if (field.includes('.')) {
      const [parent, child] = field.split('.')
      setEditForm(prev => ({
        ...prev,
        [parent]: {
          ...prev[parent],
          [child]: value
        }
      }))
    } else {
      setEditForm(prev => ({
        ...prev,
        [field]: value
      }))
    }
  }

  return (
    <div className="profile-page">
      <div className="profile-container">
        <div className="profile-header">
          <div className="profile-avatar">
            <User size={48} />
          </div>
          <h1>My Profile</h1>
          <p>Manage your personal information and preferences</p>
        </div>

        <div className="profile-content">
          <div className="profile-section">
            <div className="section-header">
              <h2>Personal Information</h2>
              <button 
                className={`edit-btn ${isEditing ? 'cancel' : 'edit'}`}
                onClick={handleEditToggle}
              >
                {isEditing ? <X size={18} /> : <Edit2 size={18} />}
                {isEditing ? 'Cancel' : 'Edit'}
              </button>
            </div>

            <div className="profile-form">
              <div className="form-group">
                <label>
                  <User size={20} />
                  Full Name
                </label>
                {isEditing ? (
                  <input
                    type="text"
                    value={editForm.name}
                    onChange={(e) => handleInputChange('name', e.target.value)}
                    className="form-input"
                  />
                ) : (
                  <span className="form-value">{profile.name}</span>
                )}
              </div>

              <div className="form-group">
                <label>
                  <Mail size={20} />
                  Email Address
                </label>
                {isEditing ? (
                  <input
                    type="email"
                    value={editForm.email}
                    onChange={(e) => handleInputChange('email', e.target.value)}
                    className="form-input"
                  />
                ) : (
                  <span className="form-value">{profile.email}</span>
                )}
              </div>

              <div className="form-group">
                <label>
                  <Phone size={20} />
                  Phone Number
                </label>
                {isEditing ? (
                  <input
                    type="tel"
                    value={editForm.phone}
                    onChange={(e) => handleInputChange('phone', e.target.value)}
                    className="form-input"
                  />
                ) : (
                  <span className="form-value">{profile.phone}</span>
                )}
              </div>
            </div>
          </div>

          <div className="profile-section">
            <div className="section-header">
              <h2>Address Information</h2>
            </div>

            <div className="profile-form">
              <div className="form-group">
                <label>
                  <MapPin size={20} />
                  Street Address
                </label>
                {isEditing ? (
                  <input
                    type="text"
                    value={editForm.address.street}
                    onChange={(e) => handleInputChange('address.street', e.target.value)}
                    className="form-input"
                  />
                ) : (
                  <span className="form-value">{profile.address.street}</span>
                )}
              </div>

              <div className="form-row">
                <div className="form-group">
                  <label>City</label>
                  {isEditing ? (
                    <input
                      type="text"
                      value={editForm.address.city}
                      onChange={(e) => handleInputChange('address.city', e.target.value)}
                      className="form-input"
                    />
                  ) : (
                    <span className="form-value">{profile.address.city}</span>
                  )}
                </div>

                <div className="form-group">
                  <label>Postal Code</label>
                  {isEditing ? (
                    <input
                      type="text"
                      value={editForm.address.code}
                      onChange={(e) => handleInputChange('address.code', e.target.value)}
                      className="form-input"
                    />
                  ) : (
                    <span className="form-value">{profile.address.code}</span>
                  )}
                </div>
              </div>

              <div className="form-group">
                <label>Country</label>
                {isEditing ? (
                  <input
                    type="text"
                    value={editForm.address.country}
                    onChange={(e) => handleInputChange('address.country', e.target.value)}
                    className="form-input"
                  />
                ) : (
                  <span className="form-value">{profile.address.country}</span>
                )}
              </div>
            </div>
          </div>

          {isEditing && (
            <div className="profile-actions">
              <button className="save-btn" onClick={handleSave}>
                <Save size={18} />
                Save Changes
              </button>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default ProfilePage

import React, { useState, useEffect } from 'react';
import api from '../api/api';
import './css/Profile.css';

function Profile() {
  const [profile, setProfile] = useState(null); // To store the fetched profile
  const [isEditing, setIsEditing] = useState(false); // To toggle update form visibility

  useEffect(() => {
    const fetchProfile = async () => {
      try {
        const response = await api.get('/profile');
        if (response.status === 200) {
          setProfile(response.data);
        } else {
          console.log('Error fetching profile');
        }
      } catch (error) {
        console.error('Error fetching profile:', error);
      }
    };

    fetchProfile();
  }, []);

  const handleChange = (e) => {
    setProfile({ ...profile, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await api.put('/profile', {
        ...profile,
        Skills: profile.Skills ? profile.Skills.split(',').map(skill => skill.trim()) : [],
      });
      alert('Profile updated successfully!');
      setIsEditing(false);
    } catch (error) {
      console.error('Error updating profile:', error);
      alert('Failed to update profile. Please try again.');
    }
  };

  return (
    <div className="profile-card">
      {!isEditing ? (
        profile ? (
          <div className="profile-details">
            <h2>{profile.FullName || 'Full Name'}</h2>
            <p><strong>Email:</strong> {profile.Email || 'N/A'}</p>
            <p><strong>Phone:</strong> {profile.PhoneNumber || 'N/A'}</p>
            <p><strong>Address:</strong> {profile.Address || 'N/A'}</p>
            <p><strong>Skills:</strong> {profile.Skills && profile.Skills.length > 0 ? profile.Skills.join(', ') : 'N/A'}</p>
            <p><strong>Course:</strong> {profile.Course || 'N/A'}</p>
            <p><strong>Specialization:</strong> {profile.Specialization || 'N/A'}</p>
            <button onClick={() => setIsEditing(true)}>Update Profile</button>
          </div>
        ) : (
          <p>Loading...</p>
        )
      ) : (
        <form onSubmit={handleSubmit} className="profile-form">
          <input
            type="text"
            name="FullName"
            value={profile?.FullName || ''}
            onChange={handleChange}
            placeholder="Full Name"
            required
          />
          <input
            type="tel"
            name="PhoneNumber"
            value={profile?.PhoneNumber || ''}
            onChange={handleChange}
            placeholder="Phone Number"
            required
          />
          <input
            type="email"
            name="Email"
            value={profile?.Email || ''}
            onChange={handleChange}
            placeholder="Email"
            required
          />
          <input
            type="text"
            name="Address"
            value={profile?.Address || ''}
            onChange={handleChange}
            placeholder="Address"
            required
          />
          <input
            type="text"
            name="Skills"
            value={profile?.Skills ? profile.Skills.join(', ') : ''}
            onChange={handleChange}
            placeholder="Skills (comma-separated)"
            required
          />
          <input
            type="text"
            name="Course"
            value={profile?.Course || ''}
            onChange={handleChange}
            placeholder="Course"
            required
          />
          <input
            type="text"
            name="Specialization"
            value={profile?.Specialization || ''}
            onChange={handleChange}
            placeholder="Specialization"
            required
          />
          <button type="submit">Save</button>
          <button type="button" onClick={() => setIsEditing(false)}>Cancel</button>
        </form>
      )}
    </div>
  );
}

export default Profile;

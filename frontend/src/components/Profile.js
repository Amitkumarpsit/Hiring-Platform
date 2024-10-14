import React, { useState } from 'react';
import { updateProfile } from '../api/api';

function Profile() {
  const [profile, setProfile] = useState({
    fullName: '',
    phoneNumber: '',
    email: '',
    address: '',
    skills: '',
    course: '',
    specialization: '',
  });

  const handleChange = (e) => {
    setProfile({ ...profile, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await updateProfile({
        ...profile,
        skills: profile.skills.split(',').map(skill => skill.trim()),
      });
      alert('Profile updated successfully!');
    } catch (error) {
      console.error('Error updating profile:', error);
      alert('Failed to update profile. Please try again.');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Update Profile</h2>
      <input
        type="text"
        name="fullName"
        value={profile.fullName}
        onChange={handleChange}
        placeholder="Full Name"
        required
      />
      <input
        type="tel"
        name="phoneNumber"
        value={profile.phoneNumber}
        onChange={handleChange}
        placeholder="Phone Number"
        required
      />
      <input
        type="email"
        name="email"
        value={profile.email}
        onChange={handleChange}
        placeholder="Email"
        required
      />
      <input
        type="text"
        name="address"
        value={profile.address}
        onChange={handleChange}
        placeholder="Address"
        required
      />
      <input
        type="text"
        name="skills"
        value={profile.skills}
        onChange={handleChange}
        placeholder="Skills (comma-separated)"
        required
      />
      <input
        type="text"
        name="course"
        value={profile.course}
        onChange={handleChange}
        placeholder="Course"
        required
      />
      <input
        type="text"
        name="specialization"
        value={profile.specialization}
        onChange={handleChange}
        placeholder="Specialization"
        required
      />
      <button type="submit">Update Profile</button>
    </form>
  );
}

export default Profile;
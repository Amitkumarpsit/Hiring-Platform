import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { register } from '../api/api';
import './css/RegisterUser.css'; // Assuming you have CSS styles for this

function RegisterUser() {
  const [userData, setUserData] = useState({
    fullName: '',
    email: '',
    phoneNumber: '',
    password: '',
    address: '',
    skills: '',
    course: '',
    specialization: '',
  });
  
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const navigate = useNavigate();

  const handleChange = (e) => {
    setUserData({ ...userData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      // Splitting the skills into an array of individual skills
      const userToRegister = {
        ...userData,
        skills: userData.skills.split(',').map(skill => skill.trim()),
      };

      // Register the user by calling the API
      const response = await register(userToRegister);

      // Show success message
      setSuccess(response.data.message);
      setError('');
    } catch (error) {
      setError('Registration failed. Please try again.');
      setSuccess('');
      console.error('Registration error:', error);
    }
  };

  return (
    <div className="register-container">
      <h2>Register</h2>
      {success && <div className="success-message">{success}</div>}
      {error && <div className="error-message">{error}</div>}
      <form onSubmit={handleSubmit} className="register-form">
        <input
          type="text"
          name="fullName"
          value={userData.fullName}
          onChange={handleChange}
          placeholder="Full Name"
          required
        />
        <input
          type="email"
          name="email"
          value={userData.email}
          onChange={handleChange}
          placeholder="Email"
          required
        />
        <input
          type="tel"
          name="phoneNumber"
          value={userData.phoneNumber}
          onChange={handleChange}
          placeholder="Phone Number"
          required
        />
        <input
          type="password"
          name="password"
          value={userData.password}
          onChange={handleChange}
          placeholder="Password"
          required
        />
        <input
          type="text"
          name="address"
          value={userData.address}
          onChange={handleChange}
          placeholder="Address"
        />
        <input
          type="text"
          name="skills"
          value={userData.skills}
          onChange={handleChange}
          placeholder="Skills (comma-separated)"
        />
        <input
          type="text"
          name="course"
          value={userData.course}
          onChange={handleChange}
          placeholder="Course"
        />
        <input
          type="text"
          name="specialization"
          value={userData.specialization}
          onChange={handleChange}
          placeholder="Specialization"
        />
        <button type="submit">Register</button>
      </form>
      <div className="login-link">
        <Link to="/login">Already have an account? Login here.</Link>
      </div>
    </div>
  );
}

export default RegisterUser;

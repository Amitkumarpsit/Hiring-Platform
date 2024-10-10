import React, { useState } from 'react';
import { postCandidate } from '../api/api';

function Profile() {
  const [candidate, setCandidate] = useState({ name: '', skills: [], resume: '' });

  const handleChange = (e) => {
    if (e.target.name === 'skills') {
      setCandidate({ ...candidate, skills: e.target.value.split(',') });
    } else {
      setCandidate({ ...candidate, [e.target.name]: e.target.value });
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await postCandidate(candidate);
      alert('Profile updated successfully!');
      setCandidate({ name: '', skills: [], resume: '' });
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
        name="name"
        value={candidate.name}
        onChange={handleChange}
        placeholder="Full Name"
        required
      />
      <input
        type="text"
        name="skills"
        value={candidate.skills.join(',')}
        onChange={handleChange}
        placeholder="Skills (comma-separated)"
        required
      />
      <textarea
        name="resume"
        value={candidate.resume}
        onChange={handleChange}
        placeholder="Paste your resume here"
        required
      />
      <button type="submit">Update Profile</button>
    </form>
  );
}

export default Profile;

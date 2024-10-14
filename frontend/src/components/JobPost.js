import React, { useState } from 'react';
import { postJob } from '../api/api';
import './JobPost.css';

function JobPost() {
  const [job, setJob] = useState({
    title: '',
    company: '',
    responsibilities: '',
    qualifications: '',
    location: '',
    category: ''
  });

  const handleChange = (e) => {
    setJob({ ...job, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await postJob(job);
      alert('Job posted successfully!');
      setJob({ title: '', company: '', responsibilities: '', qualifications: '', location: '', category: '' });
    } catch (error) {
      console.error('Error posting job:', error);
      alert('Failed to post job. Please try again.');
    }
  };

  return (
    <div className="job-post-container">
      <h2>Post a Job</h2>
      <form onSubmit={handleSubmit} className="job-post-form">
        <div className="form-group">
          <label htmlFor="title">Job Title</label>
          <input
            type="text"
            id="title"
            name="title"
            value={job.title}
            onChange={handleChange}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="company">Company</label>
          <input
            type="text"
            id="company"
            name="company"
            value={job.company}
            onChange={handleChange}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="category">Category</label>
          <select
            id="category"
            name="category"
            value={job.category}
            onChange={handleChange}
            required
          >
            <option value="">Select a category</option>
            <option value="IT">IT</option>
            <option value="Sales">Sales</option>
            <option value="Marketing">Marketing</option>
            <option value="HR">HR</option>
            <option value="Engineering">Engineering</option>
            <option value="Other">Other</option>
          </select>
        </div>
        <div className="form-group">
          <label htmlFor="location">Location</label>
          <input
            type="text"
            id="location"
            name="location"
            value={job.location}
            onChange={handleChange}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="responsibilities">Responsibilities</label>
          <textarea
            id="responsibilities"
            name="responsibilities"
            value={job.responsibilities}
            onChange={handleChange}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="qualifications">Qualifications</label>
          <textarea
            id="qualifications"
            name="qualifications"
            value={job.qualifications}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit" className="submit-button">Post Job</button>
      </form>
    </div>
  );
}

export default JobPost;
import React, { useState } from 'react';
import { postJob } from '../api/api';

function JobPost() {
  const [job, setJob] = useState({
    title: '',
    company: '',
    responsibilities: '',
    qualifications: '',
    location: ''
  });

  const handleChange = (e) => {
    setJob({ ...job, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await postJob(job);
      alert('Job posted successfully!');
      setJob({ title: '', company: '', responsibilities: '', qualifications: '', location: '' });
    } catch (error) {
      console.error('Error posting job:', error);
      alert('Failed to post job. Please try again.');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Post a Job</h2>
      <input
        type="text"
        name="title"
        value={job.title}
        onChange={handleChange}
        placeholder="Job Title"
        required
      />
      <input
        type="text"
        name="company"
        value={job.company}
        onChange={handleChange}
        placeholder="Company"
        required
      />
      <textarea
        name="responsibilities"
        value={job.responsibilities}
        onChange={handleChange}
        placeholder="Responsibilities"
        required
      />
      <textarea
        name="qualifications"
        value={job.qualifications}
        onChange={handleChange}
        placeholder="Qualifications"
        required
      />
      <input
        type="text"
        name="location"
        value={job.location}
        onChange={handleChange}
        placeholder="Location"
        required
      />
      <button type="submit">Post Job</button>
    </form>
  );
}

export default JobPost;
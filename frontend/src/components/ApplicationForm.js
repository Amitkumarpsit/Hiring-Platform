import React, { useState } from 'react';
import { submitApplication } from '../api/api';

function ApplicationForm({ jobId, onClose }) {
  const [formData, setFormData] = useState({
    fullName: '',
    email: '',
    age: '',
    course: '',
    courseEndDate: '',
    address: '',
    phoneNumber: '',
  });

  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prevData => ({
      ...prevData,
      [name]: name === 'age' ? (value === '' ? '' : parseInt(value, 10)) : value
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    try {
      const applicationData = {
        ...formData,
        jobId,
      };
      console.log('Submitting application with data:', applicationData);
      const response = await submitApplication(applicationData);
      console.log('Application submission response:', response);
      setSuccess('Application submitted successfully!');
      setTimeout(() => {
        onClose();
      }, 2000);
    } catch (error) {
      console.error('Error submitting application:', error);
      if (error.response && error.response.status === 409) {
        setError('You have already applied for this job.');
      } else {
        setError(error.response?.data?.message || 'Failed to submit application. Please try again.');
      }
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input name="fullName" value={formData.fullName} onChange={handleChange} placeholder="Full Name" required />
      <input name="email" type="email" value={formData.email} onChange={handleChange} placeholder="Email" required />
      <input name="age" type="number" value={formData.age} onChange={handleChange} placeholder="Age" required />
      <input name="course" value={formData.course} onChange={handleChange} placeholder="Course" required />
      <input name="courseEndDate" type="date" value={formData.courseEndDate} onChange={handleChange} placeholder="Course End Date" required />
      <input name="address" value={formData.address} onChange={handleChange} placeholder="Address" required />
      <input name="phoneNumber" value={formData.phoneNumber} onChange={handleChange} placeholder="Phone Number" required />
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {success && <p style={{ color: 'green' }}>{success}</p>}
      <button type="submit">Submit Application</button>
    </form>
  );
}

export default ApplicationForm;
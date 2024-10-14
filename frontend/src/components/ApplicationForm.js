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

  const handleChange = (e) => {
    const value = e.target.name === 'age' ? parseInt(e.target.value) : e.target.value;
    setFormData({ ...formData, [e.target.name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const applicationData = {
        ...formData,
        jobId,
        age: parseInt(formData.age),
      };
      console.log('Submitting application with data:', applicationData);
      const response = await submitApplication(applicationData);
      console.log('Application submission response:', response);
      alert('Application submitted successfully!');
      onClose();
    } catch (error) {
      console.error('Error submitting application:', error.response || error);
      alert('Failed to submit application. Please try again.');
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
      <button type="submit">Submit Application</button>
    </form>
  );
}

export default ApplicationForm;
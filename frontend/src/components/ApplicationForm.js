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
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: name === 'age' ? parseInt(value) : value,  // Convert age to an integer
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const applicationData = {
        ...formData,
        jobId: jobId,  // Ensure this is the correct job ID format
        courseEndDate: `${formData.courseEndDate}T00:00:00Z`,  // Convert date to full timestamp format
        age: parseInt(formData.age),
      };
      const response = await submitApplication(applicationData);
      alert('Application submitted successfully!');
      onClose();
    } catch (error) {
      console.error('Failed to submit application:', error);
      alert('Failed to submit application. Check the console for more details.');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        name="fullName"
        value={formData.fullName}
        onChange={handleChange}
        placeholder="Full Name"
        required
      />
      <input
        name="email"
        type="email"
        value={formData.email}
        onChange={handleChange}
        placeholder="Email"
        required
      />
      <input
        name="age"
        type="number"
        value={formData.age}
        onChange={handleChange}
        placeholder="Age"
        required
      />
      <input
        name="course"
        value={formData.course}
        onChange={handleChange}
        placeholder="Course"
        required
      />
      <input
        name="courseEndDate"
        type="date"
        value={formData.courseEndDate}
        onChange={handleChange}
        placeholder="Course End Date"
        required
      />
      <input
        name="address"
        value={formData.address}
        onChange={handleChange}
        placeholder="Address"
        required
      />
      <input
        name="phoneNumber"
        value={formData.phoneNumber}
        onChange={handleChange}
        placeholder="Phone Number"
        required
      />
      <button type="submit">Submit Application</button>
    </form>
  );
}

export default ApplicationForm;

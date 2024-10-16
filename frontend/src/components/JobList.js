import React, { useState, useEffect } from 'react';
import { getJobs } from '../api/api';
import ApplicationForm from './ApplicationForm';

function JobList() {
  const [jobs, setJobs] = useState([]);
  const [selectedJob, setSelectedJob] = useState(null);

  useEffect(() => {
    const fetchJobs = async () => {
      try {
        const response = await getJobs();
        setJobs(response.data);
      } catch (error) {
        console.error('Error fetching jobs:', error);
      }
    };
    fetchJobs();
  }, []);

  const handleApply = (job) => {
    setSelectedJob(job);
  };

  const handleCloseForm = () => {
    setSelectedJob(null);
  };

  return (
    <div>
      <h2>Job Listings</h2>
      {jobs.length === 0 ? (
        <p>No jobs available at the moment</p>
      ) : (
        jobs.map((job) => (
          <div key={job.ID} style={{border: '1px solid #ddd', margin: '10px 0', padding: '10px'}}>
            <h3>{job.Title}</h3>
            <p><strong>Company:</strong> {job.Company}</p>
            <p><strong>Location:</strong> {job.Location}</p>
            <p><strong>Responsibilities:</strong> {job.Responsibilities}</p>
            <p><strong>Qualifications:</strong> {job.Qualifications}</p>
            <button onClick={() => handleApply(job)}>Apply</button>
          </div>
        ))
      )}
      {selectedJob && (
        <div>
          <h3>Apply for {selectedJob.Title}</h3>
          <ApplicationForm jobId={selectedJob.ID} onClose={handleCloseForm} />
        </div>
      )}
    </div>
  );
}

export default JobList;
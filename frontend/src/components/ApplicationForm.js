import axios from 'axios';

const API_URL = 'http://localhost:8000';

const api = axios.create({
  baseURL: API_URL,
});

export const getJobs = () => api.get('/jobs');
export const postJob = (job) => api.post('/jobs/new', job);
export const getCandidates = () => api.get('/candidates');
export const postCandidate = (candidate) => api.post('/candidates/new', candidate);
export const updateProfile = (profile) => api.put('/profile', profile);

// Updated submitApplication function
export const submitApplication = async (applicationData) => {
  try {
    const response = await api.post('/applications', applicationData);
    console.log('Application submission response:', response);
    return response;
  } catch (error) {
    console.error('Error submitting application:', error.response || error);
    throw error;
  }
};

export default api;

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

export default api;
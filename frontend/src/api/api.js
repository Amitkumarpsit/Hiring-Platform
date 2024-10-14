import axios from 'axios';

const API_URL = 'http://localhost:8000';

const api = axios.create({
  baseURL: API_URL,
});

// interceptor to include the token in all requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export const getJobs = () => api.get('/jobs');
export const postJob = (job) => api.post('/jobs/new', job);
export const getCandidates = () => api.get('/candidates');
export const postCandidate = (candidate) => api.post('/candidates/new', candidate);
export const updateProfile = (profile) => api.put('/profile', profile);
export const login = (credentials) => api.post('/login', credentials);
export const register = (userData) => api.post('/register', userData);

export const submitApplication = async (applicationData) => {
  try {
    const response = await api.post('/applications', applicationData);
    console.log('Application submission response:', response);
    return response;
  } catch (error) {
    console.error('Error submitting application:', error.response ? error.response.data : error.message);
    alert('Error: ' + (error.response ? error.response.data : error.message));
    throw error;
  }
};

export default api;

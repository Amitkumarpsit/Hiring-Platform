import axios from 'axios';

const API_URL = 'http://localhost:8000';

const api = axios.create({
  baseURL: API_URL,
});

// Interceptor to include the token in all requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Existing exports
export const getJobs = () => api.get('/jobs');
export const postJob = (job) => api.post('/jobs/new', job);
export const getCandidates = () => api.get('/candidates');
export const postCandidate = (candidate) => api.post('/candidates/new', candidate);
export const updateProfile = (profile) => api.put('/profile', profile);
export const login = (credentials) => api.post('/login', credentials);
export const register = (userData) => api.post('/register', userData);

// New password reset exports
export const forgotPassword = (email) => 
  api.post('/forgot-password', { email }, {
    headers: {
      'Content-Type': 'application/json',
    }
  });
export const resetPassword = (resetToken, newPassword) => 
  api.post('/reset-password', { resetToken, newPassword });

export const submitApplication = async (applicationData) => {
  try {
    const response = await api.post('/applications', applicationData);
    return response.data;
  } catch (error) {
    console.error('Error submitting application:', error.response || error);
    throw error;
  }
};

export default api;
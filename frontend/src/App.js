import React from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import Navbar from './components/Navbar';
import Dashboard from './components/Dashboard';
import JobList from './components/JobList';
import JobPost from './components/JobPost';
import Profile from './components/Profile';
import Login from './components/Login';
import ForgotPassword from './components/ForgotPassword';
import ResetPassword from './components/ResetPassword';
import RegisterUser from './components/RegisterUser';
import VerifyEmail from './components/VerifyEmail';  // Import VerifyEmail
import './App.css';

function PrivateRoute({ children }) {
  const isAuthenticated = !!localStorage.getItem('token');
  return isAuthenticated ? children : <Navigate to="/login" />;
}

function App() {
  return (
    <Router>
      <div className="app">
        <Navbar />
        <div className="content">
          <Routes>
            <Route path="/" element={<Dashboard />} />
            <Route path="/jobs" element={<PrivateRoute><JobList /></PrivateRoute>} />
            <Route path="/post-job" element={<PrivateRoute><JobPost /></PrivateRoute>} />
            <Route path="/profile" element={<PrivateRoute><Profile /></PrivateRoute>} />
            <Route path="/login" element={<Login />} />
            <Route path="/forgot-password" element={<ForgotPassword />} />
            <Route path="/reset-password" element={<ResetPassword />} />
            <Route path="/register" element={<RegisterUser />} />
            <Route path="/verify-email" element={<VerifyEmail />} />  {/* Add VerifyEmail route */}
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;

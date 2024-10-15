import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { FaSearch, FaUser, FaSignOutAlt } from 'react-icons/fa';
import './css/Navbar.css';

const Navbar = () => {
  const navigate = useNavigate();
  const isAuthenticated = !!localStorage.getItem('token');

  const handleLogout = () => {
    localStorage.removeItem('token');
    navigate('/login');
  };

  return (
    <nav className="navbar">
      <div className="navbar-content">
        <div className="navbar-left">
          <Link to="/" className="navbar-logo">
            <div className="logo-circle">B</div>
            <span className="logo-text">BROCODES.CLUB</span>
          </Link>
        </div>
        <div className="navbar-center">
          <Link to="/" className="nav-link">HOME</Link>
          {isAuthenticated && (
            <>
              <Link to="/jobs" className="nav-link">GET JOB</Link> {/* This will only show after login */}
              <Link to="/post-job" className="nav-link">POST JOBS</Link>
              <Link to="/profile" className="nav-link">PROFILE</Link>
            </>
          )}
        </div>
        <div className="navbar-right">
          <div className="search-bar">
            <FaSearch className="search-icon" />
            <input type="text" placeholder=" " className="search-input" />
          </div>
          {isAuthenticated ? (
            <div className="user-actions">
              <FaUser className="user-icon" />
              <FaSignOutAlt className="logout-icon" onClick={handleLogout} title="Logout" />
            </div>
          ) : (
            <Link to="/login" className="login-link">
              <FaUser className="user-icon" />
              <span>Login</span>
            </Link>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
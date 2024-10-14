import React from 'react';
import { Link } from 'react-router-dom';
import { FaSearch, FaUser } from 'react-icons/fa';
import './Navbar.css';

const Navbar = () => {
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
          <Link to="/jobs" className="nav-link">GET JOB</Link>
          <Link to="/post-job" className="nav-link">POST JOBS</Link>
          <Link to="/profile" className="nav-link">PROFILE</Link>
        </div>
        <div className="navbar-right">
          <div className="search-bar">
            <FaSearch className="search-icon" />
            <input type="text" placeholder=" " className="search-input" />
          </div>
          <FaUser className="user-icon" />
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
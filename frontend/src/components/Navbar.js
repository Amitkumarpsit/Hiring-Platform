import React from 'react';
import { FaSearch, FaPlus, FaInstagram, FaUser, FaBars } from 'react-icons/fa';
import './Navbar.css'; // Importing the CSS file

const Navbar = () => {
  return (
    <nav className="navbar">
      {/* Logo Section */}
      <div className="navbar-logo">
        <h1 className="logo-text">Brocodes.Club</h1>
      </div>

      {/* Search Bar Section */}
      <div className="navbar-search">
        <FaSearch className="search-icon" />
        <input 
          type="text" 
          placeholder='Search "title card"' 
          className="search-input" 
        />
        <FaPlus className="plus-icon" />
      </div>

      {/* Navigation Links */}
      {/* <ul className="navbar-links">
        <li><Link to="/" className="nav-link">Dashboard</Link></li>
        <li><Link to="/jobs" className="nav-link">Job Listings</Link></li>
        <li><Link to="/post-job" className="nav-link">Post a Job</Link></li>
        <li><Link to="/profile" className="nav-link">Profile</Link></li>
      </ul> */}

      {/* Right Icons Section */}
      <div className="navbar-icons">
        <FaInstagram className="icon" />
        <FaUser className="icon" />
        <FaBars className="icon" />
      </div>
    </nav>
  );
};

export default Navbar;

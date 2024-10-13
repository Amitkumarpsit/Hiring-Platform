import React from 'react';
import { FaSearch, FaPlus, FaInstagram, FaUser, FaBars } from 'react-icons/fa';
import { Link } from 'react-router-dom';
import './Navbar.css';

const NavbarWithTabs = () => {
  return (
    <>
      {/* Original Navbar */}
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

        {/* Right Icons Section */}
        <div className="navbar-icons">
          <FaInstagram className="icon" />
          <FaUser className="icon" />
          <FaBars className="icon" />
        </div>
      </nav>

      {/* Tabs Bar Below the Existing Navbar with NO gap */}
      <nav className="nav-tabs">
        <ul className="navbar-links">
          <li><Link to="/" className="nav-link">Dashboard</Link></li>
          <li><Link to="/jobs" className="nav-link">Job Listings</Link></li>
          <li><Link to="/post-job" className="nav-link">Post a Job</Link></li>
          <li><Link to="/profile" className="nav-link">Profile</Link></li>
        </ul>
      </nav>

      {/* Tab Content */}
      <div className="tab-content">
        <div className="tab-pane active" id="nav-home">
          Dashboard content goes here...
        </div>
        <div className="tab-pane" id="nav-jobs">
          Job Listings content goes here...
        </div>
        <div className="tab-pane" id="nav-post-job">
          Post a Job content goes here...
        </div>
        <div className="tab-pane" id="nav-profile">
          Profile content goes here...
        </div>
      </div>
    </>
  );
};

export default NavbarWithTabs;

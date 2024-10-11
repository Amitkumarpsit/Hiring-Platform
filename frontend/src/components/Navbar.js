import React from 'react';
// import { Link } from 'react-router-dom';
import { FaSearch, FaPlus, FaInstagram, FaUser, FaBars } from 'react-icons/fa';

const Navbar = () => {
  return (
    <nav style={styles.navbar}>
      {/* Logo Section */}
      <div style={styles.navbarLogo}>
        <h1 style={styles.logoText}>Brocodes.Club</h1>
      </div>

      {/* Search Bar Section */}
      <div style={styles.navbarSearch}>
        <FaSearch style={styles.searchIcon} />
        <input 
          type="text" 
          placeholder='Search "title card"' 
          style={styles.searchInput} 
        />
        <FaPlus style={styles.plusIcon} />
      </div>

      {/* Navigation Links */}
      {/* <ul style={styles.navbarLinks}>
        <li><Link to="/" style={styles.navLink}>Dashboard</Link></li>
        <li><Link to="/jobs" style={styles.navLink}>Job Listings</Link></li>
        <li><Link to="/post-job" style={styles.navLink}>Post a Job</Link></li>
        <li><Link to="/profile" style={styles.navLink}>Profile</Link></li>
      </ul> */}

      {/* Right Icons Section */}
      <div style={styles.navbarIcons}>
        <FaInstagram style={styles.icon} />
        <FaUser style={styles.icon} />
        <FaBars style={styles.icon} />
      </div>
    </nav>
  );
};

// Styles as JavaScript object
const styles = {
  navbar: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: '10px 20px',
    backgroundColor: '#000', // Black background
  },
  navbarLogo: {
    display: 'flex',
    alignItems: 'center',
  },
  logoText: {
    color: '#FF1493', // Pink color for "B.club" text
    fontSize: '24px',
    fontWeight: 'bold',
  },
  navbarSearch: {
    display: 'flex',
    alignItems: 'center',
    border: '2px solid #FF1493', // Pink border
    padding: '5px 10px',
    borderRadius: '50px',
    backgroundColor: '#000', // Black background
  },
  searchIcon: {
    color: '#FF1493',
    marginRight: '10px',
  },
  searchInput: {
    border: 'none',
    backgroundColor: 'transparent',
    outline: 'none',
    color: '#fff',
    fontSize: '16px',
    width: '1000px',
  },
  plusIcon: {
    color: '#FF1493',
    marginLeft: '10px',
  },
  navbarLinks: {
    listStyle: 'none',
    display: 'flex',
    gap: '20px',
    marginRight: '20px',
  },
  navLink: {
    color: '#fff',
    textDecoration: 'none',
    fontSize: '16px',
  },
  navbarIcons: {
    display: 'flex',
    gap: '20px',
  },
  icon: {
    color: '#FF1493', // Pink color for icons
    fontSize: '20px',
    cursor: 'pointer',
  },
};

export default Navbar;

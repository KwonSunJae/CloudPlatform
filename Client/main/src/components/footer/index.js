// Footer.js
import React from 'react';
import './index.css';
const Footer = () => {
  return (
    <footer className="footer">
      <div className="footer-content">
        <p>© {new Date().getFullYear()} DMS Lab Cloud Platform</p>
      </div>
    </footer>
  );
};

export default Footer;
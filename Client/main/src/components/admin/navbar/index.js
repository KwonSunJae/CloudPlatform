import React from 'react';
import { Link } from 'react-router-dom'; // React Router를 사용하여 리다이렉션
import './index.css';

const AdminNavbar = () => {
  return (
    <nav className="admin-navbar">
      <h2 className="admin-navbar-title">
        <Link to="/admin">Admin Page</Link>
      </h2>
      <ul className="admin-navbar-list">
        <li className="admin-navbar-item"><Link to="/admin">유저</Link></li>
        <li className="admin-navbar-item"><Link to="/admin/vm">VM</Link></li>
        <li className="admin-navbar-item"><Link to="/admin/container">Container</Link></li>
        <li className="admin-navbar-item"><Link to="/admin/kluster">Kluster</Link></li>
        <li className="admin-navbar-item"><Link to="/admin/baremetal">Baremetal</Link></li>
      </ul>
    </nav>
  );
};

export default AdminNavbar;

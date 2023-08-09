import React from 'react';
import { Link } from 'react-router-dom'; // React Router를 사용하여 리다이렉션
import './index.css';
const NavigationBar = () => {
  return (
    <nav>
      <div className="logo">
        <h2>DMS Lab Cloud Platform</h2>
      </div>
      <ul>
        <li><Link to="/introduction">소개</Link></li>
        <li><Link to="/system">시스템</Link></li>
        <li><Link to="/monitoring">모니터링</Link></li>
        <li><Link to="/member-info">회원정보</Link></li>
      </ul>
      <div className="auth-buttons">
        <button><Link to="/login">로그인</Link></button>
        <button><Link to="/signup">회원가입</Link></button>
      </div>
    </nav>
  );
};

export default NavigationBar;
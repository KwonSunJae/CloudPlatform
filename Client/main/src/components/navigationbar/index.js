import React from 'react';
import { Link } from 'react-router-dom'; // React Router를 사용하여 리다이렉션
import { RegisterModal, LoginModal } from '../user'; // RegisterModal, LoginModal 컴포넌트를 import
import './index.css';
import { useEffect } from 'react';
const NavigationBar = () => {
  // NavigationBar 컴포넌트는 네비게이션 바를 렌더링하는 컴포넌트입니다.
  // 네비게이션 바는 다음과 같은 요소로 구성됩니다.
  // - 로고
  // - 네비게이션 메뉴(소개, 시스템, 모니터링, 회원정보)
  // - 로그인 버튼 , 클릭시 LoginModal 모달창 오픈
  // - 회원가입 버튼, 클릭시 RegisterModal 모달창 오픈
  useEffect(() => {
    const accsessToken = localStorage.getItem("accessToken");
    if (accsessToken) {
      document.getElementsByClassName("auth-buttons")[0].style.display = "none";
      document.getElementById("login-modal").style.display = "none";
      document.getElementsByClassName("logout-button")[0].style.display = "block";
    }
    else {
      document.getElementsByClassName("auth-buttons")[0].style.display = "block";
      document.getElementsByClassName("logout-button")[0].style.display = "none";
    }
  }, []);
  function openRegisterModal() {
    document.getElementById("register-modal").style.display = "block";
  }
  function openLoginModal() {
    document.getElementById("login-modal").style.display = "block";
  }
  function logout() {
    localStorage.removeItem("accessToken");
    console.log("로그아웃");
    localStorage.removeItem("refreshToken");
    window.location.href = "/";
  }

  return (
    
    <nav>
      <RegisterModal />
    <LoginModal />
      <div className="logo">
        <h2><Link to = "/" >DMS Lab Cloud Platform</Link></h2>
      </div>
      <ul>
        <li><Link to="/introduction">소개</Link></li>
        <li><Link to="/system">시스템</Link></li>
        <li><Link to="/monitoring">모니터링</Link></li>
        <li><Link to="/member-info">회원정보</Link></li>
      </ul>
      <div className="auth-buttons">
        <button onClick={openLoginModal}>로그인</button>
        <button onClick={openRegisterModal}>회원가입</button>
      </div>
      <div className="logout-button">
        <button onClick={logout}>로그아웃</button>
      </div>
      
    </nav>
  );
};

export default NavigationBar;
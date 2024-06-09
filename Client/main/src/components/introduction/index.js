import React from 'react';
import { Link } from 'react-router-dom';
import { Button } from 'react-bootstrap';
import './index.css'; // CSS 파일을 통해 스타일 적용

const ShowIntroduction = () => {
  return (
    <div className="introduction">
      {/* 이미지 추가 */}
      <div className="introduction-image">
        <img src="lab_image.jpg" alt="연구실 이미지" />
      </div>
      
      {/* 연구실 소개 내용 */}
      <div className="introduction-content">
        <h1>우리 연구실 소개</h1>
        <p>
          환영합니다! 우리 연구실은 최첨단 기술과 혁신적인 아이디어로 미래를 선도하는 연구를 진행하고 있습니다.
          다양한 분야에서 연구를 진행하며, 학문적 호기심과 창의적인 실험을 통해 새로운 지식의 창출을 목표로 합니다.
        </p>
        
        {/* 더 알아보기 버튼 */}
        <Link to="/about-us">
          <Button variant="primary">더 알아보기</Button>
        </Link>
      </div>
      
    </div>
  );
};

export default ShowIntroduction;

import React from 'react';
import { Link } from 'react-router-dom';
import { Button } from 'react-bootstrap';
import './index.css'; // CSS 파일을 통해 스타일 적용

const ShowSystem = () => {
  return (
    <div className="system">
      <div className="system-header">
        <h1>시스템 설계 및 구조</h1>
        <p>우리 시스템의 설계 및 구조를 소개합니다. 아래 이미지를 통해 전체 아키텍처를 확인할 수 있습니다.</p>
      </div>
      
      <div className="system-image">
        <img src="path/to/your/system-architecture-image.png" alt="System Architecture" />
      </div>
      
      <div className="system-content">
        <h2>시스템 아키텍처</h2>
        <p>
          본 시스템은 모듈화된 구조를 가지고 있어 확장성과 유지보수성이 뛰어납니다. 주요 구성 요소는 다음과 같습니다:
        </p>
        <ul>
          <li><strong>프론트엔드:</strong> React를 기반으로 사용자 친화적인 인터페이스를 제공합니다.</li>
          <li><strong>백엔드:</strong> Node.js와 Express를 사용하여 API 서버를 구축하였습니다.</li>
          <li><strong>데이터베이스:</strong> MongoDB를 사용하여 유연한 데이터 모델링을 지원합니다.</li>
          <li><strong>인프라:</strong> Docker와 Kubernetes를 통해 애플리케이션을 컨테이너화하고 오케스트레이션합니다.</li>
        </ul>
        
        <h2>기능 설계</h2>
        <p>
          시스템은 사용자 인증, 데이터 처리, 실시간 업데이트 등의 주요 기능을 포함하고 있습니다. 각 기능은 독립적으로 설계되어 있으며, 마이크로서비스 아키텍처를 통해 상호 작용합니다.
        </p>
        
        <Link to="/more-details">
          <Button variant="primary">자세히 보기</Button>
        </Link>
      </div>
    </div>
  );
};

export default ShowSystem;

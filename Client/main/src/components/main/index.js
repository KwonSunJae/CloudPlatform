import React from 'react';
import Slider from 'react-slick'; // 사진 슬라이드에 사용할 라이브러리 (react-slick 예시)
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';
import { Link } from 'react-router-dom';
import './index.css';
import Image1 from '../../assets/img/main.jpg';
import Image2 from '../../assets/img/meet.jpeg';
import Image3 from '../../assets/img/Image3.jpg';
import icon1 from '../../assets/icon/monitor.png';
import icon0 from '../../assets/icon/cube-3d.png';
import icon2 from '../../assets/icon/machine.png';
import icon3 from '../../assets/icon/servers.png';

const Main = () => {
  // 슬라이드로 보여줄 이미지들 배열
  const slideImages = [
    // 이미지 URL들을 넣으세요
    Image1,
    Image2,
    Image3,
  ];

  const settings = {
    dots: true,
    infinite: true,
    speed: 5000,
    slidesToShow: 2,
    slidesToScroll: 1,
    arrows: false, 
    autoplay: true,
    autoplaySpeed: 1000,
    cssEase: 'linear',
  };

  return (
    <div className="main-screen">
      <Slider {...settings}>
        {slideImages.map((image, index) => (
          <div key={index}>
            <img src={image} alt={`Slide ${index}`} />
          </div>
        ))}
      </Slider>
      {/* 아래에 클라우드 서비스 제공 항목을 추가하고, 각 항목에 맞는 리다이렉션 링크를 설정하세요 */}
      <div className="cloud-services">
        <div className="cloud-service-item">
          <img src={icon0} alt="Cluster" />
          <Link to="/cluster">Cluster</Link>
          <div className="cloud-service-description">
            클러스터 서비스는 여러 노드를 묶어 하나의 시스템처럼 작동하게 합니다.
            <Link to="/cluster" className="cloud-service-button">바로가기</Link>
          </div>
        </div>
        <div className="cloud-service-item">
          <img src={icon1} alt="Virtual Machine" />
          <Link to="/vm">Virtual Machine</Link>
          <div className="cloud-service-description">
            가상 머신 서비스는 물리적 서버를 가상화하여 여러 운영 체제를 실행할 수 있습니다.
            <Link to="/vm" className="cloud-service-button">바로가기</Link>
          </div>
        </div>
        <div className="cloud-service-item">
          <img src={icon3} alt="Physical Machine" />
          <Link to="/physical-machine">Physical Machine</Link>
          <div className="cloud-service-description">
            물리적 머신 서비스는 실제 하드웨어 서버를 임대하여 사용할 수 있는 서비스입니다.
            <Link to="/physical-machine" className="cloud-service-button">바로가기</Link>
          </div>
        </div>
        <div className="cloud-service-item">
          <img src={icon2} alt="Container" />
          <Link to="/container">Container</Link>
          <div className="cloud-service-description">
            컨테이너 서비스는 애플리케이션을 독립적으로 실행할 수 있는 경량 가상화 기술입니다.
            <Link to="/container" className="cloud-service-button">바로가기</Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Main;

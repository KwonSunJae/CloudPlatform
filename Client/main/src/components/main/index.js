import React from 'react';
import Slider from 'react-slick'; // 사진 슬라이드에 사용할 라이브러리 (react-slick 예시)
import { Link } from 'react-router-dom';
import './index.css';
import Image1 from '../../assets/img/main.jpg';
import Image2 from '../../assets/img/meet.jpeg';
import icon1 from '../../assets/icon/monitor.png';
import icon0 from '../../assets/icon/cube-3d.png';
import icon2 from '../../assets/icon/machine.png';
import icon3 from '../../assets/icon/servers.png';
const Main = () => {
  // 슬라이드로 보여줄 이미지들 배열
  const slideImages = [
    // 이미지 URL들을 넣으세요
    Image1
  ];

  const settings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
  };

  return (
    <div className="main-screen">
      <Slider {...settings}>
        {slideImages.map((image, index) => (
          <div key={index}>
            <img src={image} alt={`Slide ${index + 1}`} />
          </div>
        ))}
      </Slider>
      {/* 아래에 클라우드 서비스 제공 항목을 추가하고, 각 항목에 맞는 리다이렉션 링크를 설정하세요 */}
      <div className="cloud-services">
        <div>
          <img src={icon0} width='100px' height='100px'></img>
          <Link to="/cluster">Cluster</Link>
        </div>
        <div>
          <img src={icon1} width='100px' height='100px'></img>
          <Link to="/vm">Virtual Machine</Link>
            
        </div>
        <div>
        <img src={icon3} width='100px' height='100px'></img>
          <Link to="/physical-machine">Physical Machine</Link></div>
        <div>
        <img src={icon2} width='100px' height='100px'></img>
          <Link to="/container">Container</Link></div>
      </div>
    </div>
  );
};

export default Main;

import React from 'react';
import { Link } from 'react-router-dom';
import { Button } from 'react-bootstrap';
import './index.css'; // CSS 파일을 통해 스타일 적용

const ShowIntroduction = () => {
  return (
    <div className="introduction">
      <div className="col-xl-8 offset-xl-2 col-lg-10 offset-lg-1">
        <h2 id="introduction-to-dms-research-group">Introduction to DMS Research Group</h2>
        <p>The DMS research group is currently concentrating on research and development related to Smart Mobility Systems and Bio-Medical/Healthcare intelligent Services based on deep learning-based AI technologies on distributed systems</p>

        <h3 id="dstributed-system-technology">Dstributed System Technology</h3>
        <p>By utilizing Cloud Platform Technology, we are currently focusing on the development of "Intelligent Cloud Platform (IFC) for Smart Mobility systems and Healthcare Assistant systems, which is the core of the 4th Industrial Revolution. In particular, we are researching automatic AI platform provisiong service for "Intelligent Cooperative Heatlcare Assistant Agents for oldery people" and "Digital Twin Systems for Urban Air Mobility" application fields. We also research and develop stable and scalable system architecture and intelligent system operation management service.</p>

        <h3 id="artificial-intelligence-technology">Artificial Intelligence Technology</h3>
        <p>Various studies are in progress in relation to AI technology based on deep learning. 1) Intelligent and efficient path planning and control based on Deep Reinforcement Learning (DRL), 2) Distributed Control technology for Swarm Mobility using Distributed Multiagent DRL technology, 3) Human Pose and Activity recognition technology based on Visual Transformer, 4) Object Detection, Face and Emotion Recognition technology using computer vision and widely used deep learning vision models such as CNN, YoLo, ResNet, MobileNet, etc. 5) NLP Research on Knowledge-based Open-Domain Conversational QA and Dialog based on Transformer-based encoder/decoder languate models, 6) Research on QA dataset creation competing Question Generation and Question Answering language models, 7) Intelligent Chatbot technology supporting long-term conversation using Open Domain Knowledge etc.</p>

        <h3 id="core-research-focus">Core Research Focus</h3>
        <ul type="circle">
          <li>Intelligent Mobility Control Technologies based on Multi-Agent Deep Reinforcement Learning and Attentions between agents</li>
          <li>Human Pose, Activity, Face, Emotion Recognition models based on deep learning for computer vision</li>
          <li>Open-Domain Converstational QA and Dialog systems based on Graph Knowledge and Self-Attention models</li>
          <li>Multi-Modal Situation Detection and Context Regonition Servies for Healthcare Assistant</li>
          <li>Automatic provisioning and intelligent operation management technologies for intelligent cloud platforms (IaaS, PaaS, SaaS, FaaS)</li>
          <li>Hierarchical Distributed Intelligent Fog and Cloud Cooperative Service Platform technology based on Offloading/Caching technolog</li>
          <li>Distributed system quality evaluation technology based on stochastic reward net and discrete event simulation</li>
          <li>Efficient and reusable SW Architecture design technology suitable for domain function and quality requirements</li>
        </ul>

        <h3 id="contact">Contact</h3>

        <div className="language-plaintext highlighter-rouge">
          <div className="highlight">
            <pre className="highlight"><code>Office: Konkuk University New Engineering Building 1207
              <br/>
              Email: dkmin at konkuk.ac.kr (Prof.Dugki Min, Ph.D)
            </code></pre>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ShowIntroduction;

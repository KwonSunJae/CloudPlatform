import React from "react";
import {Container, Row, Col } from "react-bootstrap";
import Spinner from "react-bootstrap/Spinner";
import './SpinnerOurs.css'; // CSS 파일을 import합니다.

const SpinnerOurs = () => {
    return (
        <div className="spinner-container">
            <div className="spinner" role="status">
                
            </div>
        </div>
    );
}

export default SpinnerOurs;

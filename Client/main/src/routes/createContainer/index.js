import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from '../../components/navigationbar';
import Footer from "../../components/footer";
import ContainerCreatePage from "../../components/container/create";

export default function Container() {

    return (

        <div>
            <NavigationBar />
      
            <ContainerCreatePage/>
            <Footer/>
        </div>

    );
}
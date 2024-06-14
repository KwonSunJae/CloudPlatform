import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from "../../../components/navigationbar";
import Footer from "../../../components/footer";
import AdminContainer from "../../../components/admin/container";
import AdminNavbar from "../../../components/admin/navbar";



export default function AdminPageContainer() {
  
  return (
    
    <div>
    <NavigationBar />
    <AdminNavbar />
    <AdminContainer />
    <Footer />
    </div>
    
  );
}
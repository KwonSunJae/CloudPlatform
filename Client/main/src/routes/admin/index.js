import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import NavigationBar from "../../components/navigationbar";
import Footer from "../../components/footer";
import Admin from "../../components/admin/user";
import AdminNavbar from "../../components/admin/navbar";



export default function AdminPage() {
  
  return (
    
    <div>
    <NavigationBar />
    <AdminNavbar />
    <Admin />
    <Footer />
    </div>
    
  );
}
import React from "react";
import '../index.css';


// Importing all created components
import Navbar from "../components/Navbar";
import Footer from "../components/Footer";


// Pass the child props
export default function Layout({ children }) {
  return (
    <div>
      {/* Attaching all file components */}
      <Navbar />
      {children}
      <Footer/>
    </div>
  );
}
import React from "react";
import { Link } from "react-router-dom";
import "../styles/Navbar.css";
import Logo from "../assets/logo.png";

const Navbar = () => {
    return (
        <nav className="navbar">
            <Link to="/" className="logo">
                <img src={Logo} alt="Logo" />
                <span>Restaurant</span> 
            </Link>
            <div className="nav-links">
                <Link to="/">Home</Link>
                <Link to="/login">Login</Link>
                <Link to="/signup">Signup</Link>
            </div>
        </nav>
    );
};

export default Navbar;

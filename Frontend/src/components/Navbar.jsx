import React from "react";
import "./Navbar.css";

const Navbar = () => {
    return (
        <nav className="navbar">
            <h2 className="logo">99 Meal</h2>
            <div className="nav-links">
                <a href="/">Home</a>
                <a href="/login">Login</a>
                <a href="/signup">Signup</a>
            </div>
        </nav>
    );
};

export default Navbar;

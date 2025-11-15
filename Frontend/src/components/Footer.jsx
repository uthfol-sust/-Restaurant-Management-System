import React from "react";
import "../styles/Footer.css";
import { FaFacebookF, FaInstagram, FaTwitter, FaGooglePlusG, FaYoutube } from "react-icons/fa";

function Footer() {
    return (
        <footer className="footer">
            <div className="social-icons">
                <a href="#"><FaFacebookF /></a>
                <a href="#"><FaInstagram /></a>
                <a href="#"><FaTwitter /></a>
                <a href="#"><FaGooglePlusG /></a>
                <a href="#"><FaYoutube /></a>
            </div>
            <ul className="footer-menu">
                <li><a href="/">Home</a></li>
                <li><a href="/news">News</a></li>
                <li><a href="/about">About</a></li>
                <li><a href="/contact">Contact Us</a></li>
                <li><a href="/team">Our Team</a></li>
            </ul>
            <p className="footer-text">
                Copyright ©2025 | Designed by <b>Antu Kalower</b>
            </p>
        </footer>
    );
}

export default Footer;

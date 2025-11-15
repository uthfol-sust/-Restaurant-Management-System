import React, { useEffect, useState } from "react";
import '../styles/Home.css'
import "../styles/Auth.css";
import Navbar from "../components/Navbar";

const Home = () => {
    const fullText = "Welcome to 99 Meal – Your Digital Restaurant Solution!";
    const [displayText, setDisplayText] = useState("");

    useEffect(() => {
        let i = 0;

        const startTyping = () => {
            i = 0;
            setDisplayText("");

            const typingInterval = setInterval(() => {
                setDisplayText(fullText.substring(0, i));
                i++;

                if (i > fullText.length) {
                    clearInterval(typingInterval);
                    setTimeout(startTyping, 2000);
                }
            }, 60);
        };

        startTyping();
    }, []);

    return (
        <div className="App">
            <Navbar />
        <div className="home-page">
            <div className="overlay">
                <div className="hero-content">
                    <h1 className="hero-title">{displayText}</h1>
                </div>
            </div>
        </div>
        </div>
    );
};

export default Home;
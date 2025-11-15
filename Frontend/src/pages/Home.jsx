
import React, { useEffect, useState } from "react";
import '../styles/Home.css'

const Home = () => {
    const fullText = "Welcome to 99 Meal – Your Digital Restaurant Solution!";
    const [displayText, setDisplayText] = useState("");

    useEffect(() => {
        let i = 0;
        const interval = setInterval(() => {
            setDisplayText(fullText.substring(0, i));
            i++;
            if (i > fullText.length) clearInterval(interval);
        }, 60);
    }, []);

    return (
        <div className="home-page">
            <div className="overlay">
                <div className="hero-content">
                    <span className="brand">99 MEAL</span>

                    <h1 className="hero-title">{displayText}</h1>

                    <p className="hero-subtitle">
                        Fast, automated, and modern restaurant management for real-time
                        orders, kitchen updates, and admin control.
                    </p>

                    <div className="home-buttons">
                        <button onClick={() => (window.location.href = "/login")}>
                            Login
                        </button>
                        <button onClick={() => (window.location.href = "/signup")}>
                            Sign Up
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Home;

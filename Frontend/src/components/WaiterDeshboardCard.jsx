import React from "react";
import "../styles/WaiterDashboard.css";

const Card = ({ title, count, icon }) => {
    return (
        <div className="card">
            <div className="card-icon">{icon}</div>
            <div className="card-content">
                <h3>{title}</h3>
                <p>{count}</p>
            </div>
        </div>
    );
};

export default Card;

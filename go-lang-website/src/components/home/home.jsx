import React from 'react';
import './Home.css';
import Banner from '../../assets/banner.jpg'; // Adjust based on your assets structure

const Home = () => {
    return (
        <div className="home">
            <header style={{ backgroundImage: `url(${Banner})` }}>
                <h1>Hello, I'm Alec</h1>
                <p>Welcome to my personal website!</p>
            </header>
        </div>
    );
}

export default Home;

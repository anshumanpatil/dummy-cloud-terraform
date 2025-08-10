import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import styles from './Login.module.css';
import axios from "axios";
import { API_URL } from "../../constants";
const Login: React.FC = () => {
    const [email, setEmail] = useState("demo@example.com");
    const [password, setPassword] = useState("password");
    const navigate = useNavigate();
    const login = async () => {
        try {
            // Simulate a login request
            const response = await axios.post(`${API_URL}/login`, { email, password });
            if (response.status === 200) {
                navigate('/home'); // Redirect to dashboard on successful login
                return;
            } else {
                console.error('Login failed:', response.data);
                alert('Login failed. Please check your credentials.');
                return;
            }
        } catch (error) {
            console.error('Error during login:', error);
            alert('An error occurred during login. Please try again later.');
        }
    };
    // useEffect(() => {
    //     // This effect runs once when the component mounts
    //     console.log('Login component mounted');
    // }, []);

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = event.target;
        switch (name) {
            case 'email':
                setEmail(value);
                break;
            case 'password':
                setPassword(value);
                break;
            default:
                break;
        }
    };

    const handleLogin = (event: React.FormEvent) => {
        event.preventDefault();
        if (!email || !password) {
            alert('Please fill in both fields');
            return;
        }
        if (!/\S+@\S+\.\S+/.test(email)) {
            alert('Please enter a valid email address');
            return;
        }
        if (password.length < 6) {
            alert('Password must be at least 6 characters long');
            return;
        }

        login();
        console.log('Login form submitted', { email, password });
    };
    return (
        <div className={styles['login-bg']}>
            <div className={styles['login-container']}>
                <h2 className={styles['login-title']}>Login</h2>
                <form className={styles['login-form']} onSubmit={handleLogin}>
                    <label htmlFor="email" className={styles['login-label']}>Email</label>
                    <input type="email" value={email} id="email" name="email" placeholder="Enter your email" className={styles['login-input']} onChange={handleChange} />
                    <label htmlFor="password" className={styles['login-label']}>Password</label>
                    <input type="password" value={password} id="password" name="password" placeholder="Enter your password" className={styles['login-input']} onChange={handleChange} />
                    <button type="submit" className={styles['login-button']}>Login</button>
                </form>
            </div>
        </div>
    );
};

export default Login;
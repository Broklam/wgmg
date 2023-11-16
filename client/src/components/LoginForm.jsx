import React, { useState } from 'react';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const LoginForm = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });

      if (response.ok) {
        toast.success('User logged in successfully!');
      } else {
        toast.error('Login failed. Please try again.');
      }
    } catch (error) {
      console.error('An error occurred during login:', error);
      toast.error('An error occurred during login. Please try again.');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-[#ffc017]">
      <div className="bg-black p-8 rounded-lg shadow-md w-96">
        <h2 className="text-2xl font-semibold text-center text-ffc017 mb-4 text-white">
          User Login
        </h2>
        <form onSubmit={handleLogin}>
          <div className="mb-4">
            <label htmlFor="username" className="block text-white">
              Username
            </label>
            <input
              type="text"
              id="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:ring-ffc017"
              placeholder="Enter your username"
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="password" className="block text-white">
              Password
            </label>
            <input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:ring-ffc017"
              placeholder="Enter your password"
              required
            />
          </div>
          <button
            type="submit"
            className="w-full  text-white hover:text-black py-2 rounded-md bg-[#ffc117] focus:outline-none focus:ring focus:ring-yellow-300"
          >
            Login
          </button>
        </form>
      </div>
      <ToastContainer /> 
    </div>
  );
};

export default LoginForm;

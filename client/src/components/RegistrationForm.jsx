
import React, { useState } from 'react';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const RegistrationForm = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const role = 'user'; 

  const handleRegistration = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/api/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password, role }),
      });
      if (response.ok) {
        toast.success('User registered successfully!'); 
      } else {
        toast.error('Registration failed. Please try again.'); 
      }
    } catch (error) {
      console.error('An error occurred during registration:', error);
      toast.error('An error occurred during registration. Please try again.'); 
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-black">
      <div className="bg-white p-8 rounded-lg shadow-md w-96">
        <h2 className="text-2xl font-semibold text-center text-ffc017 mb-4">
          User Registration
        </h2>
        <form onSubmit={handleRegistration}>
          <div className="mb-4">
            <label htmlFor="username" className="block text-gray-600">
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
            <label htmlFor="password" className="block text-gray-600">
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
            className="w-full bg-black text-[#ffc017] hover:text-black py-2 rounded-md hover:bg-yellow-500 focus:outline-none focus:ring focus:ring-yellow-300"
          >
            Register
          </button>
        </form>
      </div>
      <ToastContainer /> 
    </div>
  );
};

export default RegistrationForm;

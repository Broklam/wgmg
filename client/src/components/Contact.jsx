import React, { useState, useEffect } from 'react';
import { motion } from 'framer-motion';



const Contact = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('');
  const [submitted, setSubmitted] = useState(false);

  useEffect(() => {
    if (submitted) {
      const timer = setTimeout(() => {
        setSubmitted(false);
      }, 5000);
      return () => clearTimeout(timer);
    }
  }, [submitted]);

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(`Name: ${name}, Email: ${email}, Message: ${message}`);
    setName('');
    setEmail('');
    setMessage('');
    setSubmitted(true);
  };

  return (
    <div className="w-full mx-auto bg-[#ffc017] text-black">
      <h2 className="text-2xl font-bold text-center py-2">If you wish to contact me:</h2>
      <form onSubmit={handleSubmit} className="shadow-md rounded px-8 pt-6">
        <div className="mb-4">
          <label className="block text-sm font-bold mb-2" htmlFor="name">
            Name
          </label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline"
            id="name"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>
        <div className="mb-4">
          <label className="block text-sm font-bold mb-2" htmlFor="email">
            Email
          </label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline"
            id="email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <div className="mb-6">
          <label className="block text-sm font-bold mb-2" htmlFor="message">
            Message
          </label>
          <textarea
            className="shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline"
            id="message"
            value={message}
            onChange={(e) => setMessage(e.target.value)}
          />
        </div>
        <div className="flex items-center justify-between">
          <button
            className="bg-black hover:bg-gray-500 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline mb-2"
            type="submit"
          >
            Send
          </button>
          <a className="inline-block align-baseline font-bold text-sm text-black hover:text-gray-500" href="mailto:ivnfrv@yahoo.com">
            Or email me directly
          </a>
        </div>
        {submitted && (
  <motion.div 
    className="p-2 mt-2"
    initial={{ opacity: 1, scale: 1 }}
    animate={{ 
      y: ["20px", "-10px", "20px", "-10px", "10px", "2000px"],
      opacity: [1, 1, 0.5, 0.3, 0.2, 0]
    }}
    transition={{ 
      y: { times: [0, 0.2, 0.4, 0.6, 0.8, 1], duration: 2, ease: "easeInOut" },
      opacity: { delay: 1.5 }
    }}
  >
    <p className="text-black">Your message has been submitted!</p>
  </motion.div>
)}




      </form>
    </div>
  );
};

export default Contact;

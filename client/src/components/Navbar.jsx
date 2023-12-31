import React, { useState } from 'react';
import { AiOutlineClose, AiOutlineMenu } from 'react-icons/ai';
import {Link} from 'react-router-dom';

const Navbar = () => {
  const [nav, setNav] = useState(false);

  const handleNav = () => {
    setNav(!nav);
  };

  return (
    <div className='flex justify-between items-center h-12  mx-auto px-4 text-white bg-black'>
      <h1 className='w-full text-3xl font-bold'>WG</h1>
      <ul className='hidden md:flex'>
        <li className='p-4'>
        <Link to="/" className=''>Home</Link>
        </li>
        <li className='p-4 '>
            <Link to="/SignUp" className=''>Registration</Link>
        </li>
        <li className='p-4'>
        <Link to="/login" className=''>Login</Link>
        </li>
        
      </ul>
      <div onClick={handleNav} className='block md:hidden'>
          {nav ? <AiOutlineClose size={20}/> : <AiOutlineMenu size={20} />}
      </div>
      <ul className={nav ? 'fixed left-0 top-12 w-[60%] max-h-[30%] border-r border-r-gray-900 bg-[#000300] ease-in-out duration-500' : 'ease-in-out duration-500 fixed left-[-100%]'}>
        
          <li className='p-4 border-b border-gray-600'>
          <Link to="/" className=''>Home</Link>
          </li>
          <li className='p-4 border-b border-gray-600'>
          <Link to="/SignUp" className=''>Registration</Link>
          </li>
          <li className='p-4 border-b border-gray-600'>
          <Link to="/Login" className=''>Login</Link>
          </li>
         
      </ul>
    </div>
  );
};

export default Navbar;
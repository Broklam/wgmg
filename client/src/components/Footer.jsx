import React from 'react';

const Footer = () => {
  return (
    <div className='flex flex-col items-center justify-center bg-black text-gray-300 py-6'>
      <div className='grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-2 gap-8 w-3/4 text-center'>
        <div>
          <h6 className='font-medium text-gray-400'>Contact</h6>
          <ul className='list-none'>
            <li className='py-2 text-sm'>Email</li>
            <li className='py-2 text-sm'>
              <a href="http://www.telegram.com" className='text-blue-500 hover:text-blue-700 transition duration-300'>Telegram</a>
            </li>
          </ul>
        </div>
        <div>
          <h6 className='font-medium text-gray-400'>Git</h6>
          <ul className='list-none'>
            <li className='py-2 text-sm'>
              <a href="http://www.github.com/Broklam/wgm" className='text-blue-500 hover:text-blue-700 transition duration-300'>Source Code</a>
            </li>
          </ul>
        </div>
      </div>
      <p className='text-gray-500 mt-6 text-sm'> Ivan Forov. PET project.</p>
    </div>
  );
};

export default Footer;

import React from 'react';

const Footer = () => {
  return (
    <div className='min-h-screen flex flex-col'>
      <div className='flex-grow'>
        {/* Your main content here */}
      </div>

      <div className='w-full mx-auto py-3 px-4 grid sm:grid-cols-2 sm:flex justify-normal lg:grid-cols-2 gap-3 text-gray-300 bg-black'>
        <div>
          <h6 className='font-medium text-gray-400 px-2'>Contact</h6>
          <ul>
            <li className='py-2 text-sm px-2'>Email</li>
            <li className='py-2 text-sm px-2'>
            <a href="http://www.telegram.com">Telegram</a>
            </li>
          </ul>
        </div>
        <div>
          <h6 className='font-medium text-gray-400 px-2'>Git</h6>
          <ul>
            <li className='py-2 px-2 text-sm'>
            <a href="http://www.github.com">Source Code</a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default Footer;



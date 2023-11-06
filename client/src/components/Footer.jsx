import React from 'react';

const Footer = () => {
  return (
    <div className='min-h-screen flex flex-col'>
      <div className='flex-grow'>
        {/* Your main content here */}
      </div>

      <div className='min-w-[1240px] mx-auto py-16 px-4 grid sm:grid-cols-2 lg:grid-cols-2 gap-3 text-gray-300 bg-black'>
        <div>
          <h6 className='font-medium text-gray-400'>Contact</h6>
          <ul>
            <li className='py-2 text-sm'>Email</li>
            <li className='py-2 text-sm'>Telegram</li>
          </ul>
        </div>
        <div>
          <h6 className='font-medium text-gray-400'>Git</h6>
          <ul>
            <li className='py-2 text-sm'>Source Code</li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default Footer;



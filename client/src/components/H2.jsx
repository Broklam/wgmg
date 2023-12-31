import React from 'react';


const H2 = () => {
  const data = [
    { text: '1.Create Account', img: './createAccount.jpg' },
    { text: '2.Set up your schedule', img: './setUp.jpg' },
    { text: '3.Invite Friends!', img: './friends.jpg' },
    { text: '4.Rock that ', img: './rock.jpg' },
  ];

  return (
    <div className="grid lg:grid-cols-4 md:grid-cols-2 sm:grid-cols-1 gap-4 bg-black text-[#ffc017] py-2">
      {data.map((item, index) => (
        <div key={index} className="flex flex-col items-center">
          <img src={item.img} alt={item.text} className="w-full h-64 object-cover transform hover:scale-105 transition-transform duration-200" />
          <p className="mt-4 text-center text-xl">{item.text}</p>
        </div>
      ))}
    </div>
  );
};

export default H2;

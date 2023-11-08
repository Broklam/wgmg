import React from 'react'

const Hero = () => {
    return (
      <div className="flex justify-between items-center bg-[#ffc017] py-10 lg:py-0 border-y border-black xl:border-hidden xl:rounded-xl">
        {/* Component content */}
        <div className="px-10 space-y-5 lg:py-6">
          <h1 className="text-6xl md:text-7xl max-w-xl font-serif w-11/12 sm:w-9/12">
            <span className="underline decoration-black decoration-4">
              WG
            </span>{" "}
            is a place to plan your chores 
          </h1>
          <h2 className="w-9/12 font-normal">
            It's easy to plan and follow chores with your neighbours.
          </h2>
          <button className="border border-black bg-white px-4 py-2 rounded-full font-medium active:scale-90 transition duration-100">
            How?    
          </button>
        </div>
        {/*Image on the right*/}
        <img
          className="hidden sm:inline-flex h-40 lg:h-80 xl:h-full"
          src="./logo.png"
          alt=""
        />
      </div>
    );
  };
  export default Hero;
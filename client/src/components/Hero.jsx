import React from 'react';
import { Link } from 'react-scroll';

const Hero = () => {
  return (
    <div className="flex justify-between items-center bg-[#ffc017] py-10 lg:py-0 border-y border-black xl:border-hidden ">
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
        <Link
          activeClass="active"
          to="H2"
          spy={true}
          smooth={true}
          offset={250}
          duration={500}
        >
          <button className="border border-black bg-white px-4 py-2 rounded-full font-medium active:scale-90 transition duration-100 animate-fade-up animate-once animate-duration-[500ms] animate-delay-100 animate-ease-linear">
            How?
          </button>
        </Link>
      </div>
      {/*Image on the right*/}
      <img
        className="hidden sm:inline-flex h-40 lg:h-80 xl:h-full"
        src="./logo.png"
        alt=""
      />
      {/* H2 component */}
      <div name="H2" className="my-16">
        {/* H2 component content... */}
      </div>
    </div>
  );
};

export default Hero;

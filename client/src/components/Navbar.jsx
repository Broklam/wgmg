import React from 'react';
import {Link} from 'react-router-dom';


const Navbar=()=>{
    return (    
                <div className='flex justify-between items-center h-24 max-w-[1240px] mx-auto px-4 text-white'>
                
                    <ul className="flex justify-center space-x-4">
                        
                       <div>
                        <Link to="/" className='text-white text-xs '>Home</Link>
                    </div>
                    <div>
                        <Link to="/about" className='text-white text-xs '>Login</Link>
                </div>
                    
                    </ul>
                   
                </div>
    )

}

export default Navbar;
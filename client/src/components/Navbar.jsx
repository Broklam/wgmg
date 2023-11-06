import React from 'react';
import {Link} from 'react-router-dom';

const Navbar=()=>{
    return (    
                <div className='navBar'>
                    <div>
                        <Link to="/">Home</Link>
                    </div>
                    <div>
                        <Link to="/about">about</Link>
                </div>
                    <div >
                        <Link to="/A">A</Link>
                    </div>
                </div>
    )

}

export default Navbar;
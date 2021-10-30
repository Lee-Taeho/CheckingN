import React from 'react';
import { Link } from 'react-router-dom';

const CardItem = props => {
    return (
         <div className="card">
          <Link className='cards_item_link' to={props.path}>
               <img src={props.img} className="card_image" alt={props.title}/>
              <h5 className="card_title">{props.title}</h5>
          </Link>
         </div>
    );
  }
  
  export default CardItem;
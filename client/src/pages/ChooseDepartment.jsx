import React  from 'react';
import DepartmentList from '../components/DepartmentList';
import './Background.css';

const ChooseDepartment = () => {
    return (
        <div className="form-signin">
        <div className='department_container'>
            <DepartmentList/>
        </div>
        </div>
        
    );
};

export default ChooseDepartment;
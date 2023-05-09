import { useState } from 'react';
import Appointment from '../../model/Appointment';

function RegisterPrice() {

    const [appointment, setAppointment] = useState<Appointment>({ id: '', accomodationId: '', start : '', end : '', priceType : '', price : 0, status: 'Free' })
    
    const createAppointment = () => {
        console.log(appointment)
        
      }
    
    return <>
        <div className="form">
            <h2 className='heading'>New accomodation</h2>
            <div className='form-fields'>
            <div className='field'>
                <label>
                Start:
                <input type="date" name="name"
                 onChange={(e) => setAppointment(prevState => ({ ...prevState, start: e.target.value }))}/>
                </label>
            </div>
            <div className='field'>
                <label>
                End:
                <input type="date" name="name" 
                onChange={(e) => setAppointment(prevState => ({ ...prevState, end: e.target.value }))}/>
                </label>
            </div>
            <div className='field'>
                <label>
                Price type:
                <select onChange={(e) => setAppointment(prevState => ({ ...prevState, priceType: e.target.value }))}>
                    <option value="PerPerson">Per person</option>
                    <option value="PerUnit">Per unit</option>
                </select>
                </label>
            </div>
            <div className='field'>
                <label>
                Price:
                <input type="number" name="name" 
                onChange={(e) => setAppointment(prevState => ({ ...prevState, price: parseInt(e.target.value) }))}/>
                </label>
            </div>
            <div className='field'>
            <button onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
            e.stopPropagation()
            createAppointment()
          }}
          >Create
          </button>
            </div>
            </div>
            
            
        </div>
    </>;
  }
export default RegisterPrice;
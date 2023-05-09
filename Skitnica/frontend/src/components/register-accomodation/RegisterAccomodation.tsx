import { useState } from 'react';
import './RegisterAccomodation.css';
import Accomodation from '../../model/Accomodation';
import accomodationService from '../../services/accomodation.service';

function RegisterAccomodation() {

    const [accomodation, setAccomodation] = useState<Accomodation>({ Id: '', Name: 'Naziv', Location : '', Facilities : '', MaxNumberOfGuests : 0, MinNumberOfGuests : 0 })
    
    const createAccomodation = () => {
        console.log(accomodation)
        accomodationService.createAccomodation(accomodation).then(() => {
                console.log('bravo');
        })
      }
    
    return <>
        <form className="form">
            <h2 className='heading'>New accomodation</h2>
            <div className='form-fields'>
            <div className='field'>
                <label>
                Name:
                <input type="text" name="name" className='input-field'
                onChange={(e) => setAccomodation(prevState => ({ ...prevState, Name: e.target.value }))}/>
                </label>
            </div>
            <div className='field'>
                <label>
                Location:
                <input type="text" name="name" 
                onChange={(e) => setAccomodation(prevState => ({ ...prevState, Location: e.target.value }))}/>
                </label>
            </div>
            <div className='field'>
                <label>
                Facilities:
                <input type="text" name="name" 
                onChange={(e) => setAccomodation(prevState => ({ ...prevState, Facilities: e.target.value }))}/>
                </label>
            </div>
            <div className='field'>
                <label>
                Minimum number of guests:
                <input type="number" name="name"   
                onChange={(e) => setAccomodation(prevState => ({ ...prevState, MinNumberOfGuests: parseInt(e.target.value) }))}/>
                </label>
            </div>
            <div className='field'>
                <label>
                Maximum number of guests:
                <input type="text" name="name" 
                onChange={(e) => setAccomodation(prevState => ({ ...prevState, MaxNumberOfGuests: parseInt(e.target.value) }))}/>
                </label>
            </div>
            <div className='field'>
            <button onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
            //e.stopPropagation()
            createAccomodation()
          }}
          >Create
          </button>
            </div>
            </div>
            
            
        </form>
    </>;
  }
export default RegisterAccomodation;
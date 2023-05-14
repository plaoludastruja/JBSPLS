import { useEffect, useState } from 'react';
import './RegisterAccomodation.css';
import Accomodation from '../../model/Accomodation';
import accomodationService from '../../services/accomodation.service';
import decodeToken from "../../services/auth.service";

function RegisterAccomodation() {

    const [accomodation, setAccomodation] = useState<Accomodation>({ id: '', name: 'Naziv', location : '', facilities : '', maxNumberOfGuests : 0, minNumberOfGuests : 0, isAutomaticApproved: 'true', hostUsername: '' })
    
    const createAccomodation = () => {
        console.log(accomodation)
        accomodationService.createAccomodation(accomodation).then(() => {
                console.log('bravo');
        })
      }
  useEffect(() => {
    setAccomodation(prevState => ({ ...prevState, hostUsername: decodeToken()?.username }))
  })
    
    return <>
        <div className="form">
            <h2 className='heading'>New accomodation</h2>
            <div className='form-fields'>
            <div className='field'>
                <label>
                Name:
                <input type="text" name="name" className='input-field'
                onChange={(e) => setAccomodation(prevState => ({ ...prevState, name: e.target.value }))}/>
                </label>
            </div>
            <div className='field'>
                <label>
                Location:
                <input type="text" name="name" 
                onChange={(e) => setAccomodation(prevState => ({ ...prevState, location: e.target.value }))}/>
                </label>
            </div>
            <div className='field'>
                <label>
                Facilities:
                <input type="text" name="name" 
                onChange={(e) => setAccomodation(prevState => ({ ...prevState, facilities: e.target.value }))}/>
                </label>
            </div>
            <div className='field'>
                <label>
                Minimum number of guests:
                <input type="number" name="name"   
                onChange={(e) => setAccomodation(prevState => ({ ...prevState, minNumberOfGuests: parseInt(e.target.value) }))}/>
                </label>
            </div>
            <div className='field'>
                <label>
                Maximum number of guests:
                <input type="text" name="name" 
                onChange={(e) => setAccomodation(prevState => ({ ...prevState, maxNumberOfGuests: parseInt(e.target.value) }))}/>
                </label>
            </div>
            <div className="field">
            <label>
              Price type:
              <select
                onChange={(e) =>
                  setAccomodation((prevState) => ({
                    ...prevState,
                    isAutomaticApproved: e.target.value ,
                  }))
                }
              >
                <option value="true">Automatic</option>
                <option value="false">Manual</option>
              </select>
            </label>
          </div>
            <div className='field'>
            <button onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
            e.stopPropagation()
            createAccomodation()
          }}
          >Create
          </button>
            </div>
            </div>
            
            
        </div>
    </>;
  }
export default RegisterAccomodation;
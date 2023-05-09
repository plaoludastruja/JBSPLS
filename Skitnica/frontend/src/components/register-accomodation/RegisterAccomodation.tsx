import './RegisterAccomodation.css';

function RegisterAccomodation() {
    return <>
        <form className="form">
            <h2 className='heading'>New accomodation</h2>
            <div className='form-fields'>
            <div className='field'>
                <label>
                Name:
                <input type="text" name="name" className='input-field'/>
                </label>
            </div>
            <div className='field'>
                <label>
                Location:
                <input type="text" name="name" />
                </label>
            </div>
            <div className='field'>
                <label>
                Facilities:
                <input type="text" name="name" />
                </label>
            </div>
            <div className='field'>
                <label>
                Minimum number of guests:
                <input type="text" name="name" />
                </label>
            </div>
            <div className='field'>
                <label>
                Maximum number of guests:
                <input type="text" name="name" />
                </label>
            </div>
            <div className='field'>
                <input type="submit" value="Submit" />
            </div>
            </div>
            
            
        </form>
    </>;
  }
export default RegisterAccomodation;
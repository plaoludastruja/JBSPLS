import React, { useState } from 'react';
import {MDBContainer, MDBCol, MDBRow, MDBBtn, MDBIcon, MDBInput, MDBCheckbox } from 'mdb-react-ui-kit';
import LoginDTO from '../../model/LoginDTO';
import userService from '../../services/user.service';
import { useNavigate } from 'react-router-dom';

function Login() {
  const navigate = useNavigate()
  const [user, setUser] = useState<LoginDTO>({
    username: "",
    password: "",
  })

  const loginUser = () => {
    userService.loginUser(user).then(() => {
      navigate('/')
    })
  }

  return (
    <MDBContainer fluid className="p-3 my-5 h-custom">

      <MDBRow>

        <MDBCol col='10' md='6'>
          <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-login-form/draw2.webp" className="img-fluid" alt="Sample image" />
        </MDBCol>

        <MDBCol col='4' md='6'>

          <div className="d-flex flex-row align-items-center justify-content-center">

            <p className="lead fw-normal mb-0 me-3">Sign in with</p>

          </div>

          <MDBInput wrapperClass='mb-4' label='Email address' onChange={(e) => setUser(prevState => ({ ...prevState, username: e.target.value }))} type='email' size="lg"/>
          <MDBInput wrapperClass='mb-4' label='Password' onChange={(e) => setUser(prevState => ({ ...prevState, password: e.target.value }))} type='password' size="lg"/>


          <div className='text-center text-md-start mt-4 pt-2'>
            <MDBBtn className="mb-0 px-5" size='lg' onClick={() => loginUser()}>Login</MDBBtn>
            <p className="small fw-bold mt-2 pt-1 mb-2">Don't have an account? <a href="/register" className="link-danger">Register</a></p>
          </div>

        </MDBCol>

      </MDBRow>

    </MDBContainer>
  );
}

export default Login;
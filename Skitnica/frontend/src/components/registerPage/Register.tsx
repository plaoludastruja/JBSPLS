import React, { useState } from "react";
import {
  MDBBtn,
  MDBContainer,
  MDBRow,
  MDBCol,
  MDBCard,
  MDBCardBody,
  MDBCardImage,
  MDBInput,
  MDBIcon,
  MDBCheckbox,
} from "mdb-react-ui-kit";
import userService from "../../services/user.service";
import User from "../../model/User";
import { useNavigate } from "react-router-dom";

function Register() {
  const navigate = useNavigate();
  const [user, setUser] = useState<User>({
    id: "",
    username: "",
    password: "",
    firstName: "",
    lastName: "",
    role: "USER",
    address: "",
    apiKey: "",
  });

  const registerUser = () => {
    userService.registerUser(user).then(() => {
      navigate("/login");
    });
  };

  const handleChangeRole = () => {
    if (user.role == "USER") {
      setUser((prevState) => ({ ...prevState, role: "HOST" }));
    } else {
      setUser((prevState) => ({ ...prevState, role: "USER" }));
    }
  };

  return (
    <MDBContainer fluid>
      <MDBCard className="text-black m-5" style={{ borderRadius: "25px" }}>
        <MDBCardBody>
          <MDBRow>
            <MDBCol
              md="10"
              lg="6"
              className="order-2 order-lg-1 d-flex flex-column align-items-center"
            >
              <p className="text-center h1 fw-bold mb-5 mx-1 mx-md-4 mt-4">
                Sign up
              </p>

              <div className="d-flex flex-row align-items-center mb-4">
                <MDBIcon fas icon="user me-3" size="lg" />
                <MDBInput
                  label="Your First Name"
                  onChange={(e) =>
                    setUser((prevState) => ({
                      ...prevState,
                      firstName: e.target.value,
                    }))
                  }
                  type="text"
                />
              </div>

              <div className="d-flex flex-row align-items-center mb-4">
                <MDBIcon fas icon="user me-3" size="lg" />
                <MDBInput
                  label="Your Last Name"
                  onChange={(e) =>
                    setUser((prevState) => ({
                      ...prevState,
                      lastName: e.target.value,
                    }))
                  }
                  type="text"
                />
              </div>

              <div className="d-flex flex-row align-items-center mb-4 ">
                <MDBIcon fas icon="map-marked-alt me-3" size="lg" />
                <MDBInput
                  label="Your Address"
                  onChange={(e) =>
                    setUser((prevState) => ({
                      ...prevState,
                      address: e.target.value,
                    }))
                  }
                  type="text"
                  className="w-100"
                />
              </div>

              <div className="d-flex flex-row align-items-center mb-4 ">
                <MDBIcon far icon="user me-3" size="lg" />
                <MDBInput
                  label="Your Username"
                  onChange={(e) =>
                    setUser((prevState) => ({
                      ...prevState,
                      username: e.target.value,
                    }))
                  }
                  type="text"
                  className="w-100"
                />
              </div>

              <div className="d-flex flex-row align-items-center mb-4">
                <MDBIcon fas icon="lock me-3" size="lg" />
                <MDBInput
                  label="Password"
                  onChange={(e) =>
                    setUser((prevState) => ({
                      ...prevState,
                      password: e.target.value,
                    }))
                  }
                  type="password"
                />
              </div>

              <div className="mb-4">
                <MDBCheckbox
                  name="flexCheck"
                  value=""
                  id="flexCheckDefault"
                  label="I want to be a host"
                  onChange={handleChangeRole}
                />
              </div>

              <MDBBtn className="mb-4" size="lg" onClick={registerUser}>
                Register
              </MDBBtn>
            </MDBCol>

            <MDBCol
              md="10"
              lg="6"
              className="order-1 order-lg-2 d-flex align-items-center"
            >
              <MDBCardImage
                src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-registration/draw1.webp"
                fluid
              />
            </MDBCol>
          </MDBRow>
        </MDBCardBody>
      </MDBCard>
    </MDBContainer>
  );
}

export default Register;

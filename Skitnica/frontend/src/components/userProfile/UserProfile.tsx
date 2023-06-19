import React, { useEffect, useState } from "react";
import {
  MDBCol,
  MDBContainer,
  MDBRow,
  MDBCard,
  MDBCardText,
  MDBCardBody,
  MDBCardImage,
  MDBBtn,
  MDBIcon,
  MDBModal,
  MDBModalDialog,
  MDBModalBody,
  MDBModalContent,
  MDBModalFooter,
  MDBModalHeader,
  MDBModalTitle,
  MDBInput,
  MDBBadge,
  MDBSwitch,
} from "mdb-react-ui-kit";
import User from "../../model/User";
import userService from "../../services/user.service";
import { removeToken } from "../../services/token.service";
import { useNavigate } from "react-router-dom";
import NotificationFilter from "../../model/NotificationFilter";

export default function UserProfile() {
  const navigate = useNavigate();
  const [basicModal, setBasicModal] = useState(false);
  const [deleteModal, setDeleteModal] = useState(false);
  const [notificationModal, setNotificationModal] = useState(false);

  const toggleShow = () => {
    setBasicModal(!basicModal);
    setUserEdit(user);
  };

  const toggleShowDelete = () => {
    setDeleteModal(!deleteModal);
  };

  const toggleShowNotification = () => {
    setNotificationModal(!notificationModal);
    setNotificationFilter(notificationFilter)
  };

  const [isBestHost, setIsBestHost] = useState(false);

  const [user, setUser] = useState<User>({
    id: "",
    username: "",
    password: "",
    firstName: "",
    lastName: "",
    role: "",
    address: "",
    apiKey: "",
  });

  const [userEdit, setUserEdit] = useState<User>({
    id: "",
    username: "",
    password: "",
    firstName: "",
    lastName: "",
    role: "",
    address: "",
    apiKey: "",
  });



  useEffect(() => {
    userService.getUserByUsername().then((response) => {
      setUser(response.data.user);
      console.log(response.data.user);
      userService.isBestHost(response.data.user.username).then((res1) => {
        console.log(res1.data)
        if (res1.data == true) {
          setIsBestHost(true);
        } else {
          setIsBestHost(false);
        }
      });
      userService.getNotificationFilter(response.data.user.username).then((res2) => {
        setNotificationFilter(res2.data.notificationFilter)
        setIsChecked1(res2.data.notificationFilter.reservation)
        setIsChecked2(res2.data.notificationFilter.rating)
        setIsChecked3(res2.data.notificationFilter.super)
      });

    });
  }, []);


  const editUser = () => {
    userService.editUser(userEdit).then(() => {
      setUser(userEdit);
      setBasicModal(!basicModal);
      removeToken();
      navigate("/login");
    });
  };

  const deleteUser = () => {
    userService.deleteUser(user.id).then(() => {
      removeToken();
      navigate("/login");
    });
  };




  const [notificationFilter, setNotificationFilter] = useState<NotificationFilter>(
    {
      id: "",
      username: "",
      reservation: true,
      rating: true,
      super: true

    }
  );

  const [isChecked1, setIsChecked1] = useState(notificationFilter.reservation);
  const handleSwitchChange1 = () => {
    setIsChecked1(!isChecked1);
  };
  const [isChecked2, setIsChecked2] = useState(notificationFilter.rating);
  const handleSwitchChange2 = () => {
    setIsChecked2(!isChecked2);
  };
  const [isChecked3, setIsChecked3] = useState(notificationFilter.super);
  const handleSwitchChange3 = () => {
    setIsChecked3(!isChecked3);
  };


  const editNotification = () => {
    const notificationFilterEdit : NotificationFilter = {
      id: notificationFilter.id ,
      username: notificationFilter.username,
      reservation: isChecked1,
      rating: isChecked2,
      super: isChecked3
    };
    userService.editNotificationFilter(notificationFilterEdit).then((res) => {
      toggleShowNotification()
      console.log(res.data.notificationFilter)
      setNotificationFilter(res.data.notificationFilter)
      setIsChecked1(res.data.notificationFilter.reservation)
        setIsChecked2(res.data.notificationFilter.rating)
        setIsChecked3(res.data.notificationFilter.super)
    });
  };

  return (
    <section style={{ backgroundColor: "#eee" }}>
      <MDBContainer className="py-5">
        <MDBRow>
          <MDBCol lg="4">
            <MDBCard className="mb-4">
              <MDBCardBody className="text-center">
                {isBestHost ? (
                  <MDBBadge className="position-absolute top-0 start-100 translate-middle p-3 bg-danger border border-light rounded-circle">
                    <i className="fas fa-star"></i>
                  </MDBBadge>
                ) : (
                  <div></div>
                )}

                <MDBCardImage
                  src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-chat/ava3.webp"
                  alt="avatar"
                  className="rounded-circle"
                  style={{ width: "150px" }}
                  fluid
                />
                {/*
                <p className="text-muted mb-1">Full Stack Developer</p>
                <p className="text-muted mb-4">Bay Area, San Francisco, CA</p>*/}
                <div className="d-flex justify-content-center mb-2">
                  <MDBBtn onClick={toggleShow}>Edit profile</MDBBtn>
                  <MDBBtn onClick={toggleShowNotification} outline className="ms-1">
                    Notification
                  </MDBBtn>
                  <MDBBtn className="ms-1" onClick={toggleShowDelete}>
                    <MDBIcon fas icon="trash" />
                  </MDBBtn>
                </div>
              </MDBCardBody>
            </MDBCard>
          </MDBCol>

          <MDBCol lg="8">
            <MDBCard className="mb-4">
              <MDBCardBody>
                <MDBRow>
                  <MDBCol sm="3">
                    <MDBCardText>Full Name</MDBCardText>
                  </MDBCol>
                  <MDBCol sm="9">
                    <MDBCardText className="text-muted">
                      {user.firstName + " " + user.lastName}
                    </MDBCardText>
                  </MDBCol>
                </MDBRow>
                <hr />
                <MDBRow>
                  <MDBCol sm="3">
                    <MDBCardText>Username</MDBCardText>
                  </MDBCol>
                  <MDBCol sm="9">
                    <MDBCardText className="text-muted">
                      {user.username}
                    </MDBCardText>
                  </MDBCol>
                </MDBRow>
                <hr />
                <MDBRow>
                  <MDBCol sm="3">
                    <MDBCardText>Role</MDBCardText>
                  </MDBCol>
                  <MDBCol sm="9">
                    <MDBCardText className="text-muted">
                      {user.role}
                    </MDBCardText>
                  </MDBCol>
                </MDBRow>
                <hr />
                <MDBRow>
                  <MDBCol sm="3">
                    <MDBCardText>Address</MDBCardText>
                  </MDBCol>
                  <MDBCol sm="9">
                    <MDBCardText className="text-muted">
                      {user.address}
                    </MDBCardText>
                  </MDBCol>
                </MDBRow>
              </MDBCardBody>
            </MDBCard>
          </MDBCol>
        </MDBRow>
      </MDBContainer>

      <MDBModal show={basicModal} tabIndex="-1">
        <MDBModalDialog>
          <MDBModalContent>
            <MDBModalHeader>
              <MDBModalTitle>Edit profile</MDBModalTitle>
              <MDBBtn
                className="btn-close"
                color="none"
                onClick={toggleShow}
              ></MDBBtn>
            </MDBModalHeader>
            <MDBModalBody>
              <MDBRow>
                <div className="align-items-center mb-4">
                  <MDBInput
                    label="Your First Name"
                    onChange={(e) =>
                      setUserEdit((prevState) => ({
                        ...prevState,
                        firstName: e.target.value,
                      }))
                    }
                    type="text"
                    value={userEdit.firstName}
                  />
                </div>

                <div className="lign-items-center mb-4">
                  <MDBInput
                    label="Your Last Name"
                    onChange={(e) =>
                      setUserEdit((prevState) => ({
                        ...prevState,
                        lastName: e.target.value,
                      }))
                    }
                    type="text"
                    value={userEdit.lastName}
                  />
                </div>

                <div className="align-items-center mb-4 ">
                  <MDBInput
                    label="Your Address"
                    onChange={(e) =>
                      setUserEdit((prevState) => ({
                        ...prevState,
                        address: e.target.value,
                      }))
                    }
                    type="text"
                    className="w-100"
                    value={userEdit.address}
                  />
                </div>

                <div className="align-items-center mb-4 ">
                  <MDBInput
                    label="Your Username"
                    onChange={(e) =>
                      setUserEdit((prevState) => ({
                        ...prevState,
                        username: e.target.value,
                      }))
                    }
                    type="text"
                    className="w-100"
                    value={userEdit.username}
                  />
                </div>

                <div className="align-items-center mb-4">
                  <MDBInput
                    label="Password"
                    onChange={(e) =>
                      setUserEdit((prevState) => ({
                        ...prevState,
                        password: e.target.value,
                      }))
                    }
                    type="password"
                    value={userEdit.password}
                  />
                </div>
              </MDBRow>
            </MDBModalBody>

            <MDBModalFooter>
              <MDBBtn color="secondary" onClick={toggleShow}>
                Close
              </MDBBtn>
              <MDBBtn onClick={editUser}>Save changes</MDBBtn>
            </MDBModalFooter>
          </MDBModalContent>
        </MDBModalDialog>
      </MDBModal>

      <MDBModal show={deleteModal} tabIndex="-1">
        <MDBModalDialog>
          <MDBModalContent>
            <MDBModalHeader>
              <MDBModalTitle>Delete profile</MDBModalTitle>
              <MDBBtn
                className="btn-close"
                color="none"
                onClick={toggleShowDelete}
              ></MDBBtn>
            </MDBModalHeader>

            <MDBModalFooter>
              <MDBBtn color="secondary" onClick={toggleShowDelete}>
                Close
              </MDBBtn>
              <MDBBtn onClick={deleteUser}>Delete</MDBBtn>
            </MDBModalFooter>
          </MDBModalContent>
        </MDBModalDialog>
      </MDBModal>


      <MDBModal show={notificationModal} tabIndex="-1">
        <MDBModalDialog>
          <MDBModalContent>
            <MDBModalHeader>
              <MDBModalTitle>Edit notification</MDBModalTitle>
              <MDBBtn
                className="btn-close"
                color="none"
                onClick={toggleShowNotification}
              ></MDBBtn>
            </MDBModalHeader>
            <MDBModalBody>
              <MDBRow>

                <MDBSwitch className="ms-1" label='Notifikacije za rezervacije' checked={isChecked1} onChange={handleSwitchChange1} />
                <MDBSwitch className="ms-1" label='Notifikacije za ocjenjivanje' checked={isChecked2} onChange={handleSwitchChange2} />
                <MDBSwitch className="ms-1" label='Notifikacije za super korisnika' checked={isChecked3} onChange={handleSwitchChange3} />
      
              </MDBRow>
            </MDBModalBody>

            <MDBModalFooter>
              <MDBBtn color="secondary" onClick={toggleShowNotification}>
                Close
              </MDBBtn>
              <MDBBtn onClick={editNotification}>Save changes</MDBBtn>
            </MDBModalFooter>
          </MDBModalContent>
        </MDBModalDialog>
      </MDBModal>
    </section>
  );
}

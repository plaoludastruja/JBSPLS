import {
  AiOutlineUserAdd as RegisterIcon,
  AiOutlineEdit,
} from "react-icons/ai";
import "./NavBar.css";
import { BsListStars, BsHouseCheck } from "react-icons/bs";
import { BiMoneyWithdraw, BiSearchAlt } from "react-icons/bi";
import {
  CgLogIn as LoginIcon,
  CgLogOut as LogoutIcon,
  CgProfile,
} from "react-icons/cg";
import { removeToken } from "../../services/token.service";
import { useEffect, useState } from "react";
import decodeToken from "../../services/auth.service";
import Notification from "../notification/Notification";

function NavBar() {
  const [isHost, setIsHost] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    if (decodeToken() != undefined) {
      setIsLoggedIn(true);
      if (decodeToken()?.role === "HOST") {
        setIsHost(true);
      }
    }
  }, []);

  const logout = () => {
    removeToken();
  };

  return (
    <nav
      className="navbar navbar-expand-lg navbar-custom shadow-5-strong gradient-custom"
      style={{ backgroundColor: "#f3cd87" }}
    >
      <button
        className="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarResponsive"
        aria-controls="navbarResponsive"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span className="navbar-toggler-icon"></span>
      </button>
      <a className="navbar-brand" href="/">
        Skitnica
      </a>
      <div className="collapse navbar-collapse" id="navbarResponsive">
        <ul className="navbar-nav mr-auto">
          {isLoggedIn ? (
            <>
              {isHost ? (
                <>
                  <li className="nav-item active">
                    <a className="nav-link" href="/create-accomodation">
                      <b>Create accomodation</b>
                      <BsHouseCheck
                        className="icon"
                        size={25}
                        color="#d88a3f"
                      />
                    </a>
                  </li>
                  <li className="nav-item active">
                    <a className="nav-link" href="/create-price">
                      <b>Create price</b>
                      <BiMoneyWithdraw
                        className="icon"
                        size={25}
                        color="#d88a3f"
                      />
                    </a>
                  </li>
                  <li className="nav-item active">
                    <a className="nav-link" href="/accomodations">
                      <b>Accomodations</b>
                      <AiOutlineEdit
                        className="icon"
                        size={25}
                        color="#d88a3f"
                      />
                    </a>
                  </li>
                  <li className="nav-item active">
                    <a className="nav-link" href="/hostsReservations">
                      <b>Reservation requests</b>
                      <BsListStars className="icon" size={25} color="#d88a3f" />
                    </a>
                  </li>
                  <li className="nav-item active">
                    <a className="nav-link" href="/profile">
                      <b>My profile</b>
                      <CgProfile className="icon" size={25} color="#d88a3f" />
                    </a>
                  </li>
                  <li className="nav-item active">
                    <a className="nav-link" href="/search">
                      <b>Search</b>
                      <BiSearchAlt className="icon" size={25} color="#d88a3f" />
                    </a>
                  </li>
                  <Notification/>
                  <li className="nav-item active">
                    <a
                      className="nav-link"
                      href="/login"
                      onClick={() => logout()}
                    >
                      <LogoutIcon className="icon" size={25} color="#d88a3f" />
                    </a>
                  </li>
                </>
              ) : (
                <>
                  <li className="nav-item active">
                    <a className="nav-link" href="/reservations">
                      <b>Reservations</b>
                      <BsListStars className="icon" size={25} color="#d88a3f" />
                    </a>
                  </li>
                  <li className="nav-item active">
                    <a className="nav-link" href="/profile">
                      <b>My profile</b>
                      <CgProfile className="icon" size={25} color="#d88a3f" />
                    </a>
                  </li>
                  <li className="nav-item active">
                    <a className="nav-link" href="/search">
                      <b>Search</b>
                      <BiSearchAlt className="icon" size={25} color="#d88a3f" />
                    </a>
                  </li>
                  <li className="nav-item active">
                    <a className="nav-link" href="/recommended">
                      <b>Recommended</b>
                      <BiSearchAlt className="icon" size={25} color="#d88a3f" />
                    </a>
                  </li>
                  <li className="nav-item active">
                  <a className="nav-link" href="/addHostGrade">
                    <b>Add Host Grade</b>
                      <BiSearchAlt className="icon" size={25} color="#d88a3f" />
                  </a>
              </li>
              <li className="nav-item active">
                  <a className="nav-link" href="/hosts">
                    <b>Hosts</b>
                      <BiSearchAlt className="icon" size={25} color="#d88a3f" />
                  </a>
              </li>
              <Notification/>
                  <li className="nav-item active">
                    <a
                      className="nav-link"
                      href="/login"
                      onClick={() => logout()}
                    >
                      <LogoutIcon className="icon" size={25} color="#d88a3f" />
                    </a>
                  </li>
                </>
              )}
            </>
          ) : (
            <>
              <li className="nav-item active">
                <a className="nav-link" href="/login">
                  <b>Login</b>
                  <LoginIcon className="icon" size={25} color="#d88a3f" />
                </a>
              </li>
              <li className="nav-item active">
                <a className="nav-link" href="/register">
                  <b>Register</b>
                  <RegisterIcon className="icon" size={25} color="#d88a3f" />
                </a>
              </li>
              <li className="nav-item active">
                  <a className="nav-link" href="/search">
                    <b>Search</b>
                      <BiSearchAlt className="icon" size={25} color="#d88a3f" />
                  </a>
              </li>
              <li className="nav-item active">
                  <a className="nav-link" href="/addHostGrade">
                    <b>Add Host Grade</b>
                      <BiSearchAlt className="icon" size={25} color="#d88a3f" />
                  </a>
              </li>
              <li className="nav-item active">
                  <a className="nav-link" href="/hosts">
                    <b>Hosts</b>
                      <BiSearchAlt className="icon" size={25} color="#d88a3f" />
                  </a>
              </li>
            </>
          )}
        </ul>
      </div>
    </nav>
  );
}
export default NavBar;

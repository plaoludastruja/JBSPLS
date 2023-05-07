import { HiOutlineHome as HomeIcon } from "react-icons/hi";
import { AiOutlineUserAdd as RegisterIcon } from "react-icons/ai";
import "./NavBar.css";

import { CgLogIn as LoginIcon, CgLogOut as LogoutIcon } from "react-icons/cg";

const logout = () => {
  //TODO
};

function NavBar() {
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
        Skitnica booking app
      </a>
      <div className="collapse navbar-collapse" id="navbarResponsive">
        <ul className="navbar-nav mr-auto">
          <li className="nav-item active">
            <a className="nav-link" href="/">
              <b>HOME</b>
              <HomeIcon className="icon" size={25} color="#d88a3f" />
            </a>
          </li>
          <li className="nav-item active">
            <a className="nav-link" href="/create-accomodation">
              <b>Create accomodation</b>
              <HomeIcon className="icon" size={25} color="#d88a3f" />
            </a>
          </li>
          <li className="nav-item active">
            <a className="nav-link" href="/login">
              <b>LOGIN</b>
              <LoginIcon className="icon" size={25} color="#d88a3f" />
            </a>
          </li>
          <li className="nav-item active">
            <a className="nav-link" href="#">
              <b>REGISTER</b>
              <RegisterIcon className="icon" size={25} color="#d88a3f" />
            </a>
          </li>
          <li className="nav-item active">
            <a
              className="nav-link"
              href="/login"
              onClick={() => {
                logout();
              }}
            >
              <b>LOGOUT</b>
              <LogoutIcon className="icon" size={25} color="#d88a3f" />
            </a>
          </li>
        </ul>
      </div>
    </nav>
  );
}
export default NavBar;

import { AiOutlineUserAdd as RegisterIcon } from "react-icons/ai";
import "./NavBar.css";
import { BsListStars, BsHouseCheck } from "react-icons/bs";
import { BiMoneyWithdraw } from "react-icons/bi";
import { CgLogIn as LoginIcon, CgLogOut as LogoutIcon } from "react-icons/cg";

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
            <a className="nav-link" href="/create-accomodation">
              <b>Create accomodation</b>
              <BsHouseCheck className="icon" size={25} color="#d88a3f" />
            </a>
          </li>
          <li className="nav-item active">
            <a className="nav-link" href="/reservations">
              <b>Reservations</b>
              <BsListStars className="icon" size={25} color="#d88a3f" />
            </a>
          </li>
          <li className="nav-item active">
            <a className="nav-link" href="/create-price">
              <b>Create price</b>
              <BiMoneyWithdraw className="icon" size={25} color="#d88a3f" />
            </a>
          </li>
          <li className="nav-item active">
            <a className="nav-link" href="/accomodations">
              <b>Accomodations</b>
              <HomeIcon className="icon" size={25} color="#d88a3f" />
            </a>
          </li>
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
            <a className="nav-link" href="/login">
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

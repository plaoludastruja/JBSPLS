import "./App.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import NavBar from "./components/navBar/NavBar";
import Login from "./components/login/Login";
import HomePage from "./components/homePage/HomePage";
import RegisterAccomodation from "./components/register-accomodation/RegisterAccomodation";
import Reservations from "./components/reservations/Reservations";
import Register from "./components/registerPage/Register";
import RegisterPrice from "./components/register-price/RegisterPrice";

function App() {
  return (
    <BrowserRouter>
      <div className="main-layout">
        <NavBar />
        <div className="container">
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/" element={<HomePage />} />
            <Route path="*" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route
              path="/create-accomodation"
              element={<RegisterAccomodation />}
            />
            <Route path="/reservations" element={<Reservations />} />
            <Route path="/create-price" element={<RegisterPrice />} />
          </Routes>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;

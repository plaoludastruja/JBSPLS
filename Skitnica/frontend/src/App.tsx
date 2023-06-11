import "./App.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import NavBar from "./components/navBar/NavBar";
import Login from "./components/login/Login";
import HomePage from "./components/homePage/HomePage";
import RegisterAccomodation from "./components/register-accomodation/RegisterAccomodation";
import Reservations from "./components/reservations/Reservations";
import Register from "./components/registerPage/Register";
import RegisterPrice from "./components/register-price/RegisterPrice";
import Accomodations from "./components/accomodations/Accomodations";
import UserProfile from "./components/userProfile/UserProfile";
import SearchAccomodations from "./components/search/SearchAccomodations";
import HostsReservations from "./components/hosts-reservations/HostsReservations";
import GradeManagement from "./components/gradeManagement/GradeManagement";
import HostsAndGrades from "./components/hostsAndGrades/HostsAndGrades";


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
            <Route path="/accomodations" element={<Accomodations />} />
            <Route path="/profile" element={<UserProfile />} />
            <Route path="/search" element={<SearchAccomodations />} />
            <Route path="/hostsReservations" element={<HostsReservations />} />
            <Route path="/addHostGrade" element={<GradeManagement />} />
            <Route path="/hosts" element={<HostsAndGrades />} />
          </Routes>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;

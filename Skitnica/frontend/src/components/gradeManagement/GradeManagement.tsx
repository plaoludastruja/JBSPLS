import { useEffect, useState } from "react";
import HostMark from "../../model/HostMark";
import { baseAxios } from "./../../services/api.service";
import reservationService from "../../services/reservation.service";
import decodeToken from "../../services/auth.service";
import { MDBBtn, MDBCard, MDBCardBody, MDBCardText, MDBCardTitle } from "mdb-react-ui-kit";
import { isArray } from "util";
import hostMarkService from "../../services/hostMark.service";
import "./GradeManagement.css";

function GradeManagement() {
  const [hostMarks, setHostMarks] = useState<HostMark[]>([]);
  const [hosts, setHosts] = useState<string[]>([]);
  const [grade, setGrade] = useState<HostMark>(
    {
      id: "",
      username: "",
      grade: 0,
      hostUsername: ""
    }
  );
  const [gradeExists, setGradeExists] = useState<Boolean>(false);
  const [rating, setRating] = useState(0);
  const [hover, setHover] = useState(0);
  const [username, setUsername] = useState<string|undefined>("");

  useEffect(() => {
    reservationService
      .getByGuest(decodeToken()?.username)
      .then((response) => {
        setHosts(response.data.usernames);
      });
      console.log(Array.isArray(hosts))
      setUsername(decodeToken()?.username)
  }, []);

  const checkExisting = (hostUsername: string): boolean => {
      hostMarkService.getByHostAndUsername(decodeToken()?.username, hostUsername).then(
        (response) => {
          console.log(response.data.hostMark)
          if(response.data.hostMark.length == 0){
            console.log("nema ocenu")
            return false
          }else{
            console.log("ima ocenu")
            return true
          }
          
        }
      )
      return true
    
  }

  const setGradeForHost = (host: string) => {
    console.log(rating)
    
    setGrade((prevState) => ({
      ...prevState,
      username: username,
      hostUsername: host,
      grade: rating
    }))
    console.log(grade)
}

  return (
    <>
      Add grade
      <div>
        {hosts.map((host) => (
          <div key={host}>
            <MDBCard>
              <MDBCardBody>
                <MDBCardTitle>{host}</MDBCardTitle>
                <MDBCardText>
                  <div style={{ display: checkExisting(host) ? 'block' : 'none' }}>
                  <div className="star-rating">
      {[...Array(5)].map((star, index) => {
        index += 1;
        return (
          <button
            type="button"
            key={index}
            className={index <= (hover || rating) ? "on" : "off"}
            onClick={() => setRating(index)}
            onMouseEnter={() => setHover(index)}
            onMouseLeave={() => setHover(rating)}
          >
            <span className="star">&#9733;</span>
          </button>
          
        );
      })}
      <button onClick={() => setGradeForHost(host)}>Set grade</button>
    </div>
                  </div>
                </MDBCardText>
                <MDBBtn onClick={() => checkExisting(host)}>
                  My grade
                </MDBBtn>
              </MDBCardBody>
            </MDBCard>
          </div>
        ))}
    </div>
    
    </>
  );
}
export default GradeManagement;

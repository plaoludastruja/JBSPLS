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
  const [hostUsername, setHostUsername] = useState<string|undefined>("");
  const [addGrade, setAddGrade] = useState(false);
  const [editGrade, setEditGrade] = useState(false);

  useEffect(() => {
    reservationService
      .getByGuest(decodeToken()?.username)
      .then((response) => {
        setHosts(response.data.usernames);
      });
      console.log(Array.isArray(hosts))
      setUsername(decodeToken()?.username)
  }, []);

  const checkExisting = (hostUsername: string) => {
    setHostUsername(hostUsername)
    var ret = false
      hostMarkService.getByHostAndUsername(decodeToken()?.username, hostUsername).then(
        (response) => {
          console.log(response.data.hostMark)
          setHostMarks(response.data.hostMark)
          if(response.data.hostMark.length != 0){
            setAddGrade(false)
            setEditGrade(true)
            setGrade((prevState) => ({
              ...prevState,
              id: response.data.hostMark[0].id,
              grade: response.data.hostMark[0].grade
            }))
            console.log("ima ocenu")
          }else{
            setAddGrade(true)
            setEditGrade(false)
            console.log("nema ocenu")
          }
          
        }
      )
      
      console.log(addGrade)
      console.log(editGrade)
      return ret
    
  }

  const sleep = (milliseconds: number) => {
    return new Promise(resolve => setTimeout(resolve, milliseconds))
}

  const setGradeForHost = () => {
    console.log(rating)
    
    /*setGrade((prevState) => ({
      ...prevState,
      username: username,
      hostUsername: host,
      grade: rating
    }))*/
    console.log(grade)
    hostMarkService.createHostGrade(grade).then(() => {
      alert("Successfully added grade!");
    });
}

const editGradeForHost = () => {
  console.log(rating)
  
  /*setGrade((prevState) => ({
      ...prevState,
      grade: rating
    }))*/
  console.log(grade)
  console.log(rating)
  hostMarkService.editHostGrade(grade).then(() => {
    alert("Successfully changed grade!");
  });
  console.log(hostMarks[0].id)
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
                  
                </MDBCardText>
                <MDBBtn onClick={() => checkExisting(host)}>
                  My grade
                </MDBBtn>
              </MDBCardBody>
            </MDBCard>
          </div>
        ))}
    </div>
    {addGrade && (
      <div>
      <div className="star-rating">
{[...Array(5)].map((star, index) => {
index += 1;
return (
<button
type="button"
key={index}
className={index <= (hover || rating) ? "on" : "off"}
onClick={function(event){
  setGrade((prevState) => ({
    ...prevState,
    username: username,
    hostUsername: hostUsername,
    grade: index
  }))
  setRating(index)
}
}
  
  
onMouseEnter={() => setHover(index)}
onMouseLeave={() => setHover(rating)}
>
<span className="star">&#9733;</span>
</button>

);
})}
<button onClick={() => setGradeForHost()}>Set grade</button>
</div>
      </div>
      )}
      {editGrade && (
        <div>
          <p>Old grade: {grade.grade}</p>

<div>
      <div className="star-rating">
{[...Array(5)].map((star, index) => {
index += 1;
return (
<button
type="button"
key={index}
className={index <= (hover || rating) ? "on" : "off"}
onClick={function(event){
  setGrade((prevState) => ({
    ...prevState,
    username: username,
    hostUsername: hostUsername,
    grade: index
  }))
  setRating(index)
}
}
  
  
onMouseEnter={() => setHover(index)}
onMouseLeave={() => setHover(rating)}
>
<span className="star">&#9733;</span>
</button>

);
})}
<button onClick={() => editGradeForHost()}>Set grade</button>
</div>
      </div>
        </div>
        

        
    )}


    
    </>
  );
  
}
export default GradeManagement;

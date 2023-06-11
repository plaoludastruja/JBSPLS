import { useEffect, useState } from "react";
import HostMark from "../../model/HostMark";
import { baseAxios } from "./../../services/api.service";
import reservationService from "../../services/reservation.service";
import decodeToken from "../../services/auth.service";
import { MDBBtn, MDBCard, MDBCardBody, MDBCardText, MDBCardTitle } from "mdb-react-ui-kit";
import { isArray } from "util";
import hostMarkService from "../../services/hostMark.service";

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
  
  useEffect(() => {
    reservationService
      .getByGuest(decodeToken()?.username)
      .then((response) => {
        setHosts(response.data.usernames);
      });
      console.log(Array.isArray(hosts))
      
  }, []);

  const something = (hostUsername: string) => {
    try{
      hostMarkService.getByHostAndUsername(decodeToken()?.username, hostUsername).then(
        (response) => {
          console.log(response.data)
          //setGradeExists(true);
        }
      )
    }catch(error){
      console.log("greska")
    }
    
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
                  <div>
                    unesi, izmeni
                  </div>
                </MDBCardText>
                <MDBBtn onClick={() => something(host)}>
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

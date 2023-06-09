import { useEffect, useState } from "react";
import HostMark from "../../model/HostMark";
import { baseAxios } from "./../../services/api.service";
import reservationService from "../../services/reservation.service";
import decodeToken from "../../services/auth.service";
import { MDBBtn, MDBCard, MDBCardBody, MDBCardText, MDBCardTitle } from "mdb-react-ui-kit";
import { isArray } from "util";

function GradeManagement() {
  const [hostMarks, setHostMarks] = useState<HostMark[]>([]);
  const [hosts, setHosts] = useState<string[]>([]);
  
  useEffect(() => {
    reservationService
      .getByGuest(decodeToken()?.username)
      .then((response) => {
        setHosts(response.data.usernames);
      });
      console.log(Array.isArray(hosts))
  }, []);

  const something = () => {

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
                <MDBBtn onClick={() => something()}>
                  Show prices
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

import { MDBBtn } from "mdb-react-ui-kit";
import HostMark from "../../model/HostMark";
import User from "../../model/User";
import userService from "../../services/user.service";
import { baseAxios } from "./../../services/api.service";
import { useEffect, useState } from "react";

function HostsAndGrades() {
    const [hosts, setHosts] = useState<User[]>([]);
    const [hostMarks, setHostMarks] = useState<HostMark[]>([]);

    useEffect(() => {
        userService
          .getHosts()
          .then((response) => {
            setHosts(response.data.users);
            //console.log(response.data.users)
          });
          
      }, []);
    
      const showGrades = (username: string) => {
        console.log(username)
      };

  return <>Hosts grades
    <div className="row">
        <div className="col-xl-8">
          <div className="card mb-4 mt-4">
            <div className="card-header">All reservations</div>
            <div className="card-body">
              <div className="table-responsive">
                <table className="table">
                  <thead>
                    <tr>
                      <th>First name</th>
                      <th>Last name</th>
                      <th>Show grades</th>
                      <th></th>
                    </tr>
                  </thead>
                  {hosts.map((host) => (
                    <tbody key={host.id}>
                      <tr>
                        <td>{host.firstName}</td>
                        <td>{host.lastName}</td>
                        <td><MDBBtn onClick={() => showGrades(host.username)}>
                              Show grades
                            </MDBBtn></td>
                        
                      </tr>
                    </tbody>
                  ))}
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
  </>;
}
export default HostsAndGrades;

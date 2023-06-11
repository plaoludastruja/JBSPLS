import { MDBBtn } from "mdb-react-ui-kit";
import HostMark from "../../model/HostMark";
import User from "../../model/User";
import userService from "../../services/user.service";
import { baseAxios } from "./../../services/api.service";
import { useEffect, useState } from "react";
import hostMarkService from "../../services/hostMark.service";

function HostsAndGrades() {
    const [hosts, setHosts] = useState<User[]>([]);
    const [hostMarks, setHostMarks] = useState<HostMark[]>([]);
    const [averageGrade, setAverageGrade] = useState<number>(0);

    const [show, setShow] = useState(false);

    useEffect(() => {
        userService
          .getHosts()
          .then((response) => {
            setHosts(response.data.users);
            //console.log(response.data.users)
          });
          
      }, []);
    
      const showGrades = (username: string) => {
        setShow(true);
        hostMarkService.getByHost(username).then((response) => {
            setHostMarks(response.data.hostMark)
            var sum = 0;
            var count = 0;
            for (var hostMark of response.data.hostMark) {
                sum += hostMark.grade;
                count++;
            }
            if(count == 0){
                setAverageGrade(0);
            }else{
                setAverageGrade(sum/count);
            }
            
        }
        )
        
        //console.log(username);
      };

  return <>Hosts grades
    <div className="row">
        <div className="col-xl-8">
          <div className="card mb-4 mt-4">
            <div className="card-header">All hosts</div>
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
      {show && (
        <div>
            <div className="row">
        <div className="col-xl-8">
          <div className="card mb-4 mt-4">
            <div className="card-header">Grades</div>
            <div className="card-body">
              <div className="table-responsive">
                <table className="table">
                  <thead>
                    <tr>
                      <th>Username</th>
                      <th>Grade</th>
                      <th>Time</th>
                    </tr>
                  </thead>
                  {hostMarks.map((grade) => (
                    <tbody key={grade.id}>
                      <tr>
                        <td>{grade.username}</td>
                        <td>{grade.grade}</td>
                        <td>vreme</td>
                        
                      </tr>
                    </tbody>
                  ))}
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div> Average grade: {averageGrade}</div>
        </div>
        
      )}
  </>;
}
export default HostsAndGrades;

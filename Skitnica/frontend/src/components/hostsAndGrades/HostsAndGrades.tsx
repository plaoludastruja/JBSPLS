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

  return <>Hosts grades</>;
}
export default HostsAndGrades;

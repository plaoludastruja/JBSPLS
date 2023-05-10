import { baseAxios } from "./../../services/api.service";

function Accomodations() {
  const send = () => {
    baseAxios.get("/user").then((res) => {
      alert(res);
    });
  };

  return (
    <>
      Accomodations
    </>
  );
}
export default Accomodations;

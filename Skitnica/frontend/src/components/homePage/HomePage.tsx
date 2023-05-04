import { baseAxios } from "./../../services/api.service";

function HomePage() {
  const send = () => {
    baseAxios.get("/user").then((res) => {
      alert(res);
    });
  };

  return (
    <>
      Home page TODO
      <button onClick={send}>Send</button>
    </>
  );
}
export default HomePage;

import { useState, useEffect } from "react";
import Appointment from "../../model/Appointment";
import Accomodation from "../../model/Accomodation";
import accomodationService from "../../services/accomodation.service";
import priceService from "../../services/price.service";
import decodeToken from "../../services/auth.service";

function RegisterPrice() {
  const [appointment, setAppointment] = useState<Appointment>({
    id: "",
    accomodationId: "",
    start: "",
    end: "",
    priceType: "PerPerson",
    price: 0,
    status: "Free",
  });
  const [accomodations, setAccomodations] = useState<Accomodation[]>([]);
  const [errorMessage, setErrorMessage] = useState("");

  useEffect(() => {
    accomodationService
      .getAccomodationsByHostUsername(decodeToken()?.username)
      .then((response) => {
        setAccomodations(response.data.accomodations);
        setAppointment((prevState) => ({
          ...prevState,
          accomodationId: response.data.accomodations[0].id,
        }));
      });
  }, []);
  //console.log(accomodations);

  const checkDates = (): boolean => {
    if (new Date(appointment.start) > new Date(appointment.end)) {
      console.log("udje ovde");
      return false;
    }
    return true;
  };
  const createAppointment = () => {
    if (
      appointment.start == "" ||
      appointment.end == "" ||
      appointment.price == 0
    ) {
      setErrorMessage("Incorrect dates");
    } else {
      if (checkDates()) {
        priceService.createAppointment(appointment).then(() => {
          setErrorMessage("Successful");
        });
      } else {
        setErrorMessage("Incorrect dates");
      }
    }
    console.log(appointment);
  };

  return (
    <>
      <div className="form">
        <h2 className="heading">New price</h2>
        <div className="form-fields">
          <div className="field">
            <label>Accomodation:</label>
            <select
              onChange={(e) =>
                setAppointment((prevState) => ({
                  ...prevState,
                  accomodationId: e.target.value,
                }))
              }
            >
              {accomodations.map((accomodation, index) => (
                <option value={accomodation.id} key={index}>
                  {accomodation.name}
                </option>
              ))}
            </select>
          </div>
          <div className="field">
            <label>
              Start:
              <input
                type="date"
                name="name"
                onChange={(e) =>
                  setAppointment((prevState) => ({
                    ...prevState,
                    start: e.target.value,
                  }))
                }
              />
            </label>
          </div>
          <div className="field">
            <label>
              End:
              <input
                type="date"
                name="name"
                onChange={(e) =>
                  setAppointment((prevState) => ({
                    ...prevState,
                    end: e.target.value,
                  }))
                }
              />
            </label>
          </div>
          <div className="field">
            <label>
              Price type:
              <select
                onChange={(e) =>
                  setAppointment((prevState) => ({
                    ...prevState,
                    priceType: e.target.value,
                  }))
                }
              >
                <option value="PerPerson">Per person</option>
                <option value="PerUnit">Per unit</option>
              </select>
            </label>
          </div>
          <div className="field">
            <label>
              Price:
              <input
                type="number"
                name="name"
                onChange={(e) =>
                  setAppointment((prevState) => ({
                    ...prevState,
                    price: parseInt(e.target.value),
                  }))
                }
              />
            </label>
          </div>
          <div className="field">
            <button
              onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
                e.stopPropagation();
                createAppointment();
              }}
            >
              Create
            </button>
            {errorMessage && (
              <label className="error-message">{errorMessage}</label>
            )}
          </div>
        </div>
      </div>
    </>
  );
}
export default RegisterPrice;

import { useEffect, useState } from "react";
import "./RegisterAccomodation.css";
import Accomodation from "../../model/Accomodation";
import accomodationService from "../../services/accomodation.service";
import decodeToken from "../../services/auth.service";
import * as base64 from "base64-js";

function RegisterAccomodation() {
  const [accomodation, setAccomodation] = useState<Accomodation>({
    id: "",
    name: "Naziv",
    location: "",
    facilities: "",
    maxNumberOfGuests: 0,
    minNumberOfGuests: 0,
    isAutomaticApproved: "true",
    hostUsername: "",
    image: "",
  });

  const createAccomodation = () => {
    console.log(accomodation);
    accomodationService.createAccomodation(accomodation).then(() => {
      alert("Successfully created!");
    });
  };
  useEffect(() => {
    setAccomodation((prevState) => ({
      ...prevState,
      hostUsername: decodeToken()?.username,
    }));
  });
  const handleImageInput = (files: any) => {
    const file = files && files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = (event) => {
        const result = event.target?.result;
        if (result && typeof result === "object") {
          const buffer = new Uint8Array(result);
          const encodedImage = base64.fromByteArray(buffer);
          setAccomodation((prevState) => ({
            ...prevState,
            image: encodedImage,
          }));
        }
      };
      reader.readAsArrayBuffer(file);
    }
  };

  return (
    <>
      <div className="form">
        <h2 className="heading">New accomodation</h2>
        <div className="form-fields">
          <div className="field">
            <label>
              Name:
              <input
                type="text"
                name="name"
                className="input-field"
                onChange={(e) =>
                  setAccomodation((prevState) => ({
                    ...prevState,
                    name: e.target.value,
                  }))
                }
              />
            </label>
          </div>
          <div className="field">
            <label>
              Location:
              <input
                type="text"
                name="name"
                onChange={(e) =>
                  setAccomodation((prevState) => ({
                    ...prevState,
                    location: e.target.value,
                  }))
                }
              />
            </label>
          </div>
          <div className="field">
            <label>
              Facilities:
              <input
                type="text"
                name="name"
                onChange={(e) =>
                  setAccomodation((prevState) => ({
                    ...prevState,
                    facilities: e.target.value,
                  }))
                }
              />
            </label>
          </div>
          <div className="field">
            <label>
              Image:
              <input
                type="file"
                name="name"
                onChange={(e) => handleImageInput(e.target.files)}
              />
            </label>
          </div>
          <div className="field">
            <label>
              Minimum number of guests:
              <input
                type="number"
                name="name"
                onChange={(e) =>
                  setAccomodation((prevState) => ({
                    ...prevState,
                    minNumberOfGuests: parseInt(e.target.value),
                  }))
                }
              />
            </label>
          </div>
          <div className="field">
            <label>
              Maximum number of guests:
              <input
                type="text"
                name="name"
                onChange={(e) =>
                  setAccomodation((prevState) => ({
                    ...prevState,
                    maxNumberOfGuests: parseInt(e.target.value),
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
                  setAccomodation((prevState) => ({
                    ...prevState,
                    isAutomaticApproved: e.target.value,
                  }))
                }
              >
                <option value="true">Automatic</option>
                <option value="false">Manual</option>
              </select>
            </label>
          </div>
          <div className="field">
            <button
              onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
                e.stopPropagation();
                createAccomodation();
              }}
            >
              Create
            </button>
          </div>
        </div>
      </div>
    </>
  );
}
export default RegisterAccomodation;

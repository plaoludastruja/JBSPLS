import React, { useEffect, useState } from 'react';
import { BiNotification } from 'react-icons/bi';
import decodeToken from "../../services/auth.service";
import {
  connect,
  NatsConnection,
  Msg,
  StringCodec
} from "nats.ws";
import { MDBModal, MDBModalDialog, MDBModalContent, MDBModalHeader, MDBModalTitle, MDBBtn, MDBModalBody, MDBRow, MDBInput, MDBModalFooter, MDBBadge, MDBListGroup, MDBListGroupItem, MDBIcon } from 'mdb-react-ui-kit';
import userService from '../../services/user.service';
import NotificationModel from '../../model/NotificationModel';
import { FiBell, FiBellOff } from "react-icons/fi";

export default function Notification() {
    const [nats, setNats] = useState<NatsConnection>();
    const sc = StringCodec();
    const [incomingNotification, setIncomingNotification] = useState(false);
    useEffect(() => {

    (async () => {
      const nc = await connect({ servers: ["ws://localhost:4222"], })
      setNats(nc)
      console.log("connected to NATS")

      const sub = nc.subscribe(decodeToken()?.username || '');
      
      (async () => {
        for await (const m of sub) {
          console.log(`[${sub.getProcessed()}]: ${sc.decode(m.data)}`);
          setIncomingNotification(true)
        }
        console.log("subscription closed");
      })();
    })();

    return () => {
      nats?.drain();
      console.log("closed NATS connection")
    }
  }, []);

  const [basicModal, setBasicModal] = useState(false);
  const [notifications, setNotifications] = useState<NotificationModel[]>([]);

  const toggleShow = () => {
    if (!basicModal) {
        userService.getAllByReceiver(decodeToken()?.username || '').then((response) => {
            setNotifications(response.data.notifications.reverse());
            setIncomingNotification(false)
          });
    }else{
        userService.readAllByUsername(decodeToken()?.username || '')
    }
    setBasicModal(!basicModal)
  }


  const renderList = notifications.map((item : NotificationModel, index) => (
    <MDBListGroupItem className='d-flex justify-content-between align-items-center'>
        {item.message}
        <MDBBadge pill light>
            {item.isRead === "true" ? <MDBIcon fas icon="check" color="success"/> : <MDBIcon fas icon="times" color='danger'/>}
        </MDBBadge>
    </MDBListGroupItem>
    
  ));

  return (
    <>
        <li className="nav-item active" onClick={toggleShow}>
            <a className="nav-link">
            {incomingNotification ? <FiBell className="icon" size={25} color="red" /> : <FiBellOff className="icon" size={25} color="#d88a3f" />}
            </a>
        </li>

        <MDBModal show={basicModal}  tabIndex="-1">
        <MDBModalDialog>
          <MDBModalContent>
            <MDBModalHeader>
              <MDBModalTitle>Notification</MDBModalTitle>
              <MDBBtn
                className="btn-close"
                color="none"
                onClick={toggleShow}
              ></MDBBtn>
            </MDBModalHeader>
            <MDBModalBody>
                <MDBListGroup style={{ minWidth: '22rem' }} light>
                    {renderList}
                </MDBListGroup>
            </MDBModalBody>

            <MDBModalFooter>
              <MDBBtn color="secondary" onClick={toggleShow}>
                Close
              </MDBBtn>
            </MDBModalFooter>
          </MDBModalContent>
        </MDBModalDialog>
      </MDBModal>
    </>
  );
}
import { faChalkboardTeacher } from "@fortawesome/free-solid-svg-icons/faChalkboardTeacher";
import { faFileImport } from "@fortawesome/free-solid-svg-icons/faFileImport";
import { faFileUpload } from "@fortawesome/free-solid-svg-icons/faFileUpload";
import { faHandsHelping } from "@fortawesome/free-solid-svg-icons/faHandsHelping";
import { faHeading } from "@fortawesome/free-solid-svg-icons/faHeading";
import { faUpload } from "@fortawesome/free-solid-svg-icons/faUpload";
import { faUserPlus } from "@fortawesome/free-solid-svg-icons/faUserPlus";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React, { useState } from "react";
import { Link, useParams } from "react-router-dom";
import { List, Menu, Sidebar } from "semantic-ui-react";
import AddTAModal from "./AddTAModal";

const sidePanelLStyle = {
  backgroundColor: "#478dcc",
  width: 250,
  top: 70
};

const sidePanelItemLStyle = {
  overflow: "auto",
  color: "white",
  fontSize: "20px"
};

function SideBar(props) {
  const [state, setState] = useState({
    activeItem: "Personal Information",
    addingTA: false
  });

  let isProfTA =
    props.courseTAs.some(ta => ta.username === props.username) || props.isProf;

  console.log(
    "{}{}{}",
    props.courseTAs.some(ta => ta.username === props.username),
    props.isProf,
    isProfTA
  );

  let { id } = useParams();
  id = id.substring(1);

  let uploadPath = isProfTA
    ? "/course:" + id + "/upload"
    : "/course:" + id + "/offer";

  return (
    <Sidebar
      as={Menu}
      animation="overlay"
      icon="labeled"
      direction="left"
      vertical
      visible={props.isMobile ? props.sidebarIsOpen : true}
      width="thin"
      style={sidePanelLStyle}
    >
      <AddTAModal
        open={state.addingTA}
        setOpen={setState}
        courseID={id}
        students={props.students}
      />
      <Menu.Item style={sidePanelItemLStyle} as="a">
        <FontAwesomeIcon size="1x" icon={faHeading} />
        <span>
          <b>&nbsp;&nbsp;{props.courseTitle ? props.courseTitle : ""}</b>
        </span>
      </Menu.Item>
      <Menu.Item style={sidePanelItemLStyle} as="a">
        <FontAwesomeIcon size="1x" icon={faChalkboardTeacher} />
        <span>
          <b>
            &nbsp;&nbsp;{props.courseProf.name ? props.courseProf.name : ""}
          </b>
        </span>
      </Menu.Item>
      <Menu.Item as="a" style={sidePanelItemLStyle}>
        <FontAwesomeIcon size="1x" icon={faHandsHelping} />
        <span>
          <b>&nbsp;&nbsp;TA List</b>
        </span>
        <List>
          {props.courseTAs.map(TA => {
            return (
              <div>
                <List.Item as="li">
                  {TA.name}
                </List.Item>
              </div>
            );
          })}
        </List>
      </Menu.Item>

      {isProfTA && (
        <Menu.Item
          positive
          onClick={() => {
            setState({ ...state, addingTA: true });
          }}
          style={sidePanelItemLStyle}
        >
          <FontAwesomeIcon size="1x" icon={faUserPlus} />
          <span>
            <b>&nbsp;&nbsp;Add TA</b>
          </span>
        </Menu.Item>
      )}
      <Link to={uploadPath + "/video"}>
        <Menu.Item style={sidePanelItemLStyle} as="a">
          <FontAwesomeIcon size="1x" icon={faUpload} />
          <span>
            <b>&nbsp;&nbsp;Upload Videos</b>
          </span>
        </Menu.Item>
      </Link>
      <Link to={uploadPath + "/attachment"}>
        <Menu.Item style={sidePanelItemLStyle} as="a">
          <FontAwesomeIcon size="1x" icon={faFileUpload} />
          <span>
            <b>&nbsp;&nbsp;Upload Attachment</b>
          </span>
        </Menu.Item>
      </Link>
      {isProfTA && (
        <Link to={"/course:" + id + "/pendings"}>
          <Menu.Item style={sidePanelItemLStyle} as="a">
            <FontAwesomeIcon size="1x" icon={faFileImport} />
            <span>
              <b>&nbsp;&nbsp;Pending Contents</b>
            </span>
          </Menu.Item>
        </Link>
        // </Link>
      )}
    </Sidebar>
  );
}

export default SideBar;

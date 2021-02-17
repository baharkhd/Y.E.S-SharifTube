import React, { useState } from "react";
import { Sidebar, Menu, Button, Icon } from "semantic-ui-react";
import { useHistory, Link } from "react-router-dom";
import AddCourseModal from "./AddCourse.js";
import JoinCourseModel from "./JoinCourse.js";
import { faChalkboardTeacher } from "@fortawesome/free-solid-svg-icons/faChalkboardTeacher";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPlusCircle } from "@fortawesome/free-solid-svg-icons/faPlusCircle";
import { faSearchPlus } from "@fortawesome/free-solid-svg-icons/faSearchPlus";

const sidePanelLStyle = {
  backgroundColor: "#5383ff",
  width: 250,
  top: 70
};

const sidePanelItemLStyle = {
  color: "white"
};

function SideBar(props) {
  const [state, setState] = useState({
    activeItem: "Personal Information",
    addingCourse: false,
    joiningCourse: false
  });

  const history = useHistory();

  const handleItemClick = (e, { name }) => setState({ activeItem: name });

  return (
    <div>
      <AddCourseModal
        addingCourse={state.addingCourse}
        setState={setState}
        makeNotif={props.makeNotif}
      />
      <JoinCourseModel
        joiningCourse={state.joiningCourse}
        setState={setState}
        username={props.username}
        makeNotif={props.makeNotif}
      />
      <Sidebar
        as={Menu}
        animation="overlay"
        icon="labeled"
        direction="left"
        vertical
        visible={props.isMobile ? props.open : true}
        width="thin"
        style={sidePanelLStyle}
      >
        {/* <Link to="/dashboard/panel"> */}
        <Menu.Item
          name="personal information"
          active={state.activeItem === "personal information"}
          // onClick={handleItemClick}
          onClick={() => {
            setState({ ...state, activeItem: "personal information" });
            history.push("/dashboard/panel");
          }}
          style={sidePanelItemLStyle}
        >
          <Icon inverted name="user" />
          Presonal Information
        </Menu.Item>
        {/* </Link> */}

        <Link to="/dashboard/courses">
          <Menu.Item
            name="classes"
            active={state.activeItem === "classes"}
            // onClick={handleItemClick}
            onClick={() => {
              history.push("/dashboard/courses");
              setState({ ...state, activeItem: "classes" });
            }}
            style={sidePanelItemLStyle}
          >
            <Icon inverted name="book" />
            Courses
          </Menu.Item>
        </Link>
        <Menu.Item
          as="a"
          positive
          onClick={() => {
            setState({ addingCourse: true });
          }}
          style={sidePanelItemLStyle}
        >
          <Icon inverted name="plus square" />
          Make A New Course
        </Menu.Item>
        <Menu.Item
          as="a"
          positive
          onClick={() => {
            setState({ joiningCourse: true });
          }}
          style={sidePanelItemLStyle}
        >
          <FontAwesomeIcon size="2x" icon={faSearchPlus} />
          <br />
          <br />
          Join Other Courses
        </Menu.Item>
      </Sidebar>
    </div>
  );
}

export default SideBar;

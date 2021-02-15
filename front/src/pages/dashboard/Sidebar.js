import React, { useState } from "react";
import { Sidebar, Menu, Button, Icon } from "semantic-ui-react";
import { useHistory, Link } from "react-router-dom";
import AddCourseModal from "./AddCourse.js";
import JoinCourseModel from "./JoinCourse.js";

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
      <AddCourseModal addingCourse={state.addingCourse} setState={setState} />
      <JoinCourseModel
        joiningCourse={state.joiningCourse}
        setState={setState}
      />
      <Sidebar
        as={Menu}
        animation="overlay"
        icon="labeled"
        direction="left"
        vertical
        visible={props.isMobile ? props.open : true}
        width="thin"
        style={{ width: 250, top: 70 }}
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
        >
          <Icon name="user" />
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
          >
            <Icon name="book" />
            Courses
          </Menu.Item>
        </Link>

        {/* 
        <Menu.Item
          name="personal information"
          active={state.activeItem === "personal information"}
          // onClick={handleItemClick}
          onClick={() => {
            history.push("/dashboard/panel");
          }}
        >
          <Icon name="user" />
          Presonal Information
        </Menu.Item>
        <Menu.Item
          name="classes"
          active={state.activeItem === "classes"}
          // onClick={handleItemClick}
          onClick={() => {
            history.push("/dashboard/courses");
          }}
        >
          <Icon name="book" />
          Courses
        </Menu.Item> */}
        <Menu.Item as="a">
          <Button
            positive
            onClick={() => {
              setState({ addingCourse: true });
            }}
          >
            Add New Class
          </Button>
        </Menu.Item>
        <Menu.Item as="a">
          <Button
            color="blue"
            onClick={() => {
              setState({ joiningCourse: true });
            }}
          >
            Join Other Classes
          </Button>
        </Menu.Item>
      </Sidebar>
    </div>
  );
}

export default SideBar;

import React, { useState } from "react";
import { Sidebar, Menu, Button, Icon } from "semantic-ui-react";

function SideBar(props) {
  const [state, setState] = useState({
    activeItem: "Personal Information"
  });

  const handleItemClick = (e, { name }) => setState({ activeItem: name });

  return (
    <Sidebar
      as={Menu}
      animation="overlay"
      icon="labeled"
      direction="left"
      vertical
      visible={props.isMobile ? props.sidebarIsOpen : true}
      width="thin"
      style={{ width: 250, top: 70 }}
    >
      <Menu.Item
        name="personal information"
        active={state.activeItem === "personal information"}
        onClick={handleItemClick}
      >
        <Icon name="user" />
        Presonal Information
      </Menu.Item>
      <Menu.Item
        name="classes"
        active={state.activeItem === "classes"}
        onClick={handleItemClick}
      >
        <Icon name="book" />
        Courses
      </Menu.Item>
      {/* <Menu.Item as="a">
        <Button
          positive
          onClick={() => {
            //   setState({ addingPost: true });
          }}
        >
          Add New Post
        </Button>
      </Menu.Item> */}
    </Sidebar>
  );
}

export default SideBar;

import React, { useState } from "react";
import { Sidebar, Menu, Button, Icon, List } from "semantic-ui-react";
import { useHistory } from "react-router-dom";

const TAs = [
  "folan1",
  "folan2",
  "folan3",
  "folan4",
  "folan5",
  "folan6",
  "folan7"
];

function SideBar(props) {
  const [state, setState] = useState({
    activeItem: "Personal Information"
  });

  const history = useHistory();

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
      <Menu.Item as="a">
        <Icon name="student" />
        Course Instructor: folan
      </Menu.Item>
      <Menu.Item as="a" >
        <Icon name="users" />
        TAs:
        <List>
          {TAs.map(TA => {
            return (
              <List.Item as="li">
                {/* <List.Icon name="user" /> */}
                {/* <List.Content>{TA}</List.Content> */}
                {TA}
              </List.Item>
            );
          })}
        </List>
      </Menu.Item>
    </Sidebar>
  );
}

export default SideBar;

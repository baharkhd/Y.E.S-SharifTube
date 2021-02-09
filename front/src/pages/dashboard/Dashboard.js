import React, { useState } from "react";
import { Sidebar, Menu, Button, Icon } from "semantic-ui-react";
import SideBar from "./Sidebar.js";
import Panel from "./Panel.js";
import Courses from "./Courses.js";

function Dashboard(props) {
  const [state, setState] = useState({
    activeItem: "Personal Information"
  });

  const handleItemClick = (e, { name }) => setState({ activeItem: name });

  return (
    <div>
      <SideBar />
      {/* <Panel /> */}
      <Courses />
    </div>
  );
}

export default Dashboard;

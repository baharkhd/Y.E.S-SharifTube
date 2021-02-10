import React, { useState } from "react";
import { Sidebar, Menu, Button, Icon } from "semantic-ui-react";
import SideBar from "./Sidebar.js";
import Panel from "./Panel.js";
import Courses from "./Courses.js";
import { Route, Switch, Link } from "react-router-dom";

function Dashboard(props) {
  const [state, setState] = useState({
    activeItem: "Personal Information"
  });

  const handleItemClick = (e, { name }) => setState({ activeItem: name });

  console.log("component:", props.component)
  const Component = props.component ? (
    props.component
  ) : (
    <Panel isMobile={props.isMobile} />
  );

  return (
    <div>
      <SideBar isMobile={props.isMobile} open={props.sidebarOpen} />
      {Component}
      {/* <Switch>
        <Route exact path="/dashboard/panel">
          <Panel isMobile={props.isMobile} />
        </Route>
        <Route exact path="/dashboard/courses">
          <Courses isMobile={props.isMobile} />
        </Route>
      </Switch> */}
    </div>
  );
}

export default Dashboard;

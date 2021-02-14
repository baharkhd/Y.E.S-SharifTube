import React, { useState } from "react";
import { Sidebar, Menu, Button, Icon } from "semantic-ui-react";
import SideBar from "./Sidebar.js";
import Panel from "./Panel.js";
import Courses from "./Courses.js";
import { Route, Switch, Link } from "react-router-dom";
import { gql, useQuery } from "@apollo/client";
import constants from "../../constants.js";

const GET_USER_QUERY = gql`
  {
    user {
      username
      name
      password
      email
      courseIDs
    }
  }
`;

function Dashboard(props) {
  const [state, setState] = useState({
    activeItem: "Personal Information"
  });

  console.log(
    "token in dashboard:",
    localStorage.getItem(constants.AUTH_TOKEN)
  );
  const { data, loading, error } = useQuery(GET_USER_QUERY);

  console.log("data:", data);
  console.log("loading:", loading);
  console.log("error:", error);

  return (
    <div>
      <SideBar isMobile={props.isMobile} open={props.sidebarOpen} />
      {!loading &&
        (!props.isCourse ? (
          <Panel isMobile={props.isMobile} user={data.user} />
        ) : (
          <Courses isMobile={props.isMobile} user={data.user} />
        ))}
    </div>
  );
}

export default Dashboard;

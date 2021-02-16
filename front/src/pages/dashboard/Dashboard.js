import React, { useState } from "react";
import { Sidebar, Menu, Button, Icon } from "semantic-ui-react";
import SideBar from "./Sidebar.js";
import Panel from "./Panel.js";
import Courses from "./Courses.js";
import { Route, Switch, Link } from "react-router-dom";
import { gql, useQuery, useMutation } from "@apollo/client";
import constants from "../../constants.js";

const GET_USER_QUERY = gql`
  {
    user {
      username
      name
      email
      courseIDs
    }
  }
`;

const Dashboard = props => {
  const [state, setState] = useState({
    activeItem: "Personal Information",
    user: undefined
  });

  console.log(
    "token in dashboard:",
    localStorage.getItem(constants.AUTH_TOKEN)
  );

  // const [test] = useMutation(GET_USER_QUERY, {
  //   refetchQueries
  // })
  const { data, loading, error } = useQuery(GET_USER_QUERY, {
    fetchPolicy: "cache-and-network",
    nextFetchPolicy: "cache-first"
  });

  console.log("user in dashboard:", state.user)
  console.log("data:", data);
  console.log("loading:", loading);
  console.log("error:", error);

  return (
    <div >
      <SideBar isMobile={props.isMobile} open={props.sidebarOpen} username={props.username} />
      {!loading &&
        (!props.isCourse ? (
          <Panel isMobile={props.isMobile} user={data.user} setState={setState} />
        ) : (
          <Courses isMobile={props.isMobile} user={data.user} />
        ))}
    </div>
  );
};

export default Dashboard;

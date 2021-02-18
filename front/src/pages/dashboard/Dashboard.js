import { gql, useQuery } from "@apollo/client";
import React, { useState } from "react";
import constants from "../../constants.js";
import Courses from "./Courses.js";
import Panel from "./Panel.js";
import SideBar from "./Sidebar.js";

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

  console.log("user in dashboard:", state.user);
  console.log("data:", data);
  console.log("loading:", loading);
  console.log("error:", error);

  return (
    <div>
      <SideBar
        isMobile={props.isMobile}
        open={props.sidebarOpen}
        username={props.username}
        makeNotif={props.makeNotif}
      />
      {!loading &&
        (!props.isCourse ? (
          <Panel
            isMobile={props.isMobile}
            user={data.user}
            setState={setState}
            makeNotif={props.makeNotif}
          />
        ) : (
          <Courses
            isMobile={props.isMobile}
            user={data.user}
            makeNotif={props.makeNotif}
          />
        ))}
    </div>
  );
};

export default Dashboard;

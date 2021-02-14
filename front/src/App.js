import "semantic-ui-css/semantic.min.css";
import React, { useState } from "react";
import { Button, Input, Segment } from "semantic-ui-react";
import { Route, Switch, useParams } from "react-router-dom";
import Header from "./pages/Header.js";
import Login from "./pages/Login/Login.js";
import Signup from "./pages/Signup/Signup.js";
import Homepage from "./pages/Homepage/Homepage.js";
import Dashboard from "./pages/dashboard/Dashboard.js";
import CourseDashboard from "./pages/CoursePage/CourseDashboard.js";
import ContentPage from "./pages/Content/ContentPage.js";
import { useMediaQuery } from "react-responsive";
import SearchTest from "./pages/SearchTest.js";
import SearchCourse from "./pages/Search/SearchCourse.js";
import SearchIndex from "./pages/Search/SearchIndex.js";
import Panel from "./pages/dashboard/Panel.js";
import Courses from "./pages/dashboard/Courses.js";
import PendingPage from "./pages/CoursePage/PendingPage.js";
import UploadPage from "./pages/CoursePage/UploadPage.js";
import useToken from "./Token/useToken.js";
import { gql, useQuery } from "@apollo/client";

const TestComponent = props => {
  let { id, test } = useParams();
  console.log("????", id.substring(1), test.substring(1));
  return (
    <div>
      <Segment>Test, {id}</Segment>
      <Segment>Test, {id}</Segment>
      <Segment>Test, {id}</Segment>
      <Segment>Test, {id}</Segment>
      <Segment>Test, {id}</Segment>
      <Segment>Test, {id}</Segment>
    </div>
  );
};

const GET_USER_QUERY = gql`
  {
    user {
      username
    }
  }
`;

function App() {
  const isMobile = useMediaQuery({
    query: "(max-device-width: 570px)"
  });

  const { data, loading, error } = useQuery(GET_USER_QUERY);
  console.log("checkkkkkkk:", data, loading, error);

  const { token, setToken } = useToken();
  console.log("token in app:", token);

  const [sidebarOpen, setSidebarOpen] = useState(false);

  return (
    <div className="App">
      <Header
        token={token}
        setToken={setToken}
        isMobile={isMobile}
        sidebarOpen={sidebarOpen}
        setSidebarOpen={setSidebarOpen}
      />
      <Switch>
        <Route exact path="/dashboard">
          <Dashboard isMobile={isMobile} sidebarOpen={sidebarOpen} />
        </Route>

        <Route exact path="/dashboard/panel">
          <Dashboard
            isMobile={isMobile}
            sidebarOpen={sidebarOpen}
            isCourse={false}
            // component={<Panel isMobile={isMobile} />}
          />
        </Route>

        <Route exact path="/dashboard/courses">
          <Dashboard
            isMobile={isMobile}
            sidebarOpen={sidebarOpen}
            isCourse={true}
            // component={<Courses isMobile={isMobile} />}
          />
        </Route>

        <Route exact path="/content">
          <ContentPage />
        </Route>
        {/* Todo: remove this part! */}
        <Route exact path="/course:id">
          <CourseDashboard
            isMobile={isMobile}
            sidebarOpen={sidebarOpen}
            username={token ? (!loading ? data.user.username : "") : ""}
          />
        </Route>
        <Route exact path="/">
          <Homepage />
        </Route>
        <Route exact path="/login">
          <Login setToken={setToken} />
        </Route>
        <Route exact path="/signup">
          <Signup />
        </Route>
        <Route exact path="/search">
          <SearchIndex />
        </Route>
        <Route exact path="/course:courseID/content:contentID">
          <ContentPage />
        </Route>
        <Route exact path="/course:courseID/pendings">
          <PendingPage />
        </Route>
        <Route exact path="/course:courseID/upload">
          <UploadPage />
        </Route>
      </Switch>
    </div>
  );
}

export default App;

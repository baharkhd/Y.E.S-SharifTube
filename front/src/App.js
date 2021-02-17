import "semantic-ui-css/semantic.min.css";
import React, { useState } from "react";
import { Button, Input, Segment } from "semantic-ui-react";
import { Route, Switch, useParams, Redirect } from "react-router-dom";
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

  const [username, setUsername] = useState("");

  const { data, loading, error } = useQuery(GET_USER_QUERY, {
    fetchPolicy: "cache-and-network"
    //   nextFetchPolicy: "cache-first"
  });
  console.log("checkkkkkkk:", data, loading, error);
  console.log("username:", username);
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
        {!token && <Redirect exact from="/dashboard" to="/login" />}
        {!token && <Redirect exact from="/dashboard/panel" to="/login" />}
        {!token && <Redirect exact from="/dashboard/courses" to="/login" />}
        {!token && <Redirect exact from="/content" to="/login" />}

        {!token && <Redirect exact from="/course:id" to="/login" />}
        {!token && (
          <Redirect
            exact
            from="/course:courseID/content:contentID"
            to="/login"
          />
        )}

        {!token && (
          <Redirect exact from="/course:courseID/pendings" to="/login" />
        )}

        {!token && (
          <Redirect exact from="/course:courseID/upload/video" to="/login" />
        )}
        {!token && (
          <Redirect
            exact
            from="/course:courseID/upload/attachment"
            to="/login"
          />
        )}
        {!token && (
          <Redirect exact from="/course:courseID/offer/video" to="/login" />
        )}
        {!token && (
          <Redirect
            exact
            from="/course:courseID/offer/attachment"
            to="/login"
          />
        )}

        {token && <Redirect exact from="/login" to="/dashboard" />}
        {token && <Redirect exact from="/signup" to="/dashboard" />}

        <Route exact path="/dashboard">
          {token && (
            <Dashboard
              isMobile={isMobile}
              sidebarOpen={sidebarOpen}
              username={
                token
                  ? !loading
                    ? data
                      ? data.user.username
                      : username
                    : ""
                  : ""
              }
            />
          )}
        </Route>

        <Route exact path="/dashboard/panel">
          {token && (
            <Dashboard
              isMobile={isMobile}
              sidebarOpen={sidebarOpen}
              isCourse={false}
              username={
                token
                  ? !loading
                    ? data
                      ? data.user.username
                      : username
                    : ""
                  : ""
              }
              // component={<Panel isMobile={isMobile} />}
            />
          )}
        </Route>

        <Route exact path="/dashboard/courses">
          {token && (
            <Dashboard
              isMobile={isMobile}
              sidebarOpen={sidebarOpen}
              isCourse={true}
              username={
                token
                  ? !loading
                    ? data
                      ? data.user.username
                      : username
                    : ""
                  : ""
              }
              // component={<Courses isMobile={isMobile} />}
            />
          )}
        </Route>

        <Route exact path="/content">
          {token && <ContentPage />}
        </Route>
        {/* Todo: remove this part! */}
        <Route exact path="/course:id">
          {token && (
            <CourseDashboard
              isMobile={isMobile}
              sidebarOpen={sidebarOpen}
              username={
                token
                  ? !loading
                    ? data
                      ? data.user.username
                      : username
                    : ""
                  : ""
              }
            />
          )}
        </Route>
        <Route exact path="/">
          <Homepage />
        </Route>
        <Route exact path="/login">
          {!token && <Login setToken={setToken} setUsername={setUsername} />}
        </Route>
        <Route exact path="/signup">
          {!token && <Signup setToken={setToken} setUsername={setUsername} />}
        </Route>
        {/* <Route exact path="/search">
          <SearchIndex />
        </Route> */}
        <Route exact path="/course:courseID/content:contentID">
          {token && <ContentPage />}
        </Route>
        <Route exact path="/course:courseID/pendings">
          {token && (
            <PendingPage
              username={
                token
                  ? !loading
                    ? data
                      ? data.user.username
                      : username
                    : ""
                  : ""
              }
            />
          )}
        </Route>
        <Route exact path="/course:courseID/upload/video">
          {token && <UploadPage fileType="video" />}
        </Route>
        <Route exact path="/course:courseID/upload/attachment">
          {token && <UploadPage fileType="attachment" />}
        </Route>
        <Route exact path="/course:courseID/offer/video">
          {token && <UploadPage />}
        </Route>
        <Route exact path="/course:courseID/offer/attachment">
          {token && <UploadPage />}
        </Route>
      </Switch>
    </div>
  );
}

export default App;

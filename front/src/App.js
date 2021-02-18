import { gql, useQuery } from "@apollo/client";
import React, { useState } from "react";
import ReactNotification, { store } from "react-notifications-component";
import "react-notifications-component/dist/theme.css";
import { useMediaQuery } from "react-responsive";
import { Redirect, Route, Switch } from "react-router-dom";
import "semantic-ui-css/semantic.min.css";
import ContentPage from "./pages/Content/ContentPage.js";
import CourseDashboard from "./pages/CoursePage/CourseDashboard.js";
import PendingPage from "./pages/CoursePage/PendingPage.js";
import UploadPage from "./pages/CoursePage/UploadPage.js";
import Dashboard from "./pages/dashboard/Dashboard.js";
import Header from "./pages/Header.js";
import Homepage from "./pages/Homepage/Homepage.js";
import Login from "./pages/Login/Login.js";
import Signup from "./pages/Signup/Signup.js";
import useToken from "./Token/useToken.js";

const GET_USER_QUERY = gql`
  {
    user {
      username
    }
  }
`;

function App() {
  const makeNotif = (title, error, type) => {
    store.addNotification({
      title: title,
      message: error,
      type: type,
      insert: "bottom",
      container: "bottom-right",
      animationIn: ["animate__animated", "animate__fadeIn"],
      animationOut: ["animate__animated", "animate__fadeOut"],
      dismiss: {
        duration: 1500,
        onScreen: false
      }
    });
  };

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
    <div className="App" style={{overflow: "hidden", height: "100%"}}>
      <ReactNotification isMobile={isMobile} />
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
              makeNotif={makeNotif}
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
              makeNotif={makeNotif}
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
              makeNotif={makeNotif}
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
          {token && <ContentPage makeNotif={makeNotif} />}
        </Route>
        {/* Todo: remove this part! */}
        <Route exact path="/course:id">
          {token && (
            <CourseDashboard
              makeNotif={makeNotif}
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
          <Homepage makeNotif={makeNotif} />
        </Route>
        <Route exact path="/login">
          {!token && (
            <Login
              makeNotif={makeNotif}
              setToken={setToken}
              setUsername={setUsername}
              makeNotif={makeNotif}
            />
          )}
        </Route>
        <Route exact path="/signup">
          {!token && (
            <Signup
              makeNotif={makeNotif}
              setToken={setToken}
              setUsername={setUsername}
            />
          )}
        </Route>
        {/* <Route exact path="/search">
          <SearchIndex />
        </Route> */}
        <Route exact path="/course:courseID/content:contentID">
          {token && <ContentPage makeNotif={makeNotif} />}
        </Route>
        <Route exact path="/course:courseID/pendings">
          {token && (
            <PendingPage
              makeNotif={makeNotif}
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
          {token && <UploadPage makeNotif={makeNotif} fileType="video" />}
        </Route>
        <Route exact path="/course:courseID/upload/attachment">
          {token && <UploadPage makeNotif={makeNotif} fileType="attachment" />}
        </Route>
        <Route exact path="/course:courseID/offer/video">
          {token && <UploadPage makeNotif={makeNotif} />}
        </Route>
        <Route exact path="/course:courseID/offer/attachment">
          {token && <UploadPage makeNotif={makeNotif} />}
        </Route>
      </Switch>
    </div>
  );
}

export default App;
